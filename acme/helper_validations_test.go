package acme

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestACME_validateKeyType(t *testing.T) {
	s := "RSA2048"

	_, errs := validateKeyType(s, "key_type")
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
}

func TestACME_validateKeyType_invalid(t *testing.T) {
	s := "512"

	_, errs := validateKeyType(s, "key_type")
	if len(errs) < 1 {
		t.Fatalf("should have given an error")
	}
}

func TestACME_validateDNSChallengeConfig(t *testing.T) {
	m := map[string]any{
		"AWS_FOO": "bar",
	}

	_, errs := validateDNSChallengeConfig(m, "config")
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
}

func TestACME_validateDNSChallengeConfig_invalid(t *testing.T) {
	s := map[string]any{
		"AWS_FOO": 1,
	}

	_, errs := validateDNSChallengeConfig(s, "config")
	if len(errs) < 1 {
		t.Fatalf("should have given an error")
	}
}

func TestACME_validateMatchDomains(t *testing.T) {
	s := []any{
		"example.com",
		"foo.example.com",
		"foo-bar.example.com",
		"foo..example.com",
		"com",
		"com.",
		"foo123.com",
		"foo_123.com",
	}

	_, errs := validateMatchDomains(s, "match_domains")

	expected := []error{
		&comparableError{s: "invalid match domain pattern: foo..example.com"},
		&comparableError{s: "invalid match domain pattern: com."},
		&comparableError{s: "invalid match domain pattern: foo_123.com"},
	}

	if diff := cmp.Diff(expected, errs, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("expected errors mismatch (-want +got):\n%s", diff)
	}
}

type comparableError struct {
	s string
}

func (e *comparableError) Error() string {
	return e.s
}

func (e *comparableError) Is(target error) bool {
	return e.s == target.Error()
}
