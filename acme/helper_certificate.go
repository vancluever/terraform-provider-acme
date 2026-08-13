package acme

import (
	"fmt"
	"time"

	"github.com/go-acme/lego/v5/certificate"
)

// certSecondsRemaining takes an certificate.Resource, parses the
// certificate, and computes the seconds that it has remaining.
func certSecondsRemaining(cert *certificate.Resource, now time.Time) (int64, error) {
	x509Certs, err := parsePEMBundle(cert.Certificate)
	if err != nil {
		return 0, err
	}
	c := x509Certs[0]

	if c.IsCA {
		return 0, fmt.Errorf("first certificate is a CA certificate")
	}

	return (c.NotAfter.Unix() - now.Unix()), nil
}

// certDaysRemaining takes an certificate.Resource, parses the
// certificate, and computes the days that it has remaining.
func certDaysRemaining(cert *certificate.Resource, now time.Time) (int64, error) {
	remaining, err := certSecondsRemaining(cert, now)
	if err != nil {
		return 0, fmt.Errorf("unable to calculate time to certificate expiry: %s", err)
	}

	return remaining / 86400, nil
}

// certLifetimeDays returns the total certificate lifetime in days.
func certLifetimeDays(cert *certificate.Resource) (float64, error) {
	x509Certs, err := parsePEMBundle(cert.Certificate)
	if err != nil {
		return 0, err
	}
	c := x509Certs[0]

	if c.IsCA {
		return 0, fmt.Errorf("first certificate is a CA certificate")
	}

	return c.NotAfter.Sub(c.NotBefore).Round(time.Hour).Hours() / 24.0, nil
}
