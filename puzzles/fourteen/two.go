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

type cycleDetector struct {
	seen [numLookBack]platform

	numSeen int
}

func newCycleDetector() cycleDetector {
	return cycleDetector{}
}

func (d *cycleDetector) see(p platform) int {
	d.seen[d.numSeen] = p
	d.numSeen++

	i := d.numSeen - 2
	i2 := d.numSeen - 3
	for i2 > 0 {
		if d.seen[i] == p && d.seen[i] == d.seen[i2] {
			return i - i2
		}
		i--
		i2 -= 2
	}

	return -1
}

func cycle(
	p *platform,
	n uint,
) {

	cd := newCycleDetector()

	var i uint
	var cycleLength int
	for i < n {
		p.rollNorth()
		p.rollWest()
		p.rollSouth()
		p.rollEast()
		i++

		cycleLength = cd.see(*p)
		if cycleLength > 0 {
			// skip ahead through all the remaining cycles
			i += ((n - i) / uint(cycleLength)) * uint(cycleLength)
			break
		}
	}

	for i < n {
		p.rollNorth()
		p.rollWest()
		p.rollSouth()
		p.rollEast()
		i++
	}

}

func (p *platform) rollSouth() {
	nextEmptySpotForCol := [size]int{}
	for i := range nextEmptySpotForCol {
		nextEmptySpotForCol[i] = size - 1
	}

	var ci int
	for ri := size - 1; ri >= 0; ri-- {
		for ci = 0; ci < size; ci++ {
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

	var ri int
	for ci := 0; ci < size; ci++ {
		for ri = 0; ri < size; ri++ {
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

	var ri int
	for ci := size - 1; ci >= 0; ci-- {
		for ri = 0; ri < size; ri++ {
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
