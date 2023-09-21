package acme

import (
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testPrivateKeyText = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA8XXIc0dO5okTzukP2USWm5tbxY6FQzzvBbOpxIfVpdKpZcKV
HfemqCZEIGu/3P3gI6rAYmDRYvLsbKSjKA5EzuUvVxrLzqPZyFI5mzF0gGEzEvYk
Z+mCPLsS5VwaXCySiz6vIBpItw6KOHByt5v8iMtgppGmjWX5N2oeVZ5314xmXFV3
OMlniyC1uLk6Y/bVtv/vK1mOATXP5vejpjBHdk/VYTTXRZp3zSZllKJbtbt2CxY4
eA55oCc9cfF46rNPsAsiH5iGbBFIIDSqscukZ9BtBZUj+kO+63he0SedzppuosKi
i9YtjgG1Mb81vgMFZ/SQeiR5FONWcH61jTSkiQIDAQABAoIBAQDJVYK8zLq3c5k2
sBLtAUnrmhFdm0b3F7neMT7fhrvYtt1U4njgMf2eu7mWpwGmTXI1i007OqudLB2D
QYxh+/PX6DYfFVLXjLwtUpKCGyyfV2z05JTaqFRWO064PKImNWxD+xKfXAtByDfs
c6bT/pcFoT+H5G7R/DNfx3ZfwfD/oo2aUCQT8PrwzQ9cjJuLYzu5Dwma29Cxtajd
Gsdrd09Qkm0PCM3c0FHG7fV3zq5SNw53tP0U0lNzSzpRiS6wmLAPDy3CcKGaj+9K
5YIE3OoQKRFn7hQkHxgnZlBJJU2BOBAOMJA6s28iRNy3pQOzR0M2kqf+YTQk/i13
if2/mvU5AoGBAPtT9XVbOu6U4Q9WyBSi5nI4AG7gHeJtPC2UWUeaCdj5FJlrEkeD
QZTzqT9KUNu5PfwgsCzCeAzZavQKXDXq7yAtIBIC8bK2sIGhM+bz7Nbu9fPrtmV0
uk5Enlpi2Y/pUFrRTn27FghZAEgWWUF2Drq0kThEZka3jXveBZ7KaHnfAoGBAPXy
3TVsw0Y34ZljmbsHAyT90ZG7FnA3PDXXHOZxEIDo89m8OTGeBW4eqhLvKa3t+thM
oUGyWTtrjKLELuGa8fiDpKq1b8NJqQYB4V0NJlfOYZ6G8Q+hrT+jXTC4+Lb7kmJq
tyIODlyg4B0GQLbFBZXc4FkwWZXxDT+JwKynh36XAoGBAOWsGhm+3yH755fO5FUH
cLRcPPkV0fmDfYThlpz6RZmENbDlyfSUHDB0Yuw1i6Lfq6dmb9jXdkG3xidx+EZF
hXTQCAitrBZ3IOG1YOrjakIYaacYdrxMaZzw1A0hXFRJEGeN8r6vYzkJrFo0IijS
LC+upy7WQujJAIB7qoMr0UHdAoGAEHTEikuRsUQR6zJ32cS5WCNHf2m2MaHwfGW9
QEn2Ybm0fzAR35kEIf8ZQBUSg9m1e/18mKm3QLuMeGOKA3xbjlY4kVd8d+OY1JcR
nilAFIXxkCrVPEeEEQr8NENcGNoyTDV5tWSdX2NAO5DsiY4bNpDFzhHnHJo5WbP8
2VCIR1cCgYAtcMtavC0nIPxmMEkpd9k+0qWcYclt73wr71sQ+kGOU1/M4g8SZh2z
QmXDkRpJf+/xpaeknf6bj6x2r7FXfVoG5vNdB+Cdn3uepkRHPLSStTIwpPVTsQVy
aTVLTgFnTNMM8whCrfR4eBwHVJIejHiA3cl5Ocq/J6u4kgtFkfwKaQ==
-----END RSA PRIVATE KEY-----
`

// testBadCertBundleText is just simply the LE test intermediate cert
const testBadCertBundleText = `
-----BEGIN CERTIFICATE-----
MIIEqzCCApOgAwIBAgIRAIvhKg5ZRO08VGQx8JdhT+UwDQYJKoZIhvcNAQELBQAw
GjEYMBYGA1UEAwwPRmFrZSBMRSBSb290IFgxMB4XDTE2MDUyMzIyMDc1OVoXDTM2
MDUyMzIyMDc1OVowIjEgMB4GA1UEAwwXRmFrZSBMRSBJbnRlcm1lZGlhdGUgWDEw
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDtWKySDn7rWZc5ggjz3ZB0
8jO4xti3uzINfD5sQ7Lj7hzetUT+wQob+iXSZkhnvx+IvdbXF5/yt8aWPpUKnPym
oLxsYiI5gQBLxNDzIec0OIaflWqAr29m7J8+NNtApEN8nZFnf3bhehZW7AxmS1m0
ZnSsdHw0Fw+bgixPg2MQ9k9oefFeqa+7Kqdlz5bbrUYV2volxhDFtnI4Mh8BiWCN
xDH1Hizq+GKCcHsinDZWurCqder/afJBnQs+SBSL6MVApHt+d35zjBD92fO2Je56
dhMfzCgOKXeJ340WhW3TjD1zqLZXeaCyUNRnfOmWZV8nEhtHOFbUCU7r/KkjMZO9
AgMBAAGjgeMwgeAwDgYDVR0PAQH/BAQDAgGGMBIGA1UdEwEB/wQIMAYBAf8CAQAw
HQYDVR0OBBYEFMDMA0a5WCDMXHJw8+EuyyCm9Wg6MHoGCCsGAQUFBwEBBG4wbDA0
BggrBgEFBQcwAYYoaHR0cDovL29jc3Auc3RnLXJvb3QteDEubGV0c2VuY3J5cHQu
b3JnLzA0BggrBgEFBQcwAoYoaHR0cDovL2NlcnQuc3RnLXJvb3QteDEubGV0c2Vu
Y3J5cHQub3JnLzAfBgNVHSMEGDAWgBTBJnSkikSg5vogKNhcI5pFiBh54DANBgkq
hkiG9w0BAQsFAAOCAgEABYSu4Il+fI0MYU42OTmEj+1HqQ5DvyAeyCA6sGuZdwjF
UGeVOv3NnLyfofuUOjEbY5irFCDtnv+0ckukUZN9lz4Q2YjWGUpW4TTu3ieTsaC9
AFvCSgNHJyWSVtWvB5XDxsqawl1KzHzzwr132bF2rtGtazSqVqK9E07sGHMCf+zp
DQVDVVGtqZPHwX3KqUtefE621b8RI6VCl4oD30Olf8pjuzG4JKBFRFclzLRjo/h7
IkkfjZ8wDa7faOjVXx6n+eUQ29cIMCzr8/rNWHS9pYGGQKJiY2xmVC9h12H99Xyf
zWE9vb5zKP3MVG6neX1hSdo7PEAb9fqRhHkqVsqUvJlIRmvXvVKTwNCP3eCjRCCI
PTAvjV+4ni786iXwwFYNz8l3PmPLCyQXWGohnJ8iBm+5nk7O2ynaPVW0U2W+pt2w
SVuvdDM5zGv2f9ltNWUiYZHJ1mmO97jSY/6YfdOUH66iRtQtDkHBRdkNBsMbD+Em
2TgBldtHNSJBfB3pm9FblgOcJ0FSWcUDWJ7vO0+NTXlgrRofRT6pVywzxVo6dND0
WzYlTWeUVsO40xJqhgUQRER9YLOLxJ0O6C8i0xFxAMKOtSdodMB3RIwt7RFQ0uyt
n5Z5MqkYhlMI3J1tPRTp1nEt9fyGspBOO05gi148Qasp+3N+svqKomoQglNoAxU=
-----END CERTIFICATE-----
`

const testPaddingBundle = `
-----BEGIN CERTIFICATE-----
MIIDSjCCAjICCQDjsMLnU/0KpzANBgkqhkiG9w0BAQsFADBnMQswCQYDVQQGEwJV
SzEQMA4GA1UECAwHRW5nbGFuZDEPMA0GA1UEBwwGTG9uZG9uMSEwHwYDVQQKDBhJ
bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0x
OTAzMDQyMjE2MzFaFw0yNDAzMDIyMjE2MzFaMGcxCzAJBgNVBAYTAlVLMRAwDgYD
VQQIDAdFbmdsYW5kMQ8wDQYDVQQHDAZMb25kb24xITAfBgNVBAoMGEludGVybmV0
IFdpZGdpdHMgUHR5IEx0ZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0A0z2uLCRrw8DXKgG1UMBRlmRt3TXEoDSPSJ
y5Awp3fJG8b5+MvU2QufUrtk3XAwG5q7dBFpV+bAQGds/1cmMNjRRXHby0mmk7v6
b57rhAuaD4VLXa7/pJDEIGDaQ7NmFe4vgO8mup5HDw7C9VZI25ou70ajZZvpBzyd
rd6pgfCq4fcCZRS56rdzcO4n48HhXHjSOaSyqHHGkLJpVc7qqc3OmJQoeUgn+1xc
MkSBEiZA3XvISHegc/5s5wB/aMDfegJEe2ZA7Ae9gAmCczPFGRtZUVU9J9UC1jVS
jIWsIyOsdHc18TxafEBjcBctpuNnvhOshbOYsd1HLbmZB/4GuwIDAQABMA0GCSqG
SIb3DQEBCwUAA4IBAQB82OHtGfw2tDsDsqj2IvkHqGw8bvxwXZ5KVRdGVg90AD+f
rFKS5qF3JxspcUjDHwFAyZo/mTXMOzRFQAytcuID4qijVRLRaM8dnFWvzwhFo0Kq
UEdVfmq2ANmhqWI5j87BoPu2GGcZ+xlzW7axl2tFOj4g1xOW1Vd/CVuPBfMHZ5JD
WyQVnPXi3plkGnIW/P5R1NHYqXIb0HW8xjzqRmbbQOW5eJ+Hy1Id4O0pnOVlNjDl
Rb4po3kWJaGezLmNO+JyUOr5HnCjLSD/4WNHGAJbfOVES4hOz+oiTlWd7HflS+00
iITbUq4IV5mAI5yceK+3rYWGYG47cu0BG9ngevUZ
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIDSjCCAjICCQDjsMLnU/0KpzANBgkqhkiG9w0BAQsFADBnMQswCQYDVQQGEwJV
SzEQMA4GA1UECAwHRW5nbGFuZDEPMA0GA1UEBwwGTG9uZG9uMSEwHwYDVQQKDBhJ
bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0x
OTAzMDQyMjE2MzFaFw0yNDAzMDIyMjE2MzFaMGcxCzAJBgNVBAYTAlVLMRAwDgYD
VQQIDAdFbmdsYW5kMQ8wDQYDVQQHDAZMb25kb24xITAfBgNVBAoMGEludGVybmV0
IFdpZGdpdHMgUHR5IEx0ZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0A0z2uLCRrw8DXKgG1UMBRlmRt3TXEoDSPSJ
y5Awp3fJG8b5+MvU2QufUrtk3XAwG5q7dBFpV+bAQGds/1cmMNjRRXHby0mmk7v6
b57rhAuaD4VLXa7/pJDEIGDaQ7NmFe4vgO8mup5HDw7C9VZI25ou70ajZZvpBzyd
rd6pgfCq4fcCZRS56rdzcO4n48HhXHjSOaSyqHHGkLJpVc7qqc3OmJQoeUgn+1xc
MkSBEiZA3XvISHegc/5s5wB/aMDfegJEe2ZA7Ae9gAmCczPFGRtZUVU9J9UC1jVS
jIWsIyOsdHc18TxafEBjcBctpuNnvhOshbOYsd1HLbmZB/4GuwIDAQABMA0GCSqG
SIb3DQEBCwUAA4IBAQB82OHtGfw2tDsDsqj2IvkHqGw8bvxwXZ5KVRdGVg90AD+f
rFKS5qF3JxspcUjDHwFAyZo/mTXMOzRFQAytcuID4qijVRLRaM8dnFWvzwhFo0Kq
UEdVfmq2ANmhqWI5j87BoPu2GGcZ+xlzW7axl2tFOj4g1xOW1Vd/CVuPBfMHZ5JD
WyQVnPXi3plkGnIW/P5R1NHYqXIb0HW8xjzqRmbbQOW5eJ+Hy1Id4O0pnOVlNjDl
Rb4po3kWJaGezLmNO+JyUOr5HnCjLSD/4WNHGAJbfOVES4hOz+oiTlWd7HflS+00
iITbUq4IV5mAI5yceK+3rYWGYG47cu0BG9ngevUZ
-----END CERTIFICATE-----`

func registrationResourceData() *schema.ResourceData {
	r := resourceACMERegistration()
	d := r.TestResourceData()

	d.SetId("regurl")
	d.Set("account_key_pem", testPrivateKeyText)
	d.Set("email_address", "nobody@example.com")

	return d
}

func blankCertificateResource() *schema.ResourceData {
	r := resourceACMECertificate()
	d := r.TestResourceData()
	d.Set("account_key_pem", testPrivateKeyText)
	return d
}

// registrationResourceDataDefaultConfig returns the schema.ResourceData for a
// registration based off of a default config.
//
// Note that the attributes in the result set may vary quite differently from
// registrationResourceData since it actually does a diff based on the
// attributes in the schema, versus working with a blank ResourceData.
func registrationResourceDataDefaultConfig(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, resourceACMERegistration().Schema, make(map[string]interface{}))
}

// certificateResourceDataDefaultConfig returns the schema.ResourceData for a
// certificate based off of a default config.
//
// Note that the attributes in the result set may vary quite differently from
// blankCertificateResource since it actually does a diff based on the
// attributes in the schema, versus working with a blank ResourceData.
func certificateResourceDataDefaultConfig(t *testing.T) *schema.ResourceData {
	// NOTE: we set cert_timeout to a static value here to try and
	// bring meaning to the TestExpandACMEClient_config_certTimeout_default test.
	// Since lego automatically sets the default to 30 seconds, we need to be
	// able to differentiate between a schema that does not have this value in
	// the client (registrations) versus one that does (certificates) as we can't
	// test on the zero value for registrations.
	return schema.TestResourceDataRaw(t, resourceACMECertificate().Schema, map[string]interface{}{"cert_timeout": 90})
}

func TestACME_expandACMEUser(t *testing.T) {
	d := registrationResourceData()
	u, err := expandACMEUser(d)
	if err != nil {
		t.Fatalf("fatal: %s", err.Error())
	}

	if u.GetEmail() != "nobody@example.com" {
		t.Fatalf("Expected email to be nobody@example.com, got %s", u.GetEmail())
	}

	key, err := privateKeyFromPEM([]byte(testPrivateKeyText))
	if err != nil {
		t.Fatalf("fatal: %s", err.Error())
	}

	if reflect.DeepEqual(key, u.GetPrivateKey()) == false {
		t.Fatalf("Expected private key to be %#v, got %#v", key, u.GetPrivateKey())
	}
}

func TestACME_expandACMEUser_badKey(t *testing.T) {
	d := registrationResourceData()
	d.Set("account_key_pem", "bad")
	_, err := expandACMEUser(d)
	if err == nil {
		t.Fatalf("expected error due to bad key")
	}
}

func TestACME_expandACMEClient_badKey(t *testing.T) {
	d := registrationResourceData()
	d.Set("account_key_pem", "bad")
	_, _, err := expandACMEClient(d, &Config{ServerURL: "https://acme-staging.api-v02.letsencrypt.org/directory"}, true)
	if err == nil {
		t.Fatalf("expected error due to bad key")
	}
}

func TestACME_certDaysRemaining_noCertData(t *testing.T) {
	c := &certificate.Resource{}
	_, err := certDaysRemaining(c)
	if err == nil {
		t.Fatalf("expected error due to bad cert data")
	}
}

func TestACME_parsePEMBundle_noData(t *testing.T) {
	b := []byte{}
	_, err := parsePEMBundle(b)
	if err == nil {
		t.Fatalf("expected error due to no PEM data")
	}
}

func TestACME_saveCertificateResource_badCert(t *testing.T) {
	b := testBadCertBundleText
	c := &certificate.Resource{
		Certificate: []byte(b),
	}
	d := blankCertificateResource()
	err := saveCertificateResource(d, c, "")
	if err == nil {
		t.Fatalf("expected error due to bad cert data")
	}
}

func TestACME_certDaysRemaining_CACert(t *testing.T) {
	b := testBadCertBundleText
	c := &certificate.Resource{
		Certificate: []byte(b),
	}
	_, err := certDaysRemaining(c)
	if err == nil {
		t.Fatalf("expected error due to cert being a CA")
	}
}

func TestACME_splitPEMBundle_noData(t *testing.T) {
	b := []byte{}
	_, _, _, err := splitPEMBundle(b)
	if err == nil {
		t.Fatalf("expected error due to no PEM data")
	}
}

func TestACME_splitPEMBundle_CAFirst(t *testing.T) {
	b := testBadCertBundleText + testBadCertBundleText
	_, _, _, err := splitPEMBundle([]byte(b))
	if err == nil {
		t.Fatalf("expected error due to CA cert being first")
	}
}

func TestACME_bundleToPKCS12_base64IsPadded(t *testing.T) {
	b := testPaddingBundle
	key := testPrivateKeyText
	pfxBase64, err := bundleToPKCS12([]byte(b), []byte(key), "")

	if err != nil {
		t.Fatalf("bad: %#v", err)
	}

	// testPaddingBundle requires padding to 4 bytes so will end in =
	if math.Remainder(float64(len(pfxBase64)), 4) != 0 && !strings.HasSuffix(string(pfxBase64), "=") {
		t.Fatalf("p12 base64 encoded certificate should be padded")
	}
}

func TestACME_splitPEMBundle_singleCert(t *testing.T) {
	b := testBadCertBundleText
	_, _, _, err := splitPEMBundle([]byte(b))
	if err == nil {
		t.Fatalf("expected error due to only one cert being present")
	}
}

func TestACME_validateKeyType(t *testing.T) {
	s := "2048"

	_, errs := validateKeyType(s, "key_type")
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
}

func TestACME_validateKeyType_invalid(t *testing.T) {
	s := "512"

	_, errs := validateKeyType(s, "key_type")
	if len(errs) < 1 {
		t.Fatalf("should have given an error")
	}
}

func TestACME_validateDNSChallengeConfig(t *testing.T) {
	m := map[string]interface{}{
		"AWS_FOO": "bar",
	}

	_, errs := validateDNSChallengeConfig(m, "config")
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
}

func TestACME_validateDNSChallengeConfig_invalid(t *testing.T) {
	s := map[string]interface{}{
		"AWS_FOO": 1,
	}

	_, errs := validateDNSChallengeConfig(s, "config")
	if len(errs) < 1 {
		t.Fatalf("should have given an error")
	}
}

func TestExpandACMEClient_config_certTimeout_default(t *testing.T) {
	testCases := []struct {
		desc     string
		f        func(t *testing.T) *schema.ResourceData
		expected time.Duration
	}{
		{
			desc:     "registration",
			f:        registrationResourceDataDefaultConfig,
			expected: time.Second * 30,
		},
		{
			desc:     "certificate",
			f:        certificateResourceDataDefaultConfig,
			expected: time.Second * 90,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			c := expandACMEClient_config(tc.f(t), &Config{}, &acmeUser{})
			if tc.expected != c.Certificate.Timeout {
				t.Fatalf("expected timeout to be %s, got %s", tc.expected, c.Certificate.Timeout)
			}
		})
	}
}
