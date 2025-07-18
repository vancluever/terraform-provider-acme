package acme

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testPrivateKeyPKCS1Text = `
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

const testPrivateKeyPKCS8Text = `
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMYmYAydUoJY0rWW
rtJqguhFlcF0Q/K47L42q2nDz3Tfg+1eZ2lFygd9rH43QkbU7lZMZ9e/A4ZVdt5i
1aKoxWIapYF3F78HBrGvG379+CAeMtSPFW2EkZpJf+Yv8GVg0C5bCBNnmNk2L20g
r+kB5wrvAh4oFd+6sr0huugLXi1/AgMBAAECgYBw40ILTcHJAwOgcqVMuWO1Iper
7Cod6V7vC0Ri8CyL4B9QJ11w61KrK90O1zqKEhtqzQLINkmyyZP8JICjt9UjC2Ll
S3ore0PP7Xd4eh6oS6hsqdMc2R63kTVA+Lr1/JYfOobioQ7d0UGw0WTrj4+L7uyW
LHyGd0iHV7mKJ5YiQQJBAOePBGC/Mu2w5d5/6gecHs7r5Ck4G7M9a40H+ptqBqYT
fXQpf+L37njnx3tZERWEfMlf68EwTZOkjThdrZCgo+ECQQDbEJ1kdV2DXI/qM0EV
RjghwIeMDLv9D/g90uXCVxw+0AAHJF3/MsMDHAoV8B7BF63h41pt9EH+sQDhnf7X
TP1fAkBpP5whzUX8u5b/1uwsoU1vh9Cg25vbkGM+Kw5BbaOwANPY5LP4GfEOi2sk
KYuWWC3P6gViPe5E2VpG8G1fe2SBAkA2+4+VhEOpUdUpQh4Gue4iwpEC3LteQ+DZ
m5JhWb3UIh6vrDgPcm0x3ZrGcNM3Qbs54/dxe4oI4+JFvoMVBNTfAkEA5gXaJYQE
T9BKDLYkDaVra4zk1hNvSAPKrHNiWjP1clCAc+lcQ9vvihSVertNGy9a8wCm+Htp
xe9MEyzqTI7zow==
-----END PRIVATE KEY-----
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
	d.Set("account_key_pem", testPrivateKeyPKCS1Text)
	d.Set("email_address", "nobody@example.com")

	return d
}

func blankCertificateResource() *schema.ResourceData {
	r := resourceACMECertificate()
	d := r.TestResourceData()
	d.Set("account_key_pem", testPrivateKeyPKCS1Text)
	return d
}

// registrationResourceDataDefaultConfig returns the schema.ResourceData for a
// registration based off of a default config.
//
// Note that the attributes in the result set may vary quite differently from
// registrationResourceData since it actually does a diff based on the
// attributes in the schema, versus working with a blank ResourceData.
func registrationResourceDataDefaultConfig(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, resourceACMERegistration().Schema, make(map[string]any))
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
	return schema.TestResourceDataRaw(t, resourceACMECertificate().Schema, map[string]any{"cert_timeout": 90})
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

	key, err := privateKeyFromPEM([]byte(testPrivateKeyPKCS1Text))
	if err != nil {
		t.Fatalf("fatal: %s", err.Error())
	}

	if reflect.DeepEqual(key, u.GetPrivateKey()) == false {
		t.Fatalf("Expected private key to be %#v, got %#v", key, u.GetPrivateKey())
	}
}

func TestACME_expandACMEUser_PKCS8(t *testing.T) {
	d := registrationResourceData()
	d.Set("account_key_pem", testPrivateKeyPKCS8Text)
	u, err := expandACMEUser(d)
	if err != nil {
		t.Fatalf("fatal: %s", err.Error())
	}

	if u.GetEmail() != "nobody@example.com" {
		t.Fatalf("Expected email to be nobody@example.com, got %s", u.GetEmail())
	}

	key, err := privateKeyFromPEM([]byte(testPrivateKeyPKCS8Text))
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
	_, err := certDaysRemaining(c, time.Now())
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
	_, err := certDaysRemaining(c, time.Now())
	if err == nil {
		t.Fatalf("expected error due to cert being a CA")
	}
}

