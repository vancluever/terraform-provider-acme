# acme_certificate

The `acme_certificate` resource can be used to create and manage an ACME TLS
certificate.

## Example

The below example creates both an account and certificate within the same
configuration. The account is created using the
[`acme_registration`][resource-registration] resource.

-> When creating accounts and certificates within the same configuration, ensure
that you reference the
[`account_key_pem`][resource-registration-account-key-pem] argument in the
`acme_registration` resource as the corresponding
[`account_key_pem`](#account_key_pem) argument in the `acme_certificate`
resource. This will ensure that the account gets created before the certificate
and avoid errors.

[resource-registration]: ./registration.md
[resource-registration-account-key-pem]: ./registration.md#account_key_pem

```hcl
provider "acme" {
  server_url = "https://acme-staging-v02.api.letsencrypt.org/directory"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.private_key.private_key_pem
  email_address   = "nobody@example.com"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = acme_registration.reg.account_key_pem
  common_name               = "www.example.com"
  subject_alternative_names = ["www2.example.com"]

  dns_challenge {
    provider = "route53"
  }
}
```

### Using an external CSR

The `acme_certificate` resource can also take an external CSR. In this example,
we create one using [`tls_cert_request`][tls-cert-request] first, before
supplying it to the [`certificate_request_pem`](#certificate_request_pem)
argument.

[tls-cert-request]: https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/cert_request

-> **NOTE:** Some current ACME CA implementations (including Let's Encrypt)
strip most of the organization information out of a certificate request
subject.  You may wish to confirm with the CA what behavior to expect when
using the `certificate_request_pem` argument with this resource.

~> **NOTE:** It is not a good practice to use the same private key for both
your account and your certificate. Make sure you use different keys.

```hcl
provider "acme" {
  server_url = "https://acme-staging-v02.api.letsencrypt.org/directory"
}

resource "tls_private_key" "reg_private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.reg_private_key.private_key_pem
  email_address   = "nobody@example.com"
}

resource "tls_private_key" "cert_private_key" {
  algorithm = "RSA"
}

resource "tls_cert_request" "req" {
  key_algorithm   = "RSA"
  private_key_pem = tls_private_key.cert_private_key.private_key_pem
  dns_names       = ["www.example.com", "www2.example.com"]

  subject {
    common_name = "www.example.com"
  }
}

resource "acme_certificate" "certificate" {
  account_key_pem         = acme_registration.reg.account_key_pem
  certificate_request_pem = tls_cert_request.req.cert_request_pem

  dns_challenge {
    provider = "route53"
  }
}
```

## Argument Reference

The resource takes the following arguments:

-> At least one challenge type (`dns_challenge`, `http_challenge`,
`http_webroot_challenge`, `http_memcached_challenge`, or `tls_challenge`) must
be specified. It's recommended you use `dns_challenge` whenever possible).

* `account_key_pem` (Required) - The private key of the account that is
  requesting the certificate. Forces a new resource when changed.
* `common_name` - The certificate's common name, the primary domain that the
  certificate will be recognized for. Required when not specifying a CSR. Forces
  a new resource when changed.
* `subject_alternative_names` - The certificate's subject alternative names,
  domains that this certificate will also be recognized for. Only valid when not
  specifying a CSR. Forces a new resource when changed.
* `key_type` - The key type for the certificate's private key. Can be one of:
  `P256` and `P384` (for ECDSA keys of respective length) or `2048`, `4096`, and
  `8192` (for RSA keys of respective length). Required when not specifying a
  CSR. The default is `2048` (RSA key of 2048 bits). Forces a new resource when
  changed.
* `certificate_request_pem` - A pre-created certificate request, such as one
  from [`tls_cert_request`][tls-cert-request], or one from an external source,
  in PEM format.  Either this, or the in-resource request options
  (`common_name`, `key_type`, and optionally `subject_alternative_names`) need
  to be specified. Forces a new resource when changed.
* `dns_challenge` (Optional) - The [DNS challenges](#using-dns-challenges) to
  use in fulfilling the request.
* `recursive_nameservers` (Optional) - The recursive nameservers that will be
  used to check for propagation of DNS challenge records, in addition to some
  in-provider checks such as zone detection. Defaults to your system-configured
  DNS resolvers.
* `disable_complete_propagation` (Optional) - Disable the requirement for full
  propagation of the TXT challenge records before proceeding with validation.
  Defaults to `false`.

-> See [About DNS propagation checks](#about-dns-propagation-checks) for details
on the `recursive_nameservers` and `disable_complete_propagation` settings.

* `pre_check_delay` (Optional) - Insert a delay after _every_ DNS challenge
  record to allow for extra time for DNS propagation before the certificate is
  requested. Use this option if you observe issues with requesting certificates
  even when DNS challenge records get added successfully. Units are in seconds.
  Defaults to 0 (no delay).

-> Be careful with `pre_check_delay` since the delay is executed _per-domain_.
Take your expected delay and divide it by the number of domains you have
configured (`common_name` + `subject_alternative_names`).

* `http_challenge` (Optional) - Defines an HTTP challenge to use in fulfilling
  the request.
* `http_webroot_challenge` (Optional) - Defines an alternate type of HTTP
  challenge that can be used to place a file at a location that can be served by
  an out-of-band webserver.
* `http_memcached_challenge` (Optional) - Defines an alternate type of HTTP
  challenge that can be used to serve up challenges to a
  [Memcached](https://memcached.org/) cluster.
* `http_s3_challenge` (Optional) - Defines an alternate type of HTTP
  challenge that can be used to serve up challenges to a
  [S3](https://aws.amazon.com/s3/) bucket.
* `tls_challenge` (Optional) - Defines a TLS challenge to use in fulfilling the
  request.

-> Only one of `http_challenge`, `http_webroot_challenge`, `http_s3_challenge`
and `http_memcached_challenge` can be defined at once. See the section on
[Using HTTP and TLS challenges](#using-http-and-tls-challenges) for more
details on using these and `tls_challenge`.

* `must_staple` (Optional) Enables the [OCSP Stapling Required][ocsp-stapling]
  TLS Security Policy extension. Certificates with this extension must include a
  valid OCSP Staple in the TLS handshake for the connection to succeed.
  Defaults to `false`. Note that this option has no effect when using an
  external CSR - it must be enabled in the CSR itself. Forces a new resource
  when changed.

[ocsp-stapling]: https://letsencrypt.org/docs/integration-guide/#implement-ocsp-stapling

-> OCSP stapling requires specific webserver configuration to support the
downloading of the staple from the CA's OCSP endpoints, and should be configured
to tolerate prolonged outages of the OCSP service. Consider this when using
`must_staple`, and only enable it if you are sure your webserver or service
provider can be configured correctly.

* `min_days_remaining` (Optional) - The minimum amount of days remaining on the
  expiration of a certificate before a renewal is attempted. The default is
  `30`. A value of less than `0` means that the certificate will never be
  renewed.
* `certificate_p12_password` - (Optional) Password to be used when generating
  the PFX file stored in [`certificate_p12`](#certificate_p12). Defaults to an
  empty string.
* `preferred_chain` - (Optional) The common name of the root of a preferred
  alternate certificate chain offered by the CA. The certificates in
  `issuer_pem` will reflect the chain requested, if available, otherwise the
  default chain will be provided. Forces a new resource when changed.

-> `preferred_chain` can be used to request alternate chains on Let's Encrypt
during the transition away from their old cross-signed intermediates. See [this
article for more
details](https://letsencrypt.org/2020/12/21/extending-android-compatibility.html).
In their example titled **"What about the alternate chain?"**, the root you
would put in to the `preferred_chain` field would be `ISRG Root X1`. The
equivalent in the [staging
environment](https://letsencrypt.org/docs/staging-environment/) is `(STAGING)
Pretend Pear X1`.

* `revoke_certificate_on_destroy` - Enables revocation of a certificate upon destroy,
which includes when a resource is re-created. Default is `true`.

* `cert_timeout` - Controls the timeout in seconds for certificate requests
  that are made after challenges are complete. Defaults to 30 seconds.

-> As mentioned, `cert_timeout` does nothing until all challenges are complete.
If you are looking to control timeouts related to a particular challenge (such
as a DNS challenge), see that challenge provider's specific options.

### Using DNS challenges

This method authenticates certificate domains by requiring the requester to
place a TXT record on the FQDNs in the certificate.

The ACME provider responds to DNS challenges automatically by utilizing one of
the supported DNS challenge providers. Most providers take credentials as
environment variables, but if you would rather use configuration for this
purpose, you can by specifying `config` blocks within a
[`dns_challenge`](#dns_challenge) block, along with the `provider` parameter.

See the DNS providers subcategory for a full list of DNS providers.

```hcl
resource "acme_certificate" "certificate" {
  #...

  dns_challenge {
    provider = "route53"

    config = {
      AWS_ACCESS_KEY_ID     = var.aws_access_key
      AWS_SECRET_ACCESS_KEY = var.aws_secret_key
      AWS_SESSION_TOKEN     = var.aws_security_token
      AWS_DEFAULT_REGION    = "us-east-1"  # OPTIONAL
    }
  }

  #...
}
```

#### Using Variable Files for Provider Arguments

Most provider arguments can be suffixed with `_FILE` to specify that you wish to
store that value in a local file. This can be useful if local storage for these
values is desired over configuration as variables or within the environment.

```hcl
resource "acme_certificate" "certificate" {
  #...

  dns_challenge {
    provider = "route53"

    config = {
      AWS_ACCESS_KEY_ID_FILE     = "/data/secrets/aws_access_key_id"
      AWS_SECRET_ACCESS_KEY_FILE = "/data/secrets/aws_secret_access_key"
      AWS_DEFAULT_REGION         = "us-east-1"
    }
  }

  #...
}
```

#### About DNS propagation checks

There are two parts to the DNS propagation check:

* A check using your system resolvers, or the settings specified in
  `recursive_nameservers`.
* A check against your domain's authoritative DNS servers.

-> The authoritative part of the DNS propagation check will almost always
require access to the outside internet. Make sure you allow the required access
accordingly, particularly in restricted networks. You can also use the
`disable_complete_propagation` setting to bypass this check altogether (see
below).

The ACME provider will normally use your system-configured DNS resolvers to
check for propagation of the TXT records before proceeding with the certificate
request. In split horizon scenarios, this check may never succeed, as the
machine running Terraform may not have visibility into these public DNS
records.

To override this default behavior, supply the `recursive_nameservers` to use as
a list in the format `host:port`:

```hcl
resource "acme_certificate" "certificate" {
  #...

  recursive_nameservers = ["8.8.8.8:53"]

  dns_challenge {
    provider = "route53"
  }

  #...
}
```

Additionally, in air-gapped scenarios, internet access to DNS servers may not be
available at all to the machine running Terraform. In this case, you can use
`disable_complete_propagation` to bypass this authoritative DNS check, ensuring
that the only propagation check being done is on the system resolver or the
resolver you configure with `recursive_nameservers`.

```hcl
resource "acme_certificate" "certificate" {
  #...

  recursive_nameservers        = ["8.8.8.8:53"]
  disable_complete_propagation = true

  dns_challenge {
    provider = "route53"
  }

  #...
}
```

~> **NOTE:** When `disable_complete_propagation` is used, you can encounter
situations where the propagation check will pass before your platform has
provisioned the DNS records on their name servers. Use this setting with care,
such as in the aforementioned air-gapped scenario where the system running
Terraform has no outbound DNS access, or for testing purposes. If you encounter
problems using this setting, consider removing it and moving your Terraform
operations to a system that can access your domain's authoritative DNS servers.

#### Using multiple primary DNS providers

The ACME provider will allow you to configure multiple DNS challenges in the
event that you have more than one primary DNS provider.

```hcl
resource "acme_certificate" "certificate" {
  #...

  dns_challenge {
    provider = "azure"
  }

  dns_challenge {
    provider = "gcloud"
  }

  dns_challenge {
    provider = "route53"
  }

  #...
}
```

Some considerations need to be kept in mind when using multiple providers:

* You cannot use more than one provider of the same type at once.
* Your NS records must be correctly configured so that each DNS challenge
  provider can correctly discover the appropriate zone to update.
* DNS propagation checks are conducted once per configured common name and
  subject alternative name, using the highest configured or default propagation
  timeout (`*_PROPAGATION_TIMEOUT`) and polling interval (`*_POLLING_INTERVAL`)
  settings.

#### Relation to Terraform provider configuration

The DNS provider configuration specified in the `acme_certificate` resource is
separate from any that you supply in a corresponding provider whose
functionality overlaps with the certificate's DNS providers.  This ensures that
there are no hard dependencies between any of these providers and the ACME
provider, but it is important to note so that configuration is supplied
correctly.

As an example, if you specify manual configuration for the [AWS
provider][tf-provider-aws] via the [`provider`][tf-providers] block instead of
the environment, you will still need to supply the configuration explicitly as
per above.

[tf-provider-aws]: https://registry.terraform.io/providers/hashicorp/aws/latest
[tf-providers]: https://www.terraform.io/docs/configuration/providers.html

Some of these providers have environment variable settings that overlap with
the ones found here, generally depending on whether or not these variables are
supported by the corresponding provider's SDK.

Check the documentation of a specific DNS provider for more details on exactly
what variables are supported.

### Using HTTP and TLS Challenges

-> It's recommended that you use [DNS challenges](#using-dns-challenges)
whenever possible to generate certificates with `acme_certificate`. Only use the
HTTP and TLS challenge types if you don't have access to do DNS challenges, and
can ensure that you can direct traffic for all domains being authorized to the
machine running Terraform, or the locations served by the
`http_webroot_challenge`, `http_s3_challenge` and `http_memcached_challenge` types. 
Additionally, these challenge types do not support wildcard domains. See the
[Let's Encrypt page on challenge types](https://letsencrypt.org/docs/challenge-types/)
for more details. These challenges have requirements that almost always exclude them from
being used on [Terraform Cloud](https://www.terraform.io/docs/cloud/) unless you
are using the [Cloud
Agents](https://www.terraform.io/docs/cloud/agents/index.html) feature.

`acme_certificate` supports HTTP and TLS challenges. The provider accomplishes
this by running a small HTTP or TLS service to serve records using the HTTP-01
or TLS-ALPN-01 challenge types. Additionally, two alternate HTTP challenge
providers exist that allow HTTP challenges to be satisfied by publishing the
challenge records either to an arbitrary filesystem location or a
[Memcached](https://memcached.org/) cluster.

#### Network Requirements for Using `http_challenge` and `tls_challenge`

`http_challenge` and `tls_challenge` by default will listen on their respective
ports (port 80 for HTTP and port 443 for TLS). These ports are _privileged_ and
will likely not be accessible by the machine running Terraform.

You can work around this by doing the following:

* On Linux, use [`setcap`](https://man7.org/linux/man-pages/man8/setcap.8.html)
  to grant escalated network privileges to either Terraform (`setcap
  'cap_net_bind_service=+eip' "$(which terraform)"`), or the provider (`setcap
  'cap_net_bind_service=+ep'
  .terraform/providers/registry.terraform.io/vancluever/acme/VERSION/ARCH/terraform-provider-acme_vVERSION`).
  Both have drawbacks: granting capabilites to Terraform itself will mean that
  Terraform core and any provider launched by it will also have the capability,
  while capabilities granted to the provider will be lost every time the
  provider is updated, or the repository is initialized with `terraform init`.
* Use proxies to direct traffic to the ports defined with the `port` option in
  the challenge clauses. If necessary, use the `proxy_header` option of
  `http_challenge` to set the header to match the host of the current FQDN being
  solved.

~> Never run Terraform (or the plugin) as root! If you cannot satisfy the
networking requirements for `http_challenge` or `tls_challenge`, consider using
the other challenge types or use [DNS challenges](#using-dns-challenges).

#### `http_challenge`

The `http_challenge` type supports standard HTTP-01 challenges.

```
resource "acme_certificate" "certificate" {
  #...

  http_challenge {
    port         = "5002"
    proxy_header = "Forwarded"
  }

  #...
}
```

The options are as follows:

* `port` (Optional) - The port that the challenge server listens on. Default: `80`.
* `proxy_header` (Optional) - The proxy header to match against. Default:
  `Host`.

The `proxy_header` option behaves differently depending on its definition:

* When set to `Host`, standard host header validation is used.
* When set to `Forwarded`, the server looks in the `Forwarded` header for a
  section matching `host=DOMAIN` where `DOMAIN` is the domain currently being
  resolved by the challenge. See [RFC 7239](https://tools.ietf.org/html/rfc7239)
  for more details.
* When set to an arbitrary header (example: `X-Forwarded-Host`), that header is
  checked for the host entry in the same way the host header would normally be
  checked.

#### `http_webroot_challenge`

Use `http_webroot_challenge` to publish a record to a location on the file
system. The record is published to `DIRECTORY/.well-known/acme-challenge/`. The
resource will request an HTTP-01 challenge for which an out-of-band process must
use this data to answer.

```
resource "acme_certificate" "certificate" {
  #...

  http_webroot_challenge {
    directory = "/a/webserver/path"
  }

  #...
}
```

The options are as follows:

* `directory` (Required) - The directory to publish the record to.

#### `http_memcached_challenge`

Use `http_memcached_challenge` to publish challenge records to a
[Memcached](https://memcached.org/) cluster. The record is published to
`/.well-known/acme-challenge/KEY`. The resource will request an HTTP-01
challenge for which an out-of-band process must use this data to answer.

See the [README.md on
lego](https://github.com/go-acme/lego/blob/4bb8bea031eb805f361c04ca222f266b9e7feced/providers/http/memcached/README.md)
for an example using Nginx.

```
resource "acme_certificate" "certificate" {
  #...

  http_memcached_challenge {
    hosts = ["127.0.0.1:11211"]
  }

  #...
}
```

#### `http_s3_challenge`

Use `http_s3_challenge` to publish challenge records to a
[S3](https://aws.amazon.com/s3/) bucket. The record is published to
`/.well-known/acme-challenge/KEY` in the bucket. The domain will need to be configured
to point to the s3 bucket, either with a reverse proxy or some application.
 The resource will request an HTTP-01
challenge for which an out-of-band process must use this data to answer.

See the [Documentation on
lego](https://github.com/go-acme/lego/blob/master/providers/http/s3/s3.toml)


```
resource "acme_certificate" "certificate" {
  #...

  http_s3_challenge {
    s3_bucket = "bucket_name"
  }

  #...
}
```

The options are as follows:

* `hosts` (Required) - The hosts to publish the record to.

#### `tls_challenge`

The `tls_challenge` type supports TLS-ALPN-01 challenges.

```
resource "acme_certificate" "certificate" {
  #...

  tls_challenge {
    port = "5001"
  }

  #...
}
```

The options are as follows:

* `port` (Optional) - The port that the challenge server listens on. Default: `443`.

## Certificate renewal

The `acme_certificate` resource handles automatic certificate renewal so long
as a plan or apply is done within the number of days specified in the
[`min_days_remaining`](#min_days_remaining) resource parameter. During refresh,
if Terraform detects that the certificate is within the expiry range specified
in `min_days_remaining`, or is already expired, Terraform will mark the
certificate to be renewed on the next apply.

Note that a value less than `0` supplied to `min_days_remaining` will cause
renewal checks to be bypassed, and the certificate will never renew.

## Attribute Reference

The following attributes are exported:

* `id` - A UUID identifying the resource in Terraform state.

-> As of provider version 2.0, this is no longer the same as `certificate_url`.
Refer to that field for the current URL of the certificate.

* `certificate_url` - The full URL of the certificate within the ACME CA.
* `certificate_domain` - The common name of the certificate.
* `private_key_pem` - The certificate's private key, in PEM format, if the
  certificate was generated from scratch and not with
  [`certificate_request_pem`](#certificate_request_pem).  If
  `certificate_request_pem` was used, this will be blank.
* `certificate_pem` - The certificate in PEM format. This does not include the
  `issuer_pem`. This certificate can be concatenated with `issuer_pem` to form
  a full chain, e.g. `"${acme_certificate.certificate.certificate_pem}${acme_certificate.certificate.issuer_pem}"`
* `issuer_pem` - The intermediate certificates of the issuer. Multiple
  certificates are concatenated in this field when there is more than one
  intermediate certificate in the chain.
* `certificate_p12` - The certificate, any intermediates, and the private key
  archived as a PFX file (PKCS12 format, generally used by Microsoft products).
  The data is base64 encoded (including padding), and its password is
  configurable via the [`certificate_p12_password`](#certificate_p12_password)
  argument. This field is empty if creating a certificate from a CSR.
* `certificate_not_after` - The expiry date of the certificate, laid out in
  RFC3339 format (`2006-01-02T15:04:05Z07:00`).
