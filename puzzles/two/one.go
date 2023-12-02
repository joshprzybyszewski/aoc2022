package two

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	max := handful{
		red:   12,
		green: 13,
		blue:  14,
	}

	lines := strings.Split(input, "\n")

	sum := 0
	for i, line := range lines {

		handfuls := strings.Split(
			line[strings.Index(line, `:`)+1:],
			";",
		)
		for _, handfulStr := range handfuls {
			handful := interpretSeen(handfulStr)
			if !isPossible(handful, max) {
				sum += (i + 1)
			}
		}

	}

	return sum, nil
}

func interpretSeen(handfulString string) handful {
	values := strings.Split(
		strings.TrimSpace(handfulString),
		",",
	)

	output := handful{}
	for _, value := range values {
		infos := strings.Split(strings.TrimSpace(value), " ")
		if len(infos) != 2 {
			panic(`ahhh: ` + value)
		}
		val, err := strconv.Atoi(infos[0])
		if err != nil {
			panic(err)
		}
		switch infos[1] {
		case `red`:
			output.red = val
		case `green`:
			output.green = val
		case `blue`:
			output.blue = val
		default:
			panic(`unknown color`)
		}
	}
	return output
}

type handful struct {
	red   int
	blue  int
	green int
}

func isPossible(
	seen, max handful,
) bool {
	return seen.red <= max.red &&
		seen.green <= max.green &&
		seen.blue <= max.blue
}
