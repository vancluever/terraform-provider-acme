package acme

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-tls/v2/tls"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

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

func init() {
	// Set TF_SCHEMA_PANIC_ON_ERROR as a sanity check on tests.
	os.Setenv("TF_SCHEMA_PANIC_ON_ERROR", "true")

	// Set lego's CA certs to pebble's CA for testing w/pebble
	os.Setenv("LEGO_CA_CERTIFICATES", filepath.Join(build.Default.GOPATH, pebbleCACerts))

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"acme": testAccProvider,
		"tls":  tls.Provider().(*schema.Provider),
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
