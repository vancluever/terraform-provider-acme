---
layout: "acme"
page_title: "ACME: Open Telekom Cloud DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-otc"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Open Telekom Cloud DNS Challenge Provider

The `otc` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Open Telekom Cloud][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://cloud.telekom.de/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "otc"
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

* `OTC_USER_NAME` - The user name to use.
* `OTC_DOMAIN_NAME` - The domain name to use.
* `OTC_PASSWORD ` - The password for the supplied user.
* `OTC_PROJECT_NAME` - The project name.
* `OTC_IDENTITY_ENDPOINT` - The identity endpoint to use.

The following additional optional variables are available:

* `OTC_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `OTC_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `OTC_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `300`).
* `OTC_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `10`).
