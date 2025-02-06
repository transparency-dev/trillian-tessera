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
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"cloud.google.com/go/spanner/apiv1/spannerpb"

	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/options"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"k8s.io/klog/v2"
)

const (
	// minCheckpointInterval is the shortest permitted interval between updating published checkpoints.
	// GCS has a rate limit 1 update per second for individual objects, but we've observed that attempting
	// to update at exactly that rate still results in the occasional refusal, so bake in a little wiggle
	// room.
	minCheckpointInterval = 1200 * time.Millisecond

	logContType      = "application/octet-stream"
	ckptContType     = "text/plain; charset=utf-8"
	logCacheControl  = "max-age=604800,immutable"
	ckptCacheControl = "no-cache"

	DefaultPushbackMaxOutstanding = 4096
	DefaultIntegrationSizeLimit   = 5 * 4096

	// SchemaCompatibilityVersion represents the expected version (e.g. layout & serialisation) of stored data.
	//
	// A binary built with a given version of the Tessera library is compatible with stored data created by a different version
	// of the library if and only if this value is the same as the compatibilityVersion stored in the Tessera table.
	//
	// NOTE: if changing this version, you need to consider whether end-users are going to update their schema instances to be
	// compatible with the new format, and provide a means to do it if so.
	SchemaCompatibilityVersion = 1
)

// Storage is a GCP based storage implementation for Tessera.
type Storage struct {
	newCP       options.NewCPFunc
	entriesPath options.EntriesPathFunc

	sequencer sequencer
	objStore  objStore

	queue *storage.Queue

	cpUpdated chan struct{}
}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, int64, error)
	setObject(ctx context.Context, obj string, data []byte, cond *gcs.Conditions, contType string, cacheCtl string) error
	lastModified(ctx context.Context, obj string) (time.Time, error)
}

// sequencer describes a type which knows how to sequence entries.
type sequencer interface {
	// assignEntries should durably allocate contiguous index numbers to the provided entries.
	assignEntries(ctx context.Context, entries []*tessera.Entry) error
	// consumeEntries should call the provided function with up to limit previously sequenced entries.
	// If the call to consumeFunc returns no error, the entries should be considered to have been consumed.
	// If any entries were successfully consumed, the implementation should also return true; this
	// serves as a weak hint that there may be more entries to be consumed.
	// If forceUpdate is true, then the consumeFunc should be called, with an empty slice of entries if
	// necessary. This allows the log self-initialise in a transactionally safe manner.
	consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error)
	// currentTree returns the sequencer's view of the current tree state.
	currentTree(ctx context.Context) (uint64, []byte, error)
}

// consumeFunc is the signature of a function which can consume entries from the sequencer and integrate
// them into the log.
// Returns the new rootHash once all passed entries have been integrated.
type consumeFunc func(ctx context.Context, from uint64, entries []storage.SequencedEntry) ([]byte, error)

// Config holds GCP project and resource configuration for a storage instance.
type Config struct {
	// Bucket is the name of the GCS bucket to use for storing log state.
	Bucket string
	// Spanner is the GCP resource URI of the spanner database instance to use.
	Spanner string
}

// New creates a new instance of the GCP based Storage.
func New(ctx context.Context, cfg Config, opts ...func(*options.StorageOptions)) (tessera.Driver, error) {
	opt := storage.ResolveStorageOptions(opts...)
	if opt.PushbackMaxOutstanding == 0 {
		opt.PushbackMaxOutstanding = DefaultPushbackMaxOutstanding
	}
	if opt.CheckpointInterval < minCheckpointInterval {
		return nil, fmt.Errorf("requested CheckpointInterval (%v) is less than minimum permitted %v", opt.CheckpointInterval, minCheckpointInterval)
	}

	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	seq, err := newSpannerSequencer(ctx, cfg.Spanner, uint64(opt.PushbackMaxOutstanding))
	if err != nil {
		return nil, fmt.Errorf("failed to create Spanner sequencer: %v", err)
	}

	r := &Storage{
		objStore: &gcsStorage{
			gcsClient: c,
			bucket:    cfg.Bucket,
		},
		sequencer:   seq,
		newCP:       opt.NewCP,
		entriesPath: opt.EntriesPath,
		cpUpdated:   make(chan struct{}),
	}
	r.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, r.sequencer.assignEntries)

	if err := r.init(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialise log storage: %v", err)
	}

	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}

			func() {
				// Don't quickloop for now, it causes issues updating checkpoint too frequently.
				cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				if _, err := r.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, r.appendEntries, false); err != nil {
					klog.Errorf("integrate: %v", err)
					return
				}
				select {
				case r.cpUpdated <- struct{}{}:
				default:
				}
			}()
		}
	}()

	go func(ctx context.Context, i time.Duration) {
		t := time.NewTicker(i)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-r.cpUpdated:
			case <-t.C:
			}
			if err := r.publishCheckpoint(ctx, i); err != nil {
				klog.Warningf("publishCheckpoint: %v", err)
			}
		}
	}(ctx, opt.CheckpointInterval)

	return r, nil
}

