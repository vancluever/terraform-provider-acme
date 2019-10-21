package storage

import (
	liquidweb "github.com/liquidweb/liquidweb-go"
	"github.com/liquidweb/liquidweb-go/types"
)

// Attachment represents the attachment details for a block volume.
type Attachment struct {
	Device   string `json:"device,omitempty"`
	Resource string `json:"resource,omitempty"`
}

// BlockVolumeParams is the set of parameters used when creating or updating a block volume
type BlockVolumeParams struct {
	Attach      string `json:"attach,omitempty"`
	CrossAttach bool   `json:"cross_attach,omitempty"`
	DetachFrom  string `json:"detach_from,omitempty"`
	Domain      string `json:"domain,omitempty"`
	NewSize     int    `json:"new_size,omitempty"`
	Region      int    `json:"region,omitempty"`
	Size        int    `json:"size,omitempty"`
	To          string `json:"to,omitempty"`
	UniqID      string `json:"uniq_id,omitempty"`
	Zone        int    `json:"zone,omitempty"`
	liquidweb.PageParams
}

// BlockVolume is the resource representing a block volume.
type BlockVolume struct {
	AttachedTo       []Attachment           `json:"attachedTo,omitempty"`
	CrossAttach      types.NumericalBoolean `json:"cross_attach,omitempty"`
	Domain           string                 `json:"domain,omitempty"`
	Label            string                 `json:"label,omitempty"`
	Size             types.FlexInt          `json:"size,omitempty"`
	Status           string                 `json:"status,omitempty"`
	UniqID           string                 `json:"uniq_id,omitempty"`
	ZoneAvailability []types.FlexInt        `json:"zoneAvailability,omitempty"`
}

// BlockVolumeList is an envelope for the API result containing either a list of block volumes or an error.
type BlockVolumeList struct {
	liquidweb.ListMeta
	Items []BlockVolume `json:"items,omitempty"`
}

// BlockVolumeDeletion represents the API result when deleting a block volume.
type BlockVolumeDeletion struct {
	Deleted string `json:"deleted"`
}

// BlockVolumeResize represents the API result when resizing a block volume.
type BlockVolumeResize struct {
	NewSize types.FlexInt `json:"new_size,omitempty"`
	OldSize types.FlexInt `json:"old_size,omitempty"`
	UniqID  string        `json:"uniq_id,omitempty"`
}
