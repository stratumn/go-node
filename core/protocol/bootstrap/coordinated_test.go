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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/protector"
	"github.com/stratumn/alice/core/protector/protectortest"
	"github.com/stratumn/alice/core/protocol/bootstrap"
	"github.com/stratumn/alice/core/protocol/bootstrap/bootstraptest"
	"github.com/stratumn/alice/core/protocol/bootstrap/proposal"
	"github.com/stratumn/alice/core/protocol/bootstrap/proposaltest"
	"github.com/stratumn/alice/core/streamutil"
	"github.com/stratumn/alice/core/streamutil/mockstream"
	"github.com/stratumn/alice/core/streamutil/streamtest"
	bootstrappb "github.com/stratumn/alice/pb/bootstrap"
	protectorpb "github.com/stratumn/alice/pb/protector"
	"github.com/stratumn/alice/test"
	"github.com/stratumn/alice/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inet "gx/ipfs/QmXoz9o2PT3tEzf7hicegwex5UgVP54n3k82K7jrWFyN86/go-libp2p-net"
	"gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	ihost "gx/ipfs/QmfZTdmunzKzAGJrSvXXQbQ5kLLUiEMX5vdwux7iXkdk7D/go-libp2p-host"
)

// TODO: remove/move to helper packages

func newNetworkConfig(t *testing.T) protector.NetworkConfig {
	config, err := protector.NewInMemoryConfig(
		context.Background(),
		protectorpb.NewNetworkConfig(protectorpb.NetworkState_BOOTSTRAP),
	)
	require.NoError(t, err, "protector.NewInMemoryConfig()")
	return config
}

func waitUntilAllowed(t *testing.T, peerID peer.ID, networkConfig protector.NetworkConfig) {
	test.WaitUntil(t, 100*time.Millisecond, 20*time.Millisecond,
		func() error {
			if !networkConfig.IsAllowed(context.Background(), peerID) {
				return errors.New("peer not allowed")
			}

			return nil
		}, "peer not allowed in time")
}

func waitUntilNotAllowed(t *testing.T, peerID peer.ID, networkConfig protector.NetworkConfig) {
	test.WaitUntil(t, 100*time.Millisecond, 20*time.Millisecond,
		func() error {
			if networkConfig.IsAllowed(context.Background(), peerID) {
				return errors.New("still allowed")
			}

			return nil
		}, "peer not removed in time")
}

func waitUntilProposed(t *testing.T, s proposal.Store, peerID peer.ID) {
	test.WaitUntil(t, 100*time.Millisecond, 10*time.Millisecond,
		func() error {
			r, _ := s.Get(context.Background(), peerID)
			if r == nil {
				return errors.New("proposal not received yet")
			}

			return nil
		}, "proposal not received in time")
}

func waitUntilDisconnected(t *testing.T, host ihost.Host, peerID peer.ID) {
	test.WaitUntil(t, 100*time.Millisecond, 10*time.Millisecond,
		func() error {
			if host.Network().Connectedness(peerID) == inet.Connected {
				return errors.New("peers still connected")
			}

			return nil
		}, "peers not disconnected in time")
}

// END TODO

func expectCoordinatedHost(host *mocks.MockHost) {
	host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatedConfigPID, gomock.Any())
	host.EXPECT().SetStreamHandler(bootstrap.PrivateCoordinatedProposePID, gomock.Any())
}

func TestCoordinated_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mocks.NewMockHost(ctrl)
	expectCoordinatedHost(host)

	mode := &protector.NetworkMode{
		CoordinatorID: test.GeneratePeerID(t),
	}
	handler := bootstrap.NewCoordinatedHandler(host, nil, mode, nil, nil)
	require.NotNil(t, handler)

	host.EXPECT().RemoveStreamHandler(bootstrap.PrivateCoordinatedConfigPID)
	host.EXPECT().RemoveStreamHandler(bootstrap.PrivateCoordinatedProposePID)

	handler.Close(context.Background())
}

type CoordinatedHandleTestCase struct {
	name          string
	coordinatorID peer.ID
	remotePeer    peer.ID
	configure     func(*testing.T, *gomock.Controller, *mocks.MockHost, *mockstream.MockCodec)
	validate      func(*testing.T, protector.NetworkConfig, proposal.Store)
	err           error
}

