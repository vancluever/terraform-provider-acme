package acme

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/rainycape/memcache"
	"software.sslmate.com/src/go-pkcs12"
)

var uuidRegexp = regexp.MustCompile(`^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$`)
var certURLRegexp = regexp.MustCompile(`^https://localhost:1400[01]/certZ/[a-z0-9]+(/alternate/\d+)?$`)

func TestAccACMECertificate_basic(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		CheckDestroy:      testAccCheckACMECertificateStatus("acme_certificate.certificate", certificateStatusRevoked),
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www", "www2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckACMECertificateStatus("acme_certificate.certificate", certificateStatusValid),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_CSR(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateCSRConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www3", "www4"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_CSR_PreferredChain(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateCSRConfigWithPreferredChain(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www3", "www4"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(alternateIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_forceRenewal(t *testing.T) {
	wantEnv := os.Environ()
	var certURL string
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateForceRenewalConfig(),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						certURL = s.RootModule().Resources["acme_certificate.certificate"].Primary.Attributes["certificate_url"]
						return nil
					},
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www6", ""),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccACMECertificateForceRenewalConfig(),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						if certURL == s.Modules[0].Resources["acme_certificate.certificate"].Primary.Attributes["certificate_url"] {
							return errors.New("certificate URL did not change")
						}

						return nil
					},
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www6", ""),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccACMECertificate_wildcard(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateWildcardConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "*", ""),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_p12Password(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigP12Password("changeit"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www12", "www13"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
			{
				Config: testAccACMECertificateConfigP12Password("changeitagain"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www12", "www13"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_preCheckDelay(t *testing.T) {
	wantEnv := os.Environ()
	var step1Start, step1End, step2Start, step2End time.Time
	const delay = 15

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				PreConfig: func() { step1Start = time.Now() },
				Config:    testAccACMECertificateConfigPreCheckDelay(0),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						step1End = time.Now()
						return nil
					},
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www16", "www17"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
			{
				Config:  testAccACMECertificateConfigPreCheckDelay(0),
				Destroy: true,
			},
			{
				PreConfig: func() { step2Start = time.Now() },
				Config:    testAccACMECertificateConfigPreCheckDelay(delay),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						step2End = time.Now()
						step1Elapsed := step1End.Sub(step1Start)
						step2Elapsed := step2End.Sub(step2Start)

						// Approximate the actual delay and expect some margin of
						// error, since it's pretty much guaranteed that the
						// elapsed time is not going to be exact, to the tune of
						// seconds on part of caching/etc.
						//
						// Additionally, we need to multiply the configured delay
						// by the number of domains we're actually configuring
						// challenges for.
						const deltaThreshold = 10

						expectedDelay := delay * 2
						actualDelay := int((step2Elapsed - step1Elapsed) / time.Second)
						delayDelta := expectedDelay - actualDelay
						if delayDelta > deltaThreshold || delayDelta < -deltaThreshold {
							return fmt.Errorf(
								"delta too large between standard and pre-check delay applies; expected %ds, got approx. %ds", expectedDelay, actualDelay)
						}

						return nil
					},
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www16", "www17"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_duplicateDomain(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigDuplicateDomain(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-dupe", "test-dupe"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_preferredChain(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigPreferredChain(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-preferred", "test-preferred2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(alternateIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_http(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigHTTP(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-http", "test-http2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_httpWebroot(t *testing.T) {
	wantEnv := os.Environ()
	closeServer, serverDir, err := testAccCheckACMECertificateWebrootTestServer()
	if err != nil {
		panic(fmt.Errorf("TestAccACMECertificate_httpWebroot: %s", err))
	}
	defer closeServer()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigHTTPWebroot(serverDir),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-webroot", "test-webroot2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_httpMemcache(t *testing.T) {
	wantEnv := os.Environ()
	closeServer, err := testAccCheckACMECertificateMemcacheTestServer()
	if err != nil {
		panic(fmt.Errorf("TestAccACMECertificate_httpMemcache: %s", err))
	}
	defer closeServer()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigHTTPMemcache(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-webroot", "test-webroot2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_httpProxy(t *testing.T) {
	wantEnv := os.Environ()
	closeServer, err := testAccCheckACMECertificateProxyTestServer()
	if err != nil {
		panic(fmt.Errorf("TestAccACMECertificate_httpProxy: %s", err))
	}
	defer closeServer()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigHTTPProxy(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-proxy", "test-proxy2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_tls(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigTLS(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-tls", "test-tls2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func TestAccACMECertificate_noRevoke(t *testing.T) {
	wantEnv := os.Environ()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		CheckDestroy:      testAccCheckACMECertificateStatus("acme_certificate.certificate", certificateStatusValid),
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigNoRevoke(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("acme_certificate.certificate", "id", uuidRegexp),
					resource.TestMatchResourceAttr("acme_certificate.certificate", "certificate_url", certURLRegexp),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "test-no-revoke", "test-no-revoke2"),
					testAccCheckACMECertificateIntermediateEqual("acme_certificate.certificate", getPebbleCertificate(mainIntermediateURL)),
					testAccCheckACMECertificateStatus("acme_certificate.certificate", certificateStatusValid),
					testAccCheckEnvironNotChanged(wantEnv),
				),
			},
		},
	})
}

func testAccCheckACMECertificateValid(n, cn, san string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find ACME certificate: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ACME certificate ID not set")
		}

		cert := rs.Primary.Attributes["certificate_pem"]
		issuer := rs.Primary.Attributes["issuer_pem"]
		key := rs.Primary.Attributes["private_key_pem"]
		x509Certs, err := parsePEMBundle([]byte(cert))
		if err != nil {
			return err
		}
		x509Cert := x509Certs[0]

		issuerCerts, err := parsePEMBundle([]byte(issuer))
		if err != nil {
			return err
		}
		issuerCert := issuerCerts[0]

		// Skip the private key test if we have an empty key. This is a legit case
		// that comes up when a CSR is supplied instead of creating a cert from
		// scratch.
		if key != "" {
			privateKey, err := privateKeyFromPEM([]byte(key))
			if err != nil {
				return err
			}

			var privPub crypto.PublicKey

			switch v := privateKey.(type) {
			case *rsa.PrivateKey:
				privPub = v.Public()
			case *ecdsa.PrivateKey:
				privPub = v.Public()
			}

			if reflect.DeepEqual(x509Cert.PublicKey, privPub) != true {
				return fmt.Errorf("Public key for cert and private key don't match: %#v, %#v", x509Cert.PublicKey, privPub)
			}

			// Test PKCS12, which is only present if there's a private key.
			if err := testFindPEMInP12(
				[]byte(rs.Primary.Attributes["certificate_p12"]),
				rs.Primary.Attributes["certificate_p12_password"],
				[]byte(cert),
				[]byte(issuer),
				[]byte(key),
			); err != nil {
				return fmt.Errorf("error validating P12 certificates: %s", err)
			}
		}

		// Ensure the issuer cert is a CA cert
		if issuerCert.IsCA == false {
			return fmt.Errorf("issuer_pem is not a CA certificate")
		}

		// domains
		domain := "." + pebbleCertDomain
		expectedCN := cn + domain
		var expectedSANs []string
		if san != "" && cn != san {
			expectedSANs = []string{cn + domain, san + domain}
		} else {
			expectedSANs = []string{cn + domain}
		}

		actualCN := x509Cert.Subject.CommonName
		actualSANs := x509Cert.DNSNames

		if expectedCN != actualCN {
			return fmt.Errorf("Expected common name to be %s, got %s", expectedCN, actualCN)
		}

		if reflect.DeepEqual(expectedSANs, actualSANs) != true {
			return fmt.Errorf("Expected SANs to be %#v, got %#v", expectedSANs, actualSANs)
		}

		// Expiry
		actualNotAfter := rs.Primary.Attributes["certificate_not_after"]
		expectedNotAfter := x509Cert.NotAfter.Format(time.RFC3339)
		if expectedNotAfter != actualNotAfter {
			return fmt.Errorf("expected certificate_not_after to be %q, got %q", expectedNotAfter, actualNotAfter)
		}

		return nil
	}
}

// testFindPEMInP12 tries to find the supplied PEM blocks in the supplied
// base64-encoded P12 content.
func testFindPEMInP12(pfxB64 []byte, password string, expected ...[]byte) error {
	pfxData := make([]byte, base64.StdEncoding.DecodedLen(len(pfxB64)))
	nBytes, err := base64.StdEncoding.Decode(pfxData, pfxB64)
	if err != nil {
		return err
	}

	// TODO: fix the ToPEM deprecation notice
	actualBlocks, err := pkcs12.ToPEM(pfxData[:nBytes], password) //nolint:staticcheck
	if err != nil {
		return err
	}

	var expectedBlocks []*pem.Block
	for i, data := range expected {
		block, _ := pem.Decode(data)
		if block == nil {
			return fmt.Errorf("bad PEM data in expected block %d", i)
		}

		expectedBlocks = append(expectedBlocks, block)
	}

	for i := 0; i < len(expectedBlocks); i++ {
		expected := expectedBlocks[i]
		for _, actual := range actualBlocks {
			if reflect.DeepEqual(expected.Bytes, actual.Bytes) {
				expectedBlocks = append(expectedBlocks[:i], expectedBlocks[i+1:]...)
				i--
			}
		}
	}

	if len(expectedBlocks) > 0 {
		return fmt.Errorf(
			"not all expected blocks were found in the PFX archive (remaining: %d, %d in archive)",
			len(expectedBlocks),
			len(actualBlocks),
		)
	}

	return nil
}

func testAccCheckACMECertificateIntermediateEqual(name string, expected *x509.Certificate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Can't find ACME certificate: %s", name)
		}

		issuer := rs.Primary.Attributes["issuer_pem"]
		issuerCerts, err := parsePEMBundle([]byte(issuer))
		if err != nil {
			return err
		}
		actual := issuerCerts[0]

		if !expected.Equal(actual) {
			return fmt.Errorf("certificate mismatch: expected issuer %q, actual %q", expected.Issuer.CommonName, actual.Issuer.CommonName)
		}

		return nil
	}
}

func testAccCheckACMECertificateWebrootTestServer() (func(), string, error) {
	dir, err := os.MkdirTemp(os.TempDir(), "terraform-provider-acme-test-webroot")
	if err != nil {
		return nil, "", err
	}

	// Would use httptest here, but this ensures we don't have to mess with the
	// default listener that would normally be crated by httptest to ensure it
	// goes to the correct place, since we need to set this to 5002.
	server := &http.Server{
		Addr:    ":5002",
		Handler: http.FileServer(http.Dir(dir)),
	}
	go server.ListenAndServe()
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		server.Shutdown(ctx)
		os.RemoveAll(dir)
	}, dir, nil
}

func testAccCheckACMECertificateMemcacheTestServer() (func(), error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/acme-challenge/", func(w http.ResponseWriter, r *http.Request) {
		client, err := memcache.New(memcacheHost)
		if err != nil {
			http.Error(w, fmt.Sprintf("memcached connect: %s", err), http.StatusBadRequest)
			return
		}

		item, err := client.Get(r.URL.Path)
		if err != nil {
			http.Error(w, fmt.Sprintf("memcached get: %s", err), http.StatusBadRequest)
			return
		}

		w.Write(item.Value)
	})

	server := &http.Server{
		Addr:    ":5002",
		Handler: mux,
	}
	go server.ListenAndServe()
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		server.Shutdown(ctx)
	}, nil
}

func testAccCheckACMECertificateProxyTestServer() (func(), error) {
	target, err := url.Parse("http://localhost:5502")
	if err != nil {
		panic(err) // No real reason to return an actual error here
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	defaultDirector := proxy.Director
	proxy.Director = func(r *http.Request) {
		r.Header.Add("Forwarded", "host="+strings.Split(r.Host, ":")[0])
		defaultDirector(r)
	}

	server := &http.Server{
		Addr:    ":5002",
		Handler: proxy,
	}
	go server.ListenAndServe()
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		server.Shutdown(ctx)
	}, nil
}

func testAccCheckACMECertificateStatus(name, expected string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Can't find ACME certificate: %s", name)
		}

		certPem := rs.Primary.Attributes["certificate_pem"]
		certs, err := parsePEMBundle([]byte(certPem))
		if err != nil {
			return err
		}
		cert := certs[0]
		actual := getStatusForCertificate(cert)

		if expected != actual {
			return fmt.Errorf("subject=%s serial=%x, expected status %q, actual %q", cert.Subject, cert.SerialNumber.Int64(), expected, actual)
		}

		return nil
	}
}

