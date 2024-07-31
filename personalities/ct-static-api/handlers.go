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
	"context"
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/certificate-transparency-go/asn1"
	"github.com/google/certificate-transparency-go/tls"
	"github.com/google/certificate-transparency-go/trillian/util"
	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/certificate-transparency-go/x509util"
	"github.com/google/trillian/monitoring"
	"github.com/transparency-dev/trillian-tessera/ctonly"
	"k8s.io/klog/v2"

	ct "github.com/google/certificate-transparency-go"
)

const (
	// HTTP content type header
	contentTypeHeader string = "Content-Type"
	// MIME content type for JSON
	contentTypeJSON string = "application/json"
)

// EntrypointName identifies a CT entrypoint as defined in section 4 of RFC 6962.
type EntrypointName string

// Constants for entrypoint names, as exposed in statistics/logging.
const (
	AddChainName    = EntrypointName("AddChain")
	AddPreChainName = EntrypointName("AddPreChain")
)

var (
	// Metrics are all per-log (label "origin"), but may also be
	// per-entrypoint (label "ep") or per-return-code (label "rc").
	once             sync.Once
	knownLogs        monitoring.Gauge     // origin => value (always 1.0)
	maxMergeDelay    monitoring.Gauge     // origin => value
	expMergeDelay    monitoring.Gauge     // origin => value
	lastSCTTimestamp monitoring.Gauge     // origin => value
	reqsCounter      monitoring.Counter   // origin, ep => value
	rspsCounter      monitoring.Counter   // origin, ep, rc => value
	rspLatency       monitoring.Histogram // origin, ep, rc => value
)

// setupMetrics initializes all the exported metrics.
func setupMetrics(mf monitoring.MetricFactory) {
	knownLogs = mf.NewGauge("known_logs", "Set to 1 for known logs", "logid")
	maxMergeDelay = mf.NewGauge("max_merge_delay", "Maximum Merge Delay in seconds", "logid")
	expMergeDelay = mf.NewGauge("expected_merge_delay", "Expected Merge Delay in seconds", "logid")
	lastSCTTimestamp = mf.NewGauge("last_sct_timestamp", "Time of last SCT in ms since epoch", "logid")
	reqsCounter = mf.NewCounter("http_reqs", "Number of requests", "logid", "ep")
	rspsCounter = mf.NewCounter("http_rsps", "Number of responses", "logid", "ep", "rc")
	rspLatency = mf.NewHistogram("http_latency", "Latency of responses in seconds", "logid", "ep", "rc")
}

// Entrypoints is a list of entrypoint names as exposed in statistics/logging.
var Entrypoints = []EntrypointName{AddChainName, AddPreChainName}

// PathHandlers maps from a path to the relevant AppHandler instance.
type PathHandlers map[string]AppHandler

// AppHandler holds a logInfo and a handler function that uses it, and is
// an implementation of the http.Handler interface.
type AppHandler struct {
	Info    *logInfo
	Handler func(context.Context, *logInfo, http.ResponseWriter, *http.Request) (int, error)
	Name    EntrypointName
	Method  string // http.MethodGet or http.MethodPost
}

