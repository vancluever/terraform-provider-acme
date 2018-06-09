test:
	go test -v . ./plugin/providers/acme 

testacc:
	TF_ACC=1 go test -v ./plugin/providers/acme $(TESTARGS) -timeout 20m

debugacc:
	TF_ACC=1 dlv test ./plugin/providers/acme -- -test.v $(TESTARGS) -test.timeout 20m

build: deps
	gox -osarch="linux/amd64 windows/amd64 darwin/amd64" \
	-output="pkg/{{.OS}}_{{.Arch}}/terraform-provider-acme" .

deps:
	go get -u github.com/mitchellh/gox

clean:
	rm -rf pkg/
