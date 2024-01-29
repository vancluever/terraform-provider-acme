---
page_title: "designate"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Designate DNSaaS for Openstack DNS Challenge Provider

The `designate` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Designate DNSaaS for Openstack](https://docs.openstack.org/designate/latest/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "designate"
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

* `OS_APPLICATION_CREDENTIAL_ID` - Application credential ID.
* `OS_APPLICATION_CREDENTIAL_NAME` - Application credential name.
* `OS_APPLICATION_CREDENTIAL_SECRET` - Application credential secret.
* `OS_AUTH_URL` - Identity endpoint URL.
* `OS_PASSWORD` - Password.
* `OS_PROJECT_NAME` - Project name.
* `OS_REGION_NAME` - Region name.
* `OS_USERNAME` - Username.
* `OS_USER_ID` - User ID.

* `DESIGNATE_POLLING_INTERVAL` - Time between DNS propagation check.
* `DESIGNATE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `DESIGNATE_TTL` - The TTL of the TXT record used for the DNS challenge.
* `OS_PROJECT_ID` - Project ID.
* `OS_TENANT_NAME` - Tenant name (deprecated see OS_PROJECT_NAME and OS_PROJECT_ID).

## Description

There are three main ways of authenticating with Designate:

1. The first one is by using the `OS_CLOUD` environment variable and a `clouds.yaml` file.
2. The second one is using your username and password, via the `OS_USERNAME`, `OS_PASSWORD` and `OS_PROJECT_NAME` environment variables.
3. The third one is by using an application credential, via the `OS_APPLICATION_CREDENTIAL_*` and `OS_USER_ID` environment variables.

For the username/password and application methods, the `OS_AUTH_URL` and `OS_REGION_NAME` environment variables are required.

For more information, you can read about the different methods of authentication with OpenStack in the Keystone's documentation and the gophercloud documentation:

- [Keystone username/password](https://docs.openstack.org/keystone/latest/user/supported_clients.html)
- [Keystone application credentials](https://docs.openstack.org/keystone/latest/user/application_credentials.html)

Public cloud providers with support for Designate:

- [Fuga Cloud](https://fuga.cloud/)

