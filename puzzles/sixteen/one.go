package sixteen

const (
	size = 110
)

func One(
	input string,
) (int, error) {
	c := newContraption(input)
	c.sendLight(coord{
		row: 0,
		col: 0,
	}, right)

	return c.numEnergized(), nil
}

type coord struct {
	row uint8
	col uint8
}

type tile uint8

const (
	empty     tile = 0
	splitUp   tile = 1 << 0
	splitWide tile = 1 << 1
	reflect1  tile = 1 << 2 // /
	reflect2  tile = 1 << 3 // \
)

type lightTile uint8

const (
	left  lightTile = 1 << 0
	right lightTile = 1 << 1
	up    lightTile = 1 << 2
	down  lightTile = 1 << 3
)

func (lt lightTile) String() string {
	switch lt {
	case left:
		return `left`
	case right:
		return `right`
	case up:
		return `up`
	case down:
		return `down`
	}
	return `none`
}

func (lt lightTile) update(
	c coord,
	t tile,
) ([2]coord, [2]lightTile) {
	sendRight := false
	sendLeft := false
	sendDown := false
	sendUp := false

	switch t {
	case empty:
		sendRight = lt == right
		sendLeft = lt == left
		sendDown = lt == down
		sendUp = lt == up
	case splitUp:
		sendDown = lt&(down|left|right) == lt
		sendUp = lt&(up|left|right) == lt
	case splitWide:
		sendLeft = lt&(up|down|left) == lt
		sendRight = lt&(up|down|right) == lt
	case reflect1:
		sendRight = lt == up
		sendLeft = lt == down
		sendDown = lt == left
		sendUp = lt == right
	case reflect2:
		sendRight = lt == down
		sendLeft = lt == up
		sendDown = lt == right
		sendUp = lt == left

	}

	outI := 0
	outCoords := [2]coord{}
	outDirs := [2]lightTile{}

	if sendRight && c.col < size-1 {
		outCoords[outI] = c
		outCoords[outI].col++
		outDirs[outI] = right
		outI++
	}
	if sendLeft && c.col > 0 {
		outCoords[outI] = c
		outCoords[outI].col--
		outDirs[outI] = left
		outI++
	}
	if sendDown && c.row < size-1 {
		outCoords[outI] = c
		outCoords[outI].row++
		outDirs[outI] = down
		outI++
	}
	if sendUp && c.row > 0 {
		outCoords[outI] = c
		outCoords[outI].row--
		outDirs[outI] = up
		outI++
	}
	return outCoords, outDirs
}

type contraption struct {
	tiles [size][size]tile

	energized [size][size]lightTile
}

func newContraption(input string) contraption {
	ri, ci := 0, 0
	c := contraption{}
	for len(input) > 0 {
		switch input[0] {
		case '.':
			c.tiles[ri][ci] = empty
		case '|':
			c.tiles[ri][ci] = splitUp
		case '-':
			c.tiles[ri][ci] = splitWide
		case '/':
			c.tiles[ri][ci] = reflect1
		case '\\':
			c.tiles[ri][ci] = reflect2
		case '\n':
			ri++
			ci = -1
		}
		ci++
		input = input[1:]
	}

	return c
}

func (c *contraption) sendLight(
	co coord,
	lt lightTile,
) {
	if lt == 0 || (c.energized[co.row][co.col])&lt == lt {
		// already energed in this way
		return
	}
	c.energized[co.row][co.col] |= lt

	cs, lts := lt.update(co, c.tiles[co.row][co.col])
	c.sendLight(cs[0], lts[0])
	c.sendLight(cs[1], lts[1])
}

func (c *contraption) numEnergized() int {
	total := 0
	for ri := 0; ri < size; ri++ {
		for ci := 0; ci < size; ci++ {
			if c.energized[ri][ci] != 0 {
				total++
			}
		}
	}
	return total
}
