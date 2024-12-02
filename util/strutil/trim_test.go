package strutil

import "testing"

func TestTrim(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{{
		input:    `1234`,
		expected: `1234`,
	}, {
		input:    ` 0123`,
		expected: `0123`,
	}, {
		input:    `2345 `,
		expected: `2345`,
	}, {
		input:    ` 4567 `,
		expected: `4567`,
	}, {
		input:    `  `,
		expected: ``,
	}}

	for _, tc := range testCases {
		actual := TrimSpaces(tc.input)
		if actual != tc.expected {
			t.Logf("Got %q, expected %q", actual, tc.expected)
			t.Fail()
		}
	}
}
