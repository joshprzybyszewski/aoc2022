package twentyone

import "fmt"

type gardenProvider struct {
	gar *garden

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
	return gardenProvider{
		start: newPrecomputedInitialGarden(initial),
		gar:   &initial,
	}
}

func (g *gardenProvider) precompute(
	entrance coord,
) {
	if entrance.row == 0 {
		if g.entrances.top[entrance.col].maxDistance == 0 {
			g.entrances.top[entrance.col] = newPrecomputedGarden(
				g.gar,
				entrance,
			)
		}

	} else if entrance.row == gridSize-1 {
		if g.entrances.bottom[entrance.col].maxDistance == 0 {
			g.entrances.bottom[entrance.col] = newPrecomputedGarden(
				g.gar,
				entrance,
			)
		}
	} else if entrance.col == 0 {
		if g.entrances.left[entrance.row].maxDistance == 0 {
			g.entrances.left[entrance.row] = newPrecomputedGarden(
				g.gar,
				entrance,
			)
		}
	} else if entrance.col == gridSize-1 {
		if g.entrances.right[entrance.row].maxDistance == 0 {
			g.entrances.right[entrance.row] = newPrecomputedGarden(
				g.gar,
				entrance,
			)
		}
	} else if entrance != g.start.entrance {
		fmt.Printf("entrance: %+v\n", entrance)
		panic(`unexpected`)
	}
}

func (g *gardenProvider) get(
	entrance coord,
	additional ...coord,
) precomputedGarden {
	if len(additional) > 0 {
		// do something else
		return newPrecomputedGardenWithManyEntrances(
			g.gar,
			entrance,
			additional,
		)
	}

	g.precompute(entrance)

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
