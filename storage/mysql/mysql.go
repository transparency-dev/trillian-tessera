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
}

// New creates a new instance of the MySQL-based Storage.
func New(ctx context.Context, db *sql.DB) (*Storage, error) {
	s := &Storage{
		db: db,
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
		if err := s.writeCheckpoint(ctx, []byte("")); err != nil {
			klog.Errorf("Failed to write initial checkpoint: %v", err)
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

// writeCheckpoint stores a raw log checkpoint.
func (s *Storage) writeCheckpoint(ctx context.Context, rawCheckpoint []byte) error {
	// Get a Tx for making transaction requests.
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			klog.Errorf("Failed to rollback in writeCheckpoint: %v", err)
		}
	}()

	if _, err := tx.ExecContext(ctx, replaceCheckpointSQL, checkpointID, rawCheckpoint); err != nil {
		klog.Errorf("Failed to execute replaceCheckpointSQL: %v", err)
		return err
	}

	// Commit the transaction.
	return tx.Commit()
}

// ReadTile returns a full tile or a partial tile at the given level and index.
func (s *Storage) ReadTile(ctx context.Context, level, index uint64) ([]byte, error) {
	row := s.db.QueryRowContext(ctx, selectSubtreeByLevelAndIndexSQL, level, index)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var tile []byte
	return tile, row.Scan(&tile)
}

// ReadEntryBundle returns the log entries at the given index.
func (s *Storage) ReadEntryBundle(ctx context.Context, index uint64) ([]byte, error) {
	row := s.db.QueryRowContext(ctx, selectTiledLeavesSQL, index)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var entryBundle []byte
	return entryBundle, row.Scan(&entryBundle)
}
