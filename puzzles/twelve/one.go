package twelve

import (
	"strings"
)

type part uint8

const (
	safe    part = 0
	broken  part = 1
	unknown part = 2
)

func (p part) toString() byte {
	switch p {
	case safe:
		return '.'
	case broken:
		return '#'
	case unknown:
		return '?'
	}
	return 'X'
}

type row struct {
	parts    [20]part
	numParts int
}

func (r row) String() string {
	var sb strings.Builder
	for i := 0; i < r.numParts; i++ {
		sb.WriteByte(r.parts[i].toString())
	}
	return sb.String()
}

func (r row) isSolution(indexes []int) bool {
	ii := 0
	cur := 0
	for i := 0; i < len(r.parts); i++ {
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

func (r row) getPossibilities(indexes []int) int {
	// fmt.Printf("-- %s %v\n", r, indexes)
	total := solveNext(r, 0, indexes)
	// fmt.Printf("   %d\n", total)
	return total
}

func solveNext(
	r row,
	i int,
	indexes []int,
) int {
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

	return solveNext(r1, i, indexes) + solveNext(r, i, indexes)
}

func getNumConfigurations(line string) int {
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

	return r.getPossibilitiesV2(indexes)
}

func One(
	input string,
) (int, error) {
	total := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		total += getNumConfigurations(input[:nli])
		input = input[nli+1:]
	}
	return total, nil
}

func (r row) isSkipPossibleV2(
	skip int,
) bool {

	var i int
	for i = 0; i < skip; i++ {
		if r.parts[i] == broken {
			return false
		}
	}

	return true
}

func (r row) isPossibleV2(
	skip int,
	busted []int,
	between []int,
	maxBI int,
) bool {

	var i int
	for i = 0; i < skip; i++ {
		if r.parts[i] == broken {
			return false
		}
	}

	var n int
	for bi := 0; bi < maxBI && bi < len(between); bi++ {
		for n = 0; n < busted[bi] && i < r.numParts; n++ {
			if r.parts[i] == safe {
				return false
			}
			i++
		}
		for n = 0; n < between[bi] && i < r.numParts; n++ {
			if r.parts[i] == broken {
				return false
			}
			i++
		}
		if i >= r.numParts {
			return false
		}
	}

	if maxBI < len(busted) {
		return true
	}

	for n = 0; n < busted[len(busted)-1]; n++ {
		if i >= r.numParts {
			return false
		}

		if r.parts[i] == safe {
			return false
		}
		i++
	}

	for ; i < r.numParts; i++ {
		if r.parts[i] == broken {
			return false
		}
	}
	return i == r.numParts
}

func (r row) getPossibilitiesV2(
	busted []int,
) int {
	between := make([]int, len(busted)-1)
	maxSkip := r.numParts - len(between)
	for _, b := range busted {
		maxSkip -= b
	}

	numPossible := 0
	for skip := 0; skip <= maxSkip; skip++ {
		if !r.isSkipPossibleV2(skip) {
			continue
		}

		numPossible += r.getPossibilitiesV2_between(skip, maxSkip, busted, between, 0)
	}

	return numPossible
}

func (r row) getPossibilitiesV2_between(
	skip, maxSkip int,
	busted, between []int,
	bi int,
) int {
	if bi >= len(between) {
		if r.isPossibleV2(skip, busted, between, len(busted)) {
			return 1
		}
		return 0
	}

	if !r.isPossibleV2(skip, busted, between, bi) {
		return 0
	}

	myBetween := make([]int, len(between))
	copy(myBetween, between)

	numPossible := 0
	for v := 1; v <= maxSkip; v++ {
		myBetween[bi] = v

		numPossible += r.getPossibilitiesV2_between(
			skip,
			maxSkip,
			busted,
			myBetween,
			bi+1,
		)
	}
	return numPossible
}
