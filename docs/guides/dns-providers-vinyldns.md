---
page_title: "vinyldns"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# VinylDNS DNS Challenge Provider

The `vinyldns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[VinylDNS](https://www.vinyldns.io).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "vinyldns"
  }
}
```
## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: ../resources/certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ../resources/certificate.md#using-variable-files-for-provider-arguments

* `VINYLDNS_ACCESS_KEY` - The VinylDNS API key.
* `VINYLDNS_HOST` - The VinylDNS API URL.
* `VINYLDNS_SECRET_KEY` - The VinylDNS API Secret key.

* `VINYLDNS_POLLING_INTERVAL` - Time between DNS propagation check.
* `VINYLDNS_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `VINYLDNS_TTL` - The TTL of the TXT record used for the DNS challenge.

The vinyldns integration makes use of dotted hostnames to ease permission management.
Users are required to have DELETE ACL level or zone admin permissions on the VinylDNS zone containing the target host.

