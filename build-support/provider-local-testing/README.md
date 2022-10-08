# Support Files for Local Testing

This directory contains support files for local provider testing.

Basically, if you are building the provider pre-release and want to test it
locally, this is the directory for you!

## Building the provider

Run `make build-pre-release` to install the binary in
`/tmp/terraform-provider-acme`.

## Provider override configuration file

You can then run `export TF_CLI_CONFIG_FILE=$(pwd)/dev.tfrc` (while in this
directory) to use the `dev.tfrc` contained here.

## Cleaning up

You can run `make clean-pre-release` to delete the
`/tmp/terraform-provider-acme` directory, or just remove it on your own.