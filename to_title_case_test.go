package strc

import "testing"

func TestToTitleCase(t *testing.T) {
	// Define a slice of test cases
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hello, World", "hello, world", "Hello, World"},
		{"The Quick Brown Fox", "the quick brown fox", "The Quick Brown Fox"},
		{"All Caps", "HELLO, WORLD", "Hello, World"},
		{"Empty String", "", ""},
		{"Complex case", "C  o mplex", "C  O Mplex"},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToTitleCase(tc.input)
			if result != tc.expected {
				t.Errorf("input \x1b[34m%q\x1b[0m expected \x1b[32m%q\x1b[0m but got \x1b[31m%q\x1b[0m", tc.input, tc.expected, result)
			}
		})
	}
}
