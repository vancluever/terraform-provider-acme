---
layout: "acme"
page_title: "ACME: Gandi DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-gandi"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Gandi DNS Challenge Provider

The `gandi` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Gandi][provider-service-page].

-> **NOTE:** This provider is for the Gandi V4 API. For the V5 API and higher
(aka LiveDNS), use the [`gandiv5`][gandiv5] provider.

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.gandi.net/en
[gandiv5]: /docs/providers/acme/dns_providers/gandiv5.html

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "gandi"
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

* `GANDI_API_KEY` - The API key to use.
