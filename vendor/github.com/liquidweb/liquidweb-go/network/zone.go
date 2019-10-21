package network

import (
	liquidweb "github.com/liquidweb/liquidweb-go"

	"github.com/liquidweb/liquidweb-go/types"
)

// Zone is a grouping of network resources.
type Zone struct {
	ID             types.FlexInt                     `json:"id,omitempty"`
	IsDefault      types.NumericalBoolean            `json:"is_default,omitempty"`
	Name           string                            `json:"name,omitempty"`
	Region         ZoneRegion                        `json:"region,omitempty"`
	Status         string                            `json:"status,omitempty"`
	ValidSourceHVS map[string]types.NumericalBoolean `json:"valid_source_hvs,omitempty"`
}

// ZoneRegion describes the zone the resources should be in.
type ZoneRegion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ZoneParams are the set of parameters you can pass to the API for Network Zones.
type ZoneParams struct {
	ID int `json:"id,omitempty"`
}

// ZoneListParams are the set of parameters you can pass to the API for listing Network Zones.
type ZoneListParams struct {
	PageNum  int    `json:"page_num,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	Region   string `json:"region,omitempty"`
}

// ZoneList is an envelope for the API result containing either a list of zones or an error.
type ZoneList struct {
	liquidweb.ListMeta
	Items []Zone `json:"items,omitempty"`
}
