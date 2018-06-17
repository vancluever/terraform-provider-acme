package acme

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns the terraform.ResourceProvider structure for the ACME
// provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACME_SERVER_URL", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"acme_registration": resourceACMERegistration(),
			"acme_certificate":  resourceACMECertificate(),
		},

		ConfigureFunc: configureProvider,
	}
}

// Config represents the configuration of the provider.
type Config struct {
	// The ACME server URL.
	ServerURL string
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return &Config{
		ServerURL: d.Get("server_url").(string),
	}, nil
}
