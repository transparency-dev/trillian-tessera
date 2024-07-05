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

// Package gcp contains a GCP-based storage implementation for Tessera.
//
// TODO: decide whether to rename this package.
//
// This storage implementation uses GCS for long-term storage and serving of
// entry bundles and log tiles, and Spanner for coordinating updates to GCS
// when multiple instances of a personality binary are running.
//
// A single GCS bucket is used to hold entry bundles and log internal tiles.
// The object keys for the bucket are selected so as to conform to the
// expected layout of a tile-based log.
//
// A Spanner database provides a transactional mechanism to allow multiple
// frontends to safely update the contents of the log.
package gcp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/spanner"
	gcs "cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"k8s.io/klog/v2"
)

// Storage is a GCP based storage implementation for Tessera.
type Storage struct {
	gcsClient *gcs.Client

	projectID string
	bucket    string

	dbPool *spanner.Client
}

// Config holds GCP project and resource configuration for a storage instance.
type Config struct {
	// ProjectID is the GCP project which hosts the storage bucket and Spanner database for the log.
	ProjectID string
	// Bucket is the name of the GCS bucket to use for storing log state.
	Bucket string
	// Spanner is the GCP resource URI of the spanner database instance to use.
	Spanner string
}

// New creates a new instance of the GCP based Storage.
func New(ctx context.Context, cfg Config) (*Storage, error) {
	c, err := gcs.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS storage: %v", err)
	}

	dbPool, err := spanner.NewClient(ctx, cfg.Spanner)
	if err != nil {
		klog.Exitf("Failed to connect to Spanner: %v", err)
	}

	r := &Storage{
		gcsClient: c,
		projectID: cfg.ProjectID,
		bucket:    cfg.Bucket,
		dbPool:    dbPool,
	}

	if err := r.initDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to init DB: %v", err)
	}

	if exists, err := r.bucketExists(ctx); err != nil {
		return nil, fmt.Errorf("failed to check whether bucket %q exists: %v", r.bucket, err)
	} else if !exists {
		return nil, fmt.Errorf("bucket %q does not exist, please create it", r.bucket)
	}

	return r, nil
}

// initDB ensures that the coordination DB is initialised correctly.
//
// The database schema consists of 3 tables:
//   - SeqCoord
//     This table only ever contains a single row which tracks the next available
//     sequence number.
//   - Seq
//     This table holds sequenced "batches" of entries. The batches are keyed
//     by the sequence number assigned to the first entry in the batch, and
//     each subsequent entry in the batch takes the numerically next sequence number.
//   - IntCoord
//     This table coordinates integration of the batches of entries stored in
//     Seq into the committed tree state.
//
// The database and schema should be created externally, e.g. by terraform.
func (s *Storage) initDB(ctx context.Context) error {

	/* Schema for reference:
	CREATE TABLE SeqCoord (
	 id INT64 NOT NULL,
	 next INT64 NOT NULL,
	) PRIMARY KEY (id);

	CREATE TABLE Seq (
		id INT64 NOT NULL,
		seq INT64 NOT NULL,
		v BYTES(MAX),
	) PRIMARY KEY (id, seq);

	CREATE TABLE IntCoord (
		id INT64 NOT NULL,
		seq INT64 NOT NULL,
	) PRIMARY KEY (id);
	*/

	// Set default values for a newly inisialised schema - these rows being present are a precondition for
	// sequencing and integration to occur.
	// Note that this will only succeed if no row exists, so there's no danger
	// of "resetting" an existing log.
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("SeqCoord", []string{"id", "next"}, []interface{}{0, 0})}); spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.Insert("IntCoord", []string{"id", "seq"}, []interface{}{0, 0})}); spanner.ErrCode(err) != codes.AlreadyExists {
		return err
	}
	return nil
}

// bucketExists tests whether the configured bucket exists.
func (s *Storage) bucketExists(ctx context.Context) (bool, error) {
	it := s.gcsClient.Buckets(ctx, s.projectID)
	for {
		bAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, err
		}
		if bAttrs.Name == s.bucket {
			return true, nil
		}
	}
	return false, nil
}

// getObject returns the data and generation of the specified object, or an error.
func (s *Storage) getObject(ctx context.Context, obj string) ([]byte, int64, error) {
	r, err := s.gcsClient.Bucket(s.bucket).Object(obj).NewReader(ctx)
	if err != nil {
		return nil, -1, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
	}
	defer r.Close()

	d, err := io.ReadAll(r)
	return d, r.Attrs.Generation, err
}

// setObject stores the provided data in the specified object, optionally gated by a condition.
//
// cond can be used to specify preconditions for the write (e.g. write iff not exists, write iff
// current generation is X, etc.), or nil can be passed if no preconditions are desired.
//
// When preconditions are specified and are not met, an error will be returned unless the currently
// stored data is bit-for-bit identical to the data to-be-written. This is intended to provide
// idempotentency for writes.
func (s *Storage) setObject(ctx context.Context, objName string, data []byte, cond *gcs.Conditions) error {
	bkt := s.gcsClient.Bucket(s.bucket)
	obj := bkt.Object(objName)

	var w *gcs.Writer
	if cond == nil {
		w = obj.NewWriter(ctx)

	} else {
		w = obj.If(*cond).NewWriter(ctx)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket, err)
	}

	if err := w.Close(); err != nil {
		// If we run into a precondition failure error, check that the object
		// which exists contains the same content that we want to write.
		// If so, we can consider this write to be idempotently successful.
		if ee, ok := err.(*googleapi.Error); ok && ee.Code == http.StatusPreconditionFailed {
			existing, existingGen, err := s.getObject(ctx, objName)
			if err != nil {
				return fmt.Errorf("failed to fetch existing content for %q (@%d): %v", objName, existingGen, err)
			}
			if !bytes.Equal(existing, data) {
				return fmt.Errorf("precondition failed: resource content for %q differs from data to-be-written", objName)
			}

			klog.V(2).Infof("setObject: identical resource already exists for %q, continuing", objName)
			return nil
		}

		return err
	}
	return nil
}
