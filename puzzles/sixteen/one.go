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

type contraption struct {
	tiles [size][size]tile

	energized [size][size]lightTile
}

func newContraption(input string) contraption {
	ri, ci := 0, 0
	c := contraption{}
	for len(input) > 0 {
		switch input[0] {
		case '.': // 10927
			c.tiles[ri][ci] = empty
		case '\\': // 306
			c.tiles[ri][ci] = reflect2
		case '/': // 301
			c.tiles[ri][ci] = reflect1
		case '-': // 289
			c.tiles[ri][ci] = splitWide
		case '|': // 277
			c.tiles[ri][ci] = splitUp
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
	if (c.energized[co.row][co.col])&lt == lt {
		// already energed in this way
		return
	}
	c.energized[co.row][co.col] |= lt

	sendRight := false
	sendLeft := false
	sendDown := false
	sendUp := false

	switch c.tiles[co.row][co.col] {
	case empty:
		sendRight = lt == right
		sendLeft = lt == left
		sendDown = lt == down
		sendUp = lt == up
	case reflect2:
		sendRight = lt == down
		sendLeft = lt == up
		sendDown = lt == right
		sendUp = lt == left
	case reflect1:
		sendRight = lt == up
		sendLeft = lt == down
		sendDown = lt == left
		sendUp = lt == right
	case splitWide:
		sendLeft = lt&(up|down|left) == lt
		sendRight = lt&(up|down|right) == lt
	case splitUp:
		sendDown = lt&(down|left|right) == lt
		sendUp = lt&(up|left|right) == lt
	}

	if sendRight && co.col < size-1 {
		tmp := co
		tmp.col++
		c.sendLight(tmp, right)
	}
	if sendLeft && co.col > 0 {
		tmp := co
		tmp.col--
		c.sendLight(tmp, left)
	}
	if sendDown && co.row < size-1 {
		tmp := co
		tmp.row++
		c.sendLight(tmp, down)
	}
	if sendUp && co.row > 0 {
		tmp := co
		tmp.row--
		c.sendLight(tmp, up)
	}
}

func (c *contraption) numEnergized() int {
	total := 0
	var ci int
	for ri := 0; ri < size; ri++ {
		for ci = 0; ci < size; ci++ {
			if c.energized[ri][ci] != 0 {
				total++
			}
		}
	}
	return total
}
