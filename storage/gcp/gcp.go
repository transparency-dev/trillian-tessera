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
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/apiv1/spannerpb"
	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/storage"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"k8s.io/klog/v2"
)

const entryBundleSize = 256

// Storage is a GCP based storage implementation for Tessera.
type Storage struct {
	gcsClient *gcs.Client

	projectID string
	bucket    string

	newCP tessera.NewCPFunc

	sequencer sequencer
	objStore  objStore

	queue *storage.Queue
}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, int64, error)
	setObject(ctx context.Context, obj string, data []byte, cond *gcs.Conditions) error
}

// sequencer describes a type which knows how to sequence entries.
type sequencer interface {
	// assignEntries should durably allocate contiguous index numbers to the provided entries.
	assignEntries(ctx context.Context, entries []*tessera.Entry) error
	// consumeEntries should call the provided function with up to limit previously sequenced entries.
	// If the call to consumeFunc returns no error, the entries should be considered to have been consumed.
	// If any entries were successfully consumed, the implementation should also return true; this
	// serves as a weak hint that there may be more entries to be consumed.
	consumeEntries(ctx context.Context, limit uint64, f consumeFunc) (bool, error)
}

// consumeFunc is the signature of a function which can consume entries from the sequencer.
type consumeFunc func(ctx context.Context, from uint64, entries []storage.SequencedEntry) error

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
func New(ctx context.Context, cfg Config, opts ...func(*tessera.StorageOptions)) (*Storage, error) {
	opt := tessera.ResolveStorageOptions(nil, opts...)

	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}
	gcsStorage, err := newGCSStorage(ctx, c, cfg.ProjectID, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS storage: %v", err)
	}
	seq, err := newSpannerSequencer(ctx, cfg.Spanner, uint64(opt.PushbackMaxOutstanding))
	if err != nil {
		return nil, fmt.Errorf("failed to create Spanner sequencer: %v", err)
	}

	r := &Storage{
		gcsClient: c,
		projectID: cfg.ProjectID,
		bucket:    cfg.Bucket,
		objStore:  gcsStorage,
		sequencer: seq,
		newCP:     opt.NewCP,
	}
	r.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, r.sequencer.assignEntries)

	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}
			for {
				cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()
				if more, err := r.sequencer.consumeEntries(cctx, 2048 /*limit*/, r.integrate); err != nil {
					klog.Errorf("integrate: %v", err)
					break
				} else if !more {
					break
				}
				klog.V(1).Info("Quickloop")
			}
		}
	}()

	return r, nil
}

// Add is the entrypoint for adding entries to a sequencing log.
func (s *Storage) Add(ctx context.Context, e *tessera.Entry) (uint64, error) {
	return s.queue.Add(ctx, e)()
}

// Get returns the requested object.
//
// This is indended to be used to proxy read requests through the personality for debug/testing purposes.
func (s *Storage) Get(ctx context.Context, path string) ([]byte, error) {
	d, _, err := s.objStore.getObject(ctx, path)
	return d, err
}

// setTile idempotently stores the provided tile at the location implied by the given level, index, and treeSize.
//
// The location to which the tile is written is defined by the tile layout spec.
func (s *Storage) setTile(ctx context.Context, level, index, logSize uint64, tile *api.HashTile) error {
	data, err := tile.MarshalText()
	if err != nil {
		return err
	}
	tPath := layout.TilePath(level, index, logSize)
	klog.V(2).Infof("StoreTile: %s (%d entries)", tPath, len(tile.Nodes))

	return s.objStore.setObject(ctx, tPath, data, &gcs.Conditions{DoesNotExist: true})
}

// getTiles returns the tiles with the given tile-coords for the specified log size.
//
// Tiles are returned in the same order as they're requested, nils represent tiles which were not found.
func (s *Storage) getTiles(ctx context.Context, tileIDs []storage.TileID, logSize uint64) ([]*api.HashTile, error) {
	r := make([]*api.HashTile, len(tileIDs))
	errG := errgroup.Group{}
	for i, id := range tileIDs {
		i := i
		id := id
		errG.Go(func() error {
			objName := layout.TilePath(id.Level, id.Index, logSize)
			data, _, err := s.objStore.getObject(ctx, objName)
			if err != nil {
				if errors.Is(err, gcs.ErrObjectNotExist) {
					// Depending on context, this may be ok.
					// We'll signal to higher levels that it wasn't found by retuning a nil for this tile.
					return nil
				}
				return err
			}
			t := &api.HashTile{}
			if err := t.UnmarshalText(data); err != nil {
				return fmt.Errorf("unmarshal(%q): %v", objName, err)
			}
			r[i] = t
			return nil
		})
	}
	if err := errG.Wait(); err != nil {
		return nil, err
	}
	return r, nil

}

