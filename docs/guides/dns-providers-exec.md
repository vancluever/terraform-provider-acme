---
page_title: "exec"
subcategory: "DNS Providers"
---

-> The following documentation is auto-generated from the ACME
provider's API library [lego](https://go-acme.github.io/lego/).  Some
sections may refer to lego directly - in most cases, these sections
apply to the Terraform provider as well.

# External program DNS Challenge Provider

The `exec` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with
External program.

[resource-acme-certificate]: ../resources/certificate.md

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: ../resources/certificate.md#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "exec"
  }
}
```

## Base Configuration

| Environment Variable Name | Description                           |
|---------------------------|---------------------------------------|
| `EXEC_MODE`               | `RAW`, none                           |
| `EXEC_PATH`               | The path of the the external program. |


## Additional Configuration

| Environment Variable Name  | Description                                                        |
|----------------------------|--------------------------------------------------------------------|
| `EXEC_POLLING_INTERVAL`    | Time between DNS propagation check in seconds (Default: 3).        |
| `EXEC_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation in seconds (Default: 60). |
| `EXEC_SEQUENCE_INTERVAL`   | Time between sequential requests in seconds (Default: 60).         |


## Description

The file name of the external program is specified in the environment variable `EXEC_PATH`.

When it is run by lego, three command-line parameters are passed to it:
The action ("present" or "cleanup"), the fully-qualified domain name and the value for the record.

For example, requesting a certificate for the domain 'my.example.org' can be achieved by calling lego as follows:

```bash
EXEC_PATH=./update-dns.sh \
lego --email you@example.com --dns exec --d my.example.org run
```

It will then call the program './update-dns.sh' with like this:

```bash
./update-dns.sh "present" "_acme-challenge.my.example.org." "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI"
```

The program then needs to make sure the record is inserted.
When it returns an error via a non-zero exit code, lego aborts.

When the record is to be removed again,
the program is called with the first command-line parameter set to `cleanup` instead of `present`.

If you want to use the raw domain, token, and keyAuth values with your program, you can set `EXEC_MODE=RAW`:

```bash
EXEC_MODE=RAW \
EXEC_PATH=./update-dns.sh \
lego --email you@example.com --dns exec -d my.example.org run
```

It will then call the program `./update-dns.sh` like this:

```bash
./update-dns.sh "present" "--" "my.example.org." "some-token" "KxAy-J3NwUmg9ZQuM-gP_Mq1nStaYSaP9tYQs5_-YsE.ksT-qywTd8058G-SHHWA3RAN72Pr0yWtPYmmY5UBpQ8"
```

## Commands

-> **NOTE**: The `--` is because the token MAY start with a `-`, and the called program may try and interpret a `-` as indicating a flag.
In the case of urfave, which is commonly used,
you can use the `--` delimiter to specify the start of positional arguments, and handle such a string safely.

### Present

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram present <FQDN> <record>`                |
| `RAW`   | `myprogram present -- <domain> <token> <key_auth>` |

### Cleanup

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram cleanup <FQDN> <record>`                |
| `RAW`   | `myprogram cleanup -- <domain> <token> <key_auth>` |