// ServeHTTP for an AppHandler invokes the underlying handler function but
// does additional common error and stats processing.
func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	label0 := a.Info.LogOrigin
	label1 := string(a.Name)
	reqsCounter.Inc(label0, label1)
	startTime := a.Info.TimeSource.Now()
	logCtx := a.Info.RequestLog.Start(r.Context())
	a.Info.RequestLog.LogOrigin(logCtx, a.Info.LogOrigin)
	defer func() {
		latency := a.Info.TimeSource.Now().Sub(startTime).Seconds()
		rspLatency.Observe(latency, label0, label1, strconv.Itoa(statusCode))
	}()
	klog.V(2).Infof("%s: request %v %q => %s", a.Info.LogOrigin, r.Method, r.URL, a.Name)
	if r.Method != a.Method {
		klog.Warningf("%s: %s wrong HTTP method: %v", a.Info.LogOrigin, a.Name, r.Method)
		a.Info.SendHTTPError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed: %s", r.Method))
		a.Info.RequestLog.Status(logCtx, http.StatusMethodNotAllowed)
		return
	}

	// For GET requests all params come as form encoded so we might as well parse them now.
	// POSTs will decode the raw request body as JSON later.
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			a.Info.SendHTTPError(w, http.StatusBadRequest, fmt.Errorf("failed to parse form data: %s", err))
			a.Info.RequestLog.Status(logCtx, http.StatusBadRequest)
			return
		}
	}

	// Many/most of the handlers forward the request on to the Log RPC server; impose a deadline
	// on this onward request.
	ctx, cancel := context.WithDeadline(logCtx, getRPCDeadlineTime(a.Info))
	defer cancel()

	var err error
	statusCode, err = a.Handler(ctx, a.Info, w, r)
	a.Info.RequestLog.Status(ctx, statusCode)
	klog.V(2).Infof("%s: %s <= st=%d", a.Info.LogOrigin, a.Name, statusCode)
	rspsCounter.Inc(label0, label1, strconv.Itoa(statusCode))
	if err != nil {
		klog.Warningf("%s: %s handler error: %v", a.Info.LogOrigin, a.Name, err)
		a.Info.SendHTTPError(w, statusCode, err)
		return
	}

	// Additional check, for consistency the handler must return an error for non-200 st
	if statusCode != http.StatusOK {
		klog.Warningf("%s: %s handler non 200 without error: %d %v", a.Info.LogOrigin, a.Name, statusCode, err)
		a.Info.SendHTTPError(w, http.StatusInternalServerError, fmt.Errorf("http handler misbehaved, st: %d", statusCode))
		return
	}
}

// CertValidationOpts contains various parameters for certificate chain validation
type CertValidationOpts struct {
	// trustedRoots is a pool of certificates that defines the roots the CT log will accept
	trustedRoots *x509util.PEMCertPool
	// currentTime is the time used for checking a certificate's validity period
	// against. If it's zero then time.Now() is used. Only for testing.
	currentTime time.Time
	// rejectExpired indicates that expired certificates will be rejected.
	rejectExpired bool
	// rejectUnexpired indicates that certificates that are currently valid or not yet valid will be rejected.
	rejectUnexpired bool
	// notAfterStart is the earliest notAfter date which will be accepted.
	// nil means no lower bound on the accepted range.
	notAfterStart *time.Time
	// notAfterLimit defines the cut off point of notAfter dates - only notAfter
	// dates strictly *before* notAfterLimit will be accepted.
	// nil means no upper bound on the accepted range.
	notAfterLimit *time.Time
	// acceptOnlyCA will reject any certificate without the CA bit set.
	acceptOnlyCA bool
	// extKeyUsages contains the list of EKUs to use during chain verification
	extKeyUsages []x509.ExtKeyUsage
	// rejectExtIds contains a list of X.509 extension IDs to reject during chain verification.
	rejectExtIds []asn1.ObjectIdentifier
}

// NewCertValidationOpts builds validation options based on parameters.
func NewCertValidationOpts(trustedRoots *x509util.PEMCertPool, currentTime time.Time, rejectExpired bool, rejectUnexpired bool, notAfterStart *time.Time, notAfterLimit *time.Time, acceptOnlyCA bool, extKeyUsages []x509.ExtKeyUsage) CertValidationOpts {
	var vOpts CertValidationOpts
	vOpts.trustedRoots = trustedRoots
	vOpts.currentTime = currentTime
	vOpts.rejectExpired = rejectExpired
	vOpts.rejectUnexpired = rejectUnexpired
	vOpts.notAfterStart = notAfterStart
	vOpts.notAfterLimit = notAfterLimit
	vOpts.acceptOnlyCA = acceptOnlyCA
	vOpts.extKeyUsages = extKeyUsages
	return vOpts
}

