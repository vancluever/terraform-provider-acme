---
page_title: "route53"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Amazon Route 53 DNS Challenge Provider

The `route53` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Amazon Route 53](https://aws.amazon.com/route53/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

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

[resource-acme-certificate-dns-challenge-arg]: ../resources/certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ../resources/certificate.md#using-variable-files-for-provider-arguments

* `AWS_ACCESS_KEY_ID` - Managed by the AWS client. Access key ID (`AWS_ACCESS_KEY_ID_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead).
* `AWS_ASSUME_ROLE_ARN` - Managed by the AWS Role ARN (`AWS_ASSUME_ROLE_ARN_FILE` is not supported).
* `AWS_EXTERNAL_ID` - Managed by STS AssumeRole API operation (`AWS_EXTERNAL_ID_FILE` is not supported).
* `AWS_HOSTED_ZONE_ID` - Override the hosted zone ID..
* `AWS_PROFILE` - Managed by the AWS client (`AWS_PROFILE_FILE` is not supported).
* `AWS_REGION` - Managed by the AWS client (`AWS_REGION_FILE` is not supported).
* `AWS_SDK_LOAD_CONFIG` - Managed by the AWS client. Retrieve the region from the CLI config file (`AWS_SDK_LOAD_CONFIG_FILE` is not supported).
* `AWS_SECRET_ACCESS_KEY` - Managed by the AWS client. Secret access key (`AWS_SECRET_ACCESS_KEY_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead).
* `AWS_WAIT_FOR_RECORD_SETS_CHANGED` - Wait for changes to be INSYNC (it can be unstable).

* `AWS_MAX_RETRIES` - The number of maximum returns the service will use to make an individual API request.
* `AWS_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 4).
* `AWS_PRIVATE_ZONE` - Set to true to use private zones only (default: use public zones only).
* `AWS_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 120).
* `AWS_SHARED_CREDENTIALS_FILE` - Managed by the AWS client. Shared credentials file..
* `AWS_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 10).

## Description

AWS Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, [`AWS_SESSION_TOKEN`]
2. Shared credentials file (defaults to `~/.aws/credentials`, profiles can be specified using `AWS_PROFILE`)
3. Amazon EC2 IAM role

The AWS Region is automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_REGION`
2. Shared configuration file if `AWS_SDK_LOAD_CONFIG` is set (defaults to `~/.aws/config`, profiles can be specified using `AWS_PROFILE`)

If `AWS_HOSTED_ZONE_ID` is not set, Lego tries to determine the correct public hosted zone via the FQDN.

See also:

- [sessions](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html)
- [Setting AWS Credentials](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials)
- [Setting AWS Region](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-the-region)

## IAM Policy Examples

### Broad privileges for testing purposes

The following [IAM policy](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html) document grants access to the required APIs needed by lego to complete the DNS challenge.
A word of caution:
These permissions grant write access to any DNS record in any hosted zone,
so it is recommended to narrow them down as much as possible if you are using this policy in production.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
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
      "Effect": "Allow",
      "Action": "route53:ListHostedZonesByName",
      "Resource": "*"
    }
  ]
}
```

### Least privilege policy for production purposes

The following AWS IAM policy document describes the least privilege permissions required for lego to complete the DNS challenge.
Write access is limited to a specified hosted zone's DNS TXT records with a key of `_acme-challenge.example.com`.
Replace `Z11111112222222333333` with your hosted zone ID and `example.com` with your domain name to use this policy.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "route53:GetChange",
      "Resource": "arn:aws:route53:::change/*"
    },
    {
      "Effect": "Allow",
      "Action": "route53:ListHostedZonesByName",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/Z11111112222222333333"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/Z11111112222222333333"
      ],
      "Condition": {
        "ForAllValues:StringEquals": {
          "route53:ChangeResourceRecordSetsNormalizedRecordNames": [
            "_acme-challenge.example.com"
          ],
          "route53:ChangeResourceRecordSetsRecordTypes": [
            "TXT"
          ]
        }
      }
    }
  ]
}
```

