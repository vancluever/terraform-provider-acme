---
layout: "acme"
page_title: "ACME: Cloudflare DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-cloudflare"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---
<br>

-> **NOTE:** The following documentation is auto-generated from the
ACME provider's API library [lego](https://go-acme.github.io/lego/).
Some sections may refer to lego directly - in most cases, these
sections apply to the Terraform provider as well.

# Cloudflare DNS Challenge Provider

The `cloudflare` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Cloudflare](https://www.cloudflare.com/dns/).

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

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

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

* `CF_API_EMAIL` - Account email.
* `CF_API_KEY` - API key.
* `CLOUDFLARE_API_KEY` - Alias to CLOUDFLARE_API_KEY.
* `CLOUDFLARE_EMAIL` - Alias to CF_API_EMAIL.

* `CLOUDFLARE_HTTP_TIMEOUT` - API request timeout.
* `CLOUDFLARE_POLLING_INTERVAL` - Time between DNS propagation check.
* `CLOUDFLARE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `CLOUDFLARE_TTL` - The TTL of the TXT record used for the DNS challenge.

The Global API Key needs to be used, not the Origin CA Key.

