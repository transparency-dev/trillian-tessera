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
)

const (
	// CheckpointPath is the location of the file containing the log checkpoint.
	CheckpointPath = "checkpoint"
)

// EntriesPathForLogIndex builds the local path at which the leaf with the given index lives in.
// Note that this will be an entry bundle containing up to 256 entries and thus multiple
// indices can map to the same output path.
//
// TODO(mhutchinson): revisit to consider how partial tile suffixes should be added.
func EntriesPathForLogIndex(seq uint64) string {
	return EntriesPath(seq / 256)
}

// EntriesPath returns the local path for the Nth entry bundle.
func EntriesPath(N uint64) string {
	return fmt.Sprintf("tile/entries%s", fmtN(N))
}

// TilePath builds the path to the subtree tile with the given level and index in tile space.
func TilePath(tileLevel, tileIndex uint64) string {
	return fmt.Sprintf("tile/%d%s", tileLevel, fmtN(tileIndex))
}

// fmtN returns the "N" part of a Tiles-spec path.
//
// N is grouped into chunks of 3 decimal digits, starting with the most significant digit, and
// padding with zeroes as necessary.
// Digit groups are prefixed with "x", except for the least-significant group which has no prefix,
// and separated with slashes.
//
// See https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md#:~:text=index%201234067%20will%20be%20encoded%20as%20x001/x234/067
func fmtN(N uint64) string {
	n := fmt.Sprintf("/%03d", N%1000)
	N /= 1000
	for N > 0 {
		n = fmt.Sprintf("/x%03d%s", N%1000, n)
		N /= 1000
	}
	return n
}
