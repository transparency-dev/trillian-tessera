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

// example-posix runs a web server that allows new entries to be POSTed to
// a tlog-tiles log stored on a posix filesystem.
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
	listen      = flag.String("listen", ":2025", "Address:port to listen on")
	pubKeyFile  = flag.String("public_key", "", "Location of public key file. If unset, uses the contents of the LOG_PUBLIC_KEY environment variable.")
	privKeyFile = flag.String("private_key", "", "Location of private key file. If unset, uses the contents of the LOG_PRIVATE_KEY environment variable.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()

	// Read log public key from file or environment variable
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
	// Read log private key from file or environment variable
	var privKey string
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
	origin := s.Name()

	writeCP := func(size uint64, root []byte) error {
		cp := &fmtlog.Checkpoint{
			Origin: origin,
			Size:   size,
			Hash:   root,
		}
		n, err := note.Sign(&note.Note{Text: string(cp.Marshal())}, s)
		if err != nil {
			return err
		}
		return posix.WriteCheckpoint(*storageDir, n)
	}
	if *initialise {
		if files, err := os.ReadDir(*storageDir); err == nil && len(files) > 0 {
			klog.Exitf("Cannot initialize a log at a non-empty directory")
		}
		if err := os.MkdirAll(*storageDir, dirPerm); err != nil {
			klog.Exitf("Failed to create log directory %q: %v", *storageDir, err)
		}
		if err := writeCP(0, rfc6962.DefaultHasher.EmptyRoot()); err != nil {
			klog.Exitf("Failed to write empty checkpoint")
		}
		klog.Infof("Initialized the log at %q", *storageDir)
	}

	// Check signatures
	v, err := note.NewVerifier(pubKey)
	if err != nil {
		klog.Exitf("Failed to instantiate Verifier: %q", err)
	}

	readCP := func() (uint64, []byte, error) {
		cpRaw, err := posix.ReadCheckpoint(*storageDir)
		if err != nil {
			klog.Exitf("Failed to read log checkpoint: %q", err)
		}
		cp, _, _, err := fmtlog.ParseCheckpoint(cpRaw, origin, v)
		if err != nil {
			return 0, []byte{}, fmt.Errorf("Failed to parse Checkpoint: %q", err)
		}
		return cp.Size, cp.Hash, nil
	}
	storage := posix.New(ctx, *storageDir, readCP, tessera.WithCheckpointSignerVerifier(s, v), tessera.WithBatching(256, time.Second))

	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				klog.Warningf("/add: %v", err)
			}
		}()
		idx, err := storage.Add(r.Context(), tessera.NewEntry(b))
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

	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

func getKeyFile(path string) (string, error) {
	k, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read key file: %w", err)
	}
	return string(k), nil
}
