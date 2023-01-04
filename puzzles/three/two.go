package three

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	var c byte
	var err error
	total := 0
	for i := 0; i+3 <= len(lines); i += 3 {
		c, err = commonLetter(lines[i : i+3])
		if err != nil {
			return 0, err
		}
		total += priority(c)
	}

	return total, nil
}

func commonLetter(
	lines []string,
) (byte, error) {
	for i := range lines[0] {
		if !strings.ContainsRune(lines[1], rune(lines[0][i])) {
			continue
		}
		if !strings.ContainsRune(lines[2], rune(lines[0][i])) {
			continue
		}
		return lines[0][i], nil
	}
	return 0, fmt.Errorf("couldn't find common element: %q %q %q", lines[0], lines[1], lines[2])
}
