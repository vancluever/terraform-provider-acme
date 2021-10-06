---
page_title: "infomaniak"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Infomaniak DNS Challenge Provider

The `infomaniak` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Infomaniak](https://www.infomaniak.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "infomaniak"
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

* `INFOMANIAK_ACCESS_TOKEN` - Access token.

* `INFOMANIAK_ENDPOINT` - https://api.infomaniak.com.
* `INFOMANIAK_HTTP_TIMEOUT` - API request timeout.
* `INFOMANIAK_POLLING_INTERVAL` - Time between DNS propagation check.
* `INFOMANIAK_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `INFOMANIAK_TTL` - The TTL of the TXT record used for the DNS challenge in seconds.

## Access token

Access token can be created at the url https://manager.infomaniak.com/v3/infomaniak-api.
You will need domain scope.

