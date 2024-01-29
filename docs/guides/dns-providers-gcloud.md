---
page_title: "gcloud"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Google Cloud DNS Challenge Provider

The `gcloud` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Google Cloud](https://cloud.google.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

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

[resource-acme-certificate-dns-challenge-arg]: ../resources/certificate.md#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: ../resources/certificate.md#using-variable-files-for-provider-arguments

* `Application Default Credentials` - [Documentation](https://cloud.google.com/docs/authentication/production#providing_credentials_to_your_application).
* `GCE_PROJECT` - Project name (by default, the project name is auto-detected by using the metadata service).
* `GCE_SERVICE_ACCOUNT` - Account.
* `GCE_SERVICE_ACCOUNT_FILE` - Account file path.

* `GCE_ALLOW_PRIVATE_ZONE` - Allows requested domain to be in private DNS zone, works only with a private ACME server (by default: false).
* `GCE_POLLING_INTERVAL` - Time between DNS propagation check.
* `GCE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `GCE_TTL` - The TTL of the TXT record used for the DNS challenge.
* `GCE_ZONE_ID` - Allows to skip the automatic detection of the zone.


