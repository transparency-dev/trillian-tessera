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

// Package mysql contains a MySQL-based storage implementation for Tessera.
package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/storage"
	"k8s.io/klog/v2"
)

const (
	selectCheckpointByIDSQL          = "SELECT `note` FROM `Checkpoint` WHERE `id` = ?"
	selectCheckpointByIDForUpdateSQL = selectCheckpointByIDSQL + " FOR UPDATE"
	replaceCheckpointSQL             = "REPLACE INTO `Checkpoint` (`id`, `note`) VALUES (?, ?)"
	selectSubtreeByLevelAndIndexSQL  = "SELECT `nodes` FROM `Subtree` WHERE `level` = ? AND `index` = ?"
	replaceSubtreeSQL                = "REPLACE INTO `Subtree` (`level`, `index`, `nodes`) VALUES (?, ?, ?)"
	selectTiledLeavesSQL             = "SELECT `data` FROM `TiledLeaves` WHERE `tile_index` = ?"
	replaceTiledLeavesSQL            = "REPLACE INTO `TiledLeaves` (`tile_index`, `data`) VALUES (?, ?)"

	checkpointID        = 0
	entryBundleSize     = 256
	defaultBatchMaxSize = entryBundleSize
	defaultQueueMaxAge  = time.Second
)

// Storage is a MySQL-based storage implementation for Tessera.
type Storage struct {
	db    *sql.DB
	queue *storage.Queue

	newCheckpoint   tessera.NewCPFunc
	parseCheckpoint tessera.ParseCPFunc
}

// New creates a new instance of the MySQL-based Storage.
func New(ctx context.Context, db *sql.DB, opts ...func(*tessera.StorageOptions)) (*Storage, error) {
	opt := tessera.ResolveStorageOptions(&tessera.StorageOptions{
		BatchMaxAge:  defaultQueueMaxAge,
		BatchMaxSize: defaultBatchMaxSize,
	}, opts...)
	s := &Storage{
		db:              db,
		newCheckpoint:   opt.NewCP,
		parseCheckpoint: opt.ParseCP,
	}
	if err := s.db.Ping(); err != nil {
		klog.Errorf("Failed to ping database: %v", err)
		return nil, err
	}

	s.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, s.sequenceBatch)

	// Initialize checkpoint if there is no row in the Checkpoint table.
	if s.newCheckpoint != nil && s.parseCheckpoint != nil {
		if _, err := s.ReadCheckpoint(ctx); err != nil {
			if err != sql.ErrNoRows {
				klog.Errorf("Failed to read checkpoint: %v", err)
				return nil, err
			}

			klog.Infof("Initializing checkpoint")
			// Get a Tx for making transaction requests.
			tx, err := s.db.BeginTx(ctx, nil)
			if err != nil {
				return nil, err
			}
			// Defer a rollback in case anything fails.
			defer func() {
				if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
					klog.Errorf("Failed to rollback in write initial checkpoint: %v", err)
				}
			}()
			if err := s.writeCheckpoint(ctx, tx, 0, rfc6962.DefaultHasher.EmptyRoot()); err != nil {
				klog.Errorf("Failed to write initial checkpoint: %v", err)
				return nil, err
			}
			// Commit the transaction.
			if err := tx.Commit(); err != nil {
				return nil, err
			}
		}
	}

	return s, nil
}

// ReadCheckpoint returns the latest stored checkpoint.
func (s *Storage) ReadCheckpoint(ctx context.Context) ([]byte, error) {
	row := s.db.QueryRowContext(ctx, selectCheckpointByIDSQL, checkpointID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var checkpoint []byte
	return checkpoint, row.Scan(&checkpoint)
}

// writeCheckpoint stores the log signed checkpoint.
func (s *Storage) writeCheckpoint(ctx context.Context, tx *sql.Tx, size uint64, rootHash []byte) error {
	rawCheckpoint, err := s.newCheckpoint(size, rootHash)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, replaceCheckpointSQL, checkpointID, rawCheckpoint); err != nil {
		klog.Errorf("Failed to execute replaceCheckpointSQL: %v", err)
		return err
	}

	return nil
}

