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
	"github.com/transparency-dev/trillian-tessera/internal/witness"
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

	DefaultIntegrationSizeLimit = 5 * 4096

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
	cfg Config
}

// sequencer describes a type which knows how to sequence entries.
//
// TODO(al): rename this as it's really more of a coordination for the log.
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
	// currentTree returns the tree state of the currently integrated tree according to the IntCoord table.
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
func New(ctx context.Context, cfg Config) (tessera.Driver, error) {
	return &Storage{
		cfg: cfg,
	}, nil
}

type LogReader struct {
	lrs            logResourceStore
	integratedSize func(context.Context) (uint64, error)
}

func (lr *LogReader) ReadCheckpoint(ctx context.Context) ([]byte, error) {
	return lr.lrs.getCheckpoint(ctx)
}

func (lr *LogReader) ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error) {
	return lr.lrs.getTile(ctx, l, i, p)
}

func (lr *LogReader) ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error) {
	return lr.lrs.getEntryBundle(ctx, i, p)
}

func (lr *LogReader) IntegratedSize(ctx context.Context) (uint64, error) {
	return lr.integratedSize(ctx)
}

func (lr *LogReader) StreamEntries(ctx context.Context, fromEntry uint64) (next func() (ri layout.RangeInfo, bundle []byte, err error), cancel func()) {
	klog.Infof("StreamEntries from %d", fromEntry)

	return streamAdaptor(ctx, lr.integratedSize, lr.lrs.getEntryBundle, fromEntry)
}

func (s *Storage) Appender(ctx context.Context, opts *tessera.AppendOptions) (*tessera.Appender, tessera.LogReader, error) {
	if opts.CheckpointInterval() < minCheckpointInterval {
		return nil, nil, fmt.Errorf("requested CheckpointInterval (%v) is less than minimum permitted %v", opts.CheckpointInterval(), minCheckpointInterval)
	}
	if opts.NewCP() == nil {
		return nil, nil, errors.New("tessera.WithCheckpointSigner must be provided in Appender()")
	}

	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	seq, err := newSpannerCoordinator(ctx, s.cfg.Spanner, uint64(opts.PushbackMaxOutstanding()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Spanner coordinator: %v", err)
	}

	a := &Appender{
		logStore: &logResourceStore{
			objStore: &gcsStorage{
				gcsClient: c,
				bucket:    s.cfg.Bucket,
			},
			entriesPath: opts.EntriesPath(),
		},
		sequencer: seq,
		cpUpdated: make(chan struct{}),
	}
	a.queue = storage.NewQueue(ctx, opts.BatchMaxAge(), opts.BatchMaxSize(), a.sequencer.assignEntries)

	if err := a.init(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to initialise log storage: %v", err)
	}

	go a.sequencerJob(ctx)
	go a.publisherJob(ctx, opts.CheckpointInterval())

	reader := &LogReader{
		lrs: *a.logStore,
		integratedSize: func(context.Context) (uint64, error) {
			s, _, err := a.sequencer.currentTree(ctx)
			return s, err
		},
	}
	wg := witness.NewWitnessGateway(opts.Witnesses(), http.DefaultClient, reader.ReadTile)
	a.newCP = func(u uint64, b []byte) ([]byte, error) {
		cp, err := opts.NewCP()(u, b)
		if err != nil {
			return cp, err
		}
		return wg.Witness(ctx, cp)
	}
	return &tessera.Appender{
		Add: a.Add,
	}, reader, nil
}

// Appender is an implementation of the Tessera appender lifecycle contract.
type Appender struct {
	newCP func(uint64, []byte) ([]byte, error)

	sequencer sequencer
	logStore  *logResourceStore

	queue *storage.Queue

	cpUpdated chan struct{}
}

// Add is the entrypoint for adding entries to a sequencing log.
func (a *Appender) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	return a.queue.Add(ctx, e)
}

// sequencerJob is a long-running function which handles the periodic integration of sequenced entries.
// Blocks until ctx is done.
func (a *Appender) sequencerJob(ctx context.Context) {
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

			if _, err := a.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, a.appendEntries, false); err != nil {
				klog.Errorf("integrate: %v", err)
				return
			}
			select {
			case a.cpUpdated <- struct{}{}:
			default:
			}
		}()
	}
}

