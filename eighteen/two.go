package eighteen

import (
	"strconv"
	"strings"
)

type coord struct {
	x, y, z int
}

func (c coord) left() (coord, bool) {
	if c.x == 0 {
		return coord{}, false
	}
	c.x--
	return c, true
}

func (c coord) right() (coord, bool) {
	if c.x == dropletSideLength-1 {
		return coord{}, false
	}
	c.x++
	return c, true
}

func (c coord) down() (coord, bool) {
	if c.y == 0 {
		return coord{}, false
	}
	c.y--
	return c, true
}

func (c coord) up() (coord, bool) {
	if c.y == dropletSideLength-1 {
		return coord{}, false
	}
	c.y++
	return c, true
}

func (c coord) front() (coord, bool) {
	if c.z == 0 {
		return coord{}, false
	}
	c.z--
	return c, true
}

func (c coord) back() (coord, bool) {
	if c.z == dropletSideLength-1 {
		return coord{}, false
	}
	c.z++
	return c, true
}

type space [dropletSideLength][dropletSideLength][dropletSideLength]bool

func (s *space) fill(
	x, y, z int,
) {
	s[x][y][z] = true
}

func (s *space) isFilled(
	c coord,
) bool {
	return s[c.x][c.y][c.z]
}

func Two(
	input string,
) (int, error) {

	s := space{}

	var i1, i2 int
	var x, y, z int
	var err error

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		i1 = 0
		i2 = strings.Index(input, `,`)
		x, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}

		i1 = i2 + 1
		i2 = i1 + strings.Index(input[i1:], `,`)
		y, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}

		i1 = i2 + 1
		i2 = nli // i1 + strings.Index(input[i1:], `,`)
		z, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}
		// translate this filled space one over in each direction.
		// This allows us to do a BFS along the outside of the droplet starting at 0,0,0
		s.fill(x+1, y+1, z+1)

		input = input[nli+1:]
	}

	// 2008, 2011 is too low
	return getSurfaceArea(&s), nil
}

func getSurfaceArea(
	s *space,
) int {
	pending := make([]coord, 0, dropletSideLength*dropletSideLength*dropletSideLength)
	sc := coord{}
	if s.isFilled(sc) {
		panic(`should not be filled`)
	}
	pending = append(pending, sc)

	var seen space
	total := 0
	var c2 coord
	var ok bool

	// Do a 3d BFS around the outside of the droplet
	for len(pending) > 0 {
		c := pending[0]
		pending = pending[1:]

		if seen.isFilled(c) {
			// already seen this coord. no work to do
			continue
		}
		// mark this coord as seen
		seen.fill(c.x, c.y, c.z)

		// add up the outside edges of the space

		c2, ok = c.left()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}

		c2, ok = c.right()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}

		c2, ok = c.up()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}

		c2, ok = c.down()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}

		c2, ok = c.front()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}

		c2, ok = c.back()
		if ok {
			if s.isFilled(c2) {
				total++
			} else {
				pending = append(pending, c2)
			}
		}
	}

	return total
}
