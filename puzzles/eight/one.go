package eight

import (
	"strings"
)

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")
	g := toGrid(lines)
	nv := numVisible(&g)

	return nv, nil
}

func toGrid(
	lines []string,
) [99][99]int {
	var output [99][99]int
	for i := range output {
		for j := range lines[i] {
			output[i][j] = int(lines[i][j] - '0')
		}
	}
	return output
}

// nolint:gocyclo yes i know
func numVisible(
	grid *[99][99]int,
) int {
	total := (4 * len(grid)) - 4

	var visibles [99][99]bool

	var l, r, u, d, e, j int

	for i := 1; i < len(grid)-1; i++ {
		l = grid[i][0]
		r = grid[i][len(grid)-1]
		u = grid[0][i]
		d = grid[len(grid)-1][i]

		for j = 1; j < len(grid)-1; j++ {
			// iterate through the middle of the forest

			// look from left
			if e = grid[i][j]; e > l {
				if !visibles[i][j] {
					total++
				}
				visibles[i][j] = true
				l = e
			}

			// look from right
			if e = grid[i][len(grid)-1-j]; e > r {
				if !visibles[i][len(grid)-1-j] {
					total++
				}
				visibles[i][len(grid)-1-j] = true
				r = e
			}

			// look from up
			if e = grid[j][i]; e > u {
				if !visibles[j][i] {
					total++
				}
				visibles[j][i] = true
				u = e
			}

			// look from down
			if e = grid[len(grid)-1-j][i]; e > d {
				if !visibles[len(grid)-1-j][i] {
					total++
				}
				visibles[len(grid)-1-j][i] = true
				d = e
			}
		}
	}

	return total
}
