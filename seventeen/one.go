package seventeen

import (
	"strings"
)

const (
	numRocksPart1 = 2022

	maxHeightPart1 = numRocksPart1*3 + (3 + 4)
)

type rock [4]uint8

func (r rock) isFarLeft() bool {
	return (r[3]&leftwall)|
		(r[2]&leftwall)|
		(r[1]&leftwall)|
		(r[0]&leftwall) > 0
}

func (r rock) isFarRight() bool {
	return (r[3]&rightwall)|
		(r[2]&rightwall)|
		(r[1]&rightwall)|
		(r[0]&rightwall) > 0
}

var (
	leftwall  uint8 = 0b01000000
	rightwall uint8 = 0b00000001

	// ####
	dash = rock{
		0b00011110, // ?0011110
		0,          // ?0000000
		0,          // ?0000000
		0,          // ?0000000
	}

	// .#.
	// ###
	// .#.
	plus = rock{
		0b00001000, // ?0001000
		0b00011100, // ?0011100
		0b00001000, // ?0001000
		0,          // ?0000000
	}

	// ..#
	// ..#
	// ###
	corner = rock{
		0b00011100, // ?0011100
		0b00000100, // ?0000100
		0b00000100, // ?0000100
		0,          // ?0000000
	}

	// #
	// #
	// #
	// #
	tower = rock{
		0b00010000, // ?0010000
		0b00010000, // ?0010000
		0b00010000, // ?0010000
		0b00010000, // ?0010000
	}

	// ##
	// ##
	square = rock{
		0b00011000, // ?0011000
		0b00011000, // ?0011000
		0,          // ?0000000
		0,          // ?0000000
	}
)

type fallingRock struct {
	rock   rock
	bottom int
}

type chamber struct {
	settled     [maxHeightPart1]uint8
	minEmptyRow int

	pending      fallingRock
	pendingIndex uint8
}

func newChamber() chamber {
	c := chamber{}

	c.addPendingRock()

	return c
}

func (c *chamber) String() string {
	var sb strings.Builder
	for r := c.minEmptyRow + 7; r >= 0; r-- {
		if r > c.minEmptyRow &&
			r > c.pending.bottom && (r >= c.pending.bottom+len(c.pending.rock) ||
			c.pending.rock[r-c.pending.bottom] == 0) {
			continue
		}

		sb.WriteString("|")
		for col := 6; col >= 0; col-- {
			if c.settled[r]&(1<<col) != 0 {
				sb.WriteByte('#')
			} else if r >= c.pending.bottom &&
				r < c.pending.bottom+len(c.pending.rock) &&
				(c.pending.rock[r-c.pending.bottom]&(1<<col) != 0) {
				sb.WriteByte('@')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("+-------+\n")

	return sb.String()
}

func (c *chamber) pushLeft() {
	if !c.canPushLeft() {
		return
	}
	// push each of the rows of the rock to the left
	for i := range c.pending.rock {
		c.pending.rock[i] = c.pending.rock[i] << 1
	}
}

func (c *chamber) canPushLeft() bool {
	if c.pending.rock.isFarLeft() {
		return false
	}

	for i := range c.pending.rock {
		if c.settled[c.pending.bottom+i]&(c.pending.rock[i]<<1) != 0 {
			return false
		}
	}

	return true
}

func (c *chamber) pushRight() {
	if !c.canPushRight() {
		return
	}
	// push each of the rows of the rock to the right
	for i := range c.pending.rock {
		c.pending.rock[i] = c.pending.rock[i] >> 1
	}
}

func (c *chamber) canPushRight() bool {
	if c.pending.rock.isFarRight() {
		return false
	}

	for i := range c.pending.rock {
		if c.settled[c.pending.bottom+i]&(c.pending.rock[i]>>1) != 0 {
			return false
		}
	}

	return true
}

func (c *chamber) fall() bool {
	if c.canFall() {
		c.pending.bottom--
		return true
	}

	// move pending to settled
	for i := range c.pending.rock {
		c.settled[c.pending.bottom+i] |= c.pending.rock[i]
	}
	// keep track of the lowest row of empty rock
	for ; c.settled[c.minEmptyRow] > 0; c.minEmptyRow++ {
	}
	// add a new pending
	c.pendingIndex++
	c.pendingIndex %= 5
	c.addPendingRock()
	return false
}

func (c *chamber) canFall() bool {
	if c.pending.bottom == 0 {
		return false
	}

	r := c.pending.bottom - 1
	for i := range c.pending.rock {
		if c.settled[r+i]&c.pending.rock[i] != 0 {
			return false
		}
	}

	return true
}

func (c *chamber) addPendingRock() {
	switch c.pendingIndex {
	case 0:
		c.pending.rock = dash
	case 1:
		c.pending.rock = plus
	case 2:
		c.pending.rock = corner
	case 3:
		c.pending.rock = tower
	case 4:
		c.pending.rock = square
	}
	c.pending.bottom = c.minEmptyRow + 3
}

func One(
	input string,
) (int, error) {
	c := newChamber()
	i := 0

	for nr := 0; nr < numRocksPart1; nr++ {
		// fmt.Printf("Starting %dth rock\n%s\n\n", nr, c.String())
		for {
			switch input[i] {
			case '<':
				c.pushLeft()
				// fmt.Printf("Pushed Left\n%s\n\n", c.String())
			case '>':
				c.pushRight()
				// fmt.Printf("Pushed Right\n%s\n\n", c.String())
			default:
				panic(input[i])
			}

			i++
			if i == len(input)-1 {
				i = 0
			}

			if !c.fall() {
				// fmt.Printf("Came to rest\n%s\n\n", c.String())
				break
			}
			// fmt.Printf("Fell\n%s\n\n", c.String())
		}
	}

	// 2978 is too low
	return c.minEmptyRow, nil
}
