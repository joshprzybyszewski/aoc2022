package thirteen

func fixSmudge(a, b uint64) bool {
	if a == b {
		panic(`dev error`)
		// return -1
	}
	c := uint64(a ^ b)

	for tmp := uint64(1); tmp <= c; tmp <<= 1 {
		if c == tmp {
			return true
		}

	}
	return false
}

func (r rect) isAlmostReflection(maxVal int, isHorizontal bool) bool {
	cmpVal := maxVal + 1
	if isHorizontal {
		if cmpVal >= r.numRows {
			return false
		}

		found := false

		for maxVal >= 0 && cmpVal < r.numRows {
			if r.rows[maxVal] != r.rows[cmpVal] {
				if !fixSmudge(r.rows[maxVal], r.rows[cmpVal]) {
					return false
				}
				if found {
					// two smudges
					return false
				}
				found = true
			}
			maxVal--
			cmpVal++
		}
		return found
	}

	if cmpVal >= r.numCols {
		return false
	}

	found := false

	for maxVal >= 0 && cmpVal < r.numCols {
		if r.cols[maxVal] != r.cols[cmpVal] {
			if !fixSmudge(r.cols[maxVal], r.cols[cmpVal]) {
				return false
			}
			if found {
				// two smudges
				return false
			}
			found = true
		}
		maxVal--
		cmpVal++
	}
	return found
}

func (r rect) getNumBeforeReflectionAfterFixingSmudge() int {
	for ri := 0; ri < r.numRows; ri++ {
		if r.isAlmostReflection(ri, true) {
			return 100 * (ri + 1)
		}
	}

	for ci := 0; ci < r.numCols; ci++ {
		if r.isAlmostReflection(ci, false) {
			return ci + 1
		}
	}
	panic(`dev error`)
	return 0
}

func Two(
	input string,
) (int, error) {

	sum := 0
	var r rect

	for len(input) > 0 {
		if input == "\n" {
			break
		}
		input, r = newRect(input)
		sum += r.getNumBeforeReflectionAfterFixingSmudge()
	}

	return sum, nil
}
