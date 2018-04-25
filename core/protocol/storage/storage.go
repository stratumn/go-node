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

package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/stratumn/alice/core/db"
	"github.com/stratumn/alice/core/protocol/storage/acl"
	"github.com/stratumn/alice/core/protocol/storage/file"
	"github.com/stratumn/alice/core/protocol/storage/p2p"
	pb "github.com/stratumn/alice/pb/storage"

	ihost "gx/ipfs/QmNmJZL7FQySMtE2BQuLMuZg2EB2CLEunJJUSVSc9YnnbV/go-libp2p-host"
	protobuf "gx/ipfs/QmRDePEiL4Yupq5EkcK3L3ko3iMgYaqUdLu7xc1kqs7dnV/go-multicodec/protobuf"
	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	inet "gx/ipfs/QmXfkENeeBvh3zYA51MaSdGUdBjhQ99cP5WQe8zgr6wchG/go-libp2p-net"
	peer "gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"
)

// ChunkSize size of a chunk of data
const ChunkSize = 1024

// Host represents an Alice host.
type Host = ihost.Host

// log is the logger for the service.
var log = logging.Logger("storage")

var (
	prefixFilesHashes = []byte("fh") // prefixFilesHashes + filehash -> filepath
)

var (
	// ErrUnauthorized is returned when a peer tries to access a file he
	// is not allowed to get
	ErrUnauthorized = errors.New("peer not authorized for requested file")
)

// Storage implements the storage protocol.
type Storage struct {
	db          db.DB
	host        Host
	chunkSize   int
	fileHandler file.Handler
	p2p         p2p.P2P
	acl         acl.ACL

	storagePath string
}

// NewStorage creates a new storage server.
func NewStorage(host Host, db db.DB, storagePath string) *Storage {
	fh := file.NewLocalFileHandler(storagePath)
	return &Storage{
		db:          db,
		host:        host,
		chunkSize:   ChunkSize,
		fileHandler: fh,
		p2p:         p2p.NewP2P(host, fh),
		acl:         acl.NewACL(db),

		storagePath: storagePath,
	}
}

// SaveFile saves a file locally.
func (s *Storage) SaveFile(ctx context.Context, chunkCh <-chan *pb.FileChunk) ([]byte, error) {
	event := log.EventBegin(ctx, "SaveFile")
	defer event.Done()

	file, err := s.fileHandler.WriteFile(ctx, chunkCh)
	if err != nil {
		event.SetError(err)
		return nil, err
	}

	var hash []byte
	if hash, err = s.IndexFile(ctx, file); err != nil {
		event.SetError(err)
		return nil, err
	}
	event.Append(&logging.Metadata{"hash": hex.EncodeToString(hash)})

	return hash, nil
}

// IndexFile adds the file hash and name to the db.
func (s *Storage) IndexFile(ctx context.Context, file *os.File) ([]byte, error) {
	// go back to the beginning of the file.
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}

	fileHash, err := mh.Encode(h.Sum(nil), mh.SHA2_256)
	if err != nil {
		return nil, err
	}

	if err = s.db.Put(append(prefixFilesHashes, fileHash...), []byte(file.Name())); err != nil {
		return nil, err
	}

	return fileHash, nil
}

// Authorize adds a list of peers to the authorized peers for a file hash.
func (s *Storage) Authorize(ctx context.Context, pids [][]byte, fileHash []byte) error {

	// Deserialize peer IDs.
	peerIds := make([]peer.ID, len(pids))
	for i, p := range pids {
		pid, err := peer.IDFromBytes(p)
		if err != nil {
			return err
		}
		peerIds[i] = pid
	}
	return s.acl.Authorize(ctx, peerIds, fileHash)
}

// PullFile pulls a file from a peer given the file hash.
func (s *Storage) PullFile(ctx context.Context, fileHash []byte, pid []byte) error {
	peerID, err := peer.IDFromBytes(pid)
	if err != nil {
		return err
	}
	return s.p2p.PullFile(ctx, fileHash, peerID)
}

// StreamHandler handles incoming messages from peers.
func (s *Storage) StreamHandler(ctx context.Context, stream inet.Stream) {
	from := stream.Conn().RemotePeer()
	defer stream.Close()

	event := log.EventBegin(ctx, "SendFile", logging.Metadata{
		"peerID": from.Pretty(),
	})
	defer event.Done()

	dec := protobuf.Multicodec(nil).Decoder(stream)
	enc := protobuf.Multicodec(nil).Encoder(stream)

	var req pb.FileInfo

	err := dec.Decode(&req)
	if err != nil {
		event.SetError(err)
		return
	}

	// Check that the peer is allowed to get the file.
	ok, err := s.acl.IsAuthorized(ctx, from, req.Hash)
	if err != nil {
		event.SetError(err)
		return
	}
	if !ok {
		event.SetError(ErrUnauthorized)
		return
	}

	path, err := s.getFilePath(ctx, req.Hash)
	if err != nil {
		event.SetError(err)
		return
	}

	if err := s.p2p.SendFile(ctx, enc, path); err != nil {
		event.SetError(err)
		return
	}

}

// ============================================================================
// =====													Helpers																	=====
// ============================================================================

func (s *Storage) getFilePath(ctx context.Context, fileHash []byte) (string, error) {
	p, err := s.db.Get(append(prefixFilesHashes, fileHash...))
	if err != nil {
		return "", err
	}

	return string(p), nil
}
