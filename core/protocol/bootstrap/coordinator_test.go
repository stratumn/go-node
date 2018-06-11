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

package bootstrap_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stratumn/alice/core/protector"
	"github.com/stratumn/alice/core/protocol/bootstrap"
	"github.com/stratumn/alice/core/protocol/bootstrap/proposal"
	"github.com/stratumn/alice/core/protocol/bootstrap/proposal/mocks"
	pb "github.com/stratumn/alice/pb/bootstrap"
	protectorpb "github.com/stratumn/alice/pb/protector"
	"github.com/stratumn/alice/test"
	"github.com/stratumn/alice/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	protobuf "gx/ipfs/QmRDePEiL4Yupq5EkcK3L3ko3iMgYaqUdLu7xc1kqs7dnV/go-multicodec/protobuf"
	"gx/ipfs/QmVxf27kucSvCLiCq6dAXjDU2WG3xZN9ae7Ny6osroP28u/yamux"
	"gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	inet "gx/ipfs/QmXoz9o2PT3tEzf7hicegwex5UgVP54n3k82K7jrWFyN86/go-libp2p-net"
	netutil "gx/ipfs/Qmb6BsZf6Y3kxffXMNTubGPF1w1bkHtpvhfYbmnwP3NQyw/go-libp2p-netutil"
	bhost "gx/ipfs/Qmc64U41EEB4nPG7wxjEqFwKJajS2f8kk5q2TvUrQf78Xu/go-libp2p-blankhost"
	"gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	"gx/ipfs/QmdeiKhUy1TVGBaKxt7y1QmBDLBdisSrLJ1x58Eoj4PXUh/go-libp2p-peerstore"
	ihost "gx/ipfs/QmfZTdmunzKzAGJrSvXXQbQ5kLLUiEMX5vdwux7iXkdk7D/go-libp2p-host"
)

func TestCoordinator_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mocks.NewMockHost(ctrl)
	host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorHandshakePID, gomock.Any()).Times(1)
	host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorProposePID, gomock.Any()).Times(1)

	handler, err := bootstrap.NewCoordinatorHandler(host, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, handler)

	host.EXPECT().RemoveStreamHandler(bootstrap.PrivateCoordinatorHandshakePID).Times(1)
	host.EXPECT().RemoveStreamHandler(bootstrap.PrivateCoordinatorProposePID).Times(1)
	handler.Close(context.Background())
}

