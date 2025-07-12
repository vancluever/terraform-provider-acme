package acme

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceACMEServerURL() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceACMEServerURLRead,
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Description: "The server URL the provider is configured with.",
				Computed:    true,
			},
		},
	}
}

func dataSourceACMEServerURLRead(d *schema.ResourceData, meta any) error {
	d.SetId(meta.(*Config).ServerURL)
	d.Set("server_url", meta.(*Config).ServerURL)
	return nil
}
