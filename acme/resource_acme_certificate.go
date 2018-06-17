package acme

import (
	"crypto/x509"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xenolf/lego/acme"
)

func resourceACMECertificate() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,

		Schema:        certificateSchemaFull(),
		SchemaVersion: 1,
		MigrateState:  resourceACMECertificateMigrateState,
	}
}

func resourceACMECertificateCreate(d *schema.ResourceData, meta interface{}) error {
	// Turn on partial state to ensure that nothing is recorded until we want it to be.
	d.Partial(true)

	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	if err = setDNSChallenge(client, d.Get("dns_challenge").(*schema.Set).List()[0].(map[string]interface{})); err != nil {
		return err
	}

	var cert *acme.CertificateResource

	if v, ok := d.GetOk("certificate_request_pem"); ok {
		var csr *x509.CertificateRequest
		csr, err = csrFromPEM([]byte(v.(string)))
		if err != nil {
			return err
		}
		cert, err = client.ObtainCertificateForCSR(*csr, true)
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

		cert, err = client.ObtainCertificate(domains, true, nil, d.Get("must_staple").(bool))
	}

	if err != nil {
		return fmt.Errorf("error creating certificate: %s", err)
	}

	// Done! save the cert
	d.Partial(false)
	d.SetId(cert.CertURL)
	saveCertificateResource(d, cert)

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
	// We use partial state to protect against losing the certificate during bad
	// renewal. min_days_remaining is a safe change to record in the state
	// however, so we allow that to be set even on error.
	d.Partial(true)
	d.SetPartial("min_days_remaining")

	// We don't need to do anything else here if the certificate hasn't been diffed
	if !d.HasChange("certificate_pem") {
		return nil
	}

	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	cert := expandCertificateResource(d)
	if err := setDNSChallenge(client, d.Get("dns_challenge").(*schema.Set).List()[0].(map[string]interface{})); err != nil {
		return err
	}
	newCert, err := client.RenewCertificate(*cert, true, d.Get("must_staple").(bool))
	if err != nil {
		return err
	}

	// Now safe to record state
	d.Partial(false)
	saveCertificateResource(d, newCert)

	return nil
}

// resourceACMECertificateDelete "deletes" the certificate by revoking it.
func resourceACMECertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	cert, ok := d.GetOk("certificate_pem")

	if ok {
		err = client.RevokeCertificate([]byte(cert.(string)))
		if err != nil {
			// Ignore conflict (409) responses, as certificate is already revoked.
			if rerr, ok := err.(acme.RemoteError); !ok || rerr.StatusCode != 409 {
				return err
			}
		}
	}

	return nil
}
