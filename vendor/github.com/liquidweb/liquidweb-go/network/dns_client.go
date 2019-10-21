package network

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
)

// DNSBackend describes the interface for interactions with the API.
type DNSBackend interface {
	Create(*DNSRecordParams) (*DNSRecord, error)
	Details(int) (*DNSRecord, error)
	List(*DNSRecordParams) (*DNSRecordList, error)
	Update(*DNSRecordParams) (*DNSRecord, error)
	Delete(*DNSRecordParams) (*DNSRecordDeletion, error)
}

// DNSClient is the backend implementation for interacting with DNS Records.
type DNSClient struct {
	Backend liquidweb.Backend
}

// Create creates a new DNS Record.
func (c *DNSClient) Create(params *DNSRecordParams) (*DNSRecord, error) {
	var result DNSRecord
	err := c.Backend.Call("v1/Network/DNS/Record/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a DNS Record.
func (c *DNSClient) Details(id int) (*DNSRecord, error) {
	var result DNSRecord
	params := DNSRecordParams{ID: id}

	err := c.Backend.Call("v1/Network/DNS/Record/details", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a list of DNS Records.
func (c *DNSClient) List(params *DNSRecordParams) (*DNSRecordList, error) {
	list := &DNSRecordList{}

	err := c.Backend.Call("v1/Network/DNS/Record/list", params, list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Update will update a DNS Record.
func (c *DNSClient) Update(params *DNSRecordParams) (*DNSRecord, error) {
	var result DNSRecord
	err := c.Backend.Call("v1/Network/DNS/Record/update", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete will delete a DNS Record.
func (c *DNSClient) Delete(params *DNSRecordParams) (*DNSRecordDeletion, error) {
	var result DNSRecordDeletion
	err := c.Backend.Call("v1/Network/DNS/Record/delete", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
