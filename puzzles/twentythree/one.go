package twentythree

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func One(
	input string,
) (int, error) {

	elves := convertInputToElfLocations(input)
	es := make([]coord, len(elves))
	populateSlice(elves, es)
	var ri uint8
	for r := 0; r < 10; r++ {
		_ = updateMap(elves, es, ri)
		ri++
		ri &= 3
	}

	min, max := getBounds(elves)

	a := (max.x - min.x + 1) * (max.y - min.y + 1)

	return int(a) - len(es), nil
}

type coord struct {
	x, y int16
}

func (c coord) nw() coord {
	c.x--
	c.y--
	return c
}

func (c coord) n() coord {
	c.y--
	return c
}

func (c coord) ne() coord {
	c.x++
	c.y--
	return c
}

func (c coord) e() coord {
	c.x++
	return c
}

func (c coord) se() coord {
	c.x++
	c.y++
	return c
}

func (c coord) s() coord {
	c.y++
	return c
}

func (c coord) sw() coord {
	c.x--
	c.y++
	return c
}

func (c coord) w() coord {
	c.x--
	return c
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
	elves map[coord]bool,
	es []coord,
	roundIndex uint8,
) bool {
	proposals := map[coord]int{}
	var proposalsLock sync.Mutex

	checkElf := func(ci int) {
		c := es[ci]
		cl := allClear

		if elves[c.nw()] {
			cl &= southEast
		}
		if elves[c.se()] {
			cl &= northWest
		}

		if (cl&northEast != 0) && elves[c.ne()] {
			cl &= southWest
		}
		if (cl&southWest != 0) && elves[c.sw()] {
			cl &= northEast
		}

		if cl == noneClear {
			// already eliminated all directions. do nothing
			return
		}

		if cl&north == north && elves[c.n()] {
			cl &= notNorth
		}
		if cl&south == south && elves[c.s()] {
			cl &= notSouth
		}
		if cl&east == east && elves[c.e()] {
			cl &= notEast
		}
		if cl&west == west && elves[c.w()] {
			cl &= notWest
		}

		if cl == allClear || cl == noneClear {
			// do nothing
			return
		}

		p := c
		switch roundIndex {
		case 0:
			if cl&north == north {
				p.y--
			} else if cl&south == south {
				p.y++
			} else if cl&west == west {
				p.x--
			} else if cl&east == east {
				p.x++
			}
		case 1:
			if cl&south == south {
				p.y++
			} else if cl&west == west {
				p.x--
			} else if cl&east == east {
				p.x++
			} else if cl&north == north {
				p.y--
			}
		case 2:
			if cl&west == west {
				p.x--
			} else if cl&east == east {
				p.x++
			} else if cl&north == north {
				p.y--
			} else if cl&south == south {
				p.y++
			}
		case 3:
			if cl&east == east {
				p.x++
			} else if cl&north == north {
				p.y--
			} else if cl&south == south {
				p.y++
			} else if cl&west == west {
				p.x--
			}
		}
		if p == c {
			return
		}
		proposalsLock.Lock()
		defer proposalsLock.Unlock()
		if _, ok := proposals[p]; ok {
			proposals[p] = -1
		} else {
			proposals[p] = ci
		}
	}

	var wg sync.WaitGroup

	work := make(chan int, len(es))
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				checkElf(w)
				wg.Done()
			}
		}()
	}

	for i := range es {
		wg.Add(1)
		work <- i
	}

	wg.Wait()
	close(work)

	if len(proposals) == 0 {
		// steady state achieved
		return true
	}

	for dst, ci := range proposals {
		if ci >= 0 {
			elves[es[ci]] = false
			es[ci] = dst
			elves[dst] = true
		}
	}

	return false
}

func getBounds(elves map[coord]bool) (coord, coord) {
	min := coord{
		x: 74,
		y: 74,
	}
	var max coord
	for e, b := range elves {
		if !b {
			continue
		}
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

func print(
	elves map[coord]bool,
) {
	min, max := getBounds(elves)

	var sb strings.Builder

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if elves[coord{
				x: x,
				y: y,
			}] {
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
	elves map[coord]bool,
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
			if elves[c] {
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
