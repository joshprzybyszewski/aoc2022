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

	/*
		When we start at (0,0):
		   max: {x:45 y:163}
		   min: {x:-110 y:-108}
	*/
	for i := range knots {
		knots[i].x = 110
		knots[i].y = 108
	}
	// 166 = 45-(-110) + 1
	// 274 = 163-(-108)+1
	var seen [166][272]bool

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
			seen[knots[9].x][knots[9].y] = true
		}
		input = input[nli+1:]
	}

	total := 0
	for i := range seen {
		for j := range seen[i] {
			if seen[i][j] {
				total++
			}
		}
	}

	return total, nil
}
