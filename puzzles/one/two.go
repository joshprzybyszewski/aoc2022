package one

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (int, error) {

	var val int
	var err error
	elves := make([]int, 0, 236)
	cur := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			elves = append(elves, cur)
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
	top3 := [3]int{-1, -1, -1}

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
