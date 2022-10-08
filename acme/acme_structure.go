package acme

//go:generate go run ../build-support/generate-dns-providers go dns_provider_factory.go

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"software.sslmate.com/src/go-pkcs12"
)

// acmeUser implements acme.User.
type acmeUser struct {

	// The email address for the account.
	Email string

	// The registration resource object.
	Registration *registration.Resource

	// The private key for the account.
	key crypto.PrivateKey
}

func (u acmeUser) GetEmail() string {
	return u.Email
}
func (u acmeUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u acmeUser) GetPrivateKey() crypto.PrivateKey {
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

// saveACMERegistration takes an *registration.Resource and sets the appropriate fields
// for a registration resource.
func saveACMERegistration(d *schema.ResourceData, reg *registration.Resource) error {
	d.Set("registration_url", reg.URI)

	return nil
}

// expandACMEClient creates a connection to an ACME server from resource data,
// and also returns the user.
//
// If loadReg is supplied, the registration information is loaded in to the
// user's registration, if it exists - if the account cannot be resolved by the
// private key, then the appropriate error is returned.
func expandACMEClient(d *schema.ResourceData, meta interface{}, loadReg bool) (*lego.Client, *acmeUser, error) {
	user, err := expandACMEUser(d)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting user data: %s", err.Error())
	}

	config := lego.NewConfig(user)
	config.CADirURL = meta.(*Config).ServerURL

	// Note this function is used by both the registration and certificate
	// resources, but key type is not necessary during registration, so
	// it's okay if it's empty for that.
	if v, ok := d.GetOk("key_type"); ok {
		config.Certificate.KeyType = certcrypto.KeyType(v.(string))
	}

	client, err := lego.NewClient(config)
	if err != nil {
		return nil, nil, err
	}

	// Populate user's registration resource if needed
	if loadReg {
		user.Registration, err = client.Registration.ResolveAccountByKey()
		if err != nil {
			return nil, nil, err
		}
	}

	return client, user, nil
}

// certificateResourceExpander is a simple interface to allow us to use the Get
// function that is in ResourceData and ResourceDiff under the same function.
type certificateResourceExpander interface {
	Get(string) interface{}
	GetOk(string) (interface{}, bool)
	GetChange(string) (interface{}, interface{})
}

// expandCertificateResource takes saved state in the certificate resource
// and returns an certificate.Resource.
func expandCertificateResource(d certificateResourceExpander) *certificate.Resource {
	cert := &certificate.Resource{
		Domain:  d.Get("certificate_domain").(string),
		CertURL: d.Get("certificate_url").(string),
	}

	// Only populate the PrivateKey or CSR fields if we have them
	if pk, ok := d.GetOk("private_key_pem"); ok {
		cert.PrivateKey = []byte(pk.(string))
	}
	if csr, ok := d.GetOk("certificate_request_pem"); ok {
		cert.CSR = []byte(csr.(string))
	}

	// There are situations now where the new certificate may be blank, which
	// signifies that the certificate needs to be renewed. In this case, we need
	// the old value here, versus the new one.
	oldCertPEM, newCertPEM := d.GetChange("certificate_pem")
	issuerPEM := d.Get("issuer_pem")
	if newCertPEM.(string) != "" {
		cert.Certificate = []byte(newCertPEM.(string) + issuerPEM.(string))
	} else {
		cert.Certificate = []byte(oldCertPEM.(string) + issuerPEM.(string))
	}
	return cert
}

// saveCertificateResource takes an certificate.Resource and sets fields.
func saveCertificateResource(d *schema.ResourceData, cert *certificate.Resource, password string) error {
	d.Set("certificate_url", cert.CertURL)
	d.Set("certificate_domain", cert.Domain)
	d.Set("private_key_pem", string(cert.PrivateKey))
	issued, issuedNotAfter, issuer, err := splitPEMBundle(cert.Certificate)
	if err != nil {
		return err
	}

	d.Set("certificate_pem", string(issued))
	d.Set("issuer_pem", string(issuer))
	d.Set("certificate_not_after", issuedNotAfter)

	// Set PKCS12 data. This is only set if there is a private key
	// present.
	if len(cert.PrivateKey) > 0 {
		pfxB64, err := bundleToPKCS12(cert.Certificate, cert.PrivateKey, password)
		if err != nil {
			return err
		}

		d.Set("certificate_p12", string(pfxB64))
	} else {
		d.Set("certificate_p12", "")
	}

	return nil
}

// certSecondsRemaining takes an certificate.Resource, parses the
// certificate, and computes the seconds that it has remaining.
func certSecondsRemaining(cert *certificate.Resource) (int64, error) {
	x509Certs, err := parsePEMBundle(cert.Certificate)
	if err != nil {
		return 0, err
	}
	c := x509Certs[0]

	if c.IsCA {
		return 0, fmt.Errorf("first certificate is a CA certificate")
	}

	expiry := c.NotAfter.Unix()
	now := time.Now().Unix()

	return (expiry - now), nil
}

// certDaysRemaining takes an certificate.Resource, parses the
// certificate, and computes the days that it has remaining.
func certDaysRemaining(cert *certificate.Resource) (int64, error) {
	remaining, err := certSecondsRemaining(cert)
	if err != nil {
		return 0, fmt.Errorf("unable to calculate time to certificate expiry: %s", err)
	}

	return remaining / 86400, nil
}

