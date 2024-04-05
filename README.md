# Version display and features count

## Version generate

```go
//go:build generate

package main

//go:generate go run thisFile.go

import (
    "fmt"
    "github.com/dyammarcano/version"
)

func main() {
    ver, err := version.NewVersion()
    if err != nil {
        fmt.Println(err)
        return
    }
	
    if err = ver.Generate(); err != nil {
        fmt.Println(err)
        return
    }
}
```
