# strpad

String padding for Go.

## Usage

```go
import (
    "fmt"
    "github.com/joshburnsxyz/strpad"
)

func main() {
    myPaddedString := strpad.Pad("Hello, world.")
    fmt.Println(myPaddedString)
    // "Hello, world." becomes " Hello, world. "
}
```

The `Pad()` function can be swapped out for `PadLeft()` or `PadRight()` using
the identical API to pad the left or right side of a string without the other.

There are also functions that follow the `PadN()` naming convention
(`PadNLeft()`, `PadNRight()`). These functions take an extra integer as a param
to indicate the amount of spaces to be padded onto the string.

```go

import (
    "fmt"
    "github.com/joshburnsxyz/strpad"
)

func main() {
    myPaddedString := strpad.PadN("Hello, world.", 2)
    fmt.Println(myPaddedString)
    // "Hello, world." becomes "  Hello, world.  "
}

```

The complete API documentation can be found on
[pkg.go.dev](https://pkg.go.dev/github.com/joshburnsxyz/strpad)

## Installation

```
go get github.com/joshburnsxyz/strpad
```

## License

This software is open-sourced under the MIT license.
