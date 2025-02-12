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
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	f_log "github.com/transparency-dev/formats/log"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

const (
	// DefaultBatchMaxSize is used by storage implementations if no WithBatching option is provided when instantiating it.
	DefaultBatchMaxSize = 256
	// DefaultBatchMaxAge is used by storage implementations if no WithBatching option is provided when instantiating it.
	DefaultBatchMaxAge = 250 * time.Millisecond
	// DefaultCheckpointInterval is used by storage implementations if no WithCheckpointInterval option is provided when instantiating it.
	DefaultCheckpointInterval = 10 * time.Second
)

// ErrPushback is returned by underlying storage implementations when there are too many
// entries with indices assigned but which have not yet been integrated into the tree.
//
// Personalities encountering this error should apply back-pressure to the source of new entries
// in an appropriate manner (e.g. for HTTP services, return a 503 with a Retry-After header).
var ErrPushback = errors.New("too many unintegrated entries")

// Driver is the implementation-specific parts of Tessera. No methods are on here as this is not for public use.
type Driver interface{}

// IndexFuture is the signature of a function which can return an assigned index or error.
//
// Implementations of this func are likely to be "futures", or a promise to return this data at
// some point in the future, and as such will block when called if the data isn't yet available.
type IndexFuture func() (uint64, error)

// Add adds a new entry to be sequenced.
// This method quickly returns an IndexFuture, which will return the index assigned
// to the new leaf. Until this is returned, the leaf is not durably added to the log,
// and terminating the process may lead to this leaf being lost.
// Once the future resolves and returns an index, the leaf is durably sequenced and will
// be preserved even in the process terminates.
//
// Once a leaf is sequenced, it will be integrated into the tree soon (generally single digit
// seconds). Until it is integrated, clients of the log will not be able to verifiably access
// this value. Personalities that require blocking until the leaf is integrated can use the
// IntegrationAwaiter to wrap the call to this method.
type AddFn func(ctx context.Context, entry *Entry) IndexFuture

// WithCheckpointSigner is an option for setting the note signer and verifier to use when creating and parsing checkpoints.
// This option is mandatory for creating logs where the checkpoint is signed locally, e.g. in
// the Appender mode. This does not need to be provided where the storage will be used to mirror
// other logs.
//
// A primary signer must be provided:
// - the primary signer is the "canonical" signing identity which should be used when creating new checkpoints.
//
// Zero or more dditional signers may also be provided.
// This enables cases like:
//   - a rolling key rotation, where checkpoints are signed by both the old and new keys for some period of time,
//   - using different signature schemes for different audiences, etc.
//
// When providing additional signers, their names MUST be identical to the primary signer name, and this name will be used
// as the checkpoint Origin line.
//
// Checkpoints signed by these signer(s) will be standard checkpoints as defined by https://c2sp.org/tlog-checkpoint.
func WithCheckpointSigner(s note.Signer, additionalSigners ...note.Signer) func(*StorageOptions) {
	origin := s.Name()
	for _, signer := range additionalSigners {
		if origin != signer.Name() {
			klog.Exitf("WithCheckpointSigner: additional signer name (%q) does not match primary signer name (%q)", signer.Name(), origin)
		}

	}
	return func(o *StorageOptions) {
		o.NewCP = func(size uint64, hash []byte) ([]byte, error) {
			// If we're signing a zero-sized tree, the tlog-checkpoint spec says (via RFC6962) that
			// the root must be SHA256 of the empty string, so we'll enforce that here:
			if size == 0 {
				emptyRoot := sha256.Sum256([]byte{})
				hash = emptyRoot[:]
			}
			cpRaw := f_log.Checkpoint{
				Origin: origin,
				Size:   size,
				Hash:   hash,
			}.Marshal()

			n, err := note.Sign(&note.Note{Text: string(cpRaw)}, append([]note.Signer{s}, additionalSigners...)...)
			if err != nil {
				return nil, fmt.Errorf("note.Sign: %w", err)
			}
			return n, nil
		}
	}
}

// WithBatching configures the batching behaviour of leaves being sequenced.
// A batch will be allowed to grow in memory until either:
//   - the number of entries in the batch reach maxSize
//   - the first entry in the batch has reached maxAge
//
// At this point the batch will be sent to the sequencer.
//
// Configuring these parameters allows the personality to tune to get the desired
// balance of sequencing latency with cost. In general, larger batches allow for
// lower cost of operation, where more frequent batches reduce the amount of time
// required for entries to be included in the log.
//
// If this option isn't provided, storage implementations with use the DefaultBatchMaxSize and DefaultBatchMaxAge consts above.
func WithBatching(maxSize uint, maxAge time.Duration) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.BatchMaxSize = maxSize
		o.BatchMaxAge = maxAge
	}
}

// WithPushback allows configuration of when the storage should start pushing back on add requests.
//
// maxOutstanding is the number of "in-flight" add requests - i.e. the number of entries with sequence numbers
// assigned, but which are not yet integrated into the log.
func WithPushback(maxOutstanding uint) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.PushbackMaxOutstanding = maxOutstanding
	}
}

// WithCheckpointInterval configures the frequency at which Tessera will attempt to create & publish
// a new checkpoint.
//
// Well behaved clients of the log will only "see" newly sequenced entries once a new checkpoint is published,
// so it's important to set that value such that it works well with your ecosystem.
//
// Regularly publishing new checkpoints:
//   - helps show that the log is "live", even if no entries are being added.
//   - enables clients of the log to reason about how frequently they need to have their
//     view of the log refreshed, which in turn helps reduce work/load across the ecosystem.
//
// Note that this option probably only makes sense for long-lived applications (e.g. HTTP servers).
//
// If this option isn't provided, storage implementations will use the DefaultCheckpointInterval const above.
func WithCheckpointInterval(interval time.Duration) func(*StorageOptions) {
	return func(o *StorageOptions) {
		o.CheckpointInterval = interval
	}
}
