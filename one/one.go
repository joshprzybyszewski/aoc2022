package one

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	var val int
	var err error

	max := -1
	cur := 0
	for _, line := range lines {
		if line == `` {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}

		val, err = strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		cur += val
	}

	return max, nil
}
