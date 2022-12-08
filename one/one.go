package one

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
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
			return ``, err
		}
		cur += val
	}

	return strconv.Itoa(max), nil
}
