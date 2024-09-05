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

// posix runs a web server that allows new entries to be POSTed to
// a tlog-tiles log stored on a posix filesystem. It allows to run
// conformance/compliance/performance tests and showing how to use
// the Tessera POSIX storage implmentation.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/mod/sumdb/note"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/posix"
	"k8s.io/klog/v2"
)

var (
	storageDir  = flag.String("storage_dir", "", "Root directory to store log data.")
	initialise  = flag.Bool("initialise", false, "Set when creating a new log to initialise the structure.")
	listen      = flag.String("listen", ":2025", "Address:port to listen on")
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

	// Create the Tessera POSIX storage, using the directory from the --storage_dir flag
	storage, err := posix.New(ctx, *storageDir, *initialise, tessera.WithCheckpointSignerVerifier(s, v), tessera.WithBatching(256, time.Second))
	if err != nil {
		klog.Exitf("Failed to construct storage: %v", err)
	}

	// Define a handler for /add that accepts POST requests and adds the POST body to the log
	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		idx, err := storage.Add(r.Context(), tessera.NewEntry(b))()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if _, err = w.Write([]byte(fmt.Sprintf("%d", idx))); err != nil {
			klog.Errorf("/add: %v", err)
			return
		}
	})
	// Proxy all GET requests to the filesystem as a lightweight file server.
	// This makes it easier to test this implementation from another machine.
	http.Handle("GET /", http.FileServer(http.Dir(*storageDir)))

	// Run the HTTP server with the single handler and block until this is terminated
	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
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