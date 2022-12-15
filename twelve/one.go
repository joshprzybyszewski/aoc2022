package twelve

import (
	"strings"
)

// [row][col]
type grid [][]int

func newGrid(input string) (grid, coord, coord) {
	lines := strings.Split(input, "\n")
	var s, e coord
	g := make(grid, len(lines)-1)
	seen := make([][]bool, len(lines)-1)
	for i, line := range lines {
		if line == `` {
			continue
		}
		g[i] = make([]int, len(line))
		seen[i] = make([]bool, len(line))
		for j, b := range line {
			if b == 'E' {
				e = coord{
					row: i,
					col: j,
				}
				g[i][j] = 25
			} else if b == 'S' {
				s = coord{
					row: i,
					col: j,
				}
				g[i][j] = 0
			} else {
				g[i][j] = int(b - 'a')
			}
		}
	}
	return g, s, e
}

func (g grid) possibleMoves(
	c coord,
) []coord {
	output := make([]coord, 0, 4)
	if c.row > 0 && g[c.row-1][c.col] <= g[c.row][c.col]+1 {
		output = append(output, coord{
			row: c.row - 1,
			col: c.col,
		})
	}
	if c.row < len(g)-1 && g[c.row+1][c.col] <= g[c.row][c.col]+1 {
		output = append(output, coord{
			row: c.row + 1,
			col: c.col,
		})
	}
	if c.col > 0 && g[c.row][c.col-1] <= g[c.row][c.col]+1 {
		output = append(output, coord{
			row: c.row,
			col: c.col - 1,
		})
	}
	if c.col < len(g[c.row])-1 && g[c.row][c.col+1] <= g[c.row][c.col]+1 {
		output = append(output, coord{
			row: c.row,
			col: c.col + 1,
		})
	}
	return output
}

type coord struct {
	row, col int
}

type step struct {
	coord

	prev *step
}

func One(
	input string,
) (int, error) {
	g, s, e := newGrid(input)
	n := getStepsBetween(g, s, e)
	return n, nil
}

func getStepsBetween(
	g grid,
	s coord,
	e coord,
) int {
	seen := make([][]bool, len(g))
	for i := range g {
		seen[i] = make([]bool, len(g[i]))
	}

	pending := make([]*step, 0, len(g)*len(g[0]))
	pending = append(pending, &step{
		coord: s,
		prev:  nil,
	})

	var final *step
	for len(pending) > 0 && final == nil {
		ps := g.possibleMoves(pending[0].coord)
		for _, p := range ps {
			if seen[p.row][p.col] {
				continue
			}
			if p == e {
				final = &step{
					coord: p,
					prev:  pending[0],
				}
				break
			}
			seen[p.row][p.col] = true
			pending = append(pending, &step{
				coord: p,
				prev:  pending[0],
			})
		}
		pending = pending[1:]
	}

	n := 0
	for s := final; s != nil; n++ {
		s = s.prev
	}

	// it's the steps between, and the first step is the starting point
	return n - 1
}
