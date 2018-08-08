## 1.0.2 (Unreleased)
## 1.0.1 (August 08, 2018)

This is release bump for the sole purpose of releasing the provider upstream. As
of this release, you will be able to fetch this project directly via Terraform!

## 1.0.0 (Jun 17, 2018)

BREAKING CHANGES:

* The provider has now been updated for ACME v2 and will no longer work for ACME
  v1. If you require v1, use version 0.6.0 of the provider.
* Existing states for `acme_registration` and `acme_certificate` will be
  preserved on update and there should be no need to re-create either
  registrations or certificates, so long as the CA supports it. Let's Encrypt
  supports these updates.
* Several fields have been removed and the resource relationships have changed.
  For full details, see the documentation.
* `server_url` is now a provider-level configuration value. The documentation
  has several full examples of this in action.
* `resource/acme_certificate`: The `http_challenge_port` and
  `tls_challenge_port` parameters have been removed. The resource now only
  supports DNS challenges, so `dns_challenge` is now a required field.
  [#40][gh-40]

IMPROVEMENTS:

* `resource/acme_certificate`: With the update to ACME v2, this resource now
  supports wildcard certificates.
* `resource/acme_registration`: This resource will now completely remove a
  registration from the ACME server when the resource is destroyed. [#39][gh-39]

BUG FIXES:

* `resource/acme_certificate`: The post-revocation OCSP validation has been
  completely removed. This should make destruction of the resource much more
  reliable. [#41][gh-41]

[gh-41]: https://github.com/vancluever/terraform-provider-acme/pull/41
[gh-40]: https://github.com/vancluever/terraform-provider-acme/pull/40
[gh-39]: https://github.com/vancluever/terraform-provider-acme/pull/39

## 0.6.0

**NOTE:** This is the last major release before 1.0.0, which will include
support for ACME v2 and will more than likely break support for ACME v1. If you
require ACME v1 after 1.0.0, use this version of the provider.

NEW DNS PROVIDERS:

The `acme_certificate` resource has had a provider refresh, with the following
new providers added:

* `bluecat`
* `cloudxns`
* `duckdns`
* `fastdns`
* `gandiv5`
* `glesys`
* `lightsail`
* `namedotcom`
* `exec`

These providers, and previous providers, have been synchronized with their state
at lego version [v0.5.0][lego-dns-providers-v0.5.0].

[lego-dns-providers-v0.5.0]: https://github.com/xenolf/lego/tree/v0.5.0/providers/dns

IMPROVEMENTS:

* `resource/acme_certificate`: This resource now supports supplying the `delete`
  [resource timeout][resource-timeouts] timeout, which controls the certificate
  revocation timeout (or more specifically, the OCSP wait timeout).
  ([#32][gh-32])
* `resource/acme_certificate`: Added alias mappings for the Azure DNS provider's
  environment variables so that the same environment variables for the
  [Terraform Azure Provider][tf-azurerm-provider] can be used with the ACME
  plugin. ([#36][gh-36])
* `resource/acme_certificate`: Already revoked certificates are ignored by the
  destroy process, ensuring that they are destroyed without error in Terraform.
  ([#33][gh-33])
* `resource/acme_certificate`: The `config` field of `dns_challenge` has now
  been marked as a sensitive field to prevent credentials from being leaked in
  output. ([#31][gh-31])

[resource-timeouts]: https://www.terraform.io/docs/configuration/resources.html#timeouts
[tf-azurerm-provider]: https://www.terraform.io/docs/providers/azurerm/index.html
[gh-36]: https://github.com/vancluever/terraform-provider-acme/pull/36
[gh-33]: https://github.com/vancluever/terraform-provider-acme/pull/33
[gh-32]: https://github.com/vancluever/terraform-provider-acme/pull/32
[gh-31]: https://github.com/vancluever/terraform-provider-acme/pull/31

## 0.5.0

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
      certificate_pem: "-----BEGIN CERTIFICATE-----
xxxxxxx
-----END CERTIFICATE-----
" => "<computed>"
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
