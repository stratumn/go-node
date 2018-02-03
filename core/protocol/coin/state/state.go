// Copyright © 2017-2018 Stratumn SAS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"encoding/binary"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	db "github.com/stratumn/alice/core/protocol/coin/db"
	pb "github.com/stratumn/alice/pb/coin"
)

var (
	// ErrAmountTooBig is returned when subtracting an amount greater than
	// the current balance.
	ErrAmountTooBig = errors.New("amount is too big")

	// ErrInvalidJobID is returned when a job ID is invalid.
	ErrInvalidJobID = errors.New("job ID is invalid")

	// ErrInconsistentTransactions is returned when the transactions to
	// roll back are inconsistent with the saved nonces.
	ErrInconsistentTransactions = errors.New("transactions are inconsistent")
)

// State stores users' account balances.
type State interface {
	Reader
	Writer
}

// Reader gives read access to users' account balances.
type Reader interface {
	// GetAccount gets the account details of a user identified
	// by his public key. It returns &pb.Account{} if the account is not
	// found.
	GetAccount(pubKey []byte) (*pb.Account, error)
}

// Writer gives write access to users' account balances.
type Writer interface {
	// UpdateAccount sets or updates the account of a user identified by
	// his public key. It should only be used for testing as it cannot be
	// rolled back.
	UpdateAccount(pubKey []byte, account *pb.Account) error

	// ProcessTransactions processes all the given transactions and updates
	// the state accordingly. It should be given a unique job ID, for
	// instance the hash of the block containing the transactions.
	ProcessTransactions(jobID []byte, txs []*pb.Transaction) error

	// RollbackTransactions rolls back transactions. The parameters are
	// expected to be identical to the ones that were given to the
	// corresponding call to ProcessTransactions().
	// You should only rollback the job corresponding to the current state.
	RollbackTransactions(jobID []byte, txs []*pb.Transaction) error
}

// NOTE: prefixes should have be the same bytesize to prevent unexpected
// behaviors when iterating over key prefixes. Smaller means less space used.
var (
	// accountPrefix is the prefix for account keys.
	accountPrefix = []byte{0}

	// prevNoncesPrefix is the prefix for previous nonces keys.
	prevNoncesPrefix = []byte{1}
)

type stateDB struct {
	mu sync.RWMutex
	db db.DB

	prefix []byte

	// jobIDLen is the bytesize of a job ID.
	//
	// It is required because to prevent collision of keys.
	//
	// For example:
	//
	//	jobA  + BC -> jobABC
	//	jobAB + C  -> jobABC
	jobIDSize int

	diff *db.Diff
}

// NewState creates a new state from a DB instance.
//
// Prefix is used to prefix keys in the database.
//
// VERY IMPORTANT NOTE: ALL THE DIFFERENT MODULES THAT SHARE THE SAME INSTANCE
// OF THE DATABASE MUST USE A UNIQUE PREFIX OF THE SAME BYTESIZE.
//
// All job IDs must have the given bytesize.
func NewState(database db.DB, prefix []byte, jobIDSize int) State {
	return &stateDB{
		db:        database,
		prefix:    prefix,
		jobIDSize: jobIDSize,
		diff:      db.NewDiff(database),
	}
}

func (s *stateDB) GetAccount(pubKey []byte) (*pb.Account, error) {
	return s.doGetAccount(s.db, pubKey)
}

// doGetAccount is able to get an account using anything that implements the
// db.Reader interface.
func (s *stateDB) doGetAccount(dbr db.Reader, pubKey []byte) (*pb.Account, error) {
	buf, err := dbr.Get(s.accountKey(pubKey))
	if errors.Cause(err) == db.ErrNotFound {
		return &pb.Account{}, nil
	}
	if err != nil {
		return nil, err
	}

	var account pb.Account
	err = proto.Unmarshal(buf, &account)

	return &account, errors.WithStack(err)
}

func (s *stateDB) UpdateAccount(pubKey []byte, account *pb.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.doUpdateAccount(s.db, pubKey, account)
}

// doGetAccount updates an account using anything that implements db.Writer.
func (s *stateDB) doUpdateAccount(dbw db.Writer, pubKey []byte, account *pb.Account) error {
	// Delete empty accounts to save space.
	if account.Balance == 0 && account.Nonce == 0 {
		return dbw.Delete(s.accountKey(pubKey))
	}

	buf, err := proto.Marshal(account)
	if err != nil {
		return errors.WithStack(err)
	}

	return dbw.Put(s.accountKey(pubKey), buf)
}

func (s *stateDB) ProcessTransactions(jobID []byte, txs []*pb.Transaction) error {
	if s.jobIDSize != len(jobID) {
		return ErrInvalidJobID
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	defer s.diff.Reset()

	// Eight bytes per transaction.
	nonces := make([]byte, len(txs)*8)

	for i, tx := range txs {
		from, err := s.doGetAccount(s.diff, tx.From)
		if err != nil {
			return err
		}

		binary.LittleEndian.PutUint64(nonces[i*8:], from.Nonce)

		// Subtract amount from sender and update nonce.
		if from.Balance < tx.Value {
			return ErrAmountTooBig
		}

		from.Balance -= tx.Value
		from.Nonce = tx.Nonce

		err = s.doUpdateAccount(s.diff, tx.From, from)
		if err != nil {
			return err
		}

		// Add amount to receiver.
		to, err := s.doGetAccount(s.diff, tx.To)
		if err != nil {
			return err
		}

		to.Balance += tx.Value

		err = s.doUpdateAccount(s.diff, tx.To, to)
		if err != nil {
			return err
		}

	}

	if err := s.diff.Put(s.prevNoncesKey(jobID), nonces); err != nil {
		return err
	}

	return s.diff.Apply()
}

func (s *stateDB) RollbackTransactions(jobID []byte, txs []*pb.Transaction) error {
	if s.jobIDSize != len(jobID) {
		return ErrInvalidJobID
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	defer s.diff.Reset()

	nonces, err := s.db.Get(s.prevNoncesKey(jobID))
	if err != nil {
		return err
	}

	if len(nonces) != len(txs)*8 {
		return ErrInconsistentTransactions
	}

	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i]

		// Add amount to sender and restore nonce.
		from, err := s.doGetAccount(s.diff, tx.From)
		if err != nil {
			return nil
		}

		from.Balance += tx.Value
		from.Nonce = binary.LittleEndian.Uint64(nonces[i*8:])

		err = s.doUpdateAccount(s.diff, tx.From, from)
		if err != nil {
			return err
		}

		// Subtract amount from received.
		to, err := s.doGetAccount(s.diff, tx.To)
		if err != nil {
			return nil
		}

		to.Balance -= tx.Value

		err = s.doUpdateAccount(s.diff, tx.To, to)
		if err != nil {
			return err
		}
	}

	if err := s.diff.Delete(s.prevNoncesKey(jobID)); err != nil {
		return err
	}

	return s.diff.Apply()
}

// accountKey returns the key corresponding to an account given its public key.
func (s *stateDB) accountKey(pubKey []byte) []byte {
	return append(append(s.prefix, accountPrefix...), pubKey...)
}

// prevNoncesKey returns the previous nonces key for the given job ID.
func (s *stateDB) prevNoncesKey(jobID []byte) []byte {
	return append(append(s.prefix, prevNoncesPrefix...), jobID...)
}
