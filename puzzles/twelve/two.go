package twelve

func Two(
	input string,
) (int, error) {
	g, _, e := newGrid(input)
	steps := paint(
		&g,
		e,
		coord{
			row: numRows + 1,
			col: numCols + 1,
		},
	)

	min := len(g)*len(g[0]) + 1
	var n int
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] != 0 {
				continue
			}
			n = steps[r][c]
			if n > 0 && n < min {
				min = n
			}
		}
	}
	return min, nil
}
