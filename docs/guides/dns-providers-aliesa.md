---
page_title: "aliesa"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# AlibabaCloud ESA DNS Challenge Provider

The `aliesa` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[AlibabaCloud ESA](https://www.alibabacloud.com/en/product/esa).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "aliesa"
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

* `ALIESA_ACCESS_KEY` - Access key ID.
* `ALIESA_RAM_ROLE` - Your instance RAM role (https://www.alibabacloud.com/help/en/ecs/user-guide/attach-an-instance-ram-role-to-an-ecs-instance).
* `ALIESA_SECRET_KEY` - Access Key secret.
* `ALIESA_SECURITY_TOKEN` - STS Security Token (optional).

* `ALIESA_HTTP_TIMEOUT` - API request timeout in seconds (Default: 30).
* `ALIESA_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `ALIESA_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 60).
* `ALIESA_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 120).


