// Copyright 2016 Google LLC. All Rights Reserved.
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

package sctfe

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/trillian/crypto/keys/pem"
	"github.com/google/trillian/monitoring"
	"golang.org/x/mod/sumdb/note"
)

func fakeCTStorage(_ context.Context, _ note.Signer) (*CTStorage, error) {
	return &CTStorage{}, nil
}

func TestSetUpInstance(t *testing.T) {
	ctx := context.Background()

	signer, err := pem.ReadPrivateKeyFile("./testdata/ct-http-server.privkey.pem", "dirk")
	if err != nil {
		t.Fatalf("Can't open key: %v", err)
	}

	var tests = []struct {
		desc             string
		origin           string
		projectID        string
		bucket           string
		spannerDB        string
		rootsPemFile     string
		extKeyUsages     string
		rejectExtensions string
		signer           crypto.Signer
		ctStorage        func(context.Context, note.Signer) (*CTStorage, error)
		wantErr          string
	}{
		{
			desc:         "valid",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			ctStorage:    fakeCTStorage,
			signer:       signer,
		},
		{
			desc:         "no-roots",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/nofile",
			ctStorage:    fakeCTStorage,
			wantErr:      "failed to read trusted roots",
			signer:       signer,
		},
		{
			desc:         "missing-root-cert",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			ctStorage:    fakeCTStorage,
			rootsPemFile: "./testdata/bogus.cert",
			signer:       signer,
			wantErr:      "failed to read trusted roots",
		},
		{
			desc:         "valid-ekus-1",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			extKeyUsages: "Any",
			signer:       signer,
			ctStorage:    fakeCTStorage,
		},
		{
			desc:         "valid-ekus-2",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			extKeyUsages: "Any,ServerAuth,TimeStamping",
			signer:       signer,
			ctStorage:    fakeCTStorage,
		},
		{
			desc:             "valid-reject-ext",
			origin:           "log",
			projectID:        "project",
			bucket:           "bucket",
			spannerDB:        "spanner",
			rootsPemFile:     "./testdata/fake-ca.cert",
			rejectExtensions: "1.2.3.4,5.6.7.8",
			signer:           signer,
			ctStorage:        fakeCTStorage,
		},
		{
			desc:             "invalid-reject-ext",
			origin:           "log",
			projectID:        "project",
			bucket:           "bucket",
			spannerDB:        "spanner",
			ctStorage:        fakeCTStorage,
			rootsPemFile:     "./testdata/fake-ca.cert",
			rejectExtensions: "1.2.3.4,one.banana.two.bananas",
			signer:           signer,
			wantErr:          "one",
		},
		{
			desc:         "missing-create-storage",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			signer:       signer,
			wantErr:      "failed to initiate storage backend",
		},
		{
			desc:         "failing-create-storage",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			signer:       signer,
			ctStorage: func(_ context.Context, _ note.Signer) (*CTStorage, error) {
				return nil, fmt.Errorf("I failed")
			},
			wantErr: "failed to initiate storage backend",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			vCfg, err := ValidateLogConfig(test.origin, test.projectID, test.bucket, test.spannerDB, test.rootsPemFile, false, false, test.extKeyUsages, test.rejectExtensions, nil, nil, signer)
			if err != nil {
				t.Fatalf("ValidateLogConfig(): %v", err)
			}
			opts := InstanceOptions{Validated: vCfg, Deadline: time.Second, MetricFactory: monitoring.InertMetricFactory{}, CreateStorage: test.ctStorage}

			if _, err := SetUpInstance(ctx, opts); err != nil {
				if test.wantErr == "" {
					t.Errorf("SetUpInstance()=_,%v; want _,nil", err)
				} else if !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("SetUpInstance()=_,%v; want err containing %q", err, test.wantErr)
				}
				return
			}
			if test.wantErr != "" {
				t.Errorf("SetUpInstance()=_,nil; want err containing %q", test.wantErr)
			}
		})
	}
}

func equivalentTimes(a *time.Time, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		// b can't be nil as it would have returned above.
		return false
	}
	return a.Equal(*b)
}

func TestSetUpInstanceSetsValidationOpts(t *testing.T) {
	ctx := context.Background()

	start := time.Unix(10000, 0)
	limit := time.Unix(12000, 0)

	signer, err := pem.ReadPrivateKeyFile("./testdata/ct-http-server.privkey.pem", "dirk")
	if err != nil {
		t.Fatalf("Can't open key: %v", err)
	}

	var tests = []struct {
		desc          string
		origin        string
		projectID     string
		bucket        string
		spannerDB     string
		rootsPemFile  string
		notAfterStart *time.Time
		notAfterLimit *time.Time
		signer        crypto.Signer
	}{
		{
			desc:         "no validation opts",
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			signer:       signer,
		},
		{
			desc:          "notAfterStart only",
			origin:        "log",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			rootsPemFile:  "./testdata/fake-ca.cert",
			notAfterStart: &start,
		},
		{
			desc:          "notAfter range",
			origin:        "log",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			rootsPemFile:  "./testdata/fake-ca.cert",
			notAfterStart: &start,
			notAfterLimit: &limit,
			signer:        signer,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			vCfg, err := ValidateLogConfig(test.origin, test.projectID, test.bucket, test.spannerDB, test.rootsPemFile, false, false, "", "", test.notAfterStart, test.notAfterLimit, signer)
			if err != nil {
				t.Fatalf("ValidateLogConfig(): %v", err)
			}
			opts := InstanceOptions{Validated: vCfg, Deadline: time.Second, MetricFactory: monitoring.InertMetricFactory{}, CreateStorage: fakeCTStorage}

			inst, err := SetUpInstance(ctx, opts)
			if err != nil {
				t.Fatalf("%v: SetUpInstance() = %v, want no error", test.desc, err)
			}
			addChainHandler, ok := inst.Handlers["/"+test.origin+ct.AddChainPath]
			if !ok {
				t.Fatal("Couldn't find AddChain handler")
			}
			gotOpts := addChainHandler.Info.validationOpts
			if got, want := gotOpts.notAfterStart, test.notAfterStart; !equivalentTimes(got, want) {
				t.Errorf("%v: handler notAfterStart %v, want %v", test.desc, got, want)
			}
			if got, want := gotOpts.notAfterLimit, test.notAfterLimit; !equivalentTimes(got, want) {
				t.Errorf("%v: handler notAfterLimit %v, want %v", test.desc, got, want)
			}
		})
	}
}

func TestErrorMasking(t *testing.T) {
	info := logInfo{}
	w := httptest.NewRecorder()
	prefix := "Internal Server Error"
	err := errors.New("well that's bad")
	info.SendHTTPError(w, 500, err)
	if got, want := w.Body.String(), fmt.Sprintf("%s\n%v\n", prefix, err); got != want {
		t.Errorf("SendHTTPError: got %s, want %s", got, want)
	}
	info.instanceOpts.MaskInternalErrors = true
	w = httptest.NewRecorder()
	info.SendHTTPError(w, 500, err)
	if got, want := w.Body.String(), prefix+"\n"; got != want {
		t.Errorf("SendHTTPError: got %s, want %s", got, want)
	}

}
