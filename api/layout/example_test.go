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

package layout_test

import (
	"fmt"

	"github.com/transparency-dev/trillian-tessera/api/layout"
)

func ExampleNodeCoordsToTileAddress() {
	var treeLevel, treeIndex uint64 = 8, 123456789
	tileLevel, tileIndex, nodeLevel, nodeIndex := layout.NodeCoordsToTileAddress(treeLevel, treeIndex)
	fmt.Printf("tile level: %d, tile index: %d, node level: %d, node index: %d", tileLevel, tileIndex, nodeLevel, nodeIndex)
	// Output: tile level: 1, tile index: 482253, node level: 0, node index: 21
}

func ExampleTilePath() {
	tilePath := layout.TilePath(0, 1234067, 315921160)
	fmt.Printf("tile path: %s", tilePath)
	// Output: tile path: tile/0/x001/x234/067.p/8
}

func ExampleEntriesPath() {
	entriesPath := layout.EntriesPath(1234067, 315921160)
	fmt.Printf("entries path: %s", entriesPath)
	// Output: entries path: tile/entries/x001/x234/067.p/8
}

func ExampleParseTileLevelIndexWidth() {
	level, index, width, _ := layout.ParseTileLevelIndexWidth("0", "x001/x234/067.p/8")
	fmt.Printf("level: %d, index: %d, width: %d", level, index, width)
	// Output: level: 0, index: 1234067, width: 8
}
