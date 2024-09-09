// Copyright 2024 The Tessera authors. All Rights Reserved.
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

package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/apiv1/spannerpb"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/modules/dedup"
	"google.golang.org/grpc/codes"
)

// NewDedupeStorage returns a struct which can be used to store identity -> index mappings backed
// by Spanner.
//
// Note that updates to this dedup storage is logically entriely separate from any updates
// happening to the log storage.
func NewDedupeStorage(ctx context.Context, spannerDB string) (*DedupStorage, error) {
	/*
	   Schema for reference:

	   	CREATE TABLE IDSeq (
	   	 id INT64 NOT NULL,
	   	 h BYTES(MAX) NOT NULL,
	   	 idx INT64 NOT NULL,
	   	) PRIMARY KEY (id, h);
	*/
	dedupDB, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Spanner: %v", err)
	}

	return &DedupStorage{
		dbPool: dedupDB,
	}, nil
}

// DedupStorage is a GCP Spanner based dedup storage implementation for SCTFE.
type DedupStorage struct {
	dbPool *spanner.Client
}

var _ dedup.BEDedupStorage = &DedupStorage{}

// Get looks up the stored index, if any, for the given identity.
func (d *DedupStorage) Get(ctx context.Context, i []byte) (uint64, bool, error) {
	var idx int64
	if row, err := d.dbPool.Single().ReadRow(ctx, "IDSeq", spanner.Key{0, i}, []string{"idx"}); err != nil {
		if c := spanner.ErrCode(err); c == codes.NotFound {
			return 0, false, nil
		}
		return 0, false, err
	} else {
		if err := row.Column(0, &idx); err != nil {
			return 0, false, fmt.Errorf("failed to read dedup index: %v", err)
		}
		idx := uint64(idx)
		return idx, true, nil
	}
}

// Add stores associations between the passed-in identities and their indices.
func (d *DedupStorage) Add(ctx context.Context, entries []dedup.LeafIdx) error {
	m := make([]*spanner.MutationGroup, 0, len(entries))
	for _, e := range entries {
		m = append(m, &spanner.MutationGroup{
			Mutations: []*spanner.Mutation{spanner.Insert("IDSeq", []string{"id", "h", "idx"}, []interface{}{0, e.LeafID, int64(e.Idx)})},
		})
	}

	i := d.dbPool.BatchWrite(ctx, m)
	return i.Do(func(r *spannerpb.BatchWriteResponse) error {
		s := r.GetStatus()
		if c := codes.Code(s.Code); c != codes.OK && c != codes.AlreadyExists {
			return fmt.Errorf("failed to write dedup record: %v (%v)", s.GetMessage(), c)
		}
		return nil
	})
}
