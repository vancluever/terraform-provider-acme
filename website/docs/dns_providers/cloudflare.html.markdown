---
layout: "acme"
page_title: "ACME: Cloudflare DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-cloudflare"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Cloudflare DNS Challenge Provider

The `cloudflare` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Cloudflare DNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.cloudflare.com/dns/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "cloudflare"
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

* `CLOUDFLARE_EMAIL` - The email address to use.
* `CLOUDFLARE_API_KEY` - The API key to use.
