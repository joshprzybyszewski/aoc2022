package twentyone

import "fmt"

type precomputedGarden struct {
	distances [gridSize][gridSize]int

	exits struct {
		left   coord
		top    coord
		right  coord
		bottom coord
	}

	entrance coord

	numEven     int
	numOdd      int
	maxDistance int
}

func newPrecomputedInitialGarden(
	gar garden,
) precomputedGarden {
	return newPrecomputedGardenWithEntranceDepth(
		&gar,
		gar.start,
		0,
	)
}

func newPrecomputedGarden(
	gar *garden,
	entrance coord,
) precomputedGarden {
	return newPrecomputedGardenWithEntranceDepth(
		gar,
		entrance,
		1,
	)
}

func newPrecomputedGardenWithEntranceDepth(
	gar *garden,
	entrance coord,
	initialDepth int,
) precomputedGarden {
	pg := precomputedGarden{
		entrance: entrance,
	}

	type pos struct {
		coord
		depth int
	}
	pending := make([]pos, 0, gridSize*gridSize)

	pending = append(pending, pos{
		coord: pg.entrance,
		depth: initialDepth,
	})

	for len(pending) > 0 {
		c := pending[0]
		pending = pending[1:]

		if pg.distances[c.row][c.col] > 0 {
			if c.depth < pg.distances[c.row][c.col] {
				fmt.Printf("c.depth: %d, actual: %d", c.depth, pg.distances[c.row][c.col])
				panic(`unexpected`)
			}
			continue
		}

		pg.distances[c.row][c.col] = c.depth

		if c.row == 0 {
			if pg.exits.top.row == 0 {
				// I can exit at this col
				pg.exits.top.row = -1
				pg.exits.top.col = c.col
			}
		} else if pg.distances[c.row-1][c.col] == 0 { // c.row MUST be > 0
			pending = append(pending, pos{
				coord: coord{
					row: c.row - 1,
					col: c.col,
				},
				depth: c.depth + 1,
			})
		}

		if c.col == 0 {
			if pg.exits.left.col == 0 {
				// I can exit at this row
				pg.exits.left.col = -1
				pg.exits.left.row = c.row
			}
		} else if pg.distances[c.row][c.col-1] == 0 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col - 1,
				},
				depth: c.depth + 1,
			})
		}
		if c.row == gridSize-1 {
			if pg.exits.bottom.row == 0 {
				// I can exit at this col
				pg.exits.bottom.row = gridSize
				pg.exits.bottom.col = c.col
			}
		} else if pg.distances[c.row+1][c.col] == 0 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row + 1,
					col: c.col,
				},
				depth: c.depth + 1,
			})
		}
		if c.col == gridSize-1 {
			if pg.exits.right.col == 0 {
				// I can exit at this row
				pg.exits.right.col = gridSize
				pg.exits.right.row = c.row
			}
		} else if pg.distances[c.row][c.col+1] == 0 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col + 1,
				},
				depth: c.depth + 1,
			})
		}
	}

	if pg.exits.left.col == 0 ||
		pg.exits.right.col == 0 ||
		pg.exits.top.row == 0 ||
		pg.exits.bottom.row == 0 {
		panic(`unexpected`)
	}
	// now they can be used as the entrance for that direction
	pg.exits.left.col = gridSize - 1
	pg.exits.right.col = 0
	pg.exits.top.row = gridSize - 1
	pg.exits.bottom.row = 0

	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] > pg.maxDistance {
				pg.maxDistance = pg.distances[ri][ci]
			}
			if pg.distances[ri][ci]%2 == 0 {
				pg.numEven++
			} else {
				pg.numOdd++
			}
		}
	}

	return pg
}

func (pg *precomputedGarden) leftDepth() int {
	return pg.distances[pg.exits.left.row][0]
}

func (pg *precomputedGarden) rightDepth() int {
	return pg.distances[pg.exits.right.row][gridSize-1]
}

func (pg *precomputedGarden) topDepth() int {
	return pg.distances[pg.exits.top.row][0]
}

func (pg *precomputedGarden) bottomDepth() int {
	return pg.distances[pg.exits.bottom.row][gridSize-1]
}

func (pg *precomputedGarden) getNumEven(
	maxDepth int,
) int {
	if maxDepth > pg.maxDistance {
		return pg.numEven
	}

	output := 0
	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] > maxDepth {
				continue
			}
			if pg.distances[ri][ci]%2 == 0 {
				pg.numEven++
			}
		}
	}
	return output
}
