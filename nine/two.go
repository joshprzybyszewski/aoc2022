package nine

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {

	lines := strings.Split(input, "\n")

	lr := longRope{}
	var q, i int
	var err error

	pos := make(map[coord]struct{}, len(lines))
	pos[lr.knots[len(lr.knots)-1]] = struct{}{}

	for _, line := range lines {
		if line == `` {
			continue
		}
		q, err = strconv.Atoi(line[2:])
		if err != nil {
			return ``, err
		}

		for i = 0; i < q; i++ {
			lr = moveLongRope(lr, direction(line[0]))
			pos[lr.knots[len(lr.knots)-1]] = struct{}{}
		}
	}

	return strconv.Itoa(len(pos)), nil
}

type longRope struct {
	knots [10]coord
}

func moveLongRope(
	lr longRope,
	d direction,
) longRope {
	switch d {
	case right:
		lr.knots[0].x++
	case left:
		lr.knots[0].x--
	case up:
		lr.knots[0].y++
	case down:
		lr.knots[0].y--
	}

	for i := 1; i < len(lr.knots); i++ {
		lr.knots[i] = moveCoord(lr.knots[i], lr.knots[i-1])
	}

	return lr
}

func moveCoord(
	c, goal coord,
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
