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

// This the tests for a MySQL+S3 AWS Tessera implementation.  It requires a
// MySQL database to successfully run the MySQL tests, otherwise they are
// skipped.  Run tests with `-parallel=1` to avoid concurent tests on the same
// database, and specifically runs of `mustDropTables`.
//
// Sample command to start a local MySQL database using Docker:
// $ docker run --name test-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_tessera -d mysql
package aws

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/tessera"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tessera/core"
	storage "github.com/transparency-dev/tessera/storage/internal"
	"k8s.io/klog/v2"
)

var (
	mySQLURI            = flag.String("mysql_uri", "root:root@tcp(localhost:3306)/test_tessera", "Connection string for a MySQL database")
	isMySQLTestOptional = flag.Bool("is_mysql_test_optional", true, "Boolean value to control whether the MySQL test is optional")
)

// TestMain inits flags and runs tests.
func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	// m.Run() will parse flags
	os.Exit(m.Run())
}

// canSkipMySQLTest checks if the test MySQL db is available and, if not, if the test can be skipped.
//
// Use this method before every MySQL test, and if it returns true, skip the test.
//
// If is_mysql_test_optional is set to true and MySQL database cannot be opened or pinged,
// the test will fail immediately. Otherwise, the test will be skipped if the test is optional
// and the database is not available.
func canSkipMySQLTest(t *testing.T, ctx context.Context) bool {
	t.Helper()

	db, err := sql.Open("mysql", *mySQLURI)
	if err != nil {
		if *isMySQLTestOptional {
			return true
		}
		t.Fatalf("failed to open MySQL test db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("failed to close MySQL database: %v", err)
		}
	}()
	if err := db.PingContext(ctx); err != nil {
		if *isMySQLTestOptional {
			return true
		}
		t.Fatalf("failed to ping MySQL test db: %v", err)
	}
	return false
}

// mustDropTables drops the `Seq`, `SeqCoord` and `IntCoord` tables.
// Call this function before every MySQL test.
func mustDropTables(t *testing.T, ctx context.Context) {
	t.Helper()

	db, err := sql.Open("mysql", *mySQLURI)
	if err != nil {
		t.Fatalf("failed to connect to db: %v", *mySQLURI)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("failed to close db: %v", err)
		}
	}()

	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `Seq`, `SeqCoord`, `IntCoord`"); err != nil {
		t.Fatalf("failed to drop all tables: %v", err)
	}
}

func TestMySQLSequencerAssignEntries(t *testing.T) {
	ctx := context.Background()
	if canSkipMySQLTest(t, ctx) {
		klog.Warningf("MySQL not available, skipping %s", t.Name())
		t.Skip("MySQL not available, skipping test")
	}
	// Clean tables in case there's already something in there.
	mustDropTables(t, ctx)

	seq, err := newMySQLSequencer(ctx, *mySQLURI, 1000, 0, 0)
	if err != nil {
		t.Fatalf("newMySQLSequencer: %v", err)
	}

	want := uint64(0)
	for chunks := range 10 {
		entries := []*core.Entry{}
		for i := range 10 + chunks {
			entries = append(entries, core.NewEntry(fmt.Appendf(nil, "item %d/%d", chunks, i)))
		}
		if err := seq.assignEntries(ctx, entries); err != nil {
			t.Fatalf("assignEntries: %v", err)
		}
		for i, e := range entries {
			if got := *e.Index(); got != want {
				t.Errorf("Chunk %d entry %d got seq %d, want %d", chunks, i, got, want)
			}
			want++
		}
	}
}

