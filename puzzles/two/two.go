package two

import (
	"strings"
)

func Two(
	input string,
) (int, error) {

	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		handfuls := strings.Split(
			line[strings.Index(line, `:`)+1:],
			";",
		)
		if len(handfuls) == 0 {
			continue
		}
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

	}

	return sum, nil
}
