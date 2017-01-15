## v0.2.1

This is a bugfix to correct #6 and ensure that TF will abort if a DNS challenge
is improperly configured (example: missing credentials). Previous to this
release if the DNS challenge could not be properly set up, the plugin would have
proceeded with an HTTP or TLS challenge.

## v0.2.0

Note that this release is built for Terraform v0.8.0 and higher - using with
v0.7.x and lower may not work. Use the v0.1.0 release instead.

 * Added the `must_staple` option - this option adds the [OCSP Stapling
   Required][1] extension to created certificates, ensuring that a valid OCSP
   Staple must be included in the TLS handshake for the connection to proceed.
   This is disabled by default. This option has no effect when being used with
   external CSRs.

[1]: https://letsencrypt.org/docs/integration-guide/#implement-ocsp-stapling

## v0.1.0

Initial release.

