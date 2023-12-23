package seventeen

import (
	"fmt"
)

const (
	maxUltraStraightLine = 10
	minUltraStraightLine = 4

	maxUltraNumPrevious     = 2 * (maxUltraStraightLine + 1 - minUltraStraightLine)
	initialUltraPendingSize = citySize * citySize * maxUltraNumPrevious
)

func Two(
	input string,
) (int, error) {

	c := newCity(input)
	dijkstraUltraHeatLossToTarget(&c)

	return c.getMinValAt(0, 0), nil
}

func (c *city) getUltraPrevious(
	pos *position,
) []position {

	output := make([]position, 0, maxUltraNumPrevious)

	if pos.heading == east || pos.heading == west {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			output = append(output,
				position{
					prev:            pos,
					row:             pos.row - int(i),
					col:             pos.col,
					totalHeatLoss:   pos.totalHeatLoss,
					heading:         south,
					leftInDirection: i,
				},
			)
		}
	}

	if pos.heading == north || pos.heading == south {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            pos,
				row:             pos.row,
				col:             pos.col - int(i),
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         east,
				leftInDirection: i,
			})
		}
	}

	if pos.heading == north || pos.heading == south {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            pos,
				row:             pos.row,
				col:             pos.col + int(i),
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         west,
				leftInDirection: i,
			})
		}
	}

	if pos.heading == east || pos.heading == west {
		for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
			output = append(output, position{
				prev:            pos,
				row:             pos.row + int(i),
				col:             pos.col,
				totalHeatLoss:   pos.totalHeatLoss,
				heading:         north,
				leftInDirection: i,
			})
		}
	}

	return output
}

func dijkstraUltraHeatLossToTarget(c *city) {
	pending := make([]position, 0, initialUltraPendingSize)

	for i := uint8(minUltraStraightLine); i <= maxUltraStraightLine; i++ {
		above := position{
			row:             citySize - 1 - int(i),
			col:             citySize - 1,
			totalHeatLoss:   0,
			heading:         south,
			leftInDirection: i,
		}

		left := position{
			row:             citySize - 1,
			col:             citySize - 1 - int(i),
			totalHeatLoss:   0,
			heading:         east,
			leftInDirection: i,
		}

		pending = append(pending,
			above,
			left,
		)
	}

	var iterated, remembered int

	// var allStarting []position

	for len(pending) > 0 {
		iterated++
		pos := pending[0]

		// if iterated%100 == 0 {
		// 	fmt.Printf("iterated:   %d\n", iterated)
		// 	fmt.Printf("remembered: %d\n", remembered)
		// 	fmt.Printf("remaining:  %d\n", len(pending))
		// 	fmt.Printf("\n")
		// 	fmt.Printf("city\n%s\n\n", c.withPos(pos))
		// 	// time.Sleep(64 * time.Millisecond)
		// }

		if c.isBetter(&pos) {
			remembered++
			c.remember(&pos)

			pending = append(pending, c.getUltraPrevious(&pos)...)
		}

		// if pos.row == 0 && pos.col == 0 {
		// 	allStarting = append(allStarting, pos)
		// }

		pending = pending[1:]
	}

	/*
		slices.SortFunc(allStarting, func(a, b position) int {
			return a.totalHeatLoss - b.totalHeatLoss
		})

		for _, pos := range allStarting {
			fmt.Printf("From Start: %d\n", pos.totalHeatLoss)
			fmt.Printf("DoubleCheck: %d\n", c.getPathHeatLoss(&pos))
			fmt.Printf("%s\n\n", c.withPos(pos))
			if pos.totalHeatLoss > 1263 {
				break
			}
			// time.Sleep(64 * time.Millisecond)
		}
	*/

	fmt.Printf("iterated:   %d\n", iterated)
	fmt.Printf("remembered: %d\n", remembered)
	fmt.Printf("city\n%s\n\n", c)
}
