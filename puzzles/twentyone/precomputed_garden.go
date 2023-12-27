package twentyone

import (
	"fmt"
	"strings"
)

type precomputedGarden struct {
	distances [gridSize][gridSize]int

	exits struct {
		left   coord
		top    coord
		right  coord
		bottom coord
	}

	additionalExits struct {
		left   []coord
		top    []coord
		right  []coord
		bottom []coord
	}

	entrance       coord
	otherEntrances []coord

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
		nil,
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
		nil,
	)
}

func newPrecomputedGardenWithManyEntrances(
	gar *garden,
	entrance coord,
	additionalEntrances []coord,
) precomputedGarden {
	return newPrecomputedGardenWithEntranceDepth(
		gar,
		entrance,
		1,
		additionalEntrances,
	)
}

func newPrecomputedGardenWithEntranceDepth(
	gar *garden,
	entrance coord,
	initialDepth int,
	additionalEntrances []coord,
) precomputedGarden {
	pg := precomputedGarden{
		entrance:       entrance,
		otherEntrances: additionalEntrances,
	}

	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if gar.tiles[ri][ci] == rocks {
				// this tile is unreachable
				pg.distances[ri][ci] = -1
			}
		}
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

	for _, entrance := range additionalEntrances {
		pending = append(pending, pos{
			coord: entrance,
			depth: initialDepth,
		})
	}

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
			} else {
				otherVal := pg.distances[0][pg.exits.top.col]
				if c.depth <= otherVal {
					exit := c.coord
					exit.row = gridSize - 1
					pg.additionalExits.top = append(pg.additionalExits.top, exit)
				}
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
			} else {
				otherVal := pg.distances[pg.exits.left.row][0]
				if c.depth <= otherVal {
					exit := c.coord
					exit.col = gridSize - 1
					pg.additionalExits.left = append(pg.additionalExits.left, exit)
				}
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
			} else {
				otherVal := pg.distances[gridSize-1][pg.exits.bottom.col]
				if c.depth <= otherVal {
					exit := c.coord
					exit.row = 0
					pg.additionalExits.bottom = append(pg.additionalExits.bottom, exit)
				}
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
			} else {
				otherVal := pg.distances[pg.exits.right.row][gridSize-1]
				if c.depth <= otherVal {
					exit := c.coord
					exit.col = 0
					pg.additionalExits.right = append(pg.additionalExits.right, exit)
				}
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
	// fmt.Printf(" getNumEven(%d)\n", maxDepth)

	if maxDepth > pg.maxDistance {
		// fmt.Printf("  pg.maxDistance < maxDepth: %d < %d\n", pg.maxDistance, maxDepth)
		// fmt.Printf(" = %d\n", pg.numEven)
		return pg.numEven
	}

	output := 0
	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] < 0 ||
				pg.distances[ri][ci] > maxDepth {
				continue
			}
			if pg.distances[ri][ci]%2 == 0 {
				output++
			}
		}
	}
	// fmt.Printf(" = %d\n", output)
	return output
}

func (pg precomputedGarden) String() string {
	var sb strings.Builder

	sb.WriteString(`Entrance: `)
	sb.WriteString(fmt.Sprintf("%+v", pg.entrance))
	sb.WriteByte('\n')

	for _, entrance := range pg.otherEntrances {
		sb.WriteString(`Additional Entrance: `)
		sb.WriteString(fmt.Sprintf("%+v", entrance))
		sb.WriteByte('\n')
	}

	sb.WriteString("Exits:\n")
	sb.WriteString(` top:`)
	sb.WriteString(fmt.Sprintf("%+v", pg.exits.top))
	for _, exit := range pg.additionalExits.top {
		sb.WriteString(fmt.Sprintf(", %+v", exit))
	}
	sb.WriteByte('\n')
	sb.WriteString(` right:`)
	sb.WriteString(fmt.Sprintf("%+v", pg.exits.right))
	for _, exit := range pg.additionalExits.right {
		sb.WriteString(fmt.Sprintf(", %+v", exit))
	}
	sb.WriteByte('\n')
	sb.WriteString(` bottom:`)
	sb.WriteString(fmt.Sprintf("%+v", pg.exits.bottom))
	for _, exit := range pg.additionalExits.bottom {
		sb.WriteString(fmt.Sprintf(", %+v", exit))
	}
	sb.WriteByte('\n')
	sb.WriteString(` left:`)
	sb.WriteString(fmt.Sprintf("%+v", pg.exits.left))
	for _, exit := range pg.additionalExits.left {
		sb.WriteString(fmt.Sprintf(", %+v", exit))
	}
	sb.WriteByte('\n')

	sb.WriteString(`Max Distance: `)
	sb.WriteString(fmt.Sprintf("%d", pg.maxDistance))
	sb.WriteByte('\n')

	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] < 0 {
				sb.WriteByte('#')
				// } else if pg.distances[ri][ci]%2 == 0 {
				// 	sb.WriteByte('E')
				// } else {
				// 	sb.WriteByte('O')
			} else {
				sb.WriteByte('0' + byte(pg.distances[ri][ci]%10))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
