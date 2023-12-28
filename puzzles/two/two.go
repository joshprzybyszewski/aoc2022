package two

import (
	"strings"
)

func Two(
	input string,
) (int, error) {

	sum := 0
	var isFirst bool
	var hi int
	var line string
	var handful, min handful

	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		isFirst = true

		line = input[strings.Index(input, `:`)+1 : nli]
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
		input = input[nli+1:]
	}

	return sum, nil
}