func TestACME_splitPEMBundle_noData(t *testing.T) {
	b := []byte{}
	_, _, _, _, err := splitPEMBundle(b)
	if err == nil {
		t.Fatalf("expected error due to no PEM data")
	}
}

func TestACME_splitPEMBundle_CAFirst(t *testing.T) {
	b := testBadCertBundleText + testBadCertBundleText
	_, _, _, _, err := splitPEMBundle([]byte(b))
	if err == nil {
		t.Fatalf("expected error due to CA cert being first")
	}
}

func TestACME_bundleToPKCS12_base64IsPadded(t *testing.T) {
	b := testPaddingBundle
	key := testPrivateKeyPKCS1Text
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
	_, _, _, _, err := splitPEMBundle([]byte(b))
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
	m := map[string]any{
		"AWS_FOO": "bar",
	}

	_, errs := validateDNSChallengeConfig(m, "config")
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
}

func TestACME_validateDNSChallengeConfig_invalid(t *testing.T) {
	s := map[string]any{
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

func TestGeneratePrivateKey(t *testing.T) {
	testCases := []struct {
		desc      string
		algo      string
		rsaBits   int
		ecCurve   string
		expectErr string
	}{
		{
			desc:    "RSA",
			algo:    keyAlgorithmRSA,
			rsaBits: 4096,
		},
		{
			desc:    "ECDSA",
			algo:    keyAlgorithmECDSA,
			ecCurve: keyECDSACurveP384,
		},
		{
			desc: "ED25519",
			algo: keyAlgorithmED25519,
		},
		{
			desc:    "RSA - odd key len",
			algo:    keyAlgorithmRSA,
			rsaBits: 1111,
		},
		{
			desc:    "RSA - odd key len",
			algo:    keyAlgorithmRSA,
			rsaBits: 1111,
		},
		{
			desc:      "ECDSA - unknown curve type",
			algo:      keyAlgorithmECDSA,
			ecCurve:   "foobar",
			expectErr: "invalid EC curve \"foobar\"",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotKey, err := generatePrivateKey(tc.algo, tc.rsaBits, tc.ecCurve)
			if err != nil {
				if tc.expectErr != "" {
					if tc.expectErr != err.Error() {
						t.Fatalf("expected error:\n\t%s\ngot error:\n\t%s", tc.expectErr, err)
					}

					return
				} else {
					t.Fatal(err)
				}
			}

			k, err := privateKeyFromPEM([]byte(gotKey))
			if err != nil {
				t.Fatal(err)
			}

			var gotKeyType string
			var gotRsaBits int
			var gotEcCurve string
			switch l := k.(type) {
			case *rsa.PrivateKey:
				gotKeyType = keyAlgorithmRSA
				gotRsaBits = l.N.BitLen()

			case *ecdsa.PrivateKey:
				gotKeyType = keyAlgorithmECDSA
				switch l.Curve {
				case elliptic.P224():
					gotEcCurve = keyECDSACurveP224
				case elliptic.P256():
					gotEcCurve = keyECDSACurveP256
				case elliptic.P384():
					gotEcCurve = keyECDSACurveP384
				case elliptic.P521():
					gotEcCurve = keyECDSACurveP521
				default:
					t.Fatalf("expected EDCSA curve %T", l.Curve)
				}

			case ed25519.PrivateKey:
				gotKeyType = keyAlgorithmED25519

			default:
				t.Fatalf("unexpected key type %T was generated", k)
			}

			if tc.algo != gotKeyType {
				t.Fatalf("expected key type %q, got %q", tc.algo, gotKeyType)
			}
			if tc.rsaBits != gotRsaBits {
				t.Fatalf("expected key type %d, got %d", tc.rsaBits, gotRsaBits)
			}
			if tc.ecCurve != gotEcCurve {
				t.Fatalf("expected EC curve to be %q, got %q", tc.algo, gotKeyType)
			}

		})
	}
}
