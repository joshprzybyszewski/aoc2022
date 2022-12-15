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

func (g grid) possibleSources(
	dest coord,
) []coord {
	output := make([]coord, 0, 4)
	minSrcVal := g[dest.row][dest.col] - 1
	if dest.row > 0 && g[dest.row-1][dest.col] >= minSrcVal {
		output = append(output, coord{
			row: dest.row - 1,
			col: dest.col,
		})
	}
	if dest.row < len(g)-1 && g[dest.row+1][dest.col] >= minSrcVal {
		output = append(output, coord{
			row: dest.row + 1,
			col: dest.col,
		})
	}
	if dest.col > 0 && g[dest.row][dest.col-1] >= minSrcVal {
		output = append(output, coord{
			row: dest.row,
			col: dest.col - 1,
		})
	}
	if dest.col < len(g[dest.row])-1 && g[dest.row][dest.col+1] >= minSrcVal {
		output = append(output, coord{
			row: dest.row,
			col: dest.col + 1,
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
	steps := paint(g, e, s)
	return steps[s.row][s.col], nil
}

func paint(
	g grid,
	zero coord,
	target coord,
) [][]int {
	output := make([][]int, len(g))
	for i := range output {
		output[i] = make([]int, len(g[i]))
		for j := range output[i] {
			output[i][j] = -1
		}
	}
	output[zero.row][zero.col] = 0

	pending := make([]coord, 0, len(g)*len(g[0]))
	pending = append(pending, zero)

	for len(pending) > 0 {
		val := output[pending[0].row][pending[0].col] + 1
		for _, s := range g.possibleSources(pending[0]) {
			if output[s.row][s.col] != -1 {
				if output[s.row][s.col] > val {
					panic(`should not have found a shorter path`)
				}
				// we already know about a path that takes fewer steps to get to the zero coord
				continue
			}
			output[s.row][s.col] = val
			pending = append(pending, s)
		}
		if pending[0] == target {
			break
		}
		pending = pending[1:]
	}

	return output
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
