---
layout: "acme"
page_title: "ACME: Vscale DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-vscale"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Vscale DNS Challenge Provider

The `vscale` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Vscale][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://vscale.io/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "vscale"
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

* `VSCALE_API_TOKEN` - The API token to use.

The following additional optional variables are available:

* `VSCALE_BASE_URL` - The base URL to use (default:
  `https://api.vscale.io/v1/domains`).
* `VSCALE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `2`).
* `VSCALE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `120`).
* `VSCALE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `60`).
* `VSCALE_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
