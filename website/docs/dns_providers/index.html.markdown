---
layout: "acme"
page_title: "Provider: ACME"
sidebar_current: "docs-acme-dns-providers"
description: |-
  Describes the DNS challenge providers that can be used with the ACME provider.
---

# acme_certificate DNS Challenge Providers

This subsection documents all of the DNS challenge providers that can be used
with the [`acme_certificate`][resource-acme-certificate] resource.

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html

For complete information on how to use these providers with the
`acme_certifiate` resource, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

Refer to a specific provider on the left sidebar for more details.

## Using Variable Files for Provider Arguments

Most provider arguments can be suffixed with `_FILE` to specify that you wish to
store that value in a local file. This can be useful if local storage for these
values is desired over configuration as variables or within the environment.

See the [example][acme-certificate-file-arg-example] in the `acme_certificate`
resource for more details.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

## Relation to Terraform provider configuration

The DNS provider configurations specified in the
[`acme_certificate`][resource-acme-certificate] resource are separate from any
that you supply in a corresponding provider whose functionality overlaps with
the certificate's DNS providers.  This ensures that there are no hard
dependencies between any of these providers and the ACME provider, but it is
important to note so that configuration is supplied correctly.

As an example, if you specify manual configuration for the [AWS
provider][tf-provider-aws] via the [`provider`][tf-providers] block instead of
the environment, you will still need to supply the configuration explicitly in
the `config` block of the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument.

[tf-provider-aws]: /docs/providers/aws/index.html
[tf-providers]: /docs/configuration/providers.html
[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

Note that some of Terraform's providers have environment variable settings that
overlap with the settings here, generally depending on whether or not these
variables are supported by the corresponding provider's SDK.

We alias certain provider environment variables so the same settings can be
supplied to both ACME and the respective native cloud provider. For specific
details, see the page for the provider in question.
