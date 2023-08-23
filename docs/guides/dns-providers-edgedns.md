---
page_title: "edgedns"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Akamai EdgeDNS DNS Challenge Provider

The `edgedns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Akamai EdgeDNS](https://www.akamai.com/us/en/products/security/edge-dns.jsp).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "edgedns"
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

* `AKAMAI_ACCESS_TOKEN` - Access token, managed by the Akamai EdgeGrid client.
* `AKAMAI_CLIENT_SECRET` - Client secret, managed by the Akamai EdgeGrid client.
* `AKAMAI_CLIENT_TOKEN` - Client token, managed by the Akamai EdgeGrid client.
* `AKAMAI_EDGERC` - Path to the .edgerc file, managed by the Akamai EdgeGrid client.
* `AKAMAI_EDGERC_SECTION` - Configuration section, managed by the Akamai EdgeGrid client.
* `AKAMAI_HOST` - API host, managed by the Akamai EdgeGrid client.

* `AKAMAI_POLLING_INTERVAL` - Time between DNS propagation check. Default: 15 seconds.
* `AKAMAI_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation. Default: 3 minutes.
* `AKAMAI_TTL` - The TTL of the TXT record used for the DNS challenge.

Akamai's credentials are automatically detected in the following locations and prioritized in the following order:

1. Section-specific environment variables (where `{SECTION}` is specified using `AKAMAI_EDGERC_SECTION`):
  - `AKAMAI_{SECTION}_HOST`
  - `AKAMAI_{SECTION}_ACCESS_TOKEN`
  - `AKAMAI_{SECTION}_CLIENT_TOKEN`
  - `AKAMAI_{SECTION}_CLIENT_SECRET`
2. If `AKAMAI_EDGERC_SECTION` is not defined or is set to `default`, environment variables:
  - `AKAMAI_HOST`
  - `AKAMAI_ACCESS_TOKEN`
  - `AKAMAI_CLIENT_TOKEN`
  - `AKAMAI_CLIENT_SECRET`
3. `.edgerc` file located at `AKAMAI_EDGERC`
  - defaults to `~/.edgerc`, sections can be specified using `AKAMAI_EDGERC_SECTION`
4. Default environment variables:
  - `AKAMAI_HOST`
  - `AKAMAI_ACCESS_TOKEN`
  - `AKAMAI_CLIENT_TOKEN`
  - `AKAMAI_CLIENT_SECRET`

See also:

- [Setting up Akamai credentials](https://developer.akamai.com/api/getting-started)
- [.edgerc Format](https://developer.akamai.com/legacy/introduction/Conf_Client.html#edgercformat)
- [API Client Authentication](https://developer.akamai.com/legacy/introduction/Client_Auth.html)
- [Config from Env](https://github.com/akamai/AkamaiOPEN-edgegrid-golang/blob/master/pkg/edgegrid/config.go#L118)

