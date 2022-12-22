package seventeen

import (
	"strings"
)

const (
	numRocksPart1 = 2022

	// needs to be at least numRocksPart1*3 + (3 + 4)
	maxChamberHeight = 8192 // 0b10000000000000 = 1 << 13
)

type rock [4]uint8

func (r *rock) isFarLeft() bool {
	// return (r[0]|r[1]|r[2]|r[3])&leftwall != 0
	return (r[3]|r[2]|r[1]|r[0])&leftwall != 0
}

func (r *rock) isFarRight() bool {
	// return (r[0]|r[1]|r[2]|r[3])&rightwall != 0
	return (r[3]|r[2]|r[1]|r[0])&rightwall != 0
}

const (
	leftwall  uint8 = 0b01000000
	rightwall uint8 = 0b00000001
	fullrow   uint8 = 0b01111111

	dashIndex   uint8 = 0
	plusIndex   uint8 = 1
	cornerIndex uint8 = 2
	towerIndex  uint8 = 3
	squareIndex uint8 = 4
)

var (
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
	settled     [maxChamberHeight]uint8
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

	if c.settled[c.pending.bottom]&(c.pending.rock[0]<<1) != 0 {
		return false
	}
	// if c.pendingIndex == dashIndex {
	// 	// the dash has no other rows
	// 	return true
	// }
	if c.settled[c.pending.bottom+1]&(c.pending.rock[1]<<1) != 0 {
		return false
	}
	// if c.pendingIndex == squareIndex {
	// 	// the square has no other rows
	// 	return true
	// }
	if c.settled[c.pending.bottom+2]&(c.pending.rock[2]<<1) != 0 {
		return false
	}
	// if c.pendingIndex != towerIndex {
	// 	// only the tower has a fourth row
	// 	return true
	// }

	return c.settled[c.pending.bottom+3]&(c.pending.rock[3]<<1) == 0
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

	if c.settled[c.pending.bottom]&(c.pending.rock[0]>>1) != 0 {
		return false
	}
	// if c.pendingIndex == dashIndex {
	// 	// the dash has no other rows
	// 	return true
	// }
	if c.settled[c.pending.bottom+1]&(c.pending.rock[1]>>1) != 0 {
		return false
	}
	// if c.pendingIndex == squareIndex {
	// 	// the square has no other rows
	// 	return true
	// }
	if c.settled[c.pending.bottom+2]&(c.pending.rock[2]>>1) != 0 {
		return false
	}
	// if c.pendingIndex != towerIndex {
	// 	// only the tower has a fourth row
	// 	return true
	// }

	return c.settled[c.pending.bottom+3]&(c.pending.rock[3]>>1) == 0
}

func (c *chamber) fall() bool {
	if c.canFall() {
		c.pending.bottom--
		return true
	}

	// move pending to settled
	c.settled[c.pending.bottom] |= c.pending.rock[0]
	c.settled[c.pending.bottom+1] |= c.pending.rock[1]
	c.settled[c.pending.bottom+2] |= c.pending.rock[2]
	c.settled[c.pending.bottom+3] |= c.pending.rock[3]
	// keep track of the lowest row of empty rock
	for c.settled[c.minEmptyRow] != 0 {
		c.minEmptyRow++
	}
	// add a new pending
	c.addPendingRock()
	return false
}

func (c *chamber) canFall() bool {
	if c.pending.bottom == 0 {
		return false
	}
	return (c.settled[c.pending.bottom-1]&c.pending.rock[0])|
		(c.settled[c.pending.bottom]&c.pending.rock[1])|
		(c.settled[c.pending.bottom+1]&c.pending.rock[2])|
		(c.settled[c.pending.bottom+2]&c.pending.rock[3]) == 0
}

func (c *chamber) addPendingRock() {
	switch c.pendingIndex {
	case dashIndex:
		c.pending.rock = dash
		c.pendingIndex = plusIndex
	case plusIndex:
		c.pending.rock = plus
		c.pendingIndex = cornerIndex
	case cornerIndex:
		c.pending.rock = corner
		c.pendingIndex = towerIndex
	case towerIndex:
		c.pending.rock = tower
		c.pendingIndex = squareIndex
	case squareIndex:
		c.pending.rock = square
		c.pendingIndex = dashIndex
	}
	c.pending.bottom = c.minEmptyRow + 3
}

func (c *chamber) reduce() int {
	r := c.getHighestFullRow()
	if r < 0 {
		// no full rows
		return 0
	}
	numRows := r + 1

	// move the rows above the full row down to the bottom
	var b int
	for t := numRows; t < c.minEmptyRow; {
		c.settled[b] = c.settled[t]
		b++
		t++
	}
	// clear out rows above the top of the rocks
	for ; b < c.minEmptyRow; b++ {
		c.settled[b] = 0
	}
	// lower the "min empty row"
	c.minEmptyRow -= numRows
	// lower the pending rock's bottom index
	c.pending.bottom -= numRows

	// return how many rows this was reduced
	return numRows
}

func (c *chamber) getHighestFullRow() int {
	var r int
	for r = c.minEmptyRow - 1; r >= 0; r-- {
		if c.settled[r] == fullrow {
			return r
		}
	}
	return r
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

	return c.minEmptyRow, nil
}
