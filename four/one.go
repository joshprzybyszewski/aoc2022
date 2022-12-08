package four

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var s1, e1, s2, e2, i1, i2 int
	var err error
	total := 0

	for _, line := range lines {
		if line == `` {
			continue
		}

		// Sscanf takes 5000 allocs, but only 87000 B
		// strings.Split takes 3000 allocs, but 112000 B
		// Using Index takes 2 allocs, and 16387 B
		i1 = 0
		i2 = strings.Index(line, `-`)
		s1, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return ``, err
		}
		i1 = i2 + 1
		i2 = strings.Index(line, `,`)
		e1, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return ``, err
		}
		i1 = i2 + 1
		i2 = strings.LastIndex(line, `-`)
		s2, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return ``, err
		}
		i1 = i2 + 1
		i2 = len(line)
		e2, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return ``, err
		}

		if (s1 >= s2 && e1 <= e2) ||
			(s2 >= s1 && e2 <= e1) {
			// the ranges are overlapping
			total++
		}
	}

	return strconv.Itoa(total), nil
}
