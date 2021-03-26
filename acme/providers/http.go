package providers

import (
	"fmt"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/hashicorp/go-multierror"
)

// HTTPProviderWrapper is a multi-provider wrapper to support multiple
// HTTP challenges.
type HTTPProviderWrapper struct {
	Providers []challenge.Provider
}

// NewHTTPProviderWrapper returns an freshly initialized
// HTTPProviderWrapper.
func NewHTTPProviderWrapper() (*HTTPProviderWrapper, error) {
	return &HTTPProviderWrapper{}, nil
}

// Present implements challenge.Provider for HTTPProviderWrapper.
func (d *HTTPProviderWrapper) Present(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.Providers {
		err = p.Present(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while presenting token for HTTP challenge: %s", err.Error()))
		}
	}

	return err
}

// CleanUp implements challenge.Provider for HTTPProviderWrapper.
func (d *HTTPProviderWrapper) CleanUp(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.Providers {
		err = p.CleanUp(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while cleaning token for HTTP challenge: %s", err.Error()))
		}
	}

	return err
}
