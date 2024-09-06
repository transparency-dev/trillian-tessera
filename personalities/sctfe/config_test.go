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
	"crypto"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/trillian/crypto/keys/pem"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func mustMarshalAny(pb proto.Message) *anypb.Any {
	ret, err := anypb.New(pb)
	if err != nil {
		panic(fmt.Sprintf("MarshalAny failed: %v", err))
	}
	return ret
}

func TestValidateLogConfig(t *testing.T) {
	signer, err := pem.ReadPrivateKeyFile("./testdata/ct-http-server.privkey.pem", "dirk")
	if err != nil {
		t.Fatalf("Can't open key: %v", err)
	}

	t100 := time.Unix(100, 0)
	t200 := time.Unix(200, 0)

	for _, tc := range []struct {
		desc            string
		cfg             *configpb.LogConfig
		origin          string
		projectID       string
		bucket          string
		spannerDB       string
		wantErr         string
		rejectExpired   bool
		rejectUnexpired bool
		extKeyUsages    string
		notAfterStart   *time.Time
		notAfterLimit   *time.Time
		signer          crypto.Signer
	}{
		{
			desc:      "empty-origin",
			wantErr:   "empty origin",
			cfg:       &configpb.LogConfig{},
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:      "empty-projectID",
			wantErr:   "empty projectID",
			cfg:       &configpb.LogConfig{},
			origin:    "testlog",
			projectID: "",
			bucket:    "bucket",
			spannerDB: "spanner",
			signer:    signer,
		},
		{
			desc:      "empty-bucket",
			wantErr:   "empty bucket",
			cfg:       &configpb.LogConfig{},
			origin:    "testlog",
			projectID: "project",
			bucket:    "",
			spannerDB: "spanner",
			signer:    signer,
		},
		{
			desc:      "empty-spannerDB",
			wantErr:   "empty spannerDB",
			cfg:       &configpb.LogConfig{},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "",
			signer:    signer,
		},
		{
			desc:            "rejecting-all",
			wantErr:         "rejecting all certificates",
			cfg:             &configpb.LogConfig{},
			origin:          "testlog",
			projectID:       "project",
			bucket:          "bucket",
			spannerDB:       "spanner",
			rejectExpired:   true,
			rejectUnexpired: true,
			signer:          signer,
		},
		{
			desc:         "unknown-ext-key-usage-1",
			wantErr:      "unknown extended key usage",
			cfg:          &configpb.LogConfig{},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "wrong_usage",
			signer:       signer,
		},
		{
			desc:         "unknown-ext-key-usage-2",
			wantErr:      "unknown extended key usage",
			cfg:          &configpb.LogConfig{},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "ClientAuth,ServerAuth,TimeStomping",
			signer:       signer,
		},
		{
			desc:         "unknown-ext-key-usage-3",
			wantErr:      "unknown extended key usage",
			cfg:          &configpb.LogConfig{},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "Any ",
			signer:       signer,
		},
		{
			desc:          "limit-before-start",
			wantErr:       "limit before start",
			cfg:           &configpb.LogConfig{},
			origin:        "testlog",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			notAfterStart: &t200,
			notAfterLimit: &t100,
			signer:        signer,
		},
		{
			desc:      "ok",
			cfg:       &configpb.LogConfig{},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
			signer:    signer,
		},
		{
			desc:         "ok-ext-key-usages",
			cfg:          &configpb.LogConfig{},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "ServerAuth,ClientAuth,OCSPSigning",
			signer:       signer,
		},
		{
			desc:          "ok-start-timestamp",
			cfg:           &configpb.LogConfig{},
			origin:        "testlog",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			notAfterStart: &t100,
			signer:        signer,
		},
		{
			desc:          "ok-limit-timestamp",
			cfg:           &configpb.LogConfig{},
			origin:        "testlog",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			notAfterStart: &t200,
			signer:        signer,
		},
		{
			desc:          "ok-range-timestamp",
			cfg:           &configpb.LogConfig{},
			origin:        "testlog",
			projectID:     "project",
			bucket:        "bucket",
			spannerDB:     "spanner",
			notAfterStart: &t100,
			notAfterLimit: &t200,
			signer:        signer,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			vc, err := ValidateLogConfig(tc.cfg, tc.origin, tc.projectID, tc.bucket, tc.spannerDB, "", tc.rejectExpired, tc.rejectUnexpired, tc.extKeyUsages, "", tc.notAfterStart, tc.notAfterLimit, signer)
			if len(tc.wantErr) == 0 && err != nil {
				t.Errorf("ValidateLogConfig()=%v, want nil", err)
			}
			if len(tc.wantErr) > 0 && (err == nil || !strings.Contains(err.Error(), tc.wantErr)) {
				t.Errorf("ValidateLogConfig()=%v, want err containing %q", err, tc.wantErr)
			}
			if err == nil && vc == nil {
				t.Error("err and ValidatedLogConfig are both nil")
			}
			// TODO(pavelkalinnikov): Test that ValidatedLogConfig is correct.
		})
	}
}