// publisherJob is a long-running function which handles the periodic publishing of checkpoints.
// Blocks until ctx is done.
func (a *Appender) publisherJob(ctx context.Context, i time.Duration) {
	t := time.NewTicker(i)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-a.cpUpdated:
		case <-t.C:
		}
		if err := a.publishCheckpoint(ctx, i); err != nil {
			klog.Warningf("publishCheckpoint failed: %v", err)
		}
	}
}

// init ensures that the storage represents a log in a valid state.
func (a *Appender) init(ctx context.Context) error {
	if _, err := a.logStore.getCheckpoint(ctx); err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			// No checkpoint exists, do a forced (possibly empty) integration to create one in a safe
			// way (setting the checkpoint directly here would not be safe as it's outside the transactional
			// framework which prevents the tree from rolling backwards or otherwise forking).
			cctx, c := context.WithTimeout(ctx, 10*time.Second)
			defer c()
			if _, err := a.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, a.appendEntries, true); err != nil {
				return fmt.Errorf("forced integrate: %v", err)
			}
			select {
			case a.cpUpdated <- struct{}{}:
			default:
			}
			return nil
		}
		return fmt.Errorf("failed to read checkpoint: %v", err)
	}

	return nil
}

func (a *Appender) publishCheckpoint(ctx context.Context, minStaleness time.Duration) error {
	m, err := a.logStore.checkpointLastModified(ctx)
	if err != nil && !errors.Is(err, gcs.ErrObjectNotExist) {
		return fmt.Errorf("lastModified(%q): %v", layout.CheckpointPath, err)
	}
	if time.Since(m) < minStaleness {
		return nil
	}

	size, root, err := a.sequencer.currentTree(ctx)
	if err != nil {
		return fmt.Errorf("currentTree: %v", err)
	}
	cpRaw, err := a.newCP(size, root)
	if err != nil {
		return fmt.Errorf("newCP: %v", err)
	}

	if err := a.logStore.setCheckpoint(ctx, cpRaw); err != nil {
		return fmt.Errorf("writeCheckpoint: %v", err)
	}
	return nil

}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, int64, error)
	setObject(ctx context.Context, obj string, data []byte, cond *gcs.Conditions, contType string, cacheCtl string) error
	lastModified(ctx context.Context, obj string) (time.Time, error)
}

// logResourceStore knows how to read and write entries which represent a tiles log inside an objStore.
type logResourceStore struct {
	objStore    objStore
	entriesPath func(uint64, uint8) string
}

func (lrs *logResourceStore) setCheckpoint(ctx context.Context, cpRaw []byte) error {
	return lrs.objStore.setObject(ctx, layout.CheckpointPath, cpRaw, nil, ckptContType, ckptCacheControl)
}

func (lrs *logResourceStore) checkpointLastModified(ctx context.Context) (time.Time, error) {
	t, err := lrs.objStore.lastModified(ctx, layout.CheckpointPath)
	return t, err
}

func (lrs *logResourceStore) getCheckpoint(ctx context.Context) ([]byte, error) {
	r, _, err := lrs.objStore.getObject(ctx, layout.CheckpointPath)
	return r, err
}

// setTile idempotently stores the provided tile at the location implied by the given level, index, and treeSize.
//
// The location to which the tile is written is defined by the tile layout spec.
func (s *logResourceStore) setTile(ctx context.Context, level, index uint64, partial uint8, data []byte) error {
	tPath := layout.TilePath(level, index, partial)
	return s.objStore.setObject(ctx, tPath, data, &gcs.Conditions{DoesNotExist: true}, logContType, logCacheControl)
}

// getTile retrieves the raw tile from the provided location.
//
// The location to which the tile is written is defined by the tile layout spec.
func (s *logResourceStore) getTile(ctx context.Context, level, index uint64, partial uint8) ([]byte, error) {
	tPath := layout.TilePath(level, index, partial)
	d, _, err := s.objStore.getObject(ctx, tPath)
	return d, err
}