func testAccCheckEnvironNotChanged(want []string) resource.TestCheckFunc {
	// Make an ignore func that allows us to ignore a few things that seem to
	// get added by TF testing after we take an environment snapshot.
	ignoreFunc := func(s string) bool {
		ignoreList := []string{
			"PLUGIN_PROTOCOL_VERSIONS",
			"CHECKPOINT_DISABLE",
		}
		for _, v := range ignoreList {
			if strings.HasPrefix(s, v+"=") {
				return true
			}
		}

		return false
	}
	return func(_ *terraform.State) error {
		if diff := cmp.Diff(want, os.Environ(), cmpopts.IgnoreSliceElements(ignoreFunc)); diff != "" {
			return fmt.Errorf("environment altered but should not have been (-want +got):\n%s", diff)
		}

		return nil
	}
}

func testAccACMECertificateConfig() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "www.${var.domain}"
  subject_alternative_names = ["www2.${var.domain}"]

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateCSRConfig() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "reg_private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.reg_private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "tls_private_key" "cert_private_key" {
  algorithm = "RSA"
}

resource "tls_cert_request" "req" {
  private_key_pem = "${tls_private_key.cert_private_key.private_key_pem}"
  dns_names       = ["www3.${var.domain}", "www4.${var.domain}"]

  subject {
    common_name  = "www3.${var.domain}"
  }
}

