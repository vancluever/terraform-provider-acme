---
layout: "acme"
page_title: "ACME: HTTP DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-httpreq"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# HTTP DNS Challenge Provider

The `httpreq` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource by interacting with
a generic HTTP endpoint.

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "httpreq"
  }
}
```

## Usage Details

The server must provide the endpoints outlined below. With the exception of
anything specified below, the client follows the default behavior outlined in
Go's [`net/http` `Client` documentation][net-http-client-doc].

[net-http-client-doc]: https://golang.org/pkg/net/http/#Client

### `POST /present`

This endpoint is used when presenting the TXT record to create. The payload can
be either in default mode, or raw mode. This is defined by the `HTTPREQ_MODE`
argument supplied to the DNS challenge. The Content-Type sent is
`application/json`. 

#### Default mode payload

```json
{
  "fqdn": "_acme-challenge.domain.",
  "value": "LHDhK3oGRvkiefQnx7OOczTY5Tic_xZ6HcMOc_gmtoM"
}
```

#### Raw mode payload

```json
{
  "domain": "domain",
  "token": "token",
  "keyAuth": "key"
}
```

### `POST /cleanup`

This endpoint is used to clean up the DNS challenge records during teardown. The
payload is exactly the same as outlined above.

## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

* `HTTPREQ_ENDPOINT` - The base URL path to use. This can include an URI base,
  example: `https://example.com/foobar`.
* `HTTPREQ_MODE` - The payload mode to use. If set to `RAW`, raw mode is used,
  otherwise the default mode is used.
* `HTTPREQ_USERNAME` - The username to use for HTTP basic authentication, if
  any.
* `HTTPREQ_PASSWORD` - The password to use for HTTP basic authentication, if
  any.

The following additional optional variables are available:

* `HTTP_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `HTTP_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
* `HTTP_HTTP_TIMEOUT` - The timeout on HTTP requests to the API (default:
  `30`).
