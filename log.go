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

type NewCPFunc func(size uint64, hash []byte) ([]byte, error)

type StorageOptions struct {
	NewCP NewCPFunc

	BatchMaxAge  time.Duration
	BatchMaxSize uint
}

func ResolveStorageOptions(defaults *StorageOptions, opts ...func(*StorageOptions)) *StorageOptions {
	if defaults == nil {
		defaults = &StorageOptions{}
	}
	for _, opt := range opts {
		opt(defaults)
	}
	return defaults
}

func WithCheckpointSigner(s note.Signer) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.NewCP = func(size uint64, hash []byte) ([]byte, error) {
			cpRaw := f_log.Checkpoint{
				Origin: s.Name(),
				Size:   size,
				Hash:   hash,
			}.Marshal()

			n, err := note.Sign(&note.Note{Text: string(cpRaw)}, s)
			if err != nil {
				return nil, fmt.Errorf("Sign: %v", err)
			}
			return n, nil
		}
	}
}

func WithBatching(maxSize uint, maxAge time.Duration) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.BatchMaxAge = maxAge
		o.BatchMaxSize = maxSize
	}
}
