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

package gcp

import (
	"context"
	"crypto/sha256"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spannertest"
	"github.com/transparency-dev/tessera"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/core"
	"github.com/transparency-dev/tessera/testonly"
	"k8s.io/klog/v2"
)

type testLookup struct {
	entryHash    []byte
	wantNotFound bool
}

func TestAntispamStorage(t *testing.T) {
	ctx := context.Background()
	closeDB := newSpannerDB(t)
	defer closeDB()

	for _, test := range []struct {
		name          string
		opts          AntispamOpts
		logEntries    [][]byte
		lookupEntries []testLookup
	}{
		{
			name: "roundtrip",
			logEntries: [][]byte{
				[]byte("one"),
				[]byte("two"),
				[]byte("three"),
			},
			lookupEntries: []testLookup{
				{
					entryHash: testIDHash([]byte("one")),
				}, {
					entryHash: testIDHash([]byte("two")),
				}, {
					entryHash: testIDHash([]byte("three")),
				}, {
					entryHash:    testIDHash([]byte("nowhere to be found")),
					wantNotFound: true,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			as, err := NewAntispam(ctx, "projects/p/instances/i/databases/d", test.opts)
			if err != nil {
				t.Fatalf("NewAntispam: %v", err)
			}

			fl, shutdown := testonly.NewTestLog(t, tessera.NewAppendOptions().WithCheckpointInterval(time.Second))
			defer func() {
				if err := shutdown(t.Context()); err != nil {
					t.Logf("shutdown: %v", err)
				}
			}()

			f := as.Follower(testBundleHasher)
			// Hack in a workaround for spannertest not supporting BatchWrites
			f.(*follower).updateIndex = updateIndexTx

			entryIndex := make(map[string]uint64)
			for i, e := range test.logEntries {
				idx, err := fl.Appender.Add(t.Context(), core.NewEntry(e))()
				if err != nil {
					t.Fatalf("Add(%d): %v", i, err)
				}
				entryIndex[string(testIDHash(e))] = idx.Index
			}

			go f.Follow(ctx, fl.LogReader)

			for {
				time.Sleep(time.Second)
				pos, err := f.EntriesProcessed(ctx)
				if err != nil {
					t.Logf("EntriesProcessed: %v", err)
					continue
				}
				sz, err := fl.LogReader.IntegratedSize(t.Context())
				if err != nil {
					t.Logf("IntegratedSize: %v", err)
					continue
				}
				klog.Infof("Wait for follower (%d) to catch up with tree (%d)", pos, sz)
				if pos >= sz {
					break
				}
			}

			for _, e := range test.lookupEntries {
				gotIndex, err := as.index(ctx, e.entryHash)
				if err != nil {
					t.Errorf("error looking up hash %x: %v", e.entryHash, err)
				}
				wantIndex := entryIndex[string(e.entryHash)]
				if gotIndex == nil {
					if !e.wantNotFound {
						t.Errorf("no index for hash %x, but expected index %d", e.entryHash, wantIndex)
					}
					continue
				}
				if *gotIndex != wantIndex {
					t.Errorf("got index %d, want %d from looking up hash %x", gotIndex, wantIndex, e.entryHash)
				}
			}
		})
	}
}

func newSpannerDB(t *testing.T) func() {
	t.Helper()
	srv, err := spannertest.NewServer("localhost:0")
	if err != nil {
		t.Fatalf("Failed to set up test spanner: %v", err)
	}
	if err := os.Setenv("SPANNER_EMULATOR_HOST", srv.Addr); err != nil {
		t.Fatalf("Setenv: %v", err)
	}
	return srv.Close
}

func testIDHash(d []byte) []byte {
	r := sha256.Sum256(d)
	return r[:]
}

func testBundleHasher(b []byte) ([][]byte, error) {
	bun := &api.EntryBundle{}
	err := bun.UnmarshalText(b)
	if err != nil {
		return nil, err
	}
	r := make([][]byte, len(bun.Entries))
	for i, e := range bun.Entries {
		r[i] = testIDHash(e)
	}
	return r, err
}

// updateIndexTx is a workaround for spannertest not supporting BatchWrites.
// We use this func as a replacement for follower's updateIndex hook, and simply commit the index
// updates inline with the larger transaction.
func updateIndexTx(_ context.Context, txn *spanner.ReadWriteTransaction, ms []*spanner.Mutation) error {
	return txn.BufferWrite(ms)
}
