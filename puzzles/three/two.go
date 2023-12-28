package three

import (
	"strings"
)

const (
	size = 140
)

type gears struct {
	nearby [size][size][3]int

	gears [size][size]bool

	gearCoords [400]coord
	gci        int
}

func newGears() gears {
	return gears{}
}

func (g *gears) addGear(row, col int, c byte) {
	if c != '*' {
		return
	}

	g.gears[row][col] = true
	g.gearCoords[g.gci] = coord{
		row: row,
		col: col,
	}
	g.gci++
}

func (g *gears) addPart(row, col int, val int) {
	if val == 0 {
		return
	}

	minCol := col - 2
	if val > 99 {
		minCol -= 2
	} else if val > 9 {
		minCol--
	}
	if minCol < 0 {
		minCol = 0
	}

	if col == size {
		col--
	}

	g.addPartValue(row, col, val)
	g.addPartValue(row, minCol, val)

	var tmpCol int

	if row > 0 {
		for tmpCol = col; tmpCol >= minCol; tmpCol-- {
			g.addPartValue(row-1, tmpCol, val)
		}
	}

	if row+1 < size {
		for tmpCol = col; tmpCol >= minCol; tmpCol-- {
			g.addPartValue(row+1, tmpCol, val)
		}
	}
}

func (g *gears) addPartValue(row, col int, val int) {
	if !g.gears[row][col] {
		return
	}

	if g.nearby[row][col][0] == 0 {
		g.nearby[row][col][0] = val
	} else if g.nearby[row][col][1] == 0 {
		g.nearby[row][col][1] = val
	} else {
		g.nearby[row][col][2] = val // we don't expect more than two parts per symbol.
	}
}

func Two(
	fullInput string,
) (int, error) {
	input := fullInput

	var row, col int
	g := newGears()

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col < nli; col++ {
			g.addGear(row, col, input[col])
		}
		row++
		input = input[nli+1:]
	}

	curNum := 0
	row = 0
	input = fullInput
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col <= nli; col++ {
			if input[col] < '0' || input[col] > '9' {
				g.addPart(row, col, curNum)
				curNum = 0
				continue
			}
			curNum *= 10
			curNum += int(input[col] - '0')
		}
		input = input[nli+1:]
		row++
	}

	total := 0

	var nums [3]int
	for i := 0; i < g.gci; i++ {
		nums = g.nearby[g.gearCoords[i].row][g.gearCoords[i].col]
		if nums[2] != 0 {
			continue
		}

		total += (nums[0] * nums[1])
	}

	return total, nil
}
