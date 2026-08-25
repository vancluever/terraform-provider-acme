package acme

import (
	"context"
	"crypto"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/lego"
	"github.com/go-acme/lego/v5/registration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// acmeUser implements acme.User.
type acmeUser struct {
	// The email address for the account.
	Email string

	// The registration resource object.
	Registration *acme.ExtendedAccount

	// The private key for the account.
	key crypto.Signer
}

func (u acmeUser) GetEmail() string {
	return u.Email
}
func (u acmeUser) GetRegistration() *acme.ExtendedAccount {
	return u.Registration
}
func (u acmeUser) GetPrivateKey() crypto.Signer {
	return u.key
}

// expandACMEUser creates a new instance of an ACME user from set
// email_address and private_key_pem fields, and a registration
// if one exists.
func expandACMEUser(d *schema.ResourceData) (*acmeUser, error) {
	key, err := privateKeyFromPEM([]byte(d.Get("account_key_pem").(string)))
	if err != nil {
		return nil, err
	}

	user := &acmeUser{
		key: key,
	}

	// only set these email if it's in the schema.
	if v, ok := d.GetOk("email_address"); ok {
		user.Email = v.(string)
	}

	return user, nil
}

// expandACMEClient creates a connection to an ACME server from resource data,
// and also returns the user.
//
// If loadReg is supplied, the registration information is loaded in to the
// user's registration, if it exists - if the account cannot be resolved by the
// private key, then the appropriate error is returned.
func expandACMEClient(d *schema.ResourceData, meta any, loadReg bool) (*lego.Client, *acmeUser, error) {
	user, err := expandACMEUser(d)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting user data: %s", err.Error())
	}

	client, err := lego.NewClient(expandACMEClient_config(d, meta, user))
	if err != nil {
		return nil, nil, err
	}

	// Populate user's registration resource if needed
	if loadReg {
		user.Registration, err = client.Registration.ResolveAccountByKey(context.TODO())
		if err != nil {
			return nil, nil, err
		}
	}

	return client, user, nil
}

func expandACMEClient_config(d *schema.ResourceData, meta any, user registration.User) *lego.Config {
	config := lego.NewConfig(user)
	config.CADirURL = meta.(*Config).ServerURL
	config.UserAgent = formatUserAgentShort()

	// Note this function is used by both the registration and certificate
	// resources, but cert timeout is not necessary during registration, so it's
	// okay if it's empty for that.
	if v, ok := d.GetOk("cert_timeout"); ok {
		config.Certificate.Timeout = time.Second * time.Duration(v.(int))
	}

	return config
}

func formatUserAgentShort() string {
	return "terraform-provider-acme" + "/" + ReleaseVersion
}

// FormatUserAgentLong returns a agent for use in client requests, we also use
// it in the -version command for the plugin.
func FormatUserAgentLong() string {
	return strings.TrimSpace(fmt.Sprintf("%s (%s; %s)", formatUserAgentShort(), runtime.GOOS, runtime.GOARCH))
}
