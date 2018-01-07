## v0.5.0 (Unreleased)

Most of the items in this release are the result of a refresh of lego, which
brings the following new features, amongst others:

 * DNSimple API now supports V2.
 * You can now supply `AWS_HOSTED_ZONE_ID` to the route53 DNS challenge to
   directly specify the zone ID for the DNS challenge, instead of getting the
   provider to try and detect it.
 * New DNS challenge providers: `azure`, `auroradns`, `dnspod`, `exoscale`,
   `godaddy`, `linode`, `rackspace`, `ns1`, and `otc`.

## v0.4.0

### General Information

 * Releases are no longer signed. SHA256SUMS are still published, however, and
   signing may come back under a more general signing key. Keep this in mind if
   you need earlier releases as well.
 * Built against Terraform v0.10.0-beta2 with the [custom diff
   patch](https://github.com/hashicorp/terraform/pull/14887). Although the
   plugin API version has not yet changed, YMMV with using this on Terraform
   versions below v0.10.0-beta2. See below for details on why we are using the
   custom diff patch.

### New Diff Behaviour for Certificate Renewals

The correctness of the certificate renewal behaviour in this resource has been a
long-running problem, due to the fact that certificates were renewed during the
refresh cycle. This caused silent updates and empty diffs unless you had
resources in the same stack that depended on the certificates. In addition to
this, this has led to issues with implementing settings like
`min_days_remaining` in a way that made its setting effective on the present run
without `ForceNew`. These issues are articulated in #13 and #15.

As of this version, these issues are no longer a problem. Using the
aforementioned custom diff patch, the certificate's expiry is now checked during
the diff phase of a `terraform plan`, articulated below:

```
The Terraform execution plan has been generated and is shown below.
Resources are shown in alphabetical order for quick scanning. Green resources
will be created (or destroyed and then created if an existing resource
exists), yellow resources are being changed in-place, and red resources
will be destroyed. Cyan entries are data sources to be read.

Note: You didn't specify an "-out" parameter to save this plan, so when
"apply" is called, Terraform can't guarantee this is what will execute.

  ~ acme_certificate.certificate
      certificate_pem: "-----BEGIN CERTIFICATE-----\nxxxxxxx\n-----END CERTIFICATE-----\n" => "<computed>"
```

If the certificate requires renewal, `certificate_pem` is set to `<computed>`
and correctly renewed during the next `terraform apply` run.

This also means that setting `min_days_remaining` no longer forces a new
resource and also works immediately - if you adjust it, its settings will work
during your next plan.

## v0.3.0

Fully updated version, supporting v0.9.0. Make sure you use this version for the
full v0.9.0 release, as v0.3.0-beta2 will not work (the plugin API version has
been incremented again). People still on versions of TF before v0.9.0 should use
a v0.2.x version.

## v0.3.0-beta2

This beta version tracks Terraform `v0.9.0`, which as of this writing (Feb 28th,
2017) is currently in beta. All that has changed so far on this side is that we
need to rebuild as the plugin API has again changed.

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