// getEntryBundle returns the serialised entry bundle at the location implied by the given index and treeSize.
//
// Returns a wrapped os.ErrNotExist if the bundle does not exist.
func (s *Storage) getEntryBundle(ctx context.Context, bundleIndex uint64, logSize uint64) ([]byte, error) {
	objName := layout.EntriesPath(bundleIndex, logSize)
	data, _, err := s.objStore.getObject(ctx, objName)
	if err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			// Return the generic NotExist error so that higher levels can differentiate
			// between this and other errors.
			return nil, fmt.Errorf("%v: %w", objName, os.ErrNotExist)
		}
		return nil, err
	}

	return data, nil
}

// setEntryBundle idempotently stores the serialised entry bundle at the location implied by the bundleIndex and treeSize.
func (s *Storage) setEntryBundle(ctx context.Context, bundleIndex uint64, logSize uint64, bundleRaw []byte) error {
	objName := layout.EntriesPath(bundleIndex, logSize)
	// Note that setObject does an idempotent interpretation of DoesNotExist - it only
	// returns an error if the named object exists _and_ contains different data to what's
	// passed in here.
	if err := s.objStore.setObject(ctx, objName, bundleRaw, &gcs.Conditions{DoesNotExist: true}); err != nil {
		return fmt.Errorf("setObject(%q): %v", objName, err)

	}
	return nil
}

// integrate incorporates the provided entries into the log starting at fromSeq.
func (s *Storage) integrate(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
	tb := storage.NewTreeBuilder(func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := s.getTiles(ctx, tileIDs, treeSize)
		if err != nil {
			return nil, fmt.Errorf("getTiles: %w", err)
		}
		return n, nil
	})

	errG := errgroup.Group{}

	errG.Go(func() error {
		if err := s.updateEntryBundles(ctx, fromSeq, entries); err != nil {
			return fmt.Errorf("updateEntryBundles: %v", err)
		}
		return nil
	})

	errG.Go(func() error {
		newSize, newRoot, tiles, err := tb.Integrate(ctx, fromSeq, entries)
		if err != nil {
			return fmt.Errorf("Integrate: %v", err)
		}
		for k, v := range tiles {
			func(ctx context.Context, k storage.TileID, v *api.HashTile) {
				errG.Go(func() error {
					return s.setTile(ctx, uint64(k.Level), k.Index, newSize, v)
				})
			}(ctx, k, v)
		}
		errG.Go(func() error {
			klog.Infof("New CP: %d, %x", newSize, newRoot)
			if s.newCP != nil {
				cpRaw, err := s.newCP(newSize, newRoot)
				if err != nil {
					return fmt.Errorf("newCP: %v", err)
				}
				if err := s.objStore.setObject(ctx, layout.CheckpointPath, cpRaw, nil); err != nil {
					switch ee := err.(type) {
					case *googleapi.Error:
						if ee.Code == http.StatusTooManyRequests {
							// TODO(al): still need to ensure it does get updated, either by the next integrate or some other mechanism.
							klog.Infof("checkpoint write rejected: %v, continuing anyway", err)
						}
					default:
						return fmt.Errorf("writeCheckpoint: %v", err)
					}
				}
			}
			return nil
		})

		return nil
	})

	return errG.Wait()
}

// updateEntryBundles adds the entries being integrated into the entry bundles.
//
// The right-most bundle will be grown, if it's partial, and/or new bundles will be created as required.
func (s *Storage) updateEntryBundles(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
	numAdded := uint64(0)
	bundleIndex, entriesInBundle := fromSeq/entryBundleSize, fromSeq%entryBundleSize
	bundleWriter := &bytes.Buffer{}
	if entriesInBundle > 0 {
		// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
		part, err := s.getEntryBundle(ctx, uint64(bundleIndex), uint64(entriesInBundle))
		if err != nil {
			return err
		}

		if _, err := bundleWriter.Write(part); err != nil {
			return fmt.Errorf("bundleWriter: %v", err)
		}
	}

	seqErr := errgroup.Group{}

	// goSetEntryBundle is a function which uses seqErr to spin off a go-routine to write out an entry bundle.
	// It's used in the for loop below.
	goSetEntryBundle := func(ctx context.Context, bundleIndex uint64, fromSeq uint64, bundleRaw []byte) {
		seqErr.Go(func() error {
			if err := s.setEntryBundle(ctx, bundleIndex, fromSeq, bundleRaw); err != nil {
				return err
			}
			return nil
		})
	}

	// Add new entries to the bundle
	for _, e := range entries {
		if _, err := bundleWriter.Write(e.BundleData); err != nil {
			return fmt.Errorf("Write: %v", err)
		}
		entriesInBundle++
		fromSeq++
		numAdded++
		if entriesInBundle == entryBundleSize {
			//  This bundle is full, so we need to write it out...
			klog.V(1).Infof("Bundle idx %d is full", bundleIndex)
			goSetEntryBundle(ctx, bundleIndex, fromSeq, bundleWriter.Bytes())
			// ... and prepare the next entry bundle for any remaining entries in the batch
			bundleIndex++
			entriesInBundle = 0
			// Don't use Reset/Truncate here - the backing []bytes is still being used by goSetEntryBundle above.
			bundleWriter = &bytes.Buffer{}
			klog.V(1).Infof("Starting bundle idx %d", bundleIndex)
		}
	}
	// If we have a partial bundle remaining once we've added all the entries from the batch,
	// this needs writing out too.
	if entriesInBundle > 0 {
		klog.V(1).Infof("Writing partial bundle idx %d.%d", bundleIndex, entriesInBundle)
		goSetEntryBundle(ctx, bundleIndex, fromSeq, bundleWriter.Bytes())
	}
	return seqErr.Wait()
}

