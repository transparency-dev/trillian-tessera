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

package tessera

import (
	"bytes"
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/client"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

func newCopier(numWorkers uint, storage MigrationWriter, getEntries client.EntryBundleFetcherFunc) *copier {
	return &copier{
		storage:    storage,
		getEntries: getEntries,
		todo:       make(chan bundle, numWorkers),
	}
}

// copier controls the migration work.
type copier struct {
	// storage is the target we're migrating to.
	storage    MigrationWriter
	getEntries client.EntryBundleFetcherFunc

	// todo contains work items to be completed.
	todo chan bundle

	// bundlesCopied is the number of entry bundles copied so far.
	bundlesCopied atomic.Uint64
}

// bundle represents the address of an individual entry bundle.
type bundle struct {
	Index   uint64
	Partial uint8
}

// Copy starts the work of copying sourceSize entries from the source to the target log.
//
// Only the entry bundles are copied as the target storage is expected to integrate them and recalculate the root.
// This is done to ensure the correctness of both the source log as well as the copy process itself.
//
// A call to this function will block until either the copying is done, or an error has occurred.
// It is an error if the resource copying completes ok but the resulting root hash does not match the provided sourceRoot.
func (c *copier) Copy(ctx context.Context, sourceSize uint64, sourceRoot []byte) error {
	klog.Infof("Starting copy; source size %d root %x", sourceSize, sourceRoot)

	// init
	targetSize, err := c.storage.IntegratedSize(ctx)
	if err != nil {
		return fmt.Errorf("size: %v", err)
	}
	if targetSize > sourceSize {
		return fmt.Errorf("target size %d > source size %d", targetSize, sourceSize)
	}

	go c.populateWork(targetSize, sourceSize)

	// Print stats
	go func() {
		bundlesToCopy := (sourceSize / layout.EntryBundleWidth) - (targetSize / layout.EntryBundleWidth) + 1
		if bundlesToCopy == 0 {
			return
		}
		for {
			time.Sleep(time.Second)
			bn := c.bundlesCopied.Load()
			bnp := float64(bn*100) / float64(bundlesToCopy)
			s, err := c.storage.IntegratedSize(ctx)
			if err != nil {
				klog.Warningf("Size: %v", err)
			}
			intp := float64(s*100) / float64(sourceSize)
			klog.Infof("integration: %d (%.2f%%)  bundles: %d (%.2f%%)", s, intp, bn, bnp)
		}
	}()

	// Do the copying
	eg := errgroup.Group{}
	for range cap(c.todo) {
		eg.Go(func() error {
			return c.worker(ctx)
		})
	}
	var root []byte
	eg.Go(func() error {
		r, err := c.storage.AwaitIntegration(ctx, sourceSize)
		if err != nil {
			return fmt.Errorf("copy failed: %v", err)
		}
		root = r
		return nil
	})

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("copy failed: %v", err)
	}

	if !bytes.Equal(root, sourceRoot) {
		return fmt.Errorf("copy completed, but local root hash %x != source root hash %x", root, sourceRoot)
	}

	klog.Infof("Copy successful.")
	return nil
}

// populateWork sends entries to the `todo` work channel.
// Each entry corresponds to an individual entryBundle which needs to be copied.
func (m *copier) populateWork(from, treeSize uint64) {
	klog.Infof("Spans for entry range [%d, %d)", from, treeSize)
	defer close(m.todo)

	for ri := range layout.Range(from, treeSize-from, treeSize) {
		m.todo <- bundle{Index: ri.Index, Partial: ri.Partial}
	}
}

// worker undertakes work items from the `todo` channel.
//
// It will attempt to retry failed operations several times before giving up, this should help
// deal with any transient errors which may occur.
func (m *copier) worker(ctx context.Context) error {
	for b := range m.todo {
		err := retry.Do(func() error {
			d, err := m.getEntries(ctx, b.Index, uint8(b.Partial))
			if err != nil {
				wErr := fmt.Errorf("failed to fetch entrybundle %d (p=%d): %v", b.Index, b.Partial, err)
				klog.Infof("%v", wErr)
				return wErr
			}
			if err := m.storage.SetEntryBundle(ctx, b.Index, b.Partial, d); err != nil {
				wErr := fmt.Errorf("failed to store entrybundle %d (p=%d): %v", b.Index, b.Partial, err)
				klog.Infof("%v", wErr)
				return wErr
			}
			m.bundlesCopied.Add(1)
			return nil
		},
			retry.Attempts(10),
			retry.DelayType(retry.BackOffDelay))
		if err != nil {
			klog.Infof("retry: %v", err)
			return err
		}
	}
	return nil
}
