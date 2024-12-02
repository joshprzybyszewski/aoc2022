package strutil

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

func TestIntBeforeSpace(t *testing.T) {
	testCases := []struct {
		input         string
		expected      int
		expectedIndex int
	}{{
		input:         `1234`,
		expected:      1234,
		expectedIndex: 4,
	}, {
		input:         `-46`,
		expected:      -46,
		expectedIndex: 3,
	}, {
		input:         `1234 more`,
		expected:      1234,
		expectedIndex: 4,
	}, {
		input:         `-46 more`,
		expected:      -46,
		expectedIndex: 3,
	}, {
		input:         `- 46 more`,
		expected:      0,
		expectedIndex: 1,
	}, {
		input:         ` 1234 more`,
		expected:      0,
		expectedIndex: 0,
	}}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual, index := IntBeforeSpace(tc.input)
			if actual != tc.expected {
				t.Logf("Got %d, expected %d", actual, tc.expected)
				t.Fail()
			}
			if index != tc.expectedIndex {
				t.Logf("Index got %d, expected %d", index, tc.expectedIndex)
				t.Fail()
			}
		})
	}
}
