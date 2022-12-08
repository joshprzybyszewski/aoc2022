package four

import (
	"fmt"
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var s1, e1, s2, e2 int
	var err error
	total := 0

	for _, line := range lines {
		if line == `` {
			continue
		}
		_, err = fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		if err != nil {
			return ``, err
		}
		if (s1 >= s2 && s1 <= e2) ||
			(e1 >= s2 && e1 <= e2) ||
			(s2 >= s1 && s2 <= e1) ||
			(e2 >= s1 && e2 <= e1) {
			// the ranges are overlapping
			total++
		}
	}

	return strconv.Itoa(total), nil
}
