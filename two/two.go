package two

import (
	"fmt"
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var s int
	var err error

	total := 0
	for _, line := range lines {
		s, err = score2(line)
		if err != nil {
			return ``, err
		}
		total += s
	}

	return strconv.Itoa(total), nil
}

func score2(
	line string,
) (int, error) {
	if line == `` {
		return 0, nil
	}

	ss, err := shapeScore2(line[0], line[2])
	if err != nil {
		return 0, err
	}

	ws, err := winScore2(line[2])
	if err != nil {
		return 0, err
	}

	return ss + ws, nil
}

func shapeScore2(
	encChar,
	encOutcome byte,
) (int, error) {

	opp := 0
	switch encChar {
	case 'A': // rock
		opp = 1
	case 'B': // paper
		opp = 2
	case 'C': // scissors
		opp = 3
	default:
		return 0, fmt.Errorf(`unsupported char: %q`, encChar)
	}

	switch encOutcome {
	case 'X': // lose
		// i'm sure i could modulo this
		mine := opp - 1
		if mine == 0 {
			return 3, nil
		}
		return mine, nil
	case 'Y': // draw
		return opp, nil
	case 'Z': // win
		// i'm sure i could modulo this
		mine := opp + 1
		if mine == 4 {
			return 1, nil
		}
		return mine, nil
	}

	return 0, fmt.Errorf(`unsupported char: %q %q`, encChar, encOutcome)
}

func winScore2(
	encOutcome byte,
) (int, error) {
	switch encOutcome {
	case 'X': // lose
		return 0, nil
	case 'Y': // draw
		return 3, nil
	case 'Z': // win
		return 6, nil
	}

	return 0, fmt.Errorf(`unsupported chars: %q`, encOutcome)
}