// Add is the entrypoint for adding entries to a sequencing log.
func (s *Storage) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	return s.queue.Add(ctx, e)
}

func (s *Storage) ReadCheckpoint(ctx context.Context) ([]byte, error) {
	return s.get(ctx, layout.CheckpointPath)
}

func (s *Storage) ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error) {
	return s.get(ctx, layout.TilePath(l, i, p))
}

func (s *Storage) ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error) {
	return s.get(ctx, s.entriesPath(i, p))
}

// get returns the requested object.
//
// This is indended to be used to proxy read requests through the personality for debug/testing purposes.
func (s *Storage) get(ctx context.Context, path string) ([]byte, error) {
	d, _, err := s.objStore.getObject(ctx, path)
	return d, err
}

// init ensures that the storage represents a log in a valid state.
func (s *Storage) init(ctx context.Context) error {
	_, err := s.get(ctx, layout.CheckpointPath)
	if err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			// No checkpoint exists, do a forced (possibly empty) integration to create one in a safe
			// way (setting the checkpoint directly here would not be safe as it's outside the transactional
			// framework which prevents the tree from rolling backwards or otherwise forking).
			cctx, c := context.WithTimeout(ctx, 10*time.Second)
			defer c()
			if _, err := s.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, s.appendEntries, true); err != nil {
				return fmt.Errorf("forced integrate: %v", err)
			}
			select {
			case s.cpUpdated <- struct{}{}:
			default:
			}
			return nil
		}
		return fmt.Errorf("failed to read checkpoint: %v", err)
	}

	return nil
}

func (s *Storage) publishCheckpoint(ctx context.Context, minStaleness time.Duration) error {
	m, err := s.objStore.lastModified(ctx, layout.CheckpointPath)
	if err != nil && !errors.Is(err, gcs.ErrObjectNotExist) {
		return fmt.Errorf("lastModified(%q): %v", layout.CheckpointPath, err)
	}
	if time.Since(m) < minStaleness {
		return nil
	}

	size, root, err := s.sequencer.currentTree(ctx)
	if err != nil {
		return fmt.Errorf("currentTree: %v", err)
	}
	cpRaw, err := s.newCP(size, root)
	if err != nil {
		return fmt.Errorf("newCP: %v", err)
	}

	if err := s.objStore.setObject(ctx, layout.CheckpointPath, cpRaw, nil, ckptContType, ckptCacheControl); err != nil {
		return fmt.Errorf("writeCheckpoint: %v", err)
	}
	return nil

}

