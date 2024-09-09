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

// The ct_server binary runs the CT personality.
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/google/trillian/crypto/keys/pem"
	"github.com/google/trillian/monitoring/opencensus"
	"github.com/google/trillian/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/modules/dedup"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/storage/bbolt"
	gcpSCTFE "github.com/transparency-dev/trillian-tessera/personalities/sctfe/storage/gcp"
	gcpTessera "github.com/transparency-dev/trillian-tessera/storage/gcp"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

func init() {
	flag.Var(&notAfterStart, "not_after_start", "Start of the range of acceptable NotAfter values, inclusive. Leaving this unset implies no lower bound to the range. RFC3339 UTC format, e.g: 2024-01-02T15:04:05Z.")
	flag.Var(&notAfterLimit, "not_after_limit", "Cut off point of notAfter dates - only notAfter dates strictly *before* notAfterLimit will be accepted. Leaving this unset means no upper bound on the accepted range. RFC3339 UTC format, e.g: 2024-01-02T15:04:05Z.")
}

// Global flags that affect all log instances.
var (
	notAfterStart timestampFlag
	notAfterLimit timestampFlag

	httpEndpoint    = flag.String("http_endpoint", "localhost:6962", "Endpoint for HTTP (host:port).")
	tlsCert         = flag.String("tls_certificate", "", "Path to server TLS certificate.")
	tlsKey          = flag.String("tls_key", "", "Path to server TLS private key.")
	metricsEndpoint = flag.String("metrics_endpoint", "", "Endpoint for serving metrics; if left empty, metrics will be visible on --http_endpoint.")
	// TODO(phboneff): delete / rename
	rpcDeadline        = flag.Duration("rpc_deadline", time.Second*10, "Deadline for backend RPC requests.")
	maskInternalErrors = flag.Bool("mask_internal_errors", false, "Don't return error strings with Internal Server Error HTTP responses.")
	tracing            = flag.Bool("tracing", false, "If true opencensus Stackdriver tracing will be enabled. See https://opencensus.io/.")
	tracingProjectID   = flag.String("tracing_project_id", "", "project ID to pass to stackdriver. Can be empty for GCP, consult docs for other platforms.")
	tracingPercent     = flag.Int("tracing_percent", 0, "Percent of requests to be traced. Zero is a special case to use the DefaultSampler.")
	dedupPath          = flag.String("dedup_path", "", "Path to the deduplication database.")
	origin             = flag.String("origin", "", "Origin of the log, for checkpoints and the monitoring prefix.")
	projectID          = flag.String("project_id", "", "GCP ProjectID.")
	bucket             = flag.String("bucket", "", "Name of the bucket to store the log in.")
	spannerDB          = flag.String("spanner_db_path", "", "Spanner database path: projects/{projectId}/instances/{instanceId}/databases/{databaseId}.")
	rootsPemFile       = flag.String("roots_pem_file", "", "Path to the file containing root certificates that are acceptable to the log. The certs are served through get-roots endpoint.")
	rejectExpired      = flag.Bool("reject_expired", false, "If true then the certificate validity period will be checked against the current time during the validation of submissions. This will cause expired certificates to be rejected.")
	rejectUnexpired    = flag.Bool("reject_unexpired", false, "If true then CTFE rejects certificates that are either currently valid or not yet valid.")
	extKeyUsages       = flag.String("ext_key_usages", "", "If set, will restrict the set of such usages that the server will accept. By default all are accepted. The values specified must be ones known to the x509 package.")
	rejectExtensions   = flag.String("reject_extension", "", "A list of X.509 extension OIDs, in dotted string form (e.g. '2.3.4.5') which, if present, should cause submissions to be rejected.")
	privKey            = flag.String("private_key", "", "Path to a private key .der file. Used to sign checkpoints and SCTs.")
	privKeyPass        = flag.String("password", "", "private_key password.")
)

