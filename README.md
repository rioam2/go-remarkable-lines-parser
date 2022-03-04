# go-remarkable-lines-parser

Golang package to parse `.rm` lines files from the reMarkable Paper Tablet.
Currently only supports `.rm` version 5 files.

## Quick Start

```bash
go get -u github.com/rioam2/go-remarkable-lines-parser
```

```go
package main

import (
	"fmt"
	"os"

	goremarkablelinesparser "github.com/rioam2/go-remarkable-lines-parser"
)

func main() {
	file, err := os.Open("./my-lines-file.rm")
	if err != nil {
		panic(err)
	}
	page, err := goremarkablelinesparser.FromReader(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(page)
}

```
