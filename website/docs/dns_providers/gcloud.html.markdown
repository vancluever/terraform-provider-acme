---
layout: "acme"
page_title: "ACME: Google Cloud DNS DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-gcloud"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Google Cloud DNS DNS Challenge Provider

The `gcloud` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Google Cloud DNS][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://cloud.google.com/dns/docs/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "gcloud"
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

* `GCE_PROJECT` - The project name.
* `GCE_SERVICE_ACCOUNT_FILE` - The path to the service account file. This is
  the same file referenced by the
  [`credentials`][tf-provider-google-credentials] option in the [Terraform
  Google provider][tf-provider-google].

[tf-provider-google-credentials]: /docs/providers/google/index.html#credentials
[tf-provider-google]: /docs/providers/google/index.html

The following additional optional variables are available:

* `GCE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `5`).
* `GCE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `180`).
* `GCE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `120`).
