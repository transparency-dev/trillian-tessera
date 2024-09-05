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

// Package integration_test contains some integration tests which are intended to
// serve as a way of checking that example binary works as intended,
// as well as providing a simple example of how to run and use it.
package integration_test

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/merkle/proof"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/client"
	"golang.org/x/mod/sumdb/note"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

var (
	runIntegrationTest = flag.Bool("run_integration_test", false, "If true, the integration tests in this package will not be skipped")
	logURL             = flag.String("log_url", "http://localhost:2024", "Log storage read root URL, e.g. https://log.server/and/path/")
	writeLogURL        = flag.String("write_log_url", "http://localhost:2024", "Log storage write root URL, e.g. https://log.server/and/path/")
	logPublicKey       = flag.String("log_public_key", "", "The log's public key value for checkpoint note verification")
	testEntrySize      = flag.Int("test_entry_size", 1024, "The number of entries to be tested in the live log integration")

	noteVerifier note.Verifier

	logReadBaseURL *url.URL
	logRead        client.Fetcher

	hc = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        256,
			MaxIdleConnsPerHost: 256,
		},
		Timeout: 60 * time.Second,
	}
)

func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	flag.Parse()

	if !*runIntegrationTest {
		klog.Warning("example binary integration tests are skipped")
		return
	}

	var err error
	noteVerifier, err = note.NewVerifier(*logPublicKey)
	if err != nil {
		klog.Fatalf("Failed to create new verifier: %v", err)
	}

	logReadBaseURL, err = url.Parse(*logURL)
	if err != nil {
		klog.Fatalf("failed to parse logURL: %v", err)
	}

	switch logReadBaseURL.Scheme {
	case "http", "https":
		logRead = httpRead
	case "file":
		logRead = fileRead
	default:
		klog.Fatalf("unsupported url scheme: %s", logReadBaseURL.Scheme)
	}

	os.Exit(m.Run())
}

func TestLiveLogIntegration(t *testing.T) {
	ctx := context.Background()
	checkpoints := make([]log.Checkpoint, *testEntrySize+1)
	var entryIndexMap sync.Map

	// Step 1 - Get checkpoint initial size for increment validation.
	var checkpointInitSize uint64
	checkpoint, _, _, err := client.FetchCheckpoint(ctx, logRead, noteVerifier, noteVerifier.Name())
	if err != nil {
		t.Errorf("client.FetchCheckpoint: %v", err)
	}
	if checkpoint == nil {
		t.Fatal("checkpoint not found")
	}
	checkpointInitSize = checkpoint.Size
	t.Logf("checkpoint initial size: %d", checkpointInitSize)
	checkpoints[0] = *checkpoint

	// Step 2 - Add entries and get new checkpoints. The entry data comes from the int loop ranging from 0 to the test entry size - 1.
	addEntriesURL, err := url.JoinPath(*writeLogURL, "add")
	if err != nil {
		t.Errorf("url.JoinPath: %v", err)
	}
	entryWriter := entryWriter{
		addURL: addEntriesURL,
	}
	errG := errgroup.Group{}
	for i := range *testEntrySize {
		errG.Go(func() error {
			index, err := entryWriter.add(ctx, []byte(fmt.Sprint(i)))
			if err != nil {
				t.Errorf("entryWriter.add: %v", err)
			}
			entryIndexMap.Store(i, index)
			checkpoint, _, _, err := client.FetchCheckpoint(ctx, logRead, noteVerifier, noteVerifier.Name())
			if err != nil {
				t.Errorf("client.FetchCheckpoint: %v", err)
			}
			if checkpoint == nil {
				// This was a t.Fatalf but that terminates the goroutine, stopping the error being returned on the next line
				t.Errorf("checkpoint not found at index: %d, test entry size: %d", index, i)
				return fmt.Errorf("failed to get checkpoint after writing entry %d (assigned sequence %d)", i, index)
			}
			checkpoints[i+1] = *checkpoint
			return err
		})
	}
	if err := errG.Wait(); err != nil {
		t.Fatalf("addEntry: %v", err)
	}

	// Step 3 - Validate checkpoint size increment.
	checkpoint, _, _, err = client.FetchCheckpoint(ctx, logRead, noteVerifier, noteVerifier.Name())
	if err != nil {
		t.Errorf("client.FetchCheckpoint: %v", err)
	}
	if checkpoint == nil {
		t.Fatal("checkpoint not found")
	}
	t.Logf("checkpoint final size: %d", checkpoint.Size)
	gotIncrease := checkpoint.Size - checkpointInitSize
	if gotIncrease != uint64(*testEntrySize) {
		t.Errorf("checkpoint size increase got: %d, want: %d", gotIncrease, *testEntrySize)
	}

	// Step 4 - Loop through the entry data index map to verify leaves and inclusion proofs.
	entryIndexMap.Range(func(k, v any) bool {
		data := k.(int)
		index := v.(uint64)

		// Step 4.1 - Get entry bundles to read back what was written, check leaves are correct.
		entryBundle, err := client.GetEntryBundle(ctx, logRead, index/256, checkpoint.Size)
		if err != nil {
			t.Errorf("client.GetEntryBundle: %v", err)
		}

		got, want := entryBundle.Entries[index%256], []byte(fmt.Sprint(data))
		if !bytes.Equal(got, want) {
			t.Errorf("Entry bundle (index: %d) got %v want %v", index, got, want)
		}

		// Step 4.2 - Test inclusion proofs.
		pb, err := client.NewProofBuilder(ctx, *checkpoint, logRead)
		if err != nil {
			t.Errorf("client.NewProofBuilder: %v", err)
		}
		ip, err := pb.InclusionProof(ctx, index)
		if err != nil {
			t.Errorf("pb.InclusionProof: %v", err)
		}
		leafHash := rfc6962.DefaultHasher.HashLeaf([]byte(fmt.Sprint(data)))
		if err := proof.VerifyInclusion(rfc6962.DefaultHasher, index, checkpoint.Size, leafHash, ip, checkpoint.Hash); err != nil {
			t.Errorf("proof.VerifyInclusion: %v", err)
		}

		return true
	})

	// Step 5 - Test consistency proofs.
	if err := client.CheckConsistency(ctx, logRead, checkpoints); err != nil {
		t.Errorf("log consistency checks failed: %v", err)
	}
}

func httpRead(ctx context.Context, path string) ([]byte, error) {
	u, err := url.JoinPath(*logURL, path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hc.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			klog.Errorf("resp.Body.Close(): %v", err)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		klog.Infof("Not found: %q", u)
		return nil, os.ErrNotExist
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("unexpected http status %q", resp.Status)
	}
	return body, nil
}

func fileRead(_ context.Context, path string) ([]byte, error) {
	return os.ReadFile(logReadBaseURL.JoinPath(path).Path)
}

type entryWriter struct {
	addURL string
}

func (w *entryWriter) add(ctx context.Context, entry []byte) (uint64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, w.addURL, bytes.NewReader(entry))
	if err != nil {
		return 0, err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(resp.Body)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			klog.Warningf("resp.Body.Close(): %v", err)
		}
	}()
	if err != nil {
		return 0, fmt.Errorf("failed to read response from %s: %w", w.addURL, err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("code: %s, path: %s, body: %s", resp.Status, w.addURL, strings.TrimSpace(string(body)))
	}
	index, err := strconv.ParseUint(string(body), 10, 64)
	if err != nil {
		return 0, err
	}

	return index, nil
}
