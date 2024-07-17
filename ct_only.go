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
	"io"

	"github.com/transparency-dev/trillian-tessera/ctonly"
)

type Storage interface {
	Add(Entry) (uint64, error)
}

func NewCertificateTransparencySequencedWriter(s Storage) func(e ctonly.Entry) (uint64, error) {
	return func(e ctonly.Entry) (uint64, error) {
		return s.Add(convertCTEntry(e))
	}
}

func convertCTEntry(e ctonly.Entry) Entry {
	r := Entry{}
	r.internal.Identity = e.Identity()
	r.indexFunc = func(idx uint64) {
		r.internal.LeafHash = e.MerkleLeafHash(idx)
		r.internal.Data = e.LeafData(idx)
	}
	r.writeBundleEntry = func(w io.Writer) error {
		if n, err := w.Write(r.internal.Data); err != nil {
			return err
		} else if l := len(r.internal.Data); n != l {
			return fmt.Errorf("short write % bytes, want %d", n, l)
		}
		return nil
	}

	return r
}
