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
	"testing"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/pkg/errors"
	pb "github.com/stratumn/alice/core/app/metrics/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	p2pmetrics "gx/ipfs/QmVvu4bS5QLfS19ePkp5Wgzn2ZUma5oXTT9BgDFyQLxUZF/go-libp2p-metrics"
)

func testGRPCServer() grpcServer {
	mtrx := newMetrics(
		p2pmetrics.NewBandwidthCounter(),
		metrics.NewInmemSink(time.Second, time.Second),
	)
	return grpcServer{func() *Metrics { return mtrx }}
}

func testGRPCServerUnavailable() grpcServer {
	return grpcServer{func() *Metrics { return nil }}
}
func TestGRPCServer_Bandwidth(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv := testGRPCServer()

	req := &pb.BandwidthReq{}
	res, err := srv.Bandwidth(ctx, req)

	require.NoError(t, err)
	require.Equal(t, uint64(0), res.TotalIn)
}

func TestGRPCServer_Bandwidth_unavailable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv := testGRPCServerUnavailable()

	req := &pb.BandwidthReq{}
	_, err := srv.Bandwidth(ctx, req)

	assert.Equal(t, ErrUnavailable, errors.Cause(err))
}