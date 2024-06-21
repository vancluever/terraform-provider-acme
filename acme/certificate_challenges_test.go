package acme

import (
	"testing"
	"time"

	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExpandDNSChallengeWrapperProvider(t *testing.T) {
	testCases := []struct {
		desc          string
		resourceData  *schema.ResourceData
		wantSeq       bool
		wantInterval  time.Duration
		wantCloserLen int
	}{
		{
			desc: "single parallel provider (route53)",
			resourceData: func() *schema.ResourceData {
				r := resourceACMECertificate()
				d := r.TestResourceData()

				d.Set("dns_challenge", []interface{}{
					map[string]interface{}{
						"provider": "route53",
					},
				})

				return d
			}(),
			wantSeq:       false,
			wantCloserLen: 1,
		},

		{
			desc: "sequential provider (exec w/defaults)",
			resourceData: func() *schema.ResourceData {
				r := resourceACMECertificate()
				d := r.TestResourceData()

				d.Set("dns_challenge", []interface{}{
					map[string]interface{}{
						"provider": "exec",
						"config": map[string]interface{}{
							"EXEC_PATH": "exit 0",
						},
					},
				})

				return d
			}(),
			wantSeq:       true,
			wantInterval:  dns01.DefaultPropagationTimeout,
			wantCloserLen: 1,
		},

		{
			desc: "sequential provider (exec w/interval)",
			resourceData: func() *schema.ResourceData {
				r := resourceACMECertificate()
				d := r.TestResourceData()

				d.Set("dns_challenge", []interface{}{
					map[string]interface{}{
						"provider": "exec",
						"config": map[string]interface{}{
							"EXEC_PATH":              "exit 0",
							"EXEC_SEQUENCE_INTERVAL": "123",
						},
					},
				})

				return d
			}(),
			wantSeq:       true,
			wantInterval:  time.Second * 123,
			wantCloserLen: 1,
		},

		{
			desc: "mixed w/defaults",
			resourceData: func() *schema.ResourceData {
				r := resourceACMECertificate()
				d := r.TestResourceData()

				d.Set("dns_challenge", []interface{}{
					map[string]interface{}{
						"provider": "route53",
					},
					map[string]interface{}{
						"provider": "exec",
						"config": map[string]interface{}{
							"EXEC_PATH": "exit 0",
						},
					},
				})

				return d
			}(),
			wantSeq:       true,
			wantInterval:  dns01.DefaultPropagationTimeout,
			wantCloserLen: 2,
		},

		{
			desc: "multiple w/varying intervals",
			resourceData: func() *schema.ResourceData {
				r := resourceACMECertificate()
				d := r.TestResourceData()

				d.Set("dns_challenge", []interface{}{
					map[string]interface{}{
						"provider": "exec",
						"config": map[string]interface{}{
							"EXEC_PATH":              "exit 0",
							"EXEC_SEQUENCE_INTERVAL": "60", // explicit default
						},
					},
					map[string]interface{}{
						"provider": "exec",
						"config": map[string]interface{}{
							"EXEC_PATH":              "exit 0",
							"EXEC_SEQUENCE_INTERVAL": "123",
						},
					},
				})

				return d
			}(),
			wantSeq:       true,
			wantInterval:  time.Second * 123,
			wantCloserLen: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, gotClosers, err := expandDNSChallengeWrapperProvider(
				tc.resourceData,
				tc.resourceData.Get("dns_challenge").([]interface{}),
			)
			if err != nil {
				t.Fatal(err)
			}

			switch g := got.(type) {
			case *DNSProviderWrapper:
				if tc.wantSeq {
					t.Fatal("expected parallel provider, got sequential")
				}
			case *DNSProviderWrapperSequential:
				if !tc.wantSeq {
					t.Fatal("expected sequential provider, got parallel")
				}

				if tc.wantInterval != g.interval {
					t.Fatalf("want interval %s, got interval %s", tc.wantInterval, g.interval)
				}
			}

			if tc.wantCloserLen != len(gotClosers) {
				t.Fatalf("want closer len %d, got closer len %d", tc.wantCloserLen, len(gotClosers))
			}
		})
	}
}
