package acme

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccACMERegistration_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckReg(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckACMERegistrationValid("acme_registration.reg", false),
		Steps: []resource.TestStep{
			{
				Config: testAccACMERegistrationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_registration.reg", "id",
						"acme_registration.reg", "registration_url",
					),
					testAccCheckACMERegistrationValid("acme_registration.reg", true),
				),
			},
		},
	})
}

func TestAccACMERegistration_eab(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckACMERegistrationValid("acme_registration.reg", false),
		Steps: []resource.TestStep{
			{
				Config: testAccACMERegistrationConfigPebble,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_registration.reg", "id",
						"acme_registration.reg", "registration_url",
					),
					testAccCheckACMERegistrationValid("acme_registration.reg", true),
				),
			},
		},
	})
}

func TestAccACMERegistration_refreshDeactivated(t *testing.T) {
	var state *terraform.State
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckReg(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMERegistrationConfig(),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						state = s
						return nil
					},
					resource.TestCheckResourceAttrPair(
						"acme_registration.reg", "id",
						"acme_registration.reg", "registration_url",
					),
					testAccCheckACMERegistrationValid("acme_registration.reg", true),
				),
			},
			{
				PreConfig: func() {
					rs := state.RootModule().Resources["acme_registration.reg"]
					d := testAccCheckACMERegistrationResourceData(rs)
					client, _, err := expandACMEClient(d, testAccProvider.Meta(), true)
					if err != nil {
						panic(err)
					}

					if err := client.Registration.DeleteRegistration(); err != nil {
						panic(err)
					}
				},
				Config:             testAccACMERegistrationConfig(),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckACMERegistrationValid(n string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find ACME registration: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ACME registration ID not set")
		}

		d := testAccCheckACMERegistrationResourceData(rs)

		client, _, err := expandACMEClient(d, testAccProvider.Meta(), true)
		if err != nil {
			if regGone(err) && !exists {
				return nil
			}
			return fmt.Errorf("Could not build ACME client off reg: %s", err.Error())
		}

		reg, err := client.Registration.QueryRegistration()
		if err != nil {
			return fmt.Errorf("Error on reg query: %s", err.Error())
		}

		actual := reg.URI
		expected := rs.Primary.ID

		if actual != expected {
			return fmt.Errorf("Expected ID to be %s, got %s", expected, actual)
		}
		return nil
	}
}

// testAccCheckACMERegistrationResourceData returns a *schema.ResourceData that should match a
// acme_registration resource.
func testAccCheckACMERegistrationResourceData(rs *terraform.ResourceState) *schema.ResourceData {
	r := resourceACMERegistration()
	d := r.TestResourceData()

	d.SetId(rs.Primary.ID)
	d.Set("account_key_pem", rs.Primary.Attributes["account_key_pem"])
	d.Set("email_address", rs.Primary.Attributes["email_address"])

	return d
}

func testAccPreCheckReg(t *testing.T) {
	if v := os.Getenv("ACME_EMAIL_ADDRESS"); v == "" {
		t.Fatal("ACME_EMAIL_ADDRESS must be set for the registration acceptance test")
	}
}

func testAccACMERegistrationConfig() string {
	return fmt.Sprintf(`
resource "tls_private_key" "private_key" {
    algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "%s"
}
`, os.Getenv("ACME_EMAIL_ADDRESS"))
}

const testAccACMERegistrationConfigPebble = `
provider "acme" {
  server_url = "https://127.0.0.1:14001/dir"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.test"
  external_account_binding {
    key_id      = "kid-1"
    hmac_base64 = "zWNDZM6eQGHWpSRTPal5eIUYFTu7EajVIoguysqZ9wG44nMEtx3MUAsUDkMTQ12W"
  }
}
`
