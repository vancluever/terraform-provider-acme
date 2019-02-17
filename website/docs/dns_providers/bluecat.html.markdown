---
layout: "acme"
page_title: "ACME: Bluecat DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-bluecat"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Bluecat DNS Challenge Provider

The `bluecat` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with [Bluecat
Address Manager][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.bluecatnetworks.com/platform/management/bluecat-address-manager/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "bluecat"
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

* `BLUECAT_SERVER_URL` - The URL for the address manager to use.
* `BLUECAT_USER_NAME` - The user name to use.
* `BLUECAT_PASSWORD` - The password to use for the supplied user name.
* `BLUECAT_CONFIG_NAME` - The configuration name to use.
* `BLUECAT_DNS_VIEW` - The DNS view to use.

The following additional optional variables are available:

* `BLUECAT_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `BLUECAT_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `BLUECAT_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
* `BLUECAT_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