// ReadTile returns a full tile or a partial tile at the given level, index and width.
//
// TODO: Handle the following scenarios:
// 1. Full tile request with full tile output: Return full tile.
// 2. Full tile request with partial tile output: Return error.
// 3. Partial tile request with full/larger partial tile output: Return trimmed partial tile with correct tile width.
// 4. Partial tile request with partial tile (same width) output: Return partial tile.
// 5. Partial tile request with smaller partial tile output: Return error.
func (s *Storage) ReadTile(ctx context.Context, level, index, width uint64) ([]byte, error) {
	row := s.db.QueryRowContext(ctx, selectSubtreeByLevelAndIndexSQL, level, index)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var tile []byte
	if err := row.Scan(&tile); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	// Return nil when returning a partial tile on a full tile request.
	if width == 256 && len(tile)/32 != int(width) {
		return nil, nil
	}

	return tile, nil
}

// writeTile replaces the tile nodes at the given level and index.
func (s *Storage) writeTile(ctx context.Context, tx *sql.Tx, level, index uint64, nodes []byte) error {
	if _, err := tx.ExecContext(ctx, replaceSubtreeSQL, level, index, nodes); err != nil {
		klog.Errorf("Failed to execute replaceSubtreeSQL: %v", err)
		return err
	}

	return nil
}

// ReadEntryBundle returns the log entries at the given index.
//
// TODO: Handle the following scenarios:
// 1. Full tile request with full tile output: Return full tile.
// 2. Full tile request with partial tile output: Return error.
// 3. Partial tile request with full/larger partial tile output: Return trimmed partial tile with correct tile width.
// 4. Partial tile request with partial tile (same width) output: Return partial tile.
// 5. Partial tile request with smaller partial tile output: Return error.
func (s *Storage) ReadEntryBundle(ctx context.Context, index uint64) ([]byte, error) {
	row := s.db.QueryRowContext(ctx, selectTiledLeavesSQL, index)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var entryBundle []byte
	if err := row.Scan(&entryBundle); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return entryBundle, nil
}

func (s *Storage) writeEntryBundle(ctx context.Context, tx *sql.Tx, index uint64, entryBundle []byte) error {
	if _, err := tx.ExecContext(ctx, replaceTiledLeavesSQL, index, entryBundle); err != nil {
		klog.Errorf("Failed to execute replaceTiledLeavesSQL: %v", err)
		return err
	}

	return nil
}

// Add is the entrypoint for adding entries to a sequencing log.
func (s *Storage) Add(ctx context.Context, entry *tessera.Entry) (uint64, error) {
	// TODO(#21): Return index if the value is already stored.

	return s.queue.Add(ctx, entry)()
}

// sequenceBatch writes the entries from the provided batch into the entry bundle files of the log.
//
// This func starts filling entries bundles at the next available slot in the log, ensuring that the
// sequenced entries are contiguous from the zeroth entry (i.e left-hand dense).
// We try to minimise the number of partially complete entry bundles by writing entries in chunks rather
// than one-by-one.
//
// TODO(#21): Separate sequencing and integration for better performance.
func (s *Storage) sequenceBatch(ctx context.Context, entries []*tessera.Entry) error {
	// Return when there is no entry to sequence.
	if len(entries) == 0 {
		return nil
	}

	// Get a Tx for making transaction requests.
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			klog.Errorf("Failed to rollback in sequenceBatch: %v", err)
		}
	}()

	// Get tree size from checkpoint. Note that "SELECT ... FOR UPDATE" is used for row-level locking.
	// TODO(#21): Optimize how we get the tree size without parsing and verifying the checkpoints every time.
	row := tx.QueryRowContext(ctx, selectCheckpointByIDForUpdateSQL, checkpointID)
	if err := row.Err(); err != nil {
		return err
	}
	var rawCheckpoint []byte
	if err := row.Scan(&rawCheckpoint); err != nil {
		return fmt.Errorf("failed to read checkpoint: %w", err)
	}
	checkpoint, err := s.parseCheckpoint(rawCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to verify checkpoint: %w", err)
	}

	// Integrate the new entries into the entry bundle (TiledLeaves table) and tile (Subtree table).
	if err := s.integrate(ctx, tx, checkpoint.Size, entries); err != nil {
		return fmt.Errorf("failed to integrate: %w", err)
	}

	// Commit the transaction.
	return tx.Commit()
}

