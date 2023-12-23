package seventeen

import (
	"fmt"
)

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

	// 1294 is too high "and it's the right answer for somebody else"
	// 1269 is too high
	// 1261 is too high

	return c.minHeatLossToTarget[0][0], nil
}

func (c *city) getUltraPrevious(
	pos position,
) []position {

	output := make([]position, 0, 32)

	// comes from the north
	if pos.heading == east || pos.heading == west {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			newPos := position{
				prev:            &pos,
				row:             pos.row - int(i),
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         south,
				leftInDirection: i,
			}
			output = append(output, newPos)
		}
	}

	// comes from the west
	if pos.heading == north || pos.heading == south {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			newPos := position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col - int(i),
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         east,
				leftInDirection: i,
			}
			output = append(output, newPos)
		}
	}

	// comes from the east
	if pos.heading == north || pos.heading == south {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			newPos := position{
				prev:            &pos,
				row:             pos.row,
				col:             pos.col + int(i),
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         west,
				leftInDirection: i,
			}
			output = append(output, newPos)
		}
	}

	// comes from the south
	if pos.heading == east || pos.heading == west {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			newPos := position{
				prev:            &pos,
				row:             pos.row + int(i),
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         north,
				leftInDirection: i,
			}
			output = append(output, newPos)
		}
	}

	return output
}

func dijkstraUltraHeatLossToTarget(c *city) {
	pending := make([]position, 0, 2048)

	for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
		above := position{
			row:             citySize - 1 - int(i),
			col:             citySize - 1,
			heading:         south,
			leftInDirection: i,
		}

		left := position{
			row:             citySize - 1,
			col:             citySize - 1 - int(i),
			heading:         east,
			leftInDirection: i,
		}

		pending = append(pending,
			above,
			left,
		)
	}

	var iterated, remembered int
	fmt.Printf("remaining:  %d\n", len(pending))

	for len(pending) > 0 {
		iterated++
		pos := pending[0]
		if iterated%10000 == 0 {
			fmt.Printf("iterated:   %d\n", iterated)
			fmt.Printf("remembered: %d\n", remembered)
			fmt.Printf("remaining:  %d\n", len(pending))
			fmt.Printf("\n")
			// fmt.Printf("city\n%s\n\n", c.withPos(pos))
			// time.Sleep(16 * time.Millisecond)
		}

		if c.isBetter(&pos) {
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
