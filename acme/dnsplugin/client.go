package dnsplugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/hashicorp/go-plugin"
	dnspluginproto "github.com/vancluever/terraform-provider-acme/v2/proto/dnsplugin/v1"
)

// NewClient creates a new DNS provider instance by dispatching to itself via
// go-plugin. The client for the new provider is returned, along with a closer
// function that should be called when done to shut down the plugin.
//
// The plugin is configured the the map passed in, this sets the local
// environment without it leaking into the parent process or any side-by-side
// providers.
func NewClient(providerName string, config map[string]string) (challenge.Provider, func(), error) {
	// Discover the path to the executable that we are running at.
	execPath, err := os.Executable()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting plugin path: %w", err)
	}

	// Create the command
	cmd := &exec.Cmd{
		Path: execPath,
		Args: []string{PluginArg, fmt.Sprintf("[%s]", providerName)},
	}

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
		if err := dnsProviderClient.Configure(providerName, config); err != nil {
			return nil, nil, fmt.Errorf("error configuring plugin: %w", err)
		}
	} else {
		return nil, nil, errors.New("internal error: returned plugin not a DnsProviderClient")
	}

	provider, ok := raw.(challenge.Provider)
	if !ok {
		return nil, nil, errors.New("internal error: returned plugin not a challenge provider")
	}

	return provider, func() { rpcClient.Close() }, nil
}

type DnsProviderClient struct {
	client dnspluginproto.DNSProviderServiceClient
}

func (m *DnsProviderClient) Configure(providerName string, config map[string]string) error {
	_, err := m.client.Configure(context.Background(), &dnspluginproto.ConfigureRequest{
		ProviderName: providerName,
		Config:       config,
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
