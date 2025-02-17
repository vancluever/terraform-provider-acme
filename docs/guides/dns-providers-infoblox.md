---
page_title: "infoblox"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Infoblox DNS Challenge Provider

The `infoblox` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Infoblox](https://www.infoblox.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "infoblox"
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

* `INFOBLOX_HOST` - Host URI.
* `INFOBLOX_PASSWORD` - Account Password.
* `INFOBLOX_USERNAME` - Account Username.

* `INFOBLOX_DNS_VIEW` - The view for the TXT records (Default: External).
* `INFOBLOX_HTTP_TIMEOUT` - API request timeout in seconds (Default: 30).
* `INFOBLOX_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `INFOBLOX_PORT` - The port for the infoblox grid manager  (Default: 443).
* `INFOBLOX_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 60).
* `INFOBLOX_SSL_VERIFY` - Whether or not to verify the TLS certificate  (Default: true).
* `INFOBLOX_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 120).
* `INFOBLOX_WAPI_VERSION` - The version of WAPI being used  (Default: 2.11).

When creating an API's user ensure it has the proper permissions for the view you are working with.

