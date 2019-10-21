package network

import (
	"github.com/liquidweb/liquidweb-go/types"
)

// VIPParams is the set of parameters used when creating or updating a VIP.
type VIPParams struct {
	Domain     string `json:"domain,omitempty"`
	Zone       int    `json:"zone,omitempty"`
	UniqID     string `json:"uniq_id,omitempty"`
	SubAccount string `json:"subaccnt,omitempty"`
}

// VIP is the resource representing a VIP entry.
type VIP struct {
	Domain       string                 `json:"domain,omitempty"`
	Active       types.NumericalBoolean `json:"active,omitempty"`
	ActiveStatus string                 `json:"activeStatus,omitempty"`
	UniqID       string                 `json:"uniq_id,omitempty"`
	Destroyed    string                 `json:"destroyed,omitempty"`
	IP           string                 `json:"ip,omitempty"`
}

// VIPDeletion represents the API result when deleting a VIP.
type VIPDeletion struct {
	Destroyed string `json:"destroyed"`
}

// VIPNewStatus represents a VIP's status when new.
const VIPNewStatus = "New"

// VIPActiveStatus represents a VIP's status when active.
const VIPActiveStatus = "Active"

// VIPDisabledStatus represents a VIP's status when disabled.
const VIPDisabledStatus = "Disabled"

// VIPTerminatedStatus represents a VIP's status when terminated.
const VIPTerminatedStatus = "Terminated"

// VIPPendingTermination represents a VIP's status when termination is pending.
const VIPPendingTermination = "Pending-Termination"

// VIPPendingActivation represents a VIP's status when activation is pending.
const VIPPendingActivation = "Pending-Activation"

// VIPPendingPayment represents a VIP's status when payment is pending.
const VIPPendingPayment = "Pending-Payment"

// VIPBucketPart represents a VIP's status when termination is pending.
const VIPBucketPart = "BucketPart"

// VIPPendingConfig represents a VIP's status when configuration is pending.
const VIPPendingConfig = "Pending-Config"

// PendingStatuses is an array of strings representing the different statuses a
// VIP can be in before it is active.
var PendingStatuses = []string{
	VIPNewStatus,
	VIPPendingTermination,
	VIPPendingActivation,
	VIPPendingPayment,
	VIPPendingConfig,
}
