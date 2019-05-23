package acme

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func testACMECertificateStateData012V3() map[string]interface{} {
	return map[string]interface{}{
		"account_key_pem":           "key",
		"common_name":               "foobar",
		"subject_alternative_names": []interface{}{"barbar", "bazbar"},
		"key_type":                  "2048",
		"certificate_request_pem":   "req",
		"min_days_remaining":        "30",
		"dns_challenge": []interface{}{
			map[string]interface{}{
				"provider":              "route53",
				"recursive_nameservers": []interface{}{"my.name.server"},
			},
		},
		"must_staple":        "0",
		"certificate_domain": "foobar",
		"private_key_pem":    "certkey",
		"certificate_pem":    "certpem",
		"certificate_url":    "certurl",
	}
}

func testACMECertificateStateData012V4() map[string]interface{} {
	return map[string]interface{}{
		"account_key_pem":           "key",
		"common_name":               "foobar",
		"subject_alternative_names": []interface{}{"barbar", "bazbar"},
		"key_type":                  "2048",
		"certificate_request_pem":   "req",
		"min_days_remaining":        "30",
		"dns_challenge": []interface{}{
			map[string]interface{}{
				"provider": "route53",
			},
		},
		"recursive_nameservers": []interface{}{"my.name.server"},
		"must_staple":           "0",
		"certificate_domain":    "foobar",
		"private_key_pem":       "certkey",
		"certificate_pem":       "certpem",
		"certificate_url":       "certurl",
	}
}

func TestResourceACMECertificateStateUpgraderV3Func(t *testing.T) {
	expected := testACMECertificateStateData012V4()
	actual, err := resourceACMECertificateStateUpgraderV3Func(testACMECertificateStateData012V3(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", spew.Sdump(expected), spew.Sdump(actual))
	}
}
