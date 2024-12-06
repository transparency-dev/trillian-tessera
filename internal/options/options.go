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

package options

import (
	"time"

	f_log "github.com/transparency-dev/formats/log"
)

// NewCPFunc is the signature of a function which knows how to format and sign checkpoints.
type NewCPFunc func(size uint64, hash []byte) ([]byte, error)

// ParseCPFunc is the signature of a function which knows how to verify and parse checkpoints.
type ParseCPFunc func(raw []byte) (*f_log.Checkpoint, error)

// EntriesPathFunc is the signature of a function which knows how to format entry bundle paths.
type EntriesPathFunc func(n uint64, p uint8) string

// StorageOptions holds optional settings for all storage implementations.
type StorageOptions struct {
	NewCP NewCPFunc

	BatchMaxAge  time.Duration
	BatchMaxSize uint

	PushbackMaxOutstanding uint

	EntriesPath EntriesPathFunc

	CheckpointInterval time.Duration
}
