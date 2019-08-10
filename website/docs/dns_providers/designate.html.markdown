---
layout: "acme"
page_title: "ACME: Designate DNSaaS for Openstack DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-designate"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Designate DNSaaS for Openstack DNS Challenge Provider

The `designate` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Designate DNSaaS for Openstack][provider-service-page].

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

* `OS_AUTH_URL` - Identity endpoint URL.
* `OS_PASSWORD` - Password.
* `OS_PROJECT_NAME` - Project name.
* `OS_REGION_NAME` - Region name.
* `OS_TENANT_NAME` - Tenant name (deprecated see OS_PROJECT_NAME and OS_PROJECT_ID).
* `OS_USERNAME` - Username.

The following additional optional variables are available:

* `DESIGNATE_POLLING_INTERVAL` - Time between DNS propagation check.
* `DESIGNATE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DESIGNATE_TTL` - The TTL of the TXT record used for the DNS challenge.
* `OS_PROJECT_ID` - Project ID.


