package fourteen

func Two(
	input string,
) (int, error) {
	p := newPlatform(input)

	p.cycle(3)

	return p.totalLoad(), nil
}

func (p *platform) cycle(
	n uint,
) {

	for i := uint(0); i < n; i++ {
		p.rollNorth()
		p.rollWest()
		p.rollSouth()
		p.rollEast()
	}
}

func (p *platform) rollSouth() {
	nextEmptySpotForCol := [size]int{}
	for i := range nextEmptySpotForCol {
		nextEmptySpotForCol[i] = size - 1
	}

	for ri := size - 1; ri >= 0; ri-- {
		for ci := 0; ci < size; ci++ {
			switch p.tiles[ri][ci] {
			case block:
				nextEmptySpotForCol[ci] = ri - 1
			case rock:
				if nextEmptySpotForCol[ci] != ri {
					p.tiles[nextEmptySpotForCol[ci]][ci] = rock
					p.tiles[ri][ci] = empty
					nextEmptySpotForCol[ci] -= 1
				} else {
					nextEmptySpotForCol[ci] = ri - 1
				}
			}
		}
	}
}

func (p *platform) rollWest() {
	nextEmptySpotForRow := [size]int{}

	for ci := 0; ci < size; ci++ {
		for ri := 0; ri < size; ri++ {
			switch p.tiles[ri][ci] {
			case block:
				nextEmptySpotForRow[ri] = ci + 1
			case rock:
				if nextEmptySpotForRow[ri] != ci {
					p.tiles[ri][nextEmptySpotForRow[ri]] = rock
					p.tiles[ri][ci] = empty
					nextEmptySpotForRow[ri] += 1
				} else {
					nextEmptySpotForRow[ri] = ci + 1
				}
			}
		}
	}

}

func (p *platform) rollEast() {
	nextEmptySpotForRow := [size]int{}
	for i := range nextEmptySpotForRow {
		nextEmptySpotForRow[i] = size - 1
	}

	for ci := size - 1; ci >= 0; ci-- {
		for ri := 0; ri < size; ri++ {
			switch p.tiles[ri][ci] {
			case block:
				nextEmptySpotForRow[ri] = ci - 1
			case rock:
				if nextEmptySpotForRow[ri] != ci {
					p.tiles[ri][nextEmptySpotForRow[ri]] = rock
					p.tiles[ri][ci] = empty
					nextEmptySpotForRow[ri] -= 1
				} else {
					nextEmptySpotForRow[ri] = ci - 1

				}
			}
		}
	}
}
