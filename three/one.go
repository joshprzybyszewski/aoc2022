package three

import (
	"fmt"
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		if line == `` {
			continue
		}
		c, err := common(line)
		if err != nil {
			return ``, err
		}
		total += priority(c)
	}

	return strconv.Itoa(total), nil
}

func common(
	line string,
) (rune, error) {
	half := len(line) / 2
	one := line[:half]
	two := line[half:]

	for _, c := range []rune(one) {
		if strings.Contains(two, string(c)) {
			return c, nil
		}
	}

	return 0, fmt.Errorf("couldn't find common element: %q %q", one, two)
}

func priority(
	c rune,
) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	return int(c-'A') + 27
}
