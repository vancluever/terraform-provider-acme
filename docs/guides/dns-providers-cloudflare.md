---
page_title: "cloudflare"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Cloudflare DNS Challenge Provider

The `cloudflare` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Cloudflare](https://www.cloudflare.com/dns/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "cloudflare"
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

* `CF_API_EMAIL` - Account email.
* `CF_API_KEY` - API key.
* `CF_DNS_API_TOKEN` - API token with DNS:Edit permission (since v3.1.0).
* `CF_ZONE_API_TOKEN` - API token with Zone:Read permission (since v3.1.0).
* `CLOUDFLARE_API_KEY` - Alias to CF_API_KEY.
* `CLOUDFLARE_DNS_API_TOKEN` - Alias to CF_DNS_API_TOKEN.
* `CLOUDFLARE_EMAIL` - Alias to CF_API_EMAIL.
* `CLOUDFLARE_ZONE_API_TOKEN` - Alias to CF_ZONE_API_TOKEN.

* `CLOUDFLARE_HTTP_TIMEOUT` - API request timeout.
* `CLOUDFLARE_POLLING_INTERVAL` - Time between DNS propagation check.
* `CLOUDFLARE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `CLOUDFLARE_TTL` - The TTL of the TXT record used for the DNS challenge.

## Description

You may use `CF_API_EMAIL` and `CF_API_KEY` to authenticate, or `CF_DNS_API_TOKEN`, or `CF_DNS_API_TOKEN` and `CF_ZONE_API_TOKEN`.

### API keys

If using API keys (`CF_API_EMAIL` and `CF_API_KEY`), the Global API Key needs to be used, not the Origin CA Key.

Please be aware, that this in principle allows Lego to read and change *everything* related to this account.

### API tokens

With API tokens (`CF_DNS_API_TOKEN`, and optionally `CF_ZONE_API_TOKEN`),
very specific access can be granted to your resources at Cloudflare.
See this [Cloudflare announcement](https://blog.cloudflare.com/api-tokens-general-availability/) for details.

The main resources Lego cares for are the DNS entries for your Zones.
It also needs to resolve a domain name to an internal Zone ID in order to manipulate DNS entries.

Hence, you should create an API token with the following permissions:

* Zone / Zone / Read
* Zone / DNS / Edit

You also need to scope the access to all your domains for this to work.
Then pass the API token as `CF_DNS_API_TOKEN` to Lego.

**Alternatively,** if you prefer a more strict set of privileges,
you can split the access tokens:

* Create one with *Zone / Zone / Read* permissions and scope it to all your zones.
  This is needed to resolve domain names to Zone IDs and can be shared among multiple Lego installations.
  Pass this API token as `CF_ZONE_API_TOKEN` to Lego.
* Create another API token with *Zone / DNS / Edit* permissions and set the scope to the domains you want to manage with a single Lego installation.
  Pass this token as `CF_DNS_API_TOKEN` to Lego.
* Repeat the previous step for each host you want to run Lego on.

This "paranoid" setup is mainly interesting for users who manage many zones/domains with a single Cloudflare account.
It follows the principle of least privilege and limits the possible damage, should one of the hosts become compromised.

