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

	rowEmptysBefore [140]int
	colEmptysBefore [140]int

	universes [431]coord
}

func newUniverse(
	input string,
) universe {
	u := universe{}

	ri, ci := 0, 0
	ui := 0

	var rowsWith, colsWith [140]bool

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for ci = 0; ci < nli; ci++ {
			if input[ci] == '.' {
				continue
			}
			u.tiles[ri][ci] = true
			rowsWith[ri] = true
			colsWith[ci] = true

			u.universes[ui] = coord{
				row: ri,
				col: ci,
			}
			ui++
		}

		ri++
		input = input[nli+1:]
	}

	for ci = 1; ci < len(rowsWith); ci++ {
		u.rowEmptysBefore[ci] = u.rowEmptysBefore[ci-1]
		if !rowsWith[ci-1] {
			u.rowEmptysBefore[ci]++
		}

		u.colEmptysBefore[ci] = u.colEmptysBefore[ci-1]
		if !colsWith[ci-1] {
			u.colEmptysBefore[ci]++
		}
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

	end.row += (u.rowEmptysBefore[end.row] - u.rowEmptysBefore[start.row]) * expansionRate
	end.col += (u.colEmptysBefore[end.col] - u.colEmptysBefore[start.col]) * expansionRate

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