resource "acme_certificate" "certificate" {
  account_key_pem         = "${acme_registration.reg.account_key_pem}"
  certificate_request_pem = "${tls_cert_request.req.cert_request_pem}"

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateCSRConfigWithPreferredChain() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "reg_private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.reg_private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "tls_private_key" "cert_private_key" {
  algorithm = "RSA"
}

resource "tls_cert_request" "req" {
  private_key_pem = "${tls_private_key.cert_private_key.private_key_pem}"
  dns_names       = ["www3.${var.domain}", "www4.${var.domain}"]

  subject {
    common_name  = "www3.${var.domain}"
  }
}

resource "acme_certificate" "certificate" {
  account_key_pem         = "${acme_registration.reg.account_key_pem}"
	certificate_request_pem = "${tls_cert_request.req.cert_request_pem}"

  recursive_nameservers        = ["%s"]
	disable_complete_propagation = true
	preferred_chain = "%s"

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		getPebbleCertificateIssuer(alternateIntermediateURL),
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateForceRenewalConfig() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem    = "${acme_registration.reg.account_key_pem}"
  common_name        = "www6.${var.domain}"
  min_days_remaining = 18250

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateWildcardConfig() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem = "${acme_registration.reg.account_key_pem}"
  common_name     = "*.${var.domain}"

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateConfigP12Password(password string) string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

variable "password" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "www12.${var.domain}"
  subject_alternative_names = ["www13.${var.domain}"]
  certificate_p12_password  = "${var.password}"

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		password,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateConfigPreCheckDelay(delay int) string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "www16.${var.domain}"
  subject_alternative_names = ["www17.${var.domain}"]

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true
  pre_check_delay              = %d

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		delay,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateConfigDuplicateDomain() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-dupe.${var.domain}"
  subject_alternative_names = ["test-dupe.${var.domain}"]

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateConfigPreferredChain() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-preferred.${var.domain}"
  subject_alternative_names = ["test-preferred2.${var.domain}"]

  recursive_nameservers        = ["%s"]
  disable_complete_propagation = true
  preferred_chain = "%s"

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		getPebbleCertificateIssuer(alternateIntermediateURL),
		pebbleChallTestDNSScriptPath,
	)
}

