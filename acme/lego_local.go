// Locally re-implemented lego functions.
//
// Code in this file has been adapted from the lego project
// (https://go-acme.github.io/lego/), governed by the MIT license, the body
// of which follows below:
//
// The MIT License (MIT)
//
// Copyright (c) 2017-2024 Ludovic Fernandez
// Copyright (c) 2015-2017 Sebastian Erhart
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package acme

import (
	"crypto"
	"fmt"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/log"
)

type localRenewOptions struct {
	certificate.RenewOptions
	UseARI bool
}

// renewWithOptions re-implements RenewWithOptions out of lego, with some
// updates to allow for the ability to take a RenewalInfo ID.
func renewWithOptions(
	c *certificate.Certifier,
	certRes certificate.Resource,
	options localRenewOptions,
) (*certificate.Resource, error) {
	// Input certificate is PEM encoded.
	// Decode it here as we may need the decoded cert later on in the renewal process.
	// The input may be a bundle or a single certificate.
	certificates, err := certcrypto.ParsePEMBundle(certRes.Certificate)
	if err != nil {
		return nil, err
	}

	x509Cert := certificates[0]
	if x509Cert.IsCA {
		return nil, fmt.Errorf("[%s] Certificate bundle starts with a CA certificate", certRes.Domain)
	}

	// This is just meant to be informal for the user.
	timeLeft := x509Cert.NotAfter.Sub(time.Now().UTC())
	log.Infof("[%s] acme: Trying renewal with %d hours remaining", certRes.Domain, int(timeLeft.Hours()))

	// We always need to request a new certificate to renew.
	// Start by checking to see if the certificate was based off a CSR,
	// and use that if it's defined.
	if len(certRes.CSR) > 0 {
		csr, errP := certcrypto.PemDecodeTox509CSR(certRes.CSR)
		if errP != nil {
			return nil, errP
		}

		request := certificate.ObtainForCSRRequest{CSR: csr}

		request.NotBefore = options.NotBefore
		request.NotAfter = options.NotAfter
		request.Bundle = options.Bundle
		request.PreferredChain = options.PreferredChain
		request.Profile = options.Profile
		request.AlwaysDeactivateAuthorizations = options.AlwaysDeactivateAuthorizations

		if options.UseARI {
			var err error
			request.ReplacesCertID, err = certificate.MakeARICertID(x509Cert)
			if err != nil {
				return nil, fmt.Errorf("error generating ARI cert ID: %w", err)
			}
		}

		return c.ObtainForCSR(request)
	}

	var privateKey crypto.PrivateKey
	if certRes.PrivateKey != nil {
		privateKey, err = certcrypto.ParsePEMPrivateKey(certRes.PrivateKey)
		if err != nil {
			return nil, err
		}
	}

	request := certificate.ObtainRequest{
		Domains:    certcrypto.ExtractDomains(x509Cert),
		PrivateKey: privateKey,
	}

	request.MustStaple = options.MustStaple
	request.NotBefore = options.NotBefore
	request.NotAfter = options.NotAfter
	request.Bundle = options.Bundle
	request.PreferredChain = options.PreferredChain
	request.EmailAddresses = options.EmailAddresses
	request.Profile = options.Profile
	request.AlwaysDeactivateAuthorizations = options.AlwaysDeactivateAuthorizations

	if options.UseARI {
		var err error
		request.ReplacesCertID, err = certificate.MakeARICertID(x509Cert)
		if err != nil {
			return nil, fmt.Errorf("error generating ARI cert ID: %w", err)
		}
	}

	return c.Obtain(request)
}
