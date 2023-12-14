package fourteen

import "fmt"

func Two(
	input string,
) (int, error) {
	p := newPlatform(input)

	cycle(p, 1000000000)

	// 104589 is too high
	// 105455 is too high

	return p.totalLoad(), nil
}

func cycle(
	p platform,
	n uint,
) {

	previouslySeen := [10000]platform{}
	psi := 0
	numSeen := 0
	lookBack := func(p platform) int {
		maxN := numSeen
		if maxN > len(previouslySeen) {
			maxN = len(previouslySeen)
		}
		i := psi
		for n := 1; n < maxN; n++ {
			if previouslySeen[i] == p {
				fmt.Printf("%s\n^ Saw %d iterations ago\n", previouslySeen[i], n)

				return n
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
		sinceLastSeen := lookBack(p)
		if sinceLastSeen > 0 {
			fmt.Printf("%s\n^ %d iterations\n", p, i)
			fmt.Printf("Saw this configuration %d times ago.\n", sinceLastSeen)
			for i+uint(sinceLastSeen) < n {
				i += uint(sinceLastSeen)
			}
			// remaining := n - i
			// numSkips := remaining / uint(sinceLastSeen)
			// fmt.Printf("Skipping ahead %d cycles.\n", numSkips)
			// i += (numSkips * uint(sinceLastSeen))
			break
		}

		if i%10000 == 0 {
			fmt.Printf("%s\n^ %d iterations\n", p, i)
		}
	}

	fmt.Printf("%s\n^ %d iterations (after all the skips)\n", p, i)

	for ; i < n; i++ {
		p.rollNorth()
		p.rollWest()
		p.rollSouth()
		p.rollEast()
	}
	fmt.Printf("%s\n^ %d iterations\n", p, i)
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
