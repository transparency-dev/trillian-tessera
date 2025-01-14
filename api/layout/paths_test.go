// Copyright 2024 Google LLC. All Rights Reserved.
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

package layout

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEntriesPathForLogIndex(t *testing.T) {
	for _, test := range []struct {
		seq      uint64
		logSize  uint64
		wantPath string
	}{
		{
			seq:      0,
			logSize:  256,
			wantPath: "tile/entries/000",
		}, {
			seq:      255,
			logSize:  256,
			wantPath: "tile/entries/000",
		}, {
			seq:      251,
			logSize:  255,
			wantPath: "tile/entries/000.p/255",
		}, {
			seq:      256,
			logSize:  512,
			wantPath: "tile/entries/001",
		}, {
			seq:      256,
			logSize:  257,
			wantPath: "tile/entries/001.p/1",
		}, {
			seq:      123456789 * 256,
			logSize:  123456790 * 256,
			wantPath: "tile/entries/x123/x456/789",
		},
	} {
		desc := fmt.Sprintf("seq %d", test.seq)
		t.Run(desc, func(t *testing.T) {
			gotPath := EntriesPathForLogIndex(test.seq, test.logSize)
			if gotPath != test.wantPath {
				t.Errorf("got file %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

func TestEntriesPath(t *testing.T) {
	for _, test := range []struct {
		N        uint64
		p        uint8
		wantPath string
		wantErr  bool
	}{
		{
			N:        0,
			wantPath: "tile/entries/000",
		},
		{
			N:        0,
			p:        8,
			wantPath: "tile/entries/000.p/8",
		}, {
			N:        255,
			wantPath: "tile/entries/255",
		}, {
			N:        255,
			p:        253,
			wantPath: "tile/entries/255.p/253",
		},
	} {
		desc := fmt.Sprintf("N %d", test.N)
		t.Run(desc, func(t *testing.T) {
			gotPath := EntriesPath(test.N, test.p)
			if gotPath != test.wantPath {
				t.Errorf("got file %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

func TestTilePath(t *testing.T) {
	for _, test := range []struct {
		level    uint64
		index    uint64
		p        uint8
		wantPath string
	}{
		{
			level:    0,
			index:    0,
			wantPath: "tile/0/000",
		}, {
			level:    0,
			index:    0,
			p:        255,
			wantPath: "tile/0/000.p/255",
		}, {
			level:    1,
			index:    0,
			wantPath: "tile/1/000",
		}, {
			level:    15,
			index:    455667,
			p:        0,
			wantPath: "tile/15/x455/667",
		}, {
			level:    15,
			index:    123456789,
			p:        41,
			wantPath: "tile/15/x123/x456/789.p/41",
		},
	} {
		desc := fmt.Sprintf("level %x index %x", test.level, test.index)
		t.Run(desc, func(t *testing.T) {
			gotPath := TilePath(test.level, test.index, test.p)
			if gotPath != test.wantPath {
				t.Errorf("Got path %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

func TestNWithSuffix(t *testing.T) {
	for _, test := range []struct {
		level    uint64
		index    uint64
		p        uint8
		wantPath string
	}{
		{
			level:    0,
			index:    0,
			wantPath: "000",
		}, {
			level:    0,
			index:    0,
			p:        255,
			wantPath: "000.p/255",
		}, {
			level:    15,
			index:    455667,
			wantPath: "x455/667",
		}, {
			level:    15,
			index:    123456789,
			p:        65,
			wantPath: "x123/x456/789.p/65",
		},
	} {
		desc := fmt.Sprintf("level %x index %x", test.level, test.index)
		t.Run(desc, func(t *testing.T) {
			gotPath := NWithSuffix(test.level, test.index, test.p)
			if gotPath != test.wantPath {
				t.Errorf("Got path %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

func TestParseTileLevelIndexPartial(t *testing.T) {
	for _, test := range []struct {
		pathLevel string
		pathIndex string
		wantLevel uint64
		wantIndex uint64
		wantP     uint8
		wantErr   bool
	}{
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067",
			wantLevel: 0,
			wantIndex: 1234067,
			wantP:     0,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p/89",
			wantLevel: 0,
			wantIndex: 1234067,
			wantP:     89,
		},
		{
			pathLevel: "63",
			pathIndex: "x999/x999/x999/x999/x999/999.p/255",
			wantLevel: 63,
			wantIndex: 999999999999999999,
			wantP:     255,
		},
		{
			pathLevel: "0",
			pathIndex: "001",
			wantLevel: 0,
			wantIndex: 1,
			wantP:     0,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001",
			wantErr:   true,
		},
		{
			pathLevel: "1",
			pathIndex: "x001/.p/abc",
			wantErr:   true,
		},
		{
			pathLevel: "64",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "-1",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "abc",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "001/002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/0002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/-002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/002.p/256",
			wantErr:   true,
		},
		{
			pathLevel: "63",
			pathIndex: "x999/x999/x999/x999/x999/x999/999.p/255",
			wantErr:   true,
		},
	} {
		desc := fmt.Sprintf("pathLevel: %q, pathIndex: %q", test.pathLevel, test.pathIndex)
		t.Run(desc, func(t *testing.T) {
			gotLevel, gotIndex, gotWidth, err := ParseTileLevelIndexPartial(test.pathLevel, test.pathIndex)
			if gotLevel != test.wantLevel {
				t.Errorf("got level %d want %d", gotLevel, test.wantLevel)
			}
			if gotIndex != test.wantIndex {
				t.Errorf("got index %d want %d", gotIndex, test.wantIndex)
			}
			if gotWidth != test.wantP {
				t.Errorf("got width %d want %d", gotWidth, test.wantP)
			}
			gotErr := err != nil
			if gotErr != test.wantErr {
				t.Errorf("got err %v want %v", gotErr, test.wantErr)
			}
		})
	}
}

func TestRange(t *testing.T) {
	for _, test := range []struct {
		from, N, treeSize uint64
		desc              string
		want              []RangeInfo
	}{
		{
			desc:     "from beyond extent",
			from:     10,
			N:        1,
			treeSize: 5,
			want:     []RangeInfo{},
		}, {
			desc:     "range end beyond extent",
			from:     3,
			N:        100,
			treeSize: 5,
			want:     []RangeInfo{{Index: 0, First: 3, N: 5 - 3, Partial: 5}},
		}, {
			desc:     "empty range",
			from:     1,
			N:        0,
			treeSize: 2,
			want:     []RangeInfo{},
		}, {
			desc:     "ok: full first bundle",
			from:     0,
			N:        256,
			treeSize: 257,
			want:     []RangeInfo{{N: 256}},
		}, {
			desc:     "ok: entire single (partial) bundle",
			from:     20,
			N:        90,
			treeSize: 111,
			want:     []RangeInfo{{Index: 0, Partial: 111, First: 20, N: 90}},
		}, {
			desc:     "ok: slice from single bundle with initial offset",
			from:     20,
			N:        90,
			treeSize: 1 << 20,
			want:     []RangeInfo{{Index: 0, Partial: 0, First: 20, N: 90}},
		}, {
			desc:     "ok: multiple bundles, first is full, last is truncated",
			from:     0,
			N:        4*256 + 42,
			treeSize: 1 << 20,
			want: []RangeInfo{
				{Index: 0, Partial: 0, First: 0, N: 256},
				{Index: 1, Partial: 0, First: 0, N: 256},
				{Index: 2, Partial: 0, First: 0, N: 256},
				{Index: 3, Partial: 0, First: 0, N: 256},
				{Index: 4, Partial: 0, First: 0, N: 42},
			},
		}, {
			desc:     "ok: multiple bundles, first is offset, last is truncated",
			from:     2,
			N:        4*256 + 4,
			treeSize: 1 << 20,
			want: []RangeInfo{
				{Index: 0, Partial: 0, First: 2, N: 256 - 2},
				{Index: 1, Partial: 0, First: 0, N: 256},
				{Index: 2, Partial: 0, First: 0, N: 256},
				{Index: 3, Partial: 0, First: 0, N: 256},
				{Index: 4, Partial: 0, First: 0, N: 6},
			},
		}, {
			desc:     "ok: offset and trucated from single bundle in middle of tree",
			from:     8*256 + 66,
			N:        4,
			treeSize: 1 << 20,
			want:     []RangeInfo{{Index: 8, Partial: 0, First: 66, N: 4}},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			i := 0
			for gotInfo := range Range(test.from, test.N, test.treeSize) {
				if d := cmp.Diff(test.want[i], gotInfo); d != "" {
					t.Fatalf("got results[%d] with diff:\n%s", i, d)
				}
				i++
			}
		})
	}
}
