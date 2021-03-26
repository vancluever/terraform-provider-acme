// Auto-generated file. Do not edit.
package acme

import (
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/exec"
	"github.com/go-acme/lego/v4/providers/http/memcached"
	"github.com/go-acme/lego/v4/providers/http/webroot"
)

// httpProviderFactoryFunc is a function that calls a provider's
// constructor and returns the provider interface.
type httpProviderFactoryFunc func(config interface{}) (challenge.Provider, error)

type WebRootArgs struct {
	Path string
}

type MemcachedArgs struct {
	Hosts []string
}

// httpProviderFactory is a factory for all of the valid HTTP providers
// supported by ACME provider.
var httpProviderFactory = map[string]httpProviderFactoryFunc{
	"exec": func(_ interface{}) (challenge.Provider, error) {
		p, err := exec.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"webroot": func(config interface{}) (challenge.Provider, error) {
		p, err := webroot.NewHTTPProvider(config.(WebRootArgs).Path)
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"memcached": func(config interface{}) (challenge.Provider, error) {
		p, err := memcached.NewMemcachedProvider(config.(MemcachedArgs).Hosts)
		if err != nil {
			return nil, err
		}

		return p, nil
	},
}
