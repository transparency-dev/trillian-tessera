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

// Package gcp contains a GCP-based storage implementation for Tessera.
//
// TODO: decide whether to rename this package.
//
// This storage implementation uses GCS for long-term storage and serving of
// entry bundles and log tiles, and Spanner for coordinating updates to GCS
// when multiple instances of a personality binary are running.
//
// A single GCS bucket is used to hold entry bundles and log internal tiles.
// The object keys for the bucket are selected so as to conform to the
// expected layout of a tile-based log.
//
// A Spanner database provides a transactional mechanism to allow multiple
// frontends to safely update the contents of the log.
package gcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/apiv1/spannerpb"
	gcs "cloud.google.com/go/storage"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/storage"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"k8s.io/klog/v2"
)

// Storage is a GCP based storage implementation for Tessera.
type Storage struct {
	gcsClient *gcs.Client

	projectID string
	bucket    string

	sequencer sequencer
	objStore  objStore

	queue *storage.Queue
}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, int64, error)
	setObject(ctx context.Context, obj string, data []byte, cond *gcs.Conditions) error
}

// coord describes a type which knows how to sequence entries.
type sequencer interface {
	assignEntries(ctx context.Context, entries [][]byte) (uint64, error)
}

// Config holds GCP project and resource configuration for a storage instance.
type Config struct {
	// ProjectID is the GCP project which hosts the storage bucket and Spanner database for the log.
	ProjectID string
	// Bucket is the name of the GCS bucket to use for storing log state.
	Bucket string
	// Spanner is the GCP resource URI of the spanner database instance to use.
	Spanner string
}

// New creates a new instance of the GCP based Storage.
func New(ctx context.Context, cfg Config) (*Storage, error) {
	c, err := gcs.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}
	gcsStorage, err := newGCSStorage(ctx, c, cfg.ProjectID, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS storage: %v", err)
	}
	seq, err := newSpannerSequencer(ctx, cfg.Spanner)
	if err != nil {
		return nil, fmt.Errorf("failed to create Spanner sequencer: %v", err)
	}

	r := &Storage{
		gcsClient: c,
		projectID: cfg.ProjectID,
		bucket:    cfg.Bucket,
		objStore:  gcsStorage,
		sequencer: seq,
	}
	// TODO(al): make queue options configurable:
	r.queue = storage.NewQueue(time.Second, 256, r.sequencer.assignEntries)

	return r, nil
}

// tileSuffix returns either the empty string or a tiles API suffix based on the passed-in tile size.
func tileSuffix(s uint64) string {
	if p := s % 256; p != 0 {
		return fmt.Sprintf(".p/%d", p)
	}
	return ""
}

// setTile stores a tile in GCS.
//
// The location to which the tile is written is defined by the tile layout spec.
func (s *Storage) setTile(ctx context.Context, level, index uint64, tile [][]byte) error {
	tileSize := uint64(len(tile))
	klog.V(2).Infof("StoreTile: level %d index %x ts: %x", level, index, tileSize)
	if tileSize == 0 || tileSize > 256 {
		return fmt.Errorf("tileSize %d must be > 0 and <= 256", tileSize)
	}
	t := slices.Concat(tile...)

	tPath := layout.TilePath(level, index) + tileSuffix(tileSize)

	return s.objStore.setObject(ctx, tPath, t, &gcs.Conditions{DoesNotExist: true})
}

// getTile returns the tile at the given tile-level and tile-index for the specified log size, or
// an error if it doesn't exist.
func (s *Storage) getTile(ctx context.Context, level, index, logSize uint64) ([][]byte, error) {
	tileSize := layout.PartialTileSize(level, index, logSize)

	objName := layout.TilePath(level, index) + tileSuffix(tileSize)
	data, _, err := s.objStore.getObject(ctx, objName)
	if err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			// Return the generic NotExist error so that higher levels can differentiate
			// between this and other errors.
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	// Recreate the tile slices
	l := len(data)
	if m := l % sha256.Size; m != 0 {
		return nil, fmt.Errorf("invalid tile data size %d", l)
	}
	t := make([][]byte, l/sha256.Size)
	for i := range t {
		t[i] = data[i*sha256.Size : (i+1)*sha256.Size]
	}
	return t, nil
}

// spannerSequencer uses Cloud Spanner to provide
// a durable and thread/multi-process safe sequencer.
type spannerSequencer struct {
	dbPool *spanner.Client
}

// new SpannerSequencer returns a new spannerSequencer struct which uses the provided
// spanner resource name for its spanner connection.
func newSpannerSequencer(ctx context.Context, spannerDB string) (*spannerSequencer, error) {
	dbPool, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Spanner: %v", err)
	}
	r := &spannerSequencer{
		dbPool: dbPool,
	}
	return r, r.initDB(ctx)
}

