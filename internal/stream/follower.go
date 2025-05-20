// Copyright 2025 The Tessera Authors. All Rights Reserved.
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

package stream

import (
	"context"
	"iter"
)

// Follower describes the contract of something which is required to track the contents of the local log.
type Follower interface {
	// Name returns a human readable name for this follower.
	Name() string

	// Follow should be implemented so as to visit entries in the log in order, using the provided
	// LogReader to access the entry bundles which contain them.
	//
	// Implementations should keep track of their progress such that they can pick-up where they left off
	// if e.g. the binary is restarted.
	Follow(context.Context, Streamer)

	// EntriesProcessed reports the progress of the follower, returning the total number of log entries
	// successfully seen/processed.
	EntriesProcessed(context.Context) (uint64, error)
}

type Streamer interface {
	// IntegratedSize returns the current size of the integrated tree.
	//
	// This tree will have in place all the static resources the returned size implies, but
	// there may not yet be a checkpoint for this size signed, witnessed, or published.
	//
	// It's ONLY safe to use this value for processes internal to the operation of the log (e.g.
	// populating antispam data structures); it MUST NOT not be used as a substitute for
	// reading the checkpoint when only data which has been publicly committed to by the
	// log should be used. If in doubt, use ReadCheckpoint instead.
	IntegratedSize(ctx context.Context) (uint64, error)

	// NextIndex returns the first as-yet unassigned index.
	//
	// In a quiescent log, this will be the same as the checkpoint size. In a log with entries actively
	// being added, this number will be higher since it will take sequenced but not-yet-integrated/not-yet-published
	// entries into account.
	NextIndex(ctx context.Context) (uint64, error)

	// StreamEntries returns an iterator over the range of requested entries.
	//
	// The iterator will yield either a Bundle struct or an error. The Bundle contains the raw serialised form
	// of the entry bundle, along with a layout.RangeInfo struct which describes which of the entries in the
	// entry bundle are part of the requested range.
	StreamEntries(ctx context.Context, startEntryIdx, N uint64) iter.Seq2[Bundle, error]
}
