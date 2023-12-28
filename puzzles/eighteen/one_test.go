package eighteen

import "testing"

type testLagoon struct {
	input       string
	expectedOne int
}

var (
	// draw a square
	tl1 = testLagoon{
		input: `R 1 (#000000)
D 1 (#000000)
L 1 (#000000)
U 1 (#000000)
`,
		expectedOne: 4,
	}
)

var (
	// draw a c kinda
	tl2 = testLagoon{
		input: `R 4 (#000000)
D 5 (#000000)
L 3 (#000000)
U 2 (#000000)
R 1 (#000000)
U 1 (#000000)
L 2 (#000000)
U 2 (#000000)
`,
		expectedOne: 20 + 7,
	}
)

var (
	// draw an icicle kinda
	tl3 = testLagoon{
		input: `D 20 (#000000)
R 1 (#000000)
U 3 (#000000)
R 2 (#000000)
U 3 (#000000)
R 4 (#000000)
U 3 (#000000)
R 8 (#000000)
U 11 (#000000)
L 15 (#000000)
`,
		expectedOne: (21 * 16) - (8 * 9) - (4 * 6) - (2 * 3),
	}
)

var (
	// draw an E
	tl4 = testLagoon{
		input: `R 10 (#000000)
D 3 (#000000)
L 8 (#000000)
D 3 (#000000)
R 4 (#000000)
D 3 (#000000)
L 3 (#000000)
D 3 (#000000)
R 7 (#000000)
D 3 (#000000)
L 10 (#000000)
U 15 (#000000)
`,
		expectedOne: (4 * 11) + (3 * 2) + (7 * 4) + (4 * 2) + (11 * 4),
	}
)

var (
	// tetris-kinda z shape
	tl5 = testLagoon{
		input: `D 3 (#000000)
L 3 (#000000)
D 5 (#000000)
R 3 (#000000)
U 3 (#000000)
R 3 (#000000)
U 5 (#000000)
L 3 (#000000)
`,
		expectedOne: (4 * 6) + (7 * 3),
	}
)

var (
	// tetris-kinda t shape
	tl6 = testLagoon{
		input: `L 3 (#000000)
D 3 (#000000)
L 3 (#000000)
D 3 (#000000)
R 3 (#000000)
D 3 (#000000)
R 3 (#000000)
U 9 (#000000)
`,
		expectedOne: (4 * 13),
	}
)

var (
	example = testLagoon{
		input: `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`,
		expectedOne: 62,
	}
)

var (
	what = testLagoon{
		input: `R 5 (#000000)
D 5 (#000000)
R 5 (#000000)
D 4 (#000000)
L 5 (#000000)
D 4 (#000000)
R 8 (#000000)
D 6 (#000000)
L 8 (#000000)
D 4 (#000000)
L 4 (#000000)
U 10 (#000000)
L 1 (#000000)
U 13 (#000000)
`,
		expectedOne: 215, // found this manually
	}
)

func TestLagoons(t *testing.T) {
	for _, tl := range []testLagoon{
		// tl1,
		// tl2,
		// tl3,
		// tl4,
		// tl5,
		// tl6,
		// example,
		what,
	} {
		act, _ := One(tl.input)
		if act != tl.expectedOne {
			l := newLagoon(tl.input, false)
			t.Logf("Expected %d, actual %d for\n%s\n", tl.expectedOne, act, l.String())
			t.Fail()
		}
		t.Fail()
	}

}
