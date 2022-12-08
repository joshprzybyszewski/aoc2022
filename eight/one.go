package eight

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")
	grid := toGrid(lines)
	visibility := toVisible(grid)
	nv := numVisible(visibility)

	return strconv.Itoa(nv), nil
}

func toGrid(
	lines []string,
) [][]int {
	output := make([][]int, len(lines)-1)
	for i := range output {
		output[i] = make([]int, len(lines[i]))
		for j := range lines[i] {
			output[i][j] = int(lines[i][j] - '0')
		}
	}
	return output
}

type visibility uint8

const (
	none     visibility = 0
	fromLeft visibility = 1 << iota
	fromRight
	fromUp
	fromDown
)

func toVisible(
	grid [][]int,
) [][]visibility {
	// [row][col]
	output := make([][]visibility, len(grid))
	for i := range grid {
		output[i] = make([]visibility, len(grid[i]))
	}

	for i := range output {
		// all the first row is visible from the top
		output[0][i] |= fromUp

		// all the last row is visible from the bottom
		output[len(output)-1][i] |= fromDown

		// all the left col is visible from the left
		output[i][0] |= fromLeft

		// all the right col is visible from the left
		output[i][len(output)-1] |= fromRight
	}

	var l, r, u, d, e, j int

	for i := 1; i < len(output)-1; i++ {
		l = grid[i][0]
		r = grid[i][len(grid)-1]
		u = grid[0][i]
		d = grid[len(grid)-1][i]

		for j = 1; j < len(output)-1; j++ {
			// iterate through the middle of the forest

			// look from left
			if e = grid[i][j]; e > l {
				output[i][j] |= fromLeft
				l = e
			}

			// look from right
			if e = grid[i][len(grid)-1-j]; e > r {
				output[i][len(grid)-1-j] |= fromRight
				r = e
			}

			// look from up
			if e = grid[j][i]; e > u {
				output[j][i] |= fromUp
				u = e
			}

			// look from down
			if e = grid[len(grid)-1-j][i]; e > d {
				output[len(grid)-1-j][i] |= fromDown
				d = e
			}
		}
	}

	return output
}

func numVisible(
	vs [][]visibility,
) int {
	total := 0

	for i := range vs {
		for j := range vs[i] {
			if vs[i][j] != none {
				total++
			}
		}
	}

	return total
}
