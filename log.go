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

package tessera

import (
	"fmt"
	"time"

	f_log "github.com/transparency-dev/formats/log"
	"golang.org/x/mod/sumdb/note"
)

// NewCPFunc is the signature of a function which knows how to format and sign checkpoints.
type NewCPFunc func(size uint64, hash []byte) ([]byte, error)

// ParseCPFunc is the signature of a function which knows how to verify and parse checkpoints.
type ParseCPFunc func(raw []byte) (*f_log.Checkpoint, error)

// StorageOptions holds optional settings for all storage implementations.
type StorageOptions struct {
	NewCP   NewCPFunc
	ParseCP ParseCPFunc

	BatchMaxAge  time.Duration
	BatchMaxSize uint
}

// ResolveStorageOptions turns a variadic array of storage options into a StorageOptions instance.
func ResolveStorageOptions(defaults *StorageOptions, opts ...func(*StorageOptions)) *StorageOptions {
	if defaults == nil {
		defaults = &StorageOptions{}
	}
	for _, opt := range opts {
		opt(defaults)
	}
	return defaults
}

// WithCheckpointSignerVerifier is an option for setting the note signer and verifier to use when creating and parsing checkpoints.
//
// Checkpoints signed by this signer and verified by this verifier will be standard checkpoints as defined by https://c2sp.org/tlog-checkpoint.
// The provided signer's name will be used as the Origin line on the checkpoint.
func WithCheckpointSignerVerifier(s note.Signer, v note.Verifier) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.NewCP = func(size uint64, hash []byte) ([]byte, error) {
			cpRaw := f_log.Checkpoint{
				Origin: s.Name(),
				Size:   size,
				Hash:   hash,
			}.Marshal()

			n, err := note.Sign(&note.Note{Text: string(cpRaw)}, s)
			if err != nil {
				return nil, fmt.Errorf("note.Sign: %w", err)
			}
			return n, nil
		}

		o.ParseCP = func(raw []byte) (*f_log.Checkpoint, error) {
			cp, _, _, err := f_log.ParseCheckpoint(raw, v.Name(), v)
			if err != nil {
				return nil, fmt.Errorf("f_log.ParseCheckpoint: %w", err)
			}
			return cp, nil
		}
	}
}

// WithBatching enables batching of write requests.
func WithBatching(maxSize uint, maxAge time.Duration) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.BatchMaxAge = maxAge
		o.BatchMaxSize = maxSize
	}
}
