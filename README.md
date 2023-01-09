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

### PadSpaceToRight

The PadSpaceToRight function pads spaces to the right of the given string.

```go
PadSpaceToRight(str string, len int) string
```

## Examples

`
strc.ToTitleCase("this does title cases") // output: This Does Title Cases
`

`
strc.PadSpaceToRight("123", 6) // output: This will add 3 spaces to the right of "123"
`

## License

The strc package is licensed under the MIT License. See the LICENSE file for details.
