package seventeen

import (
	"slices"
)

const (
	citySize = 141

	maxStraightLine = 3
)

func One(
	input string,
) (int, error) {
	c := newCity(input)
	return 0, nil
}

type city struct {
	blocks [citySize][citySize]int
}

func newCity(input string) city {
	ri, ci := 0, 0
	c := city{}
	for len(input) > 0 {
		if input[0] == '\n' {
			ri++
			ci = -1
		} else {
			c.blocks[ri][ci] = int(input[0] - '0')
		}
		ci++
		input = input[1:]
	}

	return c
}

type heading uint8

const (
	east  = 1 << 0
	south = 1 << 1
	west  = 1 << 2
	north = 1 << 3

	southeast = south | east
)

type position struct {
	row int // uint8
	col int // uint8

	lastHeading    heading
	numInDirection int

	totalHeatLoss int
}

func getMinimalHeatLoss(
	c city,
) int {

	pending := newPending()
	pending.insert(
		position{
			row:            1,
			col:            0,
			lastHeading:    south,
			numInDirection: 1,
			totalHeatLoss:  c.blocks[1][0],
		},
	)
	pending.insert(
		position{
			row:            0,
			col:            1,
			lastHeading:    east,
			numInDirection: 1,
			totalHeatLoss:  c.blocks[0][1],
		},
	)

	var pos position
	var left, straight, right *position

	for !pending.isEmpty() {
		pending.sort()
		pos = pending.pop()

		left, straight, right = getNext(c, pos)
		if left != nil {
			pending.insert(*left)
		}
		if straight != nil {
			pending.insert(*straight)
		}
		if right != nil {
			pending.insert(*right)
		}

	}

	return 0
}

func getNext(
	c city,
	pos position,
) (left, straight, right *position) {
	if pos.numInDirection < maxStraightLine {
		s := pos
		straight = &s
		s.numInDirection++

		switch s.lastHeading {
		case east:
			s.col++
			if s.col >= citySize {
				straight = nil
			}
		case south:
			s.row++
			if s.row >= citySize {
				straight = nil
			}
		case west:
			s.col--
			if s.col < 0 {
				straight = nil
			}
		case north:
			s.row--
			if s.row < 0 {
				straight = nil
			}
		}

		straight = &s
	}

	// TODO populate left and right

	return nil, straight, nil
}

type pending struct {
	all []position
}

func newPending() *pending {
	return &pending{
		all: make([]position, 0, 128),
	}
}
func (p *pending) isEmpty() bool {
	return len(p.all) == 0
}

func (p *pending) insert(
	pos position,
) {
	p.all = append(p.all, pos)
}

func (p *pending) sort() {
	slices.SortFunc(p.all, comparePositions)
}

// returns negative when a < b
func comparePositions(
	a, b position,
) int {
	// adist := (citySize - 1 - a.row) + (citySize - 1 - a.col)
	// bdist := (citySize - 1 - b.row) + (citySize - 1 - b.col)
	// adist := 2*citySize - 2 - a.row - a.col
	// bdist := 2*citySize - 2 - b.row - b.col
	adist := a.row + a.col
	bdist := b.row + b.col
	if adist != bdist {
		// return the one closest to the target, the bottom right, which means
		// the sum of the row and col will be largest
		return bdist - adist
	}

	if a.totalHeatLoss != b.totalHeatLoss {
		// if the position at a has a lower total heat loss,
		// that one should be first
		return a.totalHeatLoss - b.totalHeatLoss
	}

	if a.lastHeading != b.lastHeading {
		aGood := (a.lastHeading & southeast) == a.lastHeading
		bGood := (b.lastHeading & southeast) == b.lastHeading
		if aGood != bGood {
			if aGood {
				// a is headed southeast, b is not
				return -1
			}
			// b is headed southeast, a is not
			return 1
		}
	}

	if a.numInDirection != b.numInDirection {
		// if a has gone fewer in a given direction,
		// then that one should be first
		return a.numInDirection - b.numInDirection
	}
	return 0 // no distinguishable difference
}

func (p *pending) pop() position {
	pos := p.all[0]
	p.all = p.all[1:]
	return pos
}
