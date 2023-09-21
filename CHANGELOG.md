## 2.18.1-pre (Unreleased)

Bumped version for dev.

## 2.18.0 (September 21, 2023)

FEATURES:

* `resource/acme_certificate`: Added the `cert_timeout` option to control the
  timeout of HTTP requests used to obtain the certificate after challenges are
  complete.
  [#349](https://github.com/vancluever/terraform-provider-acme/issues/349) 

## 2.17.2 (September 20, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.14.2 See the
lego [CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.14.2/CHANGELOG.md)
for more details on additions and changes to DNS providers, and other minor
changes to the library.

## 2.17.1 (September 11, 2023)

FEATURES:

* `resource/acme_certificate`: `azuredns` DNS provider now has the same
  environment variable aliasing as the old `azure` provider.
  [#342](https://github.com/vancluever/terraform-provider-acme/issues/342) 

## 2.17.0 (August 22, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.14.0 See the
lego [CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.14.0/CHANGELOG.md)
for more details on additions and changes to DNS providers, and other minor
changes to the library.

FEATURES:

* `resource/acme_certificate`: New HTTP challenge type `http_s3_challenge`,
  which will allow publishing HTTP challenge records to an S3 bucket.
  [#330](https://github.com/vancluever/terraform-provider-acme/pull/330)

## 2.16.1 (August 11, 2023)

This change is being made to correct build issues. No other changes are being
made.

## 2.16.0 (August 11, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.13.3 See the
lego [CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.13.3/CHANGELOG.md)
for more details on additions and changes to DNS providers, and other minor
changes to the library.

## 2.15.1 (June 19, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.12.2 See the
lego [CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.12.2/CHANGELOG.md)
for more details on additions and changes to DNS providers, and other minor
changes to the library.

This lego update contains a critical fix to the `dnsmadeeasy` DNS provider,
ensuring that it does not delete all records in a zone during cleanup.
[go-acme/lego#1939](https://github.com/go-acme/lego/pull/1939)

## 2.15.0 (June 14, 2023)

This update is a full release of [2.15.0-beta1](#2150-beta1-june-13-2023),
including the lego update and the bug fix for the `recursive_nameservers`
setting. See that release for more details.

## 2.15.0-beta1 (June 13, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.12.1 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.12.1/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

BUG FIXES:

* `resource/acme_certificate`: DNS plugins should now respect the setting of
  `recursive_nameservers` again.
  [#316](https://github.com/vancluever/terraform-provider-acme/pull/316)

## 2.14.0 (May 5, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.11.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.11.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.13.1 (February 23, 2023)

This is a security fix to address
[CVE-2022-41723](https://github.com/advisories/GHSA-vvpx-j8f3-3w6h). No other
changes have been made.

## 2.13.0 (February 18, 2023)

This update is a full release of [2.13.0-beta1](#2130-beta1-february-5-2023) and
[2.13.0-beta2](#2130-beta2-february-11-2023), including the new DNS provider
plugin system, and lego v4.10.0. See those releases for more details.

BUG FIXES:

* `resource/acme_certificate`: New DNS providers and documentation for v4.10.0
  should now be properly generated.

## 2.13.0-beta2 (February 11, 2023)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.10.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.10.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

BUG FIXES:

* `resource/acme_certificate`: The new DNS plugin system now implements
  challenge provider timeouts properly.
  [#277](https://github.com/vancluever/terraform-provider-acme/pull/277)

## 2.13.0-beta1 (February 5, 2023)

NEW DNS PLUGIN SYSTEM

This update moves DNS providers for the `acme_certificate` resource into a
[go-plugin](https://github.com/hashicorp/go-plugin) backed sub-plugin built into
the provider. One provider is executed for each instance of a DNS provider
supplied in `acme_certificate`, each with its own environment. This fixes a
long-running issue where the environment variables set in the `config` parameter
of one provider in one resource would overwrite the settings of another resource
with different config settings for the same DNS provider. See
[#235](https://github.com/vancluever/terraform-provider-acme/issues/235) and
[#276](https://github.com/vancluever/terraform-provider-acme/pull/276) for more
details.

## 2.12.0 (Dec 21, 2022)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.9.1 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.9.1/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.11.1 (Oct 8, 2022)

This update is a bugfix to correct the fact that the provider was not fully
synced with the lego updates when 2.11.0 was released. No other changes were
made.

## 2.11.0 (Oct 8, 2022)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.9.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.9.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

FEATURES:

* `resource/acme_certificate`: Added the `certificate_not_after` attribute to
  show the certificate expiry date in state.
  [#264](https://github.com/vancluever/terraform-provider-acme/pull/264)

## 2.10.0 (Jul 5, 2022)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.8.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.8.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.9.0 (May 30, 2022)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.7.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.7.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.8.0 (January 24, 2022)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.6.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.6.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.7.1 (December 10, 2021)

This is a patch version bump to build with the latest version of Go to address
CVE-2021-44716.

Note that this update is being made proactively and not in response to any known
security issue. The ACME provider would normally only use `net/http` in HTTP
challenges.

No other changes are being made.

## 2.7.0 (November 15, 2021)

FEATURES:

* `resource/acme_certificate`: New flag `revoke_certificate_on_destroy` to
  control if certificates are revoked on destroy. Default is `true`, keeping
  with existing behaviour.
  [#192](https://github.com/vancluever/terraform-provider-acme/issues/192) 

## 2.6.0 (October 6, 2021)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.5.3 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.5.3/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.5.3 (August 10, 2021)

BUG FIXES:

* `resource/acme_certificate`: Corrected an issue where `preferred_chain` was
  not working for certificates that used an external CSR.
  [#199](https://github.com/vancluever/terraform-provider-acme/issues/199)

## 2.5.2 (June 9, 2021)

This is another re-issue of v2.5.0 due to a goreleaser config issue. No other
changes have been made.

## 2.5.1 (June 9, 2021)

This is a re-issue of 2.5.0 without freebsd/arm support, which has been
suspended due to build issues for the time being. No other changes have been
made.

## 2.5.0 (June 9, 2021)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.4.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.4.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.4.0 (April 4, 2021)

FEATURES:

* `resource/acme_certificate`: The resource now supports HTTP-01 and
  TLS-ALPN-01 challenges again. These are done through the `http_challenge`,
  `http_webroot_challenge`, `http_memcached_challenge`, and `tls_challenge`
  challenge types. It is still recommended that you use DNS challenges whenever
  possible. See the documentation for more details.
  [#169](https://github.com/vancluever/terraform-provider-acme/pull/169)

## 2.3.0 (March 19, 2021)

FEATURES:

* `resource/acme_certificate`: Added the `preferred_chain` option to allow for
  the selection of alternate certificate chains offered by the CA.
  [#161](https://github.com/vancluever/terraform-provider-acme/issues/161)

## 2.2.0 (March 12, 2021)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.3.1 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.3.1/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.1.2 (February 22, 2021)

BUG FIXES:

* `resource/acme_certificate`: The resource no longer always expects two
  certificates (ie: a single intermediate certificate). All intermediate
  certificates are now concatenated in `issuer_pem`. The `certificate_p12`
  should contain all issuer certificates as well.
  [#154](https://github.com/vancluever/terraform-provider-acme/issues/154)

## 2.1.1 (February 11, 2021)

This is a simple version bump to fix documentation on the Terraform Registry. No
changes are being made.

## 2.1.0 (February 10, 2021)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.2.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/16660b2/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

## 2.0.0 (December 23, 2020)

BREAKING CHANGES:

* `resource/acme_certificate`: The resource ID is now a state-local UUID, not
  the same as `certificate_url`. This is to prevent drift issues during renewal.
  If you need the URL for the current version of the certificate, use the
  `certificate_url` field.
  [#103](https://github.com/vancluever/terraform-provider-acme/issues/103)

FEATURES:

* `resource/acme_certificate`: Added the `pre_check_delay` option to allow for
  the insertion of delays in DNS challenges. This should help with DNS
  propagation issues with certain providers.
  [#111](https://github.com/vancluever/terraform-provider-acme/pull/111)
* `resource/acme_certificate`: The domain defined in the `common_name` field can
  now be specified in `subject_alternative_names`. This is a strictly semantic
  change as the CN is already included in the SAN list of issued certificates.
  [#90](https://github.com/vancluever/terraform-provider-acme/issues/90)

## 1.6.3 (November 30, 2020)

This is (yet another) simple version bump to attempt to fix documentation on the
Terraform Registry. No changes are being made.

## 1.6.2 (November 30, 2020)

This is (another) simple version bump to attempt to fix documentation on the
Terraform Registry. No changes are being made.

## 1.6.1 (November 27, 2020)

This is a simple version bump to attempt to fix documentation on the Terraform
Registry. No changes are being made.

## 1.6.0 (November 27, 2020)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v4.1.3 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v4.1.3/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library.

GENERAL NOTIFICATIONS:

* Testing of the provider has moved to use
  [pebble](https://github.com/letsencrypt/pebble/) exclusively. Tests for some
  features that are not explicitly supported by pebble or were otherwise tested
  manually have been removed. See
  [`907de66`](https://github.com/vancluever/terraform-provider-acme/commit/907de66625886fbd86b383cb158515ef458f3604)
  for more details.
* Support for Terraform 0.11 has been dropped. The provider is now only
  available on the Terraform registry.

FEATURES:

* `resource/acme_registration`: Added support for external account binding. This
  allows registrations to be linked to external accounts, commonly used by
  commercial CAs.
* `resource/acme_certificate`: Added the `disable_complete_propagation` option,
  which allows one to disable the propagation pre-check before attempting to
  complete the DNS challenge. Enabling this is only recommended for testing.

## 1.5.0 (October 21, 2019)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v3.1.0 See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v3.1.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library. ([#92](https://github.com/terraform-providers/terraform-provider-acme/issues/92))

## 1.4.0 (August 20, 2019)

LEGO UPDATE:

[lego](https://github.com/go-acme/lego) has been updated to v3.0.0 (from v2.5.0
in provider version 1.2.0). See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/v3.0.0/CHANGELOG.md) for
more details on additions and changes to DNS providers, and other minor changes
to the library. ([#82](https://github.com/terraform-providers/terraform-provider-acme/issues/82))

Starting with this release, all DNS providers and documentation for the DNS
providers will be auto-generated, so the lego CHANGELOG will be the place to
look for lego-specific changes in the future.

BUG FIXES:

* `resource/acme_certificate`: When renewing certificate, private keys and CSRs
  will now only be set if they are present in the state. This may correct some
  library-related issues during the renewal process. ([#84](https://github.com/terraform-providers/terraform-provider-acme/issues/84))
* `resource/acme_registration`: Unknown or deactivated registrations will now be
  removed from state on refresh. ([#85](https://github.com/terraform-providers/terraform-provider-acme/issues/85))

## 1.3.5 (July 30, 2019)

BUG FIXES:

* `resource/acme_certificate`: Expired certificates flagged for renewal will now
  properly renew on the next `terraform apply` on Terraform 0.12.0 and higher.
  ([#77](https://github.com/terraform-providers/terraform-provider-acme/issues/77))

## 1.3.4 (June 06, 2019)

BUG FIXES:

* `resource/acme_certificate`: All computed attributes associated with a
  certificate are now marked for re-generation when a certificate needs to be
  renewed. While this was previously happening in reality, this was not being
  reflected in the plan. ([#64](https://github.com/terraform-providers/terraform-provider-acme/issues/64))

## 1.3.3 (May 29, 2019)

This update is a Terraform 0.12 support build for the changes from 1.3.2.

## 1.3.2 (May 28, 2019)

BUG FIXES:

* `resource/acme_certificate`: Corrected an issue where `certificate_pem` would
  be deleted from state on certificate renew failure. ([#60](https://github.com/terraform-providers/terraform-provider-acme/issues/60))
* `resource/acme_certificate`: The resource will now attempt to recover the
  `certificate_pem` field from the CA if it is missing in state. ([#59](https://github.com/terraform-providers/terraform-provider-acme/issues/59))

## 1.3.1 (May 23, 2019)

IMPROVEMENTS:

* The provider now will log lego's log messages when `TF_LOG=debug` or higher is
  set. ([#53](https://github.com/terraform-providers/terraform-provider-acme/issues/53))

BUG FIXES:

* `resource/acme_certificate`: Corrected state migration issues that were
  causing the resource to not function at all in Terraform 0.12. ([#57](https://github.com/terraform-providers/terraform-provider-acme/issues/57))
* `resource/acme_certificate`: Corrected state migration issues that may have
  triggered an update of settings due to incorrect migration of the
  `recursive_nameservers` attribute on Terraform 0.11. ([#55](https://github.com/terraform-providers/terraform-provider-acme/issues/55))

## 1.3.0 (May 17, 2019)

BREAKING CHANGES: 

* `resource/acme_certificate`: The `recursive_nameservers` option for checking
  propagation of DNS challenge records has been promoted to a top-level option
  and is no longer provided as part of an individual `dns_challenge` block.
  ([#49](https://github.com/terraform-providers/terraform-provider-acme/issues/49))

FEATURES:

* `resource/acme_certificate`: This resource now supports multiple DNS
  challenges for working with multiple primary DNS providers. ([#49](https://github.com/terraform-providers/terraform-provider-acme/issues/49))

## 1.2.1 (May 14, 2019)

FEATURES:

* The plugin has been updated to support Terraform 0.12 and higher. Backwards
  compatibility has been maintained to Terraform 0.11.x and earlier versions
  that support plugin protocol version 4. ([#45](https://github.com/terraform-providers/terraform-provider-acme/issues/45))

## 1.2.0 (May 14, 2019)

LEGO UPDATE AND NEW DNS PROVIDERS:

[lego](https://github.com/go-acme/lego) has been updated to v2.5.0. See the lego
[CHANGELOG.md](https://github.com/go-acme/lego/blob/master/CHANGELOG.md#v250---2019-04-17) for more details. ([#47](https://github.com/terraform-providers/terraform-provider-acme/issues/47))

The update brings the following new DNS providers:

* `cloudns`
* `dode`
* `oraclecloud`

IMPROVEMENTS:

* `resource/acme_certificate`: The default `min_days_remaining` is now set to 30
  days, up from 7. ([#48](https://github.com/terraform-providers/terraform-provider-acme/issues/48))

## 1.1.2 (May 06, 2019)

BUG FIXES:

* `resource/acme_certificate`: Revocation on destroy now skips expired
  certificates. ([#42](https://github.com/terraform-providers/terraform-provider-acme/issues/42))

## 1.1.1 (March 12, 2019)

BUG FIXES:

* `resource/acme_certificate`: Added the optional `certificate_p12_password`
  field, used when creating the PFX bundle found in `certificate_p12`. ([#35](https://github.com/terraform-providers/terraform-provider-acme/issues/35))
* `resource/acme_certificate`: `certificate_p12` base64 data is now padded and
  should be usable by Azure services that take PKCS12 data. ([#34](https://github.com/terraform-providers/terraform-provider-acme/issues/34))

## 1.1.0 (March 01, 2019)

LEGO UPDATE AND NEW DNS PROVIDERS:

[lego](https://github.com/go-acme/lego) has been updated to v2.2.0.

As part of this update, a number of new DNS providers have been added for
`acme_certificate`:

* `acmedns`
* `alidns`
* `conoha`
* `designate`
* `dreamhost`
* `hostingde`
* `httpreq`
* `iij`
* `inwx`
* `linodev4`
* `mydnsjp`
* `netcup`
* `nifcloud`
* `sakuracloud`
* `selectel`
* `stackpath`
* `transip`
* `vegadns`
* `vscale`
* `zoneee`

Thanks very much to @yamamoto-febc ([#10](https://github.com/terraform-providers/terraform-provider-acme/issues/10)) and @bzub ([#17](https://github.com/terraform-providers/terraform-provider-acme/issues/17)), ([#18](https://github.com/terraform-providers/terraform-provider-acme/issues/18)) for the
help with documentation, code updates, and module migration work!

IMPROVEMENTS:

* `resource/acme_certificate`: Added the `recursive_nameservers` attribute to
  the `dns_challenge` block. This allows someone to specify a static resolver
  list for DNS propagation checks that will override the resolvers of the system
  running Terraform. This can be useful when dealing with split horizon DNS
  scenarios. ([#25](https://github.com/terraform-providers/terraform-provider-acme/issues/25))
* `resource/acme_certificate`: Added the `certificate_p12` output, which makes
  the certificate, intermediate CA, and private key available in a PFX PKCS12
  archive. This can be useful when working with Microsoft products.  ([#26](https://github.com/terraform-providers/terraform-provider-acme/issues/26))

BUG FIXES:

* `resource/acme_certificate`: Modifications to the `dns_challenge`
  configuration will now persist across no-op updates. Additionally,
  modification of these values will no longer force a new resource. ([#28](https://github.com/terraform-providers/terraform-provider-acme/issues/28))

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