// getTiles returns the tiles with the given tile-coords for the specified log size.
//
// Tiles are returned in the same order as they're requested, nils represent tiles which were not found.
func (s *logResourceStore) getTiles(ctx context.Context, tileIDs []storage.TileID, logSize uint64) ([]*api.HashTile, error) {
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

// getbundleFn is a function which knows how to fetch a single entry bundle from the specified address.
type getBundleFn func(ctx context.Context, bundleIdx uint64, partial uint8) ([]byte, error)

// getSizeFn is a function which knows how to return a tree size.
type getSizeFn func(ctx context.Context) (uint64, error)

// streamAdaptor uses the provided function to produce a stream of entry bundles accesible via the returned functions.
//
// Entry bundles are retuned strictly in order via consecutive calls to the returned next func.
// If the adaptor encounters an error while reading an entry bundle, the encountered error will be returned by the corresponding call to next,
// and the stream will be stopped - further calls to next will continue to return errors.
//
// When the caller has finished consuming entry bundles (either because of an error being returned via next, or having consumed all the bundles it needs),
// it MUST call the returned cancel function to release resources.
//
// This adaptor is optimised for the case where calling getBundle has some appreciable latency, and works
// around that by maintaining a read-ahead cache of subsequent bundles.
//
// TODO(al): consider whether this should be factored out as a storage mix-in.
func streamAdaptor(ctx context.Context, getSize getSizeFn, getBundle getBundleFn, fromEntry uint64) (next func() (ri layout.RangeInfo, bundle []byte, err error), cancel func()) {
	// bundleOrErr represents a fetched entry bundle and its params, or an error if we couldn't fetch it for
	// some reason.
	type bundleOrErr struct {
		ri  layout.RangeInfo
		b   []byte
		err error
	}
	// TODO(al): this should probably be configurable - it's primarily intended to act as a means to balance throughput against
	//           consumption of resources, but such balancing needs to be mindful of the nature of the source infrastructure, and
	//           how concurrent requests affect performance (e.g. GCS buckets vs. files on a single disk).
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

		// We'll keep looping around until told to exit.
		for {
			// Check afresh what size the tree is so we can keep streaming entries as the tree grows.
			treeSize, err := getSize(ctx)
			if err != nil {
				klog.Warningf("streamAdaptor: failed to get current tree size: %v", err)
				continue
			}
			klog.Infof("tick from %d to %d", fromEntry, treeSize)

			// For each bundle, pop a future into the bundles channel and kick off an async request
			// to resolve it.
			for ri := range layout.Range(fromEntry, treeSize, treeSize) {
				select {
				case <-exit:
					break
				case <-tokens:
					// We'll return a token below, once the bundle is fetched _and_ is being yielded.
				}

				c := make(chan bundleOrErr, 1)
				go func(ri layout.RangeInfo) {
					b, err := getBundle(ctx, ri.Index, ri.Partial)
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

			// Next loop, carry on from where we got to.
			fromEntry = treeSize

			select {
			case <-exit:
				klog.Infof("streamAdaptor exiting")
				return
			case <-time.After(time.Second):
				// We've caught up with and hit the end of the tree, so wait a bit before looping to avoid busy waiting.
				// TODO(al): could consider a shallow channel of sizes here.
			}
		}
	}()

	cancel = func() {
		close(exit)
	}

	var streamErr error
	next = func() (layout.RangeInfo, []byte, error) {
		if streamErr != nil {
			return layout.RangeInfo{}, nil, streamErr
		}

		f, ok := <-bundles
		if !ok {
			streamErr = tessera.ErrNoMoreEntries
			return layout.RangeInfo{}, nil, streamErr
		}
		b := f()
		if b.err != nil {
			streamErr = b.err
		}
		return b.ri, b.b, b.err
	}
	return next, cancel
}

// getEntryBundle returns the serialised entry bundle at the location described by the given index and partial size.
// A partial size of zero implies a full tile.
//
// Returns a wrapped os.ErrNotExist if the bundle does not exist.
func (s *logResourceStore) getEntryBundle(ctx context.Context, bundleIndex uint64, p uint8) ([]byte, error) {
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
func (s *logResourceStore) setEntryBundle(ctx context.Context, bundleIndex uint64, p uint8, bundleRaw []byte) error {
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
func (a *Appender) appendEntries(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
	var newRoot []byte

	errG := errgroup.Group{}

	errG.Go(func() error {
		if err := a.updateEntryBundles(ctx, fromSeq, entries); err != nil {
			return fmt.Errorf("updateEntryBundles: %v", err)
		}
		return nil
	})

	errG.Go(func() error {
		lh := make([][]byte, len(entries))
		for i, e := range entries {
			lh[i] = e.LeafHash
		}
		r, err := integrate(ctx, fromSeq, lh, a.logStore)
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
func integrate(ctx context.Context, fromSeq uint64, lh [][]byte, logStore *logResourceStore) ([]byte, error) {
	errG := errgroup.Group{}
	getTiles := func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := logStore.getTiles(ctx, tileIDs, treeSize)
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
				return logStore.setTile(ctx, k.Level, k.Index, layout.PartialTileSize(k.Level, k.Index, newSize), data)
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
func (a *Appender) updateEntryBundles(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
	if len(entries) == 0 {
		return nil
	}

	numAdded := uint64(0)
	bundleIndex, entriesInBundle := fromSeq/layout.EntryBundleWidth, fromSeq%layout.EntryBundleWidth
	bundleWriter := &bytes.Buffer{}
	if entriesInBundle > 0 {
		// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
		part, err := a.logStore.getEntryBundle(ctx, uint64(bundleIndex), uint8(entriesInBundle))
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
			if err := a.logStore.setEntryBundle(ctx, bundleIndex, p, bundleRaw); err != nil {
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

// spannerCoordinator uses Cloud Spanner to provide
// a durable and thread/multi-process safe sequencer.
type spannerCoordinator struct {
	dbPool         *spanner.Client
	maxOutstanding uint64
}

// newSpannerCoordinator returns a new spannerSequencer struct which uses the provided
// spanner resource name for its spanner connection.
func newSpannerCoordinator(ctx context.Context, spannerDB string, maxOutstanding uint64) (*spannerCoordinator, error) {
	dbPool, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Spanner: %v", err)
	}
	r := &spannerCoordinator{
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
func (s *spannerCoordinator) initDB(ctx context.Context, spannerDB string) error {
	return createAndPrepareTables(
		ctx, spannerDB,
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
func (s *spannerCoordinator) checkDataCompatibility(ctx context.Context) error {
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
func (s *spannerCoordinator) assignEntries(ctx context.Context, entries []*tessera.Entry) error {
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
			return fmt.Errorf("failed to read SeqCoord: %w", err)
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
func (s *spannerCoordinator) consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error) {
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
func (s *spannerCoordinator) currentTree(ctx context.Context) (uint64, []byte, error) {
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

// NewAntispam returns an antispam driver which uses Spanner to maintain a mapping of
// previously seen entries and their assigned indices.
//
// Note that the storage for this mapping is entirely separate and unconnected to the storage used for
// maintaining the Merkle tree.
//
// This functionality is experimental!
func NewAntispam(ctx context.Context, spannerDB string) (*AntispamStorage, error) {
	if err := createAndPrepareTables(
		ctx, spannerDB,
		[]string{
			"CREATE TABLE IF NOT EXISTS FollowCoord (id INT64 NOT NULL, nextIdx INT64 NOT NULL) PRIMARY KEY (id)",
			"CREATE TABLE IF NOT EXISTS IDSeq (id INT64 NOT NULL, h BYTES(32) NOT NULL, idx INT64 NOT NULL) PRIMARY KEY (id, h)",
		},
		[][]*spanner.Mutation{
			{spanner.Insert("FollowCoord", []string{"id", "nextIdx"}, []interface{}{0, 0})},
		},
	); err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	db, err := spanner.NewClient(ctx, spannerDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Spanner: %v", err)
	}

	r := &AntispamStorage{
		dbPool: db,
	}

	go func(ctx context.Context) {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				klog.V(1).Infof("ANTISPAM: # Writes %d, # Lookups %d, # DB hits %v", r.numWrites.Load(), r.numLookups.Load(), r.numHits.Load())
			}
		}
	}(ctx)

	return r, nil
}

type AntispamStorage struct {
	dbPool *spanner.Client

	// pushBack is used to prevent the follower from getting too far underwater.
	// Populate dynamically will set this to true/false based on how far behind the follower is from the
	// currently integrated tree size.
	// When pushBack is true, the decorator will start returning ErrPushback to all calls.
	pushBack atomic.Bool

	numLookups atomic.Uint64
	numWrites  atomic.Uint64
	numHits    atomic.Uint64
}

// index returns the index (if any) previously associated with the provided hash
func (d *AntispamStorage) index(ctx context.Context, h []byte) (*uint64, error) {
	d.numLookups.Add(1)
	var idx int64
	if row, err := d.dbPool.Single().ReadRow(ctx, "IDSeq", spanner.Key{0, h}, []string{"idx"}); err != nil {
		if c := spanner.ErrCode(err); c == codes.NotFound {
			return nil, nil
		}
		return nil, err
	} else {
		if err := row.Column(0, &idx); err != nil {
			return nil, fmt.Errorf("failed to read antispam index: %v", err)
		}
		idx := uint64(idx)
		d.numHits.Add(1)
		return &idx, nil
	}
}

// Decorator returns a function which will wrap an underlying Add delegate with
// code to dedup against the stored data.
func (d *AntispamStorage) Decorator() func(f tessera.AddFn) tessera.AddFn {
	return func(delegate tessera.AddFn) tessera.AddFn {
		return func(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
			if d.pushBack.Load() {
				// The follower is too far behind the currently integrated tree, so we're going to push back against
				// the incoming requests.
				// This should have two effects:
				//   1. The tree will cease growing, giving the follower a chance to catch up, and
				//   2. We'll stop doing lookups for each submission, freeing up Spanner CPU to focus on catching up.
				//
				// We may decide in the future that serving duplicate reads is more important than catching up as quickly
				// as possible, in which case we'd move this check down below the call to index.
				return func() (uint64, error) { return 0, tessera.ErrPushback }
			}
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

// entryStreamReader converts a stream of {RangeInfo, EntryBundle} into a stream of individually processed entries.
//
// TODO(al): Factor this out for re-use elsewhere when it's ready.
type entryStreamReader[T any] struct {
	bundleFn func([]byte) ([]T, error)
	next     func() (layout.RangeInfo, []byte, error)

	curData []T
	curRI   layout.RangeInfo
	i       uint64
}

// newEntryStreamReader creates a new stream reader which uses the provided bundleFn to process bundles into processed entries of type T
//
// Different bundleFn implementations can be provided to return raw entry bytes, parsed entry structs, or derivations of entries (e.g. hashes) as needed.
func newEntryStreamReader[T any](next func() (layout.RangeInfo, []byte, error), bundleFn func([]byte) ([]T, error)) *entryStreamReader[T] {
	return &entryStreamReader[T]{
		bundleFn: bundleFn,
		next:     next,
		i:        0,
	}
}

// Next processes and returns the next available entry in the stream along with its index in the log.
func (e *entryStreamReader[T]) Next() (uint64, T, error) {
	var t T
	if len(e.curData) == 0 {
		var err error
		var b []byte
		e.curRI, b, err = e.next()
		if err != nil {
			return 0, t, fmt.Errorf("next: %v", err)
		}
		e.curData, err = e.bundleFn(b)
		if err != nil {
			return 0, t, fmt.Errorf("bundleFn(bundleEntry @%d): %v", e.curRI.Index, err)

		}
		if e.curRI.First > 0 {
			e.curData = e.curData[e.curRI.First:]
		}
		if len(e.curData) > int(e.curRI.N) {
			e.curData = e.curData[:e.curRI.N]
		}
		e.i = 0
	}
	t, e.curData = e.curData[0], e.curData[1:]
	rIdx := e.curRI.Index*layout.EntryBundleWidth + uint64(e.curRI.First) + e.i
	e.i++
	return rIdx, t, nil
}

// Populate uses entry data from the log to populate the antispam storage.
//
// TODO(al):  add details
func (d *AntispamStorage) Populate(ctx context.Context, lr tessera.LogReader, bundleFn func([]byte) ([][]byte, error)) {
	errOutOfSync := errors.New("out-of-sync")

	t := time.NewTicker(time.Second)
	var (
		entryReader *entryStreamReader[[]byte]
		stop        func()

		curEntries [][]byte
		curIndex   uint64
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}
		size, err := lr.IntegratedSize(ctx)
		if err != nil {
			klog.Errorf("Populate: IntegratedSize(): %v", err)
			continue
		}

		// Busy loop while there's work to be done
		for workDone := true; workDone; {
			_, err = d.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
				// Figure out the last entry we used to populate our antispam storage.
				row, err := txn.ReadRowWithOptions(ctx, "FollowCoord", spanner.Key{0}, []string{"nextIdx"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
				if err != nil {
					return err
				}

				var f int64 // Spanner doesn't support uint64
				if err := row.Columns(&f); err != nil {
					return fmt.Errorf("failed to read follow coordination info: %v", err)
				}
				followFrom := uint64(f)
				if followFrom >= size {
					// Our view of the log is out of date, exit the busy loop and refresh it.
					workDone = false
					return nil
				}

				// TODO(al): Maybe make these configurable.
				const batchSize = 40
				const pushBackThreshold = batchSize * 16
				d.pushBack.Store(size-followFrom > pushBackThreshold)

				// If this is the first time around the loop we need to start the stream of entries now that we know where we want to
				// start reading from:
				if entryReader == nil {
					next, st := lr.StreamEntries(ctx, followFrom)
					stop = st
					entryReader = newEntryStreamReader(next, bundleFn)
				}

				if curIndex == followFrom && curEntries != nil {
					// Note that it's possible for Spanner to automatically retry transactions in some circumstances, when it does
					// it'll call this function again.
					// If the above condition holds, then we're in a retry situation and we must use the same data again rather
					// than continue reading entries which will take us out of sync.
				} else {
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
							return errOutOfSync
						}
						batch = append(batch, c)
					}
					curEntries = batch
					curIndex = followFrom
				}

				// Store antispam entries.
				//
				// Note that we're writing the antispam entries outside of the transaction here. The reason is because we absolutely do not want
				// the transaction to fail if there's already an entry for the same hash in the IDSeq table.
				//
				// It looks unusual, but is ok because:
				//  - individual antispam entries fails because there's already an entry for that hash is perfectly ok
				//  - we'll only continue on to update FollowCoord if no errors (other than AlreadyExists) occur while inserting entries
				//  - similarly, if we manage to insert antispam entries here, but then fail to update FollowCoord, we'll end up
				//    retrying over the same set of log entries, and then ignoring the AlreadyExists which will occur.
				//
				// Alternative approaches are:
				//  - Use InsertOrUpdate, but that will keep updating the index associated with the ID hash, and we'd rather keep serving
				//    the earliest index known for that entry.
				//  - Perform reads for each of the hashes we're about to write, and use that to filter writes.
				//    This would work, but would also incur an extra round-trip of data which isn't really necessary but would
				//    slow the process down considerably and add extra load to Spanner for no benefit.
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
							return fmt.Errorf("failed to write antispam record: %v (%v)", s.GetMessage(), c)
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
				if err != errOutOfSync {
					klog.Errorf("Failed to commit antispam population tx: %v", err)
				}
				stop()
				entryReader = nil
				continue
			}
			curEntries = nil
		}
	}
}

// MigrationTarget creates a new GCP storage for the MigrationTarget lifecycle mode.
func (s *Storage) MigrationTarget(ctx context.Context, bundleHasher tessera.UnbundlerFunc, opts *tessera.MigrationOptions) (tessera.MigrationTarget, tessera.LogReader, error) {
	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	seq, err := newSpannerCoordinator(ctx, s.cfg.Spanner, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Spanner sequencer: %v", err)
	}
	m := &MigrationStorage{
		s:            s,
		dbPool:       seq.dbPool,
		bundleHasher: bundleHasher,
		sequencer:    seq,
		logStore: &logResourceStore{
			objStore: &gcsStorage{
				gcsClient: c,
				bucket:    s.cfg.Bucket,
			},
			entriesPath: opts.EntriesPath(),
		},
	}

	r := &LogReader{
		lrs: *m.logStore,
		integratedSize: func(context.Context) (uint64, error) {
			s, _, err := m.sequencer.currentTree(ctx)
			return s, err
		},
	}
	return m, r, nil
}

// MigrationStorgage implements the tessera.MigrationTarget lifecycle contract.
type MigrationStorage struct {
	s            *Storage
	dbPool       *spanner.Client
	bundleHasher func([]byte) ([][]byte, error)
	sequencer    sequencer
	logStore     *logResourceStore
}

var _ tessera.MigrationTarget = &MigrationStorage{}

func (m *MigrationStorage) AwaitIntegration(ctx context.Context, sourceSize uint64) ([]byte, error) {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
			from, _, err := m.sequencer.currentTree(ctx)
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
	return m.logStore.setEntryBundle(ctx, index, partial, bundle)
}

func (m *MigrationStorage) IntegratedSize(ctx context.Context) (uint64, error) {
	sz, _, err := m.sequencer.currentTree(ctx)
	return sz, err
}

func (m *MigrationStorage) fetchLeafHashes(ctx context.Context, from, to, sourceSize uint64) ([][]byte, error) {
	// TODO(al): Make this configurable.
	const maxBundles = 300

	toBeAdded := sync.Map{}
	eg := errgroup.Group{}
	n := 0
	for ri := range layout.Range(from, to, sourceSize) {
		eg.Go(func() error {
			b, err := m.logStore.getEntryBundle(ctx, ri.Index, ri.Partial)
			if err != nil {
				return fmt.Errorf("getEntryBundle(%d.%d): %v", ri.Index, ri.Partial, err)
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
		newRoot, err = integrate(ctx, from, lh, m.logStore)
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
