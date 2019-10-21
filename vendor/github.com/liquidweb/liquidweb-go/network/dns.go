package network

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
	"github.com/liquidweb/liquidweb-go/types"
)

// DNSRecordParams is the set of parameters used when creating or updating a DNS record.
type DNSRecordParams struct {
	ID              int              `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Prio            int              `json:"prio,omitempty"`
	RData           string           `json:"rdata,omitempty"`
	TTL             int              `json:"ttl,omitempty"`
	Type            string           `json:"type,omitempty"`
	Zone            string           `json:"zone,omitempty"`
	ZoneID          int              `json:"zone_id,omitempty"`
	AdminEmail      string           `json:"adminEmail,omitempty"`
	Created         string           `json:"created,omitempty"`
	Exchange        string           `json:"exchange,omitempty"`
	Expiry          int              `json:"expiry,omitempty"`
	FullData        string           `json:"fullData,omitempty"`
	LastUpdated     string           `json:"last_updated,omitempty"`
	Minimum         int              `json:"minimum,omitempty"`
	Nameserver      string           `json:"nameserver,omitempty"`
	Port            int              `json:"port,omitempty"`
	RefreshInterval int              `json:"refreshInterval,omitempty"`
	RegionOverrides *RegionOverrides `json:"regionOverrides,omitempty"`
	Retry           int              `json:"retry,omitempty"`
	Serial          int              `json:"serial,omitempty"`
	Target          string           `json:"target,omitempty"`
	Weight          int              `json:"weight,omitempty"`
}

// RegionOverrides contains region data.
type RegionOverrides struct {
	RData    string
	Region   string
	RegionID int
}

// DNSRecord is the resource representing a DNS record entry.
type DNSRecord struct {
	ID              types.FlexInt    `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Prio            types.FlexInt    `json:"prio,omitempty"`
	RData           string           `json:"rdata,omitempty"`
	TTL             types.FlexInt    `json:"ttl,omitempty"`
	Type            string           `json:"type,omitempty"`
	Zone            string           `json:"zone,omitempty"`
	ZoneID          types.FlexInt    `json:"zone_id,omitempty"`
	AdminEmail      string           `json:"adminEmail,omitempty"`
	Created         string           `json:"created,omitempty"`
	Exchange        string           `json:"exchange,omitempty"`
	Expiry          types.FlexInt    `json:"expiry,omitempty"`
	FullData        string           `json:"fullData,omitempty"`
	LastUpdated     string           `json:"last_updated,omitempty"`
	Minimum         types.FlexInt    `json:"minimum,omitempty"`
	Nameserver      string           `json:"nameserver,omitempty"`
	Port            types.FlexInt    `json:"port,omitempty"`
	RefreshInterval types.FlexInt    `json:"refreshInterval,omitempty"`
	RegionOverrides *RegionOverrides `json:"regionOverrides,omitempty"`
	Retry           types.FlexInt    `json:"retry,omitempty"`
	Serial          types.FlexInt    `json:"serial,omitempty"`
	Target          string           `json:"target,omitempty"`
	Weight          types.FlexInt    `json:"weight,omitempty"`
}

// DNSRecordList is an envelope for the API result containing either a list of DNS Records or an error.
type DNSRecordList struct {
	liquidweb.ListMeta
	Items []DNSRecord `json:"items,omitempty"`
}

// DNSRecordDeletion represents the API result when deleting a DNS Record.
type DNSRecordDeletion struct {
	Deleted types.FlexInt `json:"deleted"`
}
