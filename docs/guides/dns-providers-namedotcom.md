---
page_title: "namedotcom"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Name.com DNS Challenge Provider

The `namedotcom` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Name.com](https://www.name.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "namedotcom"
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

* `NAMECOM_API_TOKEN` - API token.
* `NAMECOM_USERNAME` - Username.

* `NAMECOM_HTTP_TIMEOUT` - API request timeout in seconds (Default: 10).
* `NAMECOM_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 20).
* `NAMECOM_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 900).
* `NAMECOM_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 300).


