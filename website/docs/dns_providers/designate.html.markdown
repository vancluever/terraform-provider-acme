---
layout: "acme"
page_title: "ACME: OpenStack Designate DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-designate"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# ConoHa DNS Challenge Provider

The `designate` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with [OpenStack
Designate][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://docs.openstack.org/designate/latest/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "designate"
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

* `OS_AUTH_URL` - The Identity authentication URL.
* `OS_USERNAME` - The Username to login with.
* `OS_PASSWORD` - The Password to login with.
* `OS_TENANT_NAME` - The Name of the Tenant (Identity v2) or Project (Identity v3)
  to login with.
* `OS_REGION_NAME` - The region of the OpenStack cloud to use.

The following additional optional variables are available:

* `DESIGNATE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `10`).
* `DESIGNATE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `10`).
* `DESIGNATE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `10`).
