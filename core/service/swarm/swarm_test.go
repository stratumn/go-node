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

package swarm

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/manager/testservice"
	"github.com/stratumn/alice/core/service/metrics"
	"github.com/stratumn/alice/core/service/swarm/mockswarm"

	swarm "gx/ipfs/QmdQFrFnPrKRQtpeHKjZ3cVNwxmGKKS2TvhJTuN9C9yduh/go-libp2p-swarm"
)

func testService(ctx context.Context, t *testing.T, smuxer Transport) *Service {
	serv := &Service{}
	config := serv.Config().(Config)
	config.Addresses = []string{"/ip4/0.0.0.0/tcp/35768"}
	config.Metrics = ""

	if err := serv.SetConfig(config); err != nil {
		t.Fatalf("serv.SetConfig(config): error: %s", err)
	}

	deps := map[string]interface{}{
		"mssmux": smuxer,
	}

	if err := serv.Plug(deps); err != nil {
		t.Fatalf("serv.Plug(deps): error: %s", err)
	}

	return serv
}

func TestService_strings(t *testing.T) {
	testservice.CheckStrings(t, &Service{})
}

func TestService_Expose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	smuxer := mockswarm.NewMockTransport(ctrl)
	serv := testService(ctx, t, smuxer)
	exposed := testservice.Expose(ctx, t, serv, time.Second)

	_, ok := exposed.(*swarm.Swarm)
	if got, want := ok, true; got != want {
		t.Errorf("ok = %v want %v", got, want)
	}
}

func TestService_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	smuxer := mockswarm.NewMockTransport(ctrl)
	serv := testService(ctx, t, smuxer)
	testservice.TestRun(ctx, t, serv, time.Second)
}