// logInfo holds information for a specific log instance.
type logInfo struct {
	// LogOrigin is a pre-formatted string identifying the log for diagnostics
	LogOrigin string
	// TimeSource is a util.TimeSource that can be injected for testing
	TimeSource util.TimeSource
	// RequestLog is a logger for various request / processing / response debug
	// information.
	RequestLog RequestLog

	// Instance-wide options
	instanceOpts InstanceOptions
	// validationOpts contains the certificate chain validation parameters
	validationOpts CertValidationOpts
	// storage stores log data
	storage Storage
	// signer signs objects (e.g. STHs, SCTs) for regular logs
	signer crypto.Signer
}

// newLogInfo creates a new instance of logInfo.
func newLogInfo(
	instanceOpts InstanceOptions,
	validationOpts CertValidationOpts,
	signer crypto.Signer,
	timeSource util.TimeSource,
) *logInfo {
	vCfg := instanceOpts.Validated
	cfg := vCfg.Config

	li := &logInfo{
		LogOrigin:      cfg.Origin,
		storage:        instanceOpts.Storage,
		signer:         signer,
		TimeSource:     timeSource,
		instanceOpts:   instanceOpts,
		validationOpts: validationOpts,
		RequestLog:     instanceOpts.RequestLog,
	}

	once.Do(func() { setupMetrics(instanceOpts.MetricFactory) })
	label := cfg.Origin
	knownLogs.Set(1.0, cfg.Origin)

	maxMergeDelay.Set(float64(cfg.MaxMergeDelaySec), label)
	expMergeDelay.Set(float64(cfg.ExpectedMergeDelaySec), label)

	return li
}

// Handlers returns a map from URL paths (with the given prefix) and AppHandler instances
// to handle those entrypoints.
func (li *logInfo) Handlers(prefix string) PathHandlers {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	prefix = strings.TrimRight(prefix, "/")

	// Bind the logInfo instance to give an AppHandler instance for each endpoint.
	ph := PathHandlers{
		prefix + ct.AddChainPath:    AppHandler{Info: li, Handler: addChain, Name: AddChainName, Method: http.MethodPost},
		prefix + ct.AddPreChainPath: AppHandler{Info: li, Handler: addPreChain, Name: AddPreChainName, Method: http.MethodPost},
	}

	return ph
}

// SendHTTPError generates a custom error page to give more information on why something didn't work
func (li *logInfo) SendHTTPError(w http.ResponseWriter, statusCode int, err error) {
	errorBody := http.StatusText(statusCode)
	if !li.instanceOpts.MaskInternalErrors || statusCode != http.StatusInternalServerError {
		errorBody += fmt.Sprintf("\n%v", err)
	}
	http.Error(w, errorBody, statusCode)
}

// ParseBodyAsJSONChain tries to extract cert-chain out of request.
func ParseBodyAsJSONChain(r *http.Request) (ct.AddChainRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		klog.V(1).Infof("Failed to read request body: %v", err)
		return ct.AddChainRequest{}, err
	}

	var req ct.AddChainRequest
	if err := json.Unmarshal(body, &req); err != nil {
		klog.V(1).Infof("Failed to parse request body: %v", err)
		return ct.AddChainRequest{}, err
	}

	// The cert chain is not allowed to be empty. We'll defer other validation for later
	if len(req.Chain) == 0 {
		klog.V(1).Infof("Request chain is empty: %q", body)
		return ct.AddChainRequest{}, errors.New("cert chain was empty")
	}

	return req, nil
}

