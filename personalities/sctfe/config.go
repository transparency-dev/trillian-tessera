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
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"
)

// ValidatedLogConfig represents the LogConfig with the information that has
// been successfully parsed as a result of validating it.
type ValidatedLogConfig struct {
	Config  *LogConfig
	PubKey  crypto.PublicKey
	PrivKey proto.Message
	// If set, ExtKeyUsages will restrict the set of such usages that the
	// server will accept. By default all are accepted. The values specified
	// must be ones known to the x509 package.
	KeyUsages     []x509.ExtKeyUsage
	NotAfterStart *time.Time
	NotAfterLimit *time.Time
}

// TODO(phboneff): inline this in ValidatedLogConfig and probably inline things further
// TODO(phboneff): edit comments
type LogConfig struct {
	// origin identifies the log. It will be used in its checkpoint, and
	// is also its submission prefix, as per https://c2sp.org/static-ct-api
	Origin string
	// Path to the file containing root certificates that are acceptable to the
	// log. The certs are served through get-roots endpoint.
	RootsPemFile string
	// The private key used for signing Checkpoints or SCTs.
	PrivateKey *anypb.Any
	// The public key matching the above private key (if both are present).
	// It can be specified for the convenience of test tools, but it not used
	// by the server.
	PublicKey *keyspb.PublicKey
	// If reject_expired is true then the certificate validity period will be
	// checked against the current time during the validation of submissions.
	// This will cause expired certificates to be rejected.
	RejectExpired bool
	// If reject_unexpired is true then CTFE rejects certificates that are either
	// currently valid or not yet valid.
	RejectUnexpired bool
	// not_after_start defines the start of the range of acceptable NotAfter
	// values, inclusive.
	// Leaving this unset implies no lower bound to the range.
	NotAfterStart *timestamppb.Timestamp
	// not_after_limit defines the end of the range of acceptable NotAfter values,
	// exclusive.
	// Leaving this unset implies no upper bound to the range.
	NotAfterLimit *timestamppb.Timestamp
	// accept_only_ca controls whether or not *only* certificates with the CA bit
	// set will be accepted.
	AcceptOnlyCa bool
	// A list of X.509 extension OIDs, in dotted string form (e.g. "2.3.4.5")
	// which, if present, should cause submissions to be rejected.
	RejectExtensions []string
}

// LogConfigFromFile creates a LogConfig options from the given
// filename, which should contain text or binary-encoded protobuf configuration
// data.
func LogConfigFromFile(filename string) (*configpb.LogConfig, error) {
	cfgBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg configpb.LogConfig
	if txtErr := prototext.Unmarshal(cfgBytes, &cfg); txtErr != nil {
		if binErr := proto.Unmarshal(cfgBytes, &cfg); binErr != nil {
			return nil, fmt.Errorf("failed to parse LogConfig from %q as text protobuf (%v) or binary protobuf (%v)", filename, txtErr, binErr)
		}
	}

	return &cfg, nil
}

