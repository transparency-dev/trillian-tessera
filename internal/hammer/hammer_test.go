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

package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestLeafGenerator(t *testing.T) {
	// Always generate new values
	gN := newLeafGenerator(0, 100, 0)
	vs := make(map[string]bool)
	for range 256 {
		v := string(gN())
		vs[v] = true
	}

	// Always generate duplicate
	gD := newLeafGenerator(256, 100, 1.0)
	for range 256 {
		if !vs[string(gD())] {
			t.Error("Expected duplicate")
		}
	}
}

func TestHammerAnalyser_Stats(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var treeSize treeSizeState
	ha := newHammerAnalyser(treeSize.getSize)

	go ha.updateStatsLoop(ctx)

	time.Sleep(100 * time.Millisecond)

	baseTime := time.Now().Add(-1 * time.Minute)
	for i := 0; i < 10; i++ {
		ha.seqLeafChan <- leafTime{
			idx:        uint64(i),
			queuedAt:   baseTime,
			assignedAt: baseTime.Add(time.Duration(i) * time.Second),
		}
	}
	treeSize.setSize(10)
	time.Sleep(500 * time.Millisecond)

	avg := ha.queueTime.Avg()
	if want := float64(4500); avg != want {
		t.Errorf("integration time avg: got != want (%f != %f)", avg, want)
	}
}

type treeSizeState struct {
	size uint64
	mux  sync.RWMutex
}

func (s *treeSizeState) getSize() uint64 {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.size
}

func (s *treeSizeState) setSize(size uint64) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.size = size
}