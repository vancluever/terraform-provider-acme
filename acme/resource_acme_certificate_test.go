package acme

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"software.sslmate.com/src/go-pkcs12"
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

func TestAccACMECertificate_recursiveNameservers(t *testing.T) {
	f, err := newTestForwarder()
	if err != nil {
		t.Fatal(err)
	}

	defer f.Shutdown()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateRecursiveNameserversConfig(f.LocalAddr()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www10", "www11", false),
					f.Check(),
				),
			},
		},
	})
}

func TestAccACMECertificate_p12Password(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckCert(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigP12Password("changeit"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www12", "www13", false),
				),
			},
			{
				Config: testAccACMECertificateConfigP12Password("changeitagain"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www12", "www13", false),
				),
			},
		},
	})
}

func TestAccACMECertificate_multiProviders(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckCert(t)
			testAccPreCheckCertMultiProviders(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccACMECertificateConfigMultiProviders(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"acme_certificate.certificate", "id",
						"acme_certificate.certificate", "certificate_url",
					),
					testAccCheckACMECertificateValid("acme_certificate.certificate", "www14", "www15", false),
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

			// Test PKCS12, which is only present if there's a private key.
			if err := testFindPEMInP12(
				[]byte(rs.Primary.Attributes["certificate_p12"]),
				rs.Primary.Attributes["certificate_p12_password"],
				[]byte(cert),
				[]byte(issuer),
				[]byte(key),
			); err != nil {
				return fmt.Errorf("error validating P12 certificates: %s", err)
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

// testFindPEMInP12 tries to find the supplied PEM blocks in the supplied
// base64-encoded P12 content.
func testFindPEMInP12(pfxB64 []byte, password string, expected ...[]byte) error {
	pfxData := make([]byte, base64.StdEncoding.DecodedLen(len(pfxB64)))
	nBytes, err := base64.StdEncoding.Decode(pfxData, pfxB64)
	if err != nil {
		return err
	}

	actualBlocks, err := pkcs12.ToPEM(pfxData[:nBytes], password)
	if err != nil {
		return err
	}

	var expectedBlocks []*pem.Block
	for i, data := range expected {
		block, _ := pem.Decode(data)
		if block == nil {
			return fmt.Errorf("bad PEM data in expected block %d", i)
		}

		expectedBlocks = append(expectedBlocks, block)
	}

	for i := 0; i < len(expectedBlocks); i++ {
		expected := expectedBlocks[i]
		for _, actual := range actualBlocks {
			if reflect.DeepEqual(expected.Bytes, actual.Bytes) {
				expectedBlocks = append(expectedBlocks[:i], expectedBlocks[i+1:]...)
				i--
			}
		}
	}

	if len(expectedBlocks) > 0 {
		return fmt.Errorf(
			"not all expected blocks were found in the PFX archive (remaining: %d, %d in archive)",
			len(expectedBlocks),
			len(actualBlocks),
		)
	}

	return nil
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
		t.Skip("ACME_R53_ZONE_ID must be set for the static configuration certificate acceptance test")
	}
}

func testAccPreCheckCertMultiProviders(t *testing.T) {
	if v := os.Getenv("ACME_MULTI_PROVIDERS"); v == "" {
		t.Skip("ACME_MULTI_PROVIDERS must be set for the multiple providers certificate acceptance test")
	} else {
		providers := strings.Split(os.Getenv("ACME_MULTI_PROVIDERS"), ",")
		if len(providers) != 2 {
			t.Fatal("ACME_MULTI_PROVIDERS must specify exactly two providers")
		}
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

    config = {
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

func testAccACMECertificateRecursiveNameserversConfig(nameserver string) string {
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
  common_name               = "www10.${var.domain}"
  subject_alternative_names = ["www11.${var.domain}"]

  dns_challenge {
    provider = "route53"
  }

  recursive_nameservers = ["%s"]
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"), nameserver)
}

func testAccACMECertificateConfigP12Password(password string) string {
	return fmt.Sprintf(`
variable "email_address" {
  default = "%s"
}

variable "domain" {
  default = "%s"
}

variable "password" {
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
  common_name               = "www12.${var.domain}"
  subject_alternative_names = ["www13.${var.domain}"]
  certificate_p12_password  = "${var.password}"

  dns_challenge {
    provider = "route53"
  }
}
`,
		os.Getenv("ACME_EMAIL_ADDRESS"),
		os.Getenv("ACME_CERT_DOMAIN"),
		password,
	)
}

func testAccACMECertificateConfigMultiProviders() string {
	providers := strings.Split(os.Getenv("ACME_MULTI_PROVIDERS"), ",")
	if len(providers) < 2 {
		// This is a workaround just to make sure we don't get a panic
		// when the config is generated for the TestCase literal. This
		// test should be skipped or error out if ACME_MULTI_PROVIDERS is
		// not properly defiend.
		providers = make([]string, 2)
	}

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
  common_name               = "www14.${var.domain}"
  subject_alternative_names = ["www15.${var.domain}"]

  dns_challenge {
    provider = "%s"
  }

  dns_challenge {
    provider = "%s"
  }
}
`, os.Getenv("ACME_EMAIL_ADDRESS"), os.Getenv("ACME_CERT_DOMAIN"), providers[0], providers[1])
}
