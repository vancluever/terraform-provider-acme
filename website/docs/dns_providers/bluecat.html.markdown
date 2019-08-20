---
layout: "acme"
page_title: "ACME: Bluecat DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-bluecat"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Bluecat DNS Challenge Provider

The `bluecat` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Bluecat][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://www.bluecatnetworks.com

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "bluecat"
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

* `BLUECAT_CONFIG_NAME` - Configuration name.
* `BLUECAT_DNS_VIEW` - External DNS View Name.
* `BLUECAT_PASSWORD` - API password.
* `BLUECAT_SERVER_URL` - The server URL, should have scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve.
* `BLUECAT_USER_NAME` - API username.

The following additional optional variables are available:

* `BLUECAT_HTTP_TIMEOUT` - API request timeout.
* `BLUECAT_POLLING_INTERVAL` - Time between DNS propagation check.
* `BLUECAT_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `BLUECAT_TTL` - The TTL of the TXT record used for the DNS challenge.


