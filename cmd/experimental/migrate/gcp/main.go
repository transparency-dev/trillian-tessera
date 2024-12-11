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

// posix-migrate is a command-line tool for migrating data from a tlog-tiles
// compliant log, into a Tessera log instance.
package main

import (
	"context"
	"flag"
	"net/url"

	"github.com/transparency-dev/trillian-tessera/client"
	"github.com/transparency-dev/trillian-tessera/cmd/experimental/migrate/internal/migrate"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"k8s.io/klog/v2"
)

var (
	bucket    = flag.String("bucket", "", "Bucket to use for storing log")
	spanner   = flag.String("spanner", "", "Spanner resource URI ('projects/.../...')")
	sourceURL = flag.String("source_url", "", "Base URL for the source log.")
	stateDB   = flag.String("state_database", "migrate.sttate", "File to use for the temporary file used to track migration state.")
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

	// Create our Tessera storage backend:
	gcpCfg := storageConfigFromFlags()
	st, err := gcp.NewMigrationTarget(ctx, gcpCfg)
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}

	if err := migrate.Migrate(context.Background(), *stateDB, src.ReadCheckpoint, src.ReadTile, src.ReadEntryBundle, st); err != nil {
		klog.Exitf("Migrate failed: %v", err)
	}
}

// storageConfigFromFlags returns a gcp.Config struct populated with values
// provided via flags.
func storageConfigFromFlags() gcp.Config {
	if *bucket == "" {
		klog.Exit("--bucket must be set")
	}
	if *spanner == "" {
		klog.Exit("--spanner must be set")
	}
	return gcp.Config{
		Bucket:  *bucket,
		Spanner: *spanner,
	}
}
