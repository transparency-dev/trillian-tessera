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
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/trillian/crypto/keys"
	"github.com/google/trillian/crypto/keys/pem"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/google/trillian/monitoring"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"golang.org/x/mod/sumdb/note"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	keys.RegisterHandler(&keyspb.PEMKeyFile{}, pem.FromProto)
}

func fakeCTStorage(_ context.Context, _ note.Signer) (*CTStorage, error) {
	return &CTStorage{}, nil
}

func TestSetUpInstance(t *testing.T) {
	ctx := context.Background()

	privKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "./testdata/ct-http-server.privkey.pem", Password: "dirk"})
	missingPrivKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "./testdata/bogus.privkey.pem", Password: "dirk"})
	wrongPassPrivKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "./testdata/ct-http-server.privkey.pem", Password: "dirkly"})

	var tests = []struct {
		desc             string
		cfg              *configpb.LogConfig
		origin           string
		projectID        string
		bucket           string
		spannerDB        string
		rootsPemFile     string
		extKeyUsages     string
		rejectExtensions string
		ctStorage        func(context.Context, note.Signer) (*CTStorage, error)
		wantErr          string
	}{
		{
			desc: "valid",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			ctStorage:    fakeCTStorage,
		},
		{
			desc: "no-roots",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:    "log",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
			ctStorage: fakeCTStorage,
			wantErr:   "specify RootsPemFile",
		},
		{
			desc: "missing-root-cert",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			ctStorage:    fakeCTStorage,
			rootsPemFile: "./testdata/bogus.cert",
			wantErr:      "failed to read trusted roots",
		},
		{
			desc: "missing-privkey",
			cfg: &configpb.LogConfig{
				PrivateKey: missingPrivKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			ctStorage:    fakeCTStorage,
			wantErr:      "failed to load private key",
		},
		{
			desc: "privkey-wrong-password",
			cfg: &configpb.LogConfig{
				PrivateKey: wrongPassPrivKey,
			},
			origin:       "log",
			projectID:    "projeot",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			ctStorage:    fakeCTStorage,
			wantErr:      "failed to load private key",
		},
		{
			desc: "valid-ekus-1",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			extKeyUsages: "Any",
			ctStorage:    fakeCTStorage,
		},
		{
			desc: "valid-ekus-2",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			extKeyUsages: "Any,ServerAuth,TimeStamping",
			ctStorage:    fakeCTStorage,
		},
		{
			desc: "valid-reject-ext",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:           "log",
			projectID:        "project",
			bucket:           "bucket",
			spannerDB:        "spanner",
			rootsPemFile:     "./testdata/fake-ca.cert",
			rejectExtensions: "1.2.3.4,5.6.7.8",
			ctStorage:        fakeCTStorage,
		},
		{
			desc: "invalid-reject-ext",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:           "log",
			projectID:        "project",
			bucket:           "bucket",
			spannerDB:        "spanner",
			ctStorage:        fakeCTStorage,
			rootsPemFile:     "./testdata/fake-ca.cert",
			rejectExtensions: "1.2.3.4,one.banana.two.bananas",
			wantErr:          "one",
		},
		{
			desc: "missing-create-storage",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			wantErr:      "failed to initiate storage backend",
		},
		{
			desc: "failing-create-storage",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
			ctStorage: func(_ context.Context, _ note.Signer) (*CTStorage, error) {
				return nil, fmt.Errorf("I failed")
			},
			wantErr: "failed to initiate storage backend",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			vCfg, err := ValidateLogConfig(test.cfg, test.origin, test.projectID, test.bucket, test.spannerDB, test.rootsPemFile, false, false, test.extKeyUsages, test.rejectExtensions)
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

func equivalentTimes(a *time.Time, b *timestamppb.Timestamp) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		// b can't be nil as it would have returned above.
		return false
	}
	tsA := timestamppb.New(*a)
	return tsA.AsTime().Format(time.RFC3339Nano) == b.AsTime().Format(time.RFC3339Nano)
}

func TestSetUpInstanceSetsValidationOpts(t *testing.T) {
	ctx := context.Background()

	start := timestamppb.New(time.Unix(10000, 0))
	limit := timestamppb.New(time.Unix(12000, 0))

	privKey, err := anypb.New(&keyspb.PEMKeyFile{Path: "./testdata/ct-http-server.privkey.pem", Password: "dirk"})
	if err != nil {
		t.Fatalf("Could not marshal private key proto: %v", err)
	}
	var tests = []struct {
		desc         string
		cfg          *configpb.LogConfig
		origin       string
		projectID    string
		bucket       string
		spannerDB    string
		rootsPemFile string
	}{
		{
			desc: "no validation opts",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
		},
		{
			desc: "notAfterStart only",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: start,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
		},
		{
			desc: "notAfter range",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: start,
				NotAfterLimit: limit,
			},
			origin:       "log",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			rootsPemFile: "./testdata/fake-ca.cert",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			vCfg, err := ValidateLogConfig(test.cfg, test.origin, test.projectID, test.bucket, test.spannerDB, test.rootsPemFile, false, false, "", "")
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
			if got, want := gotOpts.notAfterStart, test.cfg.NotAfterStart; want != nil && !equivalentTimes(got, want) {
				t.Errorf("%v: handler notAfterStart %v, want %v", test.desc, got, want)
			}
			if got, want := gotOpts.notAfterLimit, test.cfg.NotAfterLimit; want != nil && !equivalentTimes(got, want) {
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
