package acme

import (
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"github.com/go-acme/lego/v3/certificate"
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/challenge/dns01"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceACMECertificate returns the current version of the
// acme_registration resource and needs to be updated when the schema
// version is incremented.
func resourceACMECertificate() *schema.Resource { return resourceACMECertificateV4() }

func resourceACMECertificateV4() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,
		MigrateState:  resourceACMECertificateMigrateState,
		SchemaVersion: 4,
		StateUpgraders: []schema.StateUpgrader{
			resourceACMECertificateStateUpgraderV3(),
		},
		Schema: map[string]*schema.Schema{
			"account_key_pem": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"common_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"certificate_request_pem"},
			},
			"subject_alternative_names": {
				Type:          schema.TypeSet,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				ForceNew:      true,
				ConflictsWith: []string{"certificate_request_pem"},
			},
			"key_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Default:       "2048",
				ConflictsWith: []string{"certificate_request_pem"},
				ValidateFunc:  validateKeyType,
			},
			"certificate_request_pem": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"common_name", "subject_alternative_names", "key_type"},
			},
			"min_days_remaining": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"dns_challenge": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Required: true,
						},
						"config": {
							Type:         schema.TypeMap,
							Optional:     true,
							ValidateFunc: validateDNSChallengeConfig,
							Sensitive:    true,
						},
					},
				},
			},
			"recursive_nameservers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"must_staple": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"certificate_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key_pem": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"certificate_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_p12": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"certificate_p12_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Default:   "",
				Sensitive: true,
			},
		},
	}
}

func resourceACMECertificateV3() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,
		MigrateState:  resourceACMECertificateMigrateState,
		SchemaVersion: 3,
		Schema: map[string]*schema.Schema{
			"account_key_pem": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"common_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"certificate_request_pem"},
			},
			"subject_alternative_names": {
				Type:          schema.TypeSet,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				ForceNew:      true,
				ConflictsWith: []string{"certificate_request_pem"},
			},
			"key_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Default:       "2048",
				ConflictsWith: []string{"certificate_request_pem"},
				ValidateFunc:  validateKeyType,
			},
			"certificate_request_pem": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"common_name", "subject_alternative_names", "key_type"},
			},
			"min_days_remaining": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"dns_challenge": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Required: true,
						},
						"config": {
							Type:         schema.TypeMap,
							Optional:     true,
							ValidateFunc: validateDNSChallengeConfig,
							Sensitive:    true,
						},
						"recursive_nameservers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"must_staple": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"certificate_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key_pem": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"certificate_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_p12": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"certificate_p12_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Default:   "",
				Sensitive: true,
			},
		},
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

func resourceACMECertificateRead(d *schema.ResourceData, meta interface{}) error {
	// This is a workaround to correct issues with some versions of the
	// resource prior to 1.3.2 where a renewal failure would possibly
	// delete the certificate.
	if _, ok := d.GetOk("certificate_pem"); !ok {
		// Try to recover the certificate from the ACME API.
		client, _, err := expandACMEClient(d, meta, true)
		if err != nil {
			return err
		}

		srcCR, err := client.Certificate.Get(d.Id(), true)
		if err != nil {
			// There are probably some cases that we will want to just drop
			// the resource if there's been an issue, but seeing as this is
			// mainly being used to recover for a bug that will be gone in
			// 1.3.2, this will probably be rare. If we start relying on
			// this behavior on a more general level, we may need to
			// investigate this more. Just error on everything for now.
			return err
		}

		dstCR := expandCertificateResource(d)
		dstCR.Certificate = srcCR.Certificate
		password := d.Get("certificate_p12_password").(string)
		if err := saveCertificateResource(d, dstCR, password); err != nil {
			return err
		}
	}

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

	expired, err := resourceACMECertificateHasExpired(d)
	if err != nil {
		return err
	}

	if expired {
		d.SetNewComputed("certificate_pem")
		d.SetNewComputed("certificate_p12")
		d.SetNewComputed("certificate_url")
		d.SetNewComputed("certificate_domain")
		d.SetNewComputed("private_key_pem")
		d.SetNewComputed("issuer_pem")
	}

	return nil
}

// resourceACMECertificateUpdate renews a certificate if it has been flagged as changed.
func resourceACMECertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	// We don't need to do anything else here if the certificate hasn't been diffed
	expired, err := resourceACMECertificateHasExpired(d)
	if err != nil {
		return err
	}

	if !expired {
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

	// Enable partial mode to protect the certificate during renewal
	d.Partial(true)

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

	// Complete, safe to turn off partial mode now.
	d.Partial(false)
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

// resourceACMECertificateHasExpired checks the acme_certificate
// resource to see if it has expired.
func resourceACMECertificateHasExpired(d certificateResourceExpander) (bool, error) {
	mindays := d.Get("min_days_remaining").(int)
	if mindays < 0 {
		log.Printf("[WARN] min_days_remaining is set to less than 0, certificate will never be renewed")
		return false, nil
	}

	cert := expandCertificateResource(d)
	remaining, err := certDaysRemaining(cert)
	if err != nil {
		return false, err
	}

	if int64(mindays) >= remaining {
		return true, nil
	}

	return false, nil
}

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
