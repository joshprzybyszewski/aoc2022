package twentyfive

import "testing"

func TestToSnafu(t *testing.T) {
	testCases := []struct {
		input int
		exp   string
	}{{
		input: 0,
		exp:   `0`,
	}, {
		input: 1,
		exp:   `1`,
	}, {
		input: 8,
		exp:   `2=`,
	}, {
		input: 20,
		exp:   `1-0`,
	}, {
		input: 2022,
		exp:   `1=11-2`,
	}, {
		input: 12345,
		exp:   `1-0---0`,
	}, {
		input: 314159265,
		exp:   `1121-1110-1=0`,
	}}

	for _, tc := range testCases {
		act := toSnafu(tc.input)
		if act != tc.exp {
			t.Logf("Expected %d to be %q, but was %q", tc.input, tc.exp, act)
			t.Fail()
		}
	}
}
