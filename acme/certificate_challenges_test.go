package acme

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-acme/lego/v5/challenge/dns01"
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

				d.Set("dns_challenge", []any{
					map[string]any{
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

				d.Set("dns_challenge", []any{
					map[string]any{
						"provider": "exec",
						"config": map[string]any{
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

				d.Set("dns_challenge", []any{
					map[string]any{
						"provider": "exec",
						"config": map[string]any{
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

				d.Set("dns_challenge", []any{
					map[string]any{
						"provider": "route53",
					},
					map[string]any{
						"provider": "exec",
						"config": map[string]any{
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

				d.Set("dns_challenge", []any{
					map[string]any{
						"provider": "exec",
						"config": map[string]any{
							"EXEC_PATH":              "exit 0",
							"EXEC_SEQUENCE_INTERVAL": "60", // explicit default
						},
					},
					map[string]any{
						"provider": "exec",
						"config": map[string]any{
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
				tc.resourceData.Get("dns_challenge").([]any),
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

func TestACME_DNSProviderWrapper_matchDomains(t *testing.T) {
	testCases := []struct {
		desc        string
		pattern     string
		domain      string
		wantLevel   int
		wantMatched bool
	}{
		{
			desc:        "basic",
			pattern:     "example.com",
			domain:      "example.com",
			wantLevel:   2,
			wantMatched: true,
		},
		{
			desc:        "mixed case",
			pattern:     "exAmpLe.com",
			domain:      "ExaMple.com",
			wantLevel:   2,
			wantMatched: true,
		},
		{
			desc:        "basic exclusionary",
			pattern:     "example.com",
			domain:      "foo.com",
			wantLevel:   0,
			wantMatched: false,
		},
		{
			desc:        "subdomain domain",
			pattern:     "example.com",
			domain:      "foo.example.com",
			wantLevel:   2,
			wantMatched: true,
		},
		{
			desc:        "subdomain pattern (matching)",
			pattern:     "foo.example.com",
			domain:      "bar.foo.example.com",
			wantLevel:   3,
			wantMatched: true,
		},
		{
			desc:        "subdomain pattern (not matching)",
			pattern:     "foo.example.com",
			domain:      "example.com",
			wantLevel:   0,
			wantMatched: false,
		},
		{
			desc:        "wildcard domain",
			pattern:     "example.com",
			domain:      "*.example.com",
			wantLevel:   2,
			wantMatched: true,
		},
		{
			// Note that this should be blocked in the validator (should not even be
			// possible to use)
			desc:        "wildcard pattern",
			pattern:     "*.example.com",
			domain:      "example.com",
			wantLevel:   0,
			wantMatched: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotLevel, gotMatched := matchDomain(tc.pattern, tc.domain)
			if tc.wantLevel != gotLevel {
				t.Fatalf("wanted level %d, got %d", tc.wantLevel, gotLevel)
			}
			if tc.wantMatched != gotMatched {
				t.Fatalf("wanted matched %t, got %t", tc.wantMatched, gotMatched)
			}
		})
	}
}

func TestACME_DNSProviderWrapper_filterMatches(t *testing.T) {
	testCases := []struct {
		desc         string
		domain       string
		matchDomains [][]string
		want         providerDomainMatchEntry
	}{
		{
			desc:   "basic",
			domain: "example.com",
			matchDomains: [][]string{
				{
					"example.com",
				},
			},
			want: providerDomainMatchEntry{
				providerIndexes: []int{0},
				level:           2,
			},
		},
		{
			desc:   "multi-list",
			domain: "foo.bar.com",
			matchDomains: [][]string{
				{
					"example.com",
					"foo.bar.com",
				},
			},
			want: providerDomainMatchEntry{
				providerIndexes: []int{0},
				level:           3,
			},
		},
		{
			desc:   "multi-provider",
			domain: "bar.com",
			matchDomains: [][]string{
				{
					"example.com",
					"foo.bar.com",
				},
				{
					"bar.com",
				},
			},
			want: providerDomainMatchEntry{
				providerIndexes: []int{1},
				level:           2,
			},
		},
		{
			desc:   "provider override",
			domain: "foo.bar.com",
			matchDomains: [][]string{
				{
					"example.com",
					"foo.bar.com",
				},
				{
					"bar.com",
				},
				{
					"foo.bar.com",
				},
			},
			want: providerDomainMatchEntry{
				providerIndexes: []int{0, 2},
				level:           3,
			},
		},
		{
			desc:         "empty matches",
			domain:       "foo.bar.com",
			matchDomains: [][]string{{}},
			want: providerDomainMatchEntry{
				providerIndexes: []int{0},
				level:           0,
			},
		},
		{
			desc:   "no matches",
			domain: "example.com",
			matchDomains: [][]string{
				{
					"foo.com",
				},
			},
			want: providerDomainMatchEntry{
				providerIndexes: []int{},
				level:           -1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			wrapper := &DNSProviderWrapper{
				matchDomains: tc.matchDomains,
			}

			wrapper.filterMatches(tc.domain)

			got := wrapper.providerDomainMatches[tc.domain]
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want: %#v, got %#v", tc.want, got)
			}
		})
	}
}
