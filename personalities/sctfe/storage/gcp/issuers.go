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

// AddIssuers stores Issuers values under their Key if there isn't an object under Key already.
func (s *IssuersStorage) AddIssuersIfNotExist(ctx context.Context, kv []sctfe.KV) error {
	// We first try and see if this issuer cert has already been stored since reads
	// are cheaper than writes.
	// TODO(phboneff): add parallel operations
	for _, kv := range kv {
		objName := s.keyToObjName(kv.K)
		obj := s.bucket.Object(objName)

		w := obj.If(gcs.Conditions{DoesNotExist: true}).NewWriter(ctx)
		w.ObjectAttrs.ContentType = s.contentType

		if _, err := w.Write(kv.V); err != nil {
			return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket.BucketName(), err)
		}

		if err := w.Close(); err != nil {
			if ee, ok := err.(*googleapi.Error); ok && ee.Code == http.StatusPreconditionFailed {
				for _, e := range ee.Errors {
					if e.Reason == "conditionNotMet" {
						klog.V(2).Infof("AddIssuersIfNotExist: object %q already exists in bucket %q, continuing", objName, s.bucket.BucketName())
						return nil
					}
				}
			}

			return fmt.Errorf("failed to close write on %q: %v", objName, err)
		}

		klog.V(2).Infof("AddIssuersIfNotExist: added %q in bucket %q", objName, s.bucket.BucketName())
	}
	return nil
}
