package nine

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {

	lines := strings.Split(input, "\n")

	var knots [10]coord
	var q, i, j int
	var err error

	pos := make(map[coord]struct{}, int(1<<13))
	pos[knots[len(knots)-1]] = struct{}{}

	for _, line := range lines {
		if line == `` {
			continue
		}
		q, err = strconv.Atoi(line[2:])
		if err != nil {
			return ``, err
		}

		for i = 0; i < q; i++ {
			switch direction(line[0]) {
			case right:
				knots[0].x++
			case left:
				knots[0].x--
			case up:
				knots[0].y++
			case down:
				knots[0].y--
			}

			for j = 1; j < len(knots); j++ {
				knots[j] = moveCoord(knots[j], knots[j-1])
			}
			pos[knots[len(knots)-1]] = struct{}{}
		}
	}

	return strconv.Itoa(len(pos)), nil
}
