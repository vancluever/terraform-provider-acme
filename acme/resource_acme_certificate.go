package acme

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-acme/lego/v4/acme"
	"github.com/go-acme/lego/v4/acme/api"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type RevocationReason string

const (
	RevocationReasonUnspecified          RevocationReason = "unspecified"
	RevocationReasonKeyCompromise        RevocationReason = "key-compromise"
	RevocationReasonCACompromise         RevocationReason = "ca-compromise"
	RevocationReasonAffiliationChanged   RevocationReason = "affiliation-changed"
	RevocationReasonSuperseded           RevocationReason = "superseded"
	RevocationReasonCessationOfOperation RevocationReason = "cessation-of-operation"
	RevocationReasonCertificateHold      RevocationReason = "certificate-hold"
	RevocationReasonRemoveFromCRL        RevocationReason = "remove-from-crl"
	RevocationReasonPrivilegeWithdrawn   RevocationReason = "privilege-withdrawn"
	RevocationReasonAACompromise         RevocationReason = "aa-compromise"
)

// resourceACMECertificate returns the current version of the
// acme_registration resource and needs to be updated when the schema
// version is incremented.
func resourceACMECertificate() *schema.Resource { return resourceACMECertificateV5() }

func resourceACMECertificateV5() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,
		MigrateState:  resourceACMECertificateMigrateState,
		SchemaVersion: 5,
		StateUpgraders: []schema.StateUpgrader{
			resourceACMECertificateStateUpgraderV4(),
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
				AtLeastOneOf:  []string{"common_name", "subject_alternative_names", "certificate_request_pem"},
				ConflictsWith: []string{"certificate_request_pem"},
			},
			"subject_alternative_names": {
				Type:          schema.TypeSet,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				ForceNew:      true,
				AtLeastOneOf:  []string{"common_name", "subject_alternative_names", "certificate_request_pem"},
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
				AtLeastOneOf:  []string{"common_name", "subject_alternative_names", "certificate_request_pem"},
				ConflictsWith: []string{"common_name", "subject_alternative_names", "key_type"},
			},
			"min_days_remaining": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"use_renewal_info": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"renewal_info_max_sleep": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 900),
			},
			"renewal_info_ignore_retry_after": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dns_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
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
			"http_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
				ConflictsWith: []string{
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
				},
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      80,
							ValidateFunc: validation.IsPortNumber,
						},
						"proxy_header": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"http_webroot_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
				ConflictsWith: []string{
					"http_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
				},
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"directory": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"http_memcached_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
				ConflictsWith: []string{"http_challenge", "http_webroot_challenge", "http_s3_challenge"},
				MaxItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hosts": {
							Type:     schema.TypeSet,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							MinItems: 1,
						},
					},
				},
			},
			"http_s3_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
				ConflictsWith: []string{"http_challenge", "http_webroot_challenge", "http_memcached_challenge"},
				MaxItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"s3_bucket": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"tls_challenge": {
				Type:     schema.TypeList,
				Optional: true,
				AtLeastOneOf: []string{
					"dns_challenge",
					"http_challenge",
					"http_webroot_challenge",
					"http_memcached_challenge",
					"http_s3_challenge",
					"tls_challenge",
				},
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      443,
							ValidateFunc: validation.IsPortNumber,
						},
					},
				},
			},
			"pre_check_delay": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"recursive_nameservers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disable_complete_propagation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"must_staple": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"preferred_chain": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"profile": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"cert_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"deactivate_authorizations": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			"certificate_not_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_serial": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_p12_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Default:   "",
				Sensitive: true,
			},
			"revoke_certificate_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"revoke_certificate_reason": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateRevocationReason,
			},
			"renewal_info_window_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_info_window_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_info_window_selected": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_info_explanation_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_info_retry_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceACMECertificateCreate(d *schema.ResourceData, meta any) error {
	// Pre-generate resource UUID here, in case there is a serious
	// issue with UUID generation that would lead to inconsistency.
	//
	// We do not use the ID of the certificate here as the IDs of
	// certificates drift during renewal (they are effectively new
	// certificates). Use the certificate_url to get the URL of the
	// current certificate instead.
	resourceUUID, err := uuid.GenerateUUID()
	if err != nil {
		return fmt.Errorf("error generating UUID for resource: %s", err)
	}

	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	dnsCloser, err := setCertificateChallengeProviders(client, d)
	defer dnsCloser()
	if err != nil {
		return err
	}

	var cert *certificate.Resource

	if v, ok := d.GetOk("certificate_request_pem"); ok {
		var csr *x509.CertificateRequest
		csr, err = csrFromPEM([]byte(v.(string)))
		if err != nil {
			return err
		}
		cert, err = client.Certificate.ObtainForCSR(certificate.ObtainForCSRRequest{
			CSR:                            csr,
			Bundle:                         true,
			PreferredChain:                 d.Get("preferred_chain").(string),
			Profile:                        d.Get("profile").(string),
			AlwaysDeactivateAuthorizations: d.Get("deactivate_authorizations").(bool),
		})
	} else {
		domains := []string{}
		cn := d.Get("common_name").(string)
		if cn != "" {
			domains = append(domains, cn)
		}

		if s, ok := d.GetOk("subject_alternative_names"); ok {
			for _, v := range stringSlice(s.(*schema.Set).List()) {
				if v != cn {
					domains = append(domains, v)
				}
			}
		}

		cert, err = client.Certificate.Obtain(certificate.ObtainRequest{
			Domains:                        domains,
			Bundle:                         true,
			MustStaple:                     d.Get("must_staple").(bool),
			PreferredChain:                 d.Get("preferred_chain").(string),
			Profile:                        d.Get("profile").(string),
			AlwaysDeactivateAuthorizations: d.Get("deactivate_authorizations").(bool),
		})
	}

	if err != nil {
		return fmt.Errorf("error creating certificate: %s", err)
	}

	d.SetId(resourceUUID)
	password := d.Get("certificate_p12_password").(string)
	if err := saveCertificateResource(d, cert, password); err != nil {
		return err
	}

	return resourceACMECertificateRead(d, meta)
}