func (ht *CoordinatedHandleTestCase) Run(
	t *testing.T,
	h func(*bootstrap.CoordinatedHandler) streamutil.AutoCloseHandler,
) {
	t.Run(ht.name, func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		host := mocks.NewMockHost(ctrl)
		expectCoordinatedHost(host)

		conn := mocks.NewMockConn(ctrl)
		conn.EXPECT().RemotePeer().Return(ht.remotePeer)

		stream := mocks.NewMockStream(ctrl)
		stream.EXPECT().Conn().Return(conn)

		codec := mockstream.NewMockCodec(ctrl)

		cfg := protectortest.NewTestNetworkConfig(t, protectorpb.NetworkState_BOOTSTRAP)
		s := proposal.NewInMemoryStore()

		ht.configure(t, ctrl, host, codec)

		mode := &protector.NetworkMode{CoordinatorID: ht.coordinatorID}
		handler := bootstrap.NewCoordinatedHandler(host, nil, mode, cfg, s).(*bootstrap.CoordinatedHandler)
		err := h(handler)(ctx, bootstraptest.NewEvent(), stream, codec)

		if ht.err != nil {
			assert.EqualError(t, err, ht.err.Error())
		} else {
			require.NoError(t, err)
			ht.validate(t, cfg, s)
		}
	})
}

func TestCoordinated_HandleConfigUpdate(t *testing.T) {
	coordinatorKey := test.GeneratePrivateKey(t)
	coordinatorID := test.GetPeerIDFromKey(t, coordinatorKey)
	peer1Key := test.GeneratePrivateKey(t)
	peer1 := test.GetPeerIDFromKey(t, peer1Key)

	testCases := []CoordinatedHandleTestCase{
		{
			"reject-not-coordinator",
			coordinatorID,
			peer1,
			func(*testing.T, *gomock.Controller, *mocks.MockHost, *mockstream.MockCodec) {
			},
			func(*testing.T, protector.NetworkConfig, proposal.Store) {},
			protector.ErrConnectionRefused,
		},
		{
			"invalid-config-content",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, _ *gomock.Controller, _ *mocks.MockHost, codec *mockstream.MockCodec) {
				codec.EXPECT().Decode(gomock.Any()).Return(errors.New("BSOD"))
			},
			func(*testing.T, protector.NetworkConfig, proposal.Store) {},
			errors.Cause(errors.New("BSOD")),
		},
		{
			"invalid-config-signature",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, _ *gomock.Controller, _ *mocks.MockHost, codec *mockstream.MockCodec) {
				cfg := &protectorpb.NetworkConfig{
					NetworkState: 42,
				}
				err := cfg.Sign(context.Background(), coordinatorKey)
				require.NoError(t, err, "cfg.Sign()")

				streamtest.ExpectDecodeConfig(t, codec, cfg)
			},
			func(t *testing.T, cfg protector.NetworkConfig, _ proposal.Store) {
				assert.Equal(t, protectorpb.NetworkState_BOOTSTRAP, cfg.NetworkState(context.Background()))
			},
			protectorpb.ErrInvalidNetworkState,
		},
		{
			"invalid-config-signer",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, _ *gomock.Controller, _ *mocks.MockHost, codec *mockstream.MockCodec) {
				cfg := protectortest.NewTestNetworkConfig(
					t,
					protectorpb.NetworkState_PROTECTED,
					peer1,
				).Copy(context.Background())

				err := cfg.Sign(context.Background(), peer1Key)
				require.NoError(t, err, "cfg.Sign()")

				streamtest.ExpectDecodeConfig(t, codec, &cfg)
			},
			func(t *testing.T, cfg protector.NetworkConfig, _ proposal.Store) {
				assert.Equal(t, protectorpb.NetworkState_BOOTSTRAP, cfg.NetworkState(context.Background()))
				assert.False(t, cfg.IsAllowed(context.Background(), peer1))
			},
			protectorpb.ErrInvalidSignature,
		},
		{
			"valid-config",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, ctrl *gomock.Controller, h *mocks.MockHost, codec *mockstream.MockCodec) {
				cfg := protectortest.NewTestNetworkConfig(
					t,
					protectorpb.NetworkState_PROTECTED,
					peer1,
				).Copy(context.Background())

				err := cfg.Sign(context.Background(), coordinatorKey)
				require.NoError(t, err, "cfg.Sign()")

				streamtest.ExpectDecodeConfig(t, codec, &cfg)

				// Disconnect from unauthorized peers
				peer2 := test.GeneratePeerID(t)
				unauthorizedConn := mocks.NewMockConn(ctrl)
				unauthorizedConn.EXPECT().RemotePeer().Return(peer2)
				unauthorizedConn.EXPECT().Close()

				network := mocks.NewMockNetwork(ctrl)
				network.EXPECT().Conns().Return([]inet.Conn{unauthorizedConn})

				h.EXPECT().Network().Return(network)
			},
			func(t *testing.T, cfg protector.NetworkConfig, _ proposal.Store) {
				assert.Equal(t, protectorpb.NetworkState_PROTECTED, cfg.NetworkState(context.Background()))
				assert.True(t, cfg.IsAllowed(context.Background(), peer1))
			},
			nil,
		},
	}

	for _, tt := range testCases {
		tt.Run(t, func(handler *bootstrap.CoordinatedHandler) streamutil.AutoCloseHandler {
			return handler.HandleConfigUpdate
		})
	}
}

