package thirteen

import (
	"strings"
)

type rect struct {
	rows [20]uint64
	cols [20]uint64

	numRows int
	numCols int
}

func (r rect) isReflection(maxVal int, isHorizontal bool) bool {
	cmpVal := maxVal + 1
	if isHorizontal {

		if cmpVal >= r.numRows {
			return false
		}

		for maxVal >= 0 && cmpVal < r.numRows {
			if r.rows[maxVal] != r.rows[cmpVal] {
				return false
			}
			maxVal--
			cmpVal++
		}
		return true
	}

	if cmpVal >= r.numCols {
		return false
	}

	for maxVal >= 0 && cmpVal < r.numCols {
		if r.cols[maxVal] != r.cols[cmpVal] {
			return false
		}
		maxVal--
		cmpVal++
	}
	return true
}

func (r rect) getNumBeforeReflection() int {
	for ri := 0; ri < r.numRows; ri++ {
		if r.isReflection(ri, true) {
			return 100 * (ri + 1)
		}
	}

	for ci := 0; ci < r.numCols; ci++ {
		if r.isReflection(ci, false) {
			return ci + 1
		}
	}
	panic(`dev error`)
	return 0
}

func newRect(
	input string,
) (string, rect) {

	var r rect
	ri, ci := 0, 0

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			r.numRows = ri
			return input[nli+1:], r
		}

		if ri == 0 {
			r.numCols = nli
		} else if r.numCols != nli {
			panic(`dev error`)
		}

		for ci = 0; ci < nli; ci++ {
			if input[ci] == '.' {
				continue
			}
			// TODO we can do fewer bit shifts if we keep track of this var elsewhere probably.
			r.rows[ri] |= 1 << ci
			r.cols[ci] |= 1 << ri
		}

		ri++
		input = input[nli+1:]
	}

	r.numRows = ri
	return input, r
}

func One(
	input string,
) (int, error) {

	sum := 0
	var r rect

	for len(input) > 0 {
		if input == "\n" {
			break
		}
		input, r = newRect(input)
		sum += r.getNumBeforeReflection()
	}

	return sum, nil
}
