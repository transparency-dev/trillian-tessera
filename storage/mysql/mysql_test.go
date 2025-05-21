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
package mysql

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/tessera"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tessera/core"
	"golang.org/x/mod/sumdb/note"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

var (
	mysqlURI            = flag.String("mysql_uri", "root:root@tcp(localhost:3306)/test_tessera", "Connection string for a MySQL database")
	isMySQLTestOptional = flag.Bool("is_mysql_test_optional", true, "Boolean value to control whether the MySQL test is optional")

	testDB     *sql.DB
	noteSigner note.Signer
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

	os.Exit(m.Run())
}

// initDatabaseSchema drops the tables and then imports the schema.
// A separate database connection is required since the schema file contains multiple statements.
// `multiStatements=true` in the data source name allows multiple statements in one query.
// This is not being used in the actual MySQL storage implementation.
func initDatabaseSchema(ctx context.Context) {
	dropTablesSQL := "DROP TABLE IF EXISTS `Checkpoint`, `Subtree`, `TiledLeaves`, `TreeState`"

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

func TestAppend(t *testing.T) {
	ctx := context.Background()

	for _, test := range []struct {
		name    string
		opts    *tessera.AppendOptions
		wantErr bool
	}{
		{
			name:    "no tessera.AppendOptions",
			opts:    tessera.NewAppendOptions(),
			wantErr: true,
		},
		{
			name: "standard tessera.WithCheckpointSigner",
			opts: tessera.NewAppendOptions().WithCheckpointSigner(noteSigner),
		},
		{
			name: "all tessera.AppendOptions",
			opts: tessera.NewAppendOptions().
				WithCheckpointSigner(noteSigner).
				WithBatching(1, 1*time.Second).
				WithPushback(10),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			drv, err := New(ctx, testDB)
			if err != nil {
				t.Fatalf("New: %v", err)
			}
			_, _, _, err = tessera.NewAppender(ctx, drv, test.opts)
			gotErr := err != nil
			if gotErr != test.wantErr {
				t.Errorf("got err: %v", err)
			}
		})
	}
}

func TestGetTile(t *testing.T) {
	ctx := context.Background()
	addFn, r, _ := newTestMySQLStorage(t, ctx)

	awaiter := tessera.NewPublicationAwaiter(ctx, r.ReadCheckpoint, 10*time.Millisecond)

	treeSize := 258

	wg := errgroup.Group{}
	for i := range treeSize {
		wg.Go(
			func() error {
				_, _, err := awaiter.Await(ctx, addFn(ctx, core.NewEntry(fmt.Appendf(nil, "TestGetTile %d", i))))
				if err != nil {
					return fmt.Errorf("failed to prep test with entry: %v", err)
				}
				return nil
			})
	}
	if err := wg.Wait(); err != nil {
		t.Fatalf("Failed to set up database with required leaves: %v", err)
	}

	for _, test := range []struct {
		name         string
		level, index uint64
		p            uint8
		wantEntries  int
		wantNotFound bool
	}{
		{
			name:  "requested partial tile for a complete tile",
			level: 0, index: 0, p: 10,
			wantEntries:  layout.TileWidth,
			wantNotFound: false,
		},
		{
			name:  "too small but that's ok",
			level: 0, index: 1, p: layout.PartialTileSize(0, 1, uint64(treeSize-1)),
			wantEntries:  2,
			wantNotFound: false,
		},
		{
			name:  "just right",
			level: 0, index: 1, p: layout.PartialTileSize(0, 1, uint64(treeSize)),
			wantEntries:  2,
			wantNotFound: false,
		},
		{
			name:  "too big",
			level: 0, index: 1, p: layout.PartialTileSize(0, 1, uint64(treeSize+1)),
			wantNotFound: true,
		},
		{
			name:  "level 1 too small",
			level: 1, index: 0, p: layout.PartialTileSize(1, 0, uint64(treeSize-1)),
			wantEntries:  1,
			wantNotFound: false,
		},
		{
			name:  "level 1 just right",
			level: 1, index: 0, p: layout.PartialTileSize(1, 0, uint64(treeSize)),
			wantEntries:  1,
			wantNotFound: false,
		},
		{
			name:  "level 1 too big",
			level: 1, index: 0, p: layout.PartialTileSize(1, 0, 550),
			wantNotFound: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			tile, err := r.ReadTile(ctx, test.level, test.index, test.p)
			if err != nil {
				if notFound, wantNotFound := errors.Is(err, fs.ErrNotExist), test.wantNotFound; notFound != wantNotFound {
					t.Errorf("wantNotFound %v but notFound %v", wantNotFound, notFound)
				}
				if test.wantNotFound {
					return
				}
				t.Errorf("got err: %v", err)
			}
			numEntries := len(tile) / sha256.Size
			if got, want := numEntries, test.wantEntries; got != want {
				t.Errorf("got %d entries, but want %d", got, want)
			}
		})
	}
}

