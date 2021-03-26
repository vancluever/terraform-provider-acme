package providers

import (
	"fmt"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/hashicorp/go-multierror"
	"time"
)

// DNSProviderWrapper is a multi-provider wrapper to support multiple
// DNS challenges.
type DNSProviderWrapper struct {
	Providers []challenge.Provider
}

// NewDNSProviderWrapper returns an freshly initialized
// DNSProviderWrapper.
func NewDNSProviderWrapper() (*DNSProviderWrapper, error) {
	return &DNSProviderWrapper{}, nil
}

// Present implements challenge.Provider for DNSProviderWrapper.
func (d *DNSProviderWrapper) Present(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.Providers {
		err = p.Present(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while presenting token for DNS challenge: %s", err.Error()))
		}
	}

	return err
}

// CleanUp implements challenge.Provider for DNSProviderWrapper.
func (d *DNSProviderWrapper) CleanUp(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.Providers {
		err = p.CleanUp(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while cleaning token for DNS challenge: %s", err.Error()))
		}
	}

	return err
}

// Timeout implements challenge.ProviderTimeout for
// DNSProviderWrapper.
//
// The highest polling interval and timeout values defined across all
// providers is used.
func (d *DNSProviderWrapper) Timeout() (time.Duration, time.Duration) {
	var timeout, interval time.Duration
	for _, p := range d.Providers {
		if pt, ok := p.(challenge.ProviderTimeout); ok {
			t, i := pt.Timeout()
			if t > timeout {
				timeout = t
			}

			if i > interval {
				interval = i
			}
		}
	}

	if timeout < 1 {
		timeout = dns01.DefaultPropagationTimeout
	}

	if interval < 1 {
		interval = dns01.DefaultPollingInterval
	}

	return timeout, interval
}
