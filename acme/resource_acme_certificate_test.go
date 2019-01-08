package acme

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/asn1"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// Constants for OCSP must staple
var (
	tlsFeatureExtensionOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 24}
	ocspMustStapleFeature  = []byte{0x30, 0x03, 0x02, 0x01, 0x05}
	envKeys                = []string{
		"AWS_PROFILE",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	}
)

func TestAccACMECertificate_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www", "www2", false),
				),
			},
		},
	})
}

func TestAccACMECertificate_CSR(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateCSRConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www3", "www4", false),
				),
			},
		},
	})
}

func TestAccACMECertificate_withDNSProviderConfig(t *testing.T) {
	// Cache credentials first and then restore them after the function ends. We
	// actually clear them after our pre-check so don't worry about that here.
	envCache := make(map[string]string)
	for _, k := range envKeys {
		envCache[k] = os.Getenv(k)
	}
	defer func() {
		for _, k := range envKeys {
			os.Setenv(k, envCache[k])
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckCert(t)
			testAccPreCheckCertZoneID(t)
			for _, k := range envKeys {
				os.Unsetenv(k)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateWithDNSProviderConfig(envCache),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www5", "", false),
				),
			},
		},
	})
}

func TestAccACMECertificate_forceRenewal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateForceRenewalConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www6", "", false),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccACMECertificateForceRenewalConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www6", "", false),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccACMECertificate_mustStaple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateMustStapleConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www8", "www9", true),
				),
			},
		},
	})
}

func TestAccACMECertificate_wildcard(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateWildcardConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckACMECertificateValid("acme_certificate.certificate", "*", "", false),
				),
			},
		},
	})
}

func testAccCheckACMECertificateValid(n, cn, san string, mustStaple bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find ACME certificate: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ACME certificate ID not set")
		}

		cert := rs.Primary.Attributes["certificate_pem"]
		issuer := rs.Primary.Attributes["issuer_pem"]
		key := rs.Primary.Attributes["private_key_pem"]
		x509Certs, err := parsePEMBundle([]byte(cert))
		if err != nil {
			return err
		}
		x509Cert := x509Certs[0]

		issuerCerts, err := parsePEMBundle([]byte(issuer))
		if err != nil {
			return err
		}
		issuerCert := issuerCerts[0]

		// Skip the private key test if we have an empty key. This is a legit case
		// that comes up when a CSR is supplied instead of creating a cert from
		// scratch.
		if key != "" {
			privateKey, err := privateKeyFromPEM([]byte(key))
			if err != nil {
				return err
			}

			var privPub crypto.PublicKey

			switch v := privateKey.(type) {
			case *rsa.PrivateKey:
				privPub = v.Public()
			case *ecdsa.PrivateKey:
				privPub = v.Public()
			}

			if reflect.DeepEqual(x509Cert.PublicKey, privPub) != true {
				return fmt.Errorf("Public key for cert and private key don't match: %#v, %#v", x509Cert.PublicKey, privPub)
			}
		}

		// Ensure the issuer cert is a CA cert
		if issuerCert.IsCA == false {
			return fmt.Errorf("issuer_pem is not a CA certificate")
		}

		// domains
		domain := "." + os.Getenv("ACME_CERT_DOMAIN")
		expectedCN := cn + domain
		var expectedSANs []string
		if san != "" {
			expectedSANs = []string{cn + domain, san + domain}
		} else {
			expectedSANs = []string{cn + domain}
		}

		actualCN := x509Cert.Subject.CommonName
		actualSANs := x509Cert.DNSNames

		if expectedCN != actualCN {
			return fmt.Errorf("Expected common name to be %s, got %s", expectedCN, actualCN)
		}

		if reflect.DeepEqual(expectedSANs, actualSANs) != true {
			return fmt.Errorf("Expected SANs to be %#v, got %#v", expectedSANs, actualSANs)
		}

		if mustStaple {
			for _, v := range x509Cert.Extensions {
				if reflect.DeepEqual(v.Id, tlsFeatureExtensionOID) && reflect.DeepEqual(v.Value, ocspMustStapleFeature) {
					goto stapleOK
				}
			}
			return fmt.Errorf("Did not find OCSP Stapling Required extension when expected")
		}

	stapleOK:

		return nil
	}
}

