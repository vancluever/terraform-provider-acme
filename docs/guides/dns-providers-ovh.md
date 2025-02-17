---
page_title: "ovh"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# OVH DNS Challenge Provider

The `ovh` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[OVH](https://www.ovh.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "ovh"
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

* `OVH_ACCESS_TOKEN` - Access token.
* `OVH_APPLICATION_KEY` - Application key (Application Key authentication).
* `OVH_APPLICATION_SECRET` - Application secret (Application Key authentication).
* `OVH_CLIENT_ID` - Client ID (OAuth2).
* `OVH_CLIENT_SECRET` - Client secret (OAuth2).
* `OVH_CONSUMER_KEY` - Consumer key (Application Key authentication).
* `OVH_ENDPOINT` - Endpoint URL (ovh-eu or ovh-ca).

* `OVH_HTTP_TIMEOUT` - API request timeout in seconds (Default: 180).
* `OVH_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `OVH_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 60).
* `OVH_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 120).

## Application Key and Secret

Application key and secret can be created by following the [OVH guide](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/).

When requesting the consumer key, the following configuration can be used to define access rights:

```json
{
  "accessRules": [
    {
      "method": "POST",
      "path": "/domain/zone/*"
    },
    {
      "method": "DELETE",
      "path": "/domain/zone/*"
    }
  ]
}
```

## OAuth2 Client Credentials

Another method for authentication is by using OAuth2 client credentials.

An IAM policy and service account can be created by following the [OVH guide](https://help.ovhcloud.com/csm/en-manage-service-account?id=kb_article_view&sysparm_article=KB0059343).

Following IAM policies need to be authorized for the affected domain:

* dnsZone:apiovh:record/create
* dnsZone:apiovh:record/delete
* dnsZone:apiovh:refresh

## Important Note

Both authentication methods cannot be used at the same time.

