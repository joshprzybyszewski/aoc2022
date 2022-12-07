package one

import (
	"fmt"
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

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

		val, err := strconv.Atoi(line)
		if err != nil {
			return ``, err
		}
		cur += val
	}

	return fmt.Sprintf("%d", max), nil
}
