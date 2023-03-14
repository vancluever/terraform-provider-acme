# acme_registration

The `acme_registration` resource can be used to create and manage accounts on an
ACME server. Once registered, the same private key that has been used for
registration can be used to request authorizations for certificates.

-> This resource is named `acme_registration` for historical reasons - in the
ACME v1 spec, a _registration_ referred to the account entity.  This resource
name is stable and more than likely will not change until a later major version
of the provider, if at all.

-> Keep in mind that when using this resource along with
[`acme_certificate`][resource-certificate] within the same configuration, a
change in the provider-level `server_url` (example: from the Let's Encrypt
staging to production environment) within the same Terraform state will result
in a resource failure, as Terraform will attempt to look for the account in the
wrong CA. Consider different workspaces per environment, and/or using [multiple
provider instances][multiple-provider-instances].

[multiple-provider-instances]: https://www.terraform.io/docs/configuration/providers.html#alias-multiple-provider-configurations
[resource-certificate]: ./certificate.md

## Example

The following creates an account off of a private key generated with the
[`tls_private_key`][resource-tls-private-key] resource.

[resource-tls-private-key]: https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key

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
```

#### Argument Reference

~> **NOTE:** All arguments in `acme_registration` force a new resource if
changed.

The resource takes the following arguments:

* `account_key_pem` (Required) - The private key used to identify the account.
* `email_address` (Required) - The contact email address for the account.
* `external_account_binding` (Optional) - An external account binding for the
  registration, usually used to link the registration with an account in a
  commercial CA. Sub-options are:
    - `key_id` (Required): The key ID for the external account binding.
    - `hmac_base64` (Required): The base64-encoded message authentication code
      for the external account binding.

#### Attribute Reference

The following attributes are exported:

* `id`: The original full URL of the account.
* `registration_url`: The current full URL of the account.

-> `id` and `registration_url` will usually be the same and will usually only
diverge when migrating protocols, ie: ACME v1 to v2.