func TestReadMissingTile(t *testing.T) {
	ctx := context.Background()
	_, r, _ := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name         string
		level, index uint64
		p            uint8
	}{
		{
			name:  "0/0/0",
			level: 0, index: 0, p: 0,
		},
		{
			name:  "123/456/789",
			level: 123, index: 456, p: 789 % layout.TileWidth,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			tile, err := r.ReadTile(ctx, test.level, test.index, test.p)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					// this is success for this test
					return
				}
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
	_, r, _ := newTestMySQLStorage(t, ctx)

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
			entryBundle, err := r.ReadEntryBundle(ctx, test.index, uint8(test.index%layout.TileWidth))
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					// this is success for this test
					return
				}
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
	addFn, _, _ := newTestMySQLStorage(t, ctx)

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
			eG := errgroup.Group{}
			for range 1024 {
				eG.Go(func() error {
					_, err := addFn(ctx, core.NewEntry(test.entry))()
					return err
				})
			}
			if err := eG.Wait(); err != nil {
				t.Errorf("got err: %v", err)
			}
		})
	}
}

func TestTileRoundTrip(t *testing.T) {
	ctx := context.Background()
	addFn, r, _ := newTestMySQLStorage(t, ctx)

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
			entryIndex, err := addFn(ctx, core.NewEntry(test.entry))()
			if err != nil {
				t.Errorf("Add got err: %v", err)
			}

			tileLevel, tileIndex, _, nodeIndex := layout.NodeCoordsToTileAddress(0, entryIndex.Index)
			tileRaw, err := r.ReadTile(ctx, tileLevel, tileIndex, uint8(nodeIndex+1))
			if err != nil {
				t.Errorf("ReadTile got err: %v", err)
			}

			tile := api.HashTile{}
			if err := tile.UnmarshalText(tileRaw); err != nil {
				t.Errorf("failed to parse tile at index %d: %v", entryIndex.Index, err)
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
	addFn, r, _ := newTestMySQLStorage(t, ctx)

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
			entryIndex, err := addFn(ctx, core.NewEntry(test.entry))()
			if err != nil {
				t.Errorf("Add got err: %v", err)
			}
			idx := entryIndex.Index
			entryBundleRaw, err := r.ReadEntryBundle(ctx, idx/layout.EntryBundleWidth, layout.PartialTileSize(0, idx, idx+1))
			if err != nil {
				t.Fatalf("ReadEntryBundle got err: %v", err)
			}

			bundle := api.EntryBundle{}
			if err := bundle.UnmarshalText(entryBundleRaw); err != nil {
				t.Errorf("failed to parse EntryBundle at index %d: %v", entryIndex.Index, err)
			}
			gotEntries := bundle.Entries
			if len(gotEntries) == 0 {
				t.Error("no entry found")
			} else {
				if !bytes.Equal(bundle.Entries[idx%layout.EntryBundleWidth], test.entry) {
					t.Errorf("got entry %v want %v", bundle.Entries[0], test.entry)
				}
			}
		})
	}
}

