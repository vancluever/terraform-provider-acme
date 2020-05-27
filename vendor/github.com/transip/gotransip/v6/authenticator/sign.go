package authenticator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	asn1Header = []byte{
		0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03,
		0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40,
	}
	// ErrDecodingPrivateKey will be thrown when an invalid private key has been given
	ErrDecodingPrivateKey = errors.New("could not decode private key")
)

func signWithKey(body []byte, key []byte) (string, error) {
	// create SHA512 hash of given parameters
	h := sha512.New()
	_, err := h.Write(body)
	if err != nil {
		return "", fmt.Errorf("signing error during request body writing: %w", err)
	}

	// prefix ASN1 header to SHA512 hash
	digest := append(asn1Header, h.Sum(nil)...)

	// prepare key struct
	block, _ := pem.Decode(key)
	if block == nil {
		return "", ErrDecodingPrivateKey
	}
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %w", err)
	}

	pkey := parsed.(*rsa.PrivateKey)

	enc, err := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.Hash(0), digest)
	if err != nil {
		return "", fmt.Errorf("could not sign data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}
