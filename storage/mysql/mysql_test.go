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

// Package mysql_test contains the tests for a MySQL-based storage implementation for Tessera.
// It requires a MySQL database to successfully run the tests. Otherwise, the tests in this file will be skipped.
//
// Sample command to start a local MySQL database using Docker:
// $ docker run --name test-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_tessera -d mysql
package mysql_test

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/storage/mysql"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	mysqlURI            = flag.String("mysql_uri", "root:root@tcp(localhost:3306)/test_tessera", "Connection string for a MySQL database")
	isMySQLTestOptional = flag.Bool("is_mysql_test_optional", true, "Boolean value to control whether the MySQL test is optional")

	testDB       *sql.DB
	noteSigner   note.Signer
	noteVerifier note.Verifier
)

const (
	testPrivateKey = "PRIVATE+KEY+transparency.dev/tessera/example+ae330e15+AXEwZQ2L6Ga3NX70ITObzyfEIketMr2o9Kc+ed/rt/QR"
	testPublicKey  = "transparency.dev/tessera/example+ae330e15+ASf4/L1zE859VqlfQgGzKy34l91Gl8W6wfwp+vKP62DW"
)

// TestMain checks whether the test MySQL database is available and starts the tests including database schema initialization.
// If is_mysql_test_optional is set to true and MySQL database cannot be opened or pinged, the test will fail immediately.
// Otherwise, the test will be skipped if the test is optional and the database is not available.
func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	db, err := sql.Open("mysql", *mysqlURI)
	if err != nil {
		if *isMySQLTestOptional {
			klog.Warning("MySQL not available, skipping all MySQL storage tests")
			return
		}
		klog.Fatalf("Failed to open MySQL test db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			klog.Warningf("Failed to close MySQL database: %v", err)
		}
	}()
	if err := db.PingContext(ctx); err != nil {
		if *isMySQLTestOptional {
			klog.Warning("MySQL not available, skipping all MySQL storage tests")
			return
		}
		klog.Fatalf("Failed to ping MySQL test db: %v", err)
	}
	testDB = db

	klog.Info("Successfully connected to MySQL test database")

	initDatabaseSchema(ctx)

	noteSigner, err = note.NewSigner(testPrivateKey)
	if err != nil {
		klog.Fatalf("Failed to create new signer: %v", err)
	}
	noteVerifier, err = note.NewVerifier(testPublicKey)
	if err != nil {
		klog.Fatalf("Failed to create new verifier: %v", err)
	}

	os.Exit(m.Run())
}

// initDatabaseSchema drops the tables and then imports the schema.
// A separate database connection is required since the schema file contains multiple statements.
// `multiStatements=true` in the data source name allows multiple statements in one query.
// This is not being used in the actual MySQL storage implementation.
func initDatabaseSchema(ctx context.Context) {
	dropTablesSQL := "DROP TABLE IF EXISTS `SequencingMetadata`, `Checkpoint`, `Subtree`, `TiledLeaves`"

	rawSchema, err := os.ReadFile("schema.sql")
	if err != nil {
		klog.Fatalf("Failed to read schema.sql: %v", err)
	}

	db, err := sql.Open("mysql", *mysqlURI+"?multiStatements=true")
	if err != nil {
		klog.Fatalf("Failed to connect to DB: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			klog.Warningf("Failed to close db: %v", err)
		}
	}()

	if _, err := db.ExecContext(ctx, dropTablesSQL); err != nil {
		klog.Fatalf("Failed to drop all tables: %v", err)
	}

	if _, err := db.ExecContext(ctx, string(rawSchema)); err != nil {
		klog.Fatalf("Failed to execute init database schema: %v", err)
	}
}

func TestNew(t *testing.T) {
	ctx := context.Background()

	for _, test := range []struct {
		name    string
		opts    []func(*tessera.StorageOptions)
		wantErr bool
	}{
		{
			name:    "no tessera.StorageOption",
			opts:    nil,
			wantErr: true,
		},
		{
			name: "standard tessera.WithCheckpointSignerVerifier",
			opts: []func(*tessera.StorageOptions){
				tessera.WithCheckpointSignerVerifier(noteSigner, noteVerifier),
			},
		},
		{
			name: "all tessera.StorageOption",
			opts: []func(*tessera.StorageOptions){
				tessera.WithCheckpointSignerVerifier(noteSigner, noteVerifier),
				tessera.WithBatching(1, 1*time.Second),
				tessera.WithPushback(10),
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := mysql.New(ctx, testDB, test.opts...)
			gotErr := err != nil
			if gotErr != test.wantErr {
				t.Errorf("got err: %v", err)
			}
		})
	}
}

