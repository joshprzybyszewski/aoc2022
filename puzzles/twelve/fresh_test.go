package twelve

import "testing"

func TestFresh(t *testing.T) {
	type testcase struct {
		input string
		exp   int
	}

	for _, tc := range []testcase{{
		input: `????#?.???? 2,2,1`,
		exp:   14,
	}, {
		input: `??? 1`,
		exp:   3,
	}, {
		input: `???? 1,1`,
		exp:   3,
	}, {
		input: `????? 1,1`,
		exp:   6,
	}, {
		input: `????.? 1,1`,
		exp:   7,
	}, {
		input: `?###??????????###??????????###??????????###??????????###???????? 3,2,1,3,2,1,3,2,1,3,2,1,3,2,1`,
		exp:   506250,
	}, {
		input: `#.?#?.#.????.???# 1,3,1,2,4`,
		exp:   3,
	}, {
		input: `?.???.?#???#???.? 3,6,1`,
		exp:   5,
	}, {
		input: `?#?...#??.# 2,1,1`,
		exp:   2,
	}} {
		p, _ := newPossibilities(tc.input)
		p.build()
		act := p.answer()
		if act != tc.exp {
			t.Logf("Expected: %d, actual: %d. For %q", tc.exp, act, tc.input)
			t.Fail()
		}
	}
}

func TestUnfold(t *testing.T) {
	type testcase struct {
		input string
		exp   string
	}

	for _, tc := range []testcase{{
		input: `# 1`,
		exp:   `#?#?#?#?# 1,1,1,1,1`,
	}} {
		act, _ := newPossibilities(tc.input)
		unfold(&act)
		exp, _ := newPossibilities(tc.exp)
		if act != exp {
			t.Logf("Expected: %+v\n, actual: %+v\n. For %q", exp, act, tc.input)
			t.Fail()
		}
	}
}