func TestCoordinated_HandlePropose(t *testing.T) {
	coordinatorKey := test.GeneratePrivateKey(t)
	coordinatorID := test.GetPeerIDFromKey(t, coordinatorKey)
	peer1 := test.GeneratePeerID(t)

	testCases := []CoordinatedHandleTestCase{
		{
			"reject-not-coordinator",
			coordinatorID,
			peer1,
			func(*testing.T, *gomock.Controller, *mocks.MockHost, *mockstream.MockCodec) {
			},
			func(*testing.T, protector.NetworkConfig, proposal.Store) {},
			protector.ErrConnectionRefused,
		},
		{
			"invalid-request",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, _ *gomock.Controller, _ *mocks.MockHost, codec *mockstream.MockCodec) {
				prop := &bootstrappb.UpdateProposal{
					UpdateType: bootstrappb.UpdateType_AddNode,
					NodeDetails: &bootstrappb.NodeIdentity{
						PeerId: []byte("b4tm4n"),
					},
				}
				streamtest.ExpectDecodeUpdateProp(t, codec, prop)
			},
			func(*testing.T, protector.NetworkConfig, proposal.Store) {},
			proposal.ErrInvalidPeerID,
		},
		{
			"valid-request",
			coordinatorID,
			coordinatorID,
			func(t *testing.T, _ *gomock.Controller, _ *mocks.MockHost, codec *mockstream.MockCodec) {
				prop := proposaltest.NewAddRequest(t, peer1)
				streamtest.ExpectDecodeUpdateProp(t, codec, prop.ToUpdateProposal())
			},
			func(t *testing.T, cfg protector.NetworkConfig, s proposal.Store) {
				r, err := s.Get(context.Background(), peer1)
				assert.NoError(t, err)
				assert.NotNil(t, r)
			},
			nil,
		},
	}

	for _, tt := range testCases {
		tt.Run(t, func(handler *bootstrap.CoordinatedHandler) streamutil.AutoCloseHandler {
			return handler.HandlePropose
		})
	}
}

