package twelve

import (
	"fmt"
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
	parts    [105]part
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
				// fmt.Printf("  %-105s %v\n", r, groups)
				num := getNum(
					r,
					0,
					groups,
					getRemainingRequired(groups),
				)
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

func getRemainingRequired(groups []int) int {
	total := 0
	if len(groups) > 0 {
		total += (len(groups) - 1)
	}
	for _, g := range groups {
		total += g
	}
	return total
}

func getNum(
	r row,
	start int,
	groups []int,
	remainingRequired int,
) int {
	// fmt.Printf("  %-105s %v\tstart = %2d\n", r, groups, start)
	if len(groups) == 0 {
		// fmt.Printf("  %-105s %v\tstart = %2d\n", r, groups, start)
		return 1
	}

	total := 0
	maxI := r.numParts
	{ // limit the max starting point
		// maxI -= remainingRequired
		if len(groups) > 0 {
			maxI -= (len(groups) - 1)
		}
		for _, g := range groups {
			maxI -= g
		}
	}

	var can bool
	var nextI int
	for i := start; i <= maxI; i++ {
		can, nextI = canPlace(r, i, groups)
		if can {
			if nextI == 1000 {
				fmt.Printf("nextI:%d\n", nextI)
			}
			// r := markGroup(r, i, groups[0])
			total += getNum(r, i+groups[0]+1, groups[1:], remainingRequired-groups[0]-1)
		}
		if r.parts[i] == broken {
			break
		}
	}

	return total
}

func canPlace(
	r row,
	start int,
	groups []int,
) (bool, int) {
	maxI := start + groups[0]
	if maxI > r.numParts {
		// return false, r.numParts
		return false, start + 1
	}
	for i := start; i < maxI; i++ {
		if r.parts[i] == safe {
			// for i < r.numParts && r.parts[i] == safe {
			// 	i++
			// }
			// return false, i + 1
			return false, start + 1
		}
	}
	if len(groups) == 1 {
		for i := maxI; i < r.numParts; i++ {
			if r.parts[i] == broken {
				// return false, r.numParts
				return false, start + 1
			}
		}
	}
	if maxI == r.numParts {
		return true, start + 1
	}

	return r.parts[maxI] != broken, start + 1

}

func markGroup(
	r row,
	start int,
	group int,
) row {
	for j := start; j < start+group; j++ {
		r.parts[j] = broken
	}
	if start+group < r.numParts {
		r.parts[start+group] = safe
	}
	return r
}
