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
	"context"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/google/go-cmp/cmp"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/options"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	entryBundleSize = 256
	logContType     = "application/octet-stream"
	ckptContType    = "text/plain; charset=utf-8"

	DefaultPushbackMaxOutstanding = 4096
	DefaultIntegrationSizeLimit   = 5 * 4096
)

// Storage is an AWS based storage implementation for Tessera.
type Storage struct {
	newCP       options.NewCPFunc
	entriesPath options.EntriesPathFunc

	sequencer sequencer
	objStore  objStore

	queue *storage.Queue
}

// objStore describes a type which can store and retrieve objects.
type objStore interface {
	getObject(ctx context.Context, obj string) ([]byte, error)
	setObject(ctx context.Context, obj string, data []byte, contType string) error
	setObjectIfNoneMatch(ctx context.Context, obj string, data []byte, contType string) error
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
}

// consumeFunc is the signature of a function which can consume entries from the sequencer.
type consumeFunc func(ctx context.Context, from uint64, entries []storage.SequencedEntry) error

// Config holds AWS project and resource configuration for a storage instance.
type Config struct {
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
func New(ctx context.Context, cfg Config, opts ...func(*options.StorageOptions)) (*Storage, error) {
	opt := storage.ResolveStorageOptions(opts...)
	if opt.PushbackMaxOutstanding == 0 {
		opt.PushbackMaxOutstanding = DefaultPushbackMaxOutstanding
	}

	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default AWS configuration: %v", err)
	}
	c := s3.NewFromConfig(sdkConfig)

	seq, err := newMySQLSequencer(ctx, cfg.DSN, uint64(opt.PushbackMaxOutstanding), cfg.MaxOpenConns, cfg.MaxIdleConns)
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL sequencer: %v", err)
	}

	r := &Storage{
		objStore: &s3Storage{
			s3Client: c,
			bucket:   cfg.Bucket,
		},
		sequencer:   seq,
		newCP:       opt.NewCP,
		entriesPath: opt.EntriesPath,
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

				if _, err := r.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, r.integrate, false); err != nil {
					klog.Errorf("integrate: %v", err)
				}
			}()
		}
	}()

	return r, nil
}

// Add is the entrypoint for adding entries to a sequencing log.
func (s *Storage) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	return s.queue.Add(ctx, e)
}

func (s *Storage) ReadCheckpoint(ctx context.Context) ([]byte, error) {
	return s.get(ctx, layout.CheckpointPath)
}

func (s *Storage) ReadTile(ctx context.Context, l, i, sz uint64) ([]byte, error) {
	return s.get(ctx, layout.TilePath(l, i, sz))
}

func (s *Storage) ReadEntryBundle(ctx context.Context, i, sz uint64) ([]byte, error) {
	return s.get(ctx, s.entriesPath(i, sz))
}

// get returns the requested object.
//
// This is indended to be used to proxy read requests through the personality for debug/testing purposes.
func (s *Storage) get(ctx context.Context, path string) ([]byte, error) {
	d, err := s.objStore.getObject(ctx, path)
	return d, err
}

// init ensures that the storage represents a log in a valid state.
func (s *Storage) init(ctx context.Context) error {
	_, err := s.get(ctx, layout.CheckpointPath)
	if err != nil {
		// Do not use errors.Is. Keep errors.As to compare by type and not by value.
		var nske *types.NoSuchKey
		if errors.As(err, &nske) {
			// No checkpoint exists, do a forced (possibly empty) integration to create one in a safe
			// way (calling updateCP directly here would not be safe as it's outside the transactional
			// framework which prevents the tree from rolling backwards or otherwise forking).
			cctx, c := context.WithTimeout(ctx, 10*time.Second)
			defer c()
			if _, err := s.sequencer.consumeEntries(cctx, DefaultIntegrationSizeLimit, s.integrate, true); err != nil {
				return fmt.Errorf("forced integrate: %v", err)
			}
			return nil
		}
		return fmt.Errorf("failed to read checkpoint: %v", err)
	}

	return nil
}

func (s *Storage) updateCP(ctx context.Context, newSize uint64, newRoot []byte) error {
	cpRaw, err := s.newCP(newSize, newRoot)
	if err != nil {
		return fmt.Errorf("newCP: %v", err)
	}
	if err := s.objStore.setObject(ctx, layout.CheckpointPath, cpRaw, ckptContType); err != nil {
		return fmt.Errorf("writeCheckpoint: %v", err)
	}
	return nil

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

	return s.objStore.setObjectIfNoneMatch(ctx, tPath, data, logContType)
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
			data, err := s.objStore.getObject(ctx, objName)
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
func (s *Storage) getEntryBundle(ctx context.Context, bundleIndex uint64, logSize uint64) ([]byte, error) {
	objName := s.entriesPath(bundleIndex, logSize)
	data, err := s.objStore.getObject(ctx, objName)
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
func (s *Storage) setEntryBundle(ctx context.Context, bundleIndex uint64, logSize uint64, bundleRaw []byte) error {
	objName := s.entriesPath(bundleIndex, logSize)
	// Note that setObject does an idempotent interpretation of IfNoneMatch - it only
	// returns an error if the named object exists _and_ contains different data to what's
	// passed in here.
	if err := s.objStore.setObjectIfNoneMatch(ctx, objName, bundleRaw, logContType); err != nil {
		return fmt.Errorf("setObjectIfNoneMatch(%q): %v", objName, err)

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
				return s.updateCP(ctx, newSize, newRoot)
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
	if len(entries) == 0 {
		return nil
	}

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
			klog.V(1).Infof("In-memory bundle idx %d is full, attempting write to S3", bundleIndex)
			goSetEntryBundle(ctx, bundleIndex, fromSeq, bundleWriter.Bytes())
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
		goSetEntryBundle(ctx, bundleIndex, fromSeq, bundleWriter.Bytes())
	}
	return seqErr.Wait()
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
		return nil, fmt.Errorf("failed to connect to MySQL db(%q)): %v", dsn, err)
	}

	if maxOpenConns > 0 {
		dbPool.SetMaxOpenConns(maxOpenConns)
	}
	if maxIdleConns >= 0 {
		dbPool.SetMaxIdleConns(maxIdleConns)
	}

	if err := dbPool.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL db(%q): %v", dsn, err)
	}

	r := &mySQLSequencer{
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
// It creates tables if they don't exist already, and inserts zero values.
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
func (s *mySQLSequencer) initDB(ctx context.Context) error {
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
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}

	// Set default values for a newly initialised schema - these rows being present are a precondition for
	// sequencing and integration to occur.
	// Note that this will only succeed if no row exists, so there's no danger
	// of "resetting" an existing log.
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO SeqCoord (id, next) VALUES (0, 0)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO IntCoord (id, seq) VALUES (0, 0)`); err != nil {
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
	row := tx.QueryRowContext(ctx, "SELECT seq FROM IntCoord WHERE id = ? FOR UPDATE", 0)
	var fromSeq uint64
	if err := row.Scan(&fromSeq); err == sql.ErrNoRows {
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
	if err := f(ctx, uint64(fromSeq), entries); err != nil {
		return false, err
	}

	// consumeFunc was successful, so we can update our coordination row, and delete the row(s) for
	// the then consumed entries.
	if _, err := tx.ExecContext(ctx, "UPDATE IntCoord SET seq=? WHERE id=?", orderCheck, 0); err != nil {
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
