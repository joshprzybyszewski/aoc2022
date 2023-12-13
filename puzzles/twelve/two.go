package twelve

import (
	"fmt"
	"strings"
	"time"
)

type unfoldedRow struct {
	parts    [105]part
	numParts int
}

func newUnfoldedRow(r row) unfoldedRow {
	// fmt.Printf("r: %s\n", r)
	// fmt.Printf("r.numParts = %d\n", r.numParts)
	ur := unfoldedRow{
		numParts: r.numParts*5 + 4,
	}
	uri := 0
	for i := 0; i < 5; i++ {
		for pi := 0; pi < r.numParts; pi++ {
			ur.parts[uri] = r.parts[pi]
			uri++
		}
		ur.parts[uri] = unknown
		uri++
	}
	return ur
}

func (r unfoldedRow) String() string {
	var sb strings.Builder
	for i := 0; i < r.numParts; i++ {
		sb.WriteByte(r.parts[i].toString())
	}
	return sb.String()
}

func (r unfoldedRow) getPossibilities(indexes []int) int {
	fmt.Printf("-- %s %v\n", r, indexes)
	total := solveNextUnfolded(r, 0, indexes)
	fmt.Printf("   %d\n", total)
	return total
}

func (r unfoldedRow) cannotBeSolution(
	indexes []int,
) bool {
	ii := 0
	cur := 0
	for i := 0; i < r.numParts; i++ {
		if r.parts[i] == unknown {
			return false
		}
		if r.parts[i] == broken {
			cur++
		} else if cur > 0 {
			if ii >= len(indexes) || cur != indexes[ii] {
				return true
			}
			ii++
			cur = 0
		}
	}
	return false
}

func (r unfoldedRow) isSolution(indexes []int) bool {
	ii := 0
	cur := 0
	for i := 0; i < r.numParts; i++ {
		if r.parts[i] == broken {
			cur++
		} else if cur > 0 {
			if ii >= len(indexes) || cur != indexes[ii] {
				return false
			}
			ii++
			cur = 0
		}
	}
	if ii < len(indexes) {
		if cur != indexes[ii] {
			return false
		}
		cur = 0
		ii++
	}
	return ii == len(indexes) && cur == 0
}

func solveNextUnfolded(
	r unfoldedRow,
	i int,
	indexes []int,
) int {
	if r.cannotBeSolution(indexes) {
		return 0
	}

	for i < len(r.parts) && r.parts[i] != unknown {
		i++
	}

	if i >= len(r.parts) {
		if r.isSolution(indexes) {
			// fmt.Printf("   %s\n", r)
			return 1
		}
		return 0
	}

	r1 := r
	r1.parts[i] = broken
	r.parts[i] = safe
	i++

	return solveNextUnfolded(r1, i, indexes) + solveNextUnfolded(r, i, indexes)
}

func Two(
	input string,
) (int, error) {
	return -42069, nil
	total := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		t0 := time.Now()
		total += getNumUnfoldedConfigurations(input[:nli])
		fmt.Printf("Took: %s\n\n", time.Since(t0))
		input = input[nli+1:]
	}
	return total, nil
}

func getNumUnfoldedConfigurations(line string) int {
	var r row
	var i int
	for i = 0; i < len(line); i++ {
		if line[i] == ' ' {
			r.numParts = i
			i++
			break
		}
		switch line[i] {
		case '?':
			r.parts[i] = unknown
		case '#':
			r.parts[i] = broken
		}
	}

	curNum := 0
	var indexes []int
	for ; i < len(line); i++ {
		if line[i] == ',' {
			indexes = append(indexes, curNum)
			curNum = 0
			continue
		}
		curNum *= 10
		curNum += int(line[i] - '0')
	}
	indexes = append(indexes, curNum)

	unfoldedIndexes := make([]int, 0, len(indexes)*5)

	for i := 0; i < 5; i++ {
		unfoldedIndexes = append(unfoldedIndexes, indexes...)
	}

	ur := newUnfoldedRow(r)
	fmt.Printf("r: %s %v\n", ur, unfoldedIndexes)

	return ur.getPossibilitiesV2(unfoldedIndexes)
}

