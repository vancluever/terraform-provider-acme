---
layout: "acme"
page_title: "ACME: ACME-DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-acmedns"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# ACME-DNS DNS Challenge Provider

The `acme-dns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[ACME-DNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://github.com/joohoi/acme-dns

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "acme-dns"
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

* `ACME_DNS_API_BASE` - The ACME-DNS API addres to use.
* `ACME_DNS_STORAGE_PATH` - The ACME-DNS JSON account data file to use.
