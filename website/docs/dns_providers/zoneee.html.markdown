---
layout: "acme"
page_title: "ACME: Zone DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-zoneee"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Zone DNS Challenge Provider

The `zoneee` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Zone][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.zone.ee/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "zoneee"
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

* `ZONEEE_API_USER` - The API user to use.
* `ZONEEE_API_KEY` - The API key to use.

The following additional optional variables are available:

* `ZONEEE_ENDPOINT` - The API endpoint to use (default:
  `https://api.zone.eu/v2/dns/`).
* `ZONEEE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `5`).
* `ZONEEE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `300`).
* `ZONEEE_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
