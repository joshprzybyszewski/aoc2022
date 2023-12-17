package seventeen

import (
	"fmt"
	"slices"
	"strings"
)

const (
	citySize = 141

	maxStraightLine = 3
)

func One(
	input string,
) (int, error) {
	c := newCity(input)
	dijkstraHeatLossToTarget(&c)
	fmt.Printf("city:\n%s\n", c)
	// 1356 is too high
	return getMinimalHeatLoss(c), nil
}

type city struct {
	blocks [citySize][citySize]int

	minHeatLossToTarget [citySize][citySize]int
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

func (c city) String() string {
	var sb strings.Builder
	for ri := 0; ri < citySize; ri++ {
		for ci := 0; ci < citySize; ci++ {
			sb.WriteString(fmt.Sprintf("%4d ", c.blocks[ri][ci]))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func dijkstraHeatLossToTarget(c *city) {
	pending := make([]position, 0, 128)

	pending = append(pending,
		position{
			row:           citySize - 2,
			col:           citySize - 1,
			totalHeatLoss: c.blocks[citySize-1][citySize-1],
		},
		position{
			row:           citySize - 1,
			col:           citySize - 2,
			totalHeatLoss: c.blocks[citySize-1][citySize-1],
		},
	)

	for len(pending) > 0 {
		pos := pending[0]
		if pos.row < 0 || pos.col < 0 {
			pending = pending[1:]
			continue
		}

		if c.minHeatLossToTarget[pos.row][pos.col] == 0 ||
			c.minHeatLossToTarget[pos.row][pos.col] > pos.totalHeatLoss {
			// if it's unset or it's more than we currently know, set it
			c.minHeatLossToTarget[pos.row][pos.col] = pos.totalHeatLoss
			pending = append(pending,
				position{
					row:           pos.row - 1,
					col:           pos.col,
					totalHeatLoss: pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				},
				position{
					row:           pos.row,
					col:           pos.col - 1,
					totalHeatLoss: pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				},
			)
		}

		pending = pending[1:]
	}
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

func (p position) String() string {
	return fmt.Sprintf("(%3d, %3d) %d", p.row, p.col, p.totalHeatLoss)
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

		fmt.Printf("Processing: %s\n", pos)

		if pending.checkSolution(pos) {
			break
		}

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

	return pending.best.totalHeatLoss
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
	}

	l := pos
	left = &l
	l.numInDirection = 0

	switch l.lastHeading {
	case south:
		l.lastHeading = east
		l.col++
		if l.col >= citySize {
			left = nil
		}
	case west:
		l.lastHeading = south
		l.row++
		if l.row >= citySize {
			left = nil
		}
	case north:
		l.lastHeading = west
		l.col--
		if l.col < 0 {
			left = nil
		}
	case east:
		l.lastHeading = north
		l.row--
		if l.row < 0 {
			left = nil
		}
	}

	r := pos
	right = &r
	r.numInDirection = 0

	switch r.lastHeading {
	case north:
		r.lastHeading = east
		r.col++
		if r.col >= citySize {
			right = nil
		}
	case east:
		r.lastHeading = south
		r.row++
		if r.row >= citySize {
			right = nil
		}
	case south:
		r.lastHeading = west
		r.col--
		if r.col < 0 {
			right = nil
		}
	case west:
		r.lastHeading = north
		r.row--
		if r.row < 0 {
			right = nil
		}
	}

	if left != nil {
		left.totalHeatLoss += c.blocks[left.row][left.col]
	}
	if straight != nil {
		straight.totalHeatLoss += c.blocks[straight.row][straight.col]
	}
	if right != nil {
		right.totalHeatLoss += c.blocks[right.row][right.col]
	}

	return left, straight, right
}

type pending struct {
	all []position

	best position
}

func newPending() *pending {
	return &pending{
		all: make([]position, 0, 128),
	}
}

func (p *pending) isEmpty() bool {
	return len(p.all) == 0
}

func (p *pending) checkSolution(
	pos position,
) bool {
	if pos.row != citySize-1 || pos.col != citySize-1 {
		return false
	}

	// TODO check if pos.heatLoss < p.best.heatLoss
	p.best = pos

	// TODO filter out
	return true
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
