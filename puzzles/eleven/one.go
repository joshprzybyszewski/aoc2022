package eleven

import "strings"

const (
	double = 1
)

type coord struct {
	row int
	col int
}

type universe struct {
	tiles [140][140]bool

	rowsWith [140]bool
	colsWith [140]bool

	universes [431]coord
}

func newUniverse(
	input string,
) universe {
	u := universe{}

	ri, ci := 0, 0
	ui := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for ci = 0; ci < nli; ci++ {
			if input[ci] == '.' {
				continue
			}
			u.tiles[ri][ci] = true
			u.rowsWith[ri] = true
			u.colsWith[ci] = true

			u.universes[ui] = coord{
				row: ri,
				col: ci,
			}
			ui++
		}

		ri++
		input = input[nli+1:]
	}

	return u
}

func (u *universe) shortestPath(
	i, j int,
	expansionRate int,
) int {
	start := u.universes[i]
	end := u.universes[j]
	if end.col < start.col {
		end.col, start.col = start.col, end.col
	}

	numExpanded := 0
	tmp := 0
	for tmp = start.row + 1; tmp < end.row; tmp++ {
		if !u.rowsWith[tmp] {
			numExpanded += expansionRate
		}
	}
	end.row += numExpanded
	numExpanded = 0

	for tmp = start.col + 1; tmp < end.col; tmp++ {
		if !u.colsWith[tmp] {
			numExpanded += expansionRate
		}
	}
	end.col += numExpanded

	return (end.col - start.col) + (end.row - start.row)
}

func One(
	input string,
) (int, error) {
	answer := solveForExpansion(input, double)
	return answer, nil
}

func solveForExpansion(
	input string,
	expansion int,
) int {
	u := newUniverse(input)

	total := 0

	var j int
	for i := 0; i < len(u.universes); i++ {
		for j = i + 1; j < len(u.universes); j++ {
			total += u.shortestPath(i, j, expansion)
		}
	}
	return total
}
