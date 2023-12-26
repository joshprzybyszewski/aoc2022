package twelve

import "fmt"

const (
	maxNumGroups  = 32
	maxLineLength = 105
)

type possibilities struct {
	line       [maxLineLength]part
	lineLength int

	possibles [maxNumGroups][maxLineLength]int
}

func newPossibilities(
	input string,
) (possibilities, []int, string) {
	var p possibilities

	var i, cur int
	var addGroup bool
	groups := make([]int, 0, 8)

	for len(input) > 0 {
		if input[0] == '\n' {
			break
		}

		if addGroup {
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
				p.line[i] = unknown
			case '#':
				p.line[i] = broken
			case ' ':
				p.lineLength = i
				addGroup = true
			}
			i++
		}
		input = input[1:]
	}

	if p.lineLength == 0 {
		return possibilities{}, nil, input
	}
	if !addGroup {
		panic(`unexpected`)
	}
	if cur == 0 {
		panic(`unexpected`)
	}

	groups = append(groups, cur)

	return p, groups, input
}

func (p *possibilities) answer(
	groups []int,
) int {
	total := 0
	for i := p.lineLength - 1; i >= 0; i-- {
		total += p.possibles[len(groups)][i]
	}
	return total
}

func (p *possibilities) build(
	groups []int,
) {
	if len(groups) > maxNumGroups {
		fmt.Printf("len(groups): %d\n", len(groups))
		panic(`unhandled`)
	}

	p.buildSubGroup(groups, true)
}

func (p *possibilities) buildSubGroup(
	groups []int,
	checkBefore bool,
) {

	if len(groups) == 0 {
		return
	}
	if len(groups) == 1 {
		group := groups[0]
		for i := p.lineLength - 1; i >= 0; i-- {
			if p.canPlace(i, group) {
				p.possibles[len(groups)][i] = 1
			}
			if p.hasBrokenInRange(i+group, p.lineLength) {
				break
			}
		}
		return
	}
	p.buildSubGroup(groups[1:], false)

	group := groups[0]
	for i := p.lineLength; i >= 0; i-- {

		prevVal := p.possibles[len(groups)-1][i]
		if prevVal == 0 {
			continue
		}

		start := i - group
		for j := start - 1; j >= 0; j-- {
			if checkBefore && p.hasBrokenInRange(0, j-1) {
				continue
			}
			if p.canPlace(j, group) {
				p.possibles[len(groups)][j] += prevVal
			}
			if p.hasBrokenInRange(j+group, start) {
				break
			}
		}
	}

}

func (p *possibilities) canPlace(
	startIndex int,
	group int,
) bool {
	if startIndex+group > p.lineLength {
		// extends beyond this line
		return false
	}
	if startIndex > 0 &&
		p.line[startIndex-1] == broken {
		// the piece before this group attempt is broken; cannot place it starting here.
		return false
	}

	if startIndex+group < p.lineLength &&
		p.line[startIndex+group] == broken {
		return false
	}

	for n := 0; n < group; n++ {
		if p.line[startIndex+n] == safe {
			return false
		}
	}
	return true
}

func (p *possibilities) hasBrokenInRange(
	startIndex, endIndex int,
) bool {
	for n := startIndex; n <= endIndex; n++ {
		if p.line[n] == broken {
			return true
		}
	}
	return false
}