// spannerSequencer uses Cloud Spanner to provide
// a durable and thread/multi-process safe sequencer.
type spannerSequencer struct {
	dbPool         *spanner.Client
	maxOutstanding uint64
}

// new SpannerSequencer returns a new spannerSequencer struct which uses the provided
// spanner resource name for its spanner connection.
func newSpannerSequencer(ctx context.Context, spannerDB string, maxOutstanding uint64) (*spannerSequencer, error) {
	dbPool, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Spanner: %v", err)
	}
	r := &spannerSequencer{
		dbPool:         dbPool,
		maxOutstanding: maxOutstanding,
	}
	if err := r.initDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to initDB: %v", err)
	}
	return r, nil
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
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("SeqCoord", []string{"id", "next"}, []interface{}{0, 0})}); err != nil && spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("IntCoord", []string{"id", "seq"}, []interface{}{0, 0})}); err != nil && spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	return nil
}

// assignEntries durably assigns each of the passed-in entries an index in the log.
//
// Entries are allocated contiguous indices, in the order in which they appear in the entries parameter.
// This is achieved by storing the passed-in entries in the Seq table in Spanner, keyed by the
// index assigned to the first entry in the batch.
func (s *spannerSequencer) assignEntries(ctx context.Context, entries []*tessera.Entry) error {
	// First grab the treeSize in a non-locking read-only fashion (we don't want to block/collide with integration).
	// We'll use this value to determine whether we need to apply back-pressure.
	var treeSize int64
	if row, err := s.dbPool.Single().ReadRow(ctx, "IntCoord", spanner.Key{0}, []string{"seq"}); err != nil {
		return err
	} else {
		if err := row.Column(0, &treeSize); err != nil {
			return fmt.Errorf("failed to read integration coordination info: %v", err)
		}
	}

	var next int64 // Unfortunately, Spanner doesn't support uint64 so we'll have to cast around a bit.

	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// First we need to grab the next available sequence number from the SeqCoord table.
		row, err := txn.ReadRowWithOptions(ctx, "SeqCoord", spanner.Key{0}, []string{"id", "next"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		if err != nil {
			return fmt.Errorf("failed to read SeqCoord: %v", err)
		}
		var id int64
		if err := row.Columns(&id, &next); err != nil {
			return fmt.Errorf("failed to parse id column: %v", err)
		}

		// Check whether there are too many outstanding entries and we should apply
		// back-pressure.
		if outstanding := next - treeSize; outstanding > int64(s.maxOutstanding) {
			return tessera.ErrPushback
		}

		next := uint64(next) // Shadow next with a uint64 version of the same value to save on casts.
		sequencedEntries := make([]storage.SequencedEntry, len(entries))
		// Assign provisional sequence numbers to entries.
		// We need to do this here in order to support serialisations which include the log position.
		for i, e := range entries {
			sequencedEntries[i] = storage.SequencedEntry{
				BundleData: e.MarshalBundleData(next + uint64(i)),
				LeafHash:   e.LeafHash(),
			}
		}

		// Flatten the entries into a single slice of bytes which we can store in the Seq.v column.
		b := &bytes.Buffer{}
		e := gob.NewEncoder(b)
		if err := e.Encode(sequencedEntries); err != nil {
			return fmt.Errorf("failed to serialise batch: %v", err)
		}
		data := b.Bytes()
		num := len(entries)

		// TODO(al): think about whether aligning bundles to tile boundaries would be a good idea or not.
		m := []*spanner.Mutation{
			// Insert our newly sequenced batch of entries into Seq,
			spanner.Insert("Seq", []string{"id", "seq", "v"}, []interface{}{0, int64(next), data}),
			// and update the next-available sequence number row in SeqCoord.
			spanner.Update("SeqCoord", []string{"id", "next"}, []interface{}{0, int64(next) + int64(num)}),
		}
		if err := txn.BufferWrite(m); err != nil {
			return fmt.Errorf("failed to apply TX: %v", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to flush batch: %w", err)
	}

	return nil
}

// consumeEntries calls f with previously sequenced entries.
//
// Once f returns without error, the entries it was called with are considered to have been consumed and are
// removed from the Seq table.
//
// Returns true if some entries were consumed as a weak signal that there may be further entries waiting to be consumed.
func (s *spannerSequencer) consumeEntries(ctx context.Context, limit uint64, f consumeFunc) (bool, error) {
	didWork := false
	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Figure out which is the starting index of sequenced entries to start consuming from.
		row, err := txn.ReadRowWithOptions(ctx, "IntCoord", spanner.Key{0}, []string{"seq"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		if err != nil {
			return err
		}
		var fromSeq int64 // Spanner doesn't support uint64
		if err := row.Column(0, &fromSeq); err != nil {
			return fmt.Errorf("failed to read integration coordination info: %v", err)
		}
		klog.V(1).Infof("Consuming from %d", fromSeq)

		// Now read the sequenced starting at the index we got above.
		rows := txn.ReadWithOptions(ctx, "Seq",
			spanner.KeyRange{Start: spanner.Key{0, fromSeq}, End: spanner.Key{0, fromSeq + int64(limit)}},
			[]string{"seq", "v"},
			&spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		defer rows.Stop()

		seqsConsumed := []int64{}
		entries := make([]storage.SequencedEntry, 0, limit)
		orderCheck := fromSeq
		for {
			row, err := rows.Next()
			if row == nil || err == iterator.Done {
				break
			}

			var vGob []byte
			var seq int64 // spanner doesn't have uint64
			if err := row.Columns(&seq, &vGob); err != nil {
				return fmt.Errorf("failed to scan seq row: %v", err)
			}

			if orderCheck != seq {
				return fmt.Errorf("integrity fail - expected seq %d, but found %d", orderCheck, seq)
			}

			g := gob.NewDecoder(bytes.NewReader(vGob))
			b := []storage.SequencedEntry{}
			if err := g.Decode(&b); err != nil {
				return fmt.Errorf("failed to deserialise v: %v", err)
			}
			entries = append(entries, b...)
			seqsConsumed = append(seqsConsumed, seq)
			orderCheck += int64(len(b))
		}
		if len(seqsConsumed) == 0 {
			klog.V(1).Info("Found no rows to sequence")
			return nil
		}

		// Call consumeFunc with the entries we've found
		if err := f(ctx, uint64(fromSeq), entries); err != nil {
			return err
		}

		// consumeFunc was successful, so we can update our coordination row, and delete the row(s) for
		// the then consumed entries.
		m := make([]*spanner.Mutation, 0)
		m = append(m, spanner.Update("IntCoord", []string{"id", "seq"}, []interface{}{0, int64(orderCheck)}))
		for _, c := range seqsConsumed {
			m = append(m, spanner.Delete("Seq", spanner.Key{0, c}))
		}
		if err := txn.BufferWrite(m); err != nil {
			return err
		}

		didWork = true
		return nil
	})
	if err != nil {
		return false, err
	}

	return didWork, nil
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
			return nil, fmt.Errorf("error scanning buckets: %v", err)
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

	d, err := io.ReadAll(r)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to read %q: %v", obj, err)
	}
	return d, r.Attrs.Generation, r.Close()
}

// setObject stores the provided data in the specified object, optionally gated by a condition.
//
// cond can be used to specify preconditions for the write (e.g. write iff not exists, write iff
// current generation is X, etc.), or nil can be passed if no preconditions are desired.
//
// Note that when preconditions are specified and are not met, an error will be returned *unless*
// the currently stored data is bit-for-bit identical to the data to-be-written.
// This is intended to provide idempotentency for writes.
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
				klog.Errorf("Resource %q non-idempotent write:\n%s", objName, cmp.Diff(existing, data))
				return fmt.Errorf("precondition failed: resource content for %q differs from data to-be-written", objName)
			}

			klog.V(2).Infof("setObject: identical resource already exists for %q, continuing", objName)
			return nil
		}

		return fmt.Errorf("failed to close write on %q: %v", objName, err)
	}
	return nil
}
