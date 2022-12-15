package fourteen

import (
	"fmt"
	"strconv"
	"strings"
)

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
	// [row][col]
	mats    [][]material
	maxRock int
	floor   int
}

func newGrid() *grid {
	return &grid{
		mats:    make([][]material, 0, 256),
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

func (g *grid) addRock(x, y int) {
	if g.floor >= 0 {
		panic(`should not have set floor before finishing adding rock`)
	}
	if y > g.maxRock {
		g.maxRock = y
	}

	g.check(x, y)

	g.mats[y][x] = rock
}

func (g *grid) addFloor() {
	g.floor = g.maxRock + 2
}

func (g *grid) addSand(x, y int) bool {
	if g.floor < 0 && y > g.maxRock {
		return false
	}
	switch g.check(x, y+1) {
	case rock, sand:
		// can't fall straight down
	default:
		// check from below
		return g.addSand(x, y+1)
	}

	switch g.check(x-1, y+1) {
	case rock, sand:
		// can't fall down to the left
	default:
		// check down to the left
		return g.addSand(x-1, y+1)
	}

	switch g.check(x+1, y+1) {
	case rock, sand:
		// can't fall down to the right
	default:
		// check down to the right
		return g.addSand(x+1, y+1)
	}

	// it has come to rest
	g.mats[y][x] = sand
	return true
}

func (g *grid) check(x, y int) material {
	if g.floor >= 0 && y == g.floor {
		return rock
	}
	if y >= len(g.mats) {
		g.mats = append(g.mats, make([][]material, y-len(g.mats)+1)...)
	}
	if x >= len(g.mats[y]) {
		g.mats[y] = append(g.mats[y], make([]material, x-len(g.mats[y])+1)...)
	}
	return g.mats[y][x]
}

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")
	g, err := getGrid(lines)
	if err != nil {
		return 0, err
	}

	// fmt.Printf("got grid\n%s\n", g)

	units := 0
	for g.addSand(500, 0) {
		units++
	}

	// fmt.Printf("%s\n", g)

	return units, nil
}

func getGrid(lines []string) (*grid, error) {
	g := newGrid()

	for _, line := range lines {
		if line == `` {
			continue
		}
		coords, err := getCoords(line)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(coords)-1; i++ {
			c := coords[i]
			next := coords[i+1]
			for {
				g.addRock(c.x, c.y)

				if c.x < next.x {
					c.x++
				} else if c.x > next.x {
					c.x--
				} else if c.y < next.y {
					c.y++
				} else if c.y > next.y {
					c.y--
				} else {
					break
				}
			}
		}
	}

	return g, nil

}

func getCoords(line string) ([]coord, error) {
	ss := strings.Split(line, ` -> `)
	output := make([]coord, 0, len(ss))
	for _, s := range ss {
		ci := strings.Index(s, `,`)
		x, err := strconv.Atoi(s[:ci])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(s[ci+1:])
		if err != nil {
			return nil, err
		}

		output = append(output, coord{
			x: x,
			y: y,
		})
	}

	return output, nil
}
