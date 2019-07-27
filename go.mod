module github.com/terraform-providers/terraform-provider-acme

go 1.12

require (
	github.com/Azure/go-autorest v12.3.0+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/go-acme/lego v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/terraform v0.12.6-0.20190729234158-cd145828c099
	github.com/miekg/dns v1.1.15
	github.com/mitchellh/copystructure v1.0.0
	github.com/terraform-providers/terraform-provider-tls v1.2.0
	software.sslmate.com/src/go-pkcs12 v0.0.0-20190322163127-6e380ad96778
)

replace github.com/go-acme/lego => github.com/vancluever/lego v0.0.0-20190729190929-f9d873e7817e

replace github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.15
