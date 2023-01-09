package strc

import (
	"strings"
	"unicode/utf8"
)

// PadSpaceToRight pads spaces to the right of the given string
//
// Example:
//   s := PadSpaceToRight("abc", 5) // s is "abc  "
func PadSpaceToRight(str string, len int) string {
	if len-utf8.RuneCountInString(str) >= 0 {
		return str + strings.Repeat(" ", len-utf8.RuneCountInString(str))
	}
	return str
}
