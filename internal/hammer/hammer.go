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

// hammer is a tool to load test a Tessera log.
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/transparency-dev/trillian-tessera/client"
	"github.com/transparency-dev/trillian-tessera/internal/hammer/loadtest"
	"golang.org/x/mod/sumdb/note"
	"golang.org/x/net/http2"

	"k8s.io/klog/v2"
)

func init() {
	flag.Var(&logURL, "log_url", "Log storage root URL (can be specified multiple times), e.g. https://log.server/and/path/")
	flag.Var(&writeLogURL, "write_log_url", "Root URL for writing to a log (can be specified multiple times), e.g. https://log.server/and/path/ (optional, defaults to log_url)")
}

var (
	logURL      multiStringFlag
	writeLogURL multiStringFlag

	logPubKey = flag.String("log_public_key", os.Getenv("TILES_LOG_PUBLIC_KEY"), "Public key for the log. This is defaulted to the environment variable TILES_LOG_PUBLIC_KEY")

	maxReadOpsPerSecond = flag.Int("max_read_ops", 20, "The maximum number of read operations per second")
	numReadersRandom    = flag.Int("num_readers_random", 4, "The number of readers looking for random leaves")
	numReadersFull      = flag.Int("num_readers_full", 4, "The number of readers downloading the whole log")

	maxWriteOpsPerSecond = flag.Int("max_write_ops", 0, "The maximum number of write operations per second")
	numWriters           = flag.Int("num_writers", 0, "The number of independent write tasks to run")

	leafMinSize = flag.Int("leaf_min_size", 0, "Minimum size in bytes of individual leaves")
	dupChance   = flag.Float64("dup_chance", 0.1, "The probability of a generated leaf being a duplicate of a previous value")

	leafWriteGoal = flag.Int64("leaf_write_goal", 0, "Exit after writing this number of leaves, or 0 to keep going indefinitely")
	maxRunTime    = flag.Duration("max_runtime", 0, "Fail after this amount of time has passed, or 0 to keep going indefinitely")

	showUI = flag.Bool("show_ui", true, "Set to false to disable the text-based UI")

	bearerToken      = flag.String("bearer_token", "", "The bearer token for auth. For GCP this is the result of `gcloud auth print-access-token`")
	bearerTokenWrite = flag.String("bearer_token_write", "", "The bearer token for auth to write. For GCP this is the result of `gcloud auth print-identity-token`. If unset will default to --bearer_token.")

	forceHTTP2 = flag.Bool("force_http2", false, "Use HTTP/2 connections *only*")

	hc = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        256,
			MaxIdleConnsPerHost: 256,
			DisableKeepAlives:   false,
		},
		Timeout: 30 * time.Second,
	}
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	if *forceHTTP2 {
		hc.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{},
		}
	}

	// If bearerTokenWrite is unset, default it to whatever bearerToken has (which may too be unset).
	if *bearerTokenWrite == "" {
		*bearerTokenWrite = *bearerToken
	}

	ctx, cancel := context.WithCancel(context.Background())

	logSigV, err := note.NewVerifier(*logPubKey)
	if err != nil {
		klog.Exitf("failed to create verifier: %v", err)
	}

	f, w, err := loadtest.NewLogClients(logURL, writeLogURL, loadtest.ClientOpts{
		Client:           hc,
		BearerToken:      *bearerToken,
		BearerTokenWrite: *bearerTokenWrite,
	})
	if err != nil {
		klog.Exit(err)
	}

	var cpRaw []byte
	cons := client.UnilateralConsensus(f.ReadCheckpoint)
	tracker, err := client.NewLogStateTracker(ctx, f.ReadTile, cpRaw, logSigV, logSigV.Name(), cons)
	if err != nil {
		klog.Exitf("Failed to create LogStateTracker: %v", err)
	}
	// Fetch initial state of log
	_, _, _, err = tracker.Update(ctx)
	if err != nil {
		klog.Exitf("Failed to get initial state of the log: %v", err)
	}

	ha := loadtest.NewHammerAnalyser(func() uint64 { return tracker.Latest().Size })
	ha.Run(ctx)

	gen := newLeafGenerator(tracker.Latest().Size, *leafMinSize, *dupChance)
	opts := loadtest.HammerOpts{
		MaxReadOpsPerSecond:  *maxReadOpsPerSecond,
		MaxWriteOpsPerSecond: *maxWriteOpsPerSecond,
		NumReadersRandom:     *numReadersRandom,
		NumReadersFull:       *numReadersFull,
		NumWriters:           *numWriters,
	}
	hammer := loadtest.NewHammer(tracker, f.ReadEntryBundle, w, gen, ha.SeqLeafChan, ha.ErrChan, opts)

	exitCode := 0
	if *leafWriteGoal > 0 {
		go func() {
			startTime := time.Now()
			goal := tracker.Latest().Size + uint64(*leafWriteGoal)
			klog.Infof("Will exit once tree size is at least %d", goal)
			tick := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-ctx.Done():
					return
				case <-tick.C:
					if tracker.Latest().Size >= goal {
						elapsed := time.Since(startTime)
						klog.Infof("Reached tree size goal of %d after %s; exiting", goal, elapsed)
						cancel()
						return
					}
				}
			}
		}()
	}
	if *maxRunTime > 0 {
		go func() {
			klog.Infof("Will fail after %s", *maxRunTime)
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(*maxRunTime):
					klog.Infof("Max runtime reached; exiting")
					exitCode = 1
					cancel()
					return
				}
			}
		}()
	}
	hammer.Run(ctx)

	if *showUI {
		c := loadtest.NewController(hammer, ha)
		c.Run(ctx)
	} else {
		<-ctx.Done()
	}
	os.Exit(exitCode)
}

// newLeafGenerator returns a function that generates values to append to a log.
// The leaves are constructed to be at least minLeafSize bytes long.
// The generator can be used by concurrent threads.
//
// dupChance provides the probability that a new leaf will be a duplicate of a previous entry.
// Leaves will be unique if dupChance is 0, and if set to 1 then all values will be duplicates.
// startSize should be set to the initial size of the log so that repeated runs of the
// hammer can start seeding leaves to avoid duplicates with previous runs.
func newLeafGenerator(startSize uint64, minLeafSize int, dupChance float64) func() []byte {
	// genLeaf MUST be determinstic given n
	genLeaf := func(n uint64) []byte {
		// Make a slice with half the number of requested bytes since we'll
		// hex-encode them below which gets us back up to the full amount.
		filler := make([]byte, minLeafSize/2)
		source := rand.New(rand.NewPCG(0, n))
		for i := range filler {
			// This throws away a lot of the generated data. An exercise to a future
			// coder is to fill in multiple bytes at a time.
			filler[i] = byte(source.Int())
		}
		return fmt.Appendf(nil, "%x %d", filler, n)
	}

	sizeLocked := startSize
	var mu sync.Mutex
	return func() []byte {
		mu.Lock()
		thisSize := sizeLocked

		if thisSize > 0 && rand.Float64() <= dupChance {
			thisSize = rand.Uint64N(thisSize)
		} else {
			sizeLocked++
		}
		mu.Unlock()

		// Do this outside of the protected block so that writers don't block on leaf generation (especially for larger leaves).
		return genLeaf(thisSize)
	}
}

// multiStringFlag allows a flag to be specified multiple times on the command
// line, and stores all of these values.
type multiStringFlag []string

func (ms *multiStringFlag) String() string {
	return strings.Join(*ms, ",")
}

func (ms *multiStringFlag) Set(w string) error {
	*ms = append(*ms, w)
	return nil
}
