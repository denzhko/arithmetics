# arithmetics/v2

A Go package providing arithmetic operations with generics support.
Designed for clarity, correctness, and modern Go usage.

## Installation

```bash
go get github.com/denzhko/arithmetics/v2@v2.0.0
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/denzhko/arithmetics/v2"
)

func main() {
    intResult := arithmetics.Add(2, 3)
    floatResult := arithmetics.Add(2.5, 3.1)

    fmt.Println(intResult)   // 5
    fmt.Println(floatResult) // 5.6
}
```

## Features

- Supports both integers and floating-point numbers
- Generic API using Go 1.18+ type parameters
- Backward-breaking changes reflected in major version (v2)
- Minimal and easy-to-use design

## Versioning

This module follows semantic versioning.
v2.0.0 introduces generic Add, breaking backward compatibility with v1.

## License

MIT
