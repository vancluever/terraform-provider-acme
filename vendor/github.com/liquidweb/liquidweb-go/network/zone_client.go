package network

import liquidweb "github.com/liquidweb/liquidweb-go"

// ZoneBackend is the interface for network zones.
type ZoneBackend interface {
	Details(int) (*Zone, error)
	List(*ZoneListParams) (*ZoneList, error)
}

// ZoneClient is the API client for network zones.
type ZoneClient struct {
	Backend liquidweb.Backend
}

// Details fetches the details for a zone.
func (c *ZoneClient) Details(id int) (*Zone, error) {
	var zoneResult *Zone
	zoneParams := ZoneParams{ID: id}

	err := c.Backend.Call("v1/Network/Zone/detail", zoneParams, zoneResult)
	if err != nil {
		return nil, err
	}
	return zoneResult, nil
}

// List returns a list of network zones.
func (c *ZoneClient) List(params *ZoneListParams) (*ZoneList, error) {
	zoneList := &ZoneList{}

	err := c.Backend.Call("v1/Network/Zone/list", params, zoneList)
	if err != nil {
		return nil, err
	}

	return zoneList, nil
}
