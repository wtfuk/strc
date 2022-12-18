# strc

The strc package provides utilities for working with strings in Go.

## Installation

To install the strc package, use the `go get` command:

`go get github.com/wtfuk/strc`

## Usage

Import the strc package in your code and call the functions as needed:

```go
package main

import (
 "fmt"

 "github.com/wtfuk/strc"
)

func main() {
 str := "hello, world"
 titleCaseStr := strc.ToTitleCase(str)
 fmt.Println(titleCaseStr) // Output: "Hello, World"
}
```

## Functions

The strc package provides the following functions:

### ToTitleCase

The ToTitleCase function converts a string to title case.

```go
ToTitleCase(str string) string
```

## Examples

For examples of how to use the strc package, see the Usage sectionf or now, more examples will be added later.

## License

The strc package is licensed under the MIT License. See the LICENSE file for details.
