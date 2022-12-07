package three

import (
	"fmt"
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	total := 0
	for i := 0; i+3 <= len(lines); i += 3 {
		c, err := commonLetter(lines[i : i+3])
		if err != nil {
			return ``, err
		}
		total += priority(c)
	}

	return strconv.Itoa(total), nil
}

func commonLetter(
	lines []string,
) (rune, error) {
	for _, c := range []rune(lines[0]) {
		if !strings.Contains(lines[1], string(c)) {
			continue
		}
		if !strings.Contains(lines[2], string(c)) {
			continue
		}
		return c, nil
	}
	return 0, fmt.Errorf("couldn't find common element: %q %q %q", lines[0], lines[1], lines[2])
}
