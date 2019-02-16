---
layout: "acme"
page_title: "ACME: TransIP DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-transip"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# TransIP DNS Challenge Provider

The `transip` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[TransIP][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.transip.nl/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "transip"
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

* `TRANSIP_ACCOUNT_NAME` - The account name to use.
* `TRANSIP_PRIVATE_KEY_PATH` - The path to the account's private key.

The following additional optional variables are available:

* `TRANSIP_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `10`).
* `TRANSIP_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for
  DNS propagation (default: `600`).
* `TRANSIP_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `10`).
