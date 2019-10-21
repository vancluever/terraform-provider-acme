package network

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
)

// VIPBackend describes the interface for interactions with the API.
type VIPBackend interface {
	Create(VIPParams) (*VIP, error)
	Destroy(string) (*VIPDeletion, error)
	Details(string) (*VIP, error)
}

// VIPClient is the backend implementation for interacting with VIP.
type VIPClient struct {
	Backend liquidweb.Backend
}

// Create creates a new VIP.
func (c *VIPClient) Create(params VIPParams) (*VIP, error) {
	var result VIP
	err := c.Backend.Call("v1/VIP/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a VIP.
func (c *VIPClient) Details(uniqID string) (*VIP, error) {
	var result VIP
	params := VIPParams{UniqID: uniqID}

	err := c.Backend.Call("v1/VIP/details", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Destroy will delete a VIP.
func (c *VIPClient) Destroy(uniqID string) (*VIPDeletion, error) {
	var result VIPDeletion
	params := VIPParams{UniqID: uniqID}

	err := c.Backend.Call("v1/VIP/destroy", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
