---
layout: "acme"
page_title: "ACME: Stackpath DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-stackpath"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Stackpath DNS Challenge Provider

The `stackpath` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Stackpath][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.stackpath.com

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "stackpath"
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

* `STACKPATH_CLIENT_ID` - The client ID to use.
* `STACKPATH_CLIENT_SECRET` - The client secret to use.
* `STACKPATH_STACK_ID` - The stack ID to use.

The following additional optional variables are available:

* `STACKPATH_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `STACKPATH_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `STACKPATH_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
