package day02

import (
	"github.com/joshprzybyszewski/aoc2022/util/lines"
	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

func Two(
	input string,
) (int, error) {

	numSafe := 0
	var l, n int
	vals := make([]int, 0, 12)
	lines.ForEach(input, func(line string) {
		vals = vals[:0]
		for {
			l, n = strutil.IntBeforeSpace(line)
			vals = append(vals, l)
			if n >= len(line) {
				break
			}
			line = line[n+1:]
		}
		if isPart2Safe(vals) {
			numSafe++
		}
	})

	return numSafe, nil
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func isPart2Safe(vals []int) bool {
	valsCpy := make([]int, len(vals)-1)
	copy(valsCpy, vals[1:])

	for i := 0; i < len(valsCpy); i++ {
		if isPart1Safe(valsCpy) {
			return true
		}
		valsCpy[i] = vals[i]
	}
	return isPart1Safe(valsCpy)
}

func isPart1Safe(vals []int) bool {
	if vals[1] > vals[0] {
		for i := 1; i < len(vals); i++ {
			if vals[i] <= vals[i-1] || vals[i] > vals[i-1]+3 {
				// isSafe = false
				return false
			}
		}
		return true
	}
	if vals[1] == vals[0] {
		return false
	}
	for i := 1; i < len(vals); i++ {
		if vals[i] >= vals[i-1] || vals[i] < vals[i-1]-3 {
			// isSafe = false
			return false
		}
	}
	return true
}
