---
layout: "acme"
page_title: "ACME: Selectel DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-selectel"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Selectel DNS Challenge Provider

The `selectel` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Selectel][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://selectel.ru/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "selectel"
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

* `SELECTEL_API_TOKEN` - The API token to use.

The following additional optional variables are available:

* `SELECTEL_BASE_URL` - The base URL to use (default:
  `https://api.selectel.ru/domains/v1`).
* `SELECTEL_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `2`).
* `SELECTEL_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for
  DNS propagation (default: `120`).
* `SELECTEL_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `60`).
* `SELECTEL_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