func TestCoordinator_HandleHandshake(t *testing.T) {
	sendHello := func(t *testing.T, stream inet.Stream) {
		enc := protobuf.Multicodec(nil).Encoder(stream)
		require.NoError(t, enc.Encode(&pb.Hello{}), "enc.Encode()")
	}

	testCases := []struct {
		name          string
		networkConfig func(context.Context, *bhost.BlankHost, *bhost.BlankHost) protector.NetworkConfig
		send          func(*testing.T, inet.Stream)
		validate      func(*testing.T, *protectorpb.NetworkConfig, *bhost.BlankHost, *bhost.BlankHost)
		receiveErr    error
	}{{
		"during-bootstrap-send-participants-to-white-listed-peer",
		func(ctx context.Context, coordinator, sender *bhost.BlankHost) protector.NetworkConfig {
			networkConfig, _ := protector.NewInMemoryConfig(
				ctx,
				protectorpb.NewNetworkConfig(protectorpb.NetworkState_BOOTSTRAP),
			)

			networkConfig.AddPeer(
				ctx,
				coordinator.ID(),
				test.GeneratePeerMultiaddrs(t, coordinator.ID()),
			)
			networkConfig.AddPeer(
				ctx,
				sender.ID(),
				test.GeneratePeerMultiaddrs(t, sender.ID()),
			)

			return networkConfig
		},
		sendHello,
		func(t *testing.T, networkConfig *protectorpb.NetworkConfig, coordinator, sender *bhost.BlankHost) {
			assert.Equal(t, protectorpb.NetworkState_BOOTSTRAP, networkConfig.NetworkState)
			assert.Len(t, networkConfig.Participants, 2)
			assert.Contains(t, networkConfig.Participants, sender.ID().Pretty())
			assert.Contains(t, networkConfig.Participants, coordinator.ID().Pretty())
		},
		nil,
	}, {
		"during-bootstrap-do-not-send-participants-to-non-white-listed-peer",
		func(ctx context.Context, coordinator, sender *bhost.BlankHost) protector.NetworkConfig {
			networkConfig, _ := protector.NewInMemoryConfig(
				ctx,
				protectorpb.NewNetworkConfig(protectorpb.NetworkState_BOOTSTRAP),
			)

			networkConfig.AddPeer(
				ctx,
				coordinator.ID(),
				test.GeneratePeerMultiaddrs(t, coordinator.ID()),
			)

			return networkConfig
		},
		sendHello,
		func(t *testing.T, networkConfig *protectorpb.NetworkConfig, coordinator, sender *bhost.BlankHost) {
			assert.Nil(t, networkConfig.Signature, "networkConfig.Signature")
			assert.Zero(t, networkConfig.NetworkState, "networkConfig.NetworkState")
			assert.Nil(t, networkConfig.Participants, "networkConfig.Participants")
		},
		nil,
	}, {
		"after-bootstrap-send-participants-to-white-listed-peer",
		func(ctx context.Context, coordinator, sender *bhost.BlankHost) protector.NetworkConfig {
			networkConfig, _ := protector.NewInMemoryConfig(
				ctx,
				protectorpb.NewNetworkConfig(protectorpb.NetworkState_PROTECTED),
			)

			networkConfig.AddPeer(
				ctx,
				sender.ID(),
				test.GeneratePeerMultiaddrs(t, sender.ID()),
			)

			return networkConfig
		},
		sendHello,
		func(t *testing.T, networkConfig *protectorpb.NetworkConfig, coordinator, sender *bhost.BlankHost) {
			assert.Equal(t, protectorpb.NetworkState_PROTECTED, networkConfig.NetworkState)
			assert.Len(t, networkConfig.Participants, 1)
			assert.Contains(t, networkConfig.Participants, sender.ID().Pretty())
		},
		nil,
	}, {
		"after-bootstrap-reject-non-white-listed-peer",
		func(ctx context.Context, coordinator, sender *bhost.BlankHost) protector.NetworkConfig {
			networkConfig, _ := protector.NewInMemoryConfig(
				ctx,
				protectorpb.NewNetworkConfig(protectorpb.NetworkState_PROTECTED),
			)

			networkConfig.AddPeer(
				ctx,
				coordinator.ID(),
				test.GeneratePeerMultiaddrs(t, coordinator.ID()),
			)

			return networkConfig
		},
		func(t *testing.T, stream inet.Stream) {
			enc := protobuf.Multicodec(nil).Encoder(stream)
			err := enc.Encode(&pb.Hello{})
			if err != nil {
				assert.EqualError(t, err, yamux.ErrConnectionReset.Error())
			}
		},
		nil,
		yamux.ErrConnectionReset,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			coordinator := bhost.NewBlankHost(netutil.GenSwarmNetwork(t, ctx))
			defer coordinator.Close()

			sender := bhost.NewBlankHost(netutil.GenSwarmNetwork(t, ctx))
			defer sender.Close()

			require.NoError(t, sender.Connect(ctx, coordinator.Peerstore().PeerInfo(coordinator.ID())))

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler, err := bootstrap.NewCoordinatorHandler(
				coordinator,
				tt.networkConfig(ctx, coordinator, sender),
				nil,
			)
			require.NoError(t, err)
			defer handler.Close(ctx)

			stream, err := sender.NewStream(ctx, coordinator.ID(), bootstrap.PrivateCoordinatorHandshakePID)
			// If the coordinator is expected to reject streams, it can happen either
			// when initiating the stream (sender.NewStream()) or when writing to it (below).
			// It depends on the underlying yamux implementation but both are ok for our usecase.
			if err != nil {
				require.Error(t, tt.receiveErr)
				require.EqualError(t, err, tt.receiveErr.Error())
				return
			}

			tt.send(t, stream)

			dec := protobuf.Multicodec(nil).Decoder(stream)
			var response protectorpb.NetworkConfig
			err = dec.Decode(&response)

			if tt.receiveErr != nil {
				assert.EqualError(t, err, tt.receiveErr.Error())
			} else {
				require.NoError(t, err)
				tt.validate(t, &response, coordinator, sender)
			}
		})
	}
}

