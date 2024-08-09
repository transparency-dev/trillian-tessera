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
	"bufio"
	"bytes"
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/certificate-transparency-go/x509util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/trillian/monitoring"
	"github.com/transparency-dev/trillian-tessera/ctonly"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/configpb"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/mockstorage"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/testdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	ct "github.com/google/certificate-transparency-go"
)

// Arbitrary time for use in tests
var fakeTime = time.Date(2016, 7, 22, 11, 01, 13, 0, time.UTC)
var fakeTimeMillis = uint64(fakeTime.UnixNano() / millisPerNano)

// The deadline should be the above bumped by 500ms
var fakeDeadlineTime = time.Date(2016, 7, 22, 11, 01, 13, 500*1000*1000, time.UTC)
var fakeTimeSource = NewFixedTimeSource(fakeTime)

type handlerTestInfo struct {
	mockCtrl *gomock.Controller
	roots    *x509util.PEMCertPool
	storage  *mockstorage.MockStorage
	li       *logInfo
}

// setupTest creates mock objects and contexts.  Caller should invoke info.mockCtrl.Finish().
func setupTest(t *testing.T, pemRoots []string, signer crypto.Signer) handlerTestInfo {
	t.Helper()
	info := handlerTestInfo{
		mockCtrl: gomock.NewController(t),
		roots:    x509util.NewPEMCertPool(),
	}

	info.storage = mockstorage.NewMockStorage(info.mockCtrl)
	vOpts := CertValidationOpts{
		trustedRoots:  info.roots,
		rejectExpired: false,
	}

	cfg := &configpb.LogConfig{Origin: "example.com"}
	vCfg := &ValidatedLogConfig{Config: cfg}
	iOpts := InstanceOptions{Validated: vCfg, Deadline: time.Millisecond * 500, MetricFactory: monitoring.InertMetricFactory{}, RequestLog: new(DefaultRequestLog)}
	info.li = newLogInfo(iOpts, vOpts, signer, fakeTimeSource, info.storage)

	for _, pemRoot := range pemRoots {
		if !info.roots.AppendCertsFromPEM([]byte(pemRoot)) {
			klog.Fatal("failed to load cert pool")
		}
	}

	return info
}

func (info handlerTestInfo) postHandlers() map[string]AppHandler {
	return map[string]AppHandler{
		"add-chain":     {Info: info.li, Handler: addChain, Name: "AddChain", Method: http.MethodPost},
		"add-pre-chain": {Info: info.li, Handler: addPreChain, Name: "AddPreChain", Method: http.MethodPost},
	}
}

func TestPostHandlersRejectGet(t *testing.T) {
	info := setupTest(t, []string{testdata.FakeCACertPEM}, nil)
	defer info.mockCtrl.Finish()

	// Anything in the post handler list should reject GET
	for path, handler := range info.postHandlers() {
		t.Run(path, func(t *testing.T) {
			s := httptest.NewServer(handler)
			defer s.Close()

			resp, err := http.Get(s.URL + "/ct/v1/" + path)
			if err != nil {
				t.Fatalf("http.Get(%s)=(_,%q); want (_,nil)", path, err)
			}
			if got, want := resp.StatusCode, http.StatusMethodNotAllowed; got != want {
				t.Errorf("http.Get(%s)=(%d,nil); want (%d,nil)", path, got, want)
			}
		})
	}
}

