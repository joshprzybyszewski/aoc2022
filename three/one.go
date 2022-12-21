package three

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (int, error) {
	var c byte
	var err error
	total := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}
		c, err = common(input[0:nli])
		if err != nil {
			return 0, err
		}
		total += priority(c)
		input = input[nli+1:]
	}

	return total, nil
}

func common(
	line string,
) (byte, error) {
	half := len(line) / 2

	for i := 0; i < half; i++ {
		if strings.ContainsRune(line[half:], rune(line[i])) {
			return line[i], nil
		}
	}

	return 0, fmt.Errorf("couldn't find common element: %q %q", line[:half], line[half:])
}

func priority(
	c byte,
) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	return int(c-'A') + 27
}