func TestService_SetConfig(t *testing.T) {
	errAny := errors.New("any error")

	tt := []struct {
		name string
		set  func(*Config)
		err  error
	}{{
		"invalid peer ID",
		func(c *Config) { c.PeerID = "alice" },
		errAny,
	}, {
		"invalid private key",
		func(c *Config) { c.PrivateKey = "alice" },
		errAny,
	}, {
		"peer ID mismatch",
		func(c *Config) {
			c.PeerID = "QmVhJVRSYHNSHgR9dJNbDxu6G7GPPqJAeiJoVRvcexGNf9"
			c.PrivateKey = "CAASqAkwggSkAgEAAoIBAQDS3P9zWlSDuHDr1oZzarlf8fTUV4IcgFPvneAXPGQOf/ENaxkXeoFqRaHyYzVrrPumN1ofGQk0zJwe5oxXoKbXeTmCIOwDowqXKH47ldzxnINJtUT0tUC3V092T5j1PfAGdlJJNf7ttQxCYd2/Uy6wfMq65yZR1GENMpRfP/lPZG+zDw2gDR4UdP4Di8napepeA+PWcQBXayWOrVSgPY/Lp6GoTqThdkoV1sV/R58w6wS3+8yjG14Vv4nBnQOKWzeLgCrZXV4X90ygFjfwks6SNWCjzY8VFP1nps1Cm7HGl8KW/lfhrLPsW2m1/LNkp1ORRd46nwu2d43V9rsfk2TbAgMBAAECggEBALviVO9YrnOLtXo+dSCEGNbnxaoqqvFwWtnoB1NM6S6OS6AULJBiBMxHbUyHn4Lo6HWrXm7VJQHZysYx0R0HNYJLLrXHVeTLstULLKc1GmBigAz+KooMsrRqZJBbzkO+w49OgHVhWlw99MV1ZFtk5/Yzy4QMEHXbUfqrdc8FNsTIb/rxTfVgbgY+vtLxYf6Lq93n1iX7s83uQMPtczB1HUP9GU4ew3EwYjgpPHnlVJeIpYhM5ngjeKWkyPk3HFbhVtihNTvnSkrJNso0X4GMd9WkM5Q/WkANFt2LxekjFLpdUrEcNe+Orh278l0LGr42p4ripQUYPg0smfokSAhAXYECgYEA8Nq3eotu3/xHn4Y+HQyqHMxFx0y2/YtO6fuFSFpZ5kX6UBVTDT/o00ikhzBmMg8JMjqFV1JBG9t0P4yHRSyRPLbN9rVhOP4FLCZ8EpnuHbSUcfLqYwXAo4BDR4DXyuyRtHDEOMqomLNc5FPI4FnwAyNZ5LcO4rE+/hY8a/W5yxsCgYEA4B97vhECKfOhHQ5icAfJizLHcdQKFIBSDj5/Tijbh/0uehbr5BZObhaZUfKU/p5HmFQhCurh6NxTr9F5BjJ8YYUzsm66kJPeoQO8IaOFKPD4bRCTbRXFu7qHjnNc6jdw1IeC68ApxSoEtxEaGJL3kzLh6fg5EXxP5qPC8eAWqUECgYEAo3E4my8tgU/YZreZROtIMRypqXI0p1+2oG9vZcbyRKJuF5Qw9MfOvjoIdDjy0LuFWRF/VN9bkYTdoRZC4T06HcJLiERTgnJWnjxLa/ALNxtItP7L8YCA1jL+9PHI/kqFIbZ4YbWcrWrh+YulwCEoD2kY4m0a69itz4zVWcm5V8sCgYALNqgLT3CLRsxF1uVn84vK8iR2doR2mCEC42+dKoApYqqDo0f0JoWQDoNnTTrVLngoj/UDRdM9wmBRiKqEe9wrSO3YPKALAcr+xWARUswjy0KyukSWDaPSC7gikXURpup3R7xuLTQp0DtiKXHjzt6iN8aD3U6FqHGa+ZCUZ4DawQKBgGB4ZgWj/MEr/tOHefNv2QGDFvf/8TAKWFSTyuOSkepcSazil+jqoJWrQdfa0Ku9kRI4b42fyjWltMkf7Nr1QcIH8Zf/GP/xfoZ138uOq95WkgT6YQEbimHZPE6ozC3zhqEgwv0MWxfP+dAJl1uyb1ffmyqDD2N5V16+0WSSBgI2"
		},
		ErrPeerIDMismatch,
	}}

	for _, test := range tt {
		serv := Service{}
		config := serv.Config().(Config)
		test.set(&config)

		err := errors.Cause(serv.SetConfig(config))
		switch {
		case err != nil && test.err == errAny:
		case err != test.err:
			t.Errorf("%s: err = %v want %v", test.name, err, test.err)
		}
	}
}

func TestService_Needs(t *testing.T) {
	tt := []struct {
		name  string
		set   func(*Config)
		needs []string
	}{{
		"stream muxer",
		func(c *Config) { c.StreamMuxer = "mysmux" },
		[]string{"mysmux"},
	}, {
		"metrics",
		func(c *Config) { c.Metrics = "mymetrics" },
		[]string{"mymetrics"},
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

	smuxer := mockswarm.NewMockTransport(ctrl)

	tt := []struct {
		name string
		set  func(*Config)
		deps map[string]interface{}
		err  error
	}{{
		"valid stream muxer",
		func(c *Config) { c.StreamMuxer = "mysmux" },
		map[string]interface{}{
			"mysmux": smuxer,
		},
		nil,
	}, {
		"invalid stream muxer",
		func(c *Config) { c.StreamMuxer = "mysmux" },
		map[string]interface{}{
			"mysmux": struct{}{},
		},
		ErrNotStreamMuxer,
	}, {
		"valid metrics",
		func(c *Config) { c.Metrics = "mymetrics" },
		map[string]interface{}{
			"mymetrics": &metrics.Metrics{},
			"mssmux":    smuxer,
		},
		nil,
	}, {
		"invalid metrics",
		func(c *Config) { c.Metrics = "mymetrics" },
		map[string]interface{}{
			"mymetrics": struct{}{},
			"mssmux":    smuxer,
		},
		ErrNotMetrics,
	}}

	for _, test := range tt {
		serv := Service{}
		config := serv.Config().(Config)
		config.Metrics = ""
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
