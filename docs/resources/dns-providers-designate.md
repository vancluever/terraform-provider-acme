---
page_title: "designate"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Designate DNSaaS for Openstack DNS Challenge Provider

The `designate` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Designate DNSaaS for Openstack](https://docs.openstack.org/designate/latest/).

[resource-acme-certificate]: ./certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ./certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "designate"
  }
}
```
## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: ./certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ./certificate.md#using-variable-files-for-provider-arguments

* `OS_AUTH_URL` - Identity endpoint URL.
* `OS_PASSWORD` - Password.
* `OS_PROJECT_NAME` - Project name.
* `OS_REGION_NAME` - Region name.
* `OS_USERNAME` - Username.

* `DESIGNATE_POLLING_INTERVAL` - Time between DNS propagation check.
* `DESIGNATE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DESIGNATE_TTL` - The TTL of the TXT record used for the DNS challenge.
* `OS_PROJECT_ID` - Project ID.
* `OS_TENANT_NAME` - Tenant name (deprecated see OS_PROJECT_NAME and OS_PROJECT_ID).


