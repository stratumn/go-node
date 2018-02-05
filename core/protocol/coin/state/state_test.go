// Copyright © 2017-2018  Stratumn SAS
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
	"testing"

	db "github.com/stratumn/alice/core/protocol/coin/db"
	pb "github.com/stratumn/alice/pb/coin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestState(t *testing.T) {
	alice := []byte("alice")
	bob := []byte("bob")
	charlie := []byte("charlie")

	tests := []struct {
		name string
		run  func(*testing.T, State)
	}{{
		"set-get",
		func(t *testing.T, s State) {
			err := s.UpdateAccount(alice, &pb.Account{Balance: 10, Nonce: 1})
			assert.NoError(t, err, "s.UpdateAccount()")
			v, err := s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount()")
			assert.Equal(t, &pb.Account{Balance: 10, Nonce: 1}, v)

			err = s.UpdateAccount(alice, &pb.Account{Balance: 20})
			assert.NoError(t, err, "s.UpdateAccount()")
			v, err = s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount()")
			assert.Equal(t, &pb.Account{Balance: 20}, v)
		},
	}, {
		"get-inexisting",
		func(t *testing.T, s State) {
			v, err := s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount()")
			assert.Equal(t, &pb.Account{}, v)
		},
	}, {
		"process-transactions",
		func(t *testing.T, s State) {
			err := s.UpdateAccount(alice, &pb.Account{Balance: 20})
			assert.NoError(t, err, "s.UpdateAccount()")

			txs := []*pb.Transaction{{
				From:  alice,
				To:    bob,
				Value: 10,
				Nonce: 1,
			}, {
				From:  bob,
				To:    charlie,
				Value: 5,
				Nonce: 2,
			}, {
				From:  charlie,
				To:    alice,
				Value: 2,
				Nonce: 3,
			}}

			err = s.ProcessTransactions([]byte("state1"), txs)
			assert.NoError(t, err)

			v, err := s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount(alice)")
			assert.Equal(t, &pb.Account{Balance: 20 - 10 + 2, Nonce: 1}, v, "s.GetAccount(alice)")

			v, err = s.GetAccount(bob)
			assert.NoError(t, err, "s.GetAccount(bob)")
			assert.Equal(t, &pb.Account{Balance: 10 - 5, Nonce: 2}, v, "s.GetAccount(bob)")

			v, err = s.GetAccount(charlie)
			assert.NoError(t, err, "s.GetAccount(charlie)")
			assert.Equal(t, &pb.Account{Balance: 5 - 2, Nonce: 3}, v, "s.GetAccount(charlie)")
		},
	}, {
		"process-transactions-invalid-state-id",
		func(t *testing.T, s State) {
			err := s.ProcessTransactions([]byte("state10"), nil)
			assert.EqualError(t, err, ErrInvalidStateID.Error(), "s.ProcessTransactions(state10)")
		},
	}, {
		"rollback-transactions",
		func(t *testing.T, s State) {
			err := s.UpdateAccount(alice, &pb.Account{Balance: 20})
			assert.NoError(t, err, "s.UpdateAccount()")

			// Process two states.

			txs1 := []*pb.Transaction{{
				From:  alice,
				To:    bob,
				Value: 10,
				Nonce: 1,
			}, {
				From:  bob,
				To:    charlie,
				Value: 5,
				Nonce: 2,
			}, {
				From:  charlie,
				To:    alice,
				Value: 2,
				Nonce: 3,
			}}

			txs2 := []*pb.Transaction{{
				From:  bob,
				To:    charlie,
				Value: 4,
				Nonce: 3,
			}, {
				From:  alice,
				To:    bob,
				Value: 3,
				Nonce: 4,
			}}

			err = s.ProcessTransactions([]byte("state1"), txs1)
			assert.NoError(t, err, "s.ProcessTransactions(state1)")

			err = s.ProcessTransactions([]byte("state2"), txs2)
			assert.NoError(t, err, "s.ProcessTransactions(state2)")

			// Rollback state2.

			err = s.RollbackTransactions([]byte("state2"), txs2)
			assert.NoError(t, err, "s.RollbackTransactions(state2)")

			v, err := s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount(alice)")
			assert.Equal(t, &pb.Account{Balance: 20 - 10 + 2, Nonce: 1}, v, "s.GetAccount(alice)")

			v, err = s.GetAccount(bob)
			assert.NoError(t, err, "s.GetAccount(bob)")
			assert.Equal(t, &pb.Account{Balance: 10 - 5, Nonce: 2}, v, "s.GetAccount(bob)")

			v, err = s.GetAccount(charlie)
			assert.NoError(t, err, "s.GetAccount(charlie)")
			assert.Equal(t, &pb.Account{Balance: 5 - 2, Nonce: 3}, v, "s.GetAccount(charlie)")

			// Rollback state1.

			err = s.RollbackTransactions([]byte("state1"), txs1)
			assert.NoError(t, err, "s.RollbackTransactions(state1)")

			v, err = s.GetAccount(alice)
			assert.NoError(t, err, "s.GetAccount(alice)")
			assert.Equal(t, &pb.Account{Balance: 20}, v, "s.GetAccount(alice)")

			v, err = s.GetAccount(bob)
			assert.NoError(t, err, "s.GetAccount(bob)")
			assert.Equal(t, &pb.Account{}, v, "s.GetAccount(bob)")

			v, err = s.GetAccount(charlie)
			assert.NoError(t, err, "s.GetAccount(charlie)")
			assert.Equal(t, &pb.Account{}, v, "s.GetAccount(charlie)")
		},
	}, {
		"rollback-transactions-invalid-state-id",
		func(t *testing.T, s State) {
			err := s.RollbackTransactions([]byte("state10"), nil)
			assert.EqualError(t, err, ErrInvalidStateID.Error(), "s.ProcessTransactions(state10)")
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memdb, err := db.NewMemDB(nil)
			require.NoError(t, err, "db.NewMemDB()")
			defer memdb.Close()

			tt.run(t, NewState(memdb, OptPrefix([]byte("test-")), OptStateIDLen(6)))
		})
	}
}
