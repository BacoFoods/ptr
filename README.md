# ptr

`ptr` is a Go library that provides a set of helper functions for easily
handling literal pointers.

## Installation

```sh
go get github.com/BacoFoods/ptr
```

## Features

- Simplifies the creation of pointers for literal values.
- Generic support for any type.

## Example

```go
package main

import (
	"fmt"
	"github.com/BacoFoods/ptr"
)

type example struct {
  Name *string
  Price *float64
}

func main() {
  e := example{
    Name: ptr.Ptr("Example"),
    Price: ptr.Ptr(42.0)
  }

  log.Println(ptr.Val(e.Name, "Default Name"))

  if ptr.EqualVals(e.Price, ptr.Ptr(82)) {
    log.Println("Not the same")
  }
}
```

## License

This library is open-source and available under the MIT License.
