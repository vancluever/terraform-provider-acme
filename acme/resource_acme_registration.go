package acme

import "github.com/hashicorp/terraform/helper/schema"

func resourceACMERegistration() *schema.Resource {
	return &schema.Resource{
		Create: resourceACMERegistrationCreate,
		Read:   resourceACMERegistrationRead,
		Delete: resourceACMERegistrationDelete,

		Schema: registrationSchemaFull(),
	}
}

func resourceACMERegistrationCreate(d *schema.ResourceData, meta interface{}) error {
	// register and agree to the TOS
	client, _, err := expandACMEClient(d, false)
	if err != nil {
		return err
	}
	_, err = client.Register(true)
	if err != nil {
		return err
	}

	return resourceACMERegistrationRead(d, meta)
}

func resourceACMERegistrationRead(d *schema.ResourceData, meta interface{}) error {
	// NOTE: This may change the ID of the resource - this is currently an
	// unfortunate consequence of the transition from ACME v1 to ACME v2.
	_, user, err := expandACMEClient(d, true)
	if err != nil {
		return err
	}

	// save the reg
	return saveACMERegistration(d, user.Registration)
}

func resourceACMERegistrationDelete(d *schema.ResourceData, meta interface{}) error {
	client, _, err := expandACMEClient(d, true)
	if err != nil {
		return err
	}

	return client.DeleteRegistration()
}
