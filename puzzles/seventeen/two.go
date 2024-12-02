package seventeen

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

	for len(pending) > 0 {
		pos := pending[0]

		if c.isBetter(&pos) {
			c.remember(&pos)

			pending = append(pending, c.getUltraPrevious(&pos)...)
		}

		pending = pending[1:]
	}
}
