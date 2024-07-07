package acme

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceACMERegistrationV1() *schema.Resource {
	return &schema.Resource{
		Create:        resourceACMERegistrationCreate,
		Read:          resourceACMERegistrationRead,
		Delete:        resourceACMERegistrationDelete,
		MigrateState:  resourceACMERegistrationMigrateState,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"account_key_pem": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
				ConflictsWith: []string{
					"account_key_algorithm",
					"account_key_ecdsa_curve",
					"account_key_rsa_bits",
				},
			},
			// https://letsencrypt.org/docs/integration-guide/#supported-key-algorithms
			// NOTE: Our internal functions support more, but we need to restrict to
			// what's listed here for Let's Encrypt Specifically. This also applies
			// to the specific RSA and ECDSA lengths/curves.
			"account_key_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice(
					[]string{keyAlgorithmRSA, keyAlgorithmECDSA},
					false,
				),
				Default:       keyAlgorithmECDSA,
				ConflictsWith: []string{"account_key_pem"},
			},
			"account_key_ecdsa_curve": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice(
					[]string{keyECDSACurveP256, keyECDSACurveP384},
					false,
				),
				Default:       keyECDSACurveP384,
				ConflictsWith: []string{"account_key_pem", "account_key_rsa_bits"},
			},
			"account_key_rsa_bits": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ValidateFunc:  validation.IntInSlice([]int{2048, 3072, 4096}),
				Default:       4096,
				ConflictsWith: []string{"account_key_pem", "account_key_ecdsa_curve"},
			},
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_account_binding": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_id": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"hmac_base64": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
							ForceNew:  true,
						},
					},
				},
			},
			"registration_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