func (r unfoldedRow) getPossibilitiesV2(
	busted []int,
) int {

	maxToSkip := r.numParts + 1 - len(busted)
	for bi := range busted {
		maxToSkip -= busted[bi]
	}

	total := 0

	for toSkip := 0; toSkip <= maxToSkip; toSkip++ {
		total += r.getPossibilitiesV2_recursive(0, toSkip, busted)
	}

	return total
}

func (r unfoldedRow) getPossibilitiesV2_recursive(
	minUncheckedIndex int,
	toSkip int,
	remainingBusted []int,
) int {
	if !r.isPossibleV2(minUncheckedIndex, toSkip, remainingBusted) {
		return 0
	}

	minUncheckedIndex += toSkip + remainingBusted[0]
	remainingBusted = remainingBusted[1:]

	if len(remainingBusted) == 0 {
		for ; minUncheckedIndex < r.numParts; minUncheckedIndex++ {
			if r.parts[minUncheckedIndex] == broken {
				return 0
			}
		}
		return 1
	}

	numberOfUnknown := 0
	for i := minUncheckedIndex; i < r.numParts && r.parts[i] == unknown; i++ {
		numberOfUnknown++
	}
	if minUncheckedIndex+numberOfUnknown < r.numParts && r.parts[minUncheckedIndex+numberOfUnknown] == broken {
		// not gonna figure out how to merg into this.
		numberOfUnknown--
	}

	// TODO
	// ???   [1,1] -> ??
	// ????  [1,1] -> ??
	// ???   [1,2] -> ???
	// ????  [1,2] -> ??
	// ????? [1,2] -> ??
	// ???   [2,1] -> ???
	// ????  [2,1] -> ??
	// ????? [2,1] -> ??

	// for i := 1; i < len(remainingBusted); i++ {
	// 	if numberOfUnknown - remainingBusted[i] -1
	// }

	// TODO we _shouldn't_ take this branch until I figure out the above.
	if numberOfUnknown > remainingBusted[0] {
		numPossible := r.getPossibilitiesV2_recursive(
			minUncheckedIndex,
			numberOfUnknown-remainingBusted[0],
			remainingBusted,
		)

		mult := numberOfUnknown - remainingBusted[0] + 1
		// ?    [1] -> 1
		// ??   [1] -> 2
		// ???  [1] -> 3
		// ???? [1] -> 4
		// ??   [2] -> 1
		// ???  [2] -> 2
		// ???? [2] -> 3
		// ???  [3] -> 1
		// ???? [3] -> 2

		return numPossible * mult
	}

	numPossible := 0
	maxToSkip := r.numParts - minUncheckedIndex - len(remainingBusted) + 1
	for bi := range remainingBusted {
		maxToSkip -= remainingBusted[bi]
	}

	for toSkip = 1; toSkip <= maxToSkip; toSkip++ {
		numPossible += r.getPossibilitiesV2_recursive(
			minUncheckedIndex,
			toSkip,
			remainingBusted,
		)
	}
	return numPossible
}

func (r unfoldedRow) isPossibleV2(
	minUncheckedIndex int,
	toSkip int,
	remainingBusted []int,
) bool {

	var n int

	i := minUncheckedIndex

	if i+toSkip+remainingBusted[0] > r.numParts {
		return false
	}
	for n = 0; n < toSkip; n++ {
		if r.parts[i] == broken {
			return false
		}
		i++
	}

	for n = 0; n < remainingBusted[0]; n++ {
		if r.parts[i] == safe {
			return false
		}
		i++
	}

	for n = 1; n < len(remainingBusted); n++ {
		i += 1 + remainingBusted[n]
	}

	return i <= r.numParts
}
