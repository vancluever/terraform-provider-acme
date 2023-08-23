---
page_title: "joker"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Joker DNS Challenge Provider

The `joker` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Joker](https://joker.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "joker"
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

* `JOKER_API_KEY` - API key (only with DMAPI mode).
* `JOKER_API_MODE` - 'DMAPI' or 'SVC'. DMAPI is for resellers accounts. (Default: DMAPI).
* `JOKER_PASSWORD` - Joker.com password.
* `JOKER_USERNAME` - Joker.com username.

* `JOKER_HTTP_TIMEOUT` - API request timeout.
* `JOKER_POLLING_INTERVAL` - Time between DNS propagation check.
* `JOKER_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `JOKER_SEQUENCE_INTERVAL` - Time between sequential requests (only with 'SVC' mode).
* `JOKER_TTL` - The TTL of the TXT record used for the DNS challenge.

## SVC mode

In the SVC mode, username and passsword are not your email and account passwords, but those displayed in Joker.com domain dashboard when enabling Dynamic DNS.

As per [Joker.com documentation](https://joker.com/faq/content/6/496/en/let_s-encrypt-support.html):

> 1. please log in at Joker.com, visit 'My Domains',
>    find the domain you want to add  Let's Encrypt certificate for, and chose "DNS" in the menu
>
> 2. on the top right, you will find the setting for 'Dynamic DNS'.
>    If not already active, please activate it.
>    It will not affect any other already existing DNS records of this domain.
>
> 3. please take a note of the credentials which are now shown as 'Dynamic DNS Authentication', consisting of a 'username' and a 'password'.
>
> 4. this is all you have to do here - and only once per domain.

