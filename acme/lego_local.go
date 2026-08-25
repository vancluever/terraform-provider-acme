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
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/certcrypto"
	"github.com/go-acme/lego/v5/certificate"
	"github.com/go-acme/lego/v5/lego"
	"github.com/go-acme/lego/v5/log"
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
		return nil, fmt.Errorf(
			"acme: Certificate bundle starts with a CA certificate, domains=%s",
			strings.Join(certRes.Domains, ", "),
		)
	}

	// This is just meant to be informal for the user.
	timeLeft := x509Cert.NotAfter.Sub(time.Now().UTC())
	log.Info(
		"acme: Trying renewal",
		log.DomainsAttr(certRes.Domains),
		log.DurationAttr("time_remaining", timeLeft),
	)

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
		request.EnableCommonName = true
		request.Profile = options.Profile
		request.AlwaysDeactivateAuthorizations = options.AlwaysDeactivateAuthorizations

		if options.UseARI {
			var err error
			request.ReplacesCertID, err = api.MakeARICertID(x509Cert)
			if err != nil {
				return nil, fmt.Errorf("error generating ARI cert ID: %w", err)
			}
		}

		return c.ObtainForCSR(context.TODO(), request)
	}

	var privateKey crypto.Signer
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
	request.EnableCommonName = true
	request.EmailAddresses = options.EmailAddresses
	request.Profile = options.Profile
	request.AlwaysDeactivateAuthorizations = options.AlwaysDeactivateAuthorizations

	if options.UseARI {
		var err error
		request.ReplacesCertID, err = api.MakeARICertID(x509Cert)
		if err != nil {
			return nil, fmt.Errorf("error generating ARI cert ID: %w", err)
		}
	}

	return c.Obtain(context.TODO(), request)
}

// getCertificateARIRenewalInfo is a client-less function that can fetch the
// ARI information for a specific certificate at a specific ACME directory URL.
//
// ARI does not require authentication, e.g., requests to the endpoint do not
// need to be signed, so this helps facilitate this by just allowing a lookup
// without the standard lego.Client.
func getCertificateARIRenewalInfo(dirURL string, cert *x509.Certificate) (*acme.ExtendedRenewalInfo, error) {
	certID, err := api.MakeARICertID(cert)
	if err != nil {
		return nil, fmt.Errorf("error making certID: %w", err)
	}

	client := createDefaultARIHTTPClient()

	ariURL, err := ariEndpointForDirectory(client, dirURL)
	if err != nil {
		return nil, fmt.Errorf("error getting ARI URL: %w", err)
	}

	return getCertificateARIRenewalInfoInner(client, ariURL, certID)
}

// ariEndpointForDirectory is a client-less function that can fetch the ARI
// endpoint for an ACME directory.
//
// ARI does not require authentication, e.g., requests to the endpoint do not
// need to be signed, so this helps facilitate this by just allowing a lookup
// without the standard lego.Client.
func ariEndpointForDirectory(client *http.Client, dirURL string) (string, error) {
	resp, err := doGetRequest(client, dirURL)
	if err != nil {
		return "", fmt.Errorf("directory request failed: %w", err)
	}

	defer resp.Body.Close()

	var resultRaw struct {
		RenewalInfo string `json:"renewalInfo"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&resultRaw); err != nil {
		return "", fmt.Errorf("failed to read directory response body: %w", err)
	}

	return resultRaw.RenewalInfo, nil
}

// getCertificateARIRenewalInfoInner is just the inner request to the ARI
// endpoint from ariEndpointForDirectory.
func getCertificateARIRenewalInfoInner(
	client *http.Client,
	ariURL string,
	ariCertID string,
) (*acme.ExtendedRenewalInfo, error) {
	resp, err := doGetRequest(client, ariURL+"/"+ariCertID)
	if err != nil {
		return nil, fmt.Errorf("ARI request failed: %w", err)
	}

	defer resp.Body.Close()

	result := new(acme.ExtendedRenewalInfo)

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to read ARI response body: %w", err)
	}

	return result, nil
}

// Note that the client bits are cribbed from lego pretty much directly, so
// that we can mirror what they do.

const (
	caCertificatesEnvVar = "LEGO_CA_CERTIFICATES"
	caSystemCertPool     = "LEGO_CA_SYSTEM_CERT_POOL"
	caServerNameEnvVar   = "LEGO_CA_SERVER_NAME"
)

func createDefaultARIHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 2 * time.Minute,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			TLSClientConfig: &tls.Config{
				ServerName: os.Getenv(caServerNameEnvVar),
				RootCAs:    initCertPool(),
			},
		},
	}
}

func initCertPool() *x509.CertPool {
	customCACertsPath := os.Getenv(caCertificatesEnvVar)
	if customCACertsPath == "" {
		return nil
	}

	useSystemCertPool, _ := strconv.ParseBool(os.Getenv(caSystemCertPool))

	caCerts := strings.Split(customCACertsPath, string(os.PathListSeparator))

	certPool, err := lego.CreateCertPool(caCerts, useSystemCertPool)
	if err != nil {
		panic(fmt.Sprintf("create certificates pool: %v", err))
	}

	return certPool
}

// Run up a GET request with the user-agent set.
func doGetRequest(client *http.Client, reqURL string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", FormatUserAgentLong())

	return client.Do(req)
}
