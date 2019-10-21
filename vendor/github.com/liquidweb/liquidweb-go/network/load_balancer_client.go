package network

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
)

// LoadBalancerBackend describes the interface for interactions with the API.
type LoadBalancerBackend interface {
	Create(LoadBalancerParams) (*LoadBalancer, error)
	Details(string) (*LoadBalancer, error)
	Update(LoadBalancerParams) (*LoadBalancer, error)
	Delete(string) (*LoadBalancerDeletion, error)
}

// LoadBalancerClient is the backend implementation for interacting with the API.
type LoadBalancerClient struct {
	Backend liquidweb.Backend
}

// Create creates a new load balancer.
func (c *LoadBalancerClient) Create(params LoadBalancerParams) (*LoadBalancer, error) {
	var result LoadBalancer
	err := c.Backend.Call("v1/Network/LoadBalancer/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a load balancer.
func (c *LoadBalancerClient) Details(uniqID string) (*LoadBalancer, error) {
	var result LoadBalancer
	params := LoadBalancerParams{UniqID: uniqID}

	err := c.Backend.Call("v1/Network/LoadBalancer/details", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update will update a load balancer.
func (c *LoadBalancerClient) Update(params LoadBalancerParams) (*LoadBalancer, error) {
	var result LoadBalancer
	err := c.Backend.Call("v1/Network/LoadBalancer/update", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete will delete a load balancer.
func (c *LoadBalancerClient) Delete(uniqID string) (*LoadBalancerDeletion, error) {
	var result LoadBalancerDeletion
	params := LoadBalancerParams{UniqID: uniqID}

	err := c.Backend.Call("v1/Network/LoadBalancer/delete", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
