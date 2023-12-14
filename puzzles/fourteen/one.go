package fourteen

import (
	"strings"
)

const (
	size = 100
)

type tile uint8

const (
	empty tile = 0
	rock  tile = 1
	block tile = 2
)

func (t tile) toByte() byte {
	switch t {
	case empty:
		return '.'
	case rock:
		return 'O'
	case block:
		return '#'
	}
	return '?'
}

type platform struct {
	tiles [size][size]tile
}

func newPlatform(input string) platform {
	ri, ci := 0, 0
	p := platform{}
	for len(input) > 0 {
		switch input[0] {
		case '.':
		case 'O':
			p.tiles[ri][ci] = rock
		case '#':
			p.tiles[ri][ci] = block
		case '\n':
			ri++
			ci = -1
		}
		ci++
		input = input[1:]
	}

	return p
}

func (p platform) String() string {
	var sb strings.Builder
	for ri := 0; ri < size; ri++ {
		for ci := 0; ci < size; ci++ {
			sb.WriteByte(p.tiles[ri][ci].toByte())
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (p *platform) rollNorth() {

	nextEmptySpotForCol := [size]int{}

	for ri := 0; ri < size; ri++ {
		for ci := 0; ci < size; ci++ {
			switch p.tiles[ri][ci] {
			case block:
				nextEmptySpotForCol[ci] = ri + 1
			case rock:
				if nextEmptySpotForCol[ci] != ri {
					p.tiles[nextEmptySpotForCol[ci]][ci] = rock
					p.tiles[ri][ci] = empty
					nextEmptySpotForCol[ci] += 1
				} else {
					nextEmptySpotForCol[ci] = ri + 1
				}
			}
		}
	}
}

func (p *platform) totalLoad() int {
	totalLoad := 0
	for ri := 0; ri < size; ri++ {
		for ci := 0; ci < size; ci++ {
			if p.tiles[ri][ci] == rock {
				totalLoad += size - ri
			}
		}
	}

	return totalLoad
}

func One(
	input string,
) (int, error) {
	p := newPlatform(input)
	p.rollNorth()

	return p.totalLoad(), nil
}
