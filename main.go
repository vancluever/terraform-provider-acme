package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vancluever/terraform-provider-acme/v2/acme"
)

func main() {
	initLegoLogger()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: acme.Provider,
	})
}
