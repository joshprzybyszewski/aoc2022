package twelve

const (
	maxNumGroups  = 32
	maxLineLength = 105
)

type possibilities struct {
	possibles [maxNumGroups][maxLineLength]int

	line       [maxLineLength]part
	lineLength int

	distToBroken [maxLineLength]int
	distToSafe   [maxLineLength]int

	groups    [maxNumGroups]int
	numGroups int
}

func newPossibilities(
	input string,
) (possibilities, string) {
	var p possibilities

	var i, cur int
	var addGroup bool

	for len(input) > 0 {
		if input[0] == '\n' {
			break
		}

		if addGroup {
			switch input[0] {
			case ',':
				// iterate past.
				p.groups[p.numGroups] = cur
				p.numGroups++
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
		return possibilities{}, input
	}
	if !addGroup {
		panic(`unexpected`)
	}
	if cur == 0 {
		panic(`unexpected`)
	}

	p.groups[p.numGroups] = cur
	p.numGroups++

	return p, input
}

func (p *possibilities) answer() int {
	total := 0
	for i := p.lineLength - 1; i >= 0; i-- {
		total += p.possibles[0][i]
	}
	return total
}

func (p *possibilities) findDistances() {
	toBroken := p.lineLength + 1
	toSafe := p.lineLength + 1
	for i := p.lineLength - 1; i >= 0; i-- {
		toBroken++
		toSafe++

		switch p.line[i] {
		case broken:
			toBroken = 0
		case safe:
			toSafe = 0
		}
		p.distToBroken[i] = toBroken
		p.distToSafe[i] = toSafe
	}
}

func (p *possibilities) build() {
	p.findDistances()
	p.buildSubGroup(0)
}

func (p *possibilities) buildSubGroup(
	gi int,
) {

	if gi >= p.numGroups {
		return
	}

	group := p.groups[gi]

	if gi == p.numGroups-1 {
		for i := p.lineLength - 1; i >= 0; i-- {
			if p.canPlace(i, group) {
				p.possibles[gi][i] = 1
			}
			if p.hasBrokenInRange(i+group, p.lineLength) {
				break
			}
		}
		return
	}
	p.buildSubGroup(gi + 1)

	for i := p.lineLength; i >= 0; i-- {

		prevVal := p.possibles[gi+1][i]
		if prevVal == 0 {
			continue
		}

		start := i - group
		for j := start - 1; j >= 0; j-- {
			if gi == 0 && (p.hasBrokenInRange(0, j-1) || p.hasBrokenInRange(j+group, i-1)) {
				continue
			}
			if p.canPlace(j, group) {
				p.possibles[gi][j] += prevVal
			}
			if p.hasBrokenInRange(j+group, i-1) {
				break
			}
		}
	}

}

func (p *possibilities) canPlace(
	startIndex int,
	group int,
) bool {

	if p.distToSafe[startIndex] < group {
		return false
	}

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
		// The piece after this group is broken; cannot place it ending here
		return false
	}

	return true
}

func (p *possibilities) hasBrokenInRange(
	startIndex, endIndex int,
) bool {
	// if p.distToBroken[startIndex] <= (endIndex - startIndex) {
	// 	return true
	// }
	for n := startIndex; n <= endIndex; n++ {
		if p.line[n] == broken {
			// panic(`should've been caught`)
			return true
		}
	}
	return false
}
