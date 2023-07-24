---
page_title: "azure"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Azure (deprecated) DNS Challenge Provider

The `azure` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Azure (deprecated)](https://azure.microsoft.com/services/dns/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "azure"
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

* `AZURE_CLIENT_ID` - Client ID.
* `AZURE_CLIENT_SECRET` - Client secret.
* `AZURE_ENVIRONMENT` - Azure environment, one of: public, usgovernment, german, and china.
* `AZURE_RESOURCE_GROUP` - Resource group.
* `AZURE_SUBSCRIPTION_ID` - Subscription ID.
* `AZURE_TENANT_ID` - Tenant ID.
* `instance metadata service` - If the credentials are **not** set via the environment, then it will attempt to get a bearer token via the [instance metadata service](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/instance-metadata-service)..

* `AZURE_METADATA_ENDPOINT` - Metadata Service endpoint URL.
* `AZURE_POLLING_INTERVAL` - Time between DNS propagation check.
* `AZURE_PRIVATE_ZONE` - Set to true to use Azure Private DNS Zones and not public.
* `AZURE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `AZURE_TTL` - The TTL of the TXT record used for the DNS challenge.
* `AZURE_ZONE_NAME` - Zone name to use inside Azure DNS service to add the TXT record in.

The following variables are **Terraform-specific** aliases for the above
configuration values:


* `ARM_CLIENT_ID` - alias for `AZURE_CLIENT_ID`.
* `ARM_CLIENT_SECRET` - alias for `AZURE_CLIENT_SECRET`.
* `ARM_RESOURCE_GROUP` - alias for `AZURE_RESOURCE_GROUP`.
* `ARM_SUBSCRIPTION_ID` - alias for `AZURE_SUBSCRIPTION_ID`.
* `ARM_TENANT_ID` - alias for `AZURE_TENANT_ID`.


