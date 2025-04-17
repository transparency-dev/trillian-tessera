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

package shizzle

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEntryMarshalBundleDelegates(t *testing.T) {
	const wantIdx = uint64(143)
	wantBundle := fmt.Appendf(nil, "Yes %d", wantIdx)

	e := NewEntry([]byte("this is data"))
	e.marshalForBundle = func(gotIdx uint64) []byte {
		if gotIdx != wantIdx {
			t.Fatalf("Got idx %d, want %d", gotIdx, wantIdx)
		}
		return wantBundle
	}

	if got, want := e.MarshalBundleData(wantIdx), wantBundle; !bytes.Equal(got, want) {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
