# arithmetics

A lightweight Go package that provides essential arithmetic operations with a clean and minimal API.

## Installation

```bash
go get github.com/denzhko/arithmetics@v1.0.1
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/denzhko/arithmetics"
)

func main() {
    result := arithmetics.Add(2, 3)
    fmt.Println(result) // 5
}
```

## Features

- Minimal and intuitive API
- Focus on correctness and simplicity
- Works with integers only

## Versioning

This module follows semantic versioning.
v1.x.x versions provide stable, integer-only addition.

## License

MIT