func testAccACMECertificateConfigHTTP() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-http.${var.domain}"
  subject_alternative_names = ["test-http2.${var.domain}"]

  http_challenge {
    port = 5002
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
	)
}

func testAccACMECertificateConfigHTTPWebroot(dir string) string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-webroot.${var.domain}"
  subject_alternative_names = ["test-webroot2.${var.domain}"]

  http_webroot_challenge {
    directory = "%s"
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		dir,
	)
}

func testAccACMECertificateConfigHTTPMemcache() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-webroot.${var.domain}"
  subject_alternative_names = ["test-webroot2.${var.domain}"]

  http_memcached_challenge {
    hosts = ["%s"]
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		memcacheHost,
	)
}

func testAccACMECertificateConfigHTTPProxy() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-proxy.${var.domain}"
  subject_alternative_names = ["test-proxy2.${var.domain}"]

  http_challenge {
    port         = 5502
    proxy_header = "Forwarded"
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
	)
}

func testAccACMECertificateConfigTLS() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "test-tls.${var.domain}"
  subject_alternative_names = ["test-tls2.${var.domain}"]

  tls_challenge {
    port = 5001
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
	)
}

func testAccACMECertificateConfigNoRevoke() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

variable "email_address" {
  default = "nobody@%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.private_key.private_key_pem
  email_address   = var.email_address
}

resource "acme_certificate" "certificate" {
  account_key_pem           = acme_registration.reg.account_key_pem
  common_name               = "test-no-revoke.${var.domain}"
  subject_alternative_names = ["test-no-revoke2.${var.domain}"]

  recursive_nameservers         = ["%s"]
  disable_complete_propagation  = true
  revoke_certificate_on_destroy = false

  dns_challenge {
    provider = "exec"
    config = {
      EXEC_PATH = "%s"
    }
  }
}
`, pebbleDirBasic,
		pebbleCertDomain,
		pebbleCertDomain,
		pebbleChallTestDNSSrv,
		pebbleChallTestDNSScriptPath,
	)
}
