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
	crand "crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/transparency-dev/trillian-tessera/client"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

func init() {
	flag.Var(&logURL, "log_url", "Log storage root URL (can be specified multiple times), e.g. https://log.server/and/path/")
}

var (
	logURL multiStringFlag

	logPubKey = flag.String("log_public_key", os.Getenv("TILES_LOG_PUBLIC_KEY"), "Public key for the log. This is defaulted to the environment variable TILES_LOG_PUBLIC_KEY")

	maxReadOpsPerSecond = flag.Int("max_read_ops", 20, "The maximum number of read operations per second")
	numReadersRandom    = flag.Int("num_readers_random", 4, "The number of readers looking for random leaves")
	numReadersFull      = flag.Int("num_readers_full", 4, "The number of readers downloading the whole log")

	maxWriteOpsPerSecond = flag.Int("max_write_ops", 0, "The maximum number of write operations per second")
	numWriters           = flag.Int("num_writers", 0, "The number of independent write tasks to run")

	leafMinSize = flag.Int("leaf_min_size", 0, "Minimum size in bytes of individual leaves")
	dupChance   = flag.Float64("dup_chance", 0.1, "The probability of a generated leaf being a duplicate of a previous value")

	showUI = flag.Bool("show_ui", true, "Set to false to disable the text-based UI")

	hc = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        256,
			MaxIdleConnsPerHost: 256,
			DisableKeepAlives:   false,
		},
		Timeout: 5 * time.Second,
	}
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()

	logSigV, err := note.NewVerifier(*logPubKey)
	if err != nil {
		klog.Exitf("failed to create verifier: %v", err)
	}

	f, w := newLogClientsFromFlags()

	var cpRaw []byte
	cons := client.UnilateralConsensus(f.Fetch)
	tracker, err := client.NewLogStateTracker(ctx, f.Fetch, cpRaw, logSigV, logSigV.Name(), cons)
	if err != nil {
		klog.Exitf("Failed to create LogStateTracker: %v", err)
	}
	// Fetch initial state of log
	_, _, _, err = tracker.Update(ctx)
	if err != nil {
		klog.Exitf("Failed to get initial state of the log: %v", err)
	}

	ha := newHammerAnalyser(func() uint64 { return tracker.LatestConsistent.Size })
	go ha.updateStatsLoop(ctx)
	go ha.errorLoop(ctx)

	gen := newLeafGenerator(tracker.LatestConsistent.Size, *leafMinSize, *dupChance)
	hammer := NewHammer(&tracker, f.Fetch, w.Write, gen, ha.seqLeafChan, ha.errChan)
	hammer.Run(ctx)

	if *showUI {
		c := newController(hammer, ha)
		c.Run(ctx)
	} else {
		<-ctx.Done()
	}
}

func NewHammer(tracker *client.LogStateTracker, f client.Fetcher, w LeafWriter, gen func() []byte, seqLeafChan chan<- leafTime, errChan chan<- error) *Hammer {
	readThrottle := NewThrottle(*maxReadOpsPerSecond)
	writeThrottle := NewThrottle(*maxWriteOpsPerSecond)

	randomReaders := newWorkerPool(func() worker {
		return NewLeafReader(tracker, f, RandomNextLeaf(), readThrottle.tokenChan, errChan)
	})
	fullReaders := newWorkerPool(func() worker {
		return NewLeafReader(tracker, f, MonotonicallyIncreasingNextLeaf(), readThrottle.tokenChan, errChan)
	})
	writers := newWorkerPool(func() worker { return NewLogWriter(w, gen, writeThrottle.tokenChan, errChan, seqLeafChan) })

	return &Hammer{
		randomReaders: randomReaders,
		fullReaders:   fullReaders,
		writers:       writers,
		readThrottle:  readThrottle,
		writeThrottle: writeThrottle,
		tracker:       tracker,
	}
}

// Hammer is responsible for coordinating the operations against the log in the form
// of write and read operations. The work of analysing the results of hammering should
// live outside of this class.
type Hammer struct {
	randomReaders workerPool
	fullReaders   workerPool
	writers       workerPool
	readThrottle  *Throttle
	writeThrottle *Throttle
	tracker       *client.LogStateTracker
}

func (h *Hammer) Run(ctx context.Context) {
	// Kick off readers & writers
	for i := 0; i < *numReadersRandom; i++ {
		h.randomReaders.Grow(ctx)
	}
	for i := 0; i < *numReadersFull; i++ {
		h.fullReaders.Grow(ctx)
	}
	for i := 0; i < *numWriters; i++ {
		h.writers.Grow(ctx)
	}

	go h.readThrottle.Run(ctx)
	go h.writeThrottle.Run(ctx)

	go h.updateCheckpointLoop(ctx)
}

func (h *Hammer) updateCheckpointLoop(ctx context.Context) {
	tick := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			size := h.tracker.LatestConsistent.Size
			_, _, _, err := h.tracker.Update(ctx)
			if err != nil {
				klog.Warning(err)
				inconsistentErr := client.ErrInconsistency{}
				if errors.As(err, &inconsistentErr) {
					klog.Fatalf("Last Good Checkpoint:\n%s\n\nFirst Bad Checkpoint:\n%s\n\n%v", string(inconsistentErr.SmallerRaw), string(inconsistentErr.LargerRaw), inconsistentErr)
				}
			}
			newSize := h.tracker.LatestConsistent.Size
			if newSize > size {
				klog.V(1).Infof("Updated checkpoint from %d to %d", size, newSize)
			}
		}
	}
}

