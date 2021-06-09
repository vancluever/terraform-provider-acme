---
page_title: "duckdns"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Duck DNS DNS Challenge Provider

The `duckdns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Duck DNS](https://www.duckdns.org/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "duckdns"
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

* `DUCKDNS_TOKEN` - Account token.

* `DUCKDNS_HTTP_TIMEOUT` - API request timeout.
* `DUCKDNS_POLLING_INTERVAL` - Time between DNS propagation check.
* `DUCKDNS_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DUCKDNS_SEQUENCE_INTERVAL` - Time between sequential requests.
* `DUCKDNS_TTL` - The TTL of the TXT record used for the DNS challenge.


