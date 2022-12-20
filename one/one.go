package one

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	var val int
	var err error

	max := -1
	cur := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			if cur > max {
				max = cur
			}
			cur = 0
		} else {
			val, err = strconv.Atoi(input[0:nli])
			if err != nil {
				return 0, err
			}
			cur += val
		}
		input = input[nli+1:]
	}

	return max, nil
}
