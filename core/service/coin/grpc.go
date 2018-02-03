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

package coin

import (
	"context"
	"crypto/sha256"
	"encoding/json"

	"github.com/pkg/errors"
	rpcpb "github.com/stratumn/alice/grpc/coin"
	pb "github.com/stratumn/alice/pb/coin"

	peer "gx/ipfs/Qma7H6RW8wRrfZpNSXwxYGcd1E149s42FpWNpDNieSVrnU/go-libp2p-peer"
)

// grpcServer is a gRPC server for the coin service.
type grpcServer struct {
	GetAccount     func([]byte) (*pb.Account, error)
	AddTransaction func(*pb.Transaction) error
}

// Account returns an account.
func (s grpcServer) Account(ctx context.Context, req *rpcpb.AccountReq) (*pb.Account, error) {
	// Just for validation.
	_, err := peer.IDFromBytes(req.PeerId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return s.GetAccount(req.PeerId)
}

// Transaction sends a coin transaction to the consensus engine.
func (s grpcServer) Transaction(ctx context.Context, req *pb.Transaction) (response *rpcpb.TransactionResp, err error) {
	err = s.AddTransaction(req)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	txHash := sha256.Sum256(b)
	txResponse := &rpcpb.TransactionResp{
		TxHash: txHash[:],
	}

	return txResponse, nil
}
