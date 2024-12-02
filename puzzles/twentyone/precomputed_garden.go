package twentyone

import (
	"fmt"
	"strings"
)

type precomputedGarden struct {
	distances [gridSize][gridSize]int

	numEven     int
	numOdd      int
	minDistance int
	maxDistance int
}

func newPrecomputedInitialGarden(
	gar garden,
) precomputedGarden {
	return newPrecomputedGardenWithStarts(
		&gar,
		map[coord]int{
			gar.start: 0,
		},
	)
}

func newPrecomputedGardenWithStarts(
	gar *garden,
	starts map[coord]int,
) precomputedGarden {
	pg := precomputedGarden{}

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

	for c := range starts {
		c := c
		pending = append(pending, pos{
			coord: c,
			depth: starts[c],
		})
	}

	for len(pending) > 0 {
		c := pending[0]
		pending = pending[1:]

		if pg.distances[c.row][c.col] > 0 {
			if c.depth >= pg.distances[c.row][c.col] {
				continue
			}
			imOdd := c.depth%2 == 1
			wasOdd := pg.distances[c.row][c.col]%2 == 1
			if imOdd != wasOdd {
				panic(`unexpected`)
			}
		}

		pg.distances[c.row][c.col] = c.depth

		if c.row == 0 {
			// do nothing
		} else if pg.distances[c.row-1][c.col] != -1 { // c.row MUST be > 0
			pending = append(pending, pos{
				coord: coord{
					row: c.row - 1,
					col: c.col,
				},
				depth: c.depth + 1,
			})
		}

		if c.col == 0 {
			// do nothing
		} else if pg.distances[c.row][c.col-1] != -1 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col - 1,
				},
				depth: c.depth + 1,
			})
		}
		if c.row == gridSize-1 {
			// do nothing
		} else if pg.distances[c.row+1][c.col] != -1 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row + 1,
					col: c.col,
				},
				depth: c.depth + 1,
			})
		}
		if c.col == gridSize-1 {
			// do nothing
		} else if pg.distances[c.row][c.col+1] != -1 {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col + 1,
				},
				depth: c.depth + 1,
			})
		}
	}

	pg.minDistance = -1
	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] > pg.maxDistance {
				pg.maxDistance = pg.distances[ri][ci]
			}
			if pg.distances[ri][ci] > 0 && (pg.minDistance == -1 || pg.distances[ri][ci] < pg.minDistance) {
				pg.minDistance = pg.distances[ri][ci]
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

func (pg precomputedGarden) lines(
	maxDepth int,
) [gridSize]string {
	var sb [gridSize]strings.Builder
	var output [gridSize]string

	for ri := range pg.distances {
		for ci := range pg.distances[ri] {
			if pg.distances[ri][ci] > maxDepth {
				sb[ri].WriteByte('.')
			} else if pg.distances[ri][ci] < 0 {
				sb[ri].WriteByte('#')
				// } else if pg.distances[ri][ci]%2 == 0 {
				// 	sb.WriteByte('E')
				// } else {
				// 	sb.WriteByte('O')
			} else {
				sb[ri].WriteByte('0' + byte(pg.distances[ri][ci]%10))
			}
		}
		output[ri] = sb[ri].String()
	}
	return output
}
