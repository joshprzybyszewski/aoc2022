package four

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {

	var s1, e1, s2, e2, i1, i2 int
	var err error
	total := 0

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		// Sscanf takes 5000 allocs, but only 87000 B
		// strings.Split takes 3000 allocs, but 112000 B
		// Using Index takes 2 allocs, and 16387 B
		i1 = 0
		i2 = strings.Index(input, `-`)
		s1, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}
		i1 = i2 + 1
		i2 = strings.Index(input, `,`)
		e1, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}
		i1 = i2 + 1
		i2 = i1 + strings.Index(input[i1:], `-`)
		s2, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}
		i1 = i2 + 1
		i2 = nli
		e2, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}

		if (s1 >= s2 && e1 <= e2) ||
			(s2 >= s1 && e2 <= e1) {
			// the ranges are completely overlapping
			total++
		}
		input = input[nli+1:]
	}

	return total, nil
}
