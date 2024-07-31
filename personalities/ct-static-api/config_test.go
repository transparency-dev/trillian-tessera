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

package ctfe

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/trillian/crypto/keyspb"
	"github.com/transparency-dev/trillian-tessera/personalities/ct-static-api/configpb"
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
	//pubKey := mustReadPublicKey("../testdata/ct-http-server.pubkey.pem")
	privKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "../testdata/ct-http-server.privkey.pem", Password: "dirk"})

	for _, tc := range []struct {
		desc    string
		cfg     *configpb.LogConfig
		wantErr string
	}{
		{
			desc:    "empty-submission-prefix",
			wantErr: "empty log origin",
			cfg:     &configpb.LogConfig{},
		},
		{
			desc:    "empty-private-key",
			wantErr: "empty private key",
			cfg:     &configpb.LogConfig{Origin: "testlog"},
		},
		{
			desc:    "invalid-private-key",
			wantErr: "invalid private key",
			cfg: &configpb.LogConfig{
				Origin:     "testlog",
				PrivateKey: &anypb.Any{},
			},
		},
		{
			desc:    "empty-storage-config",
			wantErr: "empty storage config",
			cfg: &configpb.LogConfig{
				Origin:     "testlog",
				PrivateKey: privKey,
			},
		},
		{
			desc:    "rejecting-all",
			wantErr: "rejecting all certificates",
			cfg: &configpb.LogConfig{
				Origin:          "testlog",
				RejectExpired:   true,
				RejectUnexpired: true,
				PrivateKey:      privKey,
				StorageConfig:   &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "unknown-ext-key-usage-1",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				ExtKeyUsages:  []string{"wrong_usage"},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "unknown-ext-key-usage-2",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				ExtKeyUsages:  []string{"ClientAuth", "ServerAuth", "TimeStomping"},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "unknown-ext-key-usage-3",
			wantErr: "unknown extended key usage",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				ExtKeyUsages:  []string{"Any "},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "invalid-start-timestamp",
			wantErr: "invalid start timestamp",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterStart: invalidTimestamp,
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "invalid-limit-timestamp",
			wantErr: "invalid limit timestamp",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterLimit: invalidTimestamp,
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "limit-before-start",
			wantErr: "limit before start",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 200},
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 100},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "negative-maximum-merge",
			wantErr: "negative maximum merge",
			cfg: &configpb.LogConfig{
				Origin:           "testlog",
				PrivateKey:       privKey,
				MaxMergeDelaySec: -100,
				StorageConfig:    &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "negative-expected-merge",
			wantErr: "negative expected merge",
			cfg: &configpb.LogConfig{
				Origin:                "testlog",
				PrivateKey:            privKey,
				ExpectedMergeDelaySec: -100,
				StorageConfig:         &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc:    "expected-exceeds-max",
			wantErr: "expected merge delay exceeds MMD",
			cfg: &configpb.LogConfig{
				Origin:                "testlog",
				PrivateKey:            privKey,
				MaxMergeDelaySec:      50,
				ExpectedMergeDelaySec: 100,
				StorageConfig:         &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			// Note: Substituting an arbitrary proto.Message as a PrivateKey will not
			// fail the validation because the actual key loading happens at runtime.
			// TODO(pavelkalinnikov): Decouple key protos validation and loading, and
			// make this test fail.
			desc: "ok-not-a-key",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    mustMarshalAny(&configpb.LogConfig{}),
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok-ext-key-usages",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				ExtKeyUsages:  []string{"ServerAuth", "ClientAuth", "OCSPSigning"},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok-start-timestamp",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 100},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok-limit-timestamp",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 200},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok-range-timestamp",
			cfg: &configpb.LogConfig{
				Origin:        "testlog",
				PrivateKey:    privKey,
				NotAfterStart: &timestamppb.Timestamp{Seconds: 300},
				NotAfterLimit: &timestamppb.Timestamp{Seconds: 400},
				StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
		{
			desc: "ok-merge-delay",
			cfg: &configpb.LogConfig{
				Origin:                "testlog",
				PrivateKey:            privKey,
				MaxMergeDelaySec:      86400,
				ExpectedMergeDelaySec: 7200,
				StorageConfig:         &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			vc, err := ValidateLogConfig(tc.cfg)
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

func TestValidateLogConfigSet(t *testing.T) {
	privKey := mustMarshalAny(&keyspb.PEMKeyFile{Path: "../testdata/ct-http-server.privkey.pem", Password: "dirk"})
	for _, tc := range []struct {
		desc    string
		cfg     *configpb.LogConfigSet
		wantErr string
	}{
		// TODO(phboneff): add config for multiple storage
		{
			desc:    "duplicate-prefix",
			wantErr: "duplicate origin",
			cfg: &configpb.LogConfigSet{
				Config: []*configpb.LogConfig{
					{
						Origin:        "pref1",
						StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
						PrivateKey:    privKey,
					},
					{
						Origin:        "pref1",
						StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
						PrivateKey:    privKey,
					},
				},
			},
		},
		{
			desc: "ok-all-distinct",
			cfg: &configpb.LogConfigSet{
				Config: []*configpb.LogConfig{
					{
						Origin:        "pref1",
						StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
						PrivateKey:    privKey,
					},
					{
						Origin:        "pref2",
						StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
						PrivateKey:    privKey,
					},
					{
						Origin:        "pref3",
						StorageConfig: &configpb.LogConfig_Gcp{Gcp: &configpb.GCPConfig{Bucket: "bucket", SpannerDbPath: "spanner"}},
						PrivateKey:    privKey,
					},
				},
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := ValidateLogConfigSet(tc.cfg)
			if len(tc.wantErr) == 0 && err != nil {
				t.Fatalf("ValidateLogConfigSet()=%v, want nil", err)
			}
			if len(tc.wantErr) > 0 && (err == nil || !strings.Contains(err.Error(), tc.wantErr)) {
				t.Errorf("ValidateLogConfigSet()=%v, want err containing %q", err, tc.wantErr)
			}
		})
	}
}
