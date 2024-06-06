package acme

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccACMEServerURL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMEServerURLConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.acme_server_url.url",
						"id",
						pebbleDirBasic,
					),
					resource.TestCheckResourceAttr(
						"data.acme_server_url.url",
						"server_url",
						pebbleDirBasic,
					),
					resource.TestCheckOutput(
						"server_url",
						pebbleDirBasic,
					),
				),
			},
		},
	})
}

func testAccACMEServerURLConfig() string {
	return fmt.Sprintf(`
provider "acme" {
  server_url = "%s"
}

data "acme_server_url" "url" {}

output "server_url" {
  value = data.acme_server_url.url.server_url
}
`, pebbleDirBasic)
}
