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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/certificate-transparency-go/asn1"
	"github.com/google/certificate-transparency-go/x509util"
	"github.com/google/trillian/crypto/keys"
	"github.com/google/trillian/monitoring"
	"golang.org/x/mod/sumdb/note"
)

// InstanceOptions describes the options for a log instance.
type InstanceOptions struct {
	// Validated holds the original configuration options for the log, and some
	// of its fields parsed as a result of validating it.
	Validated *ValidatedLogConfig
	// CreateStorage instantiates a Tessera storage implementation with a signer option.
	CreateStorage func(context.Context, note.Signer) (*CTStorage, error)
	// Deadline is a timeout for Tessera requests.
	Deadline time.Duration
	// MetricFactory allows creating metrics.
	MetricFactory monitoring.MetricFactory
	// ErrorMapper converts an error from an RPC request to an HTTP status, plus
	// a boolean to indicate whether the conversion succeeded.
	ErrorMapper func(error) (int, bool)
	// RequestLog provides structured logging of CTFE requests.
	RequestLog         RequestLog
	MaskInternalErrors bool
}

// Instance is a set up log/mirror instance. It must be created with the
// SetUpInstance call.
type Instance struct {
	Handlers PathHandlers
	li       *logInfo
}

// GetPublicKey returns the public key from the instance's signer.
func (i *Instance) GetPublicKey() crypto.PublicKey {
	if i.li != nil && i.li.signer != nil {
		return i.li.signer.Public()
	}
	return nil
}

// SetUpInstance sets up a log (or log mirror) instance using the provided
// configuration, and returns an object containing a set of handlers for this
// log, and an STH getter.
func SetUpInstance(ctx context.Context, opts InstanceOptions) (*Instance, error) {
	logInfo, err := setUpLogInfo(ctx, opts)
	if err != nil {
		return nil, err
	}
	handlers := logInfo.Handlers(opts.Validated.Config.Origin)
	return &Instance{Handlers: handlers, li: logInfo}, nil
}

func setUpLogInfo(ctx context.Context, opts InstanceOptions) (*logInfo, error) {
	vCfg := opts.Validated
	cfg := vCfg.Config

	// Check config validity.
	if len(cfg.RootsPemFile) == 0 {
		return nil, errors.New("need to specify RootsPemFile")
	}

	// Load the trusted roots.
	roots := x509util.NewPEMCertPool()
	for _, pemFile := range cfg.RootsPemFile {
		if err := roots.AppendCertsFromPEMFile(pemFile); err != nil {
			return nil, fmt.Errorf("failed to read trusted roots: %v", err)
		}
	}

	var signer crypto.Signer
	var err error
	if signer, err = keys.NewSigner(ctx, vCfg.PrivKey); err != nil {
		return nil, fmt.Errorf("failed to load private key: %v", err)
	}

	// TODO(phboneff): are pub keys actually used? If not, remove
	// If a public key has been configured for a log, check that it is consistent with the private key.
	if vCfg.PubKey != nil {
		switch pub := vCfg.PubKey.(type) {
		case *ecdsa.PublicKey:
			if !pub.Equal(signer.Public()) {
				return nil, errors.New("public key is not consistent with private key")
			}
		default:
			return nil, errors.New("failed to verify consistency of public key with private key")
		}
	}

	validationOpts := CertValidationOpts{
		trustedRoots:    roots,
		rejectExpired:   cfg.RejectExpired,
		rejectUnexpired: cfg.RejectUnexpired,
		notAfterStart:   vCfg.NotAfterStart,
		notAfterLimit:   vCfg.NotAfterLimit,
		acceptOnlyCA:    cfg.AcceptOnlyCa,
		extKeyUsages:    vCfg.KeyUsages,
	}
	validationOpts.rejectExtIds, err = parseOIDs(cfg.RejectExtensions)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RejectExtensions: %v", err)
	}

	logID, err := GetCTLogID(signer.Public())
	if err != nil {
		return nil, fmt.Errorf("failed to get logID for signing: %v", err)
	}
	timeSource := new(SystemTimeSource)
	ctSigner := NewCpSigner(signer, vCfg.Config.Origin, logID, timeSource)

	if opts.CreateStorage == nil {
		return nil, fmt.Errorf("failed to initiate storage backend: nil createStorage")
	}
	storage, err := opts.CreateStorage(ctx, ctSigner)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate storage backend: %v", err)
	}

	logInfo := newLogInfo(opts, validationOpts, signer, timeSource, storage)
	return logInfo, nil
}

func parseOIDs(oids []string) ([]asn1.ObjectIdentifier, error) {
	ret := make([]asn1.ObjectIdentifier, 0, len(oids))
	for _, s := range oids {
		bits := strings.Split(s, ".")
		var oid asn1.ObjectIdentifier
		for _, n := range bits {
			p, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			oid = append(oid, p)
		}
		ret = append(ret, oid)
	}
	return ret, nil
}
