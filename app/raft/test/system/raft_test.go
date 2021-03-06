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

package system

import (
	"context"
	"testing"
	"time"

	raft "github.com/stratumn/go-node/app/raft/grpc"
	"github.com/stratumn/go-node/test/session"
	"github.com/stratumn/go-node/test/system"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const numNodes = 4                        // please change thoughtfully
const startTime = 2000 * time.Millisecond // increase if test stucks
const stepTime = 2000 * time.Millisecond  // increase if test stucks

// TestRaft executes a standart raft test scenario.

func assertNodes(t *testing.T, ctx context.Context,
	nn []raft.RaftClient,
	ss []raft.StatusInfo,
	pp [][]raft.Peer,
	ee [][]raft.Entry,
) {

	assert.Equal(t, len(nn), numNodes, "node array does not match numNodes, check your tests")
	assert.Equal(t, len(ss), numNodes, "status array does not match numNodes, check your tests")
	assert.Equal(t, len(pp), numNodes, "peer array does not match numNodes, check your tests")
	assert.Equal(t, len(ee), numNodes, "log array does not match numNodes, check your tests")

	// TODO: test logs with timestamps
	for i := 0; i < numNodes; i++ {
		s, err := nn[i].Status(ctx, &raft.Empty{})
		assert.NoError(t, err)
		assert.Equalf(t, ss[i], *s, "node %d status does not match", i+1)

		ppStream, err := nn[i].Peers(ctx, &raft.Empty{})
		assert.NoError(t, err)

		p := []raft.Peer{}
		for {
			peer, err := ppStream.Recv()
			if err != nil {
				break
			}
			p = append(p, *peer)
		}

		assert.Equalf(t, pp[i], p, "node %d peer list does not match", i+1)

		eeStream, err := nn[i].Log(ctx, &raft.Empty{})
		assert.NoError(t, err)

		e := []raft.Entry{}
		for {
			entry, err := eeStream.Recv()
			if err != nil {
				break
			}
			e = append(e, *entry)
		}

		assert.Equalf(t, ee[i], e, "node %d log does not match", i+1)

	}

}

func TestRaft(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), system.MaxDuration)
	defer cancel()

	sessionConf := session.WithServices(session.SystemCfg(), "boot", "raft")

	err := session.Run(ctx, system.SessionDir, numNodes, sessionConf,
		func(ctx context.Context, set session.TestNodeSet, conns []*grpc.ClientConn) {
			assert.Equal(t, numNodes, 4, "numNodes should be changed thoughtfully")

			var err error

			n1 := raft.NewRaftClient(conns[0])
			n2 := raft.NewRaftClient(conns[1])
			n3 := raft.NewRaftClient(conns[2])
			n4 := raft.NewRaftClient(conns[3])

			a1 := []byte(set[0].PeerID())
			a2 := []byte(set[1].PeerID())
			a3 := []byte(set[2].PeerID())
			a4 := []byte(set[3].PeerID())

			e1 := []byte("\xBA\xDD\xCA\xFE")
			e2 := []byte("\xDE\xAD\xBE\xEF")
			e3 := []byte("\xFE\xED\xFA\xCE")

			nn := []raft.RaftClient{n1, n2, n3, n4}
			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{},
					[]raft.Peer{},
					[]raft.Peer{},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n1.Start(ctx, &raft.Empty{})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
					},
					[]raft.Peer{},
					[]raft.Peer{},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n1.Invite(ctx, &raft.PeerID{Address: a2})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
					},
					[]raft.Peer{},
					[]raft.Peer{},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n2.Join(ctx, &raft.PeerID{Address: a1})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
					},
					[]raft.Peer{},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n2.Invite(ctx, &raft.PeerID{Address: a3})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n3.Join(ctx, &raft.PeerID{Address: a1})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 3},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
					[]raft.Entry{},
				},
			)

			_, err = n2.Propose(ctx, &raft.Proposal{Data: e1})
			assert.NoError(t, err)
			time.Sleep(stepTime) // Make sure e1 is added before e2
			_, err = n3.Propose(ctx, &raft.Proposal{Data: e2})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 3},
					raft.StatusInfo{Running: false, Id: 0},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
					},
					[]raft.Peer{},
				},
				[][]raft.Entry{
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{},
				},
			)

			_, err = n3.Invite(ctx, &raft.PeerID{Address: a4})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n4.Join(ctx, &raft.PeerID{Address: a2})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 1},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 3},
					raft.StatusInfo{Running: true, Id: 4},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
					[]raft.Peer{
						raft.Peer{Id: 1, Address: a1},
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
				},
				[][]raft.Entry{
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
				},
			)

			_, err = n2.Expel(ctx, &raft.PeerID{Address: a1})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: false, Id: 0},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 3},
					raft.StatusInfo{Running: true, Id: 4},
				},
				[][]raft.Peer{
					[]raft.Peer{},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
					},
				},
				[][]raft.Entry{
					[]raft.Entry{},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
					},
				},
			)

			dPeersStream, err := n1.Discover(ctx, &raft.PeerID{Address: a4})
			assert.NoError(t, err)

			dPeers := []raft.Peer{}
			for {
				peer, err := dPeersStream.Recv()
				if err != nil {
					break
				}
				dPeers = append(dPeers, *peer)
			}

			assert.Equalf(t, []raft.Peer{
				raft.Peer{Id: 2, Address: a2},
				raft.Peer{Id: 3, Address: a3},
				raft.Peer{Id: 4, Address: a4},
			}, dPeers, "discovered peer list does not match")

			_, err = n2.Propose(ctx, &raft.Proposal{Data: e3})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n4.Invite(ctx, &raft.PeerID{Address: a1})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n1.Join(ctx, &raft.PeerID{Address: a2})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 5},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 3},
					raft.StatusInfo{Running: true, Id: 4},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 3, Address: a3},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
					},
				},
				[][]raft.Entry{
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
				},
			)

			_, err = n3.Stop(ctx, &raft.Empty{})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n2.Expel(ctx, &raft.PeerID{Address: a3})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n4.Invite(ctx, &raft.PeerID{Address: a3})
			assert.NoError(t, err)
			time.Sleep(stepTime)

			_, err = n3.Join(ctx, &raft.PeerID{Address: a1})
			assert.NoError(t, err)
			time.Sleep(startTime)

			assertNodes(t, ctx, nn,
				[]raft.StatusInfo{
					raft.StatusInfo{Running: true, Id: 5},
					raft.StatusInfo{Running: true, Id: 2},
					raft.StatusInfo{Running: true, Id: 6},
					raft.StatusInfo{Running: true, Id: 4},
				},
				[][]raft.Peer{
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
						raft.Peer{Id: 6, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
						raft.Peer{Id: 6, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
						raft.Peer{Id: 6, Address: a3},
					},
					[]raft.Peer{
						raft.Peer{Id: 2, Address: a2},
						raft.Peer{Id: 4, Address: a4},
						raft.Peer{Id: 5, Address: a1},
						raft.Peer{Id: 6, Address: a3},
					},
				},
				[][]raft.Entry{
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
					[]raft.Entry{
						raft.Entry{Index: 0, Data: e1},
						raft.Entry{Index: 1, Data: e2},
						raft.Entry{Index: 2, Data: e3},
					},
				},
			)

		})

	assert.NoError(t, err, "Session()")

}
