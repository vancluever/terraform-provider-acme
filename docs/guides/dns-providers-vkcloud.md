---
page_title: "vkcloud"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# VK Cloud DNS Challenge Provider

The `vkcloud` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[VK Cloud](https://mcs.mail.ru/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "vkcloud"
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

* `VK_CLOUD_PASSWORD` - Password for VK Cloud account.
* `VK_CLOUD_PROJECT_ID` - String ID of project in VK Cloud.
* `VK_CLOUD_USERNAME` - Email of VK Cloud account.

* `VK_CLOUD_DNS_ENDPOINT` - URL of DNS API. Defaults to https://mcs.mail.ru/public-dns but can be changed for usage with private clouds.
* `VK_CLOUD_DOMAIN_NAME` - Openstack users domain name. Defaults to `users` but can be changed for usage with private clouds.
* `VK_CLOUD_IDENTITY_ENDPOINT` - URL of OpenStack Auth API, Defaults to https://infra.mail.ru:35357/v3/ but can be changed for usage with private clouds.
* `VK_CLOUD_POLLING_INTERVAL` - Time between DNS propagation check.
* `VK_CLOUD_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `VK_CLOUD_TTL` - The TTL of the TXT record used for the DNS challenge.

## Credential information

You can find all required and additional information on ["Project/Keys" page](https://mcs.mail.ru/app/en/project/keys) of your cloud.

| ENV Variable               | Parameter from page |
|----------------------------|---------------------|
| VK_CLOUD_PROJECT_ID        | Project ID          |
| VK_CLOUD_USERNAME          | Username            |
| VK_CLOUD_DOMAIN_NAME       | User Domain Name    |
| VK_CLOUD_IDENTITY_ENDPOINT | Identity endpoint   |

