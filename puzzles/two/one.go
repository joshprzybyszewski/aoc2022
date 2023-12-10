package two

import (
	"strings"

	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

const (
	newline = "\n"
	comma   = ","
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

func interpretSeen(
	input string,
) handful {

	output := handful{}
	var val, valEndIndex, ci int
	for {
		input = strutil.TrimSpaces(input)
		ci = strings.Index(input, comma)
		if ci == -1 {
			ci = len(input)
		}

		val, valEndIndex = strutil.IntBeforeSpace(input[:ci])
		valEndIndex++

		if input[valEndIndex:ci] == `red` {
			if output.red != 0 {
				panic(`already red set`)
			}
			output.red = val
		} else if input[valEndIndex:ci] == `blue` {
			if output.blue != 0 {
				panic(`already blue set`)
			}
			output.blue = val
		} else if input[valEndIndex:ci] == `green` {
			if output.green != 0 {
				panic(`already green set`)
			}
			output.green = val
		} else {
			panic(`unknown line: ` + input[:ci])
		}
		if ci == len(input) {
			return output
		}
		input = input[ci+1:]
	}

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
