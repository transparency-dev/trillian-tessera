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

// posix-oneshot is a command line tool for adding entries to a local
// tlog-tiles log stored on a posix filesystem.
// The command takes a list of new entries to add to the log, and exits
// when they are successfully integrated.
// See the README in this package for more detailed usage instructions.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/mod/sumdb/note"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/posix"
	"k8s.io/klog/v2"
)

var (
	storageDir  = flag.String("storage_dir", "", "Root directory to store log data.")
	entries     = flag.String("entries", "", "File path glob of entries to add to the log.")
	privKeyFile = flag.String("private_key", "", "Location of private key file. If unset, uses the contents of the LOG_PRIVATE_KEY environment variable.")
)

const (
	// checkpointInterval is used as the value to pass to the WithCheckpointInterval option below.
	// Since this is a short-lived command-line tool, we set this to a relatively low value so that
	// the tool can publish the new checkpoint and exit relatively quickly after integrating the entries
	// into the tree.
	checkpointInterval = time.Second
)

// entryInfo binds the actual bytes to be added as a leaf with a
// user-recognisable name for the source of those bytes.
// The name is only used below in order to inform the user of the
// sequence numbers assigned to the data from the provided input files.
type entryInfo struct {
	name string
	f    tessera.IndexFuture
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	// Gather the info needed for reading/writing checkpoints
	s := getSignerOrDie()
	// Construct a new Tessera POSIX log storage, anchored at the correct directory, and initialising it if requested.
	// The options provide the checkpoint signer & verifier, and batch options.
	// In this case, we want to create a single batch containing all of the leaves being added in order to
	// add all of these leaves without creating any intermediate checkpoints.
	driver, err := posix.New(
		ctx,
		*storageDir,
	)
	if err != nil {
		klog.Exitf("Failed to construct storage: %v", err)
	}

	// Evaluate the glob provided by the --entries flag to determine the files containing leaves
	filesToAdd := readEntriesOrDie()

	appender, r, err := tessera.NewAppender(ctx, driver,
		tessera.WithCheckpointSigner(s),
		tessera.WithCheckpointInterval(checkpointInterval),
		tessera.WithBatching(uint(len(filesToAdd)), time.Second))
	if err != nil {
		klog.Exit(err)
	}
	addFn := appender.Add

	// Handle the case where no entries are to be added.
	if len(*entries) == 0 {
		klog.Info("No entries provided to integrate; exiting")
		os.Exit(0)
	}

	// We don't want to exit until our entries have been integrated into the tree, so we'll use Tessera's
	// IntegrationAwaiter to help with that.
	await := tessera.NewIntegrationAwaiter(ctx, r.ReadCheckpoint, time.Second)

	// Add each of the leaves in order, and store the futures in a slice
	// that we will check once all leaves are sent to storage.
	indexFutures := make([]entryInfo, 0, len(filesToAdd))
	for _, fp := range filesToAdd {
		b, err := os.ReadFile(fp)
		if err != nil {
			klog.Exitf("Failed to read entry file %q: %q", fp, err)
		}

		f := addFn(ctx, tessera.NewEntry(b))
		indexFutures = append(indexFutures, entryInfo{name: fp, f: f})
	}

	// Check each of the futures to ensure that the leaves are sequenced.
	for _, entry := range indexFutures {
		seq, _, err := await.Await(ctx, entry.f)
		if err != nil {
			klog.Exitf("failed to sequence %q: %q", entry.name, err)
		}
		klog.Infof("%d: %v", seq, entry.name)
	}

	// All futures have been resolved, which means the log is built and we can allow the process to terminate. Goodbye!
}

// Read log private key from file or environment variable
func getSignerOrDie() note.Signer {
	var privKey string
	var err error
	if len(*privKeyFile) > 0 {
		privKey, err = getKeyFile(*privKeyFile)
		if err != nil {
			klog.Exitf("Unable to get private key: %q", err)
		}
	} else {
		privKey = os.Getenv("LOG_PRIVATE_KEY")
		if len(privKey) == 0 {
			klog.Exit("Supply private key file path using --private_key or set LOG_PRIVATE_KEY environment variable")
		}
	}
	s, err := note.NewSigner(privKey)
	if err != nil {
		klog.Exitf("Failed to instantiate signer: %q", err)
	}
	return s
}

func getKeyFile(path string) (string, error) {
	k, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read key file: %w", err)
	}
	return string(k), nil
}

func readEntriesOrDie() []string {
	toAdd, err := filepath.Glob(*entries)
	if err != nil {
		klog.Exitf("Failed to glob entries %q: %q", *entries, err)
	}
	klog.V(1).Infof("toAdd: %v", toAdd)
	if len(toAdd) == 0 {
		klog.Exit("Sequence must be run with at least one valid entry")
	}
	return toAdd
}
