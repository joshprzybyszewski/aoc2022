package two

import (
	"strings"
)

func Two(
	input string,
) (int, error) {

	sum := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		handfuls := strings.Split(
			input[strings.Index(input, `:`)+1:nli],
			";",
		)
		min := interpretSeen(handfuls[0])
		for _, handfulStr := range handfuls {
			handful := interpretSeen(handfulStr)
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

		sum += (min.red * min.blue * min.green)
		input = input[nli+1:]
	}

	return sum, nil
}
