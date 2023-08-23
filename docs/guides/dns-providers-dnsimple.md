---
page_title: "dnsimple"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# DNSimple DNS Challenge Provider

The `dnsimple` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[DNSimple](https://dnsimple.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "dnsimple"
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

* `DNSIMPLE_OAUTH_TOKEN` - OAuth token.

* `DNSIMPLE_BASE_URL` - API endpoint URL.
* `DNSIMPLE_POLLING_INTERVAL` - Time between DNS propagation check.
* `DNSIMPLE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DNSIMPLE_TTL` - The TTL of the TXT record used for the DNS challenge.

## Description

`DNSIMPLE_BASE_URL` is optional and must be set to production (https://api.dnsimple.com).
if `DNSIMPLE_BASE_URL` is not defined or empty, the production URL is used by default.

While you can manage DNS records in the [DNSimple Sandbox environment](https://developer.dnsimple.com/sandbox/),
DNS records will not resolve, and you will not be able to satisfy the ACME DNS challenge.

To authenticate you need to provide a valid API token.
HTTP Basic Authentication is intentionally not supported.

### API tokens

You can [generate a new API token](https://support.dnsimple.com/articles/api-access-token/) from your account page.
Only Account API tokens are supported, if you try to use a User API token you will receive an error message.

