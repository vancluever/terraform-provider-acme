package dnsplugin

//go:generate go run ../../build-support/generate-dns-providers go dns_provider_factory.go

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/hashicorp/go-plugin"
	dnspluginproto "github.com/vancluever/terraform-provider-acme/v2/proto/dnsplugin/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	ProtocolVersion  = 1
	MagicCookieKey   = "TERRAFORM_PROVIDER_ACME_MAGIC_COOKIE_KEY"
	MagicCookieValue = "990EF127-D8AA-43D3-9196-9493C2D6C475"
	PluginName       = "dnsplugin"
	PluginArg        = "-dnsplugin"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  ProtocolVersion,
	MagicCookieKey:   MagicCookieKey,
	MagicCookieValue: MagicCookieValue,
}

// Serve serves the DNS plugin. This function does not retun and will cause the
// process to exit after the plugin is finished running.
func Serve() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			PluginName: &DnsPlugin{},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

// DnsPlugin the gRPC plugin for serving the DNS plugin that has been set.
type DnsPlugin struct {
	plugin.Plugin
}

func (p *DnsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	dnspluginproto.RegisterDNSProviderServiceServer(s, &DnsProviderServer{})
	return nil
}

func (p *DnsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &DnsProviderClient{client: dnspluginproto.NewDNSProviderServiceClient(c)}, nil
}

type DnsProviderServer struct {
	dnspluginproto.UnimplementedDNSProviderServiceServer

	provider challenge.Provider
}

func (s *DnsProviderServer) Configure(ctx context.Context, req *dnspluginproto.ConfigureRequest) (*dnspluginproto.ConfigureResponse, error) {
	providerFunc, ok := dnsProviderFactory[req.GetProviderName()]
	if !ok {
		return nil, fmt.Errorf("unknown provider name %q", req.GetProviderName())
	}

	if len(req.GetRecursiveNameservers()) > 0 {
		// Configure recursive nameservers in the dns01 package. Some providers use
		// functionality in dns01 that depend on these and the changes made in the
		// parent process will not propagate.
		if err := dns01.AddRecursiveNameservers(req.GetRecursiveNameservers())(nil); err != nil {
			return nil, fmt.Errorf("error setting recursive nameservers: %w", err)
		}
	}

	// Set env before configuring provider
	for k, v := range req.GetConfig() {
		os.Setenv(k, v)
	}

	var err error
	s.provider, err = providerFunc()
	if err != nil {
		return nil, fmt.Errorf("error initializing provider: %w", err)
	}

	return &dnspluginproto.ConfigureResponse{}, nil
}

func (m *DnsProviderServer) Present(ctx context.Context, req *dnspluginproto.PresentRequest) (*dnspluginproto.PresentResponse, error) {
	return &dnspluginproto.PresentResponse{}, m.provider.Present(req.GetDomain(), req.GetToken(), req.GetKeyAuth())
}

func (m *DnsProviderServer) CleanUp(ctx context.Context, req *dnspluginproto.CleanUpRequest) (*dnspluginproto.CleanUpResponse, error) {
	return &dnspluginproto.CleanUpResponse{}, m.provider.CleanUp(req.GetDomain(), req.GetToken(), req.GetKeyAuth())
}

func (m *DnsProviderServer) Timeout(ctx context.Context, req *dnspluginproto.TimeoutRequest) (*dnspluginproto.TimeoutResponse, error) {
	var timeout, interval time.Duration
	if pt, ok := m.provider.(challenge.ProviderTimeout); ok {
		timeout, interval = pt.Timeout()
	}

	return &dnspluginproto.TimeoutResponse{
		Timeout:  durationpb.New(timeout),
		Interval: durationpb.New(interval),
	}, nil
}

// helper function to map environment variables if set
func mapEnvironmentVariableValues(keyMapping map[string]string) {
	for key := range keyMapping {
		if value, ok := os.LookupEnv(key); ok {
			os.Setenv(keyMapping[key], value)
		}
	}
}
