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

// Package gcp implements SCTFE storage systems for issuers.
//
// The interfaces are defined in sctfe/storage.go
package gcp

import (
	"context"
	"fmt"
	"net/http"
	"path"

	gcs "cloud.google.com/go/storage"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"k8s.io/klog/v2"
)

// IssuersStorage is a key value store backed by GCS on GCP to store issuer chains.
type IssuersStorage struct {
	bucket      *gcs.BucketHandle
	prefix      string
	contentType string
}

// NewIssuerStorage creates a new GCSStorage.
//
// The specified bucket must exist or an error will be returned.
func NewIssuerStorage(ctx context.Context, projectID string, bucket string, prefix string, contentType string) (*IssuersStorage, error) {
	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	it := c.Buckets(ctx, projectID)
	for {
		bAttrs, err := it.Next()
		if err == iterator.Done {
			return nil, fmt.Errorf("bucket %q does not exist, please create it", bucket)
		}
		if err != nil {
			return nil, fmt.Errorf("error scanning buckets: %v", err)
		}
		if bAttrs.Name == bucket {
			break
		}
	}
	r := &IssuersStorage{
		bucket:      c.Bucket(bucket),
		prefix:      prefix,
		contentType: contentType,
	}

	return r, nil
}

// keyToObjName converts bytes to a GCS object name.
func (s *IssuersStorage) keyToObjName(key []byte) string {
	return path.Join(s.prefix, string(key))
}

// Exists checks whether a value is stored under key.
func (s *IssuersStorage) Exists(ctx context.Context, key []byte) (bool, error) {
	objName := s.keyToObjName(key)
	obj := s.bucket.Object(objName)
	_, err := obj.Attrs(ctx)
	if err == gcs.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("error fetching attributes for %q :%v", objName, err)
	}
	klog.V(2).Infof("Exists: object %q already exists in bucket %q", objName, s.bucket.BucketName())
	return true, nil
}

// AddIssuers stores all Issuers values under their Key.
//
// If there is already an object under a given key, it does not override it.
func (s *IssuersStorage) AddIssuers(ctx context.Context, kv []sctfe.KV) error {
	// We first try and see if this issuer cert has already been stored since reads
	// are cheaper than writes.
	// TODO(phboneff): monitor usage, eventually write directly depending on usage patterns
	toStore := []sctfe.KV{}
	for _, kv := range kv {
		ok, err := s.Exists(ctx, kv.K)
		if err != nil {
			return fmt.Errorf("error checking if issuer %q exists: %s", string(kv.K), err)
		}
		if !ok {
			toStore = append(toStore, kv)
		}
	}
	// TODO(phboneff): add parallel writes
	for _, kv := range toStore {
		objName := s.keyToObjName(kv.K)
		obj := s.bucket.Object(objName)

		// Don't overwrite if it already exists
		// TODO(phboneff): consider reading the object to make sure it's identical
		w := obj.If(gcs.Conditions{DoesNotExist: true}).NewWriter(ctx)
		w.ObjectAttrs.ContentType = s.contentType

		if _, err := w.Write(kv.V); err != nil {
			return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket.BucketName(), err)
		}

		if err := w.Close(); err != nil {
			// If we run into a precondition failure error, it means that the object already exists.
			if ee, ok := err.(*googleapi.Error); ok && ee.Code == http.StatusPreconditionFailed {
				klog.V(2).Infof("Add: object %q already exists in bucket %q, continuing", objName, s.bucket.BucketName())
				return nil
			}

			return fmt.Errorf("failed to close write on %q: %v", objName, err)
		}
	}
	return nil
}
