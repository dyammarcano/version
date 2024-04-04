# Version display and features count

## Version display

```go
package main

import (
    "fmt"
    "github.com/dyammarcano/version"
)

func init() {
	version.AddFeature("feature 1")
}

func main() {
    fmt.Println(version.GetVersionInfo())
}
```