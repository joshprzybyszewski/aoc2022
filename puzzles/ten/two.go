package ten

import "fmt"

func Two(
	input string,
) (int, error) {
	pm := createPipeMap(input)
	cpy := pm.onlyLoop()

	cpy.markInsides()

	fmt.Printf("only loop:\n%s\n", cpy.String())
	// return pm.stepsToFarthest(), nil

	return 0, nil
}

func (pm *pipeMap) markInsides() {
	for r := 1; r < len(pm.tiles); r++ {
		for c := 1; c < len(pm.tiles[r]); c++ {
			if pm.isInside(coord{
				row: r,
				col: c,
			}) {
				pm.tiles[r][c] |= inside
			}
		}
	}
}

func (pm *pipeMap) isLoop(
	c coord,
) bool {
	return (pm.tiles[c.row][c.col] & allDirections) != 0
}

func (pm *pipeMap) isInside(
	c coord,
) bool {
	if pm.isLoop(c) {
		return false
	}
	if c.col > 0 && pm.tiles[c.col-1][c.col].isRightInside() {
		return true
	}
	if c.row > 0 && pm.tiles[c.row-1][c.col].isBelowInside() {
		return true
	}
	// if c.row+1 < mapSize && pm.tiles[c.row+1][c.col].isAboveInside() {
	// 	return true
	// }
	// if c.col+1 < mapSize && pm.tiles[c.col+1][c.col].isLeftInside() {
	// 	return true
	// }

	return false
}

func (pm *pipeMap) onlyLoop() pipeMap {
	cpy := pipeMap{
		start: pm.start,
	}
	ends, headings := pm.getStarting()

	insides := [2]pipe{}

	var i int
	var mask pipe

	for {
		cpy.tiles[ends[0].row][ends[0].col] = insides[0] | pm.tiles[ends[0].row][ends[0].col]
		cpy.tiles[ends[1].row][ends[1].col] = insides[1] | pm.tiles[ends[1].row][ends[1].col]

		if ends[0] == ends[1] {
			return cpy
		}

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
				break
			}
		}
	}
}
