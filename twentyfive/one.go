package twentyfive

import "strings"

func One(
	input string,
) (string, error) {
	return toSnafu(sumSnafu(input)), nil
}

func sumSnafu(input string) int {
	total := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}
		total += toDecimal(input[:nli])
		input = input[nli+1:]
	}
	return total
}

func toDecimal(snafu string) int {
	v := 0
	pos := 1
	for i := len(snafu) - 1; i >= 0; i-- {
		switch snafu[i] {
		case '2':
			v += 2 * pos
		case '1':
			v += pos
		case '0':
			// do nothing
		case '-':
			v -= pos
		case '=':
			v -= 2 * pos
		default:
			panic(snafu[i])
		}

		pos *= 5
	}

	return v
}

func toSnafu(d int) string {
	if d == 0 {
		return `0`
	}

	output := ``
	prevPos := 1
	pos := 5
	total := 0
	for total != d {
		switch ((d - total) % pos) / prevPos {
		case 0:
			output = `0` + output
		case 1:
			output = `1` + output
			total += prevPos
		case 2:
			output = `2` + output
			total += prevPos + prevPos
		case 3:
			output = `=` + output
			total -= (prevPos + prevPos)
		case 4:
			output = `-` + output
			total -= prevPos
		default:
			panic(`ahh`)
		}
		prevPos = pos
		pos *= 5
	}

	switch ((d - total) % pos) / prevPos {
	case 0:
	case 1:
		output = `1` + output
	case 2:
		output = `2` + output
	default:
		panic(`ahh`)
	}

	return output
}
