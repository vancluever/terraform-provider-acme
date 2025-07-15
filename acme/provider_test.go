package acme

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vancluever/terraform-provider-acme/v2/acme/dnsplugin"
)

func testAccProviderAcme() *schema.Provider {
	return Provider()
}

func testAccProviderAcmeConfig(serverUrl string) *Config {
	return &Config{
		ServerURL: serverUrl,
	}
}

var testAccProviders = map[string]func() (*schema.Provider, error){
	"acme": func() (*schema.Provider, error) {
		return testAccProviderAcme(), nil
	},
}

// Path to the pebble CA cert list, from GOPATH
const pebbleCACerts = "src/github.com/letsencrypt/pebble/test/certs/pebble.minica.pem"

// Domain for certificates
const pebbleCertDomain = "example.test"

// URL for the non-EAB pebble directory
const pebbleDirBasic = "https://localhost:14000/dir"

// URL for the EAB pebble directory
const pebbleDirEAB = "https://localhost:14001/dir"

// URL for the static-profile pebble directory
const pebbleDirProfile = "https://localhost:14002/dir"

// Address for the challenge/test recursive nameserver
const pebbleChallTestDNSSrv = "localhost:5553"

// Relative path to the external challenge/test script
const pebbleChallTestDNSScriptPath = "../build-support/scripts/pebble-challtest-dns.sh"

// URL to the main certificate for regular tests
const mainIntermediateURL = "https://localhost:15000/intermediates/0"

// URL to the alternate certificate for preferred chain tests
const alternateIntermediateURL = "https://localhost:15000/intermediates/1"

// URL to the main certificate for static-profile tests
const profileIntermediateURL = "https://localhost:15002/intermediates/0"

// URL to cert status (non-EAB)
const certStatusURL = "https://localhost:15000/cert-status-by-serial/"

// URL to set custom ARI data (default endpoint)
const setRenewalInfoURL = "https://localhost:15000/set-renewal-info/"

// Host:port for memcached
const memcacheHost = "localhost:11211"

type setCustomARIWindowRequest struct {
	Certificate string
	ARIResponse string
}

type setCustomARIWindowRequestARIResponse struct {
	SuggestedWindow setCustomARIWindowRequestARIResponseSuggestedWindow `json:"suggestedWindow"`
	ExplanationURL  string                                              `json:"explanationURL"`
}

type setCustomARIWindowRequestARIResponseSuggestedWindow struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// getPebbleCertificate gets the certificate at the supplied URL.
func getPebbleCertificate(url string) *x509.Certificate {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(fmt.Errorf("getAlternateIntermediateCertificate: error fetching certificate: %s", err))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("getAlternateIntermediateCertificate: error reading certificate: %s", err))
	}

	certs, err := parsePEMBundle(body)
	if err != nil {
		panic(fmt.Errorf("getAlternateIntermediateCertificate: error reading PEM bundle response: %s", err))
	}

	if len(certs) != 1 {
		panic("getAlternateIntermediateCertificate: expected single certificate in intermediate chain, check pebble config")
	}

	return certs[0]
}

// getPebbleCertificateIssuer returns the issuer CN of the
// certificate at the supplied URL.
func getPebbleCertificateIssuer(url string) string {
	return getPebbleCertificate(url).Issuer.CommonName
}

const (
	certificateStatusValid   = "Valid"
	certificateStatusRevoked = "Revoked"
)

// getStatusForCertificate returns the serial number from a *x509.Certificate.
// Note this currently only works for the non-EAB endpoint.
func getStatusForCertificate(cert *x509.Certificate) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get(fmt.Sprintf("%s/%x", certStatusURL, cert.SerialNumber.Int64()))
	if err != nil {
		panic(fmt.Errorf("getStatusForCertificate: error fetching certificate status: %s", err))
	}

	if resp.StatusCode != 200 {
		panic(fmt.Errorf("getStatusForCertificate: unexpected response status: %s", resp.Status))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("getStatusForCertificate: error reading certificate status: %s", err))
	}

	var result struct {
		Status string
	}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(fmt.Errorf("getStatusForCertificate: error reading certificate status JSON: %s", err))
	}

	return result.Status
}

// setCustomARIWindow sets a custom ARI window for the supplied cert. This is
// needed to test ARI renewals correctly.
func setCustomARIWindow(certPEM string, start, end time.Time, explanationURL string) {
	ariResponse := setCustomARIWindowRequestARIResponse{
		SuggestedWindow: setCustomARIWindowRequestARIResponseSuggestedWindow{
			Start: start,
			End:   end,
		},
		ExplanationURL: explanationURL,
	}
	ariResponseBytes, err := json.Marshal(&ariResponse)
	if err != nil {
		panic(fmt.Errorf("setCustomARIWindow: error encoding custom ARI response: %w", err))
	}

	setCustomARIWindowData(certPEM, string(ariResponseBytes))
}

func setCustomARIWindowData(certPEM, data string) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	reqBody := setCustomARIWindowRequest{
		Certificate: certPEM,
		ARIResponse: data,
	}

	reqBodyBuf, err := json.Marshal(&reqBody)
	if err != nil {
		panic(fmt.Errorf("setCustomARIWindow: error encoding request body: %w", err))
	}

	resp, err := client.Post(setRenewalInfoURL, "application/json", bytes.NewReader(reqBodyBuf))
	if err != nil {
		panic(fmt.Errorf("setCustomARIWindow: error sending request: %w", err))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("setCustomARIWindow: error reading response body: %w", err))
	}

	if resp.StatusCode != 200 {
		panic(
			fmt.Errorf("setCustomARIWindow: unexpected response: %s\n\n%s\n\n%s", resp.Status, body, reqBodyBuf))
	}
}

// External providers (tls)
var testAccExternalProviders = map[string]resource.ExternalProvider{
	"tls": {
		Source: "registry.terraform.io/hashicorp/tls",
	},
}

func init() {
	// Set TF_SCHEMA_PANIC_ON_ERROR as a sanity check on tests.
	os.Setenv("TF_SCHEMA_PANIC_ON_ERROR", "true")

	// Set lego's CA certs to pebble's CA for testing w/pebble
	os.Setenv("LEGO_CA_CERTIFICATES", filepath.Join(build.Default.GOPATH, pebbleCACerts))
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestMain(m *testing.M) {
	if len(os.Args) == 2 && os.Args[1] == dnsplugin.PluginArg {
		// Start the plugin here
		dnsplugin.Serve()
	} else {
		os.Exit(m.Run())
	}
}