func TestCoordinator_HandlePropose(t *testing.T) {
	peer1 := test.GeneratePeerID(t)
	peer1Addr := test.GeneratePeerMultiaddr(t, peer1)

	var coordinator ihost.Host
	var sender ihost.Host

	testCases := []struct {
		name        string
		configure   func(protector.NetworkConfig)
		expectStore func(*testing.T, *mockproposal.MockStore)
		proposal    func() *pb.NodeIdentity
		validate    func(*testing.T, *pb.Ack)
	}{{
		"during-bootstrap-invalid-peer-id",
		func(protector.NetworkConfig) {},
		func(*testing.T, *mockproposal.MockStore) {},
		func() *pb.NodeIdentity { return &pb.NodeIdentity{PeerId: []byte("b4tm4n")} },
		func(t *testing.T, ack *pb.Ack) {
			assert.Equal(t, proposal.ErrInvalidPeerID.Error(), ack.Error)
		},
	}, {
		"during-bootstrap-missing-peer-addr",
		func(protector.NetworkConfig) {},
		func(*testing.T, *mockproposal.MockStore) {},
		func() *pb.NodeIdentity {
			return &pb.NodeIdentity{
				PeerId: []byte(peer1),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Equal(t, proposal.ErrMissingPeerAddr.Error(), ack.Error)
		},
	}, {
		"during-bootstrap-node-mismatch",
		func(protector.NetworkConfig) {},
		func(*testing.T, *mockproposal.MockStore) {},
		func() *pb.NodeIdentity {
			// Peer1 is different from the sender. During bootstrap,
			// nodes can only add themselves (not another node).
			return &pb.NodeIdentity{
				PeerId:   []byte(peer1),
				PeerAddr: peer1Addr.Bytes(),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Equal(t, proposal.ErrInvalidPeerAddr.Error(), ack.Error)
		},
	}, {
		"during-bootstrap-addr-in-peerstore",
		func(protector.NetworkConfig) {},
		func(t *testing.T, store *mockproposal.MockStore) {
			store.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1).Do(
				func(ctx context.Context, r *proposal.Request) {
					assert.Equal(t, proposal.AddNode, r.Type)
					assert.Equal(t, sender.ID(), r.PeerID)
					assert.Equal(t, sender.Addrs()[0], r.PeerAddr)
				})
		},
		func() *pb.NodeIdentity {
			coordinator.Peerstore().AddAddrs(
				sender.ID(),
				sender.Addrs(),
				peerstore.PermanentAddrTTL,
			)

			return &pb.NodeIdentity{
				PeerId: []byte(sender.ID()),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Zero(t, ack.Error)
		},
	}, {
		"during-bootstrap-addr-provided",
		func(protector.NetworkConfig) {},
		func(t *testing.T, store *mockproposal.MockStore) {
			store.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1).Do(
				func(ctx context.Context, r *proposal.Request) {
					assert.Equal(t, proposal.AddNode, r.Type)
					assert.Equal(t, sender.ID(), r.PeerID)
					assert.Equal(t, sender.Addrs()[0], r.PeerAddr)
				})
		},
		func() *pb.NodeIdentity {
			return &pb.NodeIdentity{
				PeerId:   []byte(sender.ID()),
				PeerAddr: sender.Addrs()[0].Bytes(),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Zero(t, ack.Error)
		},
	}, {
		"after-bootstrap-add-node",
		func(networkConfig protector.NetworkConfig) {
			ctx := context.Background()
			err := networkConfig.SetNetworkState(ctx, protectorpb.NetworkState_PROTECTED)
			require.NoError(t, err, "networkConfig.SetNetworkState()")

			err = networkConfig.AddPeer(ctx, sender.ID(), sender.Addrs())
			require.NoError(t, err, "networkConfig.AddPeer()")
		},
		func(t *testing.T, store *mockproposal.MockStore) {
			store.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1).Do(
				func(ctx context.Context, r *proposal.Request) {
					assert.Equal(t, proposal.AddNode, r.Type)
					assert.Equal(t, peer1, r.PeerID)
					assert.Equal(t, peer1Addr, r.PeerAddr)
				})
		},
		func() *pb.NodeIdentity {
			return &pb.NodeIdentity{
				PeerId:   []byte(peer1),
				PeerAddr: peer1Addr.Bytes(),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Zero(t, ack.Error)
		},
	}, {
		"after-bootstrap-remove-node",
		func(networkConfig protector.NetworkConfig) {
			ctx := context.Background()
			err := networkConfig.SetNetworkState(ctx, protectorpb.NetworkState_PROTECTED)
			require.NoError(t, err, "networkConfig.SetNetworkState()")

			err = networkConfig.AddPeer(ctx, sender.ID(), sender.Addrs())
			require.NoError(t, err, "networkConfig.AddPeer()")

			err = networkConfig.AddPeer(ctx, peer1, []multiaddr.Multiaddr{peer1Addr})
			require.NoError(t, err, "networkConfig.AddPeer()")
		},
		func(t *testing.T, store *mockproposal.MockStore) {
			store.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1).Do(
				func(ctx context.Context, r *proposal.Request) {
					assert.Equal(t, proposal.RemoveNode, r.Type)
					assert.Equal(t, peer1, r.PeerID)
				})
		},
		func() *pb.NodeIdentity {
			return &pb.NodeIdentity{
				PeerId: []byte(peer1),
			}
		},
		func(t *testing.T, ack *pb.Ack) {
			assert.Zero(t, ack.Error)
		},
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			coordinator = bhost.NewBlankHost(netutil.GenSwarmNetwork(t, ctx))
			defer coordinator.Close()

			sender = bhost.NewBlankHost(netutil.GenSwarmNetwork(t, ctx))
			defer sender.Close()

			require.NoError(t, sender.Connect(ctx, coordinator.Peerstore().PeerInfo(coordinator.ID())))

			store := mockproposal.NewMockStore(ctrl)
			tt.expectStore(t, store)

			networkConfig, _ := protector.NewInMemoryConfig(
				ctx,
				protectorpb.NewNetworkConfig(protectorpb.NetworkState_BOOTSTRAP),
			)
			tt.configure(networkConfig)

			handler, err := bootstrap.NewCoordinatorHandler(
				coordinator,
				networkConfig,
				store,
			)
			require.NoError(t, err)
			defer handler.Close(ctx)

			stream, err := sender.NewStream(ctx, coordinator.ID(), bootstrap.PrivateCoordinatorProposePID)
			require.NoError(t, err, "sender.NewStream()")

			enc := protobuf.Multicodec(nil).Encoder(stream)
			err = enc.Encode(tt.proposal())
			require.NoError(t, err, "enc.Encode()")

			dec := protobuf.Multicodec(nil).Decoder(stream)
			var ack pb.Ack
			err = dec.Decode(&ack)
			require.NoError(t, err, "dec.Decode()")
			tt.validate(t, &ack)
		})
	}
}

