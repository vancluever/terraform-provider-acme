---
layout: "acme"
page_title: "ACME: hosting.de DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-hostingde"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# hosting.de DNS Challenge Provider

The `hostingde` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[hosting.de][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.hosting.de

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "hostingde"
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

* `HOSTINGDE_API_KEY` - The API key to use.
* `HOSTINGDE_ZONE_NAME` - The zone name to use.

The following additional optional variables are available:

* `HOSTINGDE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `2`).
* `HOSTINGDE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `120`).
* `HOSTINGDE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
* `HOSTINGDE_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
