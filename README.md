# nekos

> 🐾 nya~

[![Go Reference](https://pkg.go.dev/badge/github.com/x6r/nekos.svg)](https://pkg.go.dev/github.com/x6r/nekos/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/x6r/nekos)](https://goreportcard.com/report/github.com/x6r/nekos)
[![Example](https://img.shields.io/badge/Example-__example%2F-C14DAA?style=flat)](https://github.com/x6r/nekos/blob/master/_example/main.go)

nekos is a feature-complete _unofficial_ [nekos.life](https://nekos.life/) api wrapper written in go. each endpoint in the api has a function that returns its response(s).

## installation

```sh
go get -u github.com/x6r/nekos
```

## usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/x6r/nekos/v2"
)

func main() {
	cat, err := nekos.CatText()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cat)
}
```

## license

[OSL-3.0](https://github.com/x6r/nekos/blob/master/LICENSE)
