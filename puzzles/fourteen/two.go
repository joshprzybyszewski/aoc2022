package fourteen

const (
	numCycles   = 1000000000
	numLookBack = numCycles / 1000000
)

func Two(
	input string,
) (int, error) {
	p := newPlatform(input)

	cycle(&p, numCycles)

	return p.totalLoad(), nil
}

func cycle(
	p *platform,
	n uint,
) {

	previouslySeen := [numLookBack]platform{}
	psi := 0
	numSeen := 0
	lookBack := func(p platform) int {
		maxN := numSeen
		if maxN > len(previouslySeen) {
			maxN = len(previouslySeen)
		}
		maxN /= 2
		i := psi
		for n := 1; n < maxN; n++ {
			if previouslySeen[i] == p {
				i2 := i + 1 - n
				if i2 < 0 {
					panic(`unsupported`)
				}
				if previouslySeen[i2] == p {
					return i - i2
				}
			}

			i--
			if i < 0 {
				i += len(previouslySeen)
			}
		}

		previouslySeen[psi] = p
		psi++
		numSeen++
		if numSeen >= len(previouslySeen) {
			panic(`unhandled`)
		}
		return -1
	}

	var i uint
	for ; i < n; i++ {
		p.rollNorth()
		p.rollWest()
		p.rollSouth()
		p.rollEast()
		cycleLength := lookBack(*p)
		if cycleLength > 0 {
			i++
			for i+uint(cycleLength) < n {
				i += uint(cycleLength)
			}
			break
		}
	}

	for ; i < n; i++ {
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