func resourceACMECertificateRead(d *schema.ResourceData, meta any) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("certificate_pem"); !ok {
		// This is a workaround to correct issues with some versions of the
		// resource prior to 1.3.2 where a renewal failure would possibly delete
		// the certificate.
		//
		// Try to recover the certificate from the ACME API.
		srcCR, err := client.Certificate.Get(d.Get("certificate_url").(string), true)
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

	if err := resourceACMECertificateRenewalInfoRefresh(d, client, time.Now()); err != nil {
		return err
	}

	return nil
}

// resourceACMECertificateCustomizeDiff checks the certificate for renewal and
// flags it as NewComputed if it needs a renewal.
func resourceACMECertificateCustomizeDiff(_ context.Context, d *schema.ResourceDiff, meta any) error {
	// Ensure duplicate providers for dns_challenge are not provided.
	providerMap := make(map[string]bool)
	for _, v := range d.Get("dns_challenge").([]any) {
		m := v.(map[string]any)
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

	shouldRenew, err := resourceACMECertificateShouldRenew(d, time.Now())
	if err != nil {
		return err
	}

	if shouldRenew {
		d.SetNewComputed("certificate_pem")
		d.SetNewComputed("certificate_p12")
		d.SetNewComputed("certificate_url")
		d.SetNewComputed("certificate_domain")
		d.SetNewComputed("certificate_not_after")
		d.SetNewComputed("private_key_pem")
		d.SetNewComputed("issuer_pem")
		d.SetNewComputed("certificate_serial")
		d.SetNewComputed("renewal_info_window_start")
		d.SetNewComputed("renewal_info_window_end")
		d.SetNewComputed("renewal_info_window_selected")
		d.SetNewComputed("renewal_info_explanation_url")
		d.SetNewComputed("renewal_info_retry_after")
	}

	return nil
}

// resourceACMECertificateUpdate renews a certificate if it has been flagged as changed.
func resourceACMECertificateUpdate(d *schema.ResourceData, meta any) error {
	shouldRenew, err := resourceACMECertificateShouldRenew(d, time.Now())
	if err != nil {
		return err
	}

	if !shouldRenew {
		// when the certificate hasn't changed but the p12 password has, we still need to regenerate the p12
		if d.HasChange("certificate_p12_password") {
			cert := expandCertificateResource(d)
			password := d.Get("certificate_p12_password").(string)
			if err := saveCertificateResource(d, cert, password); err != nil {
				return err
			}
		}
	} else {
		// Enable partial mode to protect the certificate during renewal
		d.Partial(true)

		// Sleep until renewal time if necessary (in the case of ARI)
		if err := resourceACMECertificateSleepUntilRenewalTime(d); err != nil {
			return err
		}

		client, _, err := expandACMEClient(d, meta, true)
		if err != nil {
			return err
		}

		cert := expandCertificateResource(d)

		dnsCloser, err := setCertificateChallengeProviders(client, d)
		defer dnsCloser()
		if err != nil {
			return err
		}

		newCert, err := renewWithOptions(
			client.Certificate,
			*cert,
			localRenewOptions{
				RenewOptions: certificate.RenewOptions{
					Bundle:                         true,
					PreferredChain:                 d.Get("preferred_chain").(string),
					Profile:                        d.Get("profile").(string),
					MustStaple:                     d.Get("must_staple").(bool),
					AlwaysDeactivateAuthorizations: d.Get("deactivate_authorizations").(bool),
				},
				UseARI: d.Get("use_renewal_info").(bool),
			},
		)
		if err != nil {
			return err
		}

		password := d.Get("certificate_p12_password").(string)
		if err := saveCertificateResource(d, newCert, password); err != nil {
			return err
		}

		// Complete, safe to turn off partial mode now.
		d.Partial(false)

		// Clear out ARI computed data so that it can be properly refreshed on the
		// below read.
		d.Set("renewal_info_window_start", "")
		d.Set("renewal_info_window_end", "")
		d.Set("renewal_info_window_selected", "")
		d.Set("renewal_info_explanation_url", "")
		d.Set("renewal_info_retry_after", "")
	}

	return resourceACMECertificateRead(d, meta)
}

// resourceACMECertificateDelete "deletes" the certificate by revoking it.
func resourceACMECertificateDelete(d *schema.ResourceData, meta any) error {
	if !d.Get("revoke_certificate_on_destroy").(bool) {
		return nil
	}

	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	cert := expandCertificateResource(d)
	remaining, err := certSecondsRemaining(cert, time.Now())
	if err != nil {
		return err
	}

	if remaining >= 0 {
		maybeReason, ok := d.GetOk("revoke_certificate_reason")
		if ok {
			reason := RevocationReason(maybeReason.(string))
			reasonNum, err := GetRevocationReason(reason)
			if err != nil {
				return err
			}
			return client.Certificate.RevokeWithReason(cert.Certificate, &reasonNum)
		}
		return client.Certificate.Revoke(cert.Certificate)
	}
	return nil
}

// resourceACMECertificateHasExpired checks the acme_certificate
// resource to see if it has expired.
func resourceACMECertificateHasExpired(d resourceDataOrDiff, now time.Time) (bool, error) {
	mindays := d.Get("min_days_remaining").(int)
	if mindays < 0 {
		log.Printf("[WARN] min_days_remaining is set to less than 0, certificate will never be renewed")
		return false, nil
	}

	cert := expandCertificateResource(d)
	remaining, err := certDaysRemaining(cert, now)
	if err != nil {
		return false, err
	}

	if int64(mindays) >= remaining {
		return true, nil
	}

	return false, nil
}

func resourceACMECertificatePreCheckDelay(delay int) dns01.WrapPreCheckFunc {
	// Compute a reasonable interval for the delay, max delay 10
	// seconds, minimum 2.
	var interval int
	switch {
	case delay <= 10:
		interval = 2

	case delay <= 60:
		interval = 5

	default:
		interval = 10
	}

	return func(domain, fqdn, value string, orig dns01.PreCheckFunc) (bool, error) {
		stop, err := orig(fqdn, value)
		if stop && err == nil {
			// Run the delay. TODO: Eventually make this interruptible.
			var elapsed int
			end := time.After(time.Second * time.Duration(delay))
			for {
				select {
				case <-end:
					return true, nil
				default:
				}

				remaining := delay - elapsed
				if remaining < interval {
					// To honor the specified timeout, make our next interval the
					// time remaining. Minimum one second.
					interval = max(remaining, 1)
				}

				log.Printf("[DEBUG] [%s] acme: Waiting an additional %d second(s) for DNS record propagation.", domain, remaining)
				time.Sleep(time.Second * time.Duration(interval))
				elapsed += interval
			}
		}

		// A previous pre-check failed, return and exit.
		return stop, err
	}
}

func GetRevocationReason(reason RevocationReason) (uint, error) {
	switch reason {
	case RevocationReasonUnspecified:
		return acme.CRLReasonUnspecified, nil
	case RevocationReasonKeyCompromise:
		return acme.CRLReasonKeyCompromise, nil
	case RevocationReasonCACompromise:
		return acme.CRLReasonCACompromise, nil
	case RevocationReasonAffiliationChanged:
		return acme.CRLReasonAffiliationChanged, nil
	case RevocationReasonSuperseded:
		return acme.CRLReasonSuperseded, nil
	case RevocationReasonCessationOfOperation:
		return acme.CRLReasonCessationOfOperation, nil
	case RevocationReasonCertificateHold:
		return acme.CRLReasonCertificateHold, nil
	case RevocationReasonRemoveFromCRL:
		return acme.CRLReasonRemoveFromCRL, nil
	case RevocationReasonPrivilegeWithdrawn:
		return acme.CRLReasonPrivilegeWithdrawn, nil
	case RevocationReasonAACompromise:
		return acme.CRLReasonAACompromise, nil
	default:
		return acme.CRLReasonUnspecified, fmt.Errorf("unknown revocation reason: %s", reason)
	}
}

func resourceACMECertificateRenewalInfoRefresh(
	d *schema.ResourceData,
	client *lego.Client,
	now time.Time,
) error {
	// Check to see if we have a retry-after response, if we do, honor it
	// (i.e., skip if we are before it).
	retryAfterString, ok := d.GetOk("renewal_info_retry_after")
	if ok && !d.Get("renewal_info_ignore_retry_after").(bool) {
		retryAfter, err := time.Parse(time.RFC3339, retryAfterString.(string))
		if err != nil {
			return fmt.Errorf("malformed renewal_info_retry_after: %w", err)
		} else if now.Before(retryAfter) {
			return nil
		}
	}

	// Need to grab the cert from the PEM bundle
	cb, err := parsePEMBundle([]byte(d.Get("certificate_pem").(string)))
	if err != nil {
		return err
	}
	// lego always returns the issued cert first, if the CA is first there is a problem
	if cb[0].IsCA {
		return errors.New("cannot parse PEM bundle correctly: first certificate is a CA certificate")
	}

	cert := cb[0]
	if now.After(cert.NotAfter) {
		// Early exit here, as renewalInfo does not work for expired certificates.
		// Our diff logic will start the renewal process during CustomizeDiff.
		log.Println("[WARN] certificate is expired, cannot retrieve ARI data")
		d.Set("renewal_info_window_start", "")
		d.Set("renewal_info_window_end", "")
		d.Set("renewal_info_window_selected", "")
		d.Set("renewal_info_explanation_url", "")
		d.Set("renewal_info_retry_after", "")
		return nil
	}

	renewalInfoResp, err := client.Certificate.GetRenewalInfo(certificate.RenewalInfoRequest{
		Cert: cert,
	})
	if err != nil {
		if errors.Is(err, api.ErrNoARI) {
			// No ARI detail, set blank values and return
			log.Println("[WARN] cannot retrieve ARI data as it is unsupported on the endpoint")
			d.Set("renewal_info_window_start", "")
			d.Set("renewal_info_window_end", "")
			d.Set("renewal_info_window_selected", "")
			d.Set("renewal_info_explanation_url", "")
			d.Set("renewal_info_retry_after", "")
			return nil
		}

		return err
	}

	// Select a random time from within the window to renew.
	//
	// Note that this differs from lego's logic - we don't use the provided
	// helper from the response as it will not return a renewal time at all if
	// it's too far in the future, past when the client is willing to sleep. This
	// can create some inconsistent results for us in Terraform - first, each
	// refresh will non-deterministically select different renew times, some that
	// may not fall in the max sleep, but some that may. Additionally, this full
	// reliance on state can prevent us from being able to check effectively in
	// the diff whether or not we can renew, if the presence of the selected time
	// is the sole the thing that determines it. Having a stable, one-time
	// selected timestamp saved to state allows us to run the diff with the
	// configured sleep value and get a consistent result every time. This also
	// allows us to generate diffs immediately on the setting of either
	// use_renewal_info or renewal_info_max_sleep, which would not have been
	// possible otherwise.
	windowStart := renewalInfoResp.SuggestedWindow.Start
	windowEnd := renewalInfoResp.SuggestedWindow.End
	windowSelected := windowStart
	if windowDuration := windowEnd.Sub(windowStart); windowDuration > 0 {
		randomDuration := time.Duration(rand.Int63n(int64(windowDuration)))
		windowSelected = windowSelected.Add(randomDuration)
	}

	d.Set("renewal_info_window_start", windowStart.UTC().Format(time.RFC3339))
	d.Set("renewal_info_window_end", windowEnd.UTC().Format(time.RFC3339))
	d.Set("renewal_info_explanation_url", renewalInfoResp.ExplanationURL)
	d.Set("renewal_info_retry_after", now.Add(renewalInfoResp.RetryAfter).UTC().Format(time.RFC3339))
	d.Set("renewal_info_window_selected", windowSelected.UTC().Format(time.RFC3339))

	return nil
}

func resourceACMECertificateShouldRenew(d resourceDataOrDiff, now time.Time) (bool, error) {
	if d.Get("use_renewal_info").(bool) {
		canSleep, err := resourceACMECertificateRenewalInfoCanSleepUntilSelected(d, now)
		if err != nil {
			return false, err
		}

		if canSleep {
			return true, nil
		}
	}

	return resourceACMECertificateHasExpired(d, now)
}

func resourceACMECertificateRenewalInfoCanSleepUntilSelected(
	d resourceDataOrDiff,
	now time.Time,
) (bool, error) {
	var selectedTime time.Time
	if selected := d.Get("renewal_info_window_selected").(string); selected != "" {
		var err error
		selectedTime, err = time.Parse(time.RFC3339, selected)
		if err != nil {
			return false, fmt.Errorf("malformed renewal_info_window_selected: %w", err)
		}
	} else {
		return false, errors.New(
			"renewal_info_window_selected expected to be set. This is a bug, please report it")
	}

	canSleepUntil := now.Add(time.Second * time.Duration(d.Get("renewal_info_max_sleep").(int)))
	if canSleepUntil.Before(selectedTime) {
		return false, nil
	}

	return true, nil
}

func resourceACMECertificateSleepUntilRenewalTime(d *schema.ResourceData) error {
	if !d.Get("use_renewal_info").(bool) {
		return nil
	}

	var selectedTime time.Time
	if selected := d.Get("renewal_info_window_selected").(string); selected != "" {
		var err error
		selectedTime, err = time.Parse(time.RFC3339, selected)
		if err != nil {
			return fmt.Errorf("malformed renewal_info_window_selected: %w", err)
		}
	} else {
		return errors.New("renewal_info_window_selected expected to be set. This is a bug, please report it")
	}

	sleepDuration := time.Until(selectedTime)
	log.Printf(
		"[DEBUG] sleeping %s until renewal time: %s",
		sleepDuration.Truncate(time.Second),
		selectedTime,
	)

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(sleepDuration)
		done <- true
	}()
	for {
		select {
		case <-done:
			log.Println("[DEBUG] sleep complete, proceeding with renewal")
			return nil
		case <-ticker.C:
			sleepDurationRemaining := time.Until(selectedTime)
			log.Printf(
				"[DEBUG] (%s remaining) sleeping until renewal time: %s",
				sleepDurationRemaining.Truncate(time.Second),
				selectedTime,
			)
		}
	}
}
