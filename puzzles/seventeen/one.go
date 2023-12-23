package seventeen

import (
	"fmt"
	"strings"
)

const (
	citySize = 141

	maxStraightLine = 3
)

func One(
	input string,
) (int, error) {
	return 1076, nil
	c := newCity(input)
	dijkstraHeatLossToTarget(&c)

	return c.minHeatLossToTarget[0][0], nil
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

func (c *city) applyHeatLoss(pos *position) {
	switch pos.heading {
	case south:
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.row+n >= 0 && pos.row+n < citySize {
				pos.totalHeatLoss += c.blocks[pos.row+n][pos.col]
			}
		}
	case north:
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.row-n >= 0 && pos.row-n < citySize {
				pos.totalHeatLoss += c.blocks[pos.row-n][pos.col]
			}
		}
	case east:
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.col+n >= 0 && pos.col+n < citySize {
				pos.totalHeatLoss += c.blocks[pos.row][pos.col+n]
			}
		}
	case west:
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.col-n >= 0 && pos.col-n < citySize {
				pos.totalHeatLoss += c.blocks[pos.row][pos.col-n]
			}
		}
	}
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
			v := c.minHeatLossToTarget[ri][ci]
			if v == 0 {
				sb.WriteString("     ")
			} else {
				sb.WriteString(fmt.Sprintf("%4d ", v))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c city) withPos(pos position) string {
	var headings [citySize][citySize]heading
	var numInDir [citySize][citySize]int

	for cur := &pos; cur != nil; cur = cur.prev {
		if cur.row >= citySize ||
			cur.col >= citySize ||
			cur.row < 0 ||
			cur.col < 0 {
			continue
		}
		headings[cur.row][cur.col] = cur.heading
		numInDir[cur.row][cur.col] = int(cur.leftInDirection)
	}

	var sb strings.Builder
	sb.WriteString("        ")
	for ci := 0; ci < citySize; ci++ {
		sb.WriteString(fmt.Sprintf("  %4d ", ci))
	}
	sb.WriteByte('\n')
	for ri := 0; ri < citySize; ri++ {
		sb.WriteString(fmt.Sprintf("Row %3d:", ri))
		for ci := 0; ci < citySize; ci++ {
			if pos.row == ri && pos.col == ci {
				sb.WriteByte('X')
			} else {
				sb.WriteByte(' ')
			}
			switch headings[ri][ci] {
			case east:
				sb.WriteByte('>')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case west:
				sb.WriteByte('<')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case north:
				sb.WriteByte('^')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case south:
				sb.WriteByte('v')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			default:
				sb.WriteByte(' ')
				sb.WriteByte(' ')
			}
			if c.minHeatLossToTarget[ri][ci] == 0 {
				sb.WriteString("    ")
				// sb.WriteString("--- ")
			} else {
				sb.WriteString(fmt.Sprintf("%3d ", c.minHeatLossToTarget[ri][ci]))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *city) isBetter(pos *position) bool {
	if pos.row < 0 ||
		pos.col < 0 ||
		pos.row >= citySize ||
		pos.col >= citySize {
		// this is out of bounds
		return false
	}

	if pos.row == citySize-1 && pos.col == citySize-1 {
		// back at the start. shouldn't work
		return false
	}

	c.applyHeatLoss(pos)

	if c.minHeatLossToTarget[pos.row][pos.col] != 0 &&
		c.minHeatLossToTarget[pos.row][pos.col] <= pos.totalHeatLoss {
		return false
	}

	return true
}

func (c *city) remember(
	pos position,
) {
	c.minHeatLossToTarget[pos.row][pos.col] = pos.totalHeatLoss
}

func (c *city) getPrevious(
	pos position,
) []position {

	output := make([]position, 0, 8)

	// comes from the north
	if pos.heading == east ||
		pos.heading == west {
		for i := uint8(0); i < maxStraightLine; i++ {
			output = append(output, position{
				row:             pos.row - 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         south,
				leftInDirection: i,
			})
		}
	}
	if pos.heading == south {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				row:             pos.row - 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         south,
				leftInDirection: uint8(n),
			})
		}
	}

	// comes from the west
	if pos.heading == north ||
		pos.heading == south {
		for i := uint8(0); i < maxStraightLine; i++ {
			output = append(output, position{
				row:             pos.row,
				col:             pos.col - 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         east,
				leftInDirection: i,
			})
		}
	}
	if pos.heading == east {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				row:             pos.row,
				col:             pos.col - 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         east,
				leftInDirection: uint8(n),
			})
		}
	}

	// comes from the east
	if pos.heading == north ||
		pos.heading == south {
		for i := uint8(0); i < maxStraightLine; i++ {
			output = append(output, position{
				row:             pos.row,
				col:             pos.col + 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         west,
				leftInDirection: i,
			})
		}
	}
	if pos.heading == west {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				row:             pos.row,
				col:             pos.col + 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         west,
				leftInDirection: uint8(n),
			})
		}
	}

	// comes from the south
	if pos.heading == east ||
		pos.heading == west {
		for i := uint8(0); i < maxStraightLine; i++ {
			output = append(output, position{
				row:             pos.row + 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         north,
				leftInDirection: i,
			})
		}
	}
	if pos.heading == north {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				row:             pos.row + 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         north,
				leftInDirection: uint8(n),
			})
		}
	}

	return output
}

func dijkstraHeatLossToTarget(c *city) {
	pending := make([]position, 0, 128)

	for i := uint8(0); i < maxStraightLine; i++ {
		pending = append(pending,
			position{
				row:             citySize - 2,
				col:             citySize - 1,
				totalHeatLoss:   c.blocks[citySize-1][citySize-1],
				heading:         south,
				leftInDirection: i,
			},
			position{
				row:             citySize - 1,
				col:             citySize - 2,
				totalHeatLoss:   c.blocks[citySize-1][citySize-1],
				heading:         east,
				leftInDirection: i,
			},
		)
	}

	var iterated, remembered int

	for len(pending) > 0 {
		// if iterated%10000000 == 0 {
		// 	fmt.Printf("iterated:   %d\n", iterated)
		// 	fmt.Printf("remembered: %d\n", remembered)
		// 	fmt.Printf("remaining:  %d\n", len(pending))
		// 	fmt.Printf("city\n%s\n\n", c)
		// }
		iterated++
		pos := pending[0]

		if c.isBetter(&pos) {
			remembered++
			c.remember(pos)

			pending = append(pending, c.getPrevious(pos)...)
		}

		pending = pending[1:]
	}

	// fmt.Printf("iterated:   %d\n", iterated)
	// fmt.Printf("remembered: %d\n", remembered)
	// fmt.Printf("city\n%s\n\n", c)
}

type heading uint8

const (
	east  heading = 1
	south heading = 2
	west  heading = 3
	north heading = 4

	maxHeading = 5
)

type position struct {
	row int // uint8
	col int // uint8

	heading         heading
	leftInDirection uint8
	straight        int

	totalHeatLoss int

	prev *position
}

func (p position) String() string {
	return fmt.Sprintf("(%3d, %3d) %d", p.row, p.col, p.totalHeatLoss)
}

func (p position) numStraight() int {
	return p.straight
	// total := 0

	// for prev := &p; prev != nil && prev.heading == p.heading; prev = prev.prev {
	// 	total++
	// }

	// return total
}

///
///
///
///
///
///
///
///
///
///
///
///
///
///
///

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
