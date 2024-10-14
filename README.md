# Obvious `timestamp` helper methods

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE-OF-CONDUCT.md)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
![GitHub release](https://img.shields.io/github/release/go-obvious/timestamp.svg)

Simple timestamp utilities focused on UTC

## How to Use

### Installation

```sh
go get github.com/go-obvious/timestamp
```

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/go-obvious/timestamp"
)

func main() {
    fmt.Println("Epoche:", timestamp.Milli())
}
```
