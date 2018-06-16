## [ACME Provider](README.md)

* [`acme_registration`](resource_acme_registration.md)
* [`acme_certificate`](resource_acme_certificate.md)

# `acme_registration`

The `acme_registration` resource can be used to create and manage accounts on an
ACME server. Once registered, the same private key that has been used for
registration can be used to request authorizations for certificates.

:warning: **NOTE:** This resource is named `acme_registration` for historical
reasons - in the ACME v1 spec, a _registration_ referred to the account entity.
This resource name is stable and more than likely will not change until a later
major version of the provider, if at all.

## Example

The following creates an account off of a private key generated with the
[`tls_private_key`][resource-tls-private-key] resource.

[resource-tls-private-key]: https://www.terraform.io/docs/providers/tls/r/private_key.html

```hcl
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  server_url      = "https://acme-staging-v02.api.letsencrypt.org/directory"
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}
```

#### Argument Reference

:warning: **NOTE:** All arguments in `acme_registration` force a new resource if
changed.

The resource takes the following arguments:

* `server_url` (Required) - The URL of the ACME directory endpoint.
* `account_key_pem` (Required) - The private key used to identity the account.
* `email_address` (Required) - The contact email address for the account.

#### Attribute Reference

The only attribute that is exported from this resource is the `id` of the
resource, which is set to the full URL of the account.
