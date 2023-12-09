package itoa

import "testing"

func TestInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{{
		input:    `1234`,
		expected: 1234,
	}, {
		input:    `-46`,
		expected: -46,
	}}

	for _, tc := range testCases {
		actual := Int(tc.input)
		if actual != tc.expected {
			t.Logf("Got %d, expected %d", actual, tc.expected)
			t.Fail()
		}
	}
}
