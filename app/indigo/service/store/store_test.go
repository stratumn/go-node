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

package store_test

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	storeprotocol "github.com/stratumn/alice/app/indigo/protocol/store"
	"github.com/stratumn/alice/app/indigo/protocol/store/sync"
	"github.com/stratumn/alice/app/indigo/service/store"
	"github.com/stratumn/alice/app/indigo/service/store/mockstore"
	"github.com/stratumn/alice/core/manager/testservice"
	"github.com/stratumn/alice/core/service/pubsub/mockpubsub"
	"github.com/stratumn/alice/core/service/swarm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	floodsub "gx/ipfs/QmVKrsEgixRtMWcMd6WQzuwqCUC3jfLf7Q7xcjnKoMMikS/go-libp2p-floodsub"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	crypto "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
)

func validSwarm() *swarm.Swarm {
	sk, _, _ := crypto.GenerateEd25519Key(rand.Reader)
	return &swarm.Swarm{PrivKey: sk}
}

func testService(ctx context.Context, t *testing.T, host *mockstore.MockHost) *store.Service {
	serv := &store.Service{}
	config := serv.Config().(store.Config)
	config.Version = "1.0.0"
	config.NetworkID = "42"

	require.NoError(t, serv.SetConfig(config), "serv.SetConfig(config)")

	deps := map[string]interface{}{
		"host":  host,
		"swarm": validSwarm(),
	}

	require.NoError(t, serv.Plug(deps), "serv.Plug(deps)")

	return serv
}

func TestService_Strings(t *testing.T) {
	testservice.CheckStrings(t, &store.Service{})
}

// expectHostNetwork verifies that the service joins a PoP network via floodsub.
func expectHostNetwork(ctrl *gomock.Controller, host *mockstore.MockHost) {
	net := mockpubsub.NewMockNetwork(ctrl)

	host.EXPECT().Network().Return(net)
	net.EXPECT().Notify(gomock.Any())
	host.EXPECT().SetStreamHandler(protocol.ID(floodsub.FloodSubID), gomock.Any())
	host.EXPECT().SetStreamHandler(protocol.ID(sync.SingleNodeProtocolID), gomock.Any())
	host.EXPECT().RemoveStreamHandler(protocol.ID(floodsub.FloodSubID))
	host.EXPECT().RemoveStreamHandler(protocol.ID(sync.SingleNodeProtocolID))
}

func TestService_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mockstore.NewMockHost(ctrl)
	expectHostNetwork(ctrl, host)

	serv := testService(ctx, t, host)
	testservice.TestRun(ctx, t, serv, time.Second)
}

func TestService_Run_Error(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name   string
		err    string
		config store.Config
	}{{
		"missing-network-id",
		storeprotocol.ErrInvalidNetworkID.Error(),
		store.Config{
			Version:          "1.0.0",
			StorageType:      "in-memory",
			ValidationConfig: &store.ValidationConfig{},
		},
	}, {
		"invalid-storage-type",
		store.ErrStorageNotSupported.Error(),
		store.Config{
			Version:          "1.0.0",
			StorageType:      "on-the-moon",
			NetworkID:        "42",
			ValidationConfig: &store.ValidationConfig{},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := &store.Service{}
			require.NoError(t, serv.SetConfig(tt.config), "serv.SetConfig(config)")
			assert.EqualError(t, serv.Run(ctx, func() {}, func() {}), tt.err)
		})
	}
}

func TestService_Plug(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mockstore.NewMockHost(ctrl)

	tests := []struct {
		name string
		set  func(*store.Config)
		deps map[string]interface{}
		err  error
	}{{
		"valid-private-key",
		func(c *store.Config) { c.Swarm = "myswarm" },
		map[string]interface{}{
			"host":    host,
			"myswarm": validSwarm(),
		},
		nil,
	}, {
		"missing-private-key",
		func(c *store.Config) { c.Swarm = "myswarm" },
		map[string]interface{}{
			"host":    host,
			"myswarm": &swarm.Swarm{},
		},
		store.ErrNotSwarm,
	}, {
		"missing-swarm",
		func(c *store.Config) { c.Swarm = "myswarm" },
		map[string]interface{}{
			"host": host,
		},
		store.ErrNotSwarm,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := store.Service{}
			config := serv.Config().(store.Config)
			config.Host = "host"
			tt.set(&config)

			require.NoError(t, serv.SetConfig(config), "serv.SetConfig(config)")

			err := errors.Cause(serv.Plug(tt.deps))
			switch {
			case err != tt.err:
				assert.Equal(t, tt.err, err)
			}
		})
	}
}