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

// Package aws contains an AWS-based storage implementation for Tessera.
//
// TODO: decide whether to rename this package.
//
// This storage implementation uses S3 for long-term storage and serving of
// entry bundles and log tiles, and MySQL for coordinating updates to AWS
// when multiple instances of a personality binary are running.
//
// A single S3 bucket is used to hold entry bundles and log internal tiles.
// The object keys for the bucket are selected so as to conform to the
// expected layout of a tile-based log.
//
// A MySQL database provides a transactional mechanism to allow multiple
// frontends to safely update the contents of the log.
package aws

import (
	"bytes"
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/witness"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	logContType           = "application/octet-stream"
	ckptContType          = "text/plain; charset=utf-8"
	minCheckpointInterval = time.Second

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

// Storage is an AWS based storage implementation for Tessera.
type Storage struct {
	cfg Config
}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, error)
	setObject(ctx context.Context, obj string, data []byte, contType string) error
	setObjectIfNoneMatch(ctx context.Context, obj string, data []byte, contType string) error
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

// consumeFunc is the signature of a function which can consume entries from the sequencer.
// Returns the updated root hash of the tree with the consumed entries integrated.
type consumeFunc func(ctx context.Context, from uint64, entries []storage.SequencedEntry) ([]byte, error)

// Config holds AWS project and resource configuration for a storage instance.
type Config struct {
	// SDKConfig is an optional AWS config to use when configuring service clients, e.g. to
	// use non-AWS S3 or MySQL services.
	//
	// If nil, the value from config.LoadDefaultConfig() will be used - this is the only
	// supported configuration.
	SDKConfig *aws.Config
	// S3Options is an optional function which can be used to configure the S3 library.
	// This is primarily useful when configuring the use of non-AWS S3 or MySQL services.
	//
	// If nil, the default options will be used - this is the only supported configuration.
	S3Options func(*s3.Options)
	// Bucket is the name of the S3 bucket to use for storing log state.
	Bucket string
	// DSN is the DSN of the MySQL instance to use.
	DSN string
	// Maximum connections to the MysSQL database
	MaxOpenConns int
	// Maximum idle database connections in the connection pool
	MaxIdleConns int
}

// New creates a new instance of the AWS based Storage.
//
// Storage instances created via this c'tor will participate in integrating newly sequenced entries into the log
// and periodically publishing a new checkpoint which commits to the state of the tree.
func New(ctx context.Context, cfg Config) (tessera.Driver, error) {
	if cfg.SDKConfig == nil {
		// We're running on AWS so use the SDK's default config which will will handle credentials etc.
		sdkConfig, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to load default AWS configuration: %v", err)
		}
		cfg.SDKConfig = &sdkConfig
		// We need a non-nil options func to pass in to s3.NewFromConfig below or it'll panic, so
		// we'll use a "do nothing" placeholder.
		cfg.S3Options = func(_ *s3.Options) {}
	} else {
		printDragonsWarning()
	}

	return &Storage{
		cfg: cfg,
	}, nil
}

func (s *Storage) Appender(ctx context.Context, opts *tessera.AppendOptions) (*tessera.Appender, tessera.LogReader, error) {
	pb := uint64(opts.PushbackMaxOutstanding())
	if pb == 0 {
		pb = DefaultPushbackMaxOutstanding
	}
	if opts.CheckpointInterval() < minCheckpointInterval {
		return nil, nil, fmt.Errorf("requested CheckpointInterval (%v) is less than minimum permitted %v", opts.CheckpointInterval(), minCheckpointInterval)
	}

	seq, err := newMySQLSequencer(ctx, s.cfg.DSN, pb, s.cfg.MaxOpenConns, s.cfg.MaxIdleConns)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create MySQL sequencer: %v", err)
	}

	logStore := &logResourceStore{
		objStore: &s3Storage{
			s3Client: s3.NewFromConfig(*s.cfg.SDKConfig, s.cfg.S3Options),
			bucket:   s.cfg.Bucket,
		},
		entriesPath: opts.EntriesPath(),
	}
	wg := witness.NewWitnessGateway(opts.Witnesses(), http.DefaultClient, logStore.ReadTile)
	r := &Appender{
		logStore:  logStore,
		sequencer: seq,
		newCP: func(u uint64, b []byte) ([]byte, error) {
			cp, err := opts.NewCP()(u, b)
			if err != nil {
				return cp, err
			}
			return wg.Witness(ctx, cp)
		},
		treeUpdated: make(chan struct{}),
	}
	r.queue = storage.NewQueue(ctx, opts.BatchMaxAge(), opts.BatchMaxSize(), r.sequencer.assignEntries)

	if err := r.init(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to initialise log storage: %v", err)
	}

	// Kick off go-routine which handles the integration of entries.
	go r.consumeEntriesTask(ctx)

	// Kick off go-routine which handles the publication of checkpoints.
	go r.publishCheckpointTask(ctx, opts.CheckpointInterval())

	return &tessera.Appender{
		Add: r.Add,
	}, r.logStore, nil
}

// Appender is an implementation of the Tessera appender lifecycle contract.
type Appender struct {
	newCP func(uint64, []byte) ([]byte, error)

	sequencer sequencer
	logStore  *logResourceStore

	queue *storage.Queue

	treeUpdated chan struct{}
}

// sequenceEntriesTask periodically integrates newly sequenced entries.
//
// This function does not return until the passed context is done.
func (a *Appender) consumeEntriesTask(ctx context.Context) {
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
			case a.treeUpdated <- struct{}{}:
			default:
			}
		}()
	}
}

// publishCheckpointTask periodically attempts to publish a new checkpoint representing the current state
// of the tree, once per interval.
//
// This function does not return until the passed in context is done.
func (a *Appender) publishCheckpointTask(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-a.treeUpdated:
		case <-t.C:
		}
		if err := a.publishCheckpoint(ctx, interval); err != nil {
			klog.Warningf("publishCheckpoint: %v", err)
		}
	}
}

// Add is the entrypoint for adding entries to a sequencing log.
func (a *Appender) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	return a.queue.Add(ctx, e)
}

// init ensures that the storage represents a log in a valid state.
func (a *Appender) init(ctx context.Context) error {
	_, err := a.logStore.ReadCheckpoint(ctx)
	if err != nil {
		// Do not use errors.Is. Keep errors.As to compare by type and not by value.
		var nske *types.NoSuchKey
		if errors.As(err, &nske) {
			// No checkpoint exists, do a forced (possibly empty) integration to create one in a safe
			// way (calling updateCP directly here would not be safe as it's outside the transactional
			// framework which prevents the tree from rolling backwards or otherwise forking).
			cctx, c := context.WithTimeout(ctx, 10*time.Second)
			defer c()
			if _, err := a.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, a.appendEntries, true); err != nil {
				return fmt.Errorf("forced integrate: %v", err)
			}
			select {
			case a.treeUpdated <- struct{}{}:
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
	// Do not use errors.Is. Keep errors.As to compare by type and not by value.
	var nske *types.NoSuchKey
	if err != nil && !errors.As(err, &nske) {
		return fmt.Errorf("checkpointLastModified(): %v", err)
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

// appendEntries incorporates the provided entries into the log starting at fromSeq.
//
// Returns the new root hash of the log with the entries added.
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

	err := errG.Wait()
	return newRoot, err
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
			klog.V(1).Infof("In-memory bundle idx %d is full, attempting write to S3", bundleIndex)
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
		klog.V(1).Infof("Attempting to write in-memory partial bundle idx %d.%d to S3", bundleIndex, entriesInBundle)
		goSetEntryBundle(ctx, bundleIndex, uint8(entriesInBundle), bundleWriter.Bytes())
	}
	return seqErr.Wait()
}

// logResourceStore knows how to read and write entries which represent a tiles log inside an objStore.
type logResourceStore struct {
	objStore    objStore
	entriesPath func(uint64, uint8) string
}

func (lr *logResourceStore) ReadCheckpoint(ctx context.Context) ([]byte, error) {
	return lr.get(ctx, layout.CheckpointPath)
}

func (lr *logResourceStore) ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error) {
	return lr.get(ctx, layout.TilePath(l, i, p))
}

func (lr *logResourceStore) ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error) {
	return lr.get(ctx, lr.entriesPath(i, p))
}

func (lr *logResourceStore) IntegratedSize(ctx context.Context) (uint64, error) {
	return 0, errors.New("unimplemented")
}

func (lr *logResourceStore) StreamEntries(ctx context.Context, fromEntry uint64) (next func() (ri layout.RangeInfo, bundle []byte, err error), cancel func()) {
	return func() (layout.RangeInfo, []byte, error) {
		return layout.RangeInfo{}, nil, errors.New("unimplemented")
	}, func() {}
}

// get returns the requested object.
//
// This is indended to be used to proxy read requests through the personality for debug/testing purposes.
func (s *logResourceStore) get(ctx context.Context, path string) ([]byte, error) {
	d, err := s.objStore.getObject(ctx, path)
	return d, err
}

func (lrs *logResourceStore) setCheckpoint(ctx context.Context, cpRaw []byte) error {
	return lrs.objStore.setObject(ctx, layout.CheckpointPath, cpRaw, ckptContType)
}

func (lrs *logResourceStore) checkpointLastModified(ctx context.Context) (time.Time, error) {
	t, err := lrs.objStore.lastModified(ctx, layout.CheckpointPath)
	return t, err
}

// setTile idempotently stores the provided tile at the location implied by the given level, index, and treeSize.
//
// The location to which the tile is written is defined by the tile layout spec.
func (lrs *logResourceStore) setTile(ctx context.Context, level, index, logSize uint64, tile *api.HashTile) error {
	data, err := tile.MarshalText()
	if err != nil {
		return err
	}
	tPath := layout.TilePath(level, index, layout.PartialTileSize(level, index, logSize))
	klog.V(2).Infof("StoreTile: %s (%d entries)", tPath, len(tile.Nodes))

	return lrs.objStore.setObjectIfNoneMatch(ctx, tPath, data, logContType)
}

// getTiles returns the tiles with the given tile-coords for the specified log size.
//
// Tiles are returned in the same order as they're requested, nils represent tiles which were not found.
func (lrs *logResourceStore) getTiles(ctx context.Context, tileIDs []storage.TileID, logSize uint64) ([]*api.HashTile, error) {
	r := make([]*api.HashTile, len(tileIDs))
	errG := errgroup.Group{}
	for i, id := range tileIDs {
		i := i
		id := id
		errG.Go(func() error {
			objName := layout.TilePath(id.Level, id.Index, layout.PartialTileSize(id.Level, id.Index, logSize))
			data, err := lrs.objStore.getObject(ctx, objName)
			if err != nil {
				// Do not use errors.Is. Keep errors.As to compare by type and not by value.
				var nske *types.NoSuchKey
				if errors.As(err, &nske) {
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
func (lrs *logResourceStore) getEntryBundle(ctx context.Context, bundleIndex uint64, p uint8) ([]byte, error) {
	objName := lrs.entriesPath(bundleIndex, p)
	data, err := lrs.objStore.getObject(ctx, objName)
	if err != nil {
		// Do not use errors.Is. Keep errors.As to compare by type and not by value.
		var nske *types.NoSuchKey
		if errors.As(err, &nske) {
			// Return the generic NotExist error so that higher levels can differentiate
			// between this and other errors.
			return nil, fmt.Errorf("%v: %w", objName, os.ErrNotExist)
		}
		return nil, err
	}

	return data, nil
}

// setEntryBundle idempotently stores the serialised entry bundle at the location implied by the bundleIndex and treeSize.
func (lrs *logResourceStore) setEntryBundle(ctx context.Context, bundleIndex uint64, p uint8, bundleRaw []byte) error {
	objName := lrs.entriesPath(bundleIndex, p)
	// Note that setObject does an idempotent interpretation of IfNoneMatch - it only
	// returns an error if the named object exists _and_ contains different data to what's
	// passed in here.
	if err := lrs.objStore.setObjectIfNoneMatch(ctx, objName, bundleRaw, logContType); err != nil {
		return fmt.Errorf("setObjectIfNoneMatch(%q): %v", objName, err)

	}
	return nil
}

// integrate adds the provided leaf hashes to the merkle tree, starting at the provided location.
func integrate(ctx context.Context, fromSeq uint64, lh [][]byte, lrs *logResourceStore) ([]byte, error) {
	getTiles := func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := lrs.getTiles(ctx, tileIDs, treeSize)
		if err != nil {
			return nil, fmt.Errorf("getTiles: %w", err)
		}
		return n, nil
	}

	newSize, newRoot, tiles, err := storage.Integrate(ctx, getTiles, fromSeq, lh)
	if err != nil {
		return nil, fmt.Errorf("Integrate: %v", err)
	}
	errG := errgroup.Group{}
	for k, v := range tiles {
		func(ctx context.Context, k storage.TileID, v *api.HashTile) {
			errG.Go(func() error {
				return lrs.setTile(ctx, uint64(k.Level), k.Index, newSize, v)
			})
		}(ctx, k, v)
	}
	if err := errG.Wait(); err != nil {
		return nil, err
	}
	klog.Infof("New tree: %d, %x", newSize, newRoot)
	return newRoot, nil
}

// mySQLSequencer uses MySQL to provide
// a durable and thread/multi-process safe sequencer.
type mySQLSequencer struct {
	dbPool         *sql.DB
	maxOutstanding uint64
}

// newMySQLSequencer returns a new mysqlSequencer struct which uses the provided
// DSN for its MySQL connection.
func newMySQLSequencer(ctx context.Context, dsn string, maxOutstanding uint64, maxOpenConns, maxIdleConns int) (*mySQLSequencer, error) {
	dbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL db: %v", err)
	}

	if maxOpenConns > 0 {
		dbPool.SetMaxOpenConns(maxOpenConns)
	}
	if maxIdleConns >= 0 {
		dbPool.SetMaxIdleConns(maxIdleConns)
	}

	if err := dbPool.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL db: %v", err)
	}

	r := &mySQLSequencer{
		dbPool:         dbPool,
		maxOutstanding: maxOutstanding,
	}

	if err := r.initDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to initDB: %v", err)
	}
	if err := r.checkDataCompatibility(ctx); err != nil {
		return nil, fmt.Errorf("schema is not compatible with this version of the Tessera library: %v", err)
	}
	return r, nil
}

// checkDataCompatibility compares the Tessera library SchemaCompatibilityVersion with the one stored in the
// database, and returns an error if they are not identical.
func (s *mySQLSequencer) checkDataCompatibility(ctx context.Context) error {
	row := s.dbPool.QueryRowContext(ctx, "SELECT compatibilityVersion FROM Tessera WHERE id = 0")
	var gotVersion uint64
	if err := row.Scan(&gotVersion); err != nil {
		return fmt.Errorf("failed to read schema compatibility version from DB: %v", err)
	}

	if gotVersion != SchemaCompatibilityVersion {
		return fmt.Errorf("schema compatibilityVersion (%d) != library compatibilityVersion (%d)", gotVersion, SchemaCompatibilityVersion)
	}
	return nil
}

// initDB ensures that the coordination DB is initialised correctly.
//
// It creates tables if they don't exist already, and inserts zero values.
//
// The database schema consists of 4 tables:
//   - Tessera
//     This table only ever contains a single row which tracks the compatibility
//     version of the DB schema and data formats.
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
func (s *mySQLSequencer) initDB(ctx context.Context) error {
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Tessera (
			id INT UNSIGNED NOT NULL,
			compatibilityVersion BIGINT UNSIGNED NOT NULL,
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS SeqCoord(
			id INT UNSIGNED NOT NULL,
			next BIGINT UNSIGNED NOT NULL,
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}
	// TODO(phboneff): test this with very large leaves, consider downgrading to MEDIUMBLOB.
	// Keep in mind that CT leaves can be large, as large as: https://crt.sh/?id=10751627.
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Seq(
			id INT UNSIGNED NOT NULL,
			seq BIGINT UNSIGNED NOT NULL,
			v LONGBLOB,
			PRIMARY KEY (id, seq)
		)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS IntCoord(
			id INT UNSIGNED NOT NULL,
			seq BIGINT UNSIGNED NOT NULL,
			rootHash TINYBLOB NOT NULL,
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}

	// Set default values for a newly initialised schema - these rows being present are a precondition for
	// sequencing and integration to occur.
	// Note that this will only succeed if no row exists, so there's no danger
	// of "resetting" an existing log.
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO Tessera (id, compatibilityVersion) VALUES (0, ?)`, SchemaCompatibilityVersion); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO SeqCoord (id, next) VALUES (0, 0)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO IntCoord (id, seq, rootHash) VALUES (0, 0, ?)`, rfc6962.DefaultHasher.EmptyRoot()); err != nil {
		return err
	}
	return nil
}

// assignEntries durably assigns each of the passed-in entries an index in the log.
//
// Entries are allocated contiguous indices, in the order in which they appear in the entries parameter.
// This is achieved by storing the passed-in entries in the Seq table in MySQL, keyed by the
// index assigned to the first entry in the batch.
func (s *mySQLSequencer) assignEntries(ctx context.Context, entries []*tessera.Entry) error {
	// First grab the treeSize in a non-locking read-only fashion (we don't want to block/collide with integration).
	// We'll use this value to determine whether we need to apply back-pressure.
	var treeSize uint64
	row := s.dbPool.QueryRowContext(ctx, "SELECT seq FROM IntCoord WHERE id = ?", 0)
	if err := row.Scan(&treeSize); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to read integration coordination info: %v", err)
	}

	// Now move on with sequencing in a single transaction
	tx, err := s.dbPool.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin Tx: %v", err)
	}
	defer func() {
		if tx != nil {
			if err := tx.Rollback(); err != nil {
				klog.Errorf("failed to rollback Tx: %v", err)
			}
		}
	}()

	// First we need to grab the next available sequence number from the SeqCoord table.
	var next, id uint64
	r := tx.QueryRowContext(ctx, "SELECT id, next FROM SeqCoord WHERE id = ? FOR UPDATE", 0)
	if err := r.Scan(&id, &next); err != nil {
		return fmt.Errorf("failed to read seqcoord: %v", err)
	}

	// Check whether there are too many outstanding entries and we should apply
	// back-pressure.
	if outstanding := next - treeSize; outstanding > s.maxOutstanding {
		return tessera.ErrPushback
	}

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
	num := uint64(len(entries))

	// Insert our newly sequenced batch of entries into Seq,
	if _, err := tx.ExecContext(ctx, "INSERT INTO Seq(id, seq, v) VALUES(?, ?, ?)", 0, next, data); err != nil {
		return fmt.Errorf("insert into seq: %v", err)
	}
	// and update the next-available sequence number row in SeqCoord.
	if _, err := tx.ExecContext(ctx, "UPDATE SeqCoord SET next = ? WHERE ID = ?", next+num, 0); err != nil {
		return fmt.Errorf("update seqcoord: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit Tx: %v", err)
	}
	tx = nil

	return nil
}

// consumeEntries calls f with previously sequenced entries.
//
// Once f returns without error, the entries it was called with are considered to have been consumed and are
// removed from the Seq table.
//
// Returns true if some entries were consumed as a weak signal that there may be further entries waiting to be consumed.
func (s *mySQLSequencer) consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error) {
	tx, err := s.dbPool.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("failed to begin Tx: %v", err)
	}
	defer func() {
		if tx != nil {
			if err := tx.Rollback(); err != nil {
				klog.Errorf("failed to rollback Tx: %v", err)
			}
		}
	}()

	// Figure out which is the starting index of sequenced entries to start consuming from.
	row := tx.QueryRowContext(ctx, "SELECT seq, rootHash FROM IntCoord WHERE id = ? FOR UPDATE", 0)
	var fromSeq uint64
	var rootHash []byte
	if err := row.Scan(&fromSeq, &rootHash); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to read IntCoord: %v", err)
	}
	klog.V(1).Infof("Consuming from %d", fromSeq)

	// Now read the sequenced starting at the index we got above.
	rows, err := tx.QueryContext(ctx, "SELECT seq, v FROM Seq WHERE id = ? AND seq >= ? ORDER BY seq LIMIT ? FOR UPDATE", 0, fromSeq, limit)
	if err != nil {
		return false, fmt.Errorf("failed to read Seq: %v", err)
	}
	defer rows.Close()

	// This needs to be of type `any`, to be passed to ExecContext. Only uint64s will be stored.
	seqsConsumed := []any{}
	entries := make([]storage.SequencedEntry, 0, limit)
	orderCheck := fromSeq
	for rows.Next() {

		var vGob []byte
		var seq uint64
		if err := rows.Scan(&seq, &vGob); err != nil {
			return false, fmt.Errorf("failed to scan Seq row: %v", err)
		}

		if orderCheck != seq {
			return false, fmt.Errorf("integrity fail - expected seq %d, but found %d", orderCheck, seq)
		}

		g := gob.NewDecoder(bytes.NewReader(vGob))
		b := []storage.SequencedEntry{}
		if err := g.Decode(&b); err != nil {
			return false, fmt.Errorf("failed to deserialise v from Seq: %v", err)
		}
		entries = append(entries, b...)
		seqsConsumed = append(seqsConsumed, seq)
		orderCheck += uint64(len(b))
	}
	if len(seqsConsumed) == 0 && !forceUpdate {
		klog.V(1).Info("Found no rows to sequence")
		return false, nil
	}

	// Call consumeFunc with the entries we've found
	newRoot, err := f(ctx, uint64(fromSeq), entries)
	if err != nil {
		return false, err
	}

	// consumeFunc was successful, so we can update our coordination row, and delete the row(s) for
	// the then consumed entries.
	if _, err := tx.ExecContext(ctx, "UPDATE IntCoord SET seq=?, rootHash=? WHERE id=?", orderCheck, newRoot, 0); err != nil {
		return false, fmt.Errorf("update intcoord: %v", err)
	}

	if len(seqsConsumed) > 0 {
		// TODO(phboneff): evaluate if seq BETWEEN ? AND ? is more efficient
		q := "DELETE FROM Seq WHERE id=? AND seq IN ( " + placeholder(len(seqsConsumed)) + " )"
		if _, err := tx.ExecContext(ctx, q, append([]any{0}, seqsConsumed...)...); err != nil {
			return false, fmt.Errorf("update intcoord: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("failed to commit Tx: %v", err)
	}
	tx = nil

	return true, nil
}

// currentTree returns the size and root hash of the currently integrated tree.
func (s *mySQLSequencer) currentTree(ctx context.Context) (uint64, []byte, error) {
	row := s.dbPool.QueryRowContext(ctx, "SELECT seq, rootHash FROM IntCoord WHERE id = ?", 0)
	var fromSeq uint64
	var rootHash []byte
	if err := row.Scan(&fromSeq, &rootHash); err != nil {
		return 0, nil, fmt.Errorf("failed to read IntCoord: %v", err)
	}

	return fromSeq, rootHash, nil
}

func placeholder(n int) string {
	places := make([]string, n)
	for i := 0; i < n; i++ {
		places[i] = "?"
	}
	return strings.Join(places, ",")
}

// s3Storage knows how to store and retrieve objects from S3.
type s3Storage struct {
	bucket   string
	s3Client *s3.Client
}

// getObject returns the data of the specified object, or an error.
func (s *s3Storage) getObject(ctx context.Context, obj string) ([]byte, error) {
	r, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		return nil, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("getObject: failed to read %q: %v", obj, err)
	}
	return d, r.Body.Close()
}

// setObject stores the provided data in the specified object.
func (s *s3Storage) setObject(ctx context.Context, objName string, data []byte, contType string) error {
	put := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(objName),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contType),
	}

	if _, err := s.s3Client.PutObject(ctx, put); err != nil {
		return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket, err)
	}
	return nil
}

// setObjectIfNoneMatch stores data in the specified object gated by a IfNoneMatch condition.
//
// ifNoneMatch can be used to specify the IfNoneMatch preconditions for the write, i.e write
// iff no object exists under this key already. If an object already exists under the same key,
// an error will be returned *unless*  the currently stored data is bit-for-bit identical to the
// data to-be-written. This is intended to provide idempotentency for writes.
func (s *s3Storage) setObjectIfNoneMatch(ctx context.Context, objName string, data []byte, contType string) error {
	put := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(objName),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contType),
		// "*" is the expected character for this condition
		IfNoneMatch: aws.String("*"),
	}

	if _, err := s.s3Client.PutObject(ctx, put); err != nil {

		// If we run into a precondition failure error, check that the object
		// which exists contains the same content that we want to write.
		// If so, we can consider this write to be idempotently successful.
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "PreconditionFailed" {
			existing, err := s.getObject(ctx, objName)
			if err != nil {
				return fmt.Errorf("failed to fetch existing content for %q: %v", objName, err)
			}
			if !bytes.Equal(existing, data) {
				klog.Errorf("Resource %q non-idempotent write:\n%s", objName, cmp.Diff(existing, data))
				return fmt.Errorf("precondition failed: resource content for %q differs from data to-be-written", objName)
			}

			klog.V(2).Infof("setObjectIfNoneMatch: identical resource already exists for %q, continuing", objName)
			return nil
		}

		return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket, err)
	}
	return nil
}

// lastModified returns the time the specified object was last modified, or an error
func (s *s3Storage) lastModified(ctx context.Context, obj string) (time.Time, error) {
	r, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		return time.Time{}, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
	}

	return *r.LastModified, r.Body.Close()
}

func printDragonsWarning() {
	d := `H4sIAFZYZGcAA01QMQ7EIAzbeYXV5UCqkq1bf2IFtpNuPalj334hFQdkwLGNAwBzyXnKitOiqTYj
B7ZGplWEwZhZqxZ1aKuswcD0AA4GXPUhI0MEpSd5Ow09vJ+m6rVtF6m0GDccYXDZEdp9N/g1H9Pf
Qu80vNj7tiOe0lkdc8hwZK9YxavT0+FTP++vU6DUKvpEOr1+VGTk3IBXKSX9AHz5xXRgAQAA`
	g, _ := base64.StdEncoding.DecodeString(d)
	r, _ := gzip.NewReader(bytes.NewReader(g))
	t, _ := io.ReadAll(r)
	klog.Infof("Running in non-AWS mode - see storage/aws/README.md for more details.")
	klog.Infof("Here be dragons!\n%s", t)
}
