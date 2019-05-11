---
layout: "acme"
page_title: "ACME: ConoHa DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-conoha"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# ConoHa DNS Challenge Provider

The `conoha` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with [ConoHa
DNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.conoha.jp/conoha/features/dns/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "conoha"
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

* `CONOHA_TENANT_ID` - The tenant ID to use.
* `CONOHA_API_USERNAME` - The API username to use.
* `CONOHA_API_PASSWORD` - The API password to use.

The following additional optional variables are available:

* `CONOHA_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `CONOHA_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `CONOHA_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `60`).
* `CONOHA_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
* `CONOHA_REGION` - The region to connect to (default: `tyo1`).