func TestPostHandlersFailure(t *testing.T) {
	var tests = []struct {
		descr string
		body  io.Reader
		want  int
	}{
		{"nil", nil, http.StatusBadRequest},
		{"''", strings.NewReader(""), http.StatusBadRequest},
		{"malformed-json", strings.NewReader("{ !$%^& not valid json "), http.StatusBadRequest},
		{"empty-chain", strings.NewReader(`{ "chain": [] }`), http.StatusBadRequest},
		{"wrong-chain", strings.NewReader(`{ "chain": [ "test" ] }`), http.StatusBadRequest},
	}

	info := setupTest(t, []string{testdata.FakeCACertPEM}, nil)
	defer info.mockCtrl.Finish()
	for path, handler := range info.postHandlers() {
		t.Run(path, func(t *testing.T) {
			s := httptest.NewServer(handler)

			for _, test := range tests {
				resp, err := http.Post(s.URL+"/ct/v1/"+path, "application/json", test.body)
				if err != nil {
					t.Errorf("http.Post(%s,%s)=(_,%q); want (_,nil)", path, test.descr, err)
					continue
				}
				if resp.StatusCode != test.want {
					t.Errorf("http.Post(%s,%s)=(%d,nil); want (%d,nil)", path, test.descr, resp.StatusCode, test.want)
				}
			}
		})
	}
}

func TestHandlers(t *testing.T) {
	path := "/test-prefix/ct/v1/add-chain"
	info := setupTest(t, nil, nil)
	defer info.mockCtrl.Finish()
	for _, test := range []string{
		"/test-prefix/",
		"test-prefix/",
		"/test-prefix",
		"test-prefix",
	} {
		t.Run(test, func(t *testing.T) {
			handlers := info.li.Handlers(test)
			if h, ok := handlers[path]; !ok {
				t.Errorf("Handlers(%s)[%q]=%+v; want _", test, path, h)
			} else if h.Name != "AddChain" {
				t.Errorf("Handlers(%s)[%q].Name=%q; want 'AddChain'", test, path, h.Name)
			}
			// Check each entrypoint has a handler
			if got, want := len(handlers), len(Entrypoints); got != want {
				t.Fatalf("len(Handlers(%s))=%d; want %d", test, got, want)
			}

			// We want to see the same set of handler names that we think we registered.
			var hNames []EntrypointName
			for _, v := range handlers {
				hNames = append(hNames, v.Name)
			}

			if !cmp.Equal(Entrypoints, hNames, cmpopts.SortSlices(func(n1, n2 EntrypointName) bool {
				return n1 < n2
			})) {
				t.Errorf("Handler names mismatch got: %v, want: %v", hNames, Entrypoints)
			}
		})
	}
}

func parseChain(t *testing.T, isPrecert bool, pemChain []string, root *x509.Certificate) (*ctonly.Entry, []*x509.Certificate) {
	t.Helper()
	pool := loadCertsIntoPoolOrDie(t, pemChain)
	leafChain := pool.RawCertificates()
	if !leafChain[len(leafChain)-1].Equal(root) {
		// The submitted chain may not include a root, but the generated LogLeaf will
		fullChain := make([]*x509.Certificate, len(leafChain)+1)
		copy(fullChain, leafChain)
		fullChain[len(leafChain)] = root
		leafChain = fullChain
	}
	entry, err := entryFromChain(leafChain, isPrecert, fakeTimeMillis)
	if err != nil {
		t.Fatalf("failed to create entry")
	}
	return entry, leafChain
}

