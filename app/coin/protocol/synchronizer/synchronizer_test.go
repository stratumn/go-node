// Copyright © 2017-2018 Stratumn SAS
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package synchronizer

import (
	"context"
	"errors"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"testing"

	"github.com/stratumn/go-node/app/coin/protocol/coinutil"

	"github.com/stratumn/go-node/app/coin/pb"
	tassert "github.com/stratumn/go-node/app/coin/protocol/testutil/assert"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/stratumn/go-node/app/coin/protocol/p2p/mockp2p"
	"github.com/stratumn/go-node/app/coin/protocol/synchronizer/mocksynchronizer"
	"github.com/stratumn/go-node/app/coin/protocol/testutil"
)

func TestSynchronize(t *testing.T) {

	block0 := &pb.Block{Header: &pb.Header{BlockNumber: 0}}
	hash0, err := coinutil.HashHeader(block0.Header)
	assert.NoError(t, err, "HashHeader()")

	block1 := &pb.Block{Header: &pb.Header{BlockNumber: 1, PreviousHash: hash0}}
	hash1, err := coinutil.HashHeader(block1.Header)
	assert.NoError(t, err, "HashHeader()")

	block2 := &pb.Block{Header: &pb.Header{BlockNumber: 2, PreviousHash: hash1}}

	block1bis := &pb.Block{Header: &pb.Header{BlockNumber: 1, PreviousHash: hash0, Nonce: 42}}
	hash1bis, err := coinutil.HashHeader(block1bis.Header)
	assert.NoError(t, err, "HashHeader()")

	block2bis := &pb.Block{Header: &pb.Header{BlockNumber: 2, PreviousHash: hash1bis, Nonce: 42}}
	hash2bis, err := coinutil.HashHeader(block2bis.Header)
	assert.NoError(t, err, "HashHeader()")

	block3bis := &pb.Block{Header: &pb.Header{BlockNumber: 3, PreviousHash: hash2bis, Nonce: 42}}
	hash3bis, err := coinutil.HashHeader(block3bis.Header)
	assert.NoError(t, err, "HashHeader()")

	cid, err := cid.Cast(hash3bis)
	assert.NoError(t, err, "cid.Cast()")
	pid1 := peer.ID("yankee")
	pid2 := peer.ID("zoulou")

	t.Run("synchronize-genesis-chain", func(t *testing.T) {
		/*
			the node starts with only the genesis block:
					block0

			after the sync we have:
					block0 -> block1bis -> block2bis -> block3bis
		*/
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := &testutil.SimpleChain{}

		c.AddBlock(block0)
		c.SetHead(block0)

		p2p := mockp2p.NewMockP2P(ctrl)
		dht := mocksynchronizer.NewMockContentProviderFinder(ctrl)
		sync := NewSynchronizer(p2p, dht, OptMaxBatchSizes(2))

		// Called to get the list of peers who have the block.
		dht.EXPECT().FindProviders(gomock.Any(), cid).Return([]pstore.PeerInfo{pstore.PeerInfo{ID: pid1}, pstore.PeerInfo{ID: pid2}}, nil).Times(1)

		// Get the block from the peer to check if he really has it.
		// First peer fails => call the second one.
		p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid1, []byte(hash3bis)).Return(nil, errors.New("")).Times(1)
		p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid2, []byte(hash3bis)).Return(block2bis, nil).Times(1)

		// Get the common ancestor.
		p2p.EXPECT().RequestHeadersByNumber(gomock.Any(), pid2, uint64(0), uint64(1)).Return([]*pb.Header{block0.Header}, nil).Times(1)

		// Get the blocks from that peer.
		p2p.EXPECT().RequestBlocksByNumber(gomock.Any(), pid2, uint64(1), uint64(2)).Return([]*pb.Block{block1bis, block2bis}, nil).Times(1)
		p2p.EXPECT().RequestBlocksByNumber(gomock.Any(), pid2, uint64(3), uint64(2)).Return([]*pb.Block{block3bis}, nil).Times(1)

		resCh, errCh := sync.Synchronize(context.Background(), hash3bis, c)

		tassert.WaitUntil(t, func() bool {
			// Check that we receive the blocks in this order.
			blocks := []*pb.Block{block1bis, block2bis, block3bis, nil}
			i := 0
			for {
				select {
				case b := <-resCh:
					assert.Equal(t, blocks[i], b, "resCh")
					i++
					if i == 3 {
						return true
					}
				case err := <-errCh:
					assert.NoError(t, err, "errCh")
					return false

				}
			}
		}, "channels")
	})

	t.Run("synchronize-long-chain", func(t *testing.T) {
		/*
			the node starts with:
					block0 -> block1 -> block2

			after the sync we have:
					block0 -> block1 -> block2
								\-> block1bis -> block2bis -> block3bis
		*/
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := &testutil.SimpleChain{}

		c.AddBlock(block0)
		c.SetHead(block0)

		c.AddBlock(block1)
		c.SetHead(block1)

		c.AddBlock(block2)
		c.SetHead(block2)

		p2p := mockp2p.NewMockP2P(ctrl)
		dht := mocksynchronizer.NewMockContentProviderFinder(ctrl)
		sync := NewSynchronizer(p2p, dht, OptMaxBatchSizes(2))

		// Called to get the list of peers who have the block.
		dht.EXPECT().FindProviders(gomock.Any(), cid).Return([]pstore.PeerInfo{pstore.PeerInfo{ID: pid1}, pstore.PeerInfo{ID: pid2}}, nil).Times(1)

		// Get the block from the peer to check if he really has it.
		// First peer fails => call the second one.
		p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid1, []byte(hash3bis)).Return(nil, errors.New("")).Times(1)
		p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid2, []byte(hash3bis)).Return(block2bis, nil).Times(1)

		// Get the common ancestor.
		p2p.EXPECT().RequestHeadersByNumber(gomock.Any(), pid2, uint64(1), uint64(2)).Return([]*pb.Header{block1bis.Header, block2bis.Header}, nil).Times(1)

		// Get the blocks from that peer.
		p2p.EXPECT().RequestBlocksByNumber(gomock.Any(), pid2, uint64(1), uint64(2)).Return([]*pb.Block{block1bis, block2bis}, nil).Times(1)
		p2p.EXPECT().RequestBlocksByNumber(gomock.Any(), pid2, uint64(3), uint64(2)).Return([]*pb.Block{block3bis}, nil).Times(1)

		resCh, errCh := sync.Synchronize(context.Background(), hash3bis, c)

		tassert.WaitUntil(t, func() bool {
			// Check that we receive the blocks in this order.
			blocks := []*pb.Block{block1bis, block2bis, block3bis, nil}
			i := 0
			for {
				select {
				case b := <-resCh:
					assert.Equal(t, blocks[i], b, "resCh")
					i++
					if i == 3 {
						return true
					}
				case err := <-errCh:
					assert.NoError(t, err, "errCh")
					return false

				}
			}
		}, "channels")
	})

}
func TestImpossibleSynchronize(t *testing.T) {
	/*
		Test that a node cannot force us to synchronize on a chain that has a different genesis.
		the node starts with
				block0 -> block1 -> block2

		the peer has
				block0bis -> -> block1bis -> block2bis -> block3bis
	*/
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := &testutil.SimpleChain{}
	block0 := &pb.Block{Header: &pb.Header{BlockNumber: 0}}
	hash0, err := coinutil.HashHeader(block0.Header)
	assert.NoError(t, err, "HashHeader()")
	c.AddBlock(block0)
	c.SetHead(block0)

	block1 := &pb.Block{Header: &pb.Header{BlockNumber: 1, PreviousHash: hash0}}
	hash1, err := coinutil.HashHeader(block1.Header)
	assert.NoError(t, err, "HashHeader()")
	c.AddBlock(block1)
	c.SetHead(block1)

	block2 := &pb.Block{Header: &pb.Header{BlockNumber: 2, PreviousHash: hash1}}
	c.AddBlock(block2)
	c.SetHead(block2)

	p2p := mockp2p.NewMockP2P(ctrl)
	dht := mocksynchronizer.NewMockContentProviderFinder(ctrl)
	sync := NewSynchronizer(p2p, dht, OptMaxBatchSizes(2))

	block0bis := &pb.Block{Header: &pb.Header{BlockNumber: 0, Nonce: 42}}
	hash0bis, err := coinutil.HashHeader(block0bis.Header)
	assert.NoError(t, err, "HashHeader()")

	block1bis := &pb.Block{Header: &pb.Header{BlockNumber: 1, PreviousHash: hash0bis}}
	hash1bis, err := coinutil.HashHeader(block1bis.Header)
	assert.NoError(t, err, "HashHeader()")

	block2bis := &pb.Block{Header: &pb.Header{BlockNumber: 2, PreviousHash: hash1bis}}
	hash2bis, err := coinutil.HashHeader(block2bis.Header)
	assert.NoError(t, err, "HashHeader()")

	block3bis := &pb.Block{Header: &pb.Header{BlockNumber: 3, PreviousHash: hash2bis}}
	hash3bis, err := coinutil.HashHeader(block3bis.Header)
	assert.NoError(t, err, "HashHeader()")

	assert.NoError(t, err, "hex.DecodeString()")
	cid, err := cid.Cast(hash3bis)
	assert.NoError(t, err, "cid.Cast()")

	pid1 := peer.ID("yankee")
	pid2 := peer.ID("zoulou")

	// Called to get the list of peers who have the block.
	dht.EXPECT().FindProviders(gomock.Any(), cid).Return([]pstore.PeerInfo{pstore.PeerInfo{ID: pid1}, pstore.PeerInfo{ID: pid2}}, nil).Times(1)

	// Get the block from the peer to check if he really has it.
	// First peer fails => call the second one.
	p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid1, []byte(hash3bis)).Return(nil, errors.New("")).Times(1)
	p2p.EXPECT().RequestBlockByHash(gomock.Any(), pid2, []byte(hash3bis)).Return(block2bis, nil).Times(1)

	// Get the common ancestor.
	p2p.EXPECT().RequestHeadersByNumber(gomock.Any(), pid2, uint64(1), uint64(2)).Return([]*pb.Header{block1bis.Header, block2bis.Header}, nil).Times(1)
	p2p.EXPECT().RequestHeadersByNumber(gomock.Any(), pid2, uint64(0), uint64(1)).Return([]*pb.Header{block0bis.Header, block1bis.Header}, nil).Times(1)

	resCh, errCh := sync.Synchronize(context.Background(), hash3bis, c)

	tassert.WaitUntil(t, func() bool {
		// Check that we receive the blocks in this order.
		for {
			select {
			case <-resCh:
				assert.Fail(t, "<-resCh")
				return false
			case err := <-errCh:
				assert.Error(t, err, "errCh")
				return true

			}
		}
	}, "channels")

}
