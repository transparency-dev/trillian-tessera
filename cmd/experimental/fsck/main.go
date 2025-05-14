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

// fsck is a command-line tool for checking the integrity of a tlog-tiles based log.
package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"

	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/client"
	"github.com/transparency-dev/trillian-tessera/internal/fsck"
	"k8s.io/klog/v2"
)

var (
	storageURL = flag.String("storage_url", "", "Base tlog-tiles URL")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()
	logURL, err := url.Parse(*storageURL)
	if err != nil {
		klog.Exitf("Invalid --storage_url %q: %v", *storageURL, err)
	}
	src, err := client.NewHTTPFetcher(logURL, nil)
	if err != nil {
		klog.Exitf("Failed to create HTTP fetcher: %v", err)
	}
	if err := fsck.Check(ctx, src, defaultMerkleLeafHasher); err != nil {
		klog.Exitf("fsck failed: %v", err)
	}
}

// defaultMerkleLeafHasher parses a C2SP tlog-tile bundle and returns the Merkle leaf hashes of each entry it contains.
func defaultMerkleLeafHasher(bundle []byte) ([][]byte, error) {
	eb := &api.EntryBundle{}
	if err := eb.UnmarshalText(bundle); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}
	r := make([][]byte, 0, len(eb.Entries))
	for _, e := range eb.Entries {
		h := rfc6962.DefaultHasher.HashLeaf(e)
		r = append(r, h[:])
	}
	return r, nil
}
