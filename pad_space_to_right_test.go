package strc

import "testing"

func TestPadSpaceToRight(t *testing.T) {
	// Define a slice of test cases
	testCases := []struct {
		name     string
		input    string
		len      int
		expected string
	}{
		{"Pad String", "abc", 5, "abc  "},
		{"No Padding Necessary", "abc", 2, "abc"},
		{"Pad Multibyte String", "123", 5, "123  "},
		{"No Padding Necessary for Multibyte String", "ohmy", 3, "ohmy"},
		{"Empty String", "", 5, "     "},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := PadSpaceToRight(tc.input, tc.len)
			if result != tc.expected {
				t.Errorf("input \x1b[34m%q\x1b[0m expected \x1b[32m%q\x1b[0m but got \x1b[31m%q\x1b[0m", tc.input, tc.expected, result)
			}
		})
	}
}
