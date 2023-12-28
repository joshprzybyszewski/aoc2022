package twentyone

import (
	"fmt"
	"strings"
)

const (
	gridSize = 131
)

func One(
	input string,
) (int, error) {
	g := newGarden(input)

	dijkstra(&g, 64)

	return g.numEven(), nil
}

type tile uint8

const (
	unset tile = 0
	odd   tile = 1
	even  tile = 2
	rocks tile = 3
)

func (t tile) toByte() byte {
	switch t {
	case unset:
		return '.'
	case odd:
		return 'O'
	case even:
		return 'E'
	case rocks:
		return '#'
	}
	panic(`unexpected`)
}

type coord struct {
	row int
	col int
}

func (c coord) up() coord {
	c.row--
	return c
}

func (c coord) down() coord {
	c.row++
	return c
}

func (c coord) left() coord {
	c.col--
	return c
}

func (c coord) right() coord {
	c.col++
	return c
}

type garden struct {
	tiles [gridSize][gridSize]tile

	start coord
}

func newGarden(input string) garden {
	ri, ci := 0, 0
	var g garden

	for len(input) > 0 {
		if input[0] == '\n' {
			ri++
			ci = -1
		} else {
			switch input[0] {
			case '.':
				// g.tiles[ri][ci] = unset
			case '#':
				g.tiles[ri][ci] = rocks
			case 'S':
				g.start.row = ri
				g.start.col = ci
				// g.tiles[ri][ci] = unset
			default:
				panic(`unexpected`)
			}
		}
		ci++
		input = input[1:]
	}

	return g
}

func (g garden) String() string {
	var sb strings.Builder

	for ri := 0; ri < gridSize; ri++ {
		for ci := 0; ci < gridSize; ci++ {
			sb.WriteByte(g.tiles[ri][ci].toByte())
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g *garden) numEven() int {
	total := 0

	for ri := 0; ri < gridSize; ri++ {
		for ci := 0; ci < gridSize; ci++ {
			if g.tiles[ri][ci] == even {
				total++
			}
		}
	}
	return total
}

func dijkstra(
	g *garden,
	maxDepth int,
) {
	type pos struct {
		coord
		tile  tile
		depth int
	}
	pending := make([]pos, 0, 4096)

	pending = append(pending, pos{
		coord: g.start,
		tile:  even,
	})

	for len(pending) > 0 {
		c := pending[0]
		pending = pending[1:]

		if g.tiles[c.row][c.col] != unset {
			if g.tiles[c.row][c.col] != c.tile {
				fmt.Printf("expected: %d, actual: %d", c.tile, g.tiles[c.row][c.col])
				panic(`unexpected`)
			}
			continue
		}

		if c.depth%2 == 0 {
			if c.tile != even {
				panic(`unexpected`)
			}
		} else {
			if c.tile != odd {
				panic(`unexpected`)
			}
		}

		g.tiles[c.row][c.col] = c.tile
		if c.depth == maxDepth {
			continue
		} else if c.depth > maxDepth {
			panic(`unexpected`)
		}

		next := odd
		if c.tile == odd {
			next = even
		}

		if c.row > 0 && g.tiles[c.row-1][c.col] == unset {
			pending = append(pending, pos{
				coord: coord{
					row: c.row - 1,
					col: c.col,
				},
				tile:  next,
				depth: c.depth + 1,
			})
		}
		if c.col > 0 && g.tiles[c.row][c.col-1] == unset {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col - 1,
				},
				tile:  next,
				depth: c.depth + 1,
			})
		}
		if c.row < gridSize-1 && g.tiles[c.row+1][c.col] == unset {
			pending = append(pending, pos{
				coord: coord{
					row: c.row + 1,
					col: c.col,
				},
				tile:  next,
				depth: c.depth + 1,
			})
		}
		if c.col < gridSize-1 && g.tiles[c.row][c.col+1] == unset {
			pending = append(pending, pos{
				coord: coord{
					row: c.row,
					col: c.col + 1,
				},
				tile:  next,
				depth: c.depth + 1,
			})
		}
	}
}
