[![ci-test](https://github.com/dyammarcano/version/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/dyammarcano/version/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dyammarcano/version)](https://goreportcard.com/report/github.com/dyammarcano/version)
[![Go Reference](https://pkg.go.dev/badge/github.com/dyammarcano/version.svg)](https://pkg.go.dev/github.com/dyammarcano/version)

# Version display and features names

This package is used to display the version of the application and the number of features that it has. It uses the git tags to display the version and the number of features. and it generates a file with the version json data in the project root.

## How to use

```bash
$ go get -u "github.com/dyammarcano/version"
```
## Example gen.go file

```go
//go:build generate

package main

//go:generate go run gen.go

import (
	"github.com/dyammarcano/version"
)

func main() {
	ver, err := version.NewVersion()
	if err != nil {
		panic(err)
	}

	if err = ver.Generate(); err != nil {
		panic(err)
	}
}

```

## Map feature in your project

```go
package main

import (
    "yourproject/version"
)

init() {
    version.AddFeature("feature1")
}
```

## How to generate the version file

```bash
$ go generate ./...
```

## VERSION file

```json
{"version":"v0.0.0","commitHash":"5f600de951b8a0c5bb7b035d8f95aaaaf534c9e3","date":"2024-05-15T02:24:46Z"}
```