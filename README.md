# ALS-Go Client Library for Go Lang projects
[Website](https://www.riftbit.com) |
[Contributing](https://www.riftbit.com/How-to-Contribute)

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/RiftBit/als-go-client)
[![Build Status](https://travis-ci.org/RiftBit/als-go-client.svg?branch=master)](https://travis-ci.org/RiftBit/als-go-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/RiftBit/als-go-client)](https://goreportcard.com/report/github.com/RiftBit/als-go-client)
[![Coverage Status](https://coveralls.io/repos/github/RiftBit/als-go-client/badge.svg?branch=master)](https://coveralls.io/github/RiftBit/als-go-client?branch=master)

## Example

```golang
package main

import (
	"github.com/Riftbit/als-go-client"
	"fmt"
	"github.com/Riftbit/ALS-Go/httpmodels"
)

func main() {
	conf := alsgoclient.Config{
		Uri: "http://als.local:8080/",
		Login: "ergoz",
		Password: "ergoz",
		IsAsync: false,
		Timeout: 1000,
	}
	client, err := alsgoclient.New(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("=============================")

	cats, _ := client.GetCategories()
	for _, dat := range cats.CategoriesList {
		fmt.Println(dat)
	}

	fmt.Println("=============================")

	args := httpmodels.RequestLogGetLog{Category: "capi", Limit: 10}
	result, _ := client.Get(args)
	for _,data := range result.LogList {
		fmt.Println(data.ID, "-", data.Message)
	}

	fmt.Println("=============================")
}
```