func TestCoordinator_AddNode(t *testing.T) {
	coordinatorID := test.GeneratePeerID(t)
	peer1 := test.GeneratePeerID(t)
	peer1Addrs := []multiaddr.Multiaddr{test.GeneratePeerMultiaddr(t, peer1)}
	peer2 := test.GeneratePeerID(t)

	testCases := []struct {
		name                   string
		addNodeID              peer.ID
		addNodeAddr            multiaddr.Multiaddr
		configureHost          func(*gomock.Controller, *mocks.MockHost)
		configureNetworkConfig func(*mocks.MockNetworkConfig)
		err                    error
	}{{
		"node-addr-missing",
		peer1,
		nil,
		func(_ *gomock.Controller, host *mocks.MockHost) {
			// If the peer store doesn't have an address for the node,
			// and it wasn't provided we reject the request.
			peerStore := peerstore.NewPeerstore()
			host.EXPECT().Peerstore().Times(1).Return(peerStore)
		},
		func(networkConfig *mocks.MockNetworkConfig) {
			networkConfig.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(false)
		},
		bootstrap.ErrUnknownNode,
	}, {
		"node-already-white-listed",
		peer1,
		nil,
		func(*gomock.Controller, *mocks.MockHost) {
			// If the node is already white-listed we shouldn't notify
			// participants, so we shouldn't use the host.
		},
		func(networkConfig *mocks.MockNetworkConfig) {
			networkConfig.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(true)
		},
		nil,
	}, {
		"new-node-added-from-peerstore",
		peer1,
		nil,
		func(ctrl *gomock.Controller, host *mocks.MockHost) {
			peerStore := peerstore.NewPeerstore()
			peerStore.AddAddrs(peer1, peer1Addrs, peerstore.PermanentAddrTTL)
			host.EXPECT().Peerstore().Times(1).Return(peerStore)

			stream := mocks.NewMockStream(ctrl)
			stream.EXPECT().Write(gomock.Any()).AnyTimes()
			stream.EXPECT().Close().Times(2)

			host.EXPECT().NewStream(gomock.Any(), peer1, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
			host.EXPECT().NewStream(gomock.Any(), peer2, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
		},
		func(networkConfig *mocks.MockNetworkConfig) {
			networkConfig.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(false)
			networkConfig.EXPECT().AddPeer(gomock.Any(), peer1, peer1Addrs).Times(1)
			networkConfig.EXPECT().AllowedPeers(gomock.Any()).Times(1).Return([]peer.ID{peer1, peer2})
			networkConfig.EXPECT().Copy(gomock.Any()).Times(1)
		},
		nil,
	}, {
		"new-node-added-from-addr",
		peer1,
		peer1Addrs[0],
		func(ctrl *gomock.Controller, host *mocks.MockHost) {
			peerStore := peerstore.NewPeerstore()
			host.EXPECT().Peerstore().Times(1).Return(peerStore)

			stream := mocks.NewMockStream(ctrl)
			stream.EXPECT().Write(gomock.Any()).AnyTimes()
			stream.EXPECT().Close().Times(2)

			host.EXPECT().NewStream(gomock.Any(), peer1, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
			host.EXPECT().NewStream(gomock.Any(), peer2, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
		},
		func(networkConfig *mocks.MockNetworkConfig) {
			networkConfig.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(false)
			networkConfig.EXPECT().AddPeer(gomock.Any(), peer1, peer1Addrs).Times(1)
			networkConfig.EXPECT().AllowedPeers(gomock.Any()).Times(1).Return([]peer.ID{peer1, peer2})
			networkConfig.EXPECT().Copy(gomock.Any()).Times(1)
		},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			host := mocks.NewMockHost(ctrl)
			host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorHandshakePID, gomock.Any())
			host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorProposePID, gomock.Any())
			host.EXPECT().ID().AnyTimes().Return(coordinatorID)
			tt.configureHost(ctrl, host)

			networkConfig := mocks.NewMockNetworkConfig(ctrl)
			tt.configureNetworkConfig(networkConfig)

			handler, err := bootstrap.NewCoordinatorHandler(host, networkConfig, nil)
			require.NoError(t, err, "bootstrap.NewCoordinatorHandler()")

			err = handler.AddNode(ctx, tt.addNodeID, tt.addNodeAddr, []byte("I'm batman"))
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoordinator_Accept(t *testing.T) {
	coordinatorID := test.GeneratePeerID(t)
	peer1 := test.GeneratePeerID(t)
	peer1Addr := test.GeneratePeerMultiaddr(t, peer1)
	peer2 := test.GeneratePeerID(t)

	addPeer1 := &proposal.Request{
		Type:     proposal.AddNode,
		PeerID:   peer1,
		PeerAddr: peer1Addr,
	}

	testCases := []struct {
		name      string
		acceptID  peer.ID
		configure func(*gomock.Controller, *mocks.MockHost, *mocks.MockNetworkConfig, *mockproposal.MockStore)
		err       error
	}{{
		"proposal-missing",
		peer1,
		func(ctrl *gomock.Controller, host *mocks.MockHost, cfg *mocks.MockNetworkConfig, store *mockproposal.MockStore) {
			store.EXPECT().Get(gomock.Any(), peer1).Times(1).Return(nil, nil)
		},
		proposal.ErrMissingRequest,
	}, {
		"peer-addr-missing",
		peer1,
		func(ctrl *gomock.Controller, host *mocks.MockHost, cfg *mocks.MockNetworkConfig, store *mockproposal.MockStore) {
			r := &proposal.Request{
				Type:   proposal.AddNode,
				PeerID: peer1,
			}
			store.EXPECT().Get(gomock.Any(), peer1).Times(1).Return(r, nil)
			store.EXPECT().Remove(gomock.Any(), peer1).Times(1)
		},
		proposal.ErrMissingPeerAddr,
	}, {
		"add-already-added",
		peer1,
		func(ctrl *gomock.Controller, host *mocks.MockHost, cfg *mocks.MockNetworkConfig, store *mockproposal.MockStore) {
			store.EXPECT().Get(gomock.Any(), peer1).Times(1).Return(addPeer1, nil)
			store.EXPECT().Remove(gomock.Any(), peer1).Times(1)

			cfg.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(true)
		},
		nil,
	}, {
		"add-node",
		peer1,
		func(ctrl *gomock.Controller, host *mocks.MockHost, cfg *mocks.MockNetworkConfig, store *mockproposal.MockStore) {
			store.EXPECT().Get(gomock.Any(), peer1).Times(1).Return(addPeer1, nil)
			store.EXPECT().Remove(gomock.Any(), peer1).Times(1)

			cfg.EXPECT().IsAllowed(gomock.Any(), peer1).Times(1).Return(false)
			cfg.EXPECT().AddPeer(gomock.Any(), peer1, []multiaddr.Multiaddr{addPeer1.PeerAddr}).Times(1)
			cfg.EXPECT().Copy(gomock.Any()).Times(1)
			cfg.EXPECT().AllowedPeers(gomock.Any()).Times(1).Return([]peer.ID{peer1, peer2})

			stream := mocks.NewMockStream(ctrl)
			stream.EXPECT().Write(gomock.Any()).AnyTimes()
			stream.EXPECT().Close().Times(2)

			host.EXPECT().ID().AnyTimes().Return(coordinatorID)
			host.EXPECT().NewStream(gomock.Any(), peer1, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
			host.EXPECT().NewStream(gomock.Any(), peer2, bootstrap.PrivateCoordinatedConfigPID).Times(1).Return(stream, nil)
		},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			host := mocks.NewMockHost(ctrl)
			host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorHandshakePID, gomock.Any())
			host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatorProposePID, gomock.Any())

			networkConfig := mocks.NewMockNetworkConfig(ctrl)
			store := mockproposal.NewMockStore(ctrl)

			tt.configure(ctrl, host, networkConfig, store)

			handler, err := bootstrap.NewCoordinatorHandler(host, networkConfig, store)
			require.NoError(t, err, "bootstrap.NewCoordinatorHandler()")

			err = handler.Accept(ctx, tt.acceptID)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoordinator_Reject(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	peerID := test.GeneratePeerID(t)

	host := bhost.NewBlankHost(netutil.GenSwarmNetwork(t, ctx))
	store := mockproposal.NewMockStore(ctrl)
	store.EXPECT().Remove(gomock.Any(), peerID).Times(1)

	handler, err := bootstrap.NewCoordinatorHandler(host, nil, store)
	require.NoError(t, err, "bootstrap.NewCoordinatorHandler()")

	err = handler.Reject(ctx, peerID)
	require.NoError(t, err, "handler.Reject()")
}
