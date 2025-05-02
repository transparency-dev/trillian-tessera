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
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/transparency-dev/trillian-tessera/api/layout"
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

	m := &mirror.Mirror{
		NumWorkers: *numWorkers,
		Src:        src,
		Store:      &posixStore{root: *storageDir},
	}

	// Print out stats.
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				total, done := m.Progress()
				p := float64(done*100) / float64(total)
				log.Printf("Progress %d of %d resources (%0.2f%%)", done, total, p)
			}
		}
	}()

	if err := m.Run(ctx); err != nil {
		klog.Exitf("Failed to mirror log: %v", err)
	}

	klog.Info("Log mirrored successfully.")
}

type posixStore struct {
	root string
}

func (s *posixStore) WriteCheckpoint(ctx context.Context, d []byte) error {
	return s.store(ctx, layout.CheckpointPath, d)
}

func (s *posixStore) WriteTile(ctx context.Context, l, i uint64, p uint8, d []byte) error {
	return s.store(ctx, layout.TilePath(l, i, p), d)
}

func (s *posixStore) WriteEntryBundle(ctx context.Context, i uint64, p uint8, d []byte) error {
	return s.store(ctx, layout.EntriesPath(i, p), d)
}

func (s *posixStore) store(ctx context.Context, p string, d []byte) (err error) {
	fp := filepath.Join(s.root, p)
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
