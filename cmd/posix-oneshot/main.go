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

	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/posix"
	"k8s.io/klog/v2"

	fmtlog "github.com/transparency-dev/formats/log"
)

const (
	dirPerm = 0o755
)

var (
	storageDir  = flag.String("storage_dir", "", "Root directory to store log data.")
	initialise  = flag.Bool("initialise", false, "Set when creating a new log to initialise the structure.")
	entries     = flag.String("entries", "", "File path glob of entries to add to the log.")
	pubKeyFile  = flag.String("public_key", "", "Location of public key file. If unset, uses the contents of the LOG_PUBLIC_KEY environment variable.")
	privKeyFile = flag.String("private_key", "", "Location of private key file. If unset, uses the contents of the LOG_PRIVATE_KEY environment variable.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	// Gather the info needed for reading/writing checkpoints
	v := getVerifierOrDie()
	s := getSignerOrDie()
	origin := s.Name()

	if *initialise {
		// Create the directory structure and write out an empty checkpoint
		if err := os.MkdirAll(*storageDir, dirPerm); err != nil {
			klog.Exitf("Failed to create log directory: %q", err)
		}
		// TODO(mhutchinson): This empty checkpoint initialization should live in Tessera
		emptyCP := &fmtlog.Checkpoint{
			Origin: origin,
			Size:   0,
			Hash:   rfc6962.DefaultHasher.EmptyRoot(),
		}
		n, err := note.Sign(&note.Note{Text: string(emptyCP.Marshal())}, s)
		if err != nil {
			klog.Exitf("Failed to sign empty checkpoint: %s", err)
		}
		if err := posix.WriteCheckpoint(*storageDir, n); err != nil {
			klog.Exitf("Failed to write empty checkpoint: %s", err)
		}
		// TODO(mhutchinson): This should continue if *entries is provided
		os.Exit(0)
	}

	filesToAdd := readEntriesOrDie()

	st := posix.New(ctx, *storageDir, tessera.WithCheckpointSignerVerifier(s, v), tessera.WithBatching(uint(len(filesToAdd)), time.Second))

	// sequence entries

	// entryInfo binds the actual bytes to be added as a leaf with a
	// user-recognisable name for the source of those bytes.
	// The name is only used below in order to inform the user of the
	// sequence numbers assigned to the data from the provided input files.
	type entryInfo struct {
		name string
		f    tessera.IndexFuture
	}
	indexFutures := make([]entryInfo, 0, len(filesToAdd))
	for _, fp := range filesToAdd {
		b, err := os.ReadFile(fp)
		if err != nil {
			klog.Exitf("Failed to read entry file %q: %q", fp, err)
		}

		// ask storage to sequence and we'll store the future for later
		f := st.Add(ctx, tessera.NewEntry(b))
		indexFutures = append(indexFutures, entryInfo{name: fp, f: f})
	}

	for _, entry := range indexFutures {
		seq, err := entry.f()
		if err != nil {
			klog.Exitf("failed to sequence %q: %q", entry.name, err)
		}
		klog.Infof("%d: %v", seq, entry.name)
	}
}

// Read log public key from file or environment variable
func getVerifierOrDie() note.Verifier {
	var pubKey string
	var err error
	if len(*pubKeyFile) > 0 {
		pubKey, err = getKeyFile(*pubKeyFile)
		if err != nil {
			klog.Exitf("Unable to get public key: %q", err)
		}
	} else {
		pubKey = os.Getenv("LOG_PUBLIC_KEY")
		if len(pubKey) == 0 {
			klog.Exit("Supply public key file path using --public_key or set LOG_PUBLIC_KEY environment variable")
		}
	}
	// Check signatures
	v, err := note.NewVerifier(pubKey)
	if err != nil {
		klog.Exitf("Failed to instantiate Verifier: %q", err)
	}

	return v
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
