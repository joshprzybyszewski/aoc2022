package ten

import (
	"fmt"
	"strings"
)

const (
	mapSize = 140
)

type coord struct {
	row int
	col int
}

const (
	east  pipe = 1 << 0
	north pipe = 1 << 1
	west  pipe = 1 << 2
	south pipe = 1 << 3

	start pipe = 1 << 4 /* will also be any two: north | east | south | west */

	ne pipe = north | east
	se      = south | east
	sw      = south | west
	nw      = north | west
	ew      = east | west
	ns      = north | south
)

type pipe uint8

func newPipe(b byte) pipe {
	switch b {
	case '-': // 2233
		return ew
	case '7': // 2027
		return sw
	case '|': // 2003
		return ns
	case 'F': // 1967
		return se
	case 'L': // 1946
		return ne
	case 'J': // 1917
		return nw
	case '.':
		return 0
	case 'S':
		return start
	}
	panic(`surprise`)
}

func (p pipe) String() string {
	return fmt.Sprintf("%08b", p)
}

func (p pipe) stringForMap() byte {
	if (p & start) == start {
		return 'S'
	}
	switch p {
	case ne:
		return 'L'
	case nw:
		return 'J'
	case sw:
		return '7'
	case se:
		return 'F'
	case ew:
		return '-'
	case ns:
		return '|'
	}

	return '.'
}

type pipeMap struct {
	tiles [mapSize][mapSize]pipe

	start coord
}

func (pm *pipeMap) String() string {
	var output strings.Builder
	for r := 0; r < len(pm.tiles); r++ {
		for c := 0; c < len(pm.tiles[r]); c++ {
			output.WriteByte(pm.tiles[r][c].stringForMap())
		}
		output.WriteByte('\n')
	}
	return output.String()
}

func createPipeMap(
	input string,
) pipeMap {
	pm := pipeMap{}

	ri, ci := 0, 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for ci = 0; ci < nli; ci++ {
			pm.tiles[ri][ci] = newPipe(input[ci])
			if pm.tiles[ri][ci] == start {
				pm.start = coord{
					row: ri,
					col: ci,
				}
			}
		}

		ri++
		input = input[nli+1:]
	}

	return pm
}

func (pm *pipeMap) stepsToFarthest() int {
	ends, headings := pm.getStarting()

	var i int
	var mask pipe
	numSteps := 1

	for {
		if ends[0] == ends[1] {
			return numSteps
		}

		numSteps++

		for i = 0; i < len(ends); i++ {
			switch headings[i] {
			case east:
				ends[i].col++
				mask = ^west
			case north:
				ends[i].row--
				mask = ^south
			case west:
				ends[i].col--
				mask = ^east
			case south:
				ends[i].row++
				mask = ^north
			default:
				panic(`dev error`)
			}
			headings[i] = pm.tiles[ends[i].row][ends[i].col] & mask

			if ends[0] == ends[1] {
				return numSteps
			}
		}
	}
}

func (pm *pipeMap) getStarting() ([2]coord, [2]pipe) {
	ends := [2]coord{
		pm.start,
		pm.start,
	}
	headings := [2]pipe{}

	i := 0
	if /*pm.start.col+1 < mapSize &&*/ pm.tiles[pm.start.row][pm.start.col+1]&west == west {
		ends[i].col++
		headings[i] = pm.tiles[ends[i].row][ends[i].col] & (^west)
		i++
	}

	if /*pm.start.row+1 < mapSize &&*/ pm.tiles[pm.start.row+1][pm.start.col]&north == north {
		ends[i].row++
		headings[i] = pm.tiles[ends[i].row][ends[i].col] & (^north)
		i++
	}

	if /*pm.start.col > 0 &&*/ pm.tiles[pm.start.row][pm.start.col-1]&east == east {
		ends[i].col--
		headings[i] = pm.tiles[ends[i].row][ends[i].col] & (^east)
		i++
	}

	if /*pm.start.row > 0 &&*/ pm.tiles[pm.start.row-1][pm.start.col]&south == south {
		ends[i].row--
		headings[i] = pm.tiles[ends[i].row][ends[i].col] & (^south)
		i++
	}

	return ends, headings
}

func One(
	input string,
) (int, error) {

	pm := createPipeMap(input)

	return pm.stepsToFarthest(), nil
}
