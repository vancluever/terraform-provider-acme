---
layout: "acme"
page_title: "ACME: Netcup DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-netcup"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Netcup DNS Challenge Provider

The `netcup` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Netcup][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.netcup.eu

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "netcup"
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

* `NETCUP_CUSTOMER_NUMBER` - The customer number to use.
* `NETCUP_API_KEY` - The API key to use.
* `NETCUP_API_PASSWORD` - The API password to use.

The following additional optional variables are available:

* `NETCUP_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `5`).
* `NETCUP_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `120`).
* `NETCUP_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
* `NETCUP_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `10`).
