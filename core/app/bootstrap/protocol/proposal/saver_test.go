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

package proposal_test

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stratumn/go-node/core/app/bootstrap/protocol/proposal"
	"github.com/stratumn/go-node/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaver(t *testing.T) {
	ctx := context.Background()

	dir, _ := ioutil.TempDir("", "stratumn-node")
	peer1 := test.GeneratePeerID(t)
	peer1Addr := test.GeneratePeerMultiaddr(t, peer1)

	t.Run("save-on-add-request", func(t *testing.T) {
		filePath := path.Join(dir, "save.json")
		s, err := proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		require.True(t, os.IsNotExist(err))

		err = s.AddRequest(ctx, &proposal.Request{
			Type:   proposal.RemoveNode,
			PeerID: peer1,
		})
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	})

	t.Run("save-on-add-vote", func(t *testing.T) {
		req := &proposal.Request{
			Type:      proposal.RemoveNode,
			PeerID:    peer1,
			Challenge: []byte("sp0ng3b0b"),
		}

		sm := proposal.NewInMemoryStore()
		err := sm.AddRequest(ctx, req)
		require.NoError(t, err)

		filePath := path.Join(dir, "vote.json")
		s, err := proposal.WrapWithSaver(ctx, sm, filePath)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		require.True(t, os.IsNotExist(err))

		v, err := proposal.NewVote(context.Background(), test.GeneratePrivateKey(t), req)
		require.NoError(t, err)

		err = s.AddVote(ctx, v)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	})

	t.Run("save-on-remove", func(t *testing.T) {
		filePath := path.Join(dir, "remove.json")
		s, err := proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		require.True(t, os.IsNotExist(err))

		err = s.Remove(ctx, peer1)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	})

	t.Run("save-on-list", func(t *testing.T) {
		filePath := path.Join(dir, "list.json")
		s, err := proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		require.True(t, os.IsNotExist(err))

		_, err = s.List(ctx)
		require.NoError(t, err)

		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	})

	t.Run("reject-invalid-data", func(t *testing.T) {
		filePath := path.Join(dir, "invalid.json")
		err := ioutil.WriteFile(filePath, []byte("YOLO"), 0744)
		require.NoError(t, err)

		_, err = proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		assert.Error(t, err)
	})

	t.Run("load-previous-data", func(t *testing.T) {
		filePath := path.Join(dir, "valid.json")
		s, err := proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		require.NoError(t, err)

		expires := time.Now().UTC().Add(24 * time.Hour)
		testReq := &proposal.Request{
			Type:      proposal.AddNode,
			PeerID:    peer1,
			PeerAddr:  peer1Addr,
			Info:      []byte("b4tm4n"),
			Challenge: []byte("much wow"),
			Expires:   expires,
		}

		err = s.AddRequest(ctx, testReq)
		require.NoError(t, err)

		v, err := proposal.NewVote(context.Background(), test.GeneratePrivateKey(t), testReq)
		require.NoError(t, err)

		err = s.AddVote(ctx, v)
		require.NoError(t, err)

		s2, err := proposal.WrapWithSaver(ctx, proposal.NewInMemoryStore(), filePath)
		require.NoError(t, err)

		r, err := s2.Get(ctx, peer1)
		require.NoError(t, err)
		require.NotNil(t, r)

		assert.Equal(t, proposal.AddNode, r.Type)
		assert.Equal(t, peer1, r.PeerID)
		assert.Equal(t, peer1Addr, r.PeerAddr)
		assert.Equal(t, testReq.Info, r.Info)
		assert.Equal(t, testReq.Challenge, r.Challenge)
		assert.Equal(t, expires, r.Expires)

		vv, err := s2.GetVotes(ctx, peer1)
		require.NoError(t, err)
		require.Len(t, vv, 1)

		assert.Equal(t, proposal.AddNode, vv[0].Type)
		assert.Equal(t, peer1, vv[0].PeerID)
		assert.Equal(t, testReq.Challenge, vv[0].Challenge)
		assert.NoError(t, vv[0].Verify(context.Background(), testReq))
	})
}
