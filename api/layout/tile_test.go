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

func TestNodeCoordsToTileAddress(t *testing.T) {
	for _, test := range []struct {
		treeLevel     uint64
		treeIndex     uint64
		wantTileLevel uint64
		wantTileIndex uint64
		wantNodeLevel uint
		wantNodeIndex uint64
	}{
		{
			treeLevel:     0,
			treeIndex:     0,
			wantTileLevel: 0,
			wantTileIndex: 0,
			wantNodeLevel: 0,
			wantNodeIndex: 0,
		}, {
			treeLevel:     0,
			treeIndex:     255,
			wantTileLevel: 0,
			wantTileIndex: 0,
			wantNodeLevel: 0,
			wantNodeIndex: 255,
		}, {
			treeLevel:     0,
			treeIndex:     256,
			wantTileLevel: 0,
			wantTileIndex: 1,
			wantNodeLevel: 0,
			wantNodeIndex: 0,
		}, {
			treeLevel:     1,
			treeIndex:     0,
			wantTileLevel: 0,
			wantTileIndex: 0,
			wantNodeLevel: 1,
			wantNodeIndex: 0,
		}, {
			treeLevel:     8,
			treeIndex:     0,
			wantTileLevel: 1,
			wantTileIndex: 0,
			wantNodeLevel: 0,
			wantNodeIndex: 0,
		},
	} {
		t.Run(fmt.Sprintf("%d-%d", test.treeLevel, test.treeIndex), func(t *testing.T) {
			tl, ti, nl, ni := NodeCoordsToTileAddress(test.treeLevel, test.treeIndex)
			if got, want := tl, test.wantTileLevel; got != want {
				t.Errorf("Got TileLevel %d want %d", got, want)
			}
			if got, want := ti, test.wantTileIndex; got != want {
				t.Errorf("Got TileIndex %d want %d", got, want)
			}
			if got, want := nl, test.wantNodeLevel; got != want {
				t.Errorf("Got NodeLevel %d want %d", got, want)
			}
			if got, want := ni, test.wantNodeIndex; got != want {
				t.Errorf("Got NodeIndex %d want %d", got, want)
			}
		})
	}
}