func TestMySQLSequencerPushback(t *testing.T) {
	ctx := context.Background()
	if canSkipMySQLTest(t, ctx) {
		klog.Warningf("MySQL not available, skipping %s", t.Name())
		t.Skip("MySQL not available, skipping test")
	}
	// Clean tables in case there's already something in there.
	mustDropTables(t, ctx)

	for _, test := range []struct {
		name           string
		threshold      uint64
		initialEntries int
		wantPushback   bool
	}{
		{
			name:           "no pushback: num < threshold",
			threshold:      10,
			initialEntries: 5,
		},
		{
			name:           "no pushback: num = threshold",
			threshold:      10,
			initialEntries: 10,
		},
		{
			name:           "pushback: initial > threshold",
			threshold:      10,
			initialEntries: 15,
			wantPushback:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			mustDropTables(t, ctx)

			seq, err := newMySQLSequencer(ctx, *mySQLURI, test.threshold, 0, 0)
			if err != nil {
				t.Fatalf("newMySQLSequencer: %v", err)
			}
			// Set up the test scenario with the configured number of initial outstanding entries
			entries := []*core.Entry{}
			for i := range test.initialEntries {
				entries = append(entries, core.NewEntry(fmt.Appendf(nil, "initial item %d", i)))
			}
			if err := seq.assignEntries(ctx, entries); err != nil {
				t.Fatalf("initial assignEntries: %v", err)
			}

			// Now perform the test with a single additional entry to check for pushback
			entries = []*core.Entry{core.NewEntry([]byte("additional"))}
			err = seq.assignEntries(ctx, entries)
			if gotPushback := errors.Is(err, tessera.ErrPushback); gotPushback != test.wantPushback {
				t.Fatalf("assignEntries: got pushback %t (%v), want pushback: %t", gotPushback, err, test.wantPushback)
			} else if !gotPushback && err != nil {
				t.Fatalf("assignEntries: %v", err)
			}
		})
	}
}

func TestMySQLSequencerRoundTrip(t *testing.T) {
	ctx := context.Background()
	if canSkipMySQLTest(t, ctx) {
		klog.Warningf("MySQL not available, skipping %s", t.Name())
		t.Skip("MySQL not available, skipping test")
	}
	// Clean tables in case there's already something in there.
	mustDropTables(t, ctx)

	s, err := newMySQLSequencer(ctx, *mySQLURI, 1000, 0, 0)
	if err != nil {
		t.Fatalf("newMySQLSequencer: %v", err)
	}

	seq := 0
	wantEntries := []storage.SequencedEntry{}
	for chunks := range 10 {
		entries := []*core.Entry{}
		for range 10 + chunks {
			e := core.NewEntry(fmt.Appendf(nil, "item %d", seq))
			entries = append(entries, e)
			wantEntries = append(wantEntries, storage.SequencedEntry{
				BundleData: e.MarshalBundleData(uint64(seq)),
				LeafHash:   e.LeafHash(),
			})
			seq++
		}
		if err := s.assignEntries(ctx, entries); err != nil {
			t.Fatalf("assignEntries: %v", err)
		}
	}

	seenIdx := uint64(0)
	f := func(_ context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
		if fromSeq != seenIdx {
			return nil, fmt.Errorf("f called with fromSeq %d, want %d", fromSeq, seenIdx)
		}
		for i, e := range entries {

			if got, want := e, wantEntries[i]; !reflect.DeepEqual(got, want) {
				return nil, fmt.Errorf("entry %d+%d != %d", fromSeq, i, seenIdx)
			}
			seenIdx++
		}
		return []byte("newroot"), nil
	}

	more, err := s.consumeEntries(ctx, 7, f, false)
	if err != nil {
		t.Errorf("consumeEntries: %v", err)
	}
	if !more {
		t.Errorf("more: false, expected true")
	}
}

func makeTile(t *testing.T, size uint64) *api.HashTile {
	t.Helper()
	r := &api.HashTile{Nodes: make([][]byte, size)}
	for i := uint64(0); i < size; i++ {
		h := sha256.Sum256(fmt.Appendf(nil, "%d", i))
		r.Nodes[i] = h[:]
	}
	return r
}

func TestTileRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &logResourceStore{objStore: m}

	for _, test := range []struct {
		name     string
		level    uint64
		index    uint64
		logSize  uint64
		tileSize uint64
	}{
		{
			name:     "ok",
			level:    0,
			index:    3 * layout.TileWidth,
			logSize:  3*layout.TileWidth + 20,
			tileSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantTile := makeTile(t, test.tileSize)
			if err := s.setTile(ctx, test.level, test.index, test.logSize, wantTile); err != nil {
				t.Fatalf("setTile: %v", err)
			}

			expPath := layout.TilePath(test.level, test.index, layout.PartialTileSize(test.level, test.index, test.logSize))
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want tile at %v but found none", expPath)
			}

			got, err := s.getTiles(ctx, []storage.TileID{{Level: test.level, Index: test.index}}, test.logSize)
			if err != nil {
				t.Fatalf("getTile: %v", err)
			}
			if !cmp.Equal(got[0], wantTile) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

func makeBundle(t *testing.T, idx uint64, size int) []byte {
	t.Helper()
	r := &bytes.Buffer{}
	if size == 0 {
		size = layout.EntryBundleWidth
	}
	for i := range size {
		e := core.NewEntry(fmt.Appendf(nil, "%d:%d", idx, i))
		if _, err := r.Write(e.MarshalBundleData(uint64(i))); err != nil {
			t.Fatalf("MarshalBundleEntry: %v", err)
		}
	}
	return r.Bytes()
}

func TestBundleRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &logResourceStore{
		objStore:    m,
		entriesPath: layout.EntriesPath,
	}

	for _, test := range []struct {
		name       string
		index      uint64
		p          uint8
		bundleSize int
	}{
		{
			name:       "ok",
			index:      3 * layout.EntryBundleWidth,
			p:          20,
			bundleSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantBundle := makeBundle(t, 0, test.bundleSize)
			if err := s.setEntryBundle(ctx, test.index, test.p, wantBundle); err != nil {
				t.Fatalf("setEntryBundle: %v", err)
			}

			expPath := layout.EntriesPath(test.index, test.p)
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want bundle at %v but found none", expPath)
			}

			got, err := s.getEntryBundle(ctx, test.index, test.p)
			if err != nil {
				t.Fatalf("getEntryBundle: %v", err)
			}
			if !cmp.Equal(got, wantBundle) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

func TestPublishTree(t *testing.T) {
	ctx := context.Background()
	if canSkipMySQLTest(t, ctx) {
		klog.Warningf("MySQL not available, skipping %s", t.Name())
		t.Skip("MySQL not available, skipping test")
	}

	for _, test := range []struct {
		name            string
		publishInterval time.Duration
		attempts        []time.Duration
		wantUpdates     int
	}{
		{
			name:            "works ok",
			publishInterval: 100 * time.Millisecond,
			attempts:        []time.Duration{1 * time.Second},
			wantUpdates:     1,
		}, {
			name:            "too soon, skip update",
			publishInterval: 10 * time.Second,
			attempts:        []time.Duration{100 * time.Millisecond},
			wantUpdates:     0,
		}, {
			name:            "too soon, skip update, but recovers",
			publishInterval: 2 * time.Second,
			attempts:        []time.Duration{100 * time.Millisecond, 2 * time.Second},
			wantUpdates:     1,
		}, {
			name:            "many attempts, eventually one succeeds",
			publishInterval: 1 * time.Second,
			attempts:        []time.Duration{300 * time.Millisecond, 300 * time.Millisecond, 300 * time.Millisecond, 300 * time.Millisecond},
			wantUpdates:     1,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Clean tables in case there's already something in there.
			mustDropTables(t, ctx)

			s, err := newMySQLSequencer(ctx, *mySQLURI, 1000, 0, 0)
			if err != nil {
				t.Fatalf("newMySQLSequencer: %v", err)
			}
			m := newMemObjStore()
			storage := &Appender{
				logStore: &logResourceStore{
					objStore:    m,
					entriesPath: layout.EntriesPath,
				},
				sequencer: s,
				newCP: func(_ context.Context, size uint64, hash []byte) ([]byte, error) {
					return fmt.Appendf(nil, "%d/%x,", size, hash), nil
				},
			}
			// Call init so we've got a zero-sized checkpoint to work with.
			if err := storage.init(ctx); err != nil {
				t.Fatalf("storage.init: %v", err)
			}
			if err := s.publishTree(ctx, test.publishInterval, storage.publishCheckpoint); err != nil {
				t.Fatalf("publishTree: %v", err)
			}
			cpOld := []byte("bananas")
			if err := m.setObject(ctx, layout.CheckpointPath, cpOld, "", ""); err != nil {
				t.Fatalf("setObject(bananas): %v", err)
			}
			updatesSeen := 0
			for _, d := range test.attempts {
				time.Sleep(d)
				if err := s.publishTree(ctx, test.publishInterval, storage.publishCheckpoint); err != nil {
					t.Fatalf("publishTree: %v", err)
				}
				cpNew, err := m.getObject(ctx, layout.CheckpointPath)
				if err != nil {
					t.Fatalf("getObject: %v", err)
				}
				if !bytes.Equal(cpOld, cpNew) {
					updatesSeen++
					cpOld = cpNew
				}
			}
			if updatesSeen != test.wantUpdates {
				t.Fatalf("Saw %d updates, want %d", updatesSeen, test.wantUpdates)
			}
		})
	}
}

func TestStreamEntries(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()

	logSize1 := 12345
	logSize2 := 100045

	var logSize atomic.Uint64
	logSize.Store(uint64(logSize1))

	s := logResourceStore{
		objStore:    m,
		entriesPath: layout.EntriesPath,
		integratedSize: func(ctx context.Context) (uint64, error) {
			return uint64(logSize2), nil
		},
	}

	// Populate entry bundles:
	// first to logSize1 (so we're sure we've got the partial bundle)
	for r, idx := logSize1, uint64(0); r > 0; idx++ {
		sz := min(r, layout.EntryBundleWidth)
		b := makeBundle(t, idx, sz)
		if err := s.setEntryBundle(ctx, idx, uint8(sz), b); err != nil {
			t.Fatalf("setEntryBundle(%d): %v", idx, err)
		}
		r -= sz
	}
	// Then on to logSize2
	for r, idx := logSize2, uint64(0); r > 0; idx++ {
		sz := min(r, layout.EntryBundleWidth)
		b := makeBundle(t, idx, sz)
		if err := s.setEntryBundle(ctx, idx, uint8(sz), b); err != nil {
			t.Fatalf("setEntryBundle(%d): %v", idx, err)
		}
		r -= sz
	}

	// Finally, try to stream all the bundles back.
	// We'll first try to stream up to logSize1, then when we reach it we'll
	// make the tree appear to grow to logSize2 to test resuming.
	seenEntries := uint64(0)
	next, stop := s.StreamEntries(ctx, 0)

	for {
		gotRI, _, gotErr := next()
		if gotErr != nil {
			if errors.Is(gotErr, tessera.ErrNoMoreEntries) {
				break
			}
			t.Fatalf("gotErr after %d: %v", seenEntries, gotErr)
		}
		if e := gotRI.Index*layout.EntryBundleWidth + uint64(gotRI.First); e != seenEntries {
			t.Fatalf("got idx %d, want %d", e, seenEntries)
		}
		seenEntries += uint64(gotRI.N)
		t.Logf("got RI %d / %d", gotRI.Index, seenEntries)

		switch seenEntries {
		case uint64(logSize1):
			// We've fetched all the entries from the original tree size, now we'll make
			// the tree appear to have grown to the final size.
			// The stream should start returning bundles again until we've consumed them all.
			t.Log("Reached logSize, growing tree")
			logSize.Store(uint64(logSize2))
			time.Sleep(time.Second)
		case uint64(logSize2):
			// We've seen all the entries we created, stop the iterator
			stop()
		}
	}
}

type memObjStore struct {
	sync.RWMutex
	mem map[string][]byte
}

func newMemObjStore() *memObjStore {
	return &memObjStore{
		mem: make(map[string][]byte),
	}
}

func (m *memObjStore) getObject(_ context.Context, obj string) ([]byte, error) {
	m.RLock()
	defer m.RUnlock()

	d, ok := m.mem[obj]
	if !ok {
		return nil, fmt.Errorf("obj %q not found: %w", obj, &types.NoSuchKey{})
	}
	return d, nil
}

// TODO(phboneff): add content type tests
func (m *memObjStore) setObject(_ context.Context, obj string, data []byte, _, _ string) error {
	m.Lock()
	defer m.Unlock()
	m.mem[obj] = data
	return nil
}

// TODO(phboneff): add content type tests
func (m *memObjStore) setObjectIfNoneMatch(_ context.Context, obj string, data []byte, _, _ string) error {
	m.Lock()
	defer m.Unlock()

	d, ok := m.mem[obj]
	if ok && !bytes.Equal(d, data) {
		return &smithy.GenericAPIError{Code: "PreconditionFailed"}
	}
	m.mem[obj] = data
	return nil
}
