package dnsplugin

import (
	"context"
	"fmt"
	"os"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/hashicorp/go-plugin"
	dnspluginproto "github.com/vancluever/terraform-provider-acme/v2/proto/dnsplugin/v1"
	"google.golang.org/grpc"
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

	// Set env before configuring provider
	for k, v := range req.GetConfig() {
		os.Setenv(k, v)
	}

	var err error
	s.provider, err = providerFunc()
	if err != nil {
		return nil, fmt.Errorf("erorr initializing provider: %w", err)
	}

	return &dnspluginproto.ConfigureResponse{}, nil
}

func (m *DnsProviderServer) Present(ctx context.Context, req *dnspluginproto.PresentRequest) (*dnspluginproto.PresentResponse, error) {
	return &dnspluginproto.PresentResponse{}, m.provider.Present(req.GetDomain(), req.GetToken(), req.GetKeyAuth())
}

func (m *DnsProviderServer) CleanUp(ctx context.Context, req *dnspluginproto.CleanUpRequest) (*dnspluginproto.CleanUpResponse, error) {
	return &dnspluginproto.CleanUpResponse{}, m.provider.CleanUp(req.GetDomain(), req.GetToken(), req.GetKeyAuth())
}