func TestAddChainWhitespace(t *testing.T) {
	signer, err := setupSigner(fakeSignature)
	if err != nil {
		t.Fatalf("Failed to create test signer: %v", err)
	}

	info := setupTest(t, []string{testdata.FakeCACertPEM}, signer)
	defer info.mockCtrl.Finish()

	// Throughout we use variants of a hard-coded POST body derived from a chain of:
	pemChain := []string{testdata.LeafSignedByFakeIntermediateCertPEM, testdata.FakeIntermediateCertPEM}

	// Break the JSON into chunks:
	intro := "{\"chain\""
	// followed by colon then the first line of the PEM file
	chunk1a := "[\"MIIH6DCCBtCgAwIBAgIIQoIqW4Zvv+swDQYJKoZIhvcNAQELBQAwcjELMAkGA1UE"
	// straight into rest of first entry
	chunk1b := "BhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYDVQQKDAZHb29nbGUxDDAKBgNVBAsMA0VuZzEiMCAGA1UEAwwZRmFrZUludGVybWVkaWF0ZUF1dGhvcml0eTAeFw0xNjA1MTMxNDI2NDRaFw0xOTA3MTIxNDI2NDRaMIIBWDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDU1vdW50YWluIFZpZXcxEzARBgNVBAoMCkdvb2dsZSBJbmMxFTATBgNVBAMMDCouZ29vZ2xlLmNvbTGBwzCBwAYDVQQEDIG4UkZDNTI4MCBzNC4yLjEuOSAnVGhlIHBhdGhMZW5Db25zdHJhaW50IGZpZWxkIC4uLiBnaXZlcyB0aGUgbWF4aW11bSBudW1iZXIgb2Ygbm9uLXNlbGYtaXNzdWVkIGludGVybWVkaWF0ZSBjZXJ0aWZpY2F0ZXMgdGhhdCBtYXkgZm9sbG93IHRoaXMgY2VydGlmaWNhdGUgaW4gYSB2YWxpZCBjZXJ0aWZpY2F0aW9uIHBhdGguJzEqMCgGA1UEKgwhSW50ZXJtZWRpYXRlIENBIGNlcnQgdXNlZCB0byBzaWduMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAExAk5hPUVjRJUsgKc+QHibTVH1A3QEWFmCTUdyxIUlbI//zW9Io5N/DhQLSLWmB7KoCOvpJZ+MtGCXzFX+yj/N6OCBGMwggRfMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjCCA0IGA1UdEQSCAzkwggM1ggwqLmdvb2dsZS5jb22CDSouYW5kcm9pZC5jb22CFiouYXBwZW5naW5lLmdvb2dsZS5jb22CEiouY2xvdWQuZ29vZ2xlLmNvbYIWKi5nb29nbGUtYW5hbHl0aWNzLmNvbYILKi5nb29nbGUuY2GCCyouZ29vZ2xlLmNsgg4qLmdvb2dsZS5jby5pboIOKi5nb29nbGUuY28uanCCDiouZ29vZ2xlLmNvLnVrgg8qLmdvb2dsZS5jb20uYXKCDyouZ29vZ2xlLmNvbS5hdYIPKi5nb29nbGUuY29tLmJygg8qLmdvb2dsZS5jb20uY2+CDyouZ29vZ2xlLmNvbS5teIIPKi5nb29nbGUuY29tLnRygg8qLmdvb2dsZS5jb20udm6CCyouZ29vZ2xlLmRlggsqLmdvb2dsZS5lc4ILKi5nb29nbGUuZnKCCyouZ29vZ2xlLmh1ggsqLmdvb2dsZS5pdIILKi5nb29nbGUubmyCCyouZ29vZ2xlLnBsggsqLmdvb2dsZS5wdIISKi5nb29nbGVhZGFwaXMuY29tgg8qLmdvb2dsZWFwaXMuY26CFCouZ29vZ2xlY29tbWVyY2UuY29tghEqLmdvb2dsZXZpZGVvLmNvbYIMKi5nc3RhdGljLmNugg0qLmdzdGF0aWMuY29tggoqLmd2dDEuY29tggoqLmd2dDIuY29tghQqLm1ldHJpYy5nc3RhdGljLmNvbYIMKi51cmNoaW4uY29tghAqLnVybC5nb29nbGUuY29tghYqLnlvdXR1YmUtbm9jb29raWUuY29tgg0qLnlvdXR1YmUuY29tghYqLnlvdXR1YmVlZHVjYXRpb24uY29tggsqLnl0aW1nLmNvbYIaYW5kcm9pZC5jbGllbnRzLmdvb2dsZS5jb22CC2FuZHJvaWQuY29tggRnLmNvggZnb28uZ2yCFGdvb2dsZS1hbmFseXRpY3MuY29tggpnb29nbGUuY29tghJnb29nbGVjb21tZXJjZS5jb22CCnVyY2hpbi5jb22CCHlvdXR1LmJlggt5b3V0dWJlLmNvbYIUeW91dHViZWVkdWNhdGlvbi5jb20wDAYDVR0PBAUDAweAADBoBggrBgEFBQcBAQRcMFowKwYIKwYBBQUHMAKGH2h0dHA6Ly9wa2kuZ29vZ2xlLmNvbS9HSUFHMi5jcnQwKwYIKwYBBQUHMAGGH2h0dHA6Ly9jbGllbnRzMS5nb29nbGUuY29tL29jc3AwHQYDVR0OBBYEFNv0bmPu4ty+vzhgT5gx0GRE8WPYMAwGA1UdEwEB/wQCMAAwIQYDVR0gBBowGDAMBgorBgEEAdZ5AgUBMAgGBmeBDAECAjAwBgNVHR8EKTAnMCWgI6Ahhh9odHRwOi8vcGtpLmdvb2dsZS5jb20vR0lBRzIuY3JsMA0GCSqGSIb3DQEBCwUAA4IBAQAOpm95fThLYPDBdpxOkvUkzhI0cpSVjc8cDNZ4a+5mK1A2Inq+/yLH3ZMsQIMvoDcpj7uYIr+Oxmy0i4/pHg+9it/f9cmqeawA5sqmGnSOZ/lfCYI8+bRbMIULrijCuJwjfGpZZsqOvSBuIOSzRvgGVplcs0dituT2khCFrkblwa/BqIqztvP7LuEmVpjkqt4pC3HvD0XUxs5PIdZZGInfeqymk5feReWHBuPHpPIUObKxmQt+hcw6YsHE+0B84Xtx9BMe4qqUfrqmtWXn9unBwxqSYsCqxHQpQ+70pmuBxlB9s6LStIzE9syaDmUyjxRljKAwINV6z0j7hKQ6MPpE\""
	// followed by comma then
	chunk2 := "\"MIIDnTCCAoWgAwIBAgIIQoIqW4Zvv+swDQYJKoZIhvcNAQELBQAwcTELMAkGA1UEBhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYDVQQKDAZHb29nbGUxDDAKBgNVBAsMA0VuZzEhMB8GA1UEAwwYRmFrZUNlcnRpZmljYXRlQXV0aG9yaXR5MB4XDTE2MDUxMzE0MjY0NFoXDTE5MDcxMjE0MjY0NFowcjELMAkGA1UEBhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYDVQQKDAZHb29nbGUxDDAKBgNVBAsMA0VuZzEiMCAGA1UEAwwZRmFrZUludGVybWVkaWF0ZUF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMqkDHpt6SYi1GcZyClAxr3LRDnn+oQBHbMEFUg3+lXVmEsq/xQO1s4naynV6I05676XvlMh0qPyJ+9GaBxvhHeFtGh4etQ9UEmJj55rSs50wA/IaDh+roKukQxthyTESPPgjqg+DPjh6H+h3Sn00Os6sjh3DxpOphTEsdtb7fmk8J0e2KjQQCjW/GlECzc359b9KbBwNkcAiYFayVHPLaCAdvzYVyiHgXHkEEs5FlHyhe2gNEG/81Io8c3E3DH5JhT9tmVRL3bpgpT8Kr4aoFhU2LXe45YIB1A9DjUm5TrHZ+iNtvE0YfYMR9L9C1HPppmX1CahEhTdog7laE1198UCAwEAAaM4MDYwDwYDVR0jBAgwBoAEAQIDBDASBgNVHRMBAf8ECDAGAQH/AgEAMA8GA1UdDwEB/wQFAwMH/4AwDQYJKoZIhvcNAQELBQADggEBAAHiOgwAvEzhrNMQVAz8a+SsyMIABXQ5P8WbJeHjkIipE4+5ZpkrZVXq9p8wOdkYnOHx4WNi9PVGQbLG9Iufh9fpk8cyyRWDi+V20/CNNtawMq3ClV3dWC98Tj4WX/BXDCeY2jK4jYGV+ds43HYV0ToBmvvrccq/U7zYMGFcQiKBClz5bTE+GMvrZWcO5A/Lh38i2YSF1i8SfDVnAOBlAgZmllcheHpGsWfSnduIllUvTsRvEIsaaqfVLl5QpRXBOq8tbjK85/2g6ear1oxPhJ1w9hds+WTFXkmHkWvKJebY13t3OfSjAyhaRSt8hdzDzHTFwjPjHT8h6dU7/hMdkUg=\""
	epilog := "]}\n"

	req, leafChain := parseChain(t, false, pemChain, info.roots.RawCertificates()[0])
	rsp := uint64(0)

	var tests = []struct {
		descr string
		body  string
		want  int
	}{
		{
			descr: "valid",
			body:  intro + ":" + chunk1a + chunk1b + "," + chunk2 + epilog,
			want:  http.StatusOK,
		},
		{
			descr: "valid-space-between",
			body:  intro + " : " + chunk1a + chunk1b + " , " + chunk2 + epilog,
			want:  http.StatusOK,
		},
		{
			descr: "valid-newline-between",
			body:  intro + " : " + chunk1a + chunk1b + ",\n" + chunk2 + epilog,
			want:  http.StatusOK,
		},
		{
			descr: "invalid-raw-newline-in-string",
			body:  intro + ":" + chunk1a + "\n" + chunk1b + "," + chunk2 + epilog,
			want:  http.StatusBadRequest,
		},
		{
			descr: "valid-escaped-newline-in-string",
			body:  intro + ":" + chunk1a + "\\n" + chunk1b + "," + chunk2 + epilog,
			want:  http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {
			if test.want == http.StatusOK {
				info.storage.EXPECT().AddIssuerChain(deadlineMatcher(), cmpMatcher{leafChain[1:]}).Return(nil)
				info.storage.EXPECT().Add(deadlineMatcher(), cmpMatcher{req}).Return(rsp, nil)
			}

			recorder := httptest.NewRecorder()
			handler := AppHandler{Info: info.li, Handler: addChain, Name: "AddChain", Method: http.MethodPost}
			req, err := http.NewRequest(http.MethodPost, "http://example.com/ct/v1/add-chain", strings.NewReader(test.body))
			if err != nil {
				t.Fatalf("Failed to create POST request: %v", err)
			}
			handler.ServeHTTP(recorder, req)

			if recorder.Code != test.want {
				t.Fatalf("addChain()=%d (body:%v); want %dv", recorder.Code, recorder.Body, test.want)
			}
		})
	}
}

func TestAddChain(t *testing.T) {
	var tests = []struct {
		descr string
		chain []string
		// TODO(phboneff): can this be removed?
		toSign string // hex-encoded
		want   int
		err    error
	}{
		{
			descr: "leaf-only",
			chain: []string{testdata.LeafSignedByFakeIntermediateCertPEM},
			want:  http.StatusBadRequest,
		},
		{
			descr: "wrong-entry-type",
			chain: []string{testdata.PrecertPEMValid},
			want:  http.StatusBadRequest,
		},
		{
			descr:  "backend-storage-fail",
			chain:  []string{testdata.LeafSignedByFakeIntermediateCertPEM, testdata.FakeIntermediateCertPEM},
			toSign: "1337d72a403b6539f58896decba416d5d4b3603bfa03e1f94bb9b4e898af897d",
			want:   http.StatusInternalServerError,
			err:    status.Errorf(codes.Internal, "error"),
		},
		{
			descr:  "success-without-root",
			chain:  []string{testdata.LeafSignedByFakeIntermediateCertPEM, testdata.FakeIntermediateCertPEM},
			toSign: "1337d72a403b6539f58896decba416d5d4b3603bfa03e1f94bb9b4e898af897d",
			want:   http.StatusOK,
		},
		{
			descr:  "success",
			chain:  []string{testdata.LeafSignedByFakeIntermediateCertPEM, testdata.FakeIntermediateCertPEM, testdata.FakeCACertPEM},
			toSign: "1337d72a403b6539f58896decba416d5d4b3603bfa03e1f94bb9b4e898af897d",
			want:   http.StatusOK,
		},
	}

	signer, err := setupSigner(fakeSignature)
	if err != nil {
		t.Fatalf("Failed to create test signer: %v", err)
	}

	info := setupTest(t, []string{testdata.FakeCACertPEM}, signer)
	defer info.mockCtrl.Finish()

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {
			pool := loadCertsIntoPoolOrDie(t, test.chain)
			chain := createJSONChain(t, *pool)
			if len(test.toSign) > 0 {
				req, leafChain := parseChain(t, false, test.chain, info.roots.RawCertificates()[0])
				rsp := uint64(0)
				info.storage.EXPECT().AddIssuerChain(deadlineMatcher(), cmpMatcher{leafChain[1:]}).Return(nil)
				info.storage.EXPECT().Add(deadlineMatcher(), cmpMatcher{req}).Return(rsp, test.err)
			}

			recorder := makeAddChainRequest(t, info.li, chain)
			if recorder.Code != test.want {
				t.Fatalf("addChain()=%d (body:%v); want %dv", recorder.Code, recorder.Body, test.want)
			}
			if test.want == http.StatusOK {
				var resp ct.AddChainResponse
				if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
					t.Fatalf("json.Decode(%s)=%v; want nil", recorder.Body.Bytes(), err)
				}

				if got, want := ct.Version(resp.SCTVersion), ct.V1; got != want {
					t.Errorf("resp.SCTVersion=%v; want %v", got, want)
				}
				if got, want := resp.ID, demoLogID[:]; !bytes.Equal(got, want) {
					t.Errorf("resp.ID=%v; want %v", got, want)
				}
				if got, want := resp.Timestamp, uint64(1469185273000); got != want {
					t.Errorf("resp.Timestamp=%d; want %d", got, want)
				}
				if got, want := hex.EncodeToString(resp.Signature), "040300067369676e6564"; got != want {
					t.Errorf("resp.Signature=%s; want %s", got, want)
				}
				// TODO(phboneff): check that the index is in the SCT
			}
		})
	}
}

func TestAddPrechain(t *testing.T) {
	var tests = []struct {
		descr  string
		chain  []string
		root   string
		toSign string // hex-encoded
		err    error
		want   int
	}{
		{
			descr: "leaf-signed-by-different",
			chain: []string{testdata.PrecertPEMValid, testdata.FakeIntermediateCertPEM},
			want:  http.StatusBadRequest,
		},
		{
			descr: "wrong-entry-type",
			chain: []string{testdata.TestCertPEM},
			want:  http.StatusBadRequest,
		},
		{
			descr:  "backend-storage-fail",
			chain:  []string{testdata.PrecertPEMValid, testdata.CACertPEM},
			toSign: "92ecae1a2dc67a6c5f9c96fa5cab4c2faf27c48505b696dad926f161b0ca675a",
			err:    status.Errorf(codes.Internal, "error"),
			want:   http.StatusInternalServerError,
		},
		{
			descr:  "success",
			chain:  []string{testdata.PrecertPEMValid, testdata.CACertPEM},
			toSign: "92ecae1a2dc67a6c5f9c96fa5cab4c2faf27c48505b696dad926f161b0ca675a",
			want:   http.StatusOK,
		},
		{
			descr:  "success-without-root",
			chain:  []string{testdata.PrecertPEMValid},
			toSign: "92ecae1a2dc67a6c5f9c96fa5cab4c2faf27c48505b696dad926f161b0ca675a",
			want:   http.StatusOK,
		},
		// TODO(phboneff): add a test with an intermediate
		// TODO(phboneff): add a test with a pre-issuer intermediate cert
	}

	signer, err := setupSigner(fakeSignature)
	if err != nil {
		t.Fatalf("Failed to create test signer: %v", err)
	}

	info := setupTest(t, []string{testdata.CACertPEM}, signer)
	defer info.mockCtrl.Finish()

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {
			pool := loadCertsIntoPoolOrDie(t, test.chain)
			chain := createJSONChain(t, *pool)
			if len(test.toSign) > 0 {
				req, leafChain := parseChain(t, true, test.chain, info.roots.RawCertificates()[0])
				rsp := uint64(0)
				info.storage.EXPECT().AddIssuerChain(deadlineMatcher(), cmpMatcher{leafChain[1:]}).Return(nil)
				info.storage.EXPECT().Add(deadlineMatcher(), cmpMatcher{req}).Return(rsp, test.err)
			}

			recorder := makeAddPrechainRequest(t, info.li, chain)
			if recorder.Code != test.want {
				t.Fatalf("addPrechain()=%d (body:%v); want %d", recorder.Code, recorder.Body, test.want)
			}
			if test.want == http.StatusOK {
				var resp ct.AddChainResponse
				if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
					t.Fatalf("json.Decode(%s)=%v; want nil", recorder.Body.Bytes(), err)
				}

				if got, want := ct.Version(resp.SCTVersion), ct.V1; got != want {
					t.Errorf("resp.SCTVersion=%v; want %v", got, want)
				}
				if got, want := resp.ID, demoLogID[:]; !bytes.Equal(got, want) {
					t.Errorf("resp.ID=%x; want %x", got, want)
				}
				if got, want := resp.Timestamp, uint64(1469185273000); got != want {
					t.Errorf("resp.Timestamp=%d; want %d", got, want)
				}
				if got, want := hex.EncodeToString(resp.Signature), "040300067369676e6564"; got != want {
					t.Errorf("resp.Signature=%s; want %s", got, want)
				}
			}
		})
	}
}

