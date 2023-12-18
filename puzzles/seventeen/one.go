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
	c := newCity(input)
	dijkstraHeatLossToTarget(&c)

	min := -1
	for _, vals := range c.minHeatLossToTarget[0][0] {
		for _, v := range vals {
			if min == -1 && v != 0 {
				min = v
			}
		}
	}
	// fmt.Printf("city:\n%s\n", c)
	// 1356 is too high
	// 1290 is too high
	// 1270 is too high
	return min, nil
}

type city struct {
	blocks [citySize][citySize]int

	minHeatLossToTarget [citySize][citySize][maxHeading][maxStraightLine]int
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
			v := -1
			for _, vals := range c.minHeatLossToTarget[ri][ci] {
				for _, tot := range vals {
					if tot != 0 && (tot < v || v == -1) {
						v = tot
					}
				}
			}
			sb.WriteString(fmt.Sprintf("%4d ", v)) // c.minHeatLossToTarget[ri][ci]
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *city) isBetter(pos position) bool {
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

	if c.minHeatLossToTarget[pos.row][pos.col][pos.heading][pos.leftInDirection] != 0 &&
		c.minHeatLossToTarget[pos.row][pos.col][pos.heading][pos.leftInDirection] <= pos.totalHeatLoss {
		return false
	}

	return true
}

func (c *city) remember(
	pos position,
) {
	c.minHeatLossToTarget[pos.row][pos.col][pos.heading][pos.leftInDirection] = pos.totalHeatLoss
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

		if c.isBetter(pos) {
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

	totalHeatLoss int

	prev *position
}

func (p position) String() string {
	return fmt.Sprintf("(%3d, %3d) %d", p.row, p.col, p.totalHeatLoss)
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
