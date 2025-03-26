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
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/spanner/spannertest"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"k8s.io/klog/v2"
)

type testLookup struct {
	entryHash    []byte
	wantIndex    uint64
	wantNotFound bool
}

func TestAntispamStorage(t *testing.T) {
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
					wantIndex: 0,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			as, err := NewAntispam(t.Context(), "projects/p/instances/i/databases/d", test.opts)
			if err != nil {
				t.Fatalf("NewAntispam: %v", err)
			}

			fr := newFakeLogReader(test.logEntries)

			f := as.Follower(testBundleHasher)
			go f.Follow(t.Context(), fr)

			for {
				time.Sleep(time.Second)
				pos, err := f.Position(t.Context())
				if err != nil {
					t.Logf("Position: %v", err)
					continue
				}
				klog.Infof("Wait for follower (%d) to catch up with tree (%d)", pos, fr.size)
				if pos >= fr.size {
					break
				}
			}

			for _, e := range test.lookupEntries {
				gotIndex, err := as.index(t.Context(), e.entryHash)
				if err != nil {
					t.Errorf("error looking up hash %x: %v", e.entryHash, err)
				}
				if gotIndex == nil {
					t.Errorf("no index for hash %x", e.entryHash)
					continue
				}
				if *gotIndex != e.wantIndex {
					t.Errorf("got index %d, want %d from looking up hash %x", gotIndex, e.wantIndex, e.entryHash)
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
	return bun.Entries, err
}

type fakeLogReader struct {
	bundles [][]byte
	size    uint64
}

func newFakeLogReader(data [][]byte) *fakeLogReader {
	r := &fakeLogReader{}
	c := [][]byte{}
	for _, d := range data {
		c = append(c, d)
		if len(c) == layout.EntryBundleWidth {
			b := []byte{}
			for i := range c {
				e := tessera.NewEntry(c[i])
				b = append(b, e.MarshalBundleData(r.size)...)
				r.size++
			}
			r.bundles = append(r.bundles, b)
			c = [][]byte{}
		}
	}
	if len(c) > 0 {
		b := []byte{}
		for i := range c {
			e := tessera.NewEntry(c[i])
			b = append(b, e.MarshalBundleData(r.size)...)
			r.size++
		}
		r.bundles = append(r.bundles, b)
	}
	return r
}

func (f fakeLogReader) ReadCheckpoint(_ context.Context) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (f fakeLogReader) ReadTile(_ context.Context, _, _ uint64, _ uint8) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (f fakeLogReader) ReadEntryBundle(_ context.Context, index uint64, _ uint8) ([]byte, error) {
	if index >= uint64(len(f.bundles)) {
		return nil, fmt.Errorf("no bundle at index %d: %v", index, os.ErrNotExist)
	}
	return f.bundles[index], nil
}

func (f fakeLogReader) IntegratedSize(_ context.Context) (uint64, error) {
	return f.size, nil
}

func (f fakeLogReader) StreamEntries(_ context.Context, fromEntry uint64) (func() (layout.RangeInfo, []byte, error), func()) {
	next := func() (layout.RangeInfo, []byte, error) {
		if fromEntry >= f.size {
			return layout.RangeInfo{}, nil, os.ErrNotExist
		}

		bi, offset := fromEntry/layout.EntryBundleWidth, fromEntry%layout.EntryBundleWidth
		n := layout.EntryBundleWidth - offset
		if fromEntry+n > f.size {
			n = f.size - fromEntry
		}

		ri := layout.RangeInfo{
			Index: bi,
			First: uint(offset),
			N:     uint(n),
		}
		fromEntry += n
		klog.Infof("YIELD: %v, %v", ri, f.bundles[bi])
		return ri, f.bundles[bi], nil
	}
	cancel := func() {}
	return next, cancel
}
