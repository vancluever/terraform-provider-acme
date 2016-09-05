Terraform ACME Provider
========================

This repository contains a plugin form of the ACME provider that was proposed
and submitted in [Terraform PR #7058][1].

## Why a Plugin?

As documented in the thread in the PR, making use of [`xenolf/lego`][2], while
being very helpful in creating this feature, nontheless created some coupling
of the authorization and certificate generation process that (rightfully) did
not sit well with the core TF team. Further to that, as the authorization
process allows the use of DNS provdiers that Terraform has support for, the
potential to cross provider boundaries exists.

Hence it was determined that it might not be the best fit for upstream for the
time being.

Nonetheless, the code in this repository is still completely functional and
there have been a few people that have found it useful. An external plugin was
requested, so here it is. :)

## Installing

See the [Plugin Basics][3] page of the Terraform docs to see how to plunk this
into your config. Check the [releases page][4] of this repo to get releases for
Linux, OS X, and Windows.

## Usage

The documentation from the old PR has been preserved and you can look [here][5]
to find it.

[1]: https://github.com/hashicorp/terraform/pull/7058
[2]: https://github.com/xenolf/lego
[3]: https://www.terraform.io/docs/plugins/basics.html
[4]: https://github.com/paybyphone/terraform-provider-acme/releases
[5]: website/source/docs/providers/acme
