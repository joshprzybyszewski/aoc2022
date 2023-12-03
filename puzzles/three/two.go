package three

import (
	"strings"
)

type gears map[coord][]int

func newGears() gears {
	return make(gears)
}

func (g gears) addGear(row, col int, c byte) {
	if c != '*' {
		return
	}
	g[coord{
		row: row,
		col: col,
	}] = nil
}

func (g gears) addPart(row, col int, val int) {
	c := coord{
		row: row,
		col: col,
	}
	if _, ok := g[c]; !ok {
		return
	}
	g[c] = append(g[c], val)
}

func (s gears) nearbyGears(row, col int) nearbyGears {
	output := make(nearbyGears, 8)
	// check left, then below, then right, then above
	tmp := coord{
		row: row,
		col: col - 1,
	}
	_, ok := s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.row--
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		output[tmp] = struct{}{}
	}
	return output
}

type nearbyGears map[coord]struct{}

func (g nearbyGears) add(other nearbyGears) nearbyGears {
	if g == nil {
		return other
	}
	for k := range other {
		g[k] = struct{}{}
	}
	return g
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
	var nearby nearbyGears
	row = 0
	input = fullInput
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col < nli; col++ {
			if input[col] < '0' || input[col] > '9' {
				if nearby != nil {
					for ng := range nearby {
						g.addPart(ng.row, ng.col, curNum)
					}
					nearby = nil
				}
				curNum = 0
				continue
			}
			curNum *= 10
			curNum += int(input[col] - '0')
			nearby = nearby.add(g.nearbyGears(row, col))
		}
		input = input[nli+1:]
		row++
	}

	total := 0

	for _, nums := range g {
		if len(nums) != 2 {
			continue
		}

		total += (nums[0] * nums[1])
	}

	return total, nil
}
