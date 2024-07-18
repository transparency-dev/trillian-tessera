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
	"reflect"
	"testing"
)

func TestEntryMarshalRoundTrip(t *testing.T) {
	e := NewEntry([]byte("this is data"), WithIdentity([]byte("I am who I am")))
	e.internal.LeafHash = []byte("lettuce")

	raw, err := e.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary: %v", err)
	}

	e2 := Entry{}
	if err := (&e2).UnmarshalBinary(raw); err != nil {
		t.Fatalf("UnmarshalBinary: %v", err)
	}

	if !reflect.DeepEqual(e.internal, e2.internal) {
		t.Fatalf("got %+v, want %+v", e2, e)
	}
}
