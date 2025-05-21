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

package core

import (
	"context"

	"github.com/transparency-dev/tessera/ctonly"
)

// NewCertificateTransparencyAppender returns a function which knows how to add a CT-specific entry type to the log.
//
// This entry point MUST ONLY be used for CT logs participating in the CT ecosystem.
// It should not be used as the basis for any other/new transparency application as this protocol:
// a) embodies some techniques which are not considered to be best practice (it does this to retain backawards-compatibility with RFC6962)
// b) is not compatible with the https://c2sp.org/tlog-tiles API which we _very strongly_ encourage you to use instead.
//
// Users of this MUST NOT call `Add` on the underlying Appender directly.
//
// Returns a future, which resolves to the assigned index in the log, or an error.
func NewCertificateTransparencyAppender(a AddFn) func(context.Context, *ctonly.Entry) IndexFuture {
	return func(ctx context.Context, e *ctonly.Entry) IndexFuture {
		return a(ctx, convertCTEntry(e))
	}
}

// convertCTEntry returns an Entry struct which will do the right thing for CT Static API logs.
//
// This MUST NOT be used for any other purpose.
func convertCTEntry(e *ctonly.Entry) *Entry {
	r := &Entry{}
	r.internal.Identity = e.Identity()
	r.marshalForBundle = func(idx uint64) []byte {
		r.internal.LeafHash = e.MerkleLeafHash(idx)
		r.internal.Data = e.LeafData(idx)
		return r.internal.Data
	}

	return r
}
