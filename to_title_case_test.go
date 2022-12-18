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
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		// Use t.Run to run the test in a subtest, with the test case name as the subtest name
		t.Run(tc.name, func(t *testing.T) {
			// Call the ToTitleCase function with the input string and compare the result to the expected output string
			result := ToTitleCase(tc.input)
			if result != tc.expected {
				t.Errorf("input \x1b[34m%q\x1b[0m expected \x1b[32m%q\x1b[0m but got \x1b[31m%q\x1b[0m", tc.input, tc.expected, result)
			}
		})
	}
}