func TestReadMissingTile(t *testing.T) {
	ctx := context.Background()
	s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name                string
		level, index, width uint64
	}{
		{
			name:  "0/0/0",
			level: 0, index: 0, width: 0,
		},
		{
			name:  "123/456/789",
			level: 123, index: 456, width: 789,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			tile, err := s.ReadTile(ctx, test.level, test.index, test.width)
			if err != nil {
				t.Errorf("got err: %v", err)
			}
			if tile != nil {
				t.Error("tile is not nil")
			}
		})
	}
}

func TestReadMissingEntryBundle(t *testing.T) {
	ctx := context.Background()
	s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name  string
		index uint64
	}{
		{
			name:  "0",
			index: 0,
		},
		{
			name:  "123456789",
			index: 123456789,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			entryBundle, err := s.ReadEntryBundle(ctx, test.index)
			if err != nil {
				t.Errorf("got err: %v", err)
			}
			if entryBundle != nil {
				t.Error("entryBundle is not nil")
			}
		})
	}
}

func TestParallelAdd(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name  string
		entry []byte
	}{
		{
			name:  "empty string entry",
			entry: []byte(""),
		},
		{
			name:  "123 string entry",
			entry: []byte("123"),
		},
		{
			name:  "empty byte",
			entry: []byte{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < 1024; i++ {
				go func() {
					if _, err := s.Add(ctx, tessera.NewEntry(test.entry))(); err != nil {
						t.Errorf("got err: %v", err)
					}
				}()
			}
		})
	}
}

func TestTileRoundTrip(t *testing.T) {
	ctx := context.Background()
	s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name  string
		entry []byte
	}{
		{
			name:  "empty string entry",
			entry: []byte(""),
		},
		{
			name:  "string entry",
			entry: []byte("I love Trillian Tessera"),
		},
		{
			name:  "empty byte",
			entry: []byte{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			cpSize := uint64(0)

			entryIndex, err := s.Add(ctx, tessera.NewEntry(test.entry))()
			if err != nil {
				t.Errorf("Add got err: %v", err)
			}

			for cpSize <= entryIndex {
				time.Sleep(100 * time.Millisecond)

				cpRaw, err := s.ReadCheckpoint(ctx)
				if err != nil {
					t.Errorf("ReadCheckpoint got err: %v", err)
				}
				cp, _, _, err := log.ParseCheckpoint(cpRaw, noteVerifier.Name(), noteVerifier)
				if err != nil {
					t.Errorf("log.ParseCheckpoint got err: %v", err)
				}
				cpSize = cp.Size
			}

			tileLevel, tileIndex, _, nodeIndex := layout.NodeCoordsToTileAddress(0, entryIndex)
			tileRaw, err := s.ReadTile(ctx, tileLevel, tileIndex, nodeIndex)
			if err != nil {
				t.Errorf("ReadTile got err: %v", err)
			}

			tile := api.HashTile{}
			if err := tile.UnmarshalText(tileRaw); err != nil {
				t.Errorf("failed to parse tile at index %d: %v", entryIndex, err)
			}
			nodes := tile.Nodes
			if len(nodes) == 0 {
				t.Error("no node found")
			} else {
				gotHash := nodes[nodeIndex]
				wantHash := rfc6962.DefaultHasher.HashLeaf(test.entry)
				if !bytes.Equal(gotHash, wantHash[:]) {
					t.Errorf("got hash %v want %v", gotHash, wantHash)
				}
			}
		})
	}
}

func TestEntryBundleRoundTrip(t *testing.T) {
	ctx := context.Background()
	s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name  string
		entry []byte
	}{
		{
			name:  "empty string entry",
			entry: []byte(""),
		},
		{
			name:  "string entry",
			entry: []byte("I love Trillian Tessera"),
		},
		{
			name:  "empty byte",
			entry: []byte{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			entryIndex, err := s.Add(ctx, tessera.NewEntry(test.entry))()
			if err != nil {
				t.Errorf("Add got err: %v", err)
			}
			entryBundleRaw, err := s.ReadEntryBundle(ctx, entryIndex/256)
			if err != nil {
				t.Errorf("ReadEntryBundle got err: %v", err)
			}

			bundle := api.EntryBundle{}
			if err := bundle.UnmarshalText(entryBundleRaw); err != nil {
				t.Errorf("failed to parse EntryBundle at index %d: %v", entryIndex, err)
			}
			gotEntries := bundle.Entries
			if len(gotEntries) == 0 {
				t.Error("no entry found")
			} else {
				if !bytes.Equal(bundle.Entries[entryIndex%256], test.entry) {
					t.Errorf("got entry %v want %v", bundle.Entries[0], test.entry)
				}
			}
		})
	}
}

func newTestMySQLStorage(t *testing.T, ctx context.Context) *mysql.Storage {
	t.Helper()

	s, err := mysql.New(ctx, testDB, tessera.WithCheckpointSignerVerifier(noteSigner, noteVerifier))
	if err != nil {
		t.Errorf("Failed to create mysql.Storage: %v", err)
	}

	return s
}
