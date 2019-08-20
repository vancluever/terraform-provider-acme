---
layout: "acme"
page_title: "ACME: Gandi Live DNS (v5) DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-gandiv5"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Gandi Live DNS (v5) DNS Challenge Provider

The `gandiv5` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Gandi Live DNS (v5)][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.gandi.net

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "gandiv5"
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

* `GANDIV5_API_KEY` - API key.

The following additional optional variables are available:

* `GANDIV5_HTTP_TIMEOUT` - API request timeout.
* `GANDIV5_POLLING_INTERVAL` - Time between DNS propagation check.
* `GANDIV5_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `GANDIV5_TTL` - The TTL of the TXT record used for the DNS challenge.


