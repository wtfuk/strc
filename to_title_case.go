package strc

import (
	"unicode"
)

// ToTitleCase converts a string to title case
func ToTitleCase(str string) string {

	titleCaseStr := make([]rune, len(str))
	makeUpper := true

	for i, ch := range str {
		if makeUpper && !unicode.IsSpace(ch) {
			titleCaseStr[i] = unicode.ToUpper(ch)
			makeUpper = false
		} else if unicode.IsSpace(ch) {
			makeUpper = true
			titleCaseStr[i] = ch
		} else {
			titleCaseStr[i] = unicode.ToLower(ch)
		}
	}

	return string(titleCaseStr)
}
