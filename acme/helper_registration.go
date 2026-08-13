package acme

import (
	"github.com/go-acme/lego/v5/acme"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// saveACMERegistration takes an *registration.Resource and sets the appropriate fields
// for a registration resource.
func saveACMERegistration(d *schema.ResourceData, reg *acme.ExtendedAccount) error {
	d.Set("registration_url", reg.Location)

	return nil
}