// addChainInternal is called by add-chain and add-pre-chain as the logic involved in
// processing these requests is almost identical
func addChainInternal(ctx context.Context, li *logInfo, w http.ResponseWriter, r *http.Request, isPrecert bool) (int, error) {
	var method EntrypointName

	// Check the contents of the request and convert to slice of certificates.
	addChainReq, err := ParseBodyAsJSONChain(r)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("%s: failed to parse add-chain body: %s", li.LogOrigin, err)
	}
	// Log the DERs now because they might not parse as valid X.509.
	for _, der := range addChainReq.Chain {
		li.RequestLog.AddDERToChain(ctx, der)
	}
	chain, err := verifyAddChain(li, addChainReq, isPrecert)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to verify add-chain contents: %s", err)
	}
	for _, cert := range chain {
		li.RequestLog.AddCertToChain(ctx, cert)
	}
	// Get the current time in the form used throughout RFC6962, namely milliseconds since Unix
	// epoch, and use this throughout.
	timeMillis := uint64(li.TimeSource.Now().UnixNano() / millisPerNano)

	entry, err := entryFromChain(chain, isPrecert, timeMillis)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to build MerkleTreeLeaf: %s", err)
	}

	klog.V(2).Infof("%s: %s => storage.Add", li.LogOrigin, method)
	idx, err := li.storage.Add(ctx, entry)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("couldn't store the leaf")
	}

	// Always use the returned leaf as the basis for an SCT.
	var loggedLeaf ct.MerkleTreeLeaf
	leafValue := entry.MerkleTreeLeaf(idx)
	if rest, err := tls.Unmarshal(leafValue, &loggedLeaf); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to reconstruct MerkleTreeLeaf: %s", err)
	} else if len(rest) > 0 {
		return http.StatusInternalServerError, fmt.Errorf("extra data (%d bytes) on reconstructing MerkleTreeLeaf", len(rest))
	}

	// As the Log server has definitely got the Merkle tree leaf, we can
	// generate an SCT and respond with it.
	// TODO(phboneff): this should work, but double check
	sct, err := buildV1SCT(li.signer, &loggedLeaf)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to generate SCT: %s", err)
	}
	sctBytes, err := tls.Marshal(*sct)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to marshall SCT: %s", err)
	}
	// We could possibly fail to issue the SCT after this but it's v. unlikely.
	li.RequestLog.IssueSCT(ctx, sctBytes)
	err = marshalAndWriteAddChainResponse(sct, li.signer, w)
	if err != nil {
		// reason is logged and http status is already set
		return http.StatusInternalServerError, fmt.Errorf("failed to write response: %s", err)
	}
	klog.V(3).Infof("%s: %s <= SCT", li.LogOrigin, method)
	if sct.Timestamp == timeMillis {
		lastSCTTimestamp.Set(float64(sct.Timestamp), li.LogOrigin)
	}

	return http.StatusOK, nil
}

func addChain(ctx context.Context, li *logInfo, w http.ResponseWriter, r *http.Request) (int, error) {
	return addChainInternal(ctx, li, w, r, false)
}

func addPreChain(ctx context.Context, li *logInfo, w http.ResponseWriter, r *http.Request) (int, error) {
	return addChainInternal(ctx, li, w, r, true)
}

// getRPCDeadlineTime calculates the future time an RPC should expire based on our config
func getRPCDeadlineTime(li *logInfo) time.Time {
	return li.TimeSource.Now().Add(li.instanceOpts.Deadline)
}

// verifyAddChain is used by add-chain and add-pre-chain. It does the checks that the supplied
// cert is of the correct type and chains to a trusted root.
func verifyAddChain(li *logInfo, req ct.AddChainRequest, expectingPrecert bool) ([]*x509.Certificate, error) {
	// We already checked that the chain is not empty so can move on to verification
	validPath, err := ValidateChain(req.Chain, li.validationOpts)
	if err != nil {
		// We rejected it because the cert failed checks or we could not find a path to a root etc.
		// Lots of possible causes for errors
		return nil, fmt.Errorf("chain failed to verify: %s", err)
	}

	isPrecert, err := IsPrecertificate(validPath[0])
	if err != nil {
		return nil, fmt.Errorf("precert test failed: %s", err)
	}

	// The type of the leaf must match the one the handler expects
	if isPrecert != expectingPrecert {
		if expectingPrecert {
			klog.Warningf("%s: Cert (or precert with invalid CT ext) submitted as precert chain: %q", li.LogOrigin, req.Chain)
		} else {
			klog.Warningf("%s: Precert (or cert with invalid CT ext) submitted as cert chain: %q", li.LogOrigin, req.Chain)
		}
		return nil, fmt.Errorf("cert / precert mismatch: %T", expectingPrecert)
	}

	return validPath, nil
}

