package acme

import (
	"crypto/x509"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xenolf/lego/certificate"
)

func resourceACMECertificate() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,

		Schema:        certificateSchemaFull(),
		SchemaVersion: 3,
		MigrateState:  resourceACMECertificateMigrateState,
	}
}

func resourceACMECertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	if err = setDNSChallenge(client, d.Get("dns_challenge").([]interface{})[0].(map[string]interface{})); err != nil {
		return err
	}

	var cert *certificate.Resource

	if v, ok := d.GetOk("certificate_request_pem"); ok {
		var csr *x509.CertificateRequest
		csr, err = csrFromPEM([]byte(v.(string)))
		if err != nil {
			return err
		}
		cert, err = client.Certificate.ObtainForCSR(*csr, true)
	} else {
		cn := d.Get("common_name").(string)
		domains := []string{cn}
		if s, ok := d.GetOk("subject_alternative_names"); ok {
			for _, v := range stringSlice(s.(*schema.Set).List()) {
				if v == cn {
					return fmt.Errorf("common name %s should not appear in SAN list", v)
				}
				domains = append(domains, v)
			}
		}

		cert, err = client.Certificate.Obtain(certificate.ObtainRequest{
			Domains:    domains,
			Bundle:     true,
			MustStaple: d.Get("must_staple").(bool),
		})
	}

	if err != nil {
		return fmt.Errorf("error creating certificate: %s", err)
	}

	d.SetId(cert.CertURL)
	password := d.Get("certificate_p12_password").(string)
	if err := saveCertificateResource(d, cert, password); err != nil {
		return err
	}

	return resourceACMECertificateRead(d, meta)
}

// resourceACMECertificateRead is a noop. See
// resourceACMECertificateCustomizeDiff for most of the renewal check logic.
func resourceACMECertificateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceACMECertificateCustomizeDiff checks the certificate for renewal and
// flags it as NewComputed if it needs a renewal.
func resourceACMECertificateCustomizeDiff(d *schema.ResourceDiff, meta interface{}) error {
	// There's nothing for us to do in a Create diff, so if there's no ID yet,
	// just pass this part.
	if d.Id() == "" {
		return nil
	}

	mindays := d.Get("min_days_remaining").(int)
	if mindays < 0 {
		log.Printf("[WARN] min_days_remaining is set to less than 0, certificate will never be renewed")
		return nil
	}

	cert := expandCertificateResource(d)
	remaining, err := certDaysRemaining(cert)
	if err != nil {
		return err
	}

	if int64(mindays) >= remaining {
		d.SetNewComputed("certificate_pem")
	}

	return nil
}

// resourceACMECertificateUpdate renews a certificate if it has been flagged as changed.
func resourceACMECertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	// We don't need to do anything else here if the certificate hasn't been diffed
	if !d.HasChange("certificate_pem") {
		// when the certificate hasn't changed but the p12 password has, we still need to regenerate the p12
		if d.HasChange("certificate_p12_password") {
			cert := expandCertificateResource(d)
			d.SetId(cert.CertURL)
			password := d.Get("certificate_p12_password").(string)
			if err := saveCertificateResource(d, cert, password); err != nil {
				return err
			}
		}
		return nil
	}

	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	cert := expandCertificateResource(d)
	if err := setDNSChallenge(client, d.Get("dns_challenge").([]interface{})[0].(map[string]interface{})); err != nil {
		return err
	}

	newCert, err := client.Certificate.Renew(*cert, true, d.Get("must_staple").(bool))
	if err != nil {
		return err
	}

	d.SetId(newCert.CertURL)
	password := d.Get("certificate_p12_password").(string)
	if err := saveCertificateResource(d, newCert, password); err != nil {
		return err
	}

	return nil
}

// resourceACMECertificateDelete "deletes" the certificate by revoking it.
func resourceACMECertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	cert := expandCertificateResource(d)
	remaining, err := certSecondsRemaining(cert)
	if err != nil {
		return err
	}

	if remaining >= 0 {
		if err := client.Certificate.Revoke(cert.Certificate); err != nil {
			return err
		}
	}

	return nil
}
