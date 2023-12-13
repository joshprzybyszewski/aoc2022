package twelve

import (
	"fmt"
	"strings"
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
		total += getNumUnfoldedConfigurations(input[:nli])
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

	return ur.getPossibilities(unfoldedIndexes)
}
