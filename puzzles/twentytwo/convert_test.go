package twentytwo

import "testing"

func Test_willStop(t *testing.T) {
	testCases := []struct {
		input string
		exp   bool
	}{{
		input: `0,0,1~0,0,2
0,0,3~0,0,4
`,
		exp: true,
	}, {
		input: `0,0,1~0,0,2
0,0,3~0,0,3
`,
		exp: true,
	}, {
		input: `0,0,1~0,0,2
0,0,3~0,1,3
`,
		exp: true,
	}, {
		input: `0,0,1~0,0,2
0,1,3~0,1,3
`,
		exp: false,
	}, {
		input: `0,0,1~0,2,1
0,1,3~0,1,4
`,
		exp: true,
	}}

	for _, tc := range testCases {
		blocks := convertInput(tc.input)
		if willStop(blocks[1], blocks[0]) != tc.exp {
			t.Fail()
		}
	}
}
