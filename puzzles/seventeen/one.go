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
	// fmt.Printf("city:\n%s\n", c)
	// 1356 is too high
	// 1290 is too high
	// 1270 is too high
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
	sb.WriteString("        ")
	for ci := 0; ci < citySize; ci++ {
		sb.WriteString(fmt.Sprintf("%4d ", ci))
	}
	sb.WriteByte('\n')
	for ri := 0; ri < citySize; ri++ {
		sb.WriteString(fmt.Sprintf("Row %3d:", ri))
		for ci := 0; ci < citySize; ci++ {
			sb.WriteString(fmt.Sprintf("%4d ", c.minHeatLossToTarget[ri][ci]))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func drawPath(p *position, c *city) string {
	var output [citySize][citySize]byte
	for ri := 0; ri < citySize; ri++ {
		for ci := 0; ci < citySize; ci++ {
			output[ri][ci] = '0' + byte(c.blocks[ri][ci])
		}
	}

	for p != nil {
		switch p.heading {
		case east:
			output[p.row][p.col] = '>'
		case south:
			output[p.row][p.col] = 'v'
		case west:
			output[p.row][p.col] = '<'
		case north:
			output[p.row][p.col] = '^'
		default:
			panic(`ahh`)
		}
		p = p.prev
	}

	var sb strings.Builder
	for ri := 0; ri < citySize; ri++ {
		sb.WriteString(string(output[ri][:]))
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
		if pos.row < 0 ||
			pos.col < 0 ||
			pos.row >= citySize ||
			pos.col >= citySize ||
			(pos.row == citySize-1 && pos.col == citySize-1) {
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
				position{
					row:           pos.row + 1,
					col:           pos.col,
					totalHeatLoss: pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				},
				position{
					row:           pos.row,
					col:           pos.col + 1,
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

	maxHeading = 1 << 4
)

type position struct {
	row int // uint8
	col int // uint8

	heading        heading
	numInDirection int

	totalHeatLoss int

	depth int

	prev *position
}

func (p position) canGo() heading {
	var out heading

	switch p.heading {
	case south:
		out = east | west
	case west:
		out = south | north
	case north:
		out = west | east
	case east:
		out = north | south
	}

	if p.numInDirection < maxStraightLine {
		out |= p.heading
	}

	return out
}

func (p position) String() string {
	return fmt.Sprintf("(%3d, %3d) %d", p.row, p.col, p.totalHeatLoss)
}

func getMinimalHeatLoss(
	c city,
) int {

	pending := newPending(&c)
	pending.insert(
		position{
			row:            1,
			col:            0,
			heading:        south,
			numInDirection: 1,
			totalHeatLoss:  c.blocks[1][0],
			depth:          1,
			prev:           nil,
		},
	)
	pending.insert(
		position{
			row:            0,
			col:            1,
			heading:        east,
			numInDirection: 1,
			totalHeatLoss:  c.blocks[0][1],
			depth:          1,
			prev:           nil,
		},
	)

	var pos position
	var left, straight, right *position

	for !pending.isEmpty() {
		pos = pending.pop()
		// fmt.Printf("Processing: %s\n", pos)
		// fmt.Printf("Processing:\n%s\n", drawPath(&pos, pending.city))
		// time.Sleep(time.Millisecond * 16)

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
	pos.depth++
	if pos.numInDirection < maxStraightLine {
		s := pos
		straight = &s
		s.numInDirection++
		s.prev = &pos

		switch s.heading {
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
	l.numInDirection = 1
	l.prev = &pos

	switch l.heading {
	case south:
		l.heading = east
		l.col++
		if l.col >= citySize {
			left = nil
		}
	case west:
		l.heading = south
		l.row++
		if l.row >= citySize {
			left = nil
		}
	case north:
		l.heading = west
		l.col--
		if l.col < 0 {
			left = nil
		}
	case east:
		l.heading = north
		l.row--
		if l.row < 0 {
			left = nil
		}
	}

	r := pos
	right = &r
	r.numInDirection = 1
	r.prev = &pos

	switch r.heading {
	case north:
		r.heading = east
		r.col++
		if r.col >= citySize {
			right = nil
		}
	case east:
		r.heading = south
		r.row++
		if r.row >= citySize {
			right = nil
		}
	case south:
		r.heading = west
		r.col--
		if r.col < 0 {
			right = nil
		}
	case west:
		r.heading = north
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
	city *city

	all         []position
	yetToInsert []position

	bestByBlock [citySize][citySize][maxHeading]*position

	best *position
}

func newPending(city *city) *pending {
	return &pending{
		city: city,
	}
}

func (p *pending) isEmpty() bool {
	return len(p.all) == 0 && len(p.yetToInsert) == 0
}

func (p *pending) pop() position {
	p.sort()
	pos := p.all[0]
	p.checkSolution(pos)
	p.all = p.all[1:]
	return pos
}

func (p *pending) filter() {
	for i := 0; i < len(p.all); {
		if p.cannotBeBest(p.all[i]) {
			if i < len(p.all)-1 {
				p.all[i] = p.all[len(p.all)-1]
			}
			p.all = p.all[:len(p.all)-1]
		} else {
			i++
		}
	}
}

func (p *pending) cannotBeBest(pos position) bool {
	if p.best != nil &&
		p.best.totalHeatLoss < pos.totalHeatLoss+p.city.minHeatLossToTarget[pos.row][pos.col] {
		return true
	}
	if blockBest := p.bestByBlock[pos.row][pos.col][pos.canGo()]; blockBest != nil &&
		blockBest.totalHeatLoss < pos.totalHeatLoss {
		return false
	}
	return false
}

func (p *pending) checkSolution(
	pos position,
) {
	if pos.row != citySize-1 || pos.col != citySize-1 {
		return
	}

	if p.best == nil {
		fmt.Printf("Found Solution:\n%s\nFIRST: %4d (%d pending)\n\n", drawPath(&pos, p.city), pos.totalHeatLoss, len(p.all))
		p.best = &pos
		p.filter()
		slices.SortFunc(p.all, p.comparePositions)
	} else if pos.totalHeatLoss < p.best.totalHeatLoss {
		fmt.Printf("Found New Best:\n%s\nNEW BEST: %4d (%d pending)\n\n", drawPath(&pos, p.city), pos.totalHeatLoss, len(p.all))
		p.best = &pos
		p.filter()
	}
}

func (p *pending) insert(
	pos position,
) {
	if p.cannotBeBest(pos) {
		return
	}

	p.bestByBlock[pos.row][pos.col][pos.canGo()] = &pos

	p.yetToInsert = append(p.yetToInsert, pos)
}

func (p *pending) sort() {
	if len(p.yetToInsert) == 0 {
		return
	}
	for _, e := range p.yetToInsert {
		ei := slices.IndexFunc(p.all, func(a position) bool {
			return p.comparePositions(a, e) > 0
		})
		if ei == -1 {
			p.all = append(p.all, e)
			continue
		}
		// Using slices.Insert doesn't seem to work.
		// slices.Insert(p.all, ei, e)
		after := make([]position, len(p.all)-ei)
		copy(after, p.all[ei:])
		p.all = p.all[:ei]              // trim to before
		p.all = append(p.all, e)        // insert the new element
		p.all = append(p.all, after...) // add the after ones.
	}
	p.yetToInsert = p.yetToInsert[:0]
}

// returns negative when a < b
func (p *pending) comparePositions(
	a, b position,
) int {

	// if p.best != nil {
	// 	aProjection := a.totalHeatLoss + p.city.minHeatLossToTarget[a.row][a.col]
	// 	bProjection := b.totalHeatLoss + p.city.minHeatLossToTarget[b.row][b.col]
	// 	if aProjection != bProjection {
	// 		// if the position at a has a lower projected heat loss,
	// 		// that one should be first
	// 		return aProjection - bProjection
	// 	}
	// }

	adist := a.row + a.col
	bdist := b.row + b.col

	// if p.best != nil {
	// 	// aLossPerStep := a.totalHeatLoss / adist
	// 	// bLossPerStep := b.totalHeatLoss / bdist
	// 	aLossPerStep := a.totalHeatLoss / a.depth
	// 	bLossPerStep := b.totalHeatLoss / b.depth
	// 	if aLossPerStep != bLossPerStep {
	// 		return aLossPerStep - bLossPerStep
	// 	}
	// }

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

	if a.heading != b.heading {
		aGood := (a.heading & southeast) == a.heading
		bGood := (b.heading & southeast) == b.heading
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
		// if a has gone more in a given direction,
		// then that one should be first
		return b.numInDirection - a.numInDirection
	}

	aheatdelta := p.city.blocks[a.row][a.col]
	bheatdelta := p.city.blocks[b.row][b.col]
	if aheatdelta != bheatdelta {
		return aheatdelta - bheatdelta
	}
	return 0 // no distinguishable difference
}
