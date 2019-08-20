.PHONY: clean check test build dependencies fmt

export GO111MODULE=on

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test build

clean:
	rm -rf cover.out

build: clean
	 go build -v .

test: clean
	go test -v -cover ./...

check:
	golangci-lint run

fmt:
	gofmt -s -l -w $(SRCS)

## Useful to initiate structures
gen-struct:
	echo 'package namesilo' > "gen_struct.go";
	echo '' >> "gen_struct.go";
	echo 'import "encoding/xml"' >> "gen_struct.go";
	echo '' >> "gen_struct.go";
	for i in $$(ls samples/ -1); do \
		zek -c -n $${i%.xml} "./samples/$$i" >> "gen_struct.go"; \
	done
