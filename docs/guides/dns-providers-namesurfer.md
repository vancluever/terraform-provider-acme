---
page_title: "namesurfer"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# FusionLayer NameSurfer DNS Challenge Provider

The `namesurfer` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[FusionLayer NameSurfer](https://www.fusionlayer.com/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "namesurfer"
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

* `NAMESURFER_API_KEY` - API key name.
* `NAMESURFER_API_SECRET` - API secret.
* `NAMESURFER_BASE_URL` - The base URL of NameSurfer API (jsonrpc10) endpoint URL (e.g., https://foo.example.com:8443/API/NSService_10).

* `NAMESURFER_HTTP_TIMEOUT` - API request timeout in seconds (Default: 30).
* `NAMESURFER_INSECURE_SKIP_VERIFY` - Whether to verify the API certificate.
* `NAMESURFER_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `NAMESURFER_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 120).
* `NAMESURFER_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 300).
* `NAMESURFER_VIEW` - DNS view name (optional, default: empty string).