// integrate incorporates the provided entries into the log starting at fromSeq.
func (s *Storage) integrate(ctx context.Context, tx *sql.Tx, fromSeq uint64, entries []*tessera.Entry) error {
	tb := storage.NewTreeBuilder(func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		hashTiles := make([]*api.HashTile, len(tileIDs))
		if len(tileIDs) == 0 {
			return hashTiles, nil
		}

		// Build the SQL and args to fetch the hash tiles.
		var sql strings.Builder
		args := make([]any, 0, len(tileIDs)*2)
		for i, id := range tileIDs {
			if i != 0 {
				sql.WriteString(" UNION ALL ")
			}
			_, err := sql.WriteString(selectSubtreeByLevelAndIndexSQL)
			if err != nil {
				return nil, err
			}
			args = append(args, id.Level, id.Index)
		}

		rows, err := tx.QueryContext(ctx, sql.String(), args...)
		if err != nil {
			return nil, fmt.Errorf("failed to query the hash tiles with SQL (%s): %w", sql.String(), err)
		}
		defer func() {
			if err := rows.Close(); err != nil {
				klog.Warningf("Failed to close the rows: %v", err)
			}
		}()

		i := 0
		for rows.Next() {
			var tile []byte
			if err := rows.Scan(&tile); err != nil {
				return nil, fmt.Errorf("rows.Scan: %w", err)
			}
			t := &api.HashTile{}
			if err := t.UnmarshalText(tile); err != nil {
				return nil, fmt.Errorf("api.HashTile.unmarshalText: %w", err)
			}
			hashTiles[i] = t
			i++
		}
		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("rows.Err: %w", err)
		}

		return hashTiles, nil
	})

	sequencedEntries := make([]storage.SequencedEntry, len(entries))
	// Assign provisional sequence numbers to entries.
	// We need to do this here in order to support serialisations which include the log position.
	for i, e := range entries {
		sequencedEntries[i] = storage.SequencedEntry{
			BundleData: e.MarshalBundleData(fromSeq + uint64(i)),
			LeafHash:   e.LeafHash(),
		}
	}

	// Add sequenced entries to entry bundles.
	bundleIndex, entriesInBundle := fromSeq/entryBundleSize, fromSeq%entryBundleSize
	bundleWriter := &bytes.Buffer{}

	// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
	if entriesInBundle > 0 {
		row := tx.QueryRowContext(ctx, selectTiledLeavesSQL, bundleIndex)
		if err := row.Err(); err != nil {
			return err
		}

		var partialEntryBundle []byte
		if err := row.Scan(&partialEntryBundle); err != nil {
			return fmt.Errorf("row.Scan: %w", err)
		}

		if _, err := bundleWriter.Write(partialEntryBundle); err != nil {
			return fmt.Errorf("bundleWriter: %w", err)
		}
	}

	// Add new entries to the bundle.
	for _, e := range sequencedEntries {
		if _, err := bundleWriter.Write(e.BundleData); err != nil {
			return fmt.Errorf("bundleWriter.Write: %w", err)
		}
		entriesInBundle++

		// This bundle is full, so we need to write it out.
		if entriesInBundle == entryBundleSize {
			if err := s.writeEntryBundle(ctx, tx, bundleIndex, bundleWriter.Bytes()); err != nil {
				return fmt.Errorf("writeEntryBundle: %w", err)
			}

			// Prepare the next entry bundle for any remaining entries in the batch.
			bundleIndex++
			entriesInBundle = 0
			bundleWriter = &bytes.Buffer{}
		}
	}

	// If we have a partial bundle remaining once we've added all the entries from the batch,
	// this needs writing out too.
	if entriesInBundle > 0 {
		if err := s.writeEntryBundle(ctx, tx, bundleIndex, bundleWriter.Bytes()); err != nil {
			return fmt.Errorf("writeEntryBundle: %w", err)
		}
	}

	newSize, newRoot, tiles, err := tb.Integrate(ctx, fromSeq, sequencedEntries)
	if err != nil {
		return fmt.Errorf("tb.Integrate: %v", err)
	}
	for k, v := range tiles {
		nodes, err := v.MarshalText()
		if err != nil {
			return err
		}

		if err := s.writeTile(ctx, tx, uint64(k.Level), k.Index, nodes); err != nil {
			return fmt.Errorf("failed to set tile(%v): %w", k, err)
		}
	}

	// Write new checkpoint.
	if err := s.writeCheckpoint(ctx, tx, newSize, newRoot); err != nil {
		return fmt.Errorf("writeCheckpoint: %w", err)
	}
	return nil
}