// splitPEMBundle gets a slice of x509 certificates from
// parsePEMBundle.
//
// The first certificate split is returned as the issued certificate,
// with the rest returned as the issuer (intermediate) chain.
//
// Technically, it will be possible for issuer to be empty, if there
// are zero certificates in the intermediate chain. This is highly
// unlikely, however.
func splitPEMBundle(bundle []byte) (cert []byte, certNotAfter string, issuer []byte, err error) {
	cb, err := parsePEMBundle(bundle)
	if err != nil {
		return
	}

	// lego always returns the issued cert first, if the CA is first there is a problem
	if cb[0].IsCA {
		err = fmt.Errorf("first certificate is a CA certificate")
		return
	}

	cert = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cb[0].Raw})
	certNotAfter = cb[0].NotAfter.Format(time.RFC3339)
	issuer = make([]byte, 0)
	for _, ic := range cb[1:] {
		issuer = append(issuer, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ic.Raw})...)
	}

	return
}

// bundleToPKCS12 packs an issued certificate (and any supplied
// intermediates) into a PFX file.  The private key is included in
// the archive if it is a non-zero value.
//
// The returned archive is base64-encoded.
func bundleToPKCS12(bundle, key []byte, password string) ([]byte, error) {
	cb, err := parsePEMBundle(bundle)
	if err != nil {
		return nil, err
	}

	// lego always returns the issued cert first, if the CA is first there is a problem
	if cb[0].IsCA {
		return nil, fmt.Errorf("first certificate is a CA certificate")
	}

	pk, err := privateKeyFromPEM(key)
	if err != nil {
		return nil, err
	}

	pfxData, err := pkcs12.Encode(rand.Reader, pk, cb[0], cb[1:], password)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(pfxData)))
	base64.StdEncoding.Encode(buf, pfxData)
	return buf, nil
}

// parsePEMBundle parses a certificate bundle from top to bottom and returns
// a slice of x509 certificates. This function will error if no certificates are found.
//
// TODO: This was taken from lego directly, consider exporting it there, or
// consolidating with other TF crypto functions.
func parsePEMBundle(bundle []byte) ([]*x509.Certificate, error) {
	var certificates []*x509.Certificate
	var certDERBlock *pem.Block

	for {
		certDERBlock, bundle = pem.Decode(bundle)
		if certDERBlock == nil {
			break
		}

		if certDERBlock.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(certDERBlock.Bytes)
			if err != nil {
				return nil, err
			}
			certificates = append(certificates, cert)
		}
	}

	if len(certificates) == 0 {
		return nil, errors.New("no certificates were found while parsing the bundle")
	}

	return certificates, nil
}

// helper function to map environment variables if set
func mapEnvironmentVariableValues(keyMapping map[string]string) {
	for key := range keyMapping {
		if value, ok := os.LookupEnv(key); ok {
			os.Setenv(keyMapping[key], value)
		}
	}
}

// setDNSChallenge takes a *lego.Client and the DNS challenge complex
// structure as a map[string]interface{}, and configues the client to
// only allow a DNS challenge with the configured provider.
func setDNSChallenge(client *lego.Client, m map[string]interface{}) (challenge.Provider, error) {
	var providerName string

	if v, ok := m["provider"]; ok && v.(string) != "" {
		providerName = v.(string)
	} else {
		return nil, fmt.Errorf("DNS challenge provider not defined")
	}
	// Config only needs to be set if it's defined, otherwise existing env/SDK
	// defaults are fine.
	if v, ok := m["config"]; ok {
		for k, v := range v.(map[string]interface{}) {
			os.Setenv(k, v.(string))
		}
	}

	providerFunc, ok := dnsProviderFactory[providerName]
	if !ok {
		return nil, fmt.Errorf("%s: unsupported DNS challenge provider", providerName)
	}

	return providerFunc()
}

// stringSlice converts an interface slice to a string slice.
func stringSlice(src []interface{}) []string {
	var dst []string
	for _, v := range src {
		dst = append(dst, v.(string))
	}
	return dst
}

// privateKeyFromPEM converts a PEM block into a crypto.PrivateKey.
func privateKeyFromPEM(pemData []byte) (crypto.PrivateKey, error) {
	var result *pem.Block
	rest := pemData
	for {
		result, rest = pem.Decode(rest)
		if result == nil {
			return nil, fmt.Errorf("cannot decode supplied PEM data")
		}
		switch result.Type {
		case "RSA PRIVATE KEY":
			return x509.ParsePKCS1PrivateKey(result.Bytes)
		case "EC PRIVATE KEY":
			return x509.ParseECPrivateKey(result.Bytes)
		}
	}
}

// csrFromPEM converts a PEM block into an *x509.CertificateRequest.
func csrFromPEM(pemData []byte) (*x509.CertificateRequest, error) {
	var result *pem.Block
	rest := pemData
	for {
		result, rest = pem.Decode(rest)
		if result == nil {
			return nil, fmt.Errorf("cannot decode supplied PEM data")
		}
		if result.Type == "CERTIFICATE REQUEST" {
			return x509.ParseCertificateRequest(result.Bytes)
		}
	}
}

// validateKeyType validates a key_type resource parameter is correct.
func validateKeyType(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	found := false
	for _, w := range []string{"P256", "P384", "2048", "4096", "8192"} {
		if value == w {
			found = true
		}
	}
	if !found {
		errors = append(errors, fmt.Errorf(
			"certificate key type must be one of P256, P384, 2048, 4096, or 8192"))
	}
	return
}

// validateDNSChallengeConfig ensures that the values supplied to the
// dns_challenge resource parameter in the acme_certificate resource
// are string values only.
func validateDNSChallengeConfig(v interface{}, k string) (ws []string, errors []error) {
	value := v.(map[string]interface{})
	bad := false
	for _, w := range value {
		switch w.(type) {
		case string:
			continue
		default:
			bad = true
		}
	}
	if bad {
		errors = append(errors, fmt.Errorf(
			"DNS challenge config map values must be strings only"))
	}
	return
}
