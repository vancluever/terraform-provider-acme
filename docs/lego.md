# Submitting a lego Issue

This provider integrates [lego](https://github.com/go-acme/lego) and ultimately
uses its exported primitives in a Terraform context. A large number of the
issues submitted to the Terraform ACME provider are actually issues related to
lego, particularly its DNS providers.

Before you submit an issue, especially if it is DNS provider related, please
attempt to reproduce the issue with lego's CLI tool. DNS provider-related issues
that are not very specifically related to how this provider interacts with lego
will more than likely be closed with a referral
back to lego.

## Installing and Using lego

You can view lego's installation instructions at
https://go-acme.github.io/lego/installation/. See
https://go-acme.github.io/lego/usage/cli/ for instructions on how to use the CLI
tool itself.

## Submitting an Issue to lego

[lego's issue tracker](https://github.com/go-acme/lego/issues)

You should also read their [contributing
guidelines](https://github.com/go-acme/lego/blob/master/CONTRIBUTING.md) before
submitting an issue.

Remember that the lego maintainers are not responsible for this provider, as
such it's important (and considerate!) that questions are framed in a
lego-specific context.
