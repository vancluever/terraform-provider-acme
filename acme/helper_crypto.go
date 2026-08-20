package acme

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"software.sslmate.com/src/go-pkcs12"
)

const (
	keyAlgorithmRSA     = "RSA"
	keyAlgorithmECDSA   = "ECDSA"
	keyAlgorithmED25519 = "ED25519"

	keyECDSACurveP224 = "P224"
	keyECDSACurveP256 = "P256"
	keyECDSACurveP384 = "P384"
	keyECDSACurveP521 = "P521"
)

const (
	preamblePKCS8PrivateKey = "PRIVATE KEY"
	preambleRSAPrivateKey   = "RSA PRIVATE KEY"
	preambleECPrivateKey    = "EC PRIVATE KEY"

	preambleCertificate        = "CERTIFICATE"
	preambleCertificateRequest = "CERTIFICATE REQUEST"
)

func generatePrivateKey(algo string, rsaBits int, ecCruve string) (string, error) {
	var privateKeyPem *pem.Block
	switch algo {
	case keyAlgorithmRSA:
		k, err := rsa.GenerateKey(rand.Reader, rsaBits)
		if err != nil {
			return "", fmt.Errorf("error generating RSA private key: %w", err)
		}

		privateKeyPem = &pem.Block{
			Type:  preambleRSAPrivateKey,
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}

	case keyAlgorithmECDSA:
		var k *ecdsa.PrivateKey
		var err error
		switch ecCruve {
		case keyECDSACurveP224:
			k, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
		case keyECDSACurveP256:
			k, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		case keyECDSACurveP384:
			k, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		case keyECDSACurveP521:
			k, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		default:
			return "", fmt.Errorf("invalid EC curve %q", ecCruve)
		}

		if err != nil {
			return "", fmt.Errorf("error generating ECDSA private key: %w", err)
		}

		keyBytes, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return "", fmt.Errorf("error marshaling ECDSA private key: %w", err)
		}

		privateKeyPem = &pem.Block{
			Type:  preambleECPrivateKey,
			Bytes: keyBytes,
		}

	case keyAlgorithmED25519:
		_, k, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return "", fmt.Errorf("error generating ED25519 private key: %w", err)
		}

		keyBytes, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return "", fmt.Errorf("error marshaling ED25519 private key: %w", err)
		}

		privateKeyPem = &pem.Block{
			Type:  preamblePKCS8PrivateKey,
			Bytes: keyBytes,
		}

	default:
		return "", fmt.Errorf("invalid key algorithm %q", algo)
	}

	return string(pem.EncodeToMemory(privateKeyPem)), nil
}

// privateKeyFromPEM converts a PEM block into a crypto.Signer.
func privateKeyFromPEM(pemData []byte) (crypto.Signer, error) {
	var result *pem.Block
	rest := pemData
	for {
		result, rest = pem.Decode(rest)
		if result == nil {
			return nil, fmt.Errorf("cannot decode supplied PEM data")
		}
		switch result.Type {
		case preamblePKCS8PrivateKey:
			k, err := x509.ParsePKCS8PrivateKey(result.Bytes)
			if err != nil {
				return nil, err
			}

			return k.(crypto.Signer), nil

		case preambleRSAPrivateKey:
			return x509.ParsePKCS1PrivateKey(result.Bytes)
		case preambleECPrivateKey:
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
		if result.Type == preambleCertificateRequest {
			return x509.ParseCertificateRequest(result.Bytes)
		}
	}
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
func splitPEMBundle(bundle []byte) (
	cert []byte,
	certNotBefore string,
	certNotAfter string,
	certSerial string,
	issuer []byte,
	err error,
) {
	cb, err := parsePEMBundle(bundle)
	if err != nil {
		return
	}

	// lego always returns the issued cert first, if the CA is first there is a problem
	if cb[0].IsCA {
		err = fmt.Errorf("first certificate is a CA certificate")
		return
	}

	cert = pem.EncodeToMemory(&pem.Block{Type: preambleCertificate, Bytes: cb[0].Raw})
	certNotBefore = cb[0].NotBefore.Format(time.RFC3339)
	certNotAfter = cb[0].NotAfter.Format(time.RFC3339)
	certSerial = cb[0].SerialNumber.String()
	issuer = make([]byte, 0)
	for _, ic := range cb[1:] {
		issuer = append(issuer, pem.EncodeToMemory(&pem.Block{Type: preambleCertificate, Bytes: ic.Raw})...)
	}

	return
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

		if certDERBlock.Type == preambleCertificate {
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

	pfxData, err := pkcs12.Modern2023.Encode(pk, cb[0], cb[1:], password)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(pfxData)))
	base64.StdEncoding.Encode(buf, pfxData)
	return buf, nil
}
