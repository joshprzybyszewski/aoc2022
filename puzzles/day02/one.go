package day02

import (
	"github.com/joshprzybyszewski/aoc2022/util/lines"
	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

func One(
	input string,
) (int, error) {

	numSafe := 0
	var l1, l2, n int
	var isAscending bool
	lines.ForEach(input, func(line string) {
		l1, n = strutil.IntBeforeSpace(line)
		line = line[n+1:]
		l2, n = strutil.IntBeforeSpace(line)
		isAscending = l2 > l1
		if isAscending {
			for {
				if l2 <= l1 || l2 > l1+3 {
					// isSafe = false
					return
				}
				if n >= len(line) {
					break
				}
				line = line[n+1:]
				l1 = l2
				l2, n = strutil.IntBeforeSpace(line)
			}
			numSafe++
			return
		}

		// is descending
		for {
			if l2 >= l1 || l2 < l1-3 {
				// isSafe = false
				return
			}
			if n >= len(line) {
				break
			}
			line = line[n+1:]
			l1 = l2
			l2, n = strutil.IntBeforeSpace(line)
		}
		numSafe++
	})

	return numSafe, nil
}
