// Copyright 2025 The Tessera authors. All Rights Reserved.
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

// Package aws contains an AWS-based antispam implementation for Tessera.
//
// A MySQL database provides a mechanism for maintaining an index of
// hash --> log position for detecting duplicate submissions.
package aws

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"k8s.io/klog/v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DefaultMaxBatchSize      = 64
	DefaultPushbackThreshold = 1024

	// SchemaCompatibilityVersion represents the expected version (e.g. layout & serialisation) of stored data.
	//
	// A binary built with a given version of the Tessera library is compatible with stored data created by a different version
	// of the library if and only if this value is the same as the compatibilityVersion stored in the Tessera table.
	//
	// NOTE: if changing this version, you need to consider whether end-users are going to update their schema instances to be
	// compatible with the new format, and provide a means to do it if so.
	SchemaCompatibilityVersion = 1
)

// AntispamOpts allows configuration of some tunable options.
type AntispamOpts struct {
	// MaxBatchSize is the largest number of mutations permitted in a single write operation when
	// updating the antispam index.
	//
	// Larger batches can enable (up to a point) higher throughput, but care should be taken not to
	// overload the database instance.
	MaxBatchSize uint

	// PushbackThreshold allows configuration of when to start responding to Add requests with pushback due to
	// the antispam follower falling too far behind.
	//
	// When the antispam follower is at least this many entries behind the size of the locally integrated tree,
	// the antispam decorator will return tessera.ErrPushback for every Add request.
	PushbackThreshold uint

	PushbackMaxOutstanding uint64
	MaxOpenConns           int
	MaxIdleConns           int
}

type AntispamStorage struct {
	opts AntispamOpts

	dbPool *sql.DB

	// pushBack is used to prevent the follower from getting too far underwater.
	// Populate dynamically will set this to true/false based on how far behind the follower is from the
	// currently integrated tree size.
	// When pushBack is true, the decorator will start returning ErrPushback to all calls.
	pushBack atomic.Bool

	numLookups atomic.Uint64
	numWrites  atomic.Uint64
	numHits    atomic.Uint64
}

// NewAntispam returns an antispam driver which uses a MySQL table to maintain a mapping of
// previously seen entries and their assigned indices.
//
// Note that the storage for this mapping is entirely separate and unconnected to the storage used for
// maintaining the Merkle tree.
//
// This functionality is experimental!
func NewAntispam(ctx context.Context, dsn string, opts AntispamOpts) (*AntispamStorage, error) {
	if opts.MaxBatchSize == 0 {
		opts.MaxBatchSize = DefaultMaxBatchSize
	}
	if opts.PushbackThreshold == 0 {
		opts.PushbackThreshold = DefaultPushbackThreshold
	}

	dbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL db: %v", err)
	}

	if opts.MaxOpenConns > 0 {
		dbPool.SetMaxOpenConns(opts.MaxOpenConns)
	}
	if opts.MaxIdleConns >= 0 {
		dbPool.SetMaxIdleConns(opts.MaxIdleConns)
	}

	if err := dbPool.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL db: %v", err)
	}

	r := &AntispamStorage{
		opts:   opts,
		dbPool: dbPool,
	}

	if err := r.initDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to initDB: %v", err)
	}
	if err := r.checkDataCompatibility(ctx); err != nil {
		return nil, fmt.Errorf("schema is not compatible with this version of the Tessera library: %v", err)
	}
	return r, nil
}

