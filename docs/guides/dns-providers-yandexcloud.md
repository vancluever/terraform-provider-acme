---
page_title: "yandexcloud"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Yandex Cloud DNS Challenge Provider

The `yandexcloud` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Yandex Cloud](https://cloud.yandex.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "yandexcloud"
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

* `YANDEX_CLOUD_FOLDER_ID` - The string id of folder (aka project) in Yandex Cloud.
* `YANDEX_CLOUD_IAM_TOKEN` - The base64 encoded json which contains information about iam token of service account with `dns.admin` permissions.

* `YANDEX_CLOUD_POLLING_INTERVAL` - Time between DNS propagation check.
* `YANDEX_CLOUD_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `YANDEX_CLOUD_TTL` - The TTL of the TXT record used for the DNS challenge.

## IAM Token

The simplest way to retrieve IAM access token is usage of yc-cli,
follow [docs](https://cloud.yandex.ru/docs/iam/operations/iam-token/create-for-sa) to get it

```bash
yc iam key create --service-account-name my-robot --output key.json
cat key.json | base64
```

