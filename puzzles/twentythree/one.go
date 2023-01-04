package twentythree

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (int, error) {

	space, elves := convertInputToElfLocations(input)

	w := newWorkforce(&space, elves)
	w.start()
	defer w.stop()

	var ri uint8
	for r := 0; r < 10; r++ {
		_ = updateMap(&w, ri)
		ri++
		ri &= 3
	}

	min, max := getBounds(&space)

	a := (max.x - min.x + 1) * (max.y - min.y + 1)

	return int(a) - len(elves), nil
}

type coord struct {
	x, y int
}

type clears uint8

const (
	north clears = 1 << iota
	south
	east
	west

	northWest = north | west
	northEast = north | east
	southEast = south | east
	southWest = south | west

	notNorth = ^north
	notSouth = ^south
	notEast  = ^east
	notWest  = ^west

	notNorthWest = ^(north | west)
	notNorthEast = ^(north | east)
	notSouthEast = ^(south | east)
	notSouthWest = ^(south | west)

	allClear  clears = north | south | east | west
	noneClear clears = 0
)

func updateMap(
	w *workforce,
	roundIndex uint8,
) bool {
	proposals := w.run(roundIndex)
	if len(proposals) == 0 {
		// steady state achieved
		return true
	}

	for dst, ci := range proposals {
		if ci >= 0 {
			w.space[w.elves[ci].x][w.elves[ci].y] = false
			w.elves[ci] = dst
			w.space[dst.x][dst.y] = true
		}
	}

	return false
}

func getBounds(elves *space) (coord, coord) {
	min := coord{
		x: 74,
		y: 74,
	}
	var max coord
	// TODO be smarter about this.
	for x := range elves {
		for y := range elves[x] {
			if !elves[x][y] {
				continue
			}
			if x < min.x {
				min.x = x
			}
			if x > max.x {
				max.x = x
			}
			if y < min.y {
				min.y = y
			}
			if y > max.y {
				max.y = y
			}
		}
	}
	return min, max
}

func print(
	elves *space,
) {
	min, max := getBounds(elves)

	var sb strings.Builder

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if elves[x][y] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func printWithProposals(
	elves *space,
	proposals map[coord][]coord,
) {
	min, max := getBounds(elves)
	for dst := range proposals {
		if dst.x < min.x {
			min.x = dst.x
		}
		if dst.x > max.x {
			max.x = dst.x
		}
		if dst.y < min.y {
			min.y = dst.y
		}
		if dst.y > max.y {
			max.y = dst.y
		}
	}

	var sb strings.Builder

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			c := coord{
				x: x,
				y: y,
			}
			if elves[x][y] {
				if len(proposals[c]) > 0 {
					sb.WriteByte('?')
				} else {
					sb.WriteByte('#')
				}
			} else {
				if len(proposals[c]) > 0 {
					sb.WriteByte('0' + byte(len(proposals[c])))
				} else {
					sb.WriteByte('.')
				}
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
