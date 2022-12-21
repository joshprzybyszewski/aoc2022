package nine

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (int, error) {

	var knots [10]coord
	var q, i, j int
	var err error

	pos := make(map[coord]struct{}, int(1<<13))
	pos[knots[len(knots)-1]] = struct{}{}

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}
		q, err = strconv.Atoi(input[2:nli])
		if err != nil {
			return 0, err
		}

		for i = 0; i < q; i++ {
			switch direction(input[0]) {
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
				knots[j].moveToward(knots[j-1])
			}
			pos[knots[len(knots)-1]] = struct{}{}
		}
		input = input[nli+1:]
	}

	return len(pos), nil
}
