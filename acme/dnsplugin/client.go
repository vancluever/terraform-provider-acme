package dnsplugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/hashicorp/go-plugin"
	dnspluginproto "github.com/vancluever/terraform-provider-acme/v2/proto/dnsplugin/v1"
)

// NewClient creates a new DNS provider instance by dispatching to itself via
// go-plugin. The client for the new provider is returned, along with a closer
// function that should be called when done to shut down the plugin.
//
// The plugin is initialized with the settings passed in:
//   - The environment is set with the config map.
//   - If supplied, the global recursive nameservers are also set (via the
//     dns01 package - some providers use these facilities).
func NewClient(
	providerName string,
	config map[string]string,
	recursiveNameservers []string,
) (challenge.ProviderTimeout, func(), error) {
	// Discover the path to the executable that we are running at.
	execPath, err := os.Executable()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting plugin path: %w", err)
	}

	cmd := exec.Command(execPath, "-dnsplugin")
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  Handshake,
		AutoMTLS:         true,
		Plugins:          map[string]plugin.Plugin{PluginName: &DnsPlugin{}},
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})

	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting plugin client: %w", err)
	}

	raw, err := rpcClient.Dispense(PluginName)
	if err != nil {
		return nil, nil, fmt.Errorf("error dispensing plugin: %w", err)
	}

	// First call the plugin as its gRPC server interface so that we can
	// configure it.
	if dnsProviderClient, ok := raw.(*DnsProviderClient); ok {
		if err := dnsProviderClient.Configure(providerName, config, recursiveNameservers); err != nil {
			return nil, nil, fmt.Errorf("error configuring plugin: %w", err)
		}
	} else {
		return nil, nil, errors.New("internal error: returned plugin not a DnsProviderClient")
	}

	provider, ok := raw.(challenge.ProviderTimeout)
	if !ok {
		return nil, nil, errors.New("internal error: returned plugin not a challenge provider")
	}

	return provider, func() { rpcClient.Close() }, nil
}

type DnsProviderClient struct {
	client dnspluginproto.DNSProviderServiceClient
}

func (m *DnsProviderClient) Configure(providerName string, config map[string]string, recursiveNameservers []string) error {
	_, err := m.client.Configure(context.Background(), &dnspluginproto.ConfigureRequest{
		ProviderName:         providerName,
		Config:               config,
		RecursiveNameservers: recursiveNameservers,
	})
	return err
}

func (m *DnsProviderClient) Present(domain, token, keyAuth string) error {
	_, err := m.client.Present(context.Background(), &dnspluginproto.PresentRequest{
		Domain:  domain,
		Token:   token,
		KeyAuth: keyAuth,
	})
	return err
}

func (m *DnsProviderClient) CleanUp(domain, token, keyAuth string) error {
	_, err := m.client.CleanUp(context.Background(), &dnspluginproto.CleanUpRequest{
		Domain:  domain,
		Token:   token,
		KeyAuth: keyAuth,
	})
	return err
}

func (m *DnsProviderClient) Timeout() (time.Duration, time.Duration) {
	resp, _ := m.client.Timeout(context.Background(), &dnspluginproto.TimeoutRequest{})
	return resp.GetTimeout().AsDuration(), resp.GetInterval().AsDuration()
}
