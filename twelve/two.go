package twelve

import (
	"strconv"
)

func Two(
	input string,
) (string, error) {
	g, _, e := newGrid(input)
	min := len(g)*len(g[0]) + 1
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] != 0 {
				continue
			}
			n := getStepsBetween(
				g,
				coord{
					row: r,
					col: c,
				},
				e,
			)
			if n > 0 && n < min {
				min = n
			}
		}
	}
	return strconv.Itoa(min), nil
}
