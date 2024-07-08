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

// Package layout contains routines for specifying the path layout of Tessera logs,
// which is really to say that it provides functions to calculate paths used by the
// tlog-tiles API: https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md
package layout

import (
	"fmt"
	"path/filepath"
)

const (
	// CheckpointPath is the location of the file containing the log checkpoint.
	CheckpointPath = "checkpoint"
)

// EntriesPathForLogIndex() builds the local path at which the leaf with the given index lives in.
// Note that this will be an entry bundle containing up to 256 entries and thus multiple
// indices can map to the same output path.
//
// TODO(mhutchinson): revisit to consider how partial tile suffixes should be added.
func EntriesPathForLogIndex(seq uint64) string {
	seq = seq / 256
	return EntriesPath(seq)
}

func EntriesPath(bundleIndex uint64) string {
	frag := []string{
		"tile",
		"entries",
		fmt.Sprintf("x%03x", (bundleIndex>>16)&0xff),
		fmt.Sprintf("x%03x", (bundleIndex>>8)&0xff),
		fmt.Sprintf("%03x", bundleIndex&0xff),
	}
	return filepath.Join(frag...)
}

// TilePath builds the path to the subtree tile with the given level and index in tile space.
//
// Note that NodeCoordsToTileAddress can be used to convert from node- to tile-space.
func TilePath(tileLevel, tileIndex uint64) string {
	frag := []string{
		"tile",
		fmt.Sprintf("%d", tileLevel),
		fmt.Sprintf("x%03x", (tileIndex>>16)&0xff),
		fmt.Sprintf("x%03x", (tileIndex>>8)&0xff),
		fmt.Sprintf("%03x", tileIndex&0xff),
	}
	return filepath.Join(frag...)
}
