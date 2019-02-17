---
layout: "acme"
page_title: "ACME: Linode v4 DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-v4-linode"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Linode v4 DNS Challenge Provider

The `linodev4` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Linode][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.linode.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "linodev4"
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

* `LINODE_TOKEN` - The token to use.

The following additional optional variables are available:

* `LINODE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `15`).
* `LINODE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `300`).
* `LINODE_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  no timeout).
