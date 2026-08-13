package acme

import (
	"errors"

	"github.com/go-acme/lego/v5/certificate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceDataOrDiff is a simple interface to allow us to use the Get
// function that is in ResourceData and ResourceDiff under the same function.
type resourceDataOrDiff interface {
	Get(string) any
	GetOk(string) (any, bool)
	GetChange(string) (any, any)
}

// saveCertificateResource takes an certificate.Resource and sets fields.
func saveCertificateResource(d *schema.ResourceData, cert *certificate.Resource, password string) error {
	d.Set("certificate_url", cert.CertURL)
	if len(cert.Domains) == 0 {
		return errors.New("certificate has no domains")
	}
	d.Set("certificate_domain", cert.Domains[0])
	d.Set("private_key_pem", string(cert.PrivateKey))
	issued, issuedNotBefore, issuedNotAfter, issuedSerial, issuer, err := splitPEMBundle(cert.Certificate)
	if err != nil {
		return err
	}

	d.Set("certificate_pem", string(issued))
	d.Set("issuer_pem", string(issuer))
	d.Set("certificate_not_before", issuedNotBefore)
	d.Set("certificate_not_after", issuedNotAfter)
	d.Set("certificate_serial", issuedSerial)

	// Set PKCS12 data. This is only set if there is a private key
	// present.
	if len(cert.PrivateKey) > 0 {
		pfxB64, err := bundleToPKCS12(cert.Certificate, cert.PrivateKey, password)
		if err != nil {
			return err
		}

		d.Set("certificate_p12", string(pfxB64))
	} else {
		d.Set("certificate_p12", "")
	}

	return nil
}

// expandCertificateResource takes saved state in the certificate resource
// and returns an certificate.Resource.
func expandCertificateResource(d resourceDataOrDiff) *certificate.Resource {
	cert := &certificate.Resource{
		Domains: []string{d.Get("certificate_domain").(string)},
		CertURL: d.Get("certificate_url").(string),
	}

	// Only populate the PrivateKey or CSR fields if we have them
	if pk, ok := d.GetOk("private_key_pem"); ok {
		cert.PrivateKey = []byte(pk.(string))
	}
	if csr, ok := d.GetOk("certificate_request_pem"); ok {
		cert.CSR = []byte(csr.(string))
	}

	// There are situations now where the new certificate may be blank, which
	// signifies that the certificate needs to be renewed. In this case, we need
	// the old value here, versus the new one.
	oldCertPEM, newCertPEM := d.GetChange("certificate_pem")
	issuerPEM := d.Get("issuer_pem")
	if newCertPEM.(string) != "" {
		cert.Certificate = []byte(newCertPEM.(string) + issuerPEM.(string))
	} else {
		cert.Certificate = []byte(oldCertPEM.(string) + issuerPEM.(string))
	}
	return cert
}
