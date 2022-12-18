package strc

import (
	"unicode"
)

// ToTitleCase converts a string to title case
func ToTitleCase(str string) string {
	// Make a new string with the same length as the original string
	titleCaseStr := make([]rune, len(str))

	// Set the titleCase flag to true to indicate that the next character should be uppercase
	titleCase := true

	// Iterate over each character in the string
	for i, ch := range str {
		if titleCase {
			// If the titleCase flag is true, convert the character to uppercase and reset the flag
			titleCaseStr[i] = unicode.ToUpper(ch)
			titleCase = false
		} else if unicode.IsSpace(ch) {
			// If the character is a space, set the titleCase flag to true to indicate that the next character should be uppercase
			titleCase = true
			titleCaseStr[i] = ch
		} else {
			// If the character is not a space, convert it to lowercase
			titleCaseStr[i] = unicode.ToLower(ch)
		}
	}

	// Return the title case string as a new string value
	return string(titleCaseStr)
}
