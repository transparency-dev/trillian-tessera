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

// gcp-migrate is a command-line tool for migrating data from a tlog-tiles
// compliant log, into a Tessera log instance.
package main

import (
	"context"
	"encoding/base64"
	"flag"
	"net/url"
	"strconv"
	"strings"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/client"
	"github.com/transparency-dev/trillian-tessera/cmd/experimental/migrate/internal"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"k8s.io/klog/v2"
)

var (
	bucket  = flag.String("bucket", "", "Bucket to use for storing log")
	spanner = flag.String("spanner", "", "Spanner resource URI ('projects/.../...')")

	sourceURL  = flag.String("source_url", "", "Base URL for the source log.")
	numWorkers = flag.Int("num_workers", 30, "Number of migration worker goroutines.")
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
	sourceCP, err := src.ReadCheckpoint(ctx)
	if err != nil {
		klog.Exitf("fetch initial source checkpoint: %v", err)
	}
	bits := strings.Split(string(sourceCP), "\n")
	sourceSize, err := strconv.ParseUint(bits[1], 10, 64)
	if err != nil {
		klog.Exitf("invalid CP size %q: %v", bits[1], err)
	}
	sourceRoot, err := base64.StdEncoding.DecodeString(bits[2])
	if err != nil {
		klog.Exitf("invalid checkpoint roothash %q: %v", bits[2], err)
	}

	// Create our Tessera storage backend:
	st := storageFromFlags(ctx)

	if err := internal.Migrate(context.Background(), *numWorkers, sourceSize, sourceRoot, src.ReadEntryBundle, st); err != nil {
		klog.Exitf("Migrate failed: %v", err)
	}
}

// storageFromFlags returns a MigrationTarget copnfigured for GCP using the values provided via flags.
func storageFromFlags(ctx context.Context) tessera.MigrationTarget {
	if *bucket == "" {
		klog.Exit("--bucket must be set")
	}
	if *spanner == "" {
		klog.Exit("--spanner must be set")
	}
	cfg := gcp.Config{
		Bucket:  *bucket,
		Spanner: *spanner,
	}
	driver, err := gcp.NewMigrationTarget(ctx, cfg, internal.BundleHasher)
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}
	st, err := tessera.NewMigrationTarget(driver)
	if err != nil {
		klog.Exitf("Failed to create new MigrationTarget lifecycle: %v", err)
	}
	return st
}
