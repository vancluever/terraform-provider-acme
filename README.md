Terraform ACME Provider
========================

This repository contains a plugin form of the ACME provider that was proposed
and submitted in [Terraform PR #7058][1].

The Automated Certificate Management Environment (ACME) provider is used to
interact with an ACME Certificate Authority, such as Let's Encrypt
(https://letsencrypt.org/). This provider can be used to both manage
registrations and certificates.

## About ACME

The Automated Certificate Management Environment (ACME) is an emerging
standard for the automation of a domain-validated certificate authority.
Clients set up **registrations** using a private key and contact information,
obtain **authorizations** for domains using a variety of challenges such as
HTTP, HTTPS (TLS), and DNS, with which they can request **certificates**. No
part of this process requires user interaction, a traditional blocker in
obtaining a domain validated certificate.

Currently the major ACME CA is Let's Encrypt (https://letsencrypt.org/),
but the ACME support in Terraform can be configured to use any ACME CA,
including an internal one that is set up using [Boulder][2].

You can read the ACME specification [here][3]. Note that the specification is
currently still in draft, and some features in the specification may not be
fully implemented in ACME CAs like Let's Encrypt or Boulder, and subsequently,
Terraform.

## Installing

See the [Plugin Basics][4] page of the Terraform docs to see how to plunk this
into your config. Check the [releases page][5] of this repo to get releases for
Linux, OS X, and Windows.

## Usage

The following section details the use of the provider and its two resources:
`acme_registration` and `acme_certificate`.

These docs are derived from the middleman templates that were created for the
old PR itself, and can be found in their original form [here][6].

### Note on Examples

**NOTE:** Note that usage examples use the
[Let's Encrypt staging environment][7]. If you are using Let's Encrypt, make
sure you change the URL to the correct endpoint (currently
`https://acme-v01.api.letsencrypt.org/directory`).

### Example Usage

The below example is an end-to-end demonstration of the setup of a basic
certificate, with a little help from the [`tls_private_key`][8] resource:

```
# Create the private key for the registration (not the certificate)
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

# Set up a registration using a private key from tls_private_key
resource "acme_registration" "reg" {
  server_url      = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}

# Create a certificate
resource "acme_certificate" "certificate" {
  server_url                = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem           = "${tls_private_key.private_key.private_key_pem}"
  common_name               = "www.example.com"
  subject_alternative_names = ["www2.example.com"]

  dns_challenge {
    provider = "route53"
  }

  registration_url = "${acme_registration.reg.id}"
}
```

### Registration Credentials

Note that in the above usage example, `server_url` and `account_key_pem` are
required in both resources, and are not configured in a `provider` block.
This allows Terraform the freedom to set up a registration from scratch, with
nothing needing to be done out-of-band - as seen in the example above, the
`account_key_pem` is derived from a [`tls_private_key`][8] resource.

This also means that the two resources can be de-coupled from each other -
there is no need for `acme_registration` or `acme_certificate` to appear in
the same Terraform configuration. One configuration can set up the
registration, with another setting up the certificate, using the registration
from the previous configuration, or one supplied out-of-band.

### The `acme_registration` Resource

Use this resource to create and manage an ACME registration.

**NOTE:** While the ACME draft does contain provisions for deactivating
registrations, implementation is still in development, so if this resource in
Terraform is destroyed, the registration is not completely deleted.

#### Example

The following creates an ACME registration off of a private key generated with
the [`tls_private_key`][8] resource.

```
# Create the private key for the registration (not the certificate)
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

# Set up a registration using a private key from tls_private_key
resource "acme_registration" "reg" {
  server_url      = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}
```

#### Argument Reference

The resource takes the following arguments:

 * `server_url` (Required) - The URL of the ACME directory endpoint.
 * `account_key_pem` (Required) - The private key used to sign requests. This
    is the private key that will be registered to the account.
 * `email_address` (Required) - The email address that will be attached as a
   contact to the registration.

#### Attribute Reference

The following attributes are exported:

 * `id` - The full URL of the registration. Same as `registration_url`.
 * `registration_body`: The raw body of the registration response, in JSON
   format.
 * `registration_url`: The full URL of the registration. Same as `id`.
 * `registration_new_authz_url`: The full URL to the endpoint used to create
   new authorizations.
 * `registration_tos_url`: The full URL to the CA's terms of service.

### The `acme_certificate` Resource

Use this resource to create and manage an ACME TLS certificate.

**NOTE:** Some current ACME CA implementations like [Boulder][2] strip
most of the organization information out of a certificate request's subject,
so you may wish to confirm with the CA what behaviour to expect when using the
`certificate_request_pem` argument with this resource.

#### Example

##### Full example with `common_name` and `subject_alternative_names` and DNS validation

```
# Create the private key for the registration (not the certificate)
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

# Set up a registration using a private key from tls_private_key
resource "acme_registration" "reg" {
  server_url      = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}

# Create a certificate
resource "acme_certificate" "certificate" {
  server_url                = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem           = "${tls_private_key.private_key.private_key_pem}"
  common_name               = "www.example.com"
  subject_alternative_names = ["www2.example.com"]

  dns_challenge {
    provider = "route53"
  }

  registration_url = "${acme_registration.reg.id}"
}
```

##### Above example with HTTP/TLS validation

```
# Create the private key for the registration (not the certificate)
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

# Set up a registration using a private key from tls_private_key
resource "acme_registration" "reg" {
  server_url      = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}

# Create a certificate
resource "acme_certificate" "certificate" {
  server_url                = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem           = "${tls_private_key.private_key.private_key_pem}"
  common_name               = "www.example.com"
  subject_alternative_names = ["www2.example.com"]

  http_challenge_port = 8080
  tls_challenge_port 8443 

  registration_url = "${acme_registration.reg.id}"
}
```

##### Full example with `certificate_request_pem` and DNS validation

```
resource "tls_private_key" "reg_private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  server_url      = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}

resource "tls_private_key" "cert_private_key" {
  algorithm = "RSA"
}

data "tls_cert_request" "req" {
  key_algorithm   = "RSA"
  private_key_pem = "${tls_private_key.cert_private_key.private_key_pem}"
  dns_names       = ["www.example.com", "www2.example.com"]

  subject {
    common_name  = "www.example.com"
  }
}

resource "acme_certificate" "certificate" {
  server_url       = "https://acme-staging.api.letsencrypt.org/directory"
  account_key_pem  = "${tls_private_key.reg_private_key.private_key_pem}"
  certificate_request_pem = "${data.tls_cert_request.req.cert_request_pem}"

  dns_challenge {
    provider = "route53"
  }

  registration_url = "${acme_registration.reg.id}"
}
```

#### Argument Reference

The resource takes the following arguments:

 * `server_url` (Required) - The URL of the ACME directory endpoint.
 * `account_key_pem` (Required) - The private key used to sign requests. This
    will be the private key that will be registered to the account.
 * `registration_url` (Required) - The URL that will be used to fetch the
   registrations's link to perform authorizations.
 * `common_name` - The certificate's common name, the primary domain that the 
   certificate will be recognized for. Required when not specifying a CSR.
 * `subject_alternative_names` - The certificate's subject alternative names,
   domains that this certificate will also be recognized for. Only valid when 
   not specifying a CSR.
 * `key_type` - The key type for the certificate's private key. Can be one of:
   `P256` and `P384` (for ECDSA keys of respective length) or `2048`, `4096`, 
   and `8192` (for RSA keys of respective length). Required when not
   specifying a CSR. The default is `2048` (RSA key of 2048 bits).
 * `certificate_request_pem` - A pre-created certificate request, such as one from
   [`tls_cert_request`][9], or one from an external source, in PEM format.
   Either this, or `common_name`, `key_type`, and optionally
   `subject_alternative_names` needs to be specified.
 * `min_days_remaining` (Optional) - The minimum amount of days remaining before the certificate
   expires before a renewal is attempted. The default is `7`. A value of less
   than 0 means that the certificate will never be renewed.
 * `dns_challenge` (Optional) - Select a [DNS challenge](#using-dns-challenges)
   to use in fulfilling the request. If this is used, HTTP and TLS challenges
   are disabled.
 * `http_challenge_port` (Optional) The port to use in the
   [HTTP challenge](#using-http-and-tls-challenges). Defaults to `80`.
 * `tls_challenge_port` (Optional) The port to use in the
   [TLS challenge](#using-http-and-tls-challenges). Defaults to `443`.

##### Using DNS challenges

ACME and ACME CAs such as Let's Encrypt may support [DNS challenges][10], which
allows operators to respond to authorization challenges by provisioning a TXT
record on a specific domain.

Terraform, making use of [lego][11], responds to DNS challenges automatically
by utilizing one of lego's supported [DNS challenge providers][12]. Most
providers take credentials as environment variables, but if you would rather
use configuration for this purpose, you can through specifying `config` blocks
within a `dns_challenge` block, along with the `provider` parameter.

Example with Route 53 (AWS):

```
# Configure the AWS Provider
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region = "us-east-1"
}

# Create a certificate
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "route53"
    config {
      AWS_ACCESS_KEY_ID     = "${var.aws_access_key}"
      AWS_SECRET_ACCESS_KEY = "${var.aws_secret_key}"
      AWS_DEFAULT_REGION    = "us-east-1"
    }
  }

  ...
}

```

##### Using HTTP and TLS challenges

[HTTP challenges][13] and [TLS challenges][14] work via provisioning a response
message at a specific URL within a well known URI namespace on the hosts
being requested within a certificate.

This presents a unique challenge to Terraform, as normally, Terraform is more
than likely not being run from a live webserver. It is, however, possible to
proxy these requests to the host running Terraform. In order to do this,
perform the following:

 * Set your `http_challenge_port` or `tls_challenge_port` to non-standard
   ports, or leave them if you can assign the Terraform binary the
   `cap_net_bind_service=+ep` - (Linux hosts only).
   [Example configuration here.](#above-example-with-httptls-validation)
 * Proxy the following to the host running Terraform, on the respective ports:
  * All requests on port 80 under the `/.well-known/acme-challenge/` URI
    namespace for HTTP challenges, or:
  * All TLS requests on port 443 for TLS challenges.

#### Attribute Reference

The following attributes are exported:

 * `id` - The full URL of the certificate. Same as `certificate_url`.
 * `certificate_domain` - The common name of the certificate.
 * `certificate_url` - The URL for the certificate. Same as `id`.
 * `account_ref` - The URI of the registration account for this certificate.
   should be the same as `registration_url`.
 * `private_key_pem` - The certificate's private key, in PEM format, if the
   certificate was generated from scratch and not with `certificate_request_pem`. If
   `certificate_request_pem` was used, this will be blank.
 * `certificate_pem` - The certificate in PEM format.
 * `issuer_pem` - The intermediate certificate of the issuer.

## License

```
Copyright 2016 PayByPhone Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[1]: https://github.com/hashicorp/terraform/pull/7058
[2]: https://github.com/letsencrypt/boulder
[3]: https://github.com/ietf-wg-acme/acme
[4]: https://www.terraform.io/docs/plugins/basics.html
[5]: https://github.com/paybyphone/terraform-provider-acme/releases
[6]: website/source/docs/providers/acme
[7]: https://letsencrypt.org/docs/staging-environment/
[8]: https://www.terraform.io/docs/providers/tls/r/private_key.html
[9]: https://www.terraform.io/docs/providers/tls/d/cert_request.html
[10]: https://github.com/ietf-wg-acme/acme/blob/master/draft-ietf-acme-acme.md#dns
[11]: https://github.com/xenolf/lego
[12]: https://godoc.org/github.com/xenolf/lego/providers/dns
[13]: https://github.com/ietf-wg-acme/acme/blob/master/draft-ietf-acme-acme.md#http
[14]: https://github.com/ietf-wg-acme/acme/blob/master/draft-ietf-acme-acme.md#tls-with-server-name-indication-tls-sni
