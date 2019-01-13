package acme

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xenolf/lego/registration"
)

func resourceACMERegistration() *schema.Resource {
	return &schema.Resource{
		Create: resourceACMERegistrationCreate,
		Read:   resourceACMERegistrationRead,
		Delete: resourceACMERegistrationDelete,

		Schema:        registrationSchemaFull(),
		SchemaVersion: 1,
		MigrateState:  resourceACMERegistrationMigrateState,
	}
}

func resourceACMERegistrationCreate(d *schema.ResourceData, meta interface{}) error {
	// register and agree to the TOS
	client, _, err := expandACMEClient(d, meta, false)
	if err != nil {
		return err
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{
		TermsOfServiceAgreed: true,
	})
	if err != nil {
		return err
	}
	d.SetId(reg.URI)

	return resourceACMERegistrationRead(d, meta)
}

func resourceACMERegistrationRead(d *schema.ResourceData, meta interface{}) error {
	_, user, err := expandACMEClient(d, meta, true)
	if err != nil {
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
