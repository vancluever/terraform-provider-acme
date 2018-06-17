package acme

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/terraform"
)

func testACMERegistrationStateDataV0() *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: "regurl",
		Attributes: map[string]string{
			"server_url":                 "https://acme-staging.api.letsencrypt.org/directory",
			"account_key_pem":            "key",
			"email_address":              "nobody@example.com",
			"registration_body":          "regbody",
			"registration_url":           "https://acme-staging.api.letsencrypt.org/acme/reg/123456789",
			"registration_new_authz_url": "https://acme-staging.api.letsencrypt.org/acme/new-authz",
			"registration_tos_url":       "https://letsencrypt.org/documents/LE-SA-v1.0.1-July-27-2015.pdf",
		},
	}
}

func testACMERegistrationStateDataV1() *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: "regurl",
		Attributes: map[string]string{
			"server_url":      "https://acme-staging.api.letsencrypt.org/directory",
			"account_key_pem": "key",
			"email_address":   "nobody@example.com",
		},
	}
}

func testACMECertificateStateDataV0() *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: "certurl",
		Attributes: map[string]string{
			"server_url":                  "https://acme-staging.api.letsencrypt.org/directory",
			"account_key_pem":             "key",
			"common_name":                 "foobar",
			"subject_alternative_names.#": "2",
			"subject_alternative_names.0": "barbar",
			"subject_alternative_names.1": "bazbar",
			"key_type":                    "2048",
			"certificate_request_pem":     "req",
			"min_days_remaining":          "7",
			"dns_challenge.%":             "1",
			"dns_challenge.1234.provider": "route53",
			"http_challenge_port":         "80",
			"tls_challenge_port":          "443",
			"registration_url":            "regurl",
			"must_staple":                 "0",
			"certificate_domain":          "foobar",
			"certificate_url":             "certurl",
			"account_ref":                 "regurl",
			"private_key_pem":             "certkey",
			"certificate_pem":             "certpem",
		},
	}
}

func testACMECertificateStateDataV1() *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: "certurl",
		Attributes: map[string]string{
			"server_url":                  "https://acme-staging.api.letsencrypt.org/directory",
			"account_key_pem":             "key",
			"common_name":                 "foobar",
			"subject_alternative_names.#": "2",
			"subject_alternative_names.0": "barbar",
			"subject_alternative_names.1": "bazbar",
			"key_type":                    "2048",
			"certificate_request_pem":     "req",
			"min_days_remaining":          "7",
			"dns_challenge.%":             "1",
			"dns_challenge.1234.provider": "route53",
			"must_staple":                 "0",
			"certificate_domain":          "foobar",
			"account_ref":                 "regurl",
			"private_key_pem":             "certkey",
			"certificate_pem":             "certpem",
		},
	}
}

func TestResourceACMERegistrationMigrateState(t *testing.T) {
	expected := testACMERegistrationStateDataV1()
	actual, err := resourceACMERegistrationMigrateState(0, testACMERegistrationStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}

func TestMigrateACMERegistrationStateV1(t *testing.T) {
	expected := testACMERegistrationStateDataV1()
	actual := testACMERegistrationStateDataV0()
	if err := migrateACMERegistrationStateV1(actual, nil); err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}

func TestResourceACMECertificateMigrateState(t *testing.T) {
	expected := testACMECertificateStateDataV1()
	actual, err := resourceACMECertificateMigrateState(0, testACMECertificateStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}

func TestMigrateACMECertificateStateV1(t *testing.T) {
	expected := testACMECertificateStateDataV1()
	actual := testACMECertificateStateDataV0()
	if err := migrateACMECertificateStateV1(actual, nil); err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}