// nolint:staticcheck
func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	// TODO(phboneff): move to something else, like KMS
	signer, err := pem.ReadPrivateKeyFile(*privKey, *privKeyPass)
	if err != nil {
		klog.Exitf("Can't open key: %v", err)
	}

	vCfg, err := sctfe.ValidateLogConfig(*origin, *projectID, *bucket, *spannerDB, *rootsPemFile, *rejectExpired, *rejectUnexpired, *extKeyUsages, *rejectExtensions, notAfterStart.t, notAfterLimit.t, signer)
	if err != nil {
		klog.Exitf("Invalid config: %v", err)
	}

	klog.CopyStandardLogTo("WARNING")
	klog.Info("**** CT HTTP Server Starting ****")

	metricsAt := *metricsEndpoint
	if metricsAt == "" {
		metricsAt = *httpEndpoint
	}

	// Allow cross-origin requests to all handlers registered on corsMux.
	// This is safe for CT log handlers because the log is public and
	// unauthenticated so cross-site scripting attacks are not a concern.
	corsMux := http.NewServeMux()
	corsHandler := cors.AllowAll().Handler(corsMux)
	http.Handle("/", corsHandler)

	// Register handlers for all the configured logs using the correct RPC
	// client.
	// TODO(phboneff): setupAndRegister can probably be inlined / removed later
	_, err = setupAndRegister(ctx,
		*rpcDeadline,
		vCfg,
		corsMux,
		*maskInternalErrors,
	)
	if err != nil {
		klog.Exitf("Failed to set up log instance for %+v: %v", vCfg, err)
	}

	// Return a 200 on the root, for GCE default health checking :/
	corsMux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			resp.WriteHeader(http.StatusOK)
		} else {
			resp.WriteHeader(http.StatusNotFound)
		}
	})

	// Export a healthz target.
	corsMux.HandleFunc("/healthz", func(resp http.ResponseWriter, req *http.Request) {
		// TODO(al): Wire this up to tell the truth.
		if _, err := resp.Write([]byte("ok")); err != nil {
			klog.Errorf("resp.Write(): %v", err)
		}
	})

	if metricsAt != *httpEndpoint {
		// Run a separate handler for metrics.
		go func() {
			mux := http.NewServeMux()
			mux.Handle("/metrics", promhttp.Handler())
			metricsServer := http.Server{Addr: metricsAt, Handler: mux}
			err := metricsServer.ListenAndServe()
			klog.Warningf("Metrics server exited: %v", err)
		}()
	} else {
		// Handle metrics on the DefaultServeMux.
		http.Handle("/metrics", promhttp.Handler())
	}

	// If we're enabling tracing we need to use an instrumented http.Handler.
	var handler http.Handler
	if *tracing {
		handler, err = opencensus.EnableHTTPServerTracing(*tracingProjectID, *tracingPercent)
		if err != nil {
			klog.Exitf("Failed to initialize stackdriver / opencensus tracing: %v", err)
		}
	}

	// Bring up the HTTP server and serve until we get a signal not to.
	srv := http.Server{}
	if *tlsCert != "" && *tlsKey != "" {
		cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
		if err != nil {
			klog.Errorf("failed to load TLS certificate/key: %v", err)
		}
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
		}
		srv = http.Server{Addr: *httpEndpoint, Handler: handler, TLSConfig: tlsConfig}
	} else {
		srv = http.Server{Addr: *httpEndpoint, Handler: handler}
	}
	shutdownWG := new(sync.WaitGroup)
	go awaitSignal(func() {
		shutdownWG.Add(1)
		defer shutdownWG.Done()
		// Allow 60s for any pending requests to finish then terminate any stragglers
		// TODO(phboneff): maybe wait for the sequencer queue to be empty?
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		klog.Info("Shutting down HTTP server...")
		if err := srv.Shutdown(ctx); err != nil {
			klog.Errorf("srv.Shutdown(): %v", err)
		}
		klog.Info("HTTP server shutdown")
	})

	if *tlsCert != "" && *tlsKey != "" {
		err = srv.ListenAndServeTLS("", "")
	} else {
		err = srv.ListenAndServe()
	}
	if err != http.ErrServerClosed {
		klog.Warningf("Server exited: %v", err)
	}
	// Wait will only block if the function passed to awaitSignal was called,
	// in which case it'll block until the HTTP server has gracefully shutdown
	shutdownWG.Wait()
	klog.Flush()
}

// awaitSignal waits for standard termination signals, then runs the given
// function; it should be run as a separate goroutine.
func awaitSignal(doneFn func()) {
	// Arrange notification for the standard set of signals used to terminate a server
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Now block main and wait for a signal
	sig := <-sigs
	klog.Warningf("Signal received: %v", sig)
	klog.Flush()

	doneFn()
}

func setupAndRegister(ctx context.Context, deadline time.Duration, vCfg *sctfe.ValidatedLogConfig, mux *http.ServeMux, maskInternalErrors bool) (*sctfe.Instance, error) {
	opts := sctfe.InstanceOptions{
		Validated:          vCfg,
		Deadline:           deadline,
		MetricFactory:      prometheus.MetricFactory{},
		RequestLog:         new(sctfe.DefaultRequestLog),
		MaskInternalErrors: maskInternalErrors,
		CreateStorage:      newGCPStorage,
	}

	inst, err := sctfe.SetUpInstance(ctx, opts)
	if err != nil {
		return nil, err
	}
	for path, handler := range inst.Handlers {
		mux.Handle(path, handler)
	}
	return inst, nil
}

func newGCPStorage(ctx context.Context, signer note.Signer) (*sctfe.CTStorage, error) {
	gcpCfg := gcpTessera.Config{
		ProjectID: *projectID,
		Bucket:    *bucket,
		Spanner:   *spannerDB,
	}
	tesseraStorage, err := gcpTessera.New(ctx, gcpCfg, tessera.WithCheckpointSignerVerifier(signer, nil), tessera.WithCTLayout())
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize GCP Tessera storage: %v", err)
	}

	issuerStorage, err := gcpSCTFE.NewIssuerStorage(ctx, *projectID, *bucket, "fingerprints/", "application/pkix-cert")
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize GCP issuer storage: %v", err)
	}

	// TODO: replace with a global dedup storage for GCP
	beDedupStorage, err := bbolt.NewStorage(*dedupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize BBolt deduplication database: %v", err)
	}

	fetcher, err := gcpSCTFE.GetFetcher(ctx, *bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to get a log fetcher: %v", err)
	}

	go dedup.UpdateFromLog(ctx, beDedupStorage, time.Second, fetcher, sctfe.DedupFromBundle)

	return sctfe.NewCTSTorage(tesseraStorage, issuerStorage, beDedupStorage)
}

type timestampFlag struct {
	t *time.Time
}

func (t *timestampFlag) String() string {
	if t.t != nil {
		return t.t.Format(time.RFC3339)
	}
	return ""
}

func (t *timestampFlag) Set(w string) error {
	if !strings.HasSuffix(w, "Z") {
		return fmt.Errorf("timestamps MUST be in UTC, got %v", w)
	}
	tt, err := time.Parse(time.RFC3339, w)
	if err != nil {
		return fmt.Errorf("can't parse %q as RFC3339 timestamp: %v", w, err)
	}
	t.t = &tt
	return nil
}
