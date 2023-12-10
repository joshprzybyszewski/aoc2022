package ten

import (
	"strings"
)

type coord struct {
	row int
	col int
}

const (
	north pipe = 1 << 0
	east  pipe = 1 << 1
	south pipe = 1 << 2
	west  pipe = 1 << 3

	ne pipe = north | east
	se      = south | east
	sw      = south | west
	nw      = north | west
	ew      = east | west
	ns      = north | south

	start pipe = north | east | south | west
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
		return ew
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

type pipeMap struct {
	rows [140][140]pipe

	start coord
}

func createPipeMap(
	input string,
) pipeMap {
	pm := pipeMap{}

	ri, ci := 0, 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for ci = 0; ci < nli; ci++ {
			pm.rows[ri][ci] = newPipe(input[ci])
			if pm.rows[ri][ci] == start {
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
	ends := [2]coord{}
}

func One(
	input string,
) (int, error) {

	return 0, nil
}