func testAccPreCheckCert(t *testing.T) {
	if v := os.Getenv("ACME_EMAIL_ADDRESS"); v == "" {
		t.Fatal("ACME_EMAIL_ADDRESS must be set for the certificate acceptance test")
	}
	if v := os.Getenv("ACME_CERT_DOMAIN"); v == "" {
		t.Fatal("ACME_CERT_DOMAIN must be set for the certificate acceptance test")
	}
	if v := os.Getenv("AWS_PROFILE"); v == "" {
		if v := os.Getenv("AWS_ACCESS_KEY_ID"); v == "" {
			t.Fatal("AWS_ACCESS_KEY_ID must be set for the certificate acceptance test")
		}
		if v := os.Getenv("AWS_SECRET_ACCESS_KEY"); v == "" {
			t.Fatal("AWS_SECRET_ACCESS_KEY must be set for the certificate acceptance test")
		}
	}
	if v := os.Getenv("AWS_DEFAULT_REGION"); v == "" {
		log.Println("[INFO] Test: Using us-west-2 as test region")
		os.Setenv("AWS_DEFAULT_REGION", "us-west-2")
	}
}

func testAccPreCheckCertZoneID(t *testing.T) {
	if v := os.Getenv("ACME_R53_ZONE_ID"); v == "" {
		t.Fatal("ACME_R53_ZONE_ID must be set for the static configuration certificate acceptance test")
	}
}

func testAccACMECertificateConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "www.${var.domain}"
  subject_alternative_names = ["www2.${var.domain}"]

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}

func testAccACMECertificateCSRConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "reg_private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.reg_private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "tls_private_key" "cert_private_key" {
  algorithm = "RSA"
}

resource "tls_cert_request" "req" {
  key_algorithm   = "RSA"
  private_key_pem = "${tls_private_key.cert_private_key.private_key_pem}"
  dns_names       = ["www3.${var.domain}", "www4.${var.domain}"]

  subject {
    common_name  = "www3.${var.domain}"
  }
}

resource "acme_certificate" "certificate" {
  account_key_pem         = "${acme_registration.reg.account_key_pem}"
  certificate_request_pem = "${tls_cert_request.req.cert_request_pem}"

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}

func testAccACMECertificateWithDNSProviderConfig(params map[string]string) string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem = "${acme_registration.reg.account_key_pem}"
  common_name     = "www5.${var.domain}"

  dns_challenge {
    provider = "route53"

    config {
      AWS_PROFILE           = "%s"
      AWS_ACCESS_KEY_ID     = "%s"
      AWS_SECRET_ACCESS_KEY = "%s"
      AWS_SESSION_TOKEN     = "%s"
      AWS_HOSTED_ZONE_ID    = "%s"
    }
  }
}
`,
		os.Getenv("ACME_EMAIL_ADDRESS"),
		os.Getenv("ACME_CERT_DOMAIN"),
		params["AWS_PROFILE"],
		params["AWS_ACCESS_KEY_ID"],
		params["AWS_SECRET_ACCESS_KEY"],
		params["AWS_SESSION_TOKEN"],
		os.Getenv("ACME_R53_ZONE_ID"),
	)
}

func testAccACMECertificateForceRenewalConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem    = "${acme_registration.reg.account_key_pem}"
  common_name        = "www6.${var.domain}"
  min_days_remaining = 720

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}

func testAccACMECertificateECKeyCertConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem = "${acme_registration.reg.account_key_pem}"
  common_name     = "www7.${var.domain}"

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}

func testAccACMECertificateMustStapleConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = "${acme_registration.reg.account_key_pem}"
  common_name               = "www8.${var.domain}"
  subject_alternative_names = ["www9.${var.domain}"]
  must_staple               = true

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}

func testAccACMECertificateWildcardConfig() string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "${var.email_address}"
}

resource "acme_certificate" "certificate" {
  account_key_pem = "${acme_registration.reg.account_key_pem}"
  common_name     = "*.${var.domain}"

  dns_challenge {
    provider = "route53"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"))
}
