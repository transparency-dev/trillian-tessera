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

// LeafIdx holds a LeafID and an Idx for deduplication
type LeafIdx struct {
	LeafID []byte
	Idx    uint64
}

type BEDedupStorage interface {
	Add(ctx context.Context, lidxs []LeafIdx) error
	Get(ctx context.Context, leafID []byte) (uint64, bool, error)
}

// TODO: re-architecture to prevent creating a LocaLBEDedupStorage without calling UpdateFromLog
type LocalBEDedupStorage interface {
	Add(ctx context.Context, lidxs []LeafIdx) error
	Get(ctx context.Context, leafID []byte) (uint64, bool, error)
	LogSize() (uint64, error)
}

type ParseBundleFunc func([]byte, uint64) ([]LeafIdx, error)

// UpdateFromLog synchronises a local best effort deduplication storage with a log.
func UpdateFromLog(ctx context.Context, lds LocalBEDedupStorage, t time.Duration, f client.Fetcher, pb ParseBundleFunc) {
	tck := time.NewTicker(t)
	defer tck.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tck.C:
			if err := sync(ctx, lds, pb, f); err != nil {
				klog.Warningf("error updating deduplication data: %v", err)
			}
		}
	}
}

// sync synchronises a deduplication storage with the corresponding log content.
func sync(ctx context.Context, lds LocalBEDedupStorage, pb ParseBundleFunc, f client.Fetcher) error {
	cpRaw, err := f(ctx, layout.CheckpointPath)
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
	oldSize, err := lds.LogSize()
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
			eRaw, err := f(ctx, p)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Errorf("leaf bundle at index %d not found: %v", i, err)
				}
				return fmt.Errorf("failed to fetch leaf bundle at index %d: %v", i, err)
			}
			lidxs, err := pb(eRaw, i)
			if err != nil {
				return fmt.Errorf("parseBundle(): %v", err)
			}

			if err := lds.Add(ctx, lidxs); err != nil {
				return fmt.Errorf("error storing deduplication data for tile %d: %v", i, err)
			}
			klog.V(3).Infof("LocalBEDEdup.sync(): stored dedup data for entry bundle %d, %d more bundles to go", i, ckptSize/256-i)
		}
	}
	klog.V(3).Infof("LocalBEDEdup.sync(): dedup data synced to logsize %d", ckptSize)
	return nil
}
