package client

import (
	network "github.com/liquidweb/liquidweb-go/network"
	"github.com/liquidweb/liquidweb-go/storage"
	"github.com/liquidweb/liquidweb-go/storm"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkDNS          network.DNSBackend
	NetworkLoadBalancer network.LoadBalancerBackend
	NetworkVIP          network.VIPBackend
	NetworkZone         network.ZoneBackend
	StorageBlockVolume  storage.BlockVolumeBackend

	StormConfig storm.ConfigBackend
	StormServer storm.ServerBackend
}

// NewAPI is the API client for interacting with Storm.
func NewAPI(username string, password string, url string, timeout int) (*API, error) {
	config, err := NewConfig(username, password, url, timeout, true)
	if err != nil {
		return nil, err
	}

	// Initialize http backend
	client := NewClient(config)
	api := &API{
		NetworkDNS:          &network.DNSClient{Backend: client},
		NetworkLoadBalancer: &network.LoadBalancerClient{Backend: client},
		NetworkVIP:          &network.VIPClient{Backend: client},
		NetworkZone:         &network.ZoneClient{Backend: client},
		StorageBlockVolume:  &storage.BlockVolumeClient{Backend: client},
		StormConfig:         &storm.ConfigClient{Backend: client},
		StormServer:         &storm.ServerClient{Backend: client},
	}

	return api, nil
}
