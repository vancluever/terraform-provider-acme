package acme

import (
	"github.com/go-acme/lego/v4/acme"
	"github.com/go-acme/lego/v4/registration"
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceACMERegistration returns the current version of the
// acme_registration resource and needs to be updated when the schema
// version is incremented.
func resourceACMERegistration() *schema.Resource { return resourceACMERegistrationV1() }

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
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
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

func resourceACMERegistrationCreate(d *schema.ResourceData, meta interface{}) error {
	// register and agree to the TOS
	client, _, err := expandACMEClient(d, meta, false)
	if err != nil {
		return err
	}

	var reg *registration.Resource
	// If EAB was enabled, register using EAB.
	if v, ok := d.GetOk("external_account_binding"); ok {
		reg, err = client.Registration.RegisterWithExternalAccountBinding(registration.RegisterEABOptions{
			TermsOfServiceAgreed: true,
			Kid:                  v.([]interface{})[0].(map[string]interface{})["key_id"].(string),
			HmacEncoded:          v.([]interface{})[0].(map[string]interface{})["hmac_base64"].(string),
		})
	} else {
		// Normal registration.
		reg, err = client.Registration.Register(registration.RegisterOptions{
			TermsOfServiceAgreed: true,
		})
	}

	if err != nil {
		return err
	}

	d.SetId(reg.URI)

	return resourceACMERegistrationRead(d, meta)
}

func resourceACMERegistrationRead(d *schema.ResourceData, meta interface{}) error {
	_, user, err := expandACMEClient(d, meta, true)
	if err != nil {
		if regGone(err) {
			d.SetId("")
			return nil
		}

		return err
	}

	// save the reg
	return saveACMERegistration(d, user.Registration)
}

func resourceACMERegistrationDelete(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, meta, true)
	if err != nil {
		return err
	}

	return client.Registration.DeleteRegistration()
}

func regGone(err error) bool {
	e, ok := err.(*acme.ProblemDetails)
	if !ok {
		return false
	}

	switch {
	case e.HTTPStatus == 400 && e.Type == "urn:ietf:params:acme:error:accountDoesNotExist":
		// As per RFC8555, see: no account exists when onlyReturnExisting
		// is set to true.
		return true

	case e.HTTPStatus == 403 && e.Type == "urn:ietf:params:acme:error:unauthorized":
		// Usually happens when the account has been deactivated. The URN
		// is a bit general for my liking, but it should be fine given
		// the specific nature of the request this error would be
		// returned for.
		return true
	}

	return false
}
