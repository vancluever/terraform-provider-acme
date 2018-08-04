---
layout: "acme"
page_title: "ACME: DNS Made Easy DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-dnsmadeeasy"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# DNS Made Easy DNS Challenge Provider

The `dnsmadeeasy` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[DNS Made Easy][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://dnsmadeeasy.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "dnsmadeeasy"
  }
}
```

## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

* `DNSMADEEASY_API_KEY` - The API key to use.
* `DNSMADEEASY_API_SECRET` - The secret key to use.
