---
page_title: "nicru"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# RU CENTER DNS Challenge Provider

The `nicru` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[RU CENTER](https://nic.ru/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "nicru"
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

* `NICRU_PASSWORD` - Password for an account in RU CENTER.
* `NICRU_SECRET` - Secret for application in DNS-hosting RU CENTER.
* `NICRU_SERVICE_ID` - Service ID for application in DNS-hosting RU CENTER.
* `NICRU_SERVICE_NAME` - Service Name for DNS-hosting RU CENTER.
* `NICRU_USER` - Agreement for an account in RU CENTER.

* `NICRU_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 60).
* `NICRU_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 600).
* `NICRU_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 30).

## Credential information

You can find information about service ID and secret https://www.nic.ru/manager/oauth.cgi?step=oauth.app_list

| ENV Variable        | Parameter from page            | Example           |
|---------------------|--------------------------------|-------------------|
| NICRU_USER          | Username (Number of agreement) | NNNNNNN/NIC-D     |
| NICRU_PASSWORD      | Password account               |                   |
| NICRU_SERVICE_ID    | Application ID                 | hex-based, len 32 |
| NICRU_SECRET        | Identity endpoint              | string len 91     |

