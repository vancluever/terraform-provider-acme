---
layout: "acme"
page_title: "ACME: Alibaba Cloud DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-alidns"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Alibaba Cloud DNS Challenge Provider

The `alidns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Alibaba Cloud][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.alibabacloud.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "alidns"
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

* `ALICLOUD_ACCESS_KEY` - The API key to use.
* `ALICLOUD_SECRET_KEY` - The secret key to use.
