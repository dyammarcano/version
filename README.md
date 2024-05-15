# Version display and features names

This package is used to display the version of the application and the number of features that it has. It uses the git tags to display the version and the number of features. and it generates a file with the version json data in the project root.
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