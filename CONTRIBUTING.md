# Terraform ACME Provider Public Contribution Notice

> [!NOTE]
> **The short of it:** This project does not (any longer) accept public pull
> requests. There are a few reasons for this, so read on if you want some
> reasoning, or just feel free to [open an
> issue](https://github.com/vancluever/terraform-provider-acme/issues/new/choose)
> to get your specific problem or feature request looked at.

The Terraform ACME provider is now closed to public pull requests. Any pull
request opened will be closed using the
[Vouch](https://github.com/mitchellh/vouch) workflow, with a message directing
you to this document to explain why.

Note that unlike other projects that might be implementing Vouch, it is
unlikely I will be opening this project up to external contributors. If you
have a feature request or bug, please open an issue to get further help.

**Please do not use AI to open issues**. If you are using AI or an agent to
understand an issue, take the time to comprehend the output and explain the
issue **in your own words** before submitting.

## Why?

The Terraform ACME provider is a relatively small project that does not require
much maintenance work to keep functional and up to date. While I have
considered it to be largely feature complete for a long time, I am still
interested in adding features and functionality if they make sense within the
design of the provider.

Performing this work myself allows me to continue to learn about new Terraform
features, developments in the ACME specification and our upstream use of
[lego](https://github.com/go-acme/lego), and steward the provider in a way that
will not burn me out as a maintainer.

The open source landscape has also changed significantly since this project has
started, particularly recently due to the advent of AI coding tools and agents.
Vouch has a [great section explaining
this](https://github.com/mitchellh/vouch/blob/f0591095f3c46406301604874d2482797fab7bab/README.md#why).
To add to this, on a personal level, I am not very interested in even having to
engage with content created by AI coding agents and/or other tools (see last
section). **This is not a blanket judgment of the work produced by such tools,
or those that use them.** To each their own.

## Open source, not open contribution

This project is open source and distributed under the terms of the [Mozilla
Public License](https://www.mozilla.org/en-US/MPL/2.0/). Open source, however
does not necessarily mean open contribution, a point that seems to have been in
discourse a lot recently due to the points made in the last paragraph of the
previous section.

Probably the best example of a closed-contribution project is SQLite, which has
some great rationale and statements to this end on their [copyright
page](https://www.sqlite.org/copyright.html). Echoing the sentiments
particularly in the last part of that page, feature requests and proof of
concepts are welcome, however I will more than likely implement all solutions
to these personally. 
