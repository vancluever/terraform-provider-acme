---
layout: "acme"
page_title: "ACME: DuckDNS DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-duckdns"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# DuckDNS DNS Challenge Provider

The `duckdns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[DuckDNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: http://www.duckdns.org/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "duckdns"
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

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

* `DUCKDNS_TOKEN` - The auth token to use.

The following additional optional variables are available:

* `DUCKDNS_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `DUCKDNS_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `DUCKDNS_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
* `DUCKDNS_SEQUENCE_INTERVAL` - The time between each DNS challenge (default:
  `60`).
