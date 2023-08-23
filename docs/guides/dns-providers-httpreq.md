---
page_title: "httpreq"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# HTTP request DNS Challenge Provider

The `httpreq` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
HTTP request.

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "httpreq"
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

* `HTTPREQ_ENDPOINT` - The URL of the server.
* `HTTPREQ_MODE` - `RAW`, none.

* `HTTPREQ_HTTP_TIMEOUT` - API request timeout.
* `HTTPREQ_PASSWORD` - Basic authentication password.
* `HTTPREQ_POLLING_INTERVAL` - Time between DNS propagation check.
* `HTTPREQ_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `HTTPREQ_USERNAME` - Basic authentication username.

## Description

The server must provide:

- `POST` `/present`
- `POST` `/cleanup`

The URL of the server must be defined by `HTTPREQ_ENDPOINT`.

### Mode

There are 2 modes (`HTTPREQ_MODE`):

- default mode:
```json
{
  "fqdn": "_acme-challenge.domain.",
  "value": "LHDhK3oGRvkiefQnx7OOczTY5Tic_xZ6HcMOc_gmtoM"
}
```

- `RAW`
```json
{
  "domain": "domain",
  "token": "token",
  "keyAuth": "key"
}
```

### Authentication

Basic authentication (optional) can be set with some environment variables:

- `HTTPREQ_USERNAME` and `HTTPREQ_PASSWORD`
- both values must be set, otherwise basic authentication is not defined.