func createJSONChain(t *testing.T, p x509util.PEMCertPool) io.Reader {
	t.Helper()
	var req ct.AddChainRequest
	for _, rawCert := range p.RawCertificates() {
		req.Chain = append(req.Chain, rawCert.Raw)
	}

	var buffer bytes.Buffer
	// It's tempting to avoid creating and flushing the intermediate writer but it doesn't work
	writer := bufio.NewWriter(&buffer)
	err := json.NewEncoder(writer).Encode(&req)
	if err := writer.Flush(); err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Fatalf("Failed to create test json: %v", err)
	}

	return bufio.NewReader(&buffer)
}

type dlMatcher struct {
}

func deadlineMatcher() gomock.Matcher {
	return dlMatcher{}
}

func (d dlMatcher) Matches(x interface{}) bool {
	ctx, ok := x.(context.Context)
	if !ok {
		return false
	}

	deadlineTime, ok := ctx.Deadline()
	if !ok {
		return false // we never make RPC calls without a deadline set
	}

	return deadlineTime == fakeDeadlineTime
}

func (d dlMatcher) String() string {
	return fmt.Sprintf("deadline is %v", fakeDeadlineTime)
}

func makeAddPrechainRequest(t *testing.T, li *logInfo, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()
	handler := AppHandler{Info: li, Handler: addPreChain, Name: "AddPreChain", Method: http.MethodPost}
	return makeAddChainRequestInternal(t, handler, "add-pre-chain", body)
}

func makeAddChainRequest(t *testing.T, li *logInfo, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()
	handler := AppHandler{Info: li, Handler: addChain, Name: "AddChain", Method: http.MethodPost}
	return makeAddChainRequestInternal(t, handler, "add-chain", body)
}

func makeAddChainRequestInternal(t *testing.T, handler AppHandler, path string, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://example.com/ct/v1/%s", path), body)
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	return w
}

func loadCertsIntoPoolOrDie(t *testing.T, certs []string) *x509util.PEMCertPool {
	t.Helper()
	pool := x509util.NewPEMCertPool()
	for _, cert := range certs {
		if !pool.AppendCertsFromPEM([]byte(cert)) {
			t.Fatalf("couldn't parse test certs: %v", certs)
		}
	}
	return pool
}

// cmpMatcher is a custom gomock.Matcher that uses cmp.Equal combined with a
// cmp.Comparer that knows how to properly compare proto.Message types.
type cmpMatcher struct{ want interface{} }

func (m cmpMatcher) Matches(got interface{}) bool {
	return cmp.Equal(got, m.want, cmp.Comparer(proto.Equal))
}
func (m cmpMatcher) String() string {
	return fmt.Sprintf("equals %v", m.want)
}
