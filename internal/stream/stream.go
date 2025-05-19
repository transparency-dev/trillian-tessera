// Copyright 2025 The Tessera Authors. All Rights Reserved.
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

// Package stream provides support for streaming contiguous entries from logs.
package stream

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/transparency-dev/tessera/api/layout"
	"k8s.io/klog/v2"
)

// NoMoreEntries is a sentinel error returned by StreamEntries when no more entries will be returned by calls to the next function.
var ErrNoMoreEntries = errors.New("no more entries")

// GetBundleFn is a function which knows how to fetch a single entry bundle from the specified address.
type GetBundleFn func(ctx context.Context, bundleIdx uint64, partial uint8) ([]byte, error)

// GetTreeSizeFn is a function which knows how to return a tree size.
type GetTreeSizeFn func(ctx context.Context) (uint64, error)

// StreamAdaptor uses the provided function to produce a stream of entry bundles accesible via the returned functions.
//
// Entry bundles are retuned strictly in order via consecutive calls to the returned next func.
// If the adaptor encounters an error while reading an entry bundle, the encountered error will be returned by the corresponding call to next,
// and the stream will be stopped - further calls to next will continue to return errors.
//
// When the caller has finished consuming entry bundles (either because of an error being returned via next, or having consumed all the bundles it needs),
// it MUST call the returned cancel function to release resources.
//
// This adaptor is optimised for the case where calling getBundle has some appreciable latency, and works
// around that by maintaining a read-ahead cache of subsequent bundles which is populated a number of parallel
// requests to getBundle. The request parallelism is set by the value of the numWorkers paramemter, which can be tuned
// to balance throughput against consumption of resources, but such balancing needs to be mindful of the nature of the
// source infrastructure, and how concurrent requests affect performance (e.g. GCS buckets vs. files on a single disk).
func StreamAdaptor(ctx context.Context, numWorkers uint, getSize GetTreeSizeFn, getBundle GetBundleFn, fromEntry uint64) (next func() (ri layout.RangeInfo, bundle []byte, err error), cancel func()) {
	ctx, span := tracer.Start(ctx, "tessera.storage.StreamAdaptor")
	defer span.End()

	// bundleOrErr represents a fetched entry bundle and its params, or an error if we couldn't fetch it for
	// some reason.
	type bundleOrErr struct {
		ri  layout.RangeInfo
		b   []byte
		err error
	}

	// bundles will be filled with futures for in-order entry bundles by the worker
	// go routines below.
	// This channel will be drained by the loop at the bottom of this func which
	// yields the bundles to the caller.
	bundles := make(chan func() bundleOrErr, numWorkers)
	exit := make(chan struct{})

	// Fetch entry bundle resources in parallel.
	// We use a limited number of tokens here to prevent this from
	// consuming an unbounded amount of resources.
	go func() {
		ctx, span := tracer.Start(ctx, "tessera.storage.StreamAdaptorWorker")
		defer span.End()

		defer close(bundles)

		// We'll limit ourselves to numWorkers worth of on-going work using these tokens:
		tokens := make(chan struct{}, numWorkers)
		for range numWorkers {
			tokens <- struct{}{}
		}

		// We'll keep looping around until told to exit.
		for {
			// Check afresh what size the tree is so we can keep streaming entries as the tree grows.
			treeSize, err := getSize(ctx)
			if err != nil {
				klog.Warningf("StreamAdaptor: failed to get current tree size: %v", err)
				continue
			}
			klog.V(1).Infof("StreamAdaptor: streaming from %d to %d", fromEntry, treeSize)

			// For each bundle, pop a future into the bundles channel and kick off an async request
			// to resolve it.
		rangeLoop:
			for ri := range layout.Range(fromEntry, treeSize, treeSize) {
				select {
				case <-exit:
					break rangeLoop
				case <-tokens:
					// We'll return a token below, once the bundle is fetched _and_ is being yielded.
				}

				c := make(chan bundleOrErr, 1)
				go func(ri layout.RangeInfo) {
					b, err := getBundle(ctx, ri.Index, ri.Partial)
					c <- bundleOrErr{ri: ri, b: b, err: err}
				}(ri)

				f := func() bundleOrErr {
					b := <-c
					// We're about to yield a value, so we can now return the token and unblock another fetch.
					tokens <- struct{}{}
					return b
				}

				bundles <- f
			}

			// Next loop, carry on from where we got to.
			fromEntry = treeSize

			select {
			case <-exit:
				klog.Infof("StreamAdaptor: exiting")
				return
			case <-time.After(time.Second):
				// We've caught up with and hit the end of the tree, so wait a bit before looping to avoid busy waiting.
				// TODO(al): could consider a shallow channel of sizes here.
			}
		}
	}()

	cancel = func() {
		close(exit)
	}

	var streamErr error
	next = func() (layout.RangeInfo, []byte, error) {
		if streamErr != nil {
			return layout.RangeInfo{}, nil, streamErr
		}

		f, ok := <-bundles
		if !ok {
			streamErr = ErrNoMoreEntries
			return layout.RangeInfo{}, nil, streamErr
		}
		b := f()
		if b.err != nil {
			streamErr = b.err
		}
		return b.ri, b.b, b.err
	}
	return next, cancel
}

// EntryStreamReader converts a stream of {RangeInfo, EntryBundle} into a stream of individually processed entries.
//
// This is mostly useful to Follower implementations which need to parse and consume individual entries being streamed
// from a LogReader.
type EntryStreamReader[T any] struct {
	bundleFn func([]byte) ([]T, error)
	next     func() (layout.RangeInfo, []byte, error)

	curData []T
	curRI   layout.RangeInfo
	i       uint64
}

// NewEntryStreamReader creates a new stream reader which uses the provided bundleFn to process bundles into processed entries of type T.
//
// Different bundleFn implementations can be provided to return raw entry bytes, parsed entry structs, or derivations of entries (e.g. hashes) as needed.
func NewEntryStreamReader[T any](next func() (layout.RangeInfo, []byte, error), bundleFn func([]byte) ([]T, error)) *EntryStreamReader[T] {
	return &EntryStreamReader[T]{
		bundleFn: bundleFn,
		next:     next,
		i:        0,
	}
}

// Next processes and returns the next available entry in the stream along with its index in the log.
func (e *EntryStreamReader[T]) Next() (uint64, T, error) {
	var t T
	if len(e.curData) == 0 {
		var err error
		var b []byte
		e.curRI, b, err = e.next()
		if err != nil {
			return 0, t, fmt.Errorf("next: %v", err)
		}
		e.curData, err = e.bundleFn(b)
		if err != nil {
			return 0, t, fmt.Errorf("bundleFn(bundleEntry @%d): %v", e.curRI.Index, err)

		}
		if e.curRI.First > 0 {
			e.curData = e.curData[e.curRI.First:]
		}
		if len(e.curData) > int(e.curRI.N) {
			e.curData = e.curData[:e.curRI.N]
		}
		e.i = 0
	}
	t, e.curData = e.curData[0], e.curData[1:]
	rIdx := e.curRI.Index*layout.EntryBundleWidth + uint64(e.curRI.First) + e.i
	e.i++
	return rIdx, t, nil
}
