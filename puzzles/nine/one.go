package nine

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	var head, tail coord
	var q, i int
	var err error

	/*
		When we start at (0,0):
			max: {x:52 y:166}
			min: {x:-115 y:-115}
	*/
	head.x = 115
	tail.x = 115
	head.y = 115
	tail.y = 115

	/*
		168 = 52-(-115) + 1
		282 = 166-(-115)+1
	*/
	var seen [168][282]bool

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}
		q, err = strconv.Atoi(input[2:nli])
		if err != nil {
			return 0, err
		}

		for i = 0; i < q; i++ {
			switch direction(input[0]) {
			case right:
				head.x++
			case left:
				head.x--
			case up:
				head.y++
			case down:
				head.y--
			}

			tail.moveToward(head)
			seen[tail.x][tail.y] = true
		}
		input = input[nli+1:]
	}

	total := 0
	for i := range seen {
		for j := range seen[i] {
			if seen[i][j] {
				total++
			}
		}
	}

	return total, nil
}

type coord struct {
	x, y int16
}

type direction byte

const (
	right direction = 'R'
	up    direction = 'U'
	left  direction = 'L'
	down  direction = 'D'
)

func (c *coord) moveToward(
	goal coord,
) {
	if goal.x == c.x { // same column
		// move up or down, or not at all
		if goal.y > c.y+1 {
			c.y++
		} else if goal.y < c.y-1 {
			c.y--
		}
		return
	}
	if goal.y == c.y { // same row
		// move left or right, or not at all
		if goal.x > c.x+1 {
			c.x++
		} else if goal.x < c.x-1 {
			c.x--
		}
		return
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
}
