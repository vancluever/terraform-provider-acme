---
layout: "acme"
page_title: "ACME: External program DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-exec"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# External program DNS Challenge Provider

The `exec` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
[External program][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: #

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "exec"
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


The following additional optional variables are available:



## Base Configuration

| Environment Variable Name | Description                           |
|---------------------------|---------------------------------------|
| `EXEC_MODE`               | `RAW`, none                           |
| `EXEC_PATH`               | The path of the the external program. |


## Additional Configuration

| Environment Variable Name  | Description                               |
|----------------------------|-------------------------------------------|
| `EXEC_POLLING_INTERVAL`    | Time between DNS propagation check.       |
| `EXEC_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation. |


## Description

The file name of the external program is specified in the environment variable `EXEC_PATH`.

When it is run by lego, three command-line parameters are passed to it:
The action ("present" or "cleanup"), the fully-qualified domain name and the value for the record.

For example, requesting a certificate for the domain 'foo.example.com' can be achieved by calling lego as follows:

```bash
EXEC_PATH=./update-dns.sh \
	lego --dns exec \
	--domains foo.example.com \
	--email invalid@example.com run
```

It will then call the program './update-dns.sh' with like this:

```bash
./update-dns.sh "present" "_acme-challenge.foo.example.com." "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI"
```

The program then needs to make sure the record is inserted.
When it returns an error via a non-zero exit code, lego aborts.

When the record is to be removed again,
the program is called with the first command-line parameter set to `cleanup` instead of `present`.

If you want to use the raw domain, token, and keyAuth values with your program, you can set `EXEC_MODE=RAW`:

```bash
EXEC_MODE=RAW \
EXEC_PATH=./update-dns.sh \
	lego --dns exec \
	--domains foo.example.com \
	--email invalid@example.com run
```

It will then call the program `./update-dns.sh` like this:

```bash
./update-dns.sh "present" "foo.example.com." "--" "some-token" "KxAy-J3NwUmg9ZQuM-gP_Mq1nStaYSaP9tYQs5_-YsE.ksT-qywTd8058G-SHHWA3RAN72Pr0yWtPYmmY5UBpQ8"
```

## Commands

{{% notice note %}}
The `--` is because the token MAY start with a `-`, and the called program may try and interpret a `-` as indicating a flag.
In the case of urfave, which is commonly used,
you can use the `--` delimiter to specify the start of positional arguments, and handle such a string safely.
{{% /notice %}}

### Present

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram present -- <FQDN> <record>`             |
| `RAW`   | `myprogram present -- <domain> <token> <key_auth>` |

### Cleanup

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram cleanup -- <FQDN> <record>`             |
| `RAW`   | `myprogram cleanup -- <domain> <token> <key_auth>` |

### Timeout

The command have to display propagation timeout and polling interval into Stdout.

The values must be formatted as JSON, and times are in seconds.
Example: `{"timeout": 30, "interval": 5}`

If an error occurs or if the command is not provided:
the default display propagation timeout and polling interval are used.

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram timeout`                                |
| `RAW`   | `myprogram timeout`                                |


