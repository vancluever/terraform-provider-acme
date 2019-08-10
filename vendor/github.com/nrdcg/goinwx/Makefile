.PHONY: default clean check test fmt

GOFILES := $(shell go list -f '{{range $$index, $$element := .GoFiles}}{{$$.Dir}}/{{$$element}}{{"\n"}}{{end}}' ./... | grep -v '/vendor/')

default: clean check test build

test: clean
	go test -v -cover ./...

clean:
	rm -f cover.out

build:
	go build

fmt:
	gofmt -s -l -w $(GOFILES)

check:
	golangci-lint run
