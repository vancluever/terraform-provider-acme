---
layout: "acme"
page_title: "ACME: Amazon Route 53 DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-route53"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---
<br>

-> **NOTE:** The following documentation is auto-generated from the
ACME provider's API library [lego](https://go-acme.github.io/lego/).
Some sections may refer to lego directly - in most cases, these
sections apply to the Terraform provider as well.

# Amazon Route 53 DNS Challenge Provider

The `route53` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Amazon Route 53](https://aws.amazon.com/route53/).

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

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

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

* `AWS_ACCESS_KEY_ID` - Managed by the AWS client.
* `AWS_HOSTED_ZONE_ID` - Override the hosted zone ID.
* `AWS_REGION` - Managed by the AWS client.
* `AWS_SECRET_ACCESS_KEY` - Managed by the AWS client.

* `AWS_MAX_RETRIES` - The number of maximum returns the service will use to make an individual API request.
* `AWS_POLLING_INTERVAL` - Time between DNS propagation check.
* `AWS_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `AWS_TTL` - The TTL of the TXT record used for the DNS challenge.

## Description

AWS Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`, [`AWS_SESSION_TOKEN`]
2. Shared credentials file (defaults to `~/.aws/credentials`)
3. Amazon EC2 IAM role

If `AWS_HOSTED_ZONE_ID` is not set, Lego tries to determine the correct public hosted zone via the FQDN.

See also: [sessions](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html)

## Policy

The following AWS IAM policy document describes the permissions required for lego to complete the DNS challenge.

```json
{
   "Version": "2012-10-17",
   "Statement": [
       {
           "Sid": "",
           "Effect": "Allow",
           "Action": [
               "route53:GetChange",
               "route53:ChangeResourceRecordSets",
               "route53:ListResourceRecordSets"
           ],
           "Resource": [
               "arn:aws:route53:::hostedzone/*",
               "arn:aws:route53:::change/*"
           ]
       },
       {
           "Sid": "",
           "Effect": "Allow",
           "Action": "route53:ListHostedZonesByName",
           "Resource": "*"
       }
   ]
}
```


