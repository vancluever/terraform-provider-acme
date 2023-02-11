package dnsplugin

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	dnspluginproto "github.com/vancluever/terraform-provider-acme/v2/proto/dnsplugin/v1"
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

type testDummyProviderNoTimeout struct{}

func (p *testDummyProviderNoTimeout) Present(_, _, _ string) error { return nil }
func (p *testDummyProviderNoTimeout) CleanUp(_, _, _ string) error { return nil }

func TestDnsProviderServerTimeout(t *testing.T) {
	testCases := []struct {
		desc         string
		provider     challenge.Provider
		wantTimeout  time.Duration
		wantInterval time.Duration
	}{
		{
			desc: "with timeout",
			provider: func() challenge.Provider {
				p, err := route53.NewDNSProvider()
				if err != nil {
					panic(err)
				}

				return p
			}(),
			wantTimeout:  2 * time.Minute,
			wantInterval: 4 * time.Second,
		},
		{
			desc:     "without timeout",
			provider: &testDummyProviderNoTimeout{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			server := &DnsProviderServer{provider: tc.provider}
			resp, err := server.Timeout(context.Background(), &dnspluginproto.TimeoutRequest{})
			if err != nil {
				t.Fatal(err)
			}

			if tc.wantTimeout != resp.Timeout.AsDuration() {
				t.Fatalf("want duration %s, got duration %s", tc.wantTimeout, resp.Timeout.AsDuration())
			}

			if tc.wantInterval != resp.Interval.AsDuration() {
				t.Fatalf("want duration %s, got duration %s", tc.wantTimeout, resp.Interval.AsDuration())
			}
		})
	}
}
