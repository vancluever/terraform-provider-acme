---
layout: "acme"
page_title: "ACME: DNSPod DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-dnspod"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# DNSPod DNS Challenge Provider

The `dnspod` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[DNSPod][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: http://www.dnspod.com/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "dnspod"
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

* `DNSPOD_API_KEY` - The user token.

The following additional optional variables are available:

* `DNSPOD_HTTP_TIMEOUT` - API request timeout.
* `DNSPOD_POLLING_INTERVAL` - Time between DNS propagation check.
* `DNSPOD_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DNSPOD_TTL` - The TTL of the TXT record used for the DNS challenge.


