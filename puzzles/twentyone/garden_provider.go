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

type neighbors struct {
	above *precomputedGarden
	below *precomputedGarden
	left  *precomputedGarden
	right *precomputedGarden
}

func (g *gardenProvider) getWithNeighbors(
	n neighbors,
) precomputedGarden {
	starts := make(map[coord]int, gridSize*2)
	for val := 0; val < gridSize; val++ {
		if n.above != nil {
			starts[coord{
				row: 0,
				col: val,
			}] = n.above.distances[gridSize-1][val] + 1
		}
		if n.below != nil {
			starts[coord{
				row: gridSize - 1,
				col: val,
			}] = n.below.distances[0][val] + 1
		}
		if n.left != nil {
			starts[coord{
				row: val,
				col: 0,
			}] = n.left.distances[val][gridSize-1] + 1
		}
		if n.right != nil {
			starts[coord{
				row: val,
				col: gridSize - 1,
			}] = n.right.distances[val][0] + 1
		}
	}

	if len(starts) == 0 {
		panic(`unexpected`)
	}

	if n.above != nil {
		if n.left != nil {
			if n.above.distances[gridSize-1][0] != n.left.distances[0][gridSize-1] {
				fmt.Printf("Interesting: %d, %d\n", n.above.distances[gridSize-1][0], n.left.distances[gridSize-1][0])
				panic(`unexpected`)
			}
		}
		if n.right != nil {
			if n.above.distances[gridSize-1][gridSize-1] != n.right.distances[0][0] {
				fmt.Printf("Interesting: %d, %d\n", n.above.distances[gridSize-1][gridSize-1], n.right.distances[0][0])
				panic(`unexpected`)
			}
		}
	}
	if n.below != nil {
		if n.left != nil {
			if n.below.distances[0][0] != n.left.distances[gridSize-1][gridSize-1] {
				fmt.Printf("Interesting: %d, %d\n", n.below.distances[0][0], n.left.distances[gridSize-1][0])
				panic(`unexpected`)
			}
		}
		if n.right != nil {
			if n.below.distances[0][gridSize-1] != n.right.distances[gridSize-1][0] {
				fmt.Printf("Interesting: %d, %d\n", n.below.distances[0][0], n.right.distances[0][0])
				panic(`unexpected`)
			}
		}
	}

	return newPrecomputedGardenWithStarts(
		g.gar,
		starts,
	)
}
