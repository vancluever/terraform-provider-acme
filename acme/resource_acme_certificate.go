package acme

import (
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"github.com/go-acme/lego/certificate"
	"github.com/go-acme/lego/challenge"
	"github.com/go-acme/lego/challenge/dns01"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
)

// DNSProviderWrapper is a multi-provider wrapper to support multiple
// DNS challenges.
type DNSProviderWrapper struct {
	providers []challenge.Provider
}

// NewDNSProviderWrapper returns an freshly initialized
// DNSProviderWrapper.
func NewDNSProviderWrapper() (*DNSProviderWrapper, error) {
	return &DNSProviderWrapper{}, nil
}

// Present implements challenge.Provider for DNSProviderWrapper.
func (d *DNSProviderWrapper) Present(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.providers {
		err = p.Present(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while presenting token for DNS challenge: %s", err.Error()))
		}
	}

	return err
}

// CleanUp implements challenge.Provider for DNSProviderWrapper.
func (d *DNSProviderWrapper) CleanUp(domain, token, keyAuth string) error {
	var err error
	for _, p := range d.providers {
		err = p.CleanUp(domain, token, keyAuth)
		if err != nil {
			err = multierror.Append(err, fmt.Errorf("error encountered while cleaning token for DNS challenge: %s", err.Error()))
		}
	}

	return err
}

// Timeout implements challenge.ProviderTimeout for
// DNSProviderWrapper.
//
// The highest polling interval and timeout values defined across all
// providers is used.
func (d *DNSProviderWrapper) Timeout() (time.Duration, time.Duration) {
	var timeout, interval time.Duration
	for _, p := range d.providers {
		if pt, ok := p.(challenge.ProviderTimeout); ok {
			t, i := pt.Timeout()
			if t > timeout {
				timeout = t
			}

			if i > interval {
				interval = i
			}
		}
	}

	if timeout < 1 {
		timeout = dns01.DefaultPropagationTimeout
	}

	if interval < 1 {
		interval = dns01.DefaultPollingInterval
	}

	return timeout, interval
}

func resourceACMECertificate() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,

		Schema:        certificateSchemaFull(),
		SchemaVersion: 4,
		MigrateState:  resourceACMECertificateMigrateState,
	}
}

func resourceACMECertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	provider, err := NewDNSProviderWrapper()
	if err != nil {
		return err
	}

	for _, v := range d.Get("dns_challenge").([]interface{}) {
		if p, err := setDNSChallenge(client, v.(map[string]interface{})); err == nil {
			provider.providers = append(provider.providers, p)
		} else {
			return err
		}
	}

	var opts []dns01.ChallengeOption
	if nameservers := d.Get("recursive_nameservers").([]interface{}); len(nameservers) > 0 {
		var s []string
		for _, ns := range nameservers {
			s = append(s, ns.(string))
		}

		opts = append(opts, dns01.AddRecursiveNameservers(s))
	}

	if err := client.Challenge.SetDNS01Provider(provider, opts...); err != nil {
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
	// Ensure duplicate providers for dns_challenge are not provided.
	providerMap := make(map[string]bool)
	for _, v := range d.Get("dns_challenge").([]interface{}) {
		m := v.(map[string]interface{})
		if v, ok := m["provider"]; ok && v.(string) != "" {
			provider := v.(string)
			if _, ok := providerMap[provider]; ok {
				return fmt.Errorf("duplicate dns_challenge providers: %s", provider)
			}

			providerMap[provider] = true
		} else {
			return fmt.Errorf("DNS challenge provider not defined")
		}
	}

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

	provider, err := NewDNSProviderWrapper()
	if err != nil {
		return err
	}

	for _, v := range d.Get("dns_challenge").([]interface{}) {
		if p, err := setDNSChallenge(client, v.(map[string]interface{})); err == nil {
			provider.providers = append(provider.providers, p)
		} else {
			return err
		}
	}

	var opts []dns01.ChallengeOption
	if nameservers := d.Get("recursive_nameservers").([]interface{}); len(nameservers) > 0 {
		var s []string
		for _, ns := range nameservers {
			s = append(s, ns.(string))
		}

		opts = append(opts, dns01.AddRecursiveNameservers(s))
	}

	if err := client.Challenge.SetDNS01Provider(provider, opts...); err != nil {
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
