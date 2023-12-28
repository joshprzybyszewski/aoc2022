package one

import (
	"strings"
)

func One(
	input string,
) (int, error) {

	sum := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		sum += getValue(input[:nli])
		input = input[nli+1:]
	}

	return sum, nil
}

func getValue(line string) int {
	first := -1

	var i int
	var c byte
	for i = 0; i < len(line); i++ {
		c = line[i]
		if c >= '0' && c <= '9' {
			first = int(c - '0')
			break
		}
	}
	if first == -1 {
		return 0
	}

	last := first
	for i = len(line) - 1; i >= 0; i-- {
		c = line[i]
		if c >= '0' && c <= '9' {
			last = int(c - '0')
			break
		}
	}

	return first*10 + last
}
