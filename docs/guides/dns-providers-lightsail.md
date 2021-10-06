---
page_title: "lightsail"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Amazon Lightsail DNS Challenge Provider

The `lightsail` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Amazon Lightsail](https://aws.amazon.com/lightsail/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

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

[resource-acme-certificate-dns-challenge-arg]: ../resources/certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ../resources/certificate.md#using-variable-files-for-provider-arguments

* `AWS_ACCESS_KEY_ID` - Managed by the AWS client. Access key ID (`AWS_ACCESS_KEY_ID_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead).
* `AWS_SECRET_ACCESS_KEY` - Managed by the AWS client. Secret access key (`AWS_SECRET_ACCESS_KEY_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead).
* `DNS_ZONE` - Domain name of the DNS zone.

* `AWS_SHARED_CREDENTIALS_FILE` - Managed by the AWS client. Shared credentials file..
* `LIGHTSAIL_POLLING_INTERVAL` - Time between DNS propagation check.
* `LIGHTSAIL_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.

## Description

AWS Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, [`AWS_SESSION_TOKEN`]
2. Shared credentials file (defaults to `~/.aws/credentials`, profiles can be specified using `AWS_PROFILE`)
3. Amazon EC2 IAM role

AWS region is not required to set as the Lightsail DNS zone is in global (us-east-1) region.

## Policy

The following AWS IAM policy document describes the minimum permissions required for lego to complete the DNS challenge.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "lightsail:DeleteDomainEntry",
        "lightsail:CreateDomainEntry"
      ],
      "Resource": "<Lightsail DNS zone ARN>"
    }
  ]
}
```

Replace the `Resource` value with your Lightsail DNS zone ARN.
You can retrieve the ARN using aws cli by running `aws lightsail get-domains --region us-east-1` (Lightsail web console does not show the ARN, unfortunately).
It should be in the format of `arn:aws:lightsail:global:<ACCOUNT ID>:Domain/<DOMAIN ID>`.
You also need to replace the region in the ARN to `us-east-1` (instead of `global`).

Alternatively, you can also set the `Resource` to `*` (wildcard), which allow to access all domain, but this is not recommended.

