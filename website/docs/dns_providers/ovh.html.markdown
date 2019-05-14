---
layout: "acme"
page_title: "ACME: OVH DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-ovh"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# OVH DNS Challenge Provider

The `ovh` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[OVH][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.ovh.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "ovh"
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

* `OVH_ENDPOINT` - The API endpoint to use. Can be one of `ovh-eu` or `ovh-ca`.
* `OVH_APPLICATION_KEY ` - The application key to use.
* `OVH_APPLICATION_SECRET` - The application secret to use.
* `OVH_CONSUMER_KEY` - The consumer key to use.

The following additional optional variables are available:

* `OVH_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `OVH_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `OVH_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
* `OVH_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `180`).
