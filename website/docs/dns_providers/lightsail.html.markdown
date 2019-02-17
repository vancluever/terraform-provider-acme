---
layout: "acme"
page_title: "ACME: Amazon Lightsail DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-lightsail"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Amazon Lightsail DNS Challenge Provider

The `lightsail` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Amazon Lightsail][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://aws.amazon.com/lightsail/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "lightsail"
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
* `DNS_ZONE` - The hosted zone ID to use.

The following additional optional variables are available:

* `LIGHTSAIL_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `LIGHTSAIL_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
