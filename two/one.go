package two

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		s, err := score(line)
		if err != nil {
			return ``, err
		}
		total += s
	}

	return fmt.Sprintf("%d", total), nil
}

func score(
	line string,
) (int, error) {
	if line == `` {
		return 0, nil
	}

	values := strings.Split(line, ` `)
	if len(values) != 2 {
		return 0, fmt.Errorf(`should have two values: %q`, line)
	}

	ss, err := shapeScore(values[1])
	if err != nil {
		return 0, err
	}

	ws, err := winScore(values[0], values[1])
	if err != nil {
		return 0, err
	}

	return ss + ws, nil
}

func shapeScore(
	encChar string,
) (int, error) {
	switch encChar {
	case `X`:
		return 1, nil
	case `Y`:
		return 2, nil
	case `Z`:
		return 3, nil
	}

	return 0, fmt.Errorf(`unsupported char: %q`, encChar)
}

func winScore(
	encChar1, encChar2 string,
) (int, error) {
	switch encChar2 {
	case `X`: // rock
		switch encChar1 {
		case `A`: // rock
			return 3, nil
		case `B`: // paper
			return 0, nil
		case `C`: // scissors
			return 6, nil
		}
	case `Y`: // paper
		switch encChar1 {
		case `A`: // rock
			return 6, nil
		case `B`: // paper
			return 3, nil
		case `C`: // scissors
			return 0, nil
		}
	case `Z`: // scissors
		switch encChar1 {
		case `A`: // rock
			return 0, nil
		case `B`: // paper
			return 6, nil
		case `C`: // scissors
			return 3, nil
		}
	}

	return 0, fmt.Errorf(`unsupported chars: %q %q`, encChar1, encChar2)
}
