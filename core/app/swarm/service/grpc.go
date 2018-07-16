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

package service

import (
	"context"

	"github.com/pkg/errors"

	pb "github.com/stratumn/go-indigonode/core/app/swarm/grpc"

	swarm "gx/ipfs/QmRqfgh56f8CrqpwH7D2s6t8zQRsvPoftT3sp5Y6SUhNA3/go-libp2p-swarm"
	peer "gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	"gx/ipfs/QmdeiKhUy1TVGBaKxt7y1QmBDLBdisSrLJ1x58Eoj4PXUh/go-libp2p-peerstore"
)

// grpcServer is a gRPC server for the swarm service.
type grpcServer struct {
	GetSwarm     func() *swarm.Swarm
	GetPeerStore func() peerstore.Peerstore
}

// LocalPeer returns the local peer.
func (s grpcServer) LocalPeer(ctx context.Context, req *pb.LocalPeerReq) (*pb.Peer, error) {
	swm := s.GetSwarm()
	if swm == nil {
		return nil, errors.WithStack(ErrUnavailable)
	}

	return &pb.Peer{Id: []byte(swm.LocalPeer())}, nil
}

// Peers lists the peers connected to the local peer.
func (s grpcServer) Peers(req *pb.PeersReq, ss pb.Swarm_PeersServer) error {
	swm := s.GetSwarm()
	if swm == nil {
		return errors.WithStack(ErrUnavailable)
	}

	for _, pid := range swm.Peers() {
		if err := ss.Send(&pb.Peer{Id: []byte(pid)}); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// Peers lists the connections of the swarm.
func (s grpcServer) Connections(req *pb.ConnectionsReq, ss pb.Swarm_ConnectionsServer) error {
	swm := s.GetSwarm()
	if swm == nil {
		return errors.WithStack(ErrUnavailable)
	}

	var conns []*swarm.Conn

	if len(req.PeerId) < 1 {
		conns = swm.Connections()
	} else {
		pid, err := peer.IDFromBytes(req.PeerId)
		if err != nil {
			return errors.WithStack(err)
		}

		conns = swm.ConnectionsToPeer(pid)
	}

	for _, conn := range conns {
		err := ss.Send(&pb.Connection{
			PeerId:        []byte(conn.RemotePeer()),
			LocalAddress:  conn.LocalMultiaddr().Bytes(),
			RemoteAddress: conn.RemoteMultiaddr().Bytes(),
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// Addressses lists a peer's known addresses.
func (s grpcServer) Addresses(req *pb.PeerAddrsReq, ss pb.Swarm_AddressesServer) error {
	peerStore := s.GetPeerStore()
	if peerStore == nil {
		return errors.WithStack(ErrUnavailable)
	}

	pid, err := peer.IDFromBytes(req.PeerId)
	if err != nil {
		return errors.WithStack(err)
	}

	addrs := peerStore.Addrs(pid)
	for _, addr := range addrs {
		err := ss.Send(&pb.PeerAddr{
			Address: addr.Bytes(),
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
