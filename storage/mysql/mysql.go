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
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/transparency-dev/formats/log"
	tessera "github.com/transparency-dev/trillian-tessera"
	"k8s.io/klog/v2"
)

const (
	selectCheckpointByIDSQL         = "SELECT `note` FROM `Checkpoint` WHERE `id` = ?"
	replaceCheckpointSQL            = "REPLACE INTO `Checkpoint` (`id`, `note`) VALUES (?, ?)"
	selectSubtreeByLevelAndIndexSQL = "SELECT `nodes` FROM `Subtree` WHERE `level` = ? AND `index` = ?"
	selectTiledLeavesSQL            = "SELECT `data` FROM `TiledLeaves` WHERE `tile_index` = ?"

	checkpointID = 0
)

// Storage is a MySQL-based storage implementation for Tessera.
type Storage struct {
	db *sql.DB

	newCheckpoint tessera.NewCPFunc
}

type WriteCheckpointFunc func(ctx context.Context, checkpoint log.Checkpoint) error

// New creates a new instance of the MySQL-based Storage.
func New(ctx context.Context, db *sql.DB, opts ...func(*tessera.StorageOptions)) (*Storage, error) {
	opt := tessera.ResolveStorageOptions(nil, opts...)
	s := &Storage{
		db:            db,
		newCheckpoint: opt.NewCP,
	}
	if err := s.db.Ping(); err != nil {
		klog.Errorf("Failed to ping database: %v", err)
		return nil, err
	}

	// Initialize checkpoint if there is no row in the Checkpoint table.
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
		if err := s.writeCheckpoint(ctx, tx, 0, []byte("")); err != nil {
			klog.Errorf("Failed to write initial checkpoint: %v", err)
			return nil, err
		}
		// Commit the transaction.
		if err := tx.Commit(); err != nil {
			return nil, err
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

// ReadTile returns a full tile or a partial tile at the given level and index.
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
	return entryBundle, row.Scan(&entryBundle)
}
