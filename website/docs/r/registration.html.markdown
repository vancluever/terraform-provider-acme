---
layout: "acme"
page_title: "ACME: acme_registration"
sidebar_current: "docs-acme-resource-registration"
description: |-
  Provides a resource to manage accounts on an ACME CA.
---

# acme_registration

The `acme_registration` resource can be used to create and manage accounts on an
ACME server. Once registered, the same private key that has been used for
registration can be used to request authorizations for certificates.

-> This resource is named `acme_registration` for historical reasons - in the
ACME v1 spec, a _registration_ referred to the account entity.  This resource
name is stable and more than likely will not change until a later major version
of the provider, if at all.

## Example

The following creates an account off of a private key generated with the
[`tls_private_key`][resource-tls-private-key] resource.

[resource-tls-private-key]: /docs/providers/tls/r/private_key.html

```hcl
provider "acme" {
  server_url = "https://acme-staging-v02.api.letsencrypt.org/directory"
}

resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = "${tls_private_key.private_key.private_key_pem}"
  email_address   = "nobody@example.com"
}
```

#### Argument Reference

~> **NOTE:** All arguments in `acme_registration` force a new resource if
changed.

The resource takes the following arguments:

* `account_key_pem` (Required) - The private key used to identity the account.
* `email_address` (Required) - The contact email address for the account.

#### Attribute Reference

The following attributes are exported:

* `id`: The original full URL of the account.
* `registration_url`: The current full URL of the account.

-> `id` and `registration_url` will usually be the same and will usually only
diverge when migrating protocols, ie: ACME v1 to v2.
