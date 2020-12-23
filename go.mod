module github.com/terraform-providers/terraform-provider-acme

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/davecgh/go-spew v1.1.1
	github.com/go-acme/lego/v3 v3.1.0
	github.com/go-acme/lego/v4 v4.1.3
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/terraform v0.14.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.3
	github.com/miekg/dns v1.1.31
	github.com/mitchellh/copystructure v1.0.0
	github.com/oklog/run v1.1.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	software.sslmate.com/src/go-pkcs12 v0.0.0-20190209200317-47dd539968c4
)

replace github.com/terraform-providers/terraform-provider-tls/v3 => github.com/vancluever/terraform-provider-tls/v3 v3.0.1
