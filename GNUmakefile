# Test all packages by default
TEST ?= ./...

ifeq ($(shell go env GOOS),darwin)
SEDOPTS = -i ''
else
SEDOPTS = -i
endif

GOWORKFLOWVERSION=$(shell go run ./build-support/go-version-for-workflow)

.PHONY: default
default: build

.PHONY: tools
tools:
	cd $(shell go env GOROOT) && go install github.com/hashicorp/go-bindata/go-bindata@latest && go install gotest.tools/gotestsum@latest

.PHONY: pebble-start-install
pebble-start-install: pebble-stop
	build-support/scripts/pebble-start.sh --install

.PHONY: pebble-start
pebble-start: pebble-stop
	build-support/scripts/pebble-start.sh

.PHONY: pebble-stop
pebble-stop:
	build-support/scripts/pebble-stop.sh

.PHONY: memcached-start
memcached-start: memcached-stop
	build-support/scripts/memcached-start.sh

.PHONY: memcached-stop
memcached-stop:
	build-support/scripts/memcached-stop.sh

.PHONY: stop-services
stop-services: memcached-stop pebble-stop

.PHONY: template-generate
template-generate:
	@echo "==> Re-generating templates..."
	@go generate ./build-support/generate-dns-providers

.PHONY: provider-generate
provider-generate:
	@echo "==> Re-generating Go DNS provider factory in ./acme..."
	@go generate ./acme/dnsplugin
	@go build ./acme/dnsplugin
	@go mod tidy
	@echo "==> Re-genrating documentation..."
	@rm -f docs/guides/dns-providers-*.md
	@go run ./build-support/generate-dns-providers doc docs/guides/

.PHONY: provider-generate-update
provider-generate-update: provider-generate
	test -z "$$(git diff acme docs)" || \
		{ git add acme docs && \
		git commit -m "re-generate lego DNS provider data"; }

.PHONY: proto
proto:
	cd proto/ && buf generate

.PHONY: build
build:
	go install

.PHONY: release
release:
	build-support/scripts/release.sh

.PHONY: build-pre-release
build-pre-release:
	mkdir -p /tmp/terraform-provider-acme/
	go build -o /tmp/terraform-provider-acme/terraform-provider-acme

.PHONY: clean-pre-release
clean-pre-release:
	rm -rf /tmp/terraform-provider-acme/

.PHONY: test
test:
	TF_ACC=1 gotestsum --format=short-verbose $(TEST) $(TESTARGS)

.PHONY: go-version-sync
go-version-sync:
	sed $(SEDOPTS) -e "s/go-version:.*\$$/go-version: '^$(GOWORKFLOWVERSION)'/g" .github/workflows/*.yml
	git add .github/workflows/*.yml && git commit -m "workflows: update Go to version $(GOWORKFLOWVERSION)"
