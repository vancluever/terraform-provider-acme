package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vancluever/terraform-provider-acme/acme"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: acme.Provider,
	})
}
