package strc

import (
	"reflect"
	"testing"
)

func TestPadSpaceToRight2D(t *testing.T) {
	perfectCase := [][]string{{"1", "1", "1", "1"}, {"22", "22", "22", "22"}}
	perfectCaseOut := [][]string{{"1 ", "1 ", "1 ", "1 "}, {"22", "22", "22", "22"}}
	// Define a slice of test cases
	testCases := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{"Perfect Case", perfectCase, perfectCaseOut},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := PadSpaceToRight2D(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("input \x1b[34m%q\x1b[0m expected \x1b[32m%q\x1b[0m but got \x1b[31m%q\x1b[0m", tc.input, tc.expected, result)
			}
		})
	}
}
