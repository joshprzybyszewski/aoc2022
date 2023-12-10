package ten

func Two(
	input string,
) (int, error) {
	pm := createPipeMap(input)
	cpy := pm.onlyLoop()

	numInside := 0
	crossings := 0

	for r := 1; r < len(pm.tiles); r++ {
		crossings = 0
		for c := 0; c < len(pm.tiles[r]); c++ {
			if (cpy.tiles[r][c] & north) == north {
				crossings++
				continue
			}
			if cpy.tiles[r][c] == 0 && crossings%2 == 1 {
				numInside++
			}
		}
	}

	return numInside, nil
}

func (pm *pipeMap) onlyLoop() pipeMap {
	cpy := pipeMap{
		start: pm.start,
	}
	ends, headings := pm.getStarting()
	{ // fill in the start
		cpy.tiles[cpy.start.row][cpy.start.col] = 0
		if ends[0].col == pm.start.col+1 {
			cpy.tiles[cpy.start.row][cpy.start.col] |= east
		}
		if ends[0].row == pm.start.row+1 || ends[1].row == pm.start.row+1 {
			cpy.tiles[cpy.start.row][cpy.start.col] |= south
		}
		if ends[0].col == pm.start.col-1 || ends[1].col == pm.start.col-1 {
			cpy.tiles[cpy.start.row][cpy.start.col] |= west
		}
		if ends[1].row == pm.start.row-1 {
			cpy.tiles[cpy.start.row][cpy.start.col] |= north
		}
	}

	var i int
	var mask pipe

	for {
		cpy.tiles[ends[0].row][ends[0].col] = pm.tiles[ends[0].row][ends[0].col]
		cpy.tiles[ends[1].row][ends[1].col] = pm.tiles[ends[1].row][ends[1].col]

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
