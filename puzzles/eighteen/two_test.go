package eighteen

import "testing"

func Test_getDistanceFromHexCode(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{{
		input: `#000000`,
		exp:   0,
	}, {
		input: `#000001`,
		exp:   0,
	}, {
		input: `#000010`,
		exp:   1,
	}, {
		input: `#000090`,
		exp:   9,
	}, {
		input: `#0000a0`,
		exp:   10,
	}, {
		input: `#0000f0`,
		exp:   15,
	}, {
		input: `#000100`,
		exp:   16,
	}, {
		input: `#001000`,
		exp:   256,
	}, {
		input: `#010000`,
		exp:   4096,
	}, {
		input: `#100000`,
		exp:   65536,
	}, {
		input: `#700000`,
		exp:   458752,
	}, {
		input: `#70c710`,
		exp:   461937,
	}}

	for _, tc := range testCases {
		act := getDistanceFromHexCode(tc.input)
		if act != tc.exp {
			t.Logf("Expected %d, got %d", tc.exp, act)
			t.Fail()
		}
	}
}
