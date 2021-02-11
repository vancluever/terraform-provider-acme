# Test all packages by default
TEST ?= ./...

.PHONY: default
default: build

.PHONY: tools
tools:
	cd $(shell go env GOROOT) && go get -u github.com/hashicorp/go-bindata/go-bindata gotest.tools/gotestsum

.PHONY: pebble-start-install
pebble-start-install: pebble-stop
	build-support/scripts/pebble-start.sh --install

.PHONY: pebble-start
pebble-start: pebble-stop
	build-support/scripts/pebble-start.sh

.PHONY: pebble-stop
pebble-stop:
	build-support/scripts/pebble-stop.sh

.PHONY: template-generate
template-generate:
	@echo "==> Re-generating templates..."
	@go generate ./build-support/generate-dns-providers

.PHONY: provider-generate
provider-generate:
	@echo "==> Re-generating Go DNS provider factory in ./acme..."
	@go generate ./acme
	@go build ./acme
	@go mod tidy
	@echo "==> Re-genrating documentation..."
	@rm -f docs/guides/dns-providers-*.md
	@go run ./build-support/generate-dns-providers doc docs/guides/

.PHONY: provider-generate-update
provider-generate-update: provider-generate
	test -z "$$(git diff acme docs)" || \
		{ git add acme docs && \
		git commit -m "re-generate lego DNS provider data"; }

.PHONY: build
build:
	go install

.PHONY: test
test:
	TF_LOG=debug TF_ACC=1 gotestsum --format=short-verbose $(TEST) $(TESTARGS)
