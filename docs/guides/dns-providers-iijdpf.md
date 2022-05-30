---
page_title: "iijdpf"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# IIJ DNS Platform Service DNS Challenge Provider

The `iijdpf` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[IIJ DNS Platform Service](https://www.iij.ad.jp/en/biz/dns-pfm/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "iijdpf"
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

* `IIJ_DPF_API_TOKEN` - API token.
* `IIJ_DPF_DPM_SERVICE_CODE` - IIJ Managed DNS Service's service code.

* `IIJ_DPF_API_ENDPOINT` - API endpoint URL, defaults to https://api.dns-platform.jp/dpf/v1.
* `IIJ_DPF_POLLING_INTERVAL` - Time between DNS propagation check, defaults to 5 second.
* `IIJ_DPF_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation, defaults to 660 second.
* `IIJ_DPF_TTL` - The TTL of the TXT record used for the DNS challenge, default to 300.


