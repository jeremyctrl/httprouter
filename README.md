<div align="center">

# httprouter

[![Go Reference](https://pkg.go.dev/badge/github.com/jeremyctrl/httprouter.svg)](https://pkg.go.dev/github.com/jeremyctrl/httprouter)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A tiny, fast HTTP router powered by minimal perfect hashes.

</div>

`httprouter` is a compact router that compiles route templates into minimal-perfect-hash tables at build time so lookups are constant-time.

## Get

```bash
go get github.com/jeremyctrl/httprouter
```

### Example

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/jeremyctrl/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New().
		GET("/", Index).
		GET("/hello/:name", Hello).
		Build()

	log.Fatal(http.ListenAndServe(":8080", router))
}
```