package acme

import (
	"fmt"
	"regexp"

	"github.com/go-acme/lego/v5/certcrypto"
)

// validKeyTypes returns all the valid key types as a string slice.
func validKeyTypes() []string {
	result := make([]string, 0, len(certcrypto.AllKeyTypes()))
	for _, kt := range certcrypto.AllKeyTypes() {
		result = append(result, kt.String())
	}

	return result
}

// validateKeyTypeV3Old validates a key_type resource parameter is correct.
//
// Deprecated: Do not use anymore.
func validateKeyTypeV3Old(v any, _ string) (ws []string, errors []error) {
	value := v.(string)
	found := false
	for _, w := range []string{"P256", "P384", "RSA2048", "RSA4096", "RSA8192"} {
		if value == w {
			found = true
		}
	}
	if !found {
		errors = append(errors, fmt.Errorf(
			"certificate key type must be one of P256, P384, RSA2048, RSA4096, or RSA8192"))
	}
	return
}

// validateKeyType validates a key_type resource parameter is correct.
//
// Deprecated: Do not use anymore.
func validateKeyType(v any, k string) (ws []string, errors []error) {
	value := v.(string)
	found := false
	for _, w := range []string{"P256", "P384", "2048", "4096", "8192"} {
		if value == w {
			found = true
		}
	}
	if !found {
		errors = append(errors, fmt.Errorf(
			"certificate key type must be one of P256, P384, 2048, 4096, or 8192"))
	}
	return
}

// validateDNSChallengeConfig ensures that the values supplied to the
// dns_challenge resource parameter in the acme_certificate resource
// are string values only.
func validateDNSChallengeConfig(v any, _ string) (ws []string, errors []error) {
	value := v.(map[string]any)
	bad := false
	for _, w := range value {
		switch w.(type) {
		case string:
			continue
		default:
			bad = true
		}
	}
	if bad {
		errors = append(errors, fmt.Errorf(
			"DNS challenge config map values must be strings only"))
	}
	return
}

func validateRevocationReason(v any, _ string) (ws []string, errors []error) {
	value := RevocationReason(v.(string))
	_, err := GetRevocationReason(value)
	if err != nil {
		errors = append(errors, err)
	}
	return
}

func validateMatchDomains(v any, _ string) (ws []string, errors []error) {
	domainRegexp := regexp.MustCompile(`^([-a-zA-Z0-9]+\.)*[-a-zA-Z0-9]+$`)
	value := stringSlice(v.([]any))
	for _, w := range value {
		if domainRegexp.Match([]byte(w)) {
			continue
		}
		errors = append(errors, fmt.Errorf("invalid match domain pattern: %s", w))
	}
	return
}
