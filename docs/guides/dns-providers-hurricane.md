---
page_title: "hurricane"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# Hurricane Electric DNS DNS Challenge Provider

The `hurricane` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[Hurricane Electric DNS](https://dns.he.net/).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "hurricane"
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

* `HURRICANE_TOKENS` - TXT record names and tokens.


Before using lego to request a certificate for a given domain or wildcard (such as `my.example.org` or `*.my.example.org`),
create a TXT record named `_acme-challenge.my.example.org`, and enable dynamic updates on it.
Generate a token for each URL with Hurricane Electric's UI, and copy it down.
Stick to alphanumeric tokens for greatest reliability.

To authenticate with the Hurricane Electric API,
add each record name/token pair you want to update to the `HURRICANE_TOKENS` environment variable, as shown in the examples.
Record names (without the `_acme-challenge.` component) and their tokens are separated with colons,
while the credential pairs are concatenated into a comma-separated list, like so:

```
HURRICANE_TOKENS=my.example.org:token1,demo.example.org:token2
```

If you are issuing both a wildcard certificate and a standard certificate for a given subdomain,
you should not have repeat entries for that name, as both will use the same credential.

```
HURRICANE_TOKENS=example.org:token
```

