package fourteen

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (int, error) {
	g, err := getGrid(input)
	if err != nil {
		return 0, err
	}

	units := 0
	for g.addSand(0, 500) {
		units++
	}

	return units, nil
}

type coord struct {
	x, y int
}

type material uint8

const (
	air  material = 0
	rock material = 1
	sand material = 2
)

type grid struct {

	// part 2:
	// min: {x:342 y:0}
	// max: {x:658 y:159}
	// [row][col]
	mats [159][660]material

	maxRock int
	floor   int
}

func newGrid() grid {
	return grid{
		maxRock: 0,
		floor:   -1,
	}
}

func (g *grid) String() string {
	tl, br := g.window()
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%3d", tl.x))
	sb.WriteString(`--v`)
	for c := tl.x + 1; c < br.x; c++ {
		sb.WriteByte(' ')
	}
	sb.WriteString(`v--`)
	sb.WriteString(fmt.Sprintf("%3d", br.x))
	sb.WriteByte('\n')

	for r := tl.y; r <= br.y; r++ {
		sb.WriteString(fmt.Sprintf("%3d", r))
		sb.WriteByte(' ')
		sb.WriteByte(':')
		for c := tl.x; c <= br.x; c++ {
			switch g.check(c, r) {
			case air:
				sb.WriteByte('.')
			case rock:
				sb.WriteByte('#')
			case sand:
				sb.WriteByte('o')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g *grid) window() (coord, coord) {
	var br coord
	tl := coord{
		x: -1,
		y: -1,
	}
	br.y = g.floor
	for r := range g.mats {
		for c := range g.mats[r] {
			switch g.check(c, r) {
			case rock, sand:
				if tl.x == -1 || tl.x > c {
					tl.x = c
				}
				if tl.y == -1 || tl.y > r {
					tl.y = r
				}
				if br.x < c {
					br.x = c
				}
				if br.y < r {
					br.y = r
				}
			}
		}
	}

	return tl, br
}

func (g *grid) addRock(y, x int) {
	// if g.floor >= 0 {
	// 	panic(`should not have set floor before finishing adding rock`)
	// }
	if y > g.maxRock {
		g.maxRock = y
	}

	g.mats[y][x] = rock
}

func (g *grid) addFloor() {
	g.floor = g.maxRock + 2
}

func (g *grid) addSand(y, x int) bool {
	y1 := y + 1
	if g.floor < 0 {
		if y > g.maxRock {
			return false
		}
	} else if y1 == g.floor {
		// it has come to rest
		g.mats[y][x] = sand
		return true
	}

	if g.check(y1, x) == air {
		// check from below
		return g.addSand(y1, x)
	}

	if g.check(y1, x-1) == air {
		// check down to the left
		return g.addSand(y1, x-1)
	}

	if g.check(y1, x+1) == air {
		// check down to the right
		return g.addSand(y1, x+1)
	}

	// it has come to rest
	g.mats[y][x] = sand
	return true
}

func (g *grid) check(y, x int) material {
	return g.mats[y][x]
}
