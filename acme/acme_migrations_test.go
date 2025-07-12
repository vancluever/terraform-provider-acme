package acme

import (
	"context"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func testACMECertificateStateData012V3() map[string]any {
	return map[string]any{
		"account_key_pem":           "key",
		"common_name":               "foobar",
		"subject_alternative_names": []any{"barbar", "bazbar"},
		"key_type":                  "2048",
		"certificate_request_pem":   "req",
		"min_days_remaining":        "30",
		"dns_challenge": []any{
			map[string]any{
				"provider":              "route53",
				"recursive_nameservers": []any{"my.name.server"},
			},
		},
		"must_staple":        "0",
		"certificate_domain": "foobar",
		"private_key_pem":    "certkey",
		"certificate_pem":    "certpem",
		"certificate_url":    "certurl",
	}
}

func testACMECertificateStateData012V4() map[string]any {
	return map[string]any{
		"account_key_pem":           "key",
		"common_name":               "foobar",
		"subject_alternative_names": []any{"barbar", "bazbar"},
		"key_type":                  "2048",
		"certificate_request_pem":   "req",
		"min_days_remaining":        "30",
		"dns_challenge": []any{
			map[string]any{
				"provider": "route53",
			},
		},
		"recursive_nameservers": []any{"my.name.server"},
		"must_staple":           "0",
		"certificate_domain":    "foobar",
		"private_key_pem":       "certkey",
		"certificate_pem":       "certpem",
		"certificate_url":       "certurl",
	}
}

func testACMECertificateStateData012V5() map[string]any {
	return map[string]any{
		"account_key_pem":           "key",
		"common_name":               "foobar",
		"subject_alternative_names": []any{"barbar", "bazbar"},
		"key_type":                  "2048",
		"certificate_request_pem":   "req",
		"min_days_remaining":        "30",
		"dns_challenge": []any{
			map[string]any{
				"provider": "route53",
			},
		},
		"recursive_nameservers": []any{"my.name.server"},
		"must_staple":           "0",
		"certificate_domain":    "foobar",
		"private_key_pem":       "certkey",
		"certificate_pem":       "certpem",
		"certificate_url":       "certurl",
	}
}

func testACMERegistrationStateData012V1() map[string]any {
	return map[string]any{
		"account_key_pem": "key",
		"email_address":   "hello@localhost",
		"external_account_binding": []any{
			map[string]any{
				"key_id":      "kid",
				"hmac_base64": "hmac",
			},
		},
		"id":               "id",
		"registration_url": "regurl",
	}
}

func testACMERegistrationStateData012V2() map[string]any {
	return map[string]any{
		"account_key_pem":         "key",
		"account_key_algorithm":   "ECDSA",
		"account_key_ecdsa_curve": "P384",
		"account_key_rsa_bits":    4096,
		"email_address":           "hello@localhost",
		"external_account_binding": []any{
			map[string]any{
				"key_id":      "kid",
				"hmac_base64": "hmac",
			},
		},
		"id":               "id",
		"registration_url": "regurl",
	}
}

func TestResourceACMECertificateStateUpgraderV3Func(t *testing.T) {
	expected := testACMECertificateStateData012V4()
	actual, err := resourceACMECertificateStateUpgraderV3Func(context.TODO(), testACMECertificateStateData012V3(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", spew.Sdump(expected), spew.Sdump(actual))
	}
}

func TestResourceACMECertificateStateUpgraderV4Func(t *testing.T) {
	expected := testACMECertificateStateData012V5()
	actual, err := resourceACMECertificateStateUpgraderV4Func(context.TODO(), testACMECertificateStateData012V4(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	ignore := cmpopts.IgnoreMapEntries(func(k string, _ any) bool {
		return k == "id"
	})

	if diff := cmp.Diff(expected, actual, ignore); diff != "" {
		t.Errorf("state migration v4 -> v5 mismatch (-want +got):\n%s", diff)
	}

	if id := actual["id"].(string); !uuidRegexp.MatchString(id) {
		t.Errorf("expected UUID as ID, got %q", id)
	}
}

func TestResourceACMERegistrationStateUpgraderV1Func(t *testing.T) {
	expected := testACMERegistrationStateData012V2()
	actual, err := resourceACMERegistrationStateUpgraderV1Func(context.TODO(), testACMERegistrationStateData012V1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if diff := cmp.Diff(expected, actual, nil); diff != "" {
		t.Errorf("state migration v1 -> v2 mismatch (-want +got):\n%s", diff)
	}
}