// ValidateLogConfig checks that a single log config is valid. In particular:
//   - A log has a private, and optionally a public key (both valid).
//   - Each of NotBeforeStart and NotBeforeLimit, if set, is a valid timestamp
//     proto. If both are set then NotBeforeStart <= NotBeforeLimit.
//   - Merge delays (if present) are correct.
//
// Returns the validated structures (useful to avoid double validation).
func ValidateLogConfig(cfg *configpb.LogConfig, origin string, projectID string, bucket string, spannerDB string, rootsPemFile string, rejectExpired bool, rejectUnexpired bool, extKeyUsages string, rejectExtensions string) (*ValidatedLogConfig, error) {
	if origin == "" {
		return nil, errors.New("empty origin")
	}

	// TODO(phboneff): move this logic together with the tests out of config.go and validate the flags directly
	if projectID == "" {
		return nil, errors.New("empty projectID")
	}

	if bucket == "" {
		return nil, errors.New("empty bucket")
	}

	if spannerDB == "" {
		return nil, errors.New("empty spannerDB")
	}

	lExtKeyUsages := []string{}
	lRejectExtensions := []string{}
	if extKeyUsages != "" {
		lExtKeyUsages = strings.Split(extKeyUsages, ",")
	}
	if rejectExtensions != "" {
		lRejectExtensions = strings.Split(rejectExtensions, ",")
	}

	vCfg := ValidatedLogConfig{Config: &LogConfig{
		Origin:           origin,
		RootsPemFile:     rootsPemFile,
		PrivateKey:       cfg.PrivateKey,
		PublicKey:        cfg.PublicKey,
		RejectExpired:    rejectExpired,
		RejectUnexpired:  rejectUnexpired,
		NotAfterStart:    cfg.NotAfterLimit,
		NotAfterLimit:    cfg.NotAfterLimit,
		AcceptOnlyCa:     cfg.AcceptOnlyCa,
		RejectExtensions: lRejectExtensions,
	}}

	// Validate the public key.
	if pubKey := cfg.PublicKey; pubKey != nil {
		var err error
		if vCfg.PubKey, err = x509.ParsePKIXPublicKey(pubKey.Der); err != nil {
			return nil, fmt.Errorf("x509.ParsePKIXPublicKey: %w", err)
		}
	}

	// Validate the private key.
	if cfg.PrivateKey == nil {
		return nil, errors.New("empty private key")
	}
	privKey, err := cfg.PrivateKey.UnmarshalNew()
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}
	vCfg.PrivKey = privKey

	if rejectExpired && rejectUnexpired {
		return nil, errors.New("rejecting all certificates")
	}

	// Validate the extended key usages list.
	for _, kuStr := range lExtKeyUsages {
		if ku, ok := stringToKeyUsage[kuStr]; ok {
			// If "Any" is specified, then we can ignore the entire list and
			// just disable EKU checking.
			if ku == x509.ExtKeyUsageAny {
				klog.Infof("%s: Found ExtKeyUsageAny, allowing all EKUs", origin)
				vCfg.KeyUsages = nil
				break
			}
			vCfg.KeyUsages = append(vCfg.KeyUsages, ku)
		} else {
			return nil, fmt.Errorf("unknown extended key usage: %s", kuStr)
		}
	}

	// Validate the time interval.
	start, limit := cfg.NotAfterStart, cfg.NotAfterLimit
	if start != nil {
		vCfg.NotAfterStart = &time.Time{}
		if err := start.CheckValid(); err != nil {
			return nil, fmt.Errorf("invalid start timestamp: %v", err)
		}
		*vCfg.NotAfterStart = start.AsTime()
	}
	if limit != nil {
		vCfg.NotAfterLimit = &time.Time{}
		if err := limit.CheckValid(); err != nil {
			return nil, fmt.Errorf("invalid limit timestamp: %v", err)
		}
		*vCfg.NotAfterLimit = limit.AsTime()
	}
	if start != nil && limit != nil && (*vCfg.NotAfterLimit).Before(*vCfg.NotAfterStart) {
		return nil, errors.New("limit before start")
	}

	return &vCfg, nil
}

var stringToKeyUsage = map[string]x509.ExtKeyUsage{
	"Any":                        x509.ExtKeyUsageAny,
	"ServerAuth":                 x509.ExtKeyUsageServerAuth,
	"ClientAuth":                 x509.ExtKeyUsageClientAuth,
	"CodeSigning":                x509.ExtKeyUsageCodeSigning,
	"EmailProtection":            x509.ExtKeyUsageEmailProtection,
	"IPSECEndSystem":             x509.ExtKeyUsageIPSECEndSystem,
	"IPSECTunnel":                x509.ExtKeyUsageIPSECTunnel,
	"IPSECUser":                  x509.ExtKeyUsageIPSECUser,
	"TimeStamping":               x509.ExtKeyUsageTimeStamping,
	"OCSPSigning":                x509.ExtKeyUsageOCSPSigning,
	"MicrosoftServerGatedCrypto": x509.ExtKeyUsageMicrosoftServerGatedCrypto,
	"NetscapeServerGatedCrypto":  x509.ExtKeyUsageNetscapeServerGatedCrypto,
}