func TestStreamEntries(t *testing.T) {
	ctx := context.Background()
	_, _, s := newTestMySQLStorage(t, ctx)

	for _, test := range []struct {
		name            string
		treeSize        uint64
		growTreeSize    uint64
		growAfterBundle uint64
		streamFrom      uint64
	}{
		{
			name:       "stream from zero to end of tree",
			treeSize:   10 * layout.EntryBundleWidth,
			streamFrom: 0,
		}, {
			name:            "stream from zero to end of tree, grow tree while streaming",
			treeSize:        10 * layout.EntryBundleWidth,
			growTreeSize:    20 * layout.EntryBundleWidth,
			growAfterBundle: 5,
			streamFrom:      0,
		}, {
			name:       "stream from middle to end of tree (aligned)",
			treeSize:   10 * layout.EntryBundleWidth,
			streamFrom: 5 * layout.EntryBundleWidth,
		}, {
			name:            "stream from middle to end of tree (aligned), grow while streaming",
			treeSize:        10 * layout.EntryBundleWidth,
			growTreeSize:    20 * layout.EntryBundleWidth,
			growAfterBundle: 18,
			streamFrom:      8 * layout.EntryBundleWidth,
		}, {
			name:       "stream from middle to end of tree (unaligned)",
			treeSize:   10 * layout.EntryBundleWidth,
			streamFrom: 5*layout.EntryBundleWidth + 17,
		}, {
			name:            "stream from middle to end of tree (unaligned), grow while streaming",
			treeSize:        10*layout.EntryBundleWidth + 46,
			growTreeSize:    20 * layout.EntryBundleWidth,
			growAfterBundle: 19,
			streamFrom:      5*layout.EntryBundleWidth + 17,
		}, {
			name:            "stream from middle, grow when whole previous log was streamed",
			treeSize:        10*layout.EntryBundleWidth + 41,
			growTreeSize:    20 * layout.EntryBundleWidth,
			growAfterBundle: 21,
			streamFrom:      5*layout.EntryBundleWidth + 17,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			populateEntries(t, test.treeSize, s)
			next, cancel := s.StreamEntries(context.Background(), test.streamFrom)
			wantBundle := uint64(test.streamFrom / layout.EntryBundleWidth)
			nBundles := test.treeSize / layout.EntryBundleWidth
			for i := 0; ; i++ {
				ri, b, err := next()
				if err != nil {
					t.Fatalf("expecting bundle %d, next: %v", wantBundle, err)
				}
				if ri.Index != wantBundle {
					t.Fatalf("got bundle %d, want %d", ri.Index, wantBundle)
				}
				verifyBundle(t, wantBundle, b)

				if i == int(test.growAfterBundle) && test.growTreeSize > 0 {
					populateEntries(t, test.growTreeSize, s)
					nBundles = test.growTreeSize / layout.EntryBundleWidth
				}

				wantBundle++

				t.Logf("want %d N %d", wantBundle, nBundles)
				if wantBundle >= nBundles {
					cancel()
					break
				}
			}
		})
	}
}

func verifyBundle(t *testing.T, bundleIndex uint64, b []byte) {
	t.Helper()
	for i, got := range b {
		if got, want := byte(got), byte(bundleIndex+uint64(i)); got != want {
			t.Fatalf("bundleIndex %d, entry %d is %d, want %d", bundleIndex, i, got, want)
		}
	}
}

// populateEntries creates entry bundles for a tree of the specified size, and sets the treeState to that size.
// Note that it DOES NOT currently create any tiles.
//
// To help with round-trip testing, bundles are filled with one byte per "entry" according to the size of the
// bundle implied by the provided tree size. Entry bytes within a bundle start with a value of byte(bundleIndex),
// and are increased by one for each successive entry.
func populateEntries(t *testing.T, size uint64, s *Storage) {
	t.Helper()

	tx, err := s.db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	defer func() {
		if err := tx.Commit(); err != nil {
			t.Fatalf("Commit: %v", err)
		}
	}()

	for i, rem := 0, size; rem > 0; {
		bundleIdx := i / layout.EntryBundleWidth
		bs := rem
		if bs > layout.EntryBundleWidth {
			bs = uint64(layout.EntryBundleWidth)
		}
		b := make([]byte, bs)
		for x := range b {
			b[x] = byte(bundleIdx + x)
			i++
			rem--
		}
		t.Logf("Write %d.%d", bundleIdx, bs)
		if err := s.writeEntryBundle(context.Background(), tx, uint64(bundleIdx), uint32(bs), b); err != nil {
			t.Fatalf("writeEntryBundle(@%d.%d): %v", bundleIdx, bs, err)
		}
	}
	if err := s.writeTreeState(context.Background(), tx, size, []byte("root")); err != nil {
		t.Fatalf("writeTreeState: %v", err)
	}
}

func newTestMySQLStorage(t *testing.T, ctx context.Context) (core.AddFn, tessera.LogReader, *Storage) {
	t.Helper()
	initDatabaseSchema(ctx)

	s, err := New(ctx, testDB)
	if err != nil {
		t.Fatalf("Failed to create mysql.Storage: %v", err)
	}

	a, _, r, err := tessera.NewAppender(ctx, s, tessera.NewAppendOptions().
		WithCheckpointSigner(noteSigner).
		WithCheckpointInterval(time.Second).
		WithBatching(128, 100*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	return a.Add, r, s
}
