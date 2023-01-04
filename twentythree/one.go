package twentythree

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (int, error) {

	elves := convertInputToElfLocations(input)
	// print(elves)
	for r := 0; r < 10; r++ {
		elves = getNextPositions(elves, r)
		// print(elves)
	}

	min, max := getBounds(elves)
	// fmt.Printf("min: %+v\n", min)
	// fmt.Printf("max: %+v\n", max)

	a := (max.x - min.x + 1) * (max.y - min.y + 1)

	return a - len(elves), nil
}

type coord struct {
	x, y int
}

func getNextPositions(
	elves []coord,
	roundIndex int,
) []coord {
	next := make([]coord, 0, len(elves))
	t := newKDTree(elves)

	var nearby []coord
	var n, p coord
	var right, down, left, up bool

	proposed := make(map[coord][]int)

	for i, e := range elves {
		right, down, left, up = true, true, true, true
		nearby = t.search(coordRange{
			x0: e.x - 1,
			x1: e.x + 1,
			y0: e.y - 1,
			y1: e.y + 1,
		})
		if len(nearby) == 1 {
			// none nearby: don't move.
			next = append(next, e)
			continue
		}

		for _, n = range nearby {
			if n == e {
				continue
			}
			if n.x == e.x+1 {
				right = false
			}
			if n.x == e.x-1 {
				left = false
			}
			if n.y == e.y-1 {
				up = false
			}
			if n.y == e.y+1 {
				down = false
			}
		}
		if !up && !down && !left && !right {
			// no valid moves. Don't propose
			next = append(next, e)
			continue
		}

		p = getProposal(
			e,
			roundIndex,
			up, down, left, right,
		)

		proposed[p] = append(proposed[p], i)
	}

	for c, indexes := range proposed {
		if len(indexes) == 1 {
			next = append(next, c)
			continue
		}
		for _, i := range indexes {
			next = append(next, elves[i])
		}
	}

	return next
}

func getBounds(elves []coord) (coord, coord) {
	min := coord{
		x: 74,
		y: 74,
	}
	var max coord
	for _, e := range elves {
		if e.x < min.x {
			min.x = e.x
		}
		if e.x > max.x {
			max.x = e.x
		}
		if e.y < min.y {
			min.y = e.y
		}
		if e.y > max.y {
			max.y = e.y
		}
	}
	return min, max
}

func getProposal(
	e coord,
	roundIndex int,
	up, down, left, right bool,
) coord {
	roundIndex %= 4
	if roundIndex == 0 {
		if up {
			e.y--
		} else if down {
			e.y++
		} else if left {
			e.x--
		} else if right {
			e.x++
		}
		return e
	}
	if roundIndex == 1 {
		if down {
			e.y++
		} else if left {
			e.x--
		} else if right {
			e.x++
		} else if up {
			e.y--
		}
		return e
	}
	if roundIndex == 2 {
		if left {
			e.x--
		} else if right {
			e.x++
		} else if up {
			e.y--
		} else if down {
			e.y++
		}
		return e
	}
	if right {
		e.x++
	} else if up {
		e.y--
	} else if down {
		e.y++
	} else if left {
		e.x--
	}
	return e
}

func print(
	elves []coord,
) {
	min, max := getBounds(elves)
	t := newKDTree(elves)

	var sb strings.Builder

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			n := t.search(coordRange{
				x0: x,
				x1: x,
				y0: y,
				y1: y,
			})
			if len(n) == 1 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