func (s *AntispamStorage) initDB(ctx context.Context) error {
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS AntispamMeta (
			id INT UNSIGNED NOT NULL,
			compatibilityVersion BIGINT UNSIGNED NOT NULL,
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS AntispamIDSeq (
			h TINYBLOB NOT NULL,
			idx BIGINT UNSIGNED NOT NULL,
			PRIMARY KEY (h(32))
		)`); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS AntispamFollowCoord (
			id INT UNSIGNED NOT NULL,
			nextIdx BIGINT UNSIGNED NOT NULL,
			PRIMARY KEY (id)
		)`); err != nil {
		return err
	}
	// Set default values for a newly initialised schema - these rows being present are a precondition for
	// sequencing and integration to occur.
	// Note that this will only succeed if no row exists, so there's no danger
	// of "resetting" an existing log.
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO AntispamMeta (id, compatibilityVersion) VALUES (0, ?)`, SchemaCompatibilityVersion); err != nil {
		return err
	}
	if _, err := s.dbPool.ExecContext(ctx,
		`INSERT IGNORE INTO AntispamFollowCoord (id, nextIdx) VALUES (0, 0)`); err != nil {
		return err
	}
	return nil
}

// checkDataCompatibility compares the Tessera library SchemaCompatibilityVersion with the one stored in the
// database, and returns an error if they are not identical.
func (s *AntispamStorage) checkDataCompatibility(ctx context.Context) error {
	row := s.dbPool.QueryRowContext(ctx, "SELECT compatibilityVersion FROM AntispamMeta WHERE id = 0")
	var gotVersion uint64
	if err := row.Scan(&gotVersion); err != nil {
		return fmt.Errorf("failed to read schema compatibility version from DB: %v", err)
	}

	if gotVersion != SchemaCompatibilityVersion {
		return fmt.Errorf("schema compatibilityVersion (%d) != library compatibilityVersion (%d)", gotVersion, SchemaCompatibilityVersion)
	}
	return nil
}

// index returns the index (if any) previously associated with the provided hash
func (d *AntispamStorage) index(ctx context.Context, h []byte) (*uint64, error) {
	d.numLookups.Add(1)
	row := d.dbPool.QueryRowContext(ctx, "SELECT idx FROM AntispamIDSeq WHERE h = ?", h)

	var idx uint64
	if err := row.Scan(&idx); err == sql.ErrNoRows {
		return nil, nil
	}
	d.numHits.Add(1)
	return &idx, nil
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
				//   2. We'll stop doing lookups for each submission, freeing up the DB to catch up.
				//
				// We may decide in the future that serving duplicate reads is more important than catching up as quickly
				// as possible, in which case we'd move this check down below the call to index.
				return func() (tessera.Index, error) { return tessera.Index{}, tessera.ErrPushback }
			}
			idx, err := d.index(ctx, e.Identity())
			if err != nil {
				return func() (tessera.Index, error) { return tessera.Index{}, err }
			}
			if idx != nil {
				return func() (tessera.Index, error) { return tessera.Index{Index: *idx, IsDup: true}, nil }
			}

			return delegate(ctx, e)
		}
	}
}

// Follower returns a follower which knows how to populate the antispam index.
//
// This implements tessera.Antispam.
func (d *AntispamStorage) Follower(b func([]byte) ([][]byte, error)) tessera.Follower {
	return &follower{
		as:           d,
		bundleHasher: b,
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

// follower is a struct which knows how to populate the antispam storage with identity hashes
// for entries in a log.
type follower struct {
	as           *AntispamStorage
	bundleHasher func([]byte) ([][]byte, error)
}

func (f *follower) Name() string {
	return "AWS antispam"
}

// Follow uses entry data from the log to populate the antispam storage.
func (f *follower) Follow(ctx context.Context, lr tessera.LogReader) {
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
			err = func() error {
				tx, err := f.as.dbPool.BeginTx(ctx, &sql.TxOptions{
					ReadOnly: false,
				})
				if err != nil {
					return err
				}

				row := tx.QueryRowContext(ctx, "SELECT nextIdx FROM AntispamFollowCoord WHERE id = 0")

				var followFrom uint64
				if err := row.Scan(&followFrom); err != nil {
					return err
				}

				if followFrom >= size {
					// Our view of the log is out of date, exit the busy loop and refresh it.
					workDone = false
					return nil
				}

				f.as.pushBack.Store(size-followFrom > uint64(f.as.opts.PushbackThreshold))

				// If this is the first time around the loop we need to start the stream of entries now that we know where we want to
				// start reading from:
				if entryReader == nil {
					next, st := lr.StreamEntries(ctx, followFrom)
					stop = st
					entryReader = newEntryStreamReader(next, f.bundleHasher)
				}

				bs := uint64(f.as.opts.MaxBatchSize)
				if r := size - followFrom; r < bs {
					bs = r
				}
				batch := make([][]byte, 0, bs)
				for i := range int(bs) {
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
				{
					sqlStr := "INSERT IGNORE INTO AntispamIDSeq (h, idx) VALUES "
					vals := make([]any, 0, 2*len(curEntries))
					for i, e := range curEntries {
						sqlStr += "(?, ?),"
						vals = append(vals, e, curIndex+uint64(i))
					}
					sqlStr = strings.TrimSuffix(sqlStr, ",")

					_, err := f.as.dbPool.ExecContext(ctx, sqlStr, vals...)
					if err != nil {
						return fmt.Errorf("failed to insert into AntispamIDSeq with query %q: %v", sqlStr, err)
					}
				}
				numAdded := uint64(len(curEntries))
				f.as.numWrites.Add(numAdded)
				nextIdx := uint64(followFrom + numAdded)

				// Insertion of dupe entries was successful, so update our follow coordination row:
				_, err = tx.ExecContext(ctx, "UPDATE AntispamFollowCoord SET nextIdx=? WHERE id=0", nextIdx)
				if err != nil {
					return fmt.Errorf("error updating AntispamFollowCoord: %v", err)
				}
				return tx.Commit()
			}()
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

// Position returns the index of the entry furthest from the start of the log which has been processed.
func (f *follower) Position(ctx context.Context) (uint64, error) {
	row := f.as.dbPool.QueryRowContext(ctx, "SELECT nextIdx FROM AntispamFollowCoord WHERE id = 0")

	var idx uint64
	if err := row.Scan(&idx); err != nil {
		return 0, fmt.Errorf("failed to read follow coordination info: %v", err)
	}

	return idx, nil
}
