// Copyright 2025 The Tessera authors. All Rights Reserved.
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

package core

import "context"

// AddFn adds a new entry to be sequenced.
// This method quickly returns an IndexFuture, which will return the index assigned
// to the new leaf. Until this index is obtained from the future, the leaf is not durably
// added to the log, and terminating the process may lead to this leaf being lost.
// Once the future resolves and returns an index, the leaf is durably sequenced and will
// be preserved even in the process terminates.
//
// Once a leaf is sequenced, it will be integrated into the tree soon (generally single digit
// seconds). Until it is integrated and published, clients of the log will not be able to
// verifiably access this value. Personalities that require blocking until the leaf is integrated
// can use the PublicationAwaiter to wrap the call to this method.
type AddFn func(ctx context.Context, entry *Entry) IndexFuture

// IndexFuture is the signature of a function which can return an assigned index or error.
//
// Implementations of this func are likely to be "futures", or a promise to return this data at
// some point in the future, and as such will block when called if the data isn't yet available.
type IndexFuture func() (Index, error)

// Index represents a durably assigned index for some entry.
type Index struct {
	// Index is the location in the log to which a particular entry has been assigned.
	Index uint64
	// IsDup is true if Index represents a previously assigned index for an identical entry.
	IsDup bool
}
