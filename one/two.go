package one

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	var val int
	var err error
	elves := make([]int, 0, 236)
	cur := 0
	for _, line := range lines {
		if line == `` {
			elves = append(elves, cur)
			cur = 0
			continue
		}

		val, err = strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		cur += val
	}
	top3 := [3]int{-1. - 1, -1}

	for _, e := range elves {
		if e > top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = e
		} else if e > top3[1] {
			top3[2] = top3[1]
			top3[1] = e
		} else if e > top3[2] {
			top3[2] = e
		}
	}

	return top3[0] + top3[1] + top3[2], nil
}
