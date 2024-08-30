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

// Package dedup limits the number of duplicate entries a personality allows in a Tessera log.
package dedup

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/client"
	"k8s.io/klog/v2"
)

// KV holds a LeafID and an Idx for deduplication
type KV struct {
	K []byte
	V uint64
}

type DedupStorage interface {
	Add(ctx context.Context, kvs []KV) error
	Get(ctx context.Context, key []byte) (uint64, bool, error)
}

type LocalDedupStorage interface {
	Add(ctx context.Context, kvs []KV) error
	Get(ctx context.Context, key []byte) (uint64, bool, error)
	LogSize() (uint64, error)
}

type LocalBEDedup struct {
	DedupStorage
	LogSize func() (uint64, error) // returns the largest contiguous idx Add has successfully been called with
	fetcher client.Fetcher
}

// NewLocalBestEffortDedup instantiates a local dedup storage and kicks off a synchronisation routine in the background.
func NewLocalBestEffortDedup(ctx context.Context, lds LocalDedupStorage, t time.Duration, f client.Fetcher, parseBundle func([]byte, uint64) ([]KV, error)) *LocalBEDedup {
	ret := &LocalBEDedup{DedupStorage: lds, LogSize: lds.LogSize, fetcher: f}
	go func() {
		tck := time.NewTicker(t)
		defer tck.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-tck.C:
				if err := ret.sync(ctx, parseBundle); err != nil {
					klog.Warningf("error updating deduplication data: %v", err)
				}
			}
		}
	}()
	return ret
}

// sync synchronises a deduplication storage with the corresponding log content.
func (d *LocalBEDedup) sync(ctx context.Context, parseBundle func([]byte, uint64) ([]KV, error)) error {
	cpRaw, err := d.fetcher(ctx, layout.CheckpointPath)
	if err != nil {
		return fmt.Errorf("error fetching checkpoint: %v", err)
	}
	// A https://c2sp.org/static-ct-api logsize is on the second line
	l := bytes.SplitN(cpRaw, []byte("\n"), 3)
	if len(l) < 2 {
		return errors.New("invalid checkpoint - no size")
	}
	ckptSize, err := strconv.ParseUint(string(l[1]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid checkpoint - can't extract size: %v", err)
	}
	oldSize, err := d.LogSize()
	if err != nil {
		return fmt.Errorf("OldSize(): %v", err)
	}

	// TODO(phboneff): add parallelism
	// Greatly inspired by
	// https://github.com/transparency-dev/trillian-tessera/blob/main/client/client.go
	if ckptSize > oldSize {
		klog.V(2).Infof("LocalBEDEdup.sync(): log at size %d, dedup database at size %d, startig to sync", ckptSize, oldSize)
		for i := oldSize / 256; i <= ckptSize/256; i++ {
			p := layout.EntriesPath(i, ckptSize)
			eRaw, err := d.fetcher(ctx, p)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Errorf("leaf bundle at index %d not found: %v", i, err)
				}
				return fmt.Errorf("failed to fetch leaf bundle at index %d: %v", i, err)
			}
			kvs, err := parseBundle(eRaw, i)
			if err != nil {
				return fmt.Errorf("parseBundle(): %v", err)
			}

			if err := d.Add(ctx, kvs); err != nil {
				return fmt.Errorf("error storing deduplication data for tile %d: %v", i, err)
			}
			klog.V(3).Infof("LocalBEDEdup.sync(): stored dedup data for entry bundle %d, %d more bundles to go", i, ckptSize/256-i)
		}
	}
	klog.V(3).Infof("LocalBEDEdup.sync(): dedup data synced to logsize %d", ckptSize)
	return nil
}
