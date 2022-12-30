# nekos

> 🐾 nya~

[![Latest Release](https://img.shields.io/github/release/fawni/nekos.svg)](https://github.com/fawni/nekos/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/fawni/nekos.svg)](https://pkg.go.dev/github.com/fawni/nekos/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/fawni/nekos)](https://goreportcard.com/report/github.com/fawni/nekos)

nekos is a feature-complete _unofficial_ [nekos.life](https://nekos.life/) api wrapper written in go. each endpoint in the api has a function that returns its response(s).

## installation

```sh
go get -u github.com/fawni/nekos/v2
```

## usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/fawni/nekos/v2"
)

func main() {
	cat, err := nekos.CatText()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cat) // (ฅ’ω’ฅ)
}
```

## license

[OSL-3.0](https://github.com/fawni/nekos/blob/master/LICENSE)
