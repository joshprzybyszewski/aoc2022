package twentyone

import "fmt"

type gardenProvider struct {
	gar *garden

	// entrances struct {
	// 	// the entrance is in the top row
	// 	top [gridSize]precomputedGarden
	// 	// the entrance is in the bottom row
	// 	bottom [gridSize]precomputedGarden
	// 	// the entrance is in the left column
	// 	left [gridSize]precomputedGarden
	// 	// the entrance is in the right column
	// 	right [gridSize]precomputedGarden
	// }

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

func (g *gardenProvider) get(
	entrance coord,
) precomputedGarden {
	if entrance == g.gar.start {
		return g.start
	}
	fmt.Printf("entrance: %+v\n", entrance)
	panic(`unexpected entrance`)
}

func (g *gardenProvider) getRight(
	leftColumn [gridSize]int,
) precomputedGarden {
	// TODO cache this intelligently
	return newPrecomputedGardenWithLeftColumn(
		g.gar,
		leftColumn,
	)
}

func (g *gardenProvider) getLeft(
	rightColumn [gridSize]int,
) precomputedGarden {
	// TODO cache this intelligently
	return newPrecomputedGardenWithRightColumn(
		g.gar,
		rightColumn,
	)
}

func (g *gardenProvider) getUp(
	bottomRow [gridSize]int,
) precomputedGarden {
	// TODO cache this intelligently
	return newPrecomputedGardenWithBottomRow(
		g.gar,
		bottomRow,
	)
}

func (g *gardenProvider) getDown(
	topRow [gridSize]int,
) precomputedGarden {
	// TODO cache this intelligently
	return newPrecomputedGardenWithTopRow(
		g.gar,
		topRow,
	)
}
