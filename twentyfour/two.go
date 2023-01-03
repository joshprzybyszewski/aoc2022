package twentyfour

import "fmt"

func Two(
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
	sb := p.bs

	numSteps := 0
	for p != nil {
		numSteps++
		// fmt.Println(prettyBoard(ab.getBoardAtState(p.bs), p.cur))
		p = p.prev
	}
	numSteps--

	p = navigate(&ab, s, sb, g)
	if p == nil {
		return 0, fmt.Errorf("path not found")
	}
	sb = p.bs

	for p != nil {
		numSteps++
		// fmt.Println(prettyBoard(ab.getBoardAtState(p.bs), p.cur))
		p = p.prev
	}
	numSteps--

	p = navigate(&ab, s, sb, g)
	if p == nil {
		return 0, fmt.Errorf("path not found")
	}
	sb = p.bs

	for p != nil {
		numSteps++
		// fmt.Println(prettyBoard(ab.getBoardAtState(p.bs), p.cur))
		p = p.prev
	}
	numSteps--

	return numSteps, nil
}
