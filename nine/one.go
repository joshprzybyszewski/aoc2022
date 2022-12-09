package nine

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var head, tail coord
	var q, i int
	var err error

	pos := make(map[coord]struct{}, 4096)
	record := func() {
		pos[tail] = struct{}{}
	}
	record()

	for _, line := range lines {
		if line == `` {
			continue
		}
		q, err = strconv.Atoi(line[2:])
		if err != nil {
			return ``, err
		}

		for i = 0; i < q; i++ {
			switch direction(line[0]) {
			case right:
				head.x++
			case left:
				head.x--
			case up:
				head.y++
			case down:
				head.y--
			}

			tail = moveCoord(tail, head)
			record()
		}
	}

	return strconv.Itoa(len(pos)), nil
}

type coord struct {
	x, y int
}

type direction byte

const (
	right direction = 'R'
	up    direction = 'U'
	left  direction = 'L'
	down  direction = 'D'
)

func moveCoord(
	c coord,
	goal coord,
) coord {
	if goal.x == c.x { // same column
		// move up or down, or not at all
		if goal.y > c.y+1 {
			c.y++
		} else if goal.y < c.y-1 {
			c.y--
		}
		return c
	}
	if goal.y == c.y { // same row
		// move left or right, or not at all
		if goal.x > c.x+1 {
			c.x++
		} else if goal.x < c.x-1 {
			c.x--
		}
		return c
	}

	// different row and different column: move diagonally, if at all
	if goal.y > c.y+1 {
		c.y++
		if goal.x > c.x {
			c.x++
		} else {
			c.x--
		}
	} else if goal.y < c.y-1 {
		c.y--
		if goal.x > c.x {
			c.x++
		} else {
			c.x--
		}
	} else if goal.x > c.x+1 {
		c.x++
		if goal.y > c.y {
			c.y++
		} else {
			c.y--
		}
	} else if goal.x < c.x-1 {
		c.x--
		if goal.y > c.y {
			c.y++
		} else {
			c.y--
		}
	}

	return c
}
