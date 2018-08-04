---
layout: "acme"
page_title: "ACME: Amazon Route 53 DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-route53"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Amazon Route 53 DNS Challenge Provider

The `route53` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Amazon Route 53][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://route53.microsoft.com/en-ca/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "route53"
  }
}
```

## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

-> **NOTE:** Several other options exist for configuring the AWS credential
chain. For more details, see the [AWS SDK documentation][aws-sdk-docs].

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge
[aws-sdk-docs]: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html

* `AWS_ACCESS_KEY_ID` - The AWS access key ID.
* `AWS_SECRET_ACCESS_KEY` - The AWS secret access key.
* `AWS_SESSION_TOKEN` - The session token to use, if necessary.
* `AWS_HOSTED_ZONE_ID` - The hosted zone ID to use. This can be used to
  override ACME's default domain discovery and force the provider to use a
  specific hosted zone.
