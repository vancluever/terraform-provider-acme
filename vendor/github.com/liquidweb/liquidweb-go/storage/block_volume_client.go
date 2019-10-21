package storage

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
)

// BlockVolumeBackend describes the interface for interactions with the API.
type BlockVolumeBackend interface {
	Create(*BlockVolumeParams) (*BlockVolume, error)
	Details(string) (*BlockVolume, error)
	List(*BlockVolumeParams) (*BlockVolumeList, error)
	Update(*BlockVolumeParams) (*BlockVolume, error)
	Delete(string) (*BlockVolumeDeletion, error)
	Resize(*BlockVolumeParams) (*BlockVolumeResize, error)
}

// BlockVolumeClient is the backend implementation for interacting with block volumes.
type BlockVolumeClient struct {
	Backend liquidweb.Backend
}

// Create creates a new block volume.
func (c *BlockVolumeClient) Create(params *BlockVolumeParams) (*BlockVolume, error) {
	var result BlockVolume
	err := c.Backend.Call("v1/Storage/Block/Volume/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a block volume.
func (c *BlockVolumeClient) Details(id string) (*BlockVolume, error) {
	var result BlockVolume
	params := BlockVolumeParams{UniqID: id}

	err := c.Backend.Call("v1/Storage/Block/Volume/details", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a list of block volumes.
func (c *BlockVolumeClient) List(params *BlockVolumeParams) (*BlockVolumeList, error) {
	list := &BlockVolumeList{}

	err := c.Backend.Call("v1/Storage/Block/Volume/list", params, list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// Update will update a block volume.
func (c *BlockVolumeClient) Update(params *BlockVolumeParams) (*BlockVolume, error) {
	var result BlockVolume
	err := c.Backend.Call("v1/Storage/Block/Volume/update", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete will delete a block volume.
func (c *BlockVolumeClient) Delete(id string) (*BlockVolumeDeletion, error) {
	var result BlockVolumeDeletion
	params := BlockVolumeParams{UniqID: id}

	err := c.Backend.Call("v1/Storage/Block/Volume/delete", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Resize will resize a block volume.
func (c *BlockVolumeClient) Resize(params *BlockVolumeParams) (*BlockVolumeResize, error) {
	var result BlockVolumeResize
	err := c.Backend.Call("v1/Storage/Block/Volume/resize", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
