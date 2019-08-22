---
layout: "acme"
page_title: "ACME: PowerDNS DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-pdns"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---
<br>

-> **NOTE:** The following documentation is auto-generated from the
ACME provider's API library [lego](https://go-acme.github.io/lego/).
Some sections may refer to lego directly - in most cases, these
sections apply to the Terraform provider as well.

# PowerDNS DNS Challenge Provider

The `pdns` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[PowerDNS](https://www.powerdns.com/).

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "pdns"
  }
}
```
## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

* `PDNS_API_KEY` - API key.
* `PDNS_API_URL` - API url.

* `PDNS_HTTP_TIMEOUT` - API request timeout.
* `PDNS_POLLING_INTERVAL` - Time between DNS propagation check.
* `PDNS_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `PDNS_TTL` - The TTL of the TXT record used for the DNS challenge.

## Information

Tested and confirmed to work with PowerDNS authoritative server 3.4.8 and 4.0.1. Refer to [PowerDNS documentation](https://doc.powerdns.com/md/httpapi/README/) instructions on how to enable the built-in API interface.

PowerDNS Notes:
- PowerDNS API does not currently support SSL, therefore you should take care to ensure that traffic between lego and the PowerDNS API is over a trusted network, VPN etc.
- In order to have the SOA serial automatically increment each time the `_acme-challenge` record is added/modified via the API, set `SOA-EDIT-API` to `INCEPTION-INCREMENT` for the zone in the `domainmetadata` table

