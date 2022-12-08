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

	var c byte
	var err error
	total := 0
	for _, line := range lines {
		if line == `` {
			continue
		}
		c, err = common(line)
		if err != nil {
			return ``, err
		}
		total += priority(c)
	}

	return strconv.Itoa(total), nil
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
