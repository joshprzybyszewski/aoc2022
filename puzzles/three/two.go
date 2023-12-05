package three

import (
	"strings"
)

type gears map[coord][3]int

func newGears() gears {
	return make(gears, 372)
}

func (g gears) addGear(row, col int, c byte) {
	if c != '*' {
		return
	}
	g[coord{
		row: row,
		col: col,
	}] = [3]int{}
}

func (g gears) addPart(row, col int, val int) {
	c := coord{
		row: row,
		col: col,
	}
	data, ok := g[c]
	if !ok {
		return
	}
	if data[0] == 0 {
		data[0] = val
	} else if data[1] == 0 {
		data[1] = val
	} else {
		data[2] = val // we don't expect more than two parts per symbol.
	}
	g[c] = data
}

func (s gears) addNearbyGears(
	ng nearbyGears,
	row, col int,
) {
	// check left, then below, then right, then above
	tmp := coord{
		row: row,
		col: col - 1,
	}
	_, ok := s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.row--
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		ng[tmp] = struct{}{}
	}
}

type nearbyGears map[coord]struct{}

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
	nearby := make(nearbyGears, 32)
	var ng coord
	row = 0
	input = fullInput
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col < nli; col++ {
			if input[col] < '0' || input[col] > '9' {
				for ng = range nearby {
					g.addPart(ng.row, ng.col, curNum)
					delete(nearby, ng)
				}
				curNum = 0
				continue
			}
			curNum *= 10
			curNum += int(input[col] - '0')
			g.addNearbyGears(nearby, row, col)
		}
		input = input[nli+1:]
		row++
	}

	total := 0

	for _, nums := range g {
		if nums[2] != 0 {
			continue
		}

		total += (nums[0] * nums[1])
	}

	return total, nil
}
