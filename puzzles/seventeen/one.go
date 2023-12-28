package seventeen

const (
	citySize = 141

	minStraightLine = 1
	maxStraightLine = 3

	maxNumPrevious     = 2 * (maxUltraStraightLine + 1 - minUltraStraightLine)
	initialPendingSize = citySize * citySize * maxUltraNumPrevious
)

func One(
	input string,
) (int, error) {
	c := newCity(input)
	dijkstraHeatLossToTarget(&c)

	return c.getMinValAt(0, 0), nil
}

type city struct {
	blocks [citySize][citySize]int

	minHeatLossToTarget [citySize][citySize][maxHeading]int
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
		if pos.row+int(pos.leftInDirection) >= citySize {
			panic(`unexpected`)
		}
		for n := 1; n <= int(pos.leftInDirection); n++ {
			pos.totalHeatLoss += c.blocks[pos.row+n][pos.col]
		}
	case north:
		if pos.row-int(pos.leftInDirection) < 0 {
			panic(`unexpected`)
		}

		for n := 1; n <= int(pos.leftInDirection); n++ {
			pos.totalHeatLoss += c.blocks[pos.row-n][pos.col]
		}
	case east:
		if pos.col+int(pos.leftInDirection) >= citySize {
			panic(`unexpected`)
		}
		for n := 1; n <= int(pos.leftInDirection); n++ {
			pos.totalHeatLoss += c.blocks[pos.row][pos.col+n]
		}
	case west:
		if pos.col-int(pos.leftInDirection) < 0 {
			panic(`unexpected`)
		}

		for n := 1; n <= int(pos.leftInDirection); n++ {
			pos.totalHeatLoss += c.blocks[pos.row][pos.col-n]
		}
	}
}

func (c *city) getMinValAt(ri, ci int) int {
	min := 0
	for _, v := range c.minHeatLossToTarget[ri][ci] {
		if v == 0 {
			continue
		}
		if min == 0 || v < min {
			min = v
		}
	}
	return min
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

	val := c.minHeatLossToTarget[pos.row][pos.col][pos.heading]
	if val != 0 &&
		val <= pos.totalHeatLoss {
		return false
	}

	val = c.minHeatLossToTarget[pos.row][pos.col][pos.heading.opposite()]
	if val != 0 &&
		val <= pos.totalHeatLoss {
		return false
	}

	return true
}

func (c *city) remember(
	pos *position,
) {
	c.minHeatLossToTarget[pos.row][pos.col][pos.heading] = pos.totalHeatLoss
	c.minHeatLossToTarget[pos.row][pos.col][pos.heading.opposite()] = pos.totalHeatLoss
}

func (c *city) getPrevious(
	pos *position,
) []position {

	output := make([]position, 0, maxNumPrevious)

	if pos.heading == east || pos.heading == west {
		for i := uint8(minStraightLine); i <= maxStraightLine; i++ {
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
		for i := uint8(minStraightLine); i <= maxStraightLine; i++ {
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
		for i := uint8(minStraightLine); i <= maxStraightLine; i++ {
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
		for i := uint8(minStraightLine); i <= maxStraightLine; i++ {
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

func dijkstraHeatLossToTarget(c *city) {
	pending := make([]position, 0, initialPendingSize)

	for i := uint8(minStraightLine); i <= maxStraightLine; i++ {
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

		if c.isBetter(&pos) {
			remembered++
			c.remember(&pos)

			pending = append(pending, c.getPrevious(&pos)...)
		}

		pending = pending[1:]
	}
}

type heading uint8

const (
	east  heading = 1
	south heading = 2
	west  heading = 3
	north heading = 4

	maxHeading = 5
)

func (h heading) opposite() heading {
	// 1 -> 3 : 0b0001 -> 0b0011
	// 2 -> 4 : 0b0010 -> 0b0100
	// 3 -> 1 : 0b0011 -> 0b0001
	// 4 -> 2 : 0b0100 -> 0b0010
	h += 2

	if h > 4 {
		h -= 4
	}
	return h
}

type position struct {
	row int // uint8
	col int // uint8

	heading         heading
	leftInDirection uint8
	straight        int

	totalHeatLoss int

	prev *position
}
