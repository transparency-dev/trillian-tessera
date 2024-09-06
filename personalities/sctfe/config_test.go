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
	"fmt"
	"strings"
	"testing"

	"github.com/google/trillian/crypto/keyspb"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	invalidTimestamp = &timestamppb.Timestamp{Nanos: int32(1e9)}
)

func mustMarshalAny(pb proto.Message) *anypb.Any {
	ret, err := anypb.New(pb)
	if err != nil {
		panic(fmt.Sprintf("MarshalAny failed: %v", err))
	}
	return ret
}

func TestValidateLogConfig(t *testing.T) {
	privKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "../testdata/ct-http-server.privkey.pem", Password: "dirk"})

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
			desc:      "empty-private-key",
			wantErr:   "empty private key",
			cfg:       &configpb.LogConfig{},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:    "invalid-private-key",
			wantErr: "invalid private key",
			cfg: &configpb.LogConfig{
				PrivateKey: &anypb.Any{},
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:    "empty-projectID",
			wantErr: "empty projectID",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:    "testlog",
			projectID: "",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:    "empty-bucket",
			wantErr: "empty bucket",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "",
			spannerDB: "spanner",
		},
		{
			desc:    "empty-spannerDB",
			wantErr: "empty spannerDB",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "",
		},
		{
			desc:    "rejecting-all",
			wantErr: "rejecting all certificates",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:          "testlog",
			projectID:       "project",
			bucket:          "bucket",
			spannerDB:       "spanner",
			rejectExpired:   true,
			rejectUnexpired: true,
		},
		{
			desc:    "unknown-ext-key-usage-1",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "wrong_usage",
		},
		{
			desc:    "unknown-ext-key-usage-2",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "ClientAuth,ServerAuth,TimeStomping",
		},
		{
			desc:    "unknown-ext-key-usage-3",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "Any ",
		},
		{
			desc:    "invalid-start-timestamp",
			wantErr: "invalid start timestamp",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: invalidTimestamp,
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:    "invalid-limit-timestamp",
			wantErr: "invalid limit timestamp",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterLimit: invalidTimestamp,
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc:    "limit-before-start",
			wantErr: "limit before start",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 200},
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 100},
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc: "ok",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			// Note: Substituting an arbitrary proto.Message as a PrivateKey will not
			// fail the validation because the actual key loading happens at runtime.
			// TODO(pavelkalinnikov): Decouple key protos validation and loading, and
			// make this test fail.
			desc: "ok-not-a-key",
			cfg: &configpb.LogConfig{
				PrivateKey: mustMarshalAny(&configpb.LogConfig{}),
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc: "ok-ext-key-usages",
			cfg: &configpb.LogConfig{
				PrivateKey: privKey,
			},
			origin:       "testlog",
			projectID:    "project",
			bucket:       "bucket",
			spannerDB:    "spanner",
			extKeyUsages: "ServerAuth,ClientAuth,OCSPSigning",
		},
		{
			desc: "ok-start-timestamp",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 100},
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc: "ok-limit-timestamp",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 200},
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
		{
			desc: "ok-range-timestamp",
			cfg: &configpb.LogConfig{
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 300},
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 400},
			},
			origin:    "testlog",
			projectID: "project",
			bucket:    "bucket",
			spannerDB: "spanner",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			vc, err := ValidateLogConfig(tc.cfg, tc.origin, tc.projectID, tc.bucket, tc.spannerDB, "", tc.rejectExpired, tc.rejectUnexpired, tc.extKeyUsages, "")
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
