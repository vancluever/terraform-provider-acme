package acme

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider

// Path to the pebble CA cert list, from GOPATH
const pebbleCACerts = "src/github.com/letsencrypt/pebble/test/certs/pebble.minica.pem"

// Domain for certificates
const pebbleCertDomain = "example.test"

// URL for the non-EAB pebble directory
const pebbleDirBasic = "https://localhost:14000/dir"

// URL for the EAB pebble directory
const pebbleDirEAB = "https://localhost:14001/dir"

// Address for the challenge/test recursive nameserver
const pebbleChallTestDNSSrv = "localhost:5553"

// Relative path to the external challenge/test script
const pebbleChallTestDNSScriptPath = "../build-support/scripts/pebble-challtest-dns.sh"

// URL to the main certificate for regular tests
const mainIntermediateURL = "https://localhost:15000/intermediates/0"

// URL to the alternate certificate for preferred chain tests
const alternateIntermediateURL = "https://localhost:15000/intermediates/1"

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

	testAccProviders = map[string]*schema.Provider{
		"acme": Provider(),
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
