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
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/trillian/crypto/keys"
	"github.com/google/trillian/crypto/keys/der"
	"github.com/google/trillian/crypto/keys/pem"
	"github.com/google/trillian/crypto/keys/pkcs11"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/google/trillian/monitoring/opencensus"
	"github.com/google/trillian/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	tdnote "github.com/transparency-dev/formats/note"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/modules/dedup"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/storage/bbolt"
	gcpSCTFE "github.com/transparency-dev/trillian-tessera/personalities/sctfe/storage/gcp"
	gcpTessera "github.com/transparency-dev/trillian-tessera/storage/gcp"
	"golang.org/x/mod/sumdb/note"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// Global flags that affect all log instances.
var (
	httpEndpoint       = flag.String("http_endpoint", "localhost:6962", "Endpoint for HTTP (host:port)")
	tlsCert            = flag.String("tls_certificate", "", "Path to server TLS certificate")
	tlsKey             = flag.String("tls_key", "", "Path to server TLS private key")
	metricsEndpoint    = flag.String("metrics_endpoint", "", "Endpoint for serving metrics; if left empty, metrics will be visible on --http_endpoint")
	rpcDeadline        = flag.Duration("rpc_deadline", time.Second*10, "Deadline for backend RPC requests")
	logConfig          = flag.String("log_config", "", "File holding log config in text proto format")
	maskInternalErrors = flag.Bool("mask_internal_errors", false, "Don't return error strings with Internal Server Error HTTP responses")
	tracing            = flag.Bool("tracing", false, "If true opencensus Stackdriver tracing will be enabled. See https://opencensus.io/.")
	tracingProjectID   = flag.String("tracing_project_id", "", "project ID to pass to stackdriver. Can be empty for GCP, consult docs for other platforms.")
	tracingPercent     = flag.Int("tracing_percent", 0, "Percent of requests to be traced. Zero is a special case to use the DefaultSampler")
	pkcs11ModulePath   = flag.String("pkcs11_module_path", "", "Path to the PKCS#11 module to use for keys that use the PKCS#11 interface")
	// This should be specified in the config proto, but this proto is going to go away in favour of flags, so let's put this one here directly.
	// TODO: remove comment above when the config proto has been deleted.
	dedupPath = flag.String("dedup_path", "", "Path to the deduplication database")
)

// nolint:staticcheck
func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	keys.RegisterHandler(&keyspb.PEMKeyFile{}, pem.FromProto)
	keys.RegisterHandler(&keyspb.PrivateKey{}, der.FromProto)
	keys.RegisterHandler(&keyspb.PKCS11Config{}, func(ctx context.Context, pb proto.Message) (crypto.Signer, error) {
		if cfg, ok := pb.(*keyspb.PKCS11Config); ok {
			return pkcs11.FromConfig(*pkcs11ModulePath, cfg)
		}
		return nil, fmt.Errorf("pkcs11: got %T, want *keyspb.PKCS11Config", pb)
	})

	cfgs, err := sctfe.LogConfigSetFromFile(*logConfig)
	if err != nil {
		klog.Exitf("Failed to read config: %v", err)
	}

	vCfgs, err := sctfe.ValidateLogConfigSet(cfgs)
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
	var publicKeys []crypto.PublicKey
	for _, vc := range vCfgs {
		inst, err := setupAndRegister(ctx,
			*rpcDeadline,
			vc,
			corsMux,
			*maskInternalErrors,
		)
		if err != nil {
			klog.Exitf("Failed to set up log instance for %+v: %v", cfgs, err)
		}

		// Ensure that this log does not share the same private key as any other
		// log that has already been set up and registered.
		if publicKey := inst.GetPublicKey(); publicKey != nil {
			for _, p := range publicKeys {
				switch pub := publicKey.(type) {
				case *ecdsa.PublicKey:
					if pub.Equal(p) {
						klog.Exitf("Same private key used by more than one log")
					}
				case ed25519.PublicKey:
					if pub.Equal(p) {
						klog.Exitf("Same private key used by more than one log")
					}
				case *rsa.PublicKey:
					if pub.Equal(p) {
						klog.Exitf("Same private key used by more than one log")
					}
				}
			}
			publicKeys = append(publicKeys, publicKey)
		}
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
	}

	switch vCfg.Config.StorageConfig.(type) {
	case *configpb.LogConfig_Gcp:
		klog.Info("Found GCP storage config, will set up GCP tessera storage")
		opts.CreateStorage = newGCPStorage
	default:
		return nil, fmt.Errorf("unrecognized storage config")
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

func newGCPStorage(ctx context.Context, vCfg *sctfe.ValidatedLogConfig, signer note.Signer) (*sctfe.CTStorage, error) {
	cfg := vCfg.Config.GetGcp()
	gcpCfg := gcpTessera.Config{
		ProjectID: cfg.ProjectId,
		Bucket:    cfg.Bucket,
		Spanner:   cfg.SpannerDbPath,
	}
	tesseraStorage, err := gcpTessera.New(ctx, gcpCfg, tessera.WithCheckpointSignerVerifier(signer, nil), tessera.WithCTLayout())
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize GCP Tessera storage: %v", err)
	}

	issuerStorage, err := gcpSCTFE.NewIssuerStorage(ctx, cfg.ProjectId, cfg.Bucket, "fingerprints/", "application/pkix-cert")
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize GCP issuer storage: %v", err)
	}

	dedupStorage, err := bbolt.NewStorage(*dedupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize BBolt deduplication database")
	}

	fetcher, err := gcpSCTFE.GetFetcher(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to get a log fetcher")
	}

	verifierString, err := tdnote.RFC6962VerifierString(vCfg.Config.Origin, vCfg.PubKey)
	if err != nil {
		return nil, fmt.Errorf("error creating static-ct-api checkpoint verifier string: %v", err)

	}
	verifier, err := tdnote.NewRFC6962Verifier(verifierString)
	if err != nil {
		return nil, fmt.Errorf("error creating static-ct-api checkpoint verifier: %v", err)

	}

	localDedup := dedup.NewLocalBestEffortDedup(ctx, dedupStorage, time.Second, fetcher, verifier, vCfg.Config.Origin, sctfe.DedupFromBundle)

	return sctfe.NewCTSTorage(tesseraStorage, issuerStorage, localDedup)
}
