package day01

import (
	"slices"

	"github.com/joshprzybyszewski/aoc2022/util/lines"
	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

func One(
	input string,
) (int, error) {

	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	lines.ForEach(input,
		func(line string) {
			l, n := strutil.IntBeforeSpace(line)
			r, _ := strutil.IntTrimSpace(line[n:])
			left = append(left, l)
			right = append(right, r)
		},
	)

	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			sum -= diff
		} else {
			sum += diff
		}
	}
	return sum, nil
}