// setTile idempotently stores the provided tile at the location implied by the given level, index, and treeSize.
//
// The location to which the tile is written is defined by the tile layout spec.
func (s *Storage) setTile(ctx context.Context, level, index uint64, partial uint8, data []byte) error {
	tPath := layout.TilePath(level, index, partial)
	return s.objStore.setObject(ctx, tPath, data, &gcs.Conditions{DoesNotExist: true}, logContType, logCacheControl)
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
			objName := layout.TilePath(id.Level, id.Index, layout.PartialTileSize(id.Level, id.Index, logSize))
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

func (s *Storage) State(ctx context.Context) (uint64, error) {
	size, _, err := s.sequencer.currentTree(ctx)
	return size, err
}

// StreamEntryRange provides a mechanism to quickly read sequential entry bundles covering a range of entries [fromEntry, fromEntry+N).
//
// Returns a "next" function, which can be used to retrieve subsequent entry bundles, and a "stop" function which must be
// called when no further bundles are required.
//
// This implementation is intended to be relatively performant compared to the naive approach of serially fetching
// and yielding each bundle in turn, and is intended for use cases where the caller is actively consuming large
// sections of the log contents.
//
// TODO(al): consider whether this should be factored out as a storage mix-in.
func (s *Storage) StreamEntryRange(ctx context.Context, fromEntry, N, treeSize uint64) (func() (ri layout.RangeInfo, bundle []byte, err error), func()) {
	klog.V(1).Infof("StreamEntryRange from %d, N %d, treeSize %d", fromEntry, N, treeSize)

	// bundleOrErr represents a fetched entry bunlde and its params, or an error if we couldn't fetch it for
	// some reason.
	type bundleOrErr struct {
		ri  layout.RangeInfo
		b   []byte
		err error
	}
	// TODO(al): this should probably be configurable
	nWorkers := 10

	// bundles will be filled with futures for in-order entry bundles by the worker
	// go routines below.
	// This channel will be drained by the loop at the bottom of this func which
	// yields the bundles to the caller.
	bundles := make(chan func() bundleOrErr, nWorkers)
	exit := make(chan struct{})

	// Fetch entry bundle resources in parallel.
	// We use a limited number of tokens here to prevent this from
	// consuming an unbounded amount of resources.
	go func() {
		defer close(bundles)

		// We'll limit ourselves to nWorkers worth of on-going work using these tokens:
		tokens := make(chan struct{}, nWorkers)
		for range nWorkers {
			tokens <- struct{}{}
		}

		// For each bundle, pop a future into the bundles channel and kick off an async request
		// to resolve it.
		for ri := range layout.Range(fromEntry, N, treeSize) {
			select {
			case <-exit:
				return
			case <-tokens:
				// We'll return a token below, once the bundle is fetched _and_ is being yielded.
			}

			c := make(chan bundleOrErr, 1)
			go func(ri layout.RangeInfo) {
				b, err := s.getEntryBundle(ctx, ri.Index, ri.Partial)
				c <- bundleOrErr{ri: ri, b: b, err: err}
			}(ri)

			f := func() bundleOrErr {
				b := <-c
				// We're about to yield a value, so we can now return the token and unblock another fetch.
				tokens <- struct{}{}
				return b
			}

			bundles <- f
		}
	}()

	stop := func() {
		close(exit)
	}

	next := func() (layout.RangeInfo, []byte, error) {
		select {
		case f, ok := <-bundles:
			if !ok {
				return layout.RangeInfo{}, nil, errors.New("no more bundles")
			}
			b := f()
			return b.ri, b.b, b.err
		}
	}
	return next, stop
}

// getEntryBundle returns the serialised entry bundle at the location described by the given index and partial size.
// A partial size of zero implies a full tile.
//
// Returns a wrapped os.ErrNotExist if the bundle does not exist.
func (s *Storage) getEntryBundle(ctx context.Context, bundleIndex uint64, p uint8) ([]byte, error) {
	objName := s.entriesPath(bundleIndex, p)
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
func (s *Storage) setEntryBundle(ctx context.Context, bundleIndex uint64, p uint8, bundleRaw []byte) error {
	objName := s.entriesPath(bundleIndex, p)
	// Note that setObject does an idempotent interpretation of DoesNotExist - it only
	// returns an error if the named object exists _and_ contains different data to what's
	// passed in here.
	if err := s.objStore.setObject(ctx, objName, bundleRaw, &gcs.Conditions{DoesNotExist: true}, logContType, logCacheControl); err != nil {
		return fmt.Errorf("setObject(%q): %v", objName, err)

	}
	return nil
}

// appendEntries incorporates the provided entries into the log starting at fromSeq.
func (s *Storage) appendEntries(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
	var newRoot []byte

	errG := errgroup.Group{}

	errG.Go(func() error {
		if err := s.updateEntryBundles(ctx, fromSeq, entries); err != nil {
			return fmt.Errorf("updateEntryBundles: %v", err)
		}
		return nil
	})

	errG.Go(func() error {
		lh := make([][]byte, len(entries))
		for i, e := range entries {
			lh[i] = e.LeafHash
		}
		r, err := s.integrate(ctx, fromSeq, lh)
		if err != nil {
			return fmt.Errorf("integrate: %v", err)
		}
		newRoot = r
		return nil
	})
	if err := errG.Wait(); err != nil {
		return nil, err
	}
	return newRoot, nil
}

// integrate adds the provided leaf hashes to the merkle tree, starting at the provided location.
func (s *Storage) integrate(ctx context.Context, fromSeq uint64, lh [][]byte) ([]byte, error) {
	errG := errgroup.Group{}
	getTiles := func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := s.getTiles(ctx, tileIDs, treeSize)
		if err != nil {
			return nil, fmt.Errorf("getTiles: %w", err)
		}
		return n, nil
	}

	newSize, newRoot, tiles, err := storage.Integrate(ctx, getTiles, fromSeq, lh)
	if err != nil {
		return nil, fmt.Errorf("Integrate: %v", err)
	}
	for k, v := range tiles {
		func(ctx context.Context, k storage.TileID, v *api.HashTile) {
			errG.Go(func() error {
				data, err := v.MarshalText()
				if err != nil {
					return err
				}
				return s.setTile(ctx, k.Level, k.Index, layout.PartialTileSize(k.Level, k.Index, newSize), data)
			})
		}(ctx, k, v)
	}
	if err := errG.Wait(); err != nil {
		return nil, err
	}
	klog.Infof("New tree: %d, %x", newSize, newRoot)

	return newRoot, nil
}

// updateEntryBundles adds the entries being integrated into the entry bundles.
//
// The right-most bundle will be grown, if it's partial, and/or new bundles will be created as required.
func (s *Storage) updateEntryBundles(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
	if len(entries) == 0 {
		return nil
	}

	numAdded := uint64(0)
	bundleIndex, entriesInBundle := fromSeq/layout.EntryBundleWidth, fromSeq%layout.EntryBundleWidth
	bundleWriter := &bytes.Buffer{}
	if entriesInBundle > 0 {
		// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
		part, err := s.getEntryBundle(ctx, uint64(bundleIndex), uint8(entriesInBundle))
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
	goSetEntryBundle := func(ctx context.Context, bundleIndex uint64, p uint8, bundleRaw []byte) {
		seqErr.Go(func() error {
			if err := s.setEntryBundle(ctx, bundleIndex, p, bundleRaw); err != nil {
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
		if entriesInBundle == layout.EntryBundleWidth {
			//  This bundle is full, so we need to write it out...
			klog.V(1).Infof("In-memory bundle idx %d is full, attempting write to GCS", bundleIndex)
			goSetEntryBundle(ctx, bundleIndex, 0, bundleWriter.Bytes())
			// ... and prepare the next entry bundle for any remaining entries in the batch
			bundleIndex++
			entriesInBundle = 0
			// Don't use Reset/Truncate here - the backing []bytes is still being used by goSetEntryBundle above.
			bundleWriter = &bytes.Buffer{}
			klog.V(1).Infof("Starting to fill in-memory bundle idx %d", bundleIndex)
		}
	}
	// If we have a partial bundle remaining once we've added all the entries from the batch,
	// this needs writing out too.
	if entriesInBundle > 0 {
		klog.V(1).Infof("Attempting to write in-memory partial bundle idx %d.%d to GCS", bundleIndex, entriesInBundle)
		goSetEntryBundle(ctx, bundleIndex, uint8(entriesInBundle), bundleWriter.Bytes())
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
	if err := r.initDB(ctx, spannerDB); err != nil {
		return nil, fmt.Errorf("failed to initDB: %v", err)
	}
	if err := r.checkDataCompatibility(ctx); err != nil {
		return nil, fmt.Errorf("schema is not compatible with this version of the Tessera library: %v", err)
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
func (s *spannerSequencer) initDB(ctx context.Context, spannerDB string) error {
	return createAndPrepareTables(ctx, spannerDB,
		[]string{
			"CREATE TABLE IF NOT EXISTS Tessera (id INT64 NOT NULL, compatibilityVersion INT64 NOT NULL) PRIMARY KEY (id)",
			"CREATE TABLE IF NOT EXISTS SeqCoord (id INT64 NOT NULL, next INT64 NOT NULL,) PRIMARY KEY (id)",
			"CREATE TABLE IF NOT EXISTS Seq (id INT64 NOT NULL, seq INT64 NOT NULL, v BYTES(MAX),) PRIMARY KEY (id, seq)",
			"CREATE TABLE IF NOT EXISTS IntCoord (id INT64 NOT NULL, seq INT64 NOT NULL, rootHash BYTES(32)) PRIMARY KEY (id)",
		},
		[][]*spanner.Mutation{
			{spanner.Insert("Tessera", []string{"id", "compatibilityVersion"}, []interface{}{0, SchemaCompatibilityVersion})},
			{spanner.Insert("SeqCoord", []string{"id", "next"}, []interface{}{0, 0})},
			{spanner.Insert("IntCoord", []string{"id", "seq", "rootHash"}, []interface{}{0, 0, rfc6962.DefaultHasher.EmptyRoot()})},
		},
	)
}

// checkDataCompatibility compares the Tessera library SchemaCompatibilityVersion with the one stored in the
// database, and returns an error if they are not identical.
func (s *spannerSequencer) checkDataCompatibility(ctx context.Context) error {
	row, err := s.dbPool.Single().ReadRow(ctx, "Tessera", spanner.Key{0}, []string{"compatibilityVersion"})
	if err != nil {
		return fmt.Errorf("failed to read schema compatibilityVersion: %v", err)
	}
	var compat int64
	if err := row.Columns(&compat); err != nil {
		return fmt.Errorf("failed to scan schema compatibilityVersion: %v", err)
	}

	if compat != SchemaCompatibilityVersion {
		return fmt.Errorf("schema compatibilityVersion (%d) != library compatibilityVersion (%d)", compat, SchemaCompatibilityVersion)
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
func (s *spannerSequencer) consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error) {
	didWork := false
	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Figure out which is the starting index of sequenced entries to start consuming from.
		row, err := txn.ReadRowWithOptions(ctx, "IntCoord", spanner.Key{0}, []string{"seq", "rootHash"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		if err != nil {
			return err
		}
		var fromSeq int64 // Spanner doesn't support uint64
		var rootHash []byte
		if err := row.Columns(&fromSeq, &rootHash); err != nil {
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
		if len(seqsConsumed) == 0 && !forceUpdate {
			klog.V(1).Info("Found no rows to sequence")
			return nil
		}

		// Call consumeFunc with the entries we've found
		newRoot, err := f(ctx, uint64(fromSeq), entries)
		if err != nil {
			return err
		}

		// consumeFunc was successful, so we can update our coordination row, and delete the row(s) for
		// the then consumed entries.
		m := make([]*spanner.Mutation, 0)
		m = append(m, spanner.Update("IntCoord", []string{"id", "seq", "rootHash"}, []interface{}{0, int64(orderCheck), newRoot}))
		for _, c := range seqsConsumed {
			m = append(m, spanner.Delete("Seq", spanner.Key{0, c}))
		}
		if len(m) > 0 {
			if err := txn.BufferWrite(m); err != nil {
				return err
			}
		}

		didWork = true
		return nil
	})
	if err != nil {
		return false, err
	}

	return didWork, nil
}

// currentTree returns the size and root hash of the currently integrated tree.
func (s *spannerSequencer) currentTree(ctx context.Context) (uint64, []byte, error) {
	row, err := s.dbPool.Single().ReadRow(ctx, "IntCoord", spanner.Key{0}, []string{"seq", "rootHash"})
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read IntCoord: %v", err)
	}
	var fromSeq int64 // Spanner doesn't support uint64
	var rootHash []byte
	if err := row.Columns(&fromSeq, &rootHash); err != nil {
		return 0, nil, fmt.Errorf("failed to read integration coordination info: %v", err)
	}

	return uint64(fromSeq), rootHash, nil
}

// gcsStorage knows how to store and retrieve objects from GCS.
type gcsStorage struct {
	bucket    string
	gcsClient *gcs.Client
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
func (s *gcsStorage) setObject(ctx context.Context, objName string, data []byte, cond *gcs.Conditions, contType string, cacheCtl string) error {
	bkt := s.gcsClient.Bucket(s.bucket)
	obj := bkt.Object(objName)

	var w *gcs.Writer
	if cond == nil {
		w = obj.NewWriter(ctx)

	} else {
		w = obj.If(*cond).NewWriter(ctx)
	}
	w.ObjectAttrs.ContentType = contType
	w.ObjectAttrs.CacheControl = cacheCtl
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

func (s *gcsStorage) lastModified(ctx context.Context, obj string) (time.Time, error) {
	r, err := s.gcsClient.Bucket(s.bucket).Object(obj).NewReader(ctx)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
	}
	return r.Attrs.LastModified, r.Close()
}

// NewDedupe returns a dedupe driver which uses Spanner to maintain a mapping of
// previously seen entries and their assigned indices.
//
// Note that the storage for this mapping is entirely separate and unconnected to the storage used for
// maintaining the Merkle tree.
//
// This functionality is experimental!
func NewDedupe(ctx context.Context, spannerDB string) (*DedupStorage, error) {
	/*
		Schema for reference:
			CREATE TABLE FollowCoord (
				id INT64 NOT NULL,
				nextIdx INT64 NOT NULL,
			) PRIMARY KEY (id);

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

	// Initialise the DB if necessary
	if _, err := dedupDB.Apply(ctx, []*spanner.Mutation{spanner.Insert("FollowCoord", []string{"id", "nextIdx"}, []interface{}{0, 0})}); err != nil && spanner.ErrCode(err) != codes.AlreadyExists {
		return nil, fmt.Errorf("failed to initialise dedupDB: %v:", err)
	}

	r := &DedupStorage{
		ctx:    ctx,
		dbPool: dedupDB,
	}

	go func(ctx context.Context) {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				klog.V(1).Infof("DEDUP: # Writes %d, # Lookups %d, # DB hits %v", r.numWrites.Load(), r.numLookups.Load(), r.numDBDedups.Load())
			}
		}
	}(ctx)

	return r, nil
}

type DedupStorage struct {
	ctx    context.Context
	dbPool *spanner.Client

	numLookups  atomic.Uint64
	numWrites   atomic.Uint64
	numDBDedups atomic.Uint64
}

// index returns the index (if any) previously associated with the provided hash
func (d *DedupStorage) index(ctx context.Context, h []byte) (*uint64, error) {
	d.numLookups.Add(1)
	var idx int64
	if row, err := d.dbPool.Single().ReadRow(ctx, "IDSeq", spanner.Key{0, h}, []string{"idx"}); err != nil {
		if c := spanner.ErrCode(err); c == codes.NotFound {
			return nil, nil
		}
		return nil, err
	} else {
		if err := row.Column(0, &idx); err != nil {
			return nil, fmt.Errorf("failed to read dedup index: %v", err)
		}
		idx := uint64(idx)
		d.numDBDedups.Add(1)
		return &idx, nil
	}
}

// storeMappings stores the associations between the keys and IDs in a non-atomic fashion
// (i.e. it does not store all or none in a transactional sense).
//
// Returns an error if one or more mappings cannot be stored.
func (d *DedupStorage) storeMappings(ctx context.Context, entries []dedupeMapping) error {
	m := make([]*spanner.MutationGroup, 0, len(entries))
	for _, e := range entries {
		m = append(m, &spanner.MutationGroup{
			Mutations: []*spanner.Mutation{spanner.Insert("IDSeq", []string{"id", "h", "idx"}, []interface{}{0, e.ID, int64(e.Idx)})},
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

// dedupeMapping represents an ID -> index mapping.
type dedupeMapping struct {
	ID  []byte
	Idx uint64
}

// Decorator returns a function which will wrap an underlying Add delegate with
// code to dedup against the stored data.
func (d *DedupStorage) Decorator() func(f tessera.AddFn) tessera.AddFn {
	return func(delegate tessera.AddFn) tessera.AddFn {
		return func(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
			idx, err := d.index(ctx, e.Identity())
			if err != nil {
				return func() (uint64, error) { return 0, err }
			}
			if idx != nil {
				return func() (uint64, error) { return *idx, nil }
			}

			return delegate(ctx, e)
		}
	}
}

func NewEntryReader(next func() (layout.RangeInfo, []byte, error), bundleFn BundleHasherFunc) *EntryReader {
	return &EntryReader{
		bundleFn: bundleFn,
		next:     next,
		i:        0,
	}
}

type EntryReader struct {
	bundleFn BundleHasherFunc
	next     func() (layout.RangeInfo, []byte, error)

	curData [][]byte
	curRI   layout.RangeInfo
	i       uint64
}

func (e *EntryReader) Next() (uint64, []byte, error) {
	if len(e.curData) == 0 {
		var err error
		var b []byte
		e.curRI, b, err = e.next()
		if err != nil {
			return 0, nil, fmt.Errorf("next: %v", err)
		}
		e.curData, err = e.bundleFn(b)
		if err != nil {
			return 0, nil, fmt.Errorf("bundleFn(bundleEntry @%d): %v", e.curRI.Index, err)

		}
		if e.curRI.First > 0 {
			e.curData = e.curData[e.curRI.First:]
		}
		if len(e.curData) > int(e.curRI.N) {
			e.curData = e.curData[:e.curRI.N]
		}
		e.i = 0
	}
	var r []byte
	r, e.curData = e.curData[0], e.curData[1:]
	rIdx := e.curRI.Index*layout.EntryBundleWidth + uint64(e.curRI.First) + e.i
	e.i++
	return rIdx, r, nil
}

// LogFollower provides read-only access to the log with an API tailored to bulk in-order
// reads of entry bundles.
//
// TODO(al): factor this out into higher layer when it's ready.
type LogFollower interface {
	// State returns the size of the currently integrated tree.
	// Note that this _may_ be larger than the currently _published_ checkpoint.
	State(ctx context.Context) (uint64, error)

	// StreamEntryBundles returns functions which act like a pull iterator for subsequent entry bundles starting at the given index.
	//
	// Implementations must:
	//  - truncate the requested range if any or all of it is beyond the extent of the currently integrated tree.
	//  - cease iterating if next() produces an error, or stop is called. next should continue to return an error if called again after either of these cases.
	StreamEntryRange(ctx context.Context, fromIdx, N, treeSize uint64) (next func() (layout.RangeInfo, []byte, error), stop func())
}

// Populate uses entry data from the log to populate the dedupe storage.
//
// TODO(al):  add details
func (d *DedupStorage) Populate(ctx context.Context, lf LogFollower, bundleFn BundleHasherFunc) error {
	t := time.NewTicker(time.Second)
	var (
		entryReader *EntryReader
		stop        func()

		curEntries [][]byte
		curIndex   uint64
	)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
		}
		size, err := lf.State(ctx)
		if err != nil {
			klog.Errorf("Populate: State(): %v", err)
			continue
		}

		// Busy loop while there's work to be done
		for workDone := true; workDone; {
			_, err = d.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
				// Figure out the last entry we used to populate our dedup storage.
				row, err := txn.ReadRowWithOptions(ctx, "FollowCoord", spanner.Key{0}, []string{"nextIdx"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
				if err != nil {
					return err
				}

				var f int64 // Spanner doesn't support uint64
				if err := row.Columns(&f); err != nil {
					return fmt.Errorf("failed to read follow coordination info: %v", err)
				}
				followFrom := uint64(f)
				if followFrom == size {
					workDone = false
					return nil
				}

				// If this is the first time around the loop we need to start the stream of entries now that we know where we want to
				// start reading from:
				if entryReader == nil {
					next, st := lf.StreamEntryRange(ctx, followFrom, size-followFrom, size)
					stop = st
					entryReader = NewEntryReader(next, bundleFn)
				}

				if curIndex == followFrom && curEntries != nil {
					// We're recovering from a previous failed attempt
				} else {
					const batchSize = 64
					bs := uint64(batchSize)
					if r := size - followFrom; r < bs {
						bs = r
					}
					batch := make([][]byte, 0, bs)
					for i := 0; i < int(bs); i++ {
						idx, c, err := entryReader.Next()
						if err != nil {
							return fmt.Errorf("entryReader.next: %v", err)
						}
						if wantIdx := followFrom + uint64(i); idx != wantIdx {
							// We're out of sync
							return fmt.Errorf("out of sync (%d != %d), bailing", idx, wantIdx)
						}
						batch = append(batch, c)
					}
					curEntries = batch
					curIndex = followFrom
				}

				// Store dedup entries.
				//
				// Note that we're writing the dedup entries outside of the transaction here.
				// This looks weird, but is ok because we don't mind if one or more of the individual dedupe entries fails because there's already
				// an entry for that hash, and we'll only continue on to update our FollowCoord if no errors (other than AlreadyExists) occur while
				// inserting entries.
				//
				// Alternative approaches are:
				//  - Use InsertOrUpdate, but we'd rather keep serving the earliest index known for that entry.
				//  - Perform Reads for each of the hashes we're about to write, and use that to filter writes.
				//    This would work, but would incur an extra round-trip of data which isn't really necessary.
				{
					m := make([]*spanner.MutationGroup, 0, len(curEntries))
					for i, e := range curEntries {
						m = append(m, &spanner.MutationGroup{
							Mutations: []*spanner.Mutation{spanner.Insert("IDSeq", []string{"id", "h", "idx"}, []interface{}{0, e, int64(curIndex + uint64(i))})},
						})
					}

					i := d.dbPool.BatchWrite(ctx, m)
					err := i.Do(func(r *spannerpb.BatchWriteResponse) error {
						s := r.GetStatus()
						if c := codes.Code(s.Code); c != codes.OK && c != codes.AlreadyExists {
							return fmt.Errorf("failed to write dedup record: %v (%v)", s.GetMessage(), c)
						}
						return nil
					})
					if err != nil {
						return err
					}
				}

				numAdded := uint64(len(curEntries))
				d.numWrites.Add(numAdded)

				// Insertion of dupe entries was successful, so update our follow coordination row:
				m := make([]*spanner.Mutation, 0)
				m = append(m, spanner.Update("FollowCoord", []string{"id", "nextIdx"}, []interface{}{0, int64(followFrom + numAdded)}))

				return txn.BufferWrite(m)
			})
			if err != nil {
				klog.Errorf("Failed to commit dedupe population tx: %v", err)
				stop()
				entryReader = nil
				continue
			}
			curEntries = nil
		}
	}
}

// BundleHasherFunc is the signature of a function which knows how to parse an entry bundle and calculate leaf hashes for its entries.
type BundleHasherFunc func(entryBundle []byte) (LeafHashes [][]byte, err error)

// NewMigrationTarget creates a new GCP storage for the MigrationTarget lifecycle mode.
func NewMigrationTarget(ctx context.Context, cfg Config, bundleHasher BundleHasherFunc, opts ...func(*options.StorageOptions)) (*MigrationStorage, error) {
	opt := storage.ResolveStorageOptions(opts...)
	if opt.PushbackMaxOutstanding == 0 {
		opt.PushbackMaxOutstanding = DefaultPushbackMaxOutstanding
	}

	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	seq, err := newSpannerSequencer(ctx, cfg.Spanner, uint64(opt.PushbackMaxOutstanding))
	if err != nil {
		return nil, fmt.Errorf("failed to create Spanner sequencer: %v", err)
	}

	r := &Storage{
		objStore: &gcsStorage{
			gcsClient: c,
			bucket:    cfg.Bucket,
		},
		sequencer:   seq,
		newCP:       opt.NewCP,
		entriesPath: opt.EntriesPath,
	}

	m := &MigrationStorage{
		s:            r,
		dbPool:       seq.dbPool,
		bundleHasher: bundleHasher,
	}

	return m, nil
}

type MigrationStorage struct {
	s            *Storage
	dbPool       *spanner.Client
	bundleHasher BundleHasherFunc
}

func (m *MigrationStorage) AwaitIntegration(ctx context.Context, sourceSize uint64) ([]byte, error) {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
			from, _, err := m.s.sequencer.currentTree(ctx)
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				klog.Warningf("readTreeState: %v", err)
				continue
			}
			klog.Infof("Integrate from %d (Target %d)", from, sourceSize)
			newSize, newRoot, err := m.buildTree(ctx, sourceSize)
			if err != nil {
				klog.Warningf("integrate: %v", err)
			}
			if newSize == sourceSize {
				klog.Infof("Integrated to %d with roothash %x", newSize, newRoot)
				return newRoot, nil
			}
		}
	}
}

func (m *MigrationStorage) SetEntryBundle(ctx context.Context, index uint64, partial uint8, bundle []byte) error {
	return m.s.setEntryBundle(ctx, index, partial, bundle)
}

func (m *MigrationStorage) State(ctx context.Context) (uint64, []byte, error) {
	return m.s.sequencer.currentTree(ctx)
}

func (m *MigrationStorage) fetchLeafHashes(ctx context.Context, from, to, sourceSize uint64) ([][]byte, error) {
	// TODO(al): Make this configurable.
	const maxBundles = 300

	toBeAdded := sync.Map{}
	eg := errgroup.Group{}
	n := 0
	for ri := range layout.Range(from, to, sourceSize) {
		eg.Go(func() error {
			b, err := m.s.ReadEntryBundle(ctx, ri.Index, ri.Partial)
			if err != nil {
				return fmt.Errorf("ReadEntryBundle(%d.%d): %v", ri.Index, ri.Partial, err)
			}

			bh, err := m.bundleHasher(b)
			if err != nil {
				return fmt.Errorf("bundleHasherFunc for bundle index %d: %v", ri.Index, err)
			}
			toBeAdded.Store(ri.Index, bh[ri.First:ri.First+ri.N])
			return nil
		})
		n++
		if n >= maxBundles {
			break
		}
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	lh := make([][]byte, 0, maxBundles)
	for i := from / layout.EntryBundleWidth; ; i++ {
		v, ok := toBeAdded.LoadAndDelete(i)
		if !ok {
			break
		}
		bh := v.([][]byte)
		lh = append(lh, bh...)
	}

	return lh, nil
}

func (m *MigrationStorage) buildTree(ctx context.Context, sourceSize uint64) (uint64, []byte, error) {
	var newSize uint64
	var newRoot []byte

	_, err := m.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Figure out which is the starting index of sequenced entries to start consuming from.
		row, err := txn.ReadRowWithOptions(ctx, "IntCoord", spanner.Key{0}, []string{"seq", "rootHash"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
		if err != nil {
			return err
		}
		var fromSeq int64 // Spanner doesn't support uint64
		var rootHash []byte
		if err := row.Columns(&fromSeq, &rootHash); err != nil {
			return fmt.Errorf("failed to read integration coordination info: %v", err)
		}

		from := uint64(fromSeq)
		klog.V(1).Infof("Integrating from %d", from)
		lh, err := m.fetchLeafHashes(ctx, from, sourceSize, sourceSize)
		if err != nil {
			return fmt.Errorf("fetchLeafHashes(%d, %d, %d): %v", from, sourceSize, sourceSize, err)
		}

		if len(lh) == 0 {
			klog.Infof("Integrate: nothing to do, nothing done")
			return nil
		}

		added := uint64(len(lh))
		klog.Infof("Integrate: adding %d entries to existing tree size %d", len(lh), from)
		newRoot, err = m.s.integrate(ctx, from, lh)
		if err != nil {
			klog.Warningf("integrate failed: %v", err)
			return fmt.Errorf("Integrate failed: %v", err)
		}
		newSize = from + added
		klog.Infof("Integrate: added %d entries", added)

		// integration was successful, so we can update our coordination row
		m := make([]*spanner.Mutation, 0)
		m = append(m, spanner.Update("IntCoord", []string{"id", "seq", "rootHash"}, []interface{}{0, int64(from + added), newRoot}))
		return txn.BufferWrite(m)
	})

	if err != nil {
		return 0, nil, err
	}
	return newSize, newRoot, nil
}

// createAndPrepareTables applies the passed in list of DDL statements and groups of mutations.
//
// This is intended to be used to create and initialise Spanner instances on first use.
// DDL should likely be of the form "CREATE TABLE IF NOT EXISTS".
// Mutation groups should likey be one or more spanner.Insert operations - AlreadyExists errors will be silently ignored.
func createAndPrepareTables(ctx context.Context, spannerDB string, ddl []string, mutations [][]*spanner.Mutation) error {
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	op, err := adminClient.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
		Database:   spannerDB,
		Statements: ddl,
	})
	if err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}
	if err := op.Wait(ctx); err != nil {
		return err
	}
	adminClient.Close()

	dbPool, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return fmt.Errorf("failed to connect to Spanner: %v", err)
	}
	defer dbPool.Close()

	// Set default values for a newly initialised schema using passed in mutation groups.
	// Note that this will only succeed if no row exists, so there's no danger of "resetting" an existing log.
	for _, mg := range mutations {
		if _, err := dbPool.Apply(ctx, mg); err != nil && spanner.ErrCode(err) != codes.AlreadyExists {
			return err
		}
	}
	return nil
}
