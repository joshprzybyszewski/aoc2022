package one

import (
	"strings"
)

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		sum += getValue([]byte(line))
	}

	return sum, nil
}

func getValue(line []byte) int {
	first, last := -1, -1

	for _, c := range line {
		if c >= '0' && c <= '9' {
			if first == -1 {
				first = int(c - '0')
			}
			last = int(c - '0')
			continue
		}

	}
	if first == -1 {
		return 0
	}

	return first*10 + last
}
