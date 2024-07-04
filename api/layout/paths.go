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

// EntriesPath builds the local path at which the leaf with the given index lives in.
// Note that this will be a leaf tile containing up to 256 entries and thus multiple
// indices can map to the same output path.
func EntriesPath(seq uint64) string {
	seq = seq / 256
	frag := []string{
		"tile",
		"entries",
		fmt.Sprintf("x%03x", (seq>>16)&0xff),
		fmt.Sprintf("x%03x", (seq>>8)&0xff),
		fmt.Sprintf("%03x", seq&0xff),
	}
	return filepath.Join(frag...)
}

// TilePath builds the directory path and relative filename for the subtree tile with the
// given level and index.
func TilePath(level, index uint64) string {
	seq := index / 256
	frag := []string{
		"tile",
		fmt.Sprintf("%d", level),
		fmt.Sprintf("x%03x", (seq>>16)&0xff),
		fmt.Sprintf("x%03x", (seq>>8)&0xff),
		fmt.Sprintf("%03x", seq&0xff),
	}
	return filepath.Join(frag...)
}
