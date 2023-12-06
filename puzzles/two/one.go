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

	i := 0
	sum := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		handfuls := strings.Split(
			input[strings.Index(input, `:`)+1:nli],
			";",
		)
		hasImpossible := false
		for _, handfulStr := range handfuls {
			handful := interpretSeen(handfulStr)
			if !isPossible(handful, max) {
				hasImpossible = true
				break
			}
		}

		if !hasImpossible {
			sum += (i + 1)
		}
		input = input[nli+1:]
		i++
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
			if output.red != 0 {
				panic(`already red set`)
			}
			output.red = val
		case `green`:
			if output.green != 0 {
				panic(`already green set`)
			}

			output.green = val
		case `blue`:
			if output.blue != 0 {
				panic(`already blue set`)
			}

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
