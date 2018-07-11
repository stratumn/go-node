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

// Package postgresauditstore implements the audit.Store interface.
// It stores links and their evidences in a PostgreSQL database.
package postgresauditstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/stratumn/go-indigocore/cs"
	"github.com/stratumn/go-indigocore/postgresstore"
	"github.com/stratumn/go-indigonode/app/indigo/protocol/store/audit"
	"github.com/stratumn/go-indigonode/app/indigo/protocol/store/constants"
	"github.com/stratumn/go-indigonode/core/monitoring"

	peer "gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
)

// PostgresAuditStore implements the audit.Store interface.
// It stores segments in a PostgreSQL database.
type PostgresAuditStore struct {
	*reader
	*writer

	db    *sql.DB
	stmts *stmts
}

// New creates a new PostgresAuditStore.
func New(config *postgresstore.Config) (*PostgresAuditStore, error) {
	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, err
	}

	return &PostgresAuditStore{
		db: db,
	}, nil
}

// AddSegment stores a segment and its signature in the DB.
func (s *PostgresAuditStore) AddSegment(ctx context.Context, segment *cs.Segment) error {
	ctx, span := monitoring.StartSpan(ctx, "indigo.store.audit", "AddSegment")
	defer span.End()

	span.AddStringAttribute("link_hash", segment.GetLinkHashString())

	peerID, err := constants.GetLinkNodeID(&segment.Link)
	if err != nil {
		span.SetUnknownError(err)
		return err
	}

	span.SetPeerID(peerID)

	lh, err := s.createLink(ctx, &segment.Link, peerID)
	if err != nil {
		span.SetUnknownError(err)
		return errors.WithStack(err)
	}

	for _, evidence := range segment.Meta.Evidences {
		if err := s.addEvidence(ctx, lh, evidence); err != nil {
			span.SetUnknownError(err)
			return err
		}
	}

	return nil
}

// GetByPeer returns segments saved in the database.
func (s *PostgresAuditStore) GetByPeer(ctx context.Context, peerID peer.ID, p audit.Pagination) (cs.SegmentSlice, error) {
	ctx, span := monitoring.StartSpan(ctx, "indigo.store.audit", "GetByPeer", monitoring.SpanOptionPeerID(peerID))
	defer span.End()

	p.InitIfInvalid()

	res, err := s.FindSegments(ctx, &audit.SegmentFilter{
		PeerID:     peerID,
		Pagination: p,
	})
	if err != nil {
		span.SetUnknownError(err)
	}

	return res, err
}

// Create creates the database tables and indexes.
// Note that the actual database itself needs to be created before calling Create().
func (s *PostgresAuditStore) Create() error {
	for _, query := range sqlCreate {
		if _, err := s.db.Exec(query); err != nil {
			return err
		}
	}
	return nil
}

// Prepare prepares the database statements.
// It should be called once before interacting with segments.
// It assumes the tables have been created using Create().
func (s *PostgresAuditStore) Prepare() error {
	stmts, err := newStmts(s.db)
	if err != nil {
		return err
	}
	s.stmts = stmts
	s.reader = &reader{stmts: s.stmts.readStmts}
	s.writer = &writer{stmts: s.stmts.writeStmts}
	return nil
}

// Drop drops the database tables and indexes.
func (s *PostgresAuditStore) Drop() error {
	for _, query := range sqlDrop {
		if _, err := s.db.Exec(query); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the database connection.
func (s *PostgresAuditStore) Close() error {
	return s.db.Close()
}
