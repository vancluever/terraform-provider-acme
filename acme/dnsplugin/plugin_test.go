package dnsplugin

import (
	"os"
	"testing"

	"github.com/go-acme/lego/v4/challenge"
)

var _ = challenge.ProviderTimeout((*DnsProviderClient)(nil))

func TestMapEnvironmentVariableValues(t *testing.T) {
	oldFoo := os.Getenv("ACME_ENV_TEST_FOO")
	oldBar := os.Getenv("ACME_ENV_TEST_BAR")
	defer os.Setenv("ACME_ENV_TEST_FOO", oldFoo)
	defer os.Setenv("ACME_ENV_TEST_BAR", oldBar)

	expected := "test"
	os.Setenv("ACME_ENV_TEST_FOO", expected)
	mapEnvironmentVariableValues(map[string]string{
		"ACME_ENV_TEST_FOO": "ACME_ENV_TEST_BAR",
	})

	actual := os.Getenv("ACME_ENV_TEST_BAR")
	if expected != actual {
		t.Fatalf("expected ACME_ENV_TEST_BAR to be %q, got %q", expected, actual)
	}
}
