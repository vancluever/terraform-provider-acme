---
page_title: "selfhostde"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# SelfHost.(de|eu) DNS Challenge Provider

The `selfhostde` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[SelfHost.(de|eu)](https://www.selfhost.de).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "selfhostde"
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

* `SELFHOSTDE_PASSWORD` - Password.
* `SELFHOSTDE_RECORDS_MAPPING` - Record IDs mapping with domains (ex: example.com:123:456,example.org:789,foo.example.com:147).
* `SELFHOSTDE_USERNAME` - Username.

* `SELFHOSTDE_HTTP_TIMEOUT` - API request timeout.
* `SELFHOSTDE_POLLING_INTERVAL` - Time between DNS propagation check.
* `SELFHOSTDE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `SELFHOSTDE_TTL` - The TTL of the TXT record used for the DNS challenge.

SelfHost.de doesn't have an API to create or delete TXT records,
there is only an "unofficial" and undocumented endpoint to update an existing TXT record.

So, before using lego to request a certificate for a given domain or wildcard (such as `my.example.org` or `*.my.example.org`),
you must create:

- one TXT record named `_acme-challenge.my.example.org` if you are **not** using wildcard for this domain.
- two TXT records named `_acme-challenge.my.example.org` if you are using wildcard for this domain.

After that you must edit the TXT record(s) to get the ID(s).

You then must prepare the `SELFHOSTDE_RECORDS_MAPPING` environment variable with the following format:

```
<domain_A>:<record_id_A1>:<record_id_A2>,<domain_B>:<record_id_B1>:<record_id_B2>,<domain_C>:<record_id_C1>:<record_id_C2>
```

where each group of domain + record ID(s) is separated with a comma (`,`),
and the domain and record ID(s) are separated with a colon (`:`).

For example, if you want to create or renew a certificate for `my.example.org`, `*.my.example.org`, and `other.example.org`,
you would need:

- two separate records for `_acme-challenge.my.example.org`
- and another separate record for `_acme-challenge.other.example.org`

The resulting environment variable would then be: `SELFHOSTDE_RECORDS_MAPPING=my.example.com:123:456,other.example.com:789`


