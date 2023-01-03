package twentyfour

import "fmt"

func One(
	input string,
) (int, error) {

	b := getBoard(input)
	ab := newAllBoards(b)

	s := position{
		row: 0,
	}
	g := position{
		row: len(b) - 1,
	}
	for c := range b[0] {
		if b[s.row][c] == empty {
			s.col = c
		}
		if b[g.row][c] == empty {
			g.col = c
		}
	}

	p := navigate(&ab, s, 0, g)
	if p == nil {
		return 0, fmt.Errorf("path not found")
	}

	numSteps := 0
	for p != nil {
		numSteps++
		p = p.prev
	}
	numSteps--

	return numSteps, nil
}
