module github.com/terraform-providers/terraform-provider-acme

go 1.15

require (
	github.com/Azure/azure-sdk-for-go v45.0.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.3 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.0 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/aws/aws-sdk-go v1.31.9 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/go-acme/lego/v3 v3.1.0
	github.com/go-acme/lego/v4 v4.1.3
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/gophercloud/gophercloud v0.10.1-0.20200424014253-c3bfe50899e5 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/hcl/v2 v2.6.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.3
	github.com/miekg/dns v1.1.31
	github.com/mitchellh/copystructure v1.0.0
	github.com/oklog/run v1.1.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/zclconf/go-cty v1.5.1 // indirect
	software.sslmate.com/src/go-pkcs12 v0.0.0-20190209200317-47dd539968c4
)

replace github.com/terraform-providers/terraform-provider-tls/v3 => github.com/vancluever/terraform-provider-tls/v3 v3.0.1
