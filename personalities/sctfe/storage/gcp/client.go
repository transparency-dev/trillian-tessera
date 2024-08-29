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
	"io"

	gcs "cloud.google.com/go/storage"
)

// GetFetcher returns a GCS read function for objects in a given bucket.
func GetFetcher(ctx context.Context, bucket string) (func(ctx context.Context, path string) ([]byte, error), error) {
	c, err := gcs.NewClient(ctx, gcs.WithJSONReads())
	if err != nil {
		return func(context.Context, string) ([]byte, error) { return nil, nil }, nil
	}
	return func(ctx context.Context, path string) ([]byte, error) {
		r, err := c.Bucket(bucket).Object(path).NewReader(ctx)
		if err != nil {
			return nil, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", path, bucket, err)
		}

		d, err := io.ReadAll(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read %q: %v", path, err)
		}
		return d, r.Close()
	}, nil
}
