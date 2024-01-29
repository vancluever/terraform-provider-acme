---
page_title: "efficientip"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Efficient IP DNS Challenge Provider

The `efficientip` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Efficient IP](https://efficientip.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "efficientip"
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

* `EFFICIENTIP_DNS_NAME` - DNS name (ex: dns.smart).
* `EFFICIENTIP_HOSTNAME` - Hostname (ex: foo.example.com).
* `EFFICIENTIP_PASSWORD` - Password.
* `EFFICIENTIP_USERNAME` - Username.

* `EFFICIENTIP_HTTP_TIMEOUT` - API request timeout.
* `EFFICIENTIP_INSECURE_SKIP_VERIFY` - Whether or not to verify EfficientIP API certificate.
* `EFFICIENTIP_POLLING_INTERVAL` - Time between DNS propagation check.
* `EFFICIENTIP_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `EFFICIENTIP_TTL` - The TTL of the TXT record used for the DNS challenge.
* `EFFICIENTIP_VIEW_NAME` - View name (ex: external).


