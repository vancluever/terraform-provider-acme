---
layout: "acme"
page_title: "ACME: AuroraDNS DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-auroradns"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# AuroraDNS DNS Challenge Provider

The `auroradns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[AuroraDNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://auroradns.microsoft.com/en-ca/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "auroradns"
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

* `AURORA_USER_ID` - The user ID to use.
* `AURORA_KEY` - The key to use.
