package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-acme/acme"
)

func main() {
	initLegoLogger()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: acme.Provider,
	})
}
