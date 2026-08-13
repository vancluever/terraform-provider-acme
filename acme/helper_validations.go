package acme

import "fmt"

// validateKeyType validates a key_type resource parameter is correct.
func validateKeyType(v any, k string) (ws []string, errors []error) {
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

// validateDNSChallengeConfig ensures that the values supplied to the
// dns_challenge resource parameter in the acme_certificate resource
// are string values only.
func validateDNSChallengeConfig(v any, k string) (ws []string, errors []error) {
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

func validateRevocationReason(v any, k string) (ws []string, errors []error) {
	value := RevocationReason(v.(string))
	_, err := GetRevocationReason(value)
	if err != nil {
		errors = append(errors, err)
	}
	return
}
