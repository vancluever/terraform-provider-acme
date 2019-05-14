---
layout: "acme"
page_title: "ACME: Exec DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-exec"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Exec DNS Challenge Provider

The `exec` DNS challenge provider can be used to perform DNS challenges for the
[`acme_certificate`][resource-acme-certificate] resource, using a custom
external script.

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "exec"

    config = {
      "EXEC_PATH" = "./update-dns.sh"
    }
  }
}
```

## Usage Details

The file name of the external script is specified in the environment variable
`EXEC_PATH`. When it is run by Terraform, four command-line parameters are passed
to it: The action ("present" or "cleanup"), the fully-qualified domain name,
the value for the record, and the TTL.

In the above basic example, the `update-dns.sh` script would be called in the
following fashion:

```
./update-dns.sh "present" "_acme-challenge.foo.example.com." "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI" "120"
```

If the script returns a non-zero return code, the execution of the update is
considered to have failed, and Terraform will return an error.

When the record is to be removed, the script is called again, with the first
command-line parameter set to "cleanup" instead of "present".

### Using raw values

If you want to use the raw domain, token, and keyAuth values with your script,
you can set `EXEC_MODE` to `RAW`. When used like this, `update-dns.sh` will be
called in the following way:

```
./update-dns.sh "present" "foo.example.com." "--" "some-token" "KxAy-J3NwUmg9ZQuM-gP_Mq1nStaYSaP9tYQs5_-YsE.ksT-qywTd8058G-SHHWA3RAN72Pr0yWtPYmmY5UBpQ8"
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

-> **NOTE:** Due to the nature of the `exec` provider, it's recommended that
these be supplied as explicit `config` values.

* `EXEC_MODE` - Send the raw domain, token, and keyAuth values to the external
  script. The only usable value here is `RAW`.
* `EXEC_PATH` - The path to the external script to call.

The following additional optional variables are available:

* `EXEC_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `60`).
* `EXEC_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `60`).