// initDB ensures that the coordination DB is initialised correctly.
//
// The database schema consists of 3 tables:
//   - SeqCoord
//     This table only ever contains a single row which tracks the next available
//     sequence number.
//   - Seq
//     This table holds sequenced "batches" of entries. The batches are keyed
//     by the sequence number assigned to the first entry in the batch, and
//     each subsequent entry in the batch takes the numerically next sequence number.
//   - IntCoord
//     This table coordinates integration of the batches of entries stored in
//     Seq into the committed tree state.
//
// The database and schema should be created externally, e.g. by terraform.
func (s *spannerSequencer) initDB(ctx context.Context) error {

	/* Schema for reference:
	CREATE TABLE SeqCoord (
	 id INT64 NOT NULL,
	 next INT64 NOT NULL,
	) PRIMARY KEY (id);

	CREATE TABLE Seq (
		id INT64 NOT NULL,
		seq INT64 NOT NULL,
		v BYTES(MAX),
	) PRIMARY KEY (id, seq);

	CREATE TABLE IntCoord (
		id INT64 NOT NULL,
		seq INT64 NOT NULL,
	) PRIMARY KEY (id);
	*/

	// Set default values for a newly inisialised schema - these rows being present are a precondition for
	// sequencing and integration to occur.
	// Note that this will only succeed if no row exists, so there's no danger
	// of "resetting" an existing log.
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("SeqCoord", []string{"id", "next"}, []interface{}{0, 0})}); spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("IntCoord", []string{"id", "seq"}, []interface{}{0, 0})}); spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	return nil
}

// assignEntries durably assigns each of the passed-in entries an index in the log.
//
// Entries are allocated contiguous indices, in the order in which they appear in the entries parameter.
// This is achieved by storing the passed-in entries in the Seq table in Spanner, keyed by the
// index assigned to the first entry in the batch.
func (s *spannerSequencer) assignEntries(ctx context.Context, entries [][]byte) (uint64, error) {
	// Flatted the entries into a single slice of bytes which we can store in the Seq.v column.
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	if err := e.Encode(entries); err != nil {
		return 0, fmt.Errorf("failed to serialise batch: %v", err)
	}
	data := b.Bytes()
	num := len(entries)

	var next int64 // Unfortunately, Spanner doesn't support uint64 so we'll have to cast around a bit.

	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// First we need to grab the next available sequence number from the SeqCoord table.
		row, err := txn.ReadRowWithOptions(ctx, "SeqCoord", spanner.Key{0}, []string{"id", "next"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		if err != nil {
			return err
		}
		var id int64
		if err := row.Columns(&id, &next); err != nil {
			return err
		}

		next := uint64(next) // Shadow next with a uint64 version of the same value to save on casts.
		// TODO(al): think about whether aligning bundles to tile boundaries would be a good idea or not.
		m := []*spanner.Mutation{
			// Insert our newly sequenced batch of entries into Seq,
			spanner.Insert("Seq", []string{"id", "seq", "v"}, []interface{}{0, int64(next), data}),
			// and update the next-available sequence number row in SeqCoord.
			spanner.Update("SeqCoord", []string{"id", "next"}, []interface{}{0, int64(next) + int64(num)}),
		}
		if err := txn.BufferWrite(m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to flush batch: %v", err)
	}

	return uint64(next), nil
}

// gcsStorage knows how to store and retrieve objects from GCS.
type gcsStorage struct {
	bucket    string
	gcsClient *gcs.Client
}

// newGCSStorage creates a new gcsStorage.
//
// The specified bucket must exist or an error will be returned.
func newGCSStorage(ctx context.Context, c *gcs.Client, projectID string, bucket string) (*gcsStorage, error) {
	it := c.Buckets(ctx, projectID)
	for {
		bAttrs, err := it.Next()
		if err == iterator.Done {
			return nil, fmt.Errorf("bucket %q does not exist, please create it", bucket)
		}
		if err != nil {
			return nil, err
		}
		if bAttrs.Name == bucket {
			break
		}
	}
	r := &gcsStorage{
		gcsClient: c,
		bucket:    bucket,
	}

	return r, nil
}

// getObject returns the data and generation of the specified object, or an error.
func (s *gcsStorage) getObject(ctx context.Context, obj string) ([]byte, int64, error) {
	r, err := s.gcsClient.Bucket(s.bucket).Object(obj).NewReader(ctx)
	if err != nil {
		return nil, -1, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
	}
	defer r.Close()

	d, err := io.ReadAll(r)
	return d, r.Attrs.Generation, err
}

// setObject stores the provided data in the specified object, optionally gated by a condition.
//
// cond can be used to specify preconditions for the write (e.g. write iff not exists, write iff
// current generation is X, etc.), or nil can be passed if no preconditions are desired.
//
// When preconditions are specified and are not met, an error will be returned unless the currently
// stored data is bit-for-bit identical to the data to-be-written. This is intended to provide
// idempotentency for writes.
func (s *gcsStorage) setObject(ctx context.Context, objName string, data []byte, cond *gcs.Conditions) error {
	bkt := s.gcsClient.Bucket(s.bucket)
	obj := bkt.Object(objName)

	var w *gcs.Writer
	if cond == nil {
		w = obj.NewWriter(ctx)

	} else {
		w = obj.If(*cond).NewWriter(ctx)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket, err)
	}

	if err := w.Close(); err != nil {
		// If we run into a precondition failure error, check that the object
		// which exists contains the same content that we want to write.
		// If so, we can consider this write to be idempotently successful.
		if ee, ok := err.(*googleapi.Error); ok && ee.Code == http.StatusPreconditionFailed {
			existing, existingGen, err := s.getObject(ctx, objName)
			if err != nil {
				return fmt.Errorf("failed to fetch existing content for %q (@%d): %v", objName, existingGen, err)
			}
			if !bytes.Equal(existing, data) {
				return fmt.Errorf("precondition failed: resource content for %q differs from data to-be-written", objName)
			}

			klog.V(2).Infof("setObject: identical resource already exists for %q, continuing", objName)
			return nil
		}

		return err
	}
	return nil
}
