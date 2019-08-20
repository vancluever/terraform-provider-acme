---
layout: "acme"
page_title: "ACME: Rackspace DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-rackspace"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Rackspace DNS Challenge Provider

The `rackspace` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Rackspace][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.rackspace.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "rackspace"
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

* `RACKSPACE_API_KEY` - API key.
* `RACKSPACE_USER` - API user.

The following additional optional variables are available:

* `RACKSPACE_HTTP_TIMEOUT` - API request timeout.
* `RACKSPACE_POLLING_INTERVAL` - Time between DNS propagation check.
* `RACKSPACE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `RACKSPACE_TTL` - The TTL of the TXT record used for the DNS challenge.


