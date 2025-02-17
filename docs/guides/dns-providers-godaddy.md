---
page_title: "godaddy"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Go Daddy DNS Challenge Provider

The `godaddy` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Go Daddy](https://godaddy.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "godaddy"
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

* `GODADDY_API_KEY` - API key.
* `GODADDY_API_SECRET` - API secret.

* `GODADDY_HTTP_TIMEOUT` - API request timeout in seconds (Default: 30).
* `GODADDY_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `GODADDY_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 120).
* `GODADDY_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 600).

GoDaddy has recently (2024-04) updated the account requirements to access parts of their production Domains API:

- Availability API: Limited to accounts with 50 or more domains.
- Management and DNS APIs: Limited to accounts with 10 or more domains and/or an active Discount Domain Club plan.

https://community.letsencrypt.org/t/getting-unauthorized-url-error-while-trying-to-get-cert-for-subdomains/217329/12

