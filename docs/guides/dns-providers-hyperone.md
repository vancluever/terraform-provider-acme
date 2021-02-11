---
page_title: "hyperone"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# HyperOne DNS Challenge Provider

The `hyperone` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[HyperOne](https://www.hyperone.com).

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "hyperone"
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


* `HYPERONE_API_URL` - Allows to pass custom API Endpoint to be used in the challenge (default https://api.hyperone.com/v2).
* `HYPERONE_LOCATION_ID` - Specifies location (region) to be used in API calls. (default pl-waw-1).
* `HYPERONE_PASSPORT_LOCATION` - Allows to pass custom passport file location (default ~/.h1/passport.json).
* `HYPERONE_POLLING_INTERVAL` - Time between DNS propagation check.
* `HYPERONE_PROPAGATION_TIMEOUT` - Maximum waiting time for DNS propagation.
* `HYPERONE_TTL` - The TTL of the TXT record used for the DNS challenge.

## Description

Default configuration does not require any additional environment variables,
just a passport file in `~/.h1/passport.json` location.

### Generating passport file using H1 CLI

To use this application you have to generate passport file for `sa`:

```
h1 iam project sa credential generate --name my-passport --project <project ID> --sa <sa ID> --passport-output-file ~/.h1/passport.json
```

### Required permissions

The application requires following permissions:
-  `dns/zone/list`
-  `dns/zone.recordset/list`
-  `dns/zone.recordset/create`
-  `dns/zone.recordset/delete`
-  `dns/zone.record/create`
-  `dns/zone.record/list`
-  `dns/zone.record/delete`

All required permissions are available via platform role `tool.lego`.

