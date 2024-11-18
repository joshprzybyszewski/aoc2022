package day02

import (
	"strings"

	"github.com/joshprzybyszewski/aoc2022/util/lines"
)

func Two(
	input string,
) (int, error) {

	sum := 0
	var isFirst bool
	var hi int
	var handful, min handful

	lines.ForEach(input, func(line string) {
		isFirst = true

		line = line[strings.Index(line, `:`)+1:]
		for {
			hi = strings.Index(line, semicolon)
			if hi == -1 {
				hi = len(line)
			}

			handful = interpretSeen(line[:hi])
			if isFirst {
				min = handful
				isFirst = false
			} else {
				if handful.red > min.red {
					min.red = handful.red
				}
				if handful.green > min.green {
					min.green = handful.green
				}
				if handful.blue > min.blue {
					min.blue = handful.blue
				}
			}

			if hi == len(line) {
				break
			}

			line = line[hi+1:]
		}

		sum += (min.red * min.blue * min.green)
	})

	return sum, nil
}
