---
layout: "acme"
page_title: "ACME: INWX DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-inwx"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# INWX DNS Challenge Provider

The `inwx` DNS challenge provider can be used to perform DNS challenges for the
[`acme_certificate`][resource-acme-certificate] resource with
[INWX][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.inwx.com/en/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "inwx"
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

* `INWX_USERNAME` - The API username to use.
* `INWX_PASSWORD` - The API password to use.

The following additional optional variables are available:

* `INWX_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `INWX_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `INWX_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `300`).
* `INWX_SANDBOX` - Whether or not to use sandbox mode (default: `false`).