func newHammerAnalyser(treeSizeFn func() uint64) *HammerAnalyser {
	leafSampleChan := make(chan leafTime, 100)
	errChan := make(chan error, 20)
	return &HammerAnalyser{
		treeSizeFn:      treeSizeFn,
		seqLeafChan:     leafSampleChan,
		errChan:         errChan,
		integrationTime: movingaverage.Concurrent(movingaverage.New(30)),
		queueTime:       movingaverage.Concurrent(movingaverage.New(30)),
	}
}

// HammerAnalyser is responsible for measuring and interpreting the result of hammering.
type HammerAnalyser struct {
	treeSizeFn  func() uint64
	seqLeafChan chan leafTime
	errChan     chan error

	queueTime       *movingaverage.ConcurrentMovingAverage
	integrationTime *movingaverage.ConcurrentMovingAverage
}

func (a *HammerAnalyser) updateStatsLoop(ctx context.Context) {
	tick := time.NewTicker(100 * time.Millisecond)
	size := a.treeSizeFn()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
		}
		newSize := a.treeSizeFn()
		if newSize <= size {
			continue
		}
		now := time.Now()
		totalLatency := time.Duration(0)
		queueLatency := time.Duration(0)
		numLeaves := 0
		var sample *leafTime
	ReadLoop:
		for {
			if sample == nil {
				select {
				case l, ok := <-a.seqLeafChan:
					if !ok {
						break ReadLoop
					}
					sample = &l
				default:
					break ReadLoop
				}
			}
			// Stop considering leaf times once we've caught up with that cross
			// either the current checkpoint or "now":
			// - leaves with indices beyond the tree size we're considering are not integrated yet, so we can't calculate their TTI
			// - leaves which were queued before "now", but not assigned by "now" should also be ignored as they don't fall into this epoch (and would contribute a -ve latency if they were included).
			if sample.idx >= newSize || sample.assignedAt.After(now) {
				break
			}
			queueLatency += sample.assignedAt.Sub(sample.queuedAt)
			// totalLatency is skewed towards being higher than perhaps it may technically be by:
			// - the tick interval of this goroutine,
			// - the tick interval of the goroutine which updates the LogStateTracker,
			// - any latency in writes to the log becoming visible for reads.
			// But it's probably good enough for now.
			totalLatency += now.Sub(sample.queuedAt)

			numLeaves++
			sample = nil
		}
		if numLeaves > 0 {
			a.integrationTime.Add(float64(totalLatency/time.Millisecond) / float64(numLeaves))
			a.queueTime.Add(float64(queueLatency/time.Millisecond) / float64(numLeaves))
		}
	}
}

func (a *HammerAnalyser) errorLoop(ctx context.Context) {
	tick := time.NewTicker(time.Second)
	pbCount := 0
	for {
		select {
		case <-ctx.Done(): //context cancelled
			return
		case <-tick.C:
			if pbCount > 0 {
				klog.Warningf("%d requests received pushback from log", pbCount)
				pbCount = 0
			}
		case err := <-a.errChan:
			if errors.Is(err, ErrRetry) {
				pbCount++
				continue
			}
			klog.Warning(err)
		}
	}
}

// newLeafGenerator returns a function that generates values to append to a log.
// The leaves are constructed to be at least minLeafSize bytes long.
// dupChance provides the probability that a new leaf will be a duplicate of a previous entry.
// Leaves will be unique if dupChance is 0, and if set to 1 then all values will be duplicates.
// startSize should be set to the initial size of the log so that repeated runs of the
// hammer can start seeding leaves to avoid duplicates with previous runs.
func newLeafGenerator(startSize uint64, minLeafSize int, dupChance float64) func() []byte {
	genLeaf := func(n uint64) []byte {
		// Make a slice with half the number of requested bytes since we'll
		// hex-encode them below which gets us back up to the full amount.
		filler := make([]byte, minLeafSize/2)
		_, _ = crand.Read(filler)
		return []byte(fmt.Sprintf("%x %d", filler, n))
	}

	nextLeaf := genLeaf(startSize)
	return func() []byte {
		if rand.Float64() <= dupChance {
			// This one will actually be unique, but the next iteration will
			// duplicate it. In future, this duplication could be randomly
			// selected to include really old leaves too, to test long-term
			// deduplication in the log (if it supports  that).
			return nextLeaf
		}

		startSize++
		r := nextLeaf
		nextLeaf = genLeaf(startSize)
		return r
	}
}

func NewThrottle(opsPerSecond int) *Throttle {
	return &Throttle{
		opsPerSecond: opsPerSecond,
		tokenChan:    make(chan bool, opsPerSecond),
	}
}

type Throttle struct {
	opsPerSecond int
	tokenChan    chan bool

	oversupply int
}

func (t *Throttle) Increase() {
	tokenCount := t.opsPerSecond
	delta := float64(tokenCount) * 0.1
	if delta < 1 {
		delta = 1
	}
	t.opsPerSecond = tokenCount + int(delta)
}

func (t *Throttle) Decrease() {
	tokenCount := t.opsPerSecond
	if tokenCount <= 1 {
		return
	}
	delta := float64(tokenCount) * 0.1
	if delta < 1 {
		delta = 1
	}
	t.opsPerSecond = tokenCount - int(delta)
}

func (t *Throttle) Run(ctx context.Context) {
	interval := time.Second
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done(): //context cancelled
			return
		case <-ticker.C:
			tokenCount := t.opsPerSecond
			timeout := time.After(interval)
		Loop:
			for i := 0; i < t.opsPerSecond; i++ {
				select {
				case t.tokenChan <- true:
					tokenCount--
				case <-timeout:
					break Loop
				}
			}
			t.oversupply = tokenCount
		}
	}
}

func (t *Throttle) String() string {
	return fmt.Sprintf("Current max: %d/s. Oversupply in last second: %d", t.opsPerSecond, t.oversupply)
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
