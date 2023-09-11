---
page_title: "azuredns"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Azure DNS DNS Challenge Provider

The `azuredns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Azure DNS](https://azure.microsoft.com/services/dns/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "azuredns"
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
* `AZURE_RESOURCE_GROUP` - DNS zone resource group.
* `AZURE_SUBSCRIPTION_ID` - DNS zone subscription ID.
* `AZURE_TENANT_ID` - Tenant ID.

* `AZURE_ENVIRONMENT` - Azure environment, one of: public, usgovernment, and china.
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

## Description

Azure Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables for client secret: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_SECRET`
2. Environment variables for client certificate: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_CERTIFICATE_PATH`
3. Workload identity for resources hosted in Azure environment (see below)
4. Shared credentials file (defaults to `~/.azure`), used by Azure CLI

Link:
- [Azure Authentication](https://learn.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication)

### Workload identity

#### Azure Managed Identity

The Azure Managed Identity service allows linking Azure AD identities to Azure resources, without needing to manually manage client IDs and secrets.

Workloads with a Managed Identity can manage their own certificates, with permissions on specific domain names set using IAM assignments.
For this to work, the Managed Identity requires the **Reader** role on the target DNS Zone,
and the **DNS Zone Contributor** on the relevant `_acme-challenge` TXT records.

For example, to allow a Managed Identity to create a certificate for "fw01.lab.example.com", using Azure CLI:

```bash
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
export AZURE_RESOURCE_GROUP="rg1"
export SERVICE_PRINCIPAL_ID="00000000-0000-0000-0000-000000000000"

export AZURE_DNS_ZONE="lab.example.com"
export AZ_HOSTNAME="fw01"
export AZ_RECORD_SET="_acme-challenge.${AZ_HOSTNAME}"

az role assignment create \
--assignee "${SERVICE_PRINCIPAL_ID}" \
--role "Reader" \
--scope "/subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/${AZURE_RESOURCE_GROUP}/providers/Microsoft.Network/dnszones/${AZURE_DNS_ZONE}"

az role assignment create \
--assignee "${SERVICE_PRINCIPAL_ID}" \
--role "DNS Zone Contributor" \
--scope "/subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/${AZURE_RESOURCE_GROUP}/providers/Microsoft.Network/dnszones/${AZURE_DNS_ZONE}/TXT/${AZ_RECORD_SET}"
```

#### Azure Managed Identity (with Azure Arc)

The Azure Arc agent provides the ability to use a Managed Identity on resources hosted outside of Azure
(such as on-prem virtual machines, or VMs in another cloud provider).

While the upstream `azidentity` SDK will try to automatically identify and use the Azure Arc metadata service,
if you get `azuredns: DefaultAzureCredential: failed to acquire a token.` error messages,
you may need to set the environment variables:
  * `IMDS_ENDPOINT=http://localhost:40342`
  * `IDENTITY_ENDPOINT=http://localhost:40342/metadata/identity/oauth2/token`

#### Workload identity for AKS

Workload identity allows workloads running Azure Kubernetes Services (AKS) clusters to authenticate as an Azure AD application identity using federated credentials.

This must be configured in kubernetes workload deployment in one hand and on the Azure AD application registration in the other hand.

Here is a summary of the steps to follow to use it :
* create a `ServiceAccount` resource, add following annotations to reference the targeted Azure AD application registration : `azure.workload.identity/client-id` and `azure.workload.identity/tenant-id`.
* on the `Deployment` resource you must reference the previous `ServiceAccount` and add the following label : `azure.workload.identity/use: "true"`.
* create a fedreated credentials of type `Kubernetes accessing Azure resources`, add the cluster issuer URL  and add the namespace and name of your kubernetes service account.

Link :
- [Azure AD Workload identity](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html)


