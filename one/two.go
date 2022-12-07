package one

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var elves []int
	cur := 0
	for _, line := range lines {
		if line == `` {
			elves = append(elves, cur)
			cur = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return ``, err
		}
		cur += val
	}
	sort.Ints(elves)
	total := 0
	for e := 1; e <= 3; e++ {
		i := len(elves) - e
		if i < 0 {
			break
		}
		total += elves[i]
	}

	return fmt.Sprintf("%d", total), nil
}
