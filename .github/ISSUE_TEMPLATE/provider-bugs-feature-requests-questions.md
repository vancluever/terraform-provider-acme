---
name: Provider bugs/feature requests/questions
about: Submit a provider-related issue (for DNS provider issues, see "Submitting a lego Issue").
title: ''
labels: ''
assignees: ''

---

**Please delete this message before you submit your issue.**

Thanks for your interest in using the Terraform ACME provider!

Please read and consider the following before submitting your report.

This provider integrates [lego](https://github.com/go-acme/lego) and ultimately uses its exported primitives in a Terraform context. A large number of the issues submitted to this repository are actually issues related to lego, particularly its DNS providers.

Before you submit an issue:

* Search the issue tracker here to see if your problem has been answered (even for lego-related issues; there is plenty of history and options that have been added to work around various DNS issues that neither we nor lego can fix).
* If your problem is not assuredly related to this provider itself (such as in the case with DNS provider issues), please attempt to reproduce the issue with [lego's CLI tool](https://go-acme.github.io/lego/installation/).

DNS provider-related issues that are not specifically related to how this provider interacts with lego will more than likely be closed with a referral back to lego. Please read our [instructions](https://github.com/vancluever/terraform-provider-acme/blob/master/docs/lego.md) on submitting a lego issue if you need to.

Please also understand that there are other issues that we will not be able to really troubleshoot for you or reliably provide support on:

* Transient network issues: Please check your network configuration, contact the CA you are working with, and/or contact your service provider(s).
* General Terraform configuration issues: Please refer to the [Terraform community page](https://www.terraform.io/community) for links on where best to direct your request.

Thanks!
