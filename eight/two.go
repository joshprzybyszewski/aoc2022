package eight

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")
	g := toGrid(lines)
	max := getMaxSceneryScore(&g)

	return strconv.Itoa(max), nil
}

func getMaxSceneryScore(
	grid *[99][99]int,
) int {

	// all of the trees on the edge have a viewing score of zero in one direction
	max := 0

	var j, k, ss int

	for i := 1; i < len(grid)-1; i++ {
		for j = 1; j < len(grid)-1; j++ {
			// iterate through the middle of the forest

			// look to the right
			for k = 1; j+k < len(grid)-1 && grid[i][j+k] < grid[i][j]; k++ {
			}
			ss = k

			// look to the left
			for k = 1; j-k > 0 && grid[i][j-k] < grid[i][j]; k++ {
			}
			ss *= k

			// look down
			for k = 1; i+k < len(grid)-1 && grid[i+k][j] < grid[i][j]; k++ {
			}
			ss *= k

			// look up
			for k = 1; i-k > 0 && grid[i-k][j] < grid[i][j]; k++ {
			}
			ss *= k

			if ss > max {
				max = ss
			}
		}
	}

	return max
}
