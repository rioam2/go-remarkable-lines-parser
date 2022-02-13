# go-remarkable-lines-parser

Golang package to parse `.rm` lines files from the reMarkable Paper Tablet

## Quick Start

```bash
go get -u github.com/rioam2/go-remarkable-lines-parser
```

```go
package main

import (
  "fmt"
  goremarkablelinesparser "github.com/rioam2/go-remarkable-lines-parser"
)

func main() {
  file := os.Open("./my-lines-file.rm")
  page, err := goremarkablelinesparser.FromReader(file)
  if err != nil {
    panic(err)
  }
  fmt.Println(page)
}
```
