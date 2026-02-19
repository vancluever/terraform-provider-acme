---
page_title: "bluecatv2"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Bluecat v2 DNS Challenge Provider

The `bluecatv2` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Bluecat v2](https://www.bluecatnetworks.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "bluecatv2"
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

* `BLUECATV2_CONFIG_NAME` - Configuration name.
* `BLUECATV2_PASSWORD` - API password.
* `BLUECATV2_USERNAME` - API username.
* `BLUECATV2_VIEW_NAME` - DNS View Name.
* `BLUECAT_SERVER_URL` - The server URL: it should have a scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve.

* `BLUECATV2_HTTP_TIMEOUT` - API request timeout in seconds (Default: 30).
* `BLUECATV2_POLLING_INTERVAL` - Time between DNS propagation check in seconds (Default: 2).
* `BLUECATV2_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation in seconds (Default: 60).
* `BLUECATV2_SKIP_DEPLOY` - Skip quick deployements.
* `BLUECATV2_TTL` - The TTL of the TXT record used for the DNS challenge in seconds (Default: 120).


