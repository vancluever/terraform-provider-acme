---
page_title: "ovh"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# OVH DNS Challenge Provider

The `ovh` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[OVH](https://www.ovh.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

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

[resource-acme-certificate-dns-challenge-arg]: ../resources/certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ../resources/certificate.md#using-variable-files-for-provider-arguments

* `OVH_APPLICATION_KEY` - Application key.
* `OVH_APPLICATION_SECRET` - Application secret.
* `OVH_CONSUMER_KEY` - Consumer key.
* `OVH_ENDPOINT` - Endpoint URL (ovh-eu or ovh-ca).

* `OVH_HTTP_TIMEOUT` - API request timeout.
* `OVH_POLLING_INTERVAL` - Time between DNS propagation check.
* `OVH_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `OVH_TTL` - The TTL of the TXT record used for the DNS challenge.

## Application Key and Secret

Application key and secret can be created by following the [OVH guide](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/).

When requesting the consumer key, the following configuration can be used to define access rights:

```json
{
  "accessRules": [
    {
      "method": "POST",
      "path": "/domain/zone/*"
    },
    {
      "method": "DELETE",
      "path": "/domain/zone/*"
    }
  ]
}
```

