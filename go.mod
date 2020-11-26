module github.com/terraform-providers/terraform-provider-acme

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Unknwon/com v0.0.0-20151008135407-28b053d5a292 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dnaeon/go-vcr v0.0.0-20180920040454-5637cf3d8a31 // indirect
	github.com/go-acme/lego/v3 v3.1.0
	github.com/go-acme/lego/v4 v4.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl2 v0.0.0-20190725010614-0c3fe388e450 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform v0.13.5
	github.com/hashicorp/terraform-plugin-sdk v1.16.0
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/marstr/guid v1.1.0 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/miekg/dns v1.1.31
	github.com/mitchellh/copystructure v1.0.0
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/terraform-providers/terraform-provider-openstack v1.15.0 // indirect
	github.com/terraform-providers/terraform-provider-tls/v2 v2.2.1
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/build v0.0.0-20190111050920-041ab4dc3f9d // indirect
	k8s.io/client-go v11.0.0+incompatible // indirect
	software.sslmate.com/src/go-pkcs12 v0.0.0-20190209200317-47dd539968c4
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.3.0+incompatible
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.14
	github.com/ldez/go-auroradns => github.com/ldez/go-auroradns/v2 v2.0.2
	github.com/terraform-providers/terraform-provider-tls/v2 => github.com/vancluever/terraform-provider-tls/v2 v2.2.1
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