// marshalAndWriteAddChainResponse is used by add-chain and add-pre-chain to create and write
// the JSON response to the client
func marshalAndWriteAddChainResponse(sct *ct.SignedCertificateTimestamp, signer crypto.Signer, w http.ResponseWriter) error {
	logID, err := GetCTLogID(signer.Public())
	if err != nil {
		return fmt.Errorf("failed to marshal logID: %s", err)
	}
	sig, err := tls.Marshal(sct.Signature)
	if err != nil {
		return fmt.Errorf("failed to marshal signature: %s", err)
	}

	rsp := ct.AddChainResponse{
		SCTVersion: sct.SCTVersion,
		Timestamp:  sct.Timestamp,
		ID:         logID[:],
		Extensions: base64.StdEncoding.EncodeToString(sct.Extensions),
		Signature:  sig,
	}

	w.Header().Set(contentTypeHeader, contentTypeJSON)
	jsonData, err := json.Marshal(&rsp)
	if err != nil {
		return fmt.Errorf("failed to marshal add-chain: %s", err)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write add-chain resp: %s", err)
	}

	return nil
}

// entryFromChain generates an Entry from a chain and timestamp.
// copied from certificate-transparency-go/serialization.go
// TODO(phboneff): move in a different file maybe?
func entryFromChain(chain []*x509.Certificate, isPrecert bool, timestamp uint64) (*ctonly.Entry, error) {
	leaf := ctonly.Entry{
		IsPrecert: isPrecert,
		Timestamp: timestamp,
	}
	if !isPrecert {
		leaf.Certificate = chain[0].Raw
		return &leaf, nil
	}

	// Pre-certs are more complicated. First, parse the leaf pre-cert and its
	// putative issuer.
	if len(chain) < 2 {
		return nil, fmt.Errorf("no issuer cert available for precert leaf building")
	}
	issuer := chain[1]
	cert := chain[0]

	var preIssuer *x509.Certificate
	if IsPreIssuer(issuer) {
		// Replace the cert's issuance information with details from the pre-issuer.
		preIssuer = issuer

		// The issuer of the pre-cert is not going to be the issuer of the final
		// cert.  Change to use the final issuer's key hash.
		if len(chain) < 3 {
			return nil, fmt.Errorf("no issuer cert available for pre-issuer")
		}
		issuer = chain[2]
	}

	// Next, post-process the DER-encoded TBSCertificate, to remove the CT poison
	// extension and possibly update the issuer field.
	defangedTBS, err := x509.BuildPrecertTBS(cert.RawTBSCertificate, preIssuer)
	if err != nil {
		return nil, fmt.Errorf("failed to remove poison extension: %v", err)
	}

	leaf.Precertificate = cert.Raw
	leaf.PrecertSigningCert = issuer.Raw
	leaf.Certificate = defangedTBS

	issuerKeyHash := sha256.Sum256(issuer.RawSubjectPublicKeyInfo)
	leaf.IssuerKeyHash = issuerKeyHash[:]
	return &leaf, nil
}

// IsPreIssuer indicates whether a certificate is a pre-cert issuer with the specific
// certificate transparency extended key usage.
// copied form certificate-transparency-go/serialization.go
func IsPreIssuer(issuer *x509.Certificate) bool {
	for _, eku := range issuer.ExtKeyUsage {
		if eku == x509.ExtKeyUsageCertificateTransparency {
			return true
		}
	}
	return false
}
