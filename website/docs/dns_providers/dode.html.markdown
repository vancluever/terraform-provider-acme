---
layout: "acme"
page_title: "ACME: Domain Offensive DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-dode"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Domain Offensive DNS Challenge Provider

The `dode` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Domain Offensive][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.do.de/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "dode"
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

* `DODE_TOKEN` - The API token to use.

The following additional optional variables are available:

* `DODE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `2`).
* `DODE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `DODE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
* `DODE_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
* `DODE_SEQUENCE_INTERVAL` - The time between each DNS challenge (default:
  `60`).
