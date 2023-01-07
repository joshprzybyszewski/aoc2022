package twentyfour

import "fmt"

func Two(
	input string,
) (int, error) {
	b := getBoard(input)
	ab := populatedAllBoards(&b)

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
	if p.numSteps == -1 {
		return 0, fmt.Errorf("path not found")
	}
	numSteps := p.numSteps

	p = navigate(&ab, s, p.bs, g)
	if p.numSteps == -1 {
		return 0, fmt.Errorf("path not found")
	}
	numSteps += p.numSteps

	p = navigate(&ab, s, p.bs, g)
	if p.numSteps == -1 {
		return 0, fmt.Errorf("path not found")
	}
	numSteps += p.numSteps

	return numSteps, nil
}
