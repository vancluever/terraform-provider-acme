package acme

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceACMECertificateV6() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMECertificateCreate,
		Read:          resourceACMECertificateRead,
		CustomizeDiff: resourceACMECertificateCustomizeDiff,
		Update:        resourceACMECertificateUpdate,
		Delete:        resourceACMECertificateDelete,
		MigrateState:  resourceACMECertificateMigrateState,
		SchemaVersion: 6,
		StateUpgraders: []schema.StateUpgrader{
			resourceACMECertificateStateUpgraderV5(),
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
				Default:       "RSA2048",
				ConflictsWith: []string{"certificate_request_pem"},
				ValidateFunc:  validateKeyTypeV3Old,
			},
			"certificate_request_pem": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				AtLeastOneOf:  []string{"common_name", "subject_alternative_names", "certificate_request_pem"},
				ConflictsWith: []string{"common_name", "subject_alternative_names", "key_type"},
			},
			"validity_days": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},
			"min_days_remaining": {
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       30,
				ConflictsWith: []string{"min_days_dynamic"},
			},
			"min_days_dynamic": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"min_days_remaining"},
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
						"match_domains": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       0,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"propagation_wait"},
			},
			"recursive_nameservers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disable_authoritative_propagation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"propagation_wait": {
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       0,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"pre_check_delay"},
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
			"certificate_not_before": {
				Type:     schema.TypeString,
				Computed: true,
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
			"validity_days": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},
			"min_days_remaining": {
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       30,
				ConflictsWith: []string{"min_days_dynamic"},
			},
			"min_days_dynamic": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"min_days_remaining"},
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
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       0,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"propagation_wait"},
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
			"propagation_wait": {
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       0,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"pre_check_delay"},
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
			"certificate_not_before": {
				Type:     schema.TypeString,
				Computed: true,
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
