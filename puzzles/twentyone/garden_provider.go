package twentyone

import "fmt"

type gardenProvider struct {
	entrances struct {
		// the entrance is in the top row
		top [gridSize]precomputedGarden
		// the entrance is in the bottom row
		bottom [gridSize]precomputedGarden
		// the entrance is in the left column
		left [gridSize]precomputedGarden
		// the entrance is in the right column
		right [gridSize]precomputedGarden
	}

	start precomputedGarden
}

func newGardenProvider(
	initial garden,
) gardenProvider {
	g := gardenProvider{
		start: newPrecomputedInitialGarden(initial),
	}
	g.precompute(&initial)

	return g
}

func (g *gardenProvider) precompute(gar *garden) {
	for val := 0; val < gridSize; val++ {
		g.entrances.top[val] = newPrecomputedGarden(
			gar,
			coord{
				row: 0,
				col: val,
			},
		)
		g.entrances.bottom[val] = newPrecomputedGarden(
			gar,
			coord{
				row: gridSize - 1,
				col: val,
			},
		)
		g.entrances.left[val] = newPrecomputedGarden(
			gar,
			coord{
				row: val,
				col: 0,
			},
		)
		g.entrances.right[val] = newPrecomputedGarden(
			gar,
			coord{
				row: val,
				col: gridSize - 1,
			},
		)
	}
}

func (g *gardenProvider) get(entrance coord) precomputedGarden {
	if entrance.row == 0 {
		return g.entrances.top[entrance.col]
	}
	if entrance.col == 0 {
		return g.entrances.left[entrance.row]
	}
	if entrance.row == gridSize-1 {
		return g.entrances.bottom[entrance.col]
	}
	if entrance.col == gridSize-1 {
		return g.entrances.right[entrance.row]
	}
	if entrance == g.start.entrance {
		return g.start
	}
	fmt.Printf("entrance: %+v\n", entrance)
	panic(`unexpected entrance`)
}
