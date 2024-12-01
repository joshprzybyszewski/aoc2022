package day01

import (
	"github.com/joshprzybyszewski/aoc2022/util/lines"
	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

func Two(
	input string,
) (int, error) {

	left := make([]int, 0, 1000)
	right := make(map[int]int, 1000)

	lines.ForEach(input,
		func(line string) {
			l, n := strutil.IntBeforeSpace(line)
			r, _ := strutil.IntTrimSpace(line[n:])
			left = append(left, l)
			right[r]++
		},
	)

	sum := 0
	for _, l := range left {
		sum += l * right[l]
	}

	return sum, nil
}
