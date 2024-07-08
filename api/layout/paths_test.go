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
)

func TestEntriesPathForLogIndex(t *testing.T) {
	for _, test := range []struct {
		seq      uint64
		wantPath string
	}{
		{
			seq:      0,
			wantPath: "tile/entries/000",
		}, {
			seq:      255,
			wantPath: "tile/entries/000",
		}, {
			seq:      256,
			wantPath: "tile/entries/001",
		}, {
			seq:      123456789 * 256,
			wantPath: "tile/entries/x123/x456/789",
		},
	} {
		desc := fmt.Sprintf("seq %d", test.seq)
		t.Run(desc, func(t *testing.T) {
			gotPath := EntriesPathForLogIndex(test.seq)
			if gotPath != test.wantPath {
				t.Errorf("got file %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

func TestEntriesPath(t *testing.T) {
	for _, test := range []struct {
		N        uint64
		wantPath string
	}{
		{
			N:        0,
			wantPath: "tile/entries/000",
		}, {
			N:        255,
			wantPath: "tile/entries/255",
		}, {
			N:        256,
			wantPath: "tile/entries/256",
		}, {
			N:        123456789000,
			wantPath: "tile/entries/x123/x456/x789/000",
		},
	} {
		desc := fmt.Sprintf("N %d", test.N)
		t.Run(desc, func(t *testing.T) {
			gotPath := EntriesPath(test.N)
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
		wantPath string
	}{
		{
			level:    0,
			index:    0,
			wantPath: "tile/0/000",
		}, {
			level:    15,
			index:    455667,
			wantPath: "tile/15/x455/667",
		}, {
			level:    3,
			index:    1234567,
			wantPath: "tile/3/x001/x234/567",
		}, {
			level:    15,
			index:    123456789,
			wantPath: "tile/15/x123/x456/789",
		},
	} {
		desc := fmt.Sprintf("level %x index %x", test.level, test.index)
		t.Run(desc, func(t *testing.T) {
			gotPath := TilePath(test.level, test.index)
			if gotPath != test.wantPath {
				t.Errorf("Got path %q want %q", gotPath, test.wantPath)
			}
		})
	}
}
