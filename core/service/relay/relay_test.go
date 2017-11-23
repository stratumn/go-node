// Copyright © 2017  Stratumn SAS
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

package relay

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/manager/testservice"
	"github.com/stratumn/alice/core/service/relay/mockrelay"

	testutil "gx/ipfs/QmQGX417WoxKxDJeHqouMEmmH4G1RCENNSzkZYHrXy3Xb3/go-libp2p-netutil"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	circuit "gx/ipfs/Qmf7GSJ4omRJsvA9uzTqzbnVhq4RWLPzjzW4xJzUta4dKE/go-libp2p-circuit"
)

func testService(ctx context.Context, t *testing.T, host Host) *Service {
	serv := &Service{}
	config := serv.Config().(Config)

	if err := serv.SetConfig(config); err != nil {
		t.Fatalf("serv.SetConfig(config): error: %s", err)
	}

	deps := map[string]interface{}{
		"host": host,
	}

	if err := serv.Plug(deps); err != nil {
		t.Fatalf("serv.Plug(deps): error: %s", err)
	}

	return serv
}

func expectHost(ctx context.Context, t *testing.T, host *mockrelay.MockHost) {
	swm := testutil.GenSwarmNetwork(t, ctx)

	host.EXPECT().ID().Return(swm.LocalPeer()).AnyTimes()
	host.EXPECT().Network().Return(swm).AnyTimes()
	host.EXPECT().SetStreamHandler(protocol.ID(circuit.ProtoID), gomock.Any())
	host.EXPECT().RemoveStreamHandler(protocol.ID(circuit.ProtoID))
}

func TestService_strings(t *testing.T) {
	testservice.CheckStrings(t, &Service{})
}

func TestService_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mockrelay.NewMockHost(ctrl)
	expectHost(ctx, t, host)

	serv := testService(ctx, t, host)
	testservice.TestRun(ctx, t, serv, time.Second)
}

func TestService_Needs(t *testing.T) {
	tt := []struct {
		name  string
		set   func(*Config)
		needs []string
	}{{
		"host",
		func(c *Config) { c.Host = "myhost" },
		[]string{"myhost"},
	}}

	for _, test := range tt {
		serv := Service{}
		config := serv.Config().(Config)
		test.set(&config)

		if err := serv.SetConfig(config); err != nil {
			t.Errorf("%s: serv.SetConfig(config): error: %s", test.name, err)
			continue
		}

		needs := serv.Needs()
		for _, n := range test.needs {
			if _, ok := needs[n]; !ok {
				t.Errorf("%s: needs[%q] = nil want struct{}{}", test.name, n)
			}
		}
	}
}

func TestService_Plug(t *testing.T) {
	errAny := errors.New("any error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	host := mockrelay.NewMockHost(ctrl)

	tt := []struct {
		name string
		set  func(*Config)
		deps map[string]interface{}
		err  error
	}{{
		"valid host",
		func(c *Config) { c.Host = "myhost" },
		map[string]interface{}{
			"myhost": host,
		},
		nil,
	}, {
		"invalid host",
		func(c *Config) { c.Host = "myhost" },
		map[string]interface{}{
			"myhost": struct{}{},
		},
		ErrNotHost,
	}}

	for _, test := range tt {
		serv := Service{}
		config := serv.Config().(Config)
		test.set(&config)

		if err := serv.SetConfig(config); err != nil {
			t.Errorf("%s: serv.SetConfig(config): error: %s", test.name, err)
			continue
		}

		err := errors.Cause(serv.Plug(test.deps))
		switch {
		case err != nil && test.err == errAny:
		case err != test.err:
			t.Errorf("%s: err = %v want %v", test.name, err, test.err)
		}
	}
}