func TestCoordinated_Handshake(t *testing.T) {
	t.Run("coordinator-unavailable", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		testNetwork := bootstraptest.NewTestNetwork(ctx, t)
		defer testNetwork.Close()

		unavailableCoordinatorID := test.GeneratePeerID(t)

		config := newNetworkConfig(t)
		handler, err := testNetwork.AddCoordinatedNode(
			unavailableCoordinatorID,
			config,
		)

		assert.EqualError(t, err, protector.ErrConnectionRefused.Error())
		assert.Nil(t, handler)
	})

	t.Run("coordinator-closes-conn", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		testNetwork := bootstraptest.NewTestNetwork(ctx, t)
		defer testNetwork.Close()

		coordinatorConfig := newNetworkConfig(t)
		coordinatorConfig.SetNetworkState(ctx, protectorpb.NetworkState_PROTECTED)

		_, err := testNetwork.AddCoordinatorNode(coordinatorConfig)
		require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

		coordinatorID := testNetwork.CoordinatorID()
		coordinatedConfig := newNetworkConfig(t)
		handler, err := testNetwork.AddCoordinatedNode(
			coordinatorID,
			coordinatedConfig,
		)

		assert.EqualError(t, err, protector.ErrConnectionRefused.Error())
		assert.Nil(t, handler)
	})

	t.Run("coordinator-invalid-signature", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		testNetwork := bootstraptest.NewTestNetwork(ctx, t)
		defer testNetwork.Close()

		coordinatorConfig := newNetworkConfig(t)
		_, err := testNetwork.AddCoordinatorNode(coordinatorConfig)
		require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

		coordinatorID := testNetwork.CoordinatorID()
		coordinatedConfig := newNetworkConfig(t)
		coordinatedNode, connect := testNetwork.PrepareCoordinatedNode(
			coordinatorID,
			coordinatedConfig,
		)

		err = coordinatorConfig.AddPeer(ctx, coordinatedNode.ID(), coordinatedNode.Addrs())
		require.NoError(t, err, "coordinatorConfig.AddPeer()")

		unknownKey := test.GeneratePrivateKey(t)
		err = coordinatorConfig.Sign(ctx, unknownKey)
		require.NoError(t, err, "coordinatorConfig.Sign()")

		handler, err := connect()
		assert.EqualError(t, err, protectorpb.ErrInvalidSignature.Error())
		assert.Nil(t, handler)
	})

	t.Run("coordinator-empty-config", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		testNetwork := bootstraptest.NewTestNetwork(ctx, t)
		defer testNetwork.Close()

		coordinatorConfig := newNetworkConfig(t)
		_, err := testNetwork.AddCoordinatorNode(coordinatorConfig)
		require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

		err = coordinatorConfig.Sign(ctx, testNetwork.CoordinatorKey())
		require.NoError(t, err, "coordinatorConfig.Sign()")

		coordinatedConfig := newNetworkConfig(t)
		_, connect := testNetwork.PrepareCoordinatedNode(
			testNetwork.CoordinatorID(),
			coordinatedConfig,
		)

		handler, err := connect()
		assert.NoError(t, err)
		assert.NotNil(t, handler)

		// When the coordinator returns an empty config, this is not a handshake error.
		// It means we're not whitelisted yet, but the network is still bootstrapping.
		assert.Len(t, coordinatedConfig.AllowedPeers(ctx), 0)
	})

	t.Run("coordinator-valid-config", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		testNetwork := bootstraptest.NewTestNetwork(ctx, t)
		defer testNetwork.Close()

		coordinatorConfig := newNetworkConfig(t)
		_, err := testNetwork.AddCoordinatorNode(coordinatorConfig)
		require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

		coordinatedConfig := newNetworkConfig(t)
		coordinatedNode, connect := testNetwork.PrepareCoordinatedNode(
			testNetwork.CoordinatorID(),
			coordinatedConfig,
		)

		err = coordinatorConfig.AddPeer(
			ctx,
			testNetwork.CoordinatorID(),
			test.GeneratePeerMultiaddrs(t, testNetwork.CoordinatorID()),
		)
		require.NoError(t, err, "coordinatorConfig.AddPeer")

		err = coordinatorConfig.AddPeer(
			ctx,
			coordinatedNode.ID(),
			coordinatedNode.Addrs(),
		)
		require.NoError(t, err, "coordinatorConfig.AddPeer")

		err = coordinatorConfig.Sign(ctx, testNetwork.CoordinatorKey())
		require.NoError(t, err, "coordinatorConfig.Sign()")

		handler, err := connect()
		assert.NoError(t, err)
		assert.NotNil(t, handler)

		assert.Len(t, coordinatedConfig.AllowedPeers(ctx), 2)
		assert.True(t, coordinatedConfig.IsAllowed(ctx, testNetwork.CoordinatorID()))
		assert.True(t, coordinatedConfig.IsAllowed(ctx, coordinatedNode.ID()))
	})
}

func TestCoordinated_AddNode(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testNetwork := bootstraptest.NewTestNetwork(ctx, t)
	defer testNetwork.Close()

	coordinatorHandler, err := testNetwork.AddCoordinatorNode(newNetworkConfig(t))
	require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

	networkConfig := newNetworkConfig(t)
	host, connect := testNetwork.PrepareCoordinatedNode(
		testNetwork.CoordinatorID(),
		networkConfig,
	)

	handler, err := connect()
	assert.NoError(t, err)
	assert.NotNil(t, handler)

	err = handler.AddNode(ctx, host.ID(), host.Addrs()[0], []byte("trust me, I'm b4tm4n"))
	require.NoError(t, err, "handler.AddNode()")

	// We shouldn't allow the node until the coordinator validates it.
	assert.False(t, networkConfig.IsAllowed(ctx, host.ID()))

	err = coordinatorHandler.Accept(ctx, host.ID())
	require.NoError(t, err, "coordinatorHandler.Accept()")

	waitUntilAllowed(t, host.ID(), networkConfig)
}

func TestCoordinated_RemoveNode(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testNetwork := bootstraptest.NewTestNetwork(ctx, t)
	defer testNetwork.Close()

	coordinatorConfig := newNetworkConfig(t)
	coordinatorHandler, err := testNetwork.AddCoordinatorNode(coordinatorConfig)
	require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

	coordinatedConfig := newNetworkConfig(t)

	h1, connect1 := testNetwork.PrepareCoordinatedNode(testNetwork.CoordinatorID(), coordinatedConfig)
	handler1, err := connect1()
	assert.NoError(t, err)
	assert.NotNil(t, handler1)

	h2, connect2 := testNetwork.PrepareCoordinatedNode(testNetwork.CoordinatorID(), coordinatedConfig)
	handler2, err := connect2()
	assert.NoError(t, err)
	assert.NotNil(t, handler2)

	err = coordinatorHandler.Accept(ctx, h1.ID())
	require.NoError(t, err, "coordinatorHandler.Accept()")

	err = coordinatorHandler.Accept(ctx, h2.ID())
	require.NoError(t, err, "coordinatorHandler.Accept()")

	err = coordinatorHandler.CompleteBootstrap(ctx)
	require.NoError(t, err, "coordinatorHandler.CompleteBootstrap()")

	err = h1.Connect(ctx, h2.Peerstore().PeerInfo(h2.ID()))
	require.NoError(t, err, "h1.Connect(h2)")

	assert.Equal(t, inet.Connected, h1.Network().Connectedness(h2.ID()))

	err = handler1.RemoveNode(ctx, h2.ID())
	require.NoError(t, err, "handler.RemoveNode()")

	waitUntilProposed(t, testNetwork.CoordinatorStore(), h2.ID())
	waitUntilProposed(t, testNetwork.CoordinatedStore(h1.ID()), h2.ID())

	// We shouldn't remove the node until enough votes are in.
	assert.True(t, coordinatedConfig.IsAllowed(ctx, h2.ID()))

	err = handler1.Accept(ctx, h2.ID())
	require.NoError(t, err, "handler1.Accept()")

	waitUntilNotAllowed(t, h2.ID(), coordinatedConfig)
	waitUntilDisconnected(t, h1, h2.ID())
}

func TestCoordinated_Accept(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testNetwork := bootstraptest.NewTestNetwork(ctx, t)
	defer testNetwork.Close()

	_, err := testNetwork.AddCoordinatorNode(newNetworkConfig(t))
	require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

	networkConfig := newNetworkConfig(t)
	host, connect := testNetwork.PrepareCoordinatedNode(
		testNetwork.CoordinatorID(),
		networkConfig,
	)

	handler, err := connect()
	assert.NoError(t, err)
	assert.NotNil(t, handler)

	propStore := testNetwork.CoordinatedStore(host.ID())

	t.Run("missing-request", func(t *testing.T) {
		err = handler.Accept(ctx, test.GeneratePeerID(t))
		assert.EqualError(t, err, proposal.ErrMissingRequest.Error())
	})

	t.Run("add-request", func(t *testing.T) {
		peerID := test.GeneratePeerID(t)
		err = propStore.AddRequest(ctx, &proposal.Request{
			Type:     proposal.AddNode,
			PeerID:   peerID,
			PeerAddr: test.GeneratePeerMultiaddr(t, peerID),
		})
		require.NoError(t, err)

		err = handler.Accept(ctx, peerID)
		assert.EqualError(t, err, bootstrap.ErrInvalidOperation.Error())
	})

	t.Run("remove-request-vote", func(t *testing.T) {
		peerID := test.GeneratePeerID(t)
		req := &proposal.Request{
			Type:      proposal.RemoveNode,
			PeerID:    peerID,
			Challenge: []byte("such challenge"),
		}

		err = propStore.AddRequest(ctx, req)
		require.NoError(t, err)

		err = handler.Accept(ctx, peerID)
		require.NoError(t, err)

		// Should have been removed from the store once accepted.
		r, _ := propStore.Get(ctx, peerID)
		require.Nil(t, r)
	})
}

func TestCoordinated_Reject(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testNetwork := bootstraptest.NewTestNetwork(ctx, t)
	defer testNetwork.Close()

	_, err := testNetwork.AddCoordinatorNode(newNetworkConfig(t))
	require.NoError(t, err, "testNetwork.AddCoordinatorNode()")

	networkConfig := newNetworkConfig(t)
	host, connect := testNetwork.PrepareCoordinatedNode(
		testNetwork.CoordinatorID(),
		networkConfig,
	)

	handler, err := connect()
	assert.NoError(t, err)
	assert.NotNil(t, handler)

	propStore := testNetwork.CoordinatedStore(host.ID())

	peerID := test.GeneratePeerID(t)
	err = handler.Reject(ctx, peerID)
	require.NoError(t, err)

	err = propStore.AddRequest(ctx, &proposal.Request{
		Type:      proposal.RemoveNode,
		PeerID:    peerID,
		Challenge: []byte("much chall3ng3"),
	})
	require.NoError(t, err)

	err = handler.Reject(ctx, peerID)
	require.NoError(t, err)

	r, _ := propStore.Get(ctx, peerID)
	assert.Nil(t, r)
}
