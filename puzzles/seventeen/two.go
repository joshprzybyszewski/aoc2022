package seventeen

import "fmt"

const (
	maxUltraStraightLine = 10
	minUltraStraightLine = 4

	requiredBeforeTurn = minUltraStraightLine - 1
)

func Two(
	input string,
) (int, error) {

	c := newCity(input)
	dijkstraUltraHeatLossToTarget(&c)

	min := -1
	for _, vals := range c.minHeatLossToTarget[0][0] {
		for _, v := range vals {
			if min == -1 && v != 0 {
				min = v
			}
		}
	}

	return min, nil
}

func (c *city) getUltraPrevious(
	pos position,
) []position {

	output := make([]position, 0, 32)

	// comes from the north
	if (pos.heading == east ||
		pos.heading == west) && pos.numStraight() > requiredBeforeTurn {
		for i := uint8(0); i < maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row - 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         south,
				leftInDirection: i,
				straight:        1,
			})
		}
	}
	if pos.heading == south {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row - 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         south,
				leftInDirection: uint8(n),
				straight:        pos.straight + 1,
			})
		}
	}

	// comes from the west
	if (pos.heading == north ||
		pos.heading == south) && pos.numStraight() > requiredBeforeTurn {
		for i := uint8(0); i < maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col - 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         east,
				leftInDirection: i,
				straight:        1,
			})
		}
	}
	if pos.heading == east {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col - 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         east,
				leftInDirection: uint8(n),
				straight:        pos.straight + 1,
			})
		}
	}

	// comes from the east
	if (pos.heading == north ||
		pos.heading == south) && pos.numStraight() > requiredBeforeTurn {
		for i := uint8(0); i < maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col + 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         west,
				leftInDirection: i,
				straight:        1,
			})
		}
	}
	if pos.heading == west {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col + 1,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         west,
				leftInDirection: uint8(n),
				straight:        pos.straight + 1,
			})
		}
	}

	// comes from the south
	if (pos.heading == east ||
		pos.heading == west) && pos.numStraight() > requiredBeforeTurn {
		for i := uint8(0); i < maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row + 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         north,
				leftInDirection: i,
				straight:        1,
			})
		}
	}
	if pos.heading == north {
		for n := int(pos.leftInDirection) - 1; n >= 0; n-- {
			output = append(output, position{
				prev:            &pos,
				row:             pos.row + 1,
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss + c.blocks[pos.row][pos.col],
				heading:         north,
				leftInDirection: uint8(n),
				straight:        pos.straight + 1,
			})
		}
	}

	return output
}

func dijkstraUltraHeatLossToTarget(c *city) {
	pending := make([]position, 0, 2048)

	for i := uint8(0); i < maxUltraStraightLine; i++ {
		pending = append(pending,
			position{
				row:             citySize - 2,
				col:             citySize - 1,
				totalHeatLoss:   c.blocks[citySize-1][citySize-1],
				heading:         south,
				leftInDirection: i,
				straight:        1,
			},
			position{
				row:             citySize - 1,
				col:             citySize - 2,
				totalHeatLoss:   c.blocks[citySize-1][citySize-1],
				heading:         east,
				leftInDirection: i,
				straight:        1,
			},
		)
	}

	var iterated, remembered int
	fmt.Printf("remaining:  %d\n", len(pending))

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

			pending = append(pending, c.getUltraPrevious(pos)...)
		}

		pending = pending[1:]
	}

	fmt.Printf("iterated:   %d\n", iterated)
	fmt.Printf("remembered: %d\n", remembered)
	fmt.Printf("city\n%s\n\n", c)
}
