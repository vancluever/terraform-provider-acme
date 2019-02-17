---
layout: "acme"
page_title: "ACME: DigitalOcean DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-digitalocean"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# DigitalOcean DNS Challenge Provider

The `digitalocean` DNS challenge provider can be used to perform DNS challenges
for the [`acme_certificate`][resource-acme-certificate] resource with
[DigitalOcean][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.digitalocean.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "digitalocean"
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

* `DO_AUTH_TOKEN` - The auth token to use.

The following additional optional variables are available:

* `DO_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `5`).
* `DO_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `DO_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `30`).
* `DO_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
