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

func One(
	input string,
) (int, error) {
	total := 0
	var i, cur int
	var r row
	var addGroup bool
	groups := make([]int, 0, 8)
	for len(input) > 0 {
		if input[0] == '\n' {
			if r.numParts > 0 {
				groups = append(groups, cur)
				// fmt.Printf("  %20s %v\n", r, groups)
				num := getNum(r, 0, groups)
				// fmt.Printf("ANSWER: %d\n\n\n", num)
				total += num
			}
			i = 0
			cur = 0
			r = row{}
			addGroup = false
			groups = groups[:0]
		} else if addGroup {
			switch input[0] {
			case ',':
				// iterate past.
				groups = append(groups, cur)
				cur = 0
			default:
				cur *= 10
				cur += int(input[0] - '0')
			}
		} else {
			switch input[0] {
			case '?':
				r.parts[i] = unknown
			case '#':
				r.parts[i] = broken
			case ' ':
				r.numParts = i
				addGroup = true
			}
			i++
		}
		input = input[1:]
	}
	return total, nil
}

func getNum(
	r row,
	start int,
	groups []int,
) int {
	// fmt.Printf("  %20s %v\tstart = %2d\n", r, groups, start)
	if len(groups) == 0 {
		// fmt.Printf("  %20s %v\tstart = %2d\n", r, groups, start)
		return 1
	}

	total := 0
	maxI := r.numParts
	{ // limit the max starting point
		if len(groups) > 0 {
			maxI -= (len(groups) - 1)
		}
		for _, g := range groups {
			maxI -= g
		}
	}

	for i := start; i <= maxI; i++ {
		if canPlace(r, i, groups) {
			rCpy := r
			for j := i; j < i+groups[0]; j++ {
				rCpy.parts[j] = broken
			}
			if i+groups[0] < r.numParts {
				rCpy.parts[i+groups[0]] = safe
			}
			total += getNum(rCpy, i+groups[0]+1, groups[1:])

		}
		if r.parts[i] == broken {
			// TODO this can be faster
			break
		}
	}

	return total
}

func canPlace(
	r row,
	start int,
	groups []int,
) bool {
	maxI := start + groups[0]
	if maxI > r.numParts {
		return false
	}
	for i := start; i < maxI; i++ {
		if r.parts[i] == safe {
			return false
		}
	}
	if len(groups) == 1 {
		for i := maxI; i < r.numParts; i++ {
			if r.parts[i] == broken {
				return false
			}
		}
	}
	return maxI == r.numParts || r.parts[maxI] != broken
}

func hasBroken(
	r row,
	start int,
) bool {
	for i := start; i < r.numParts; i++ {
		if r.parts[i] == broken {
			return true
		}
	}
	return false
}
