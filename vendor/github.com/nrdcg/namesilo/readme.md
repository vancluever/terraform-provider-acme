# Go library for accessing the Namesilo API

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/nrdcg/namesilo.svg)](https://github.com/nrdcg/namesilo/releases)
[![GoDoc](https://godoc.org/github.com/nrdcg/namesilo?status.svg)](https://godoc.org/github.com/nrdcg/namesilo)
[![Build Status](https://travis-ci.com/nrdcg/namesilo.svg?branch=master)](https://travis-ci.com/nrdcg/namesilo)

A Namesilo API client written in Go.

namesilo is a Go client library for accessing the Namesilo API.

## Example


```go
package main

import (
	"fmt"
	"log"

	"github.com/nrdcg/namesilo"
)

func main() {
	transport, err := namesilo.NewTokenTransport("1234")
	if err != nil {
		log.Fatal(err)
	}

	client := namesilo.NewClient(transport.Client())

	params := &namesilo.AddAccountFundsParams{
		Amount:    "1000000",
		PaymentID: "acbd",
	}

	funds, err := client.AddAccountFunds(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(funds)
}
```

## API Documentation

- [API docs](https://www.namesilo.com/api_reference.php)
