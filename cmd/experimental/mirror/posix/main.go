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

// mirror/posix is a command-line tool for mirroring a tlog-tiles compliant log
// into a POSIX filesystem.
package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/transparency-dev/trillian-tessera/client"
	mirror "github.com/transparency-dev/trillian-tessera/cmd/experimental/mirror/internal"
	"k8s.io/klog/v2"
)

var (
	storageDir = flag.String("storage_dir", "", "Root directory to store log data.")
	sourceURL  = flag.String("source_url", "", "Base URL for the source log.")
	numWorkers = flag.Uint("num_workers", 30, "Number of migration worker goroutines.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	srcURL, err := url.Parse(*sourceURL)
	if err != nil {
		klog.Exitf("Invalid --source_url %q: %v", *sourceURL, err)
	}
	src, err := client.NewHTTPFetcher(srcURL, nil)
	if err != nil {
		klog.Exitf("Failed to create HTTP fetcher: %v", err)
	}

	if err := mirror.Mirror(ctx, src, *numWorkers, storeFunc(*storageDir)); err != nil {
		klog.Exitf("Failed to mirror log: %v", err)
	}

	klog.Exitf("Log mirrored successfully.")
}

func storeFunc(root string) mirror.StoreFn {
	return func(ctx context.Context, p string, d []byte) (err error) {
		fp := filepath.Join(root, p)
		if err := os.MkdirAll(filepath.Dir(fp), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(fp, d, 0o644); err != nil {
			if errors.Is(err, os.ErrExist) {
				if e, err := os.ReadFile(fp); err != nil {
					return fmt.Errorf("%q already exists, but couldn't read it: %v", fp, err)
				} else if !bytes.Equal(d, e) {
					return fmt.Errorf("%q already exists, but contains different data", fp)
				}
				return nil
			}
			return err
		}
		return nil
	}
}
