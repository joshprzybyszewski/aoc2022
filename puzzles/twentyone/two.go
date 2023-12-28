package twentyone

import "fmt"

func Two(
	input string,
) (int, error) {
	for _, depth := range []int{
		6,
		10,
		50,
		100,
		// 500,
		// 1000,
		// 5000,
	} {
		answer := getAnswerFromGrid(input, depth)
		fmt.Printf("Got %10d in %4d steps\n", answer, depth)
	}

	return 0, nil
}

func getAnswerFromGrid(
	input string,
	depth int,
) int {
	initial := newGarden(input)

	gardenProvider := newGardenProvider(initial)

	galaxy := newGalaxyBuilder(&gardenProvider)

	galaxy.populate(depth)
	// galaxy.populate(&gardenProvider, stepGoal)

	return galaxy.totalEven
}

const (
	stepGoal = 26501365
)

type galaxyBuilder struct {
	gp *gardenProvider

	maxDepth int

	totalEven int
}

func newGalaxyBuilder(
	gp *gardenProvider,
) galaxyBuilder {
	return galaxyBuilder{
		gp: gp,
	}
}

func (gb *galaxyBuilder) populate(
	maxDepth int,
) {
	if gb.maxDepth != 0 {
		panic(`unexpected`)
	}
	gb.maxDepth = maxDepth

	// get the starting plot
	groundzero := gb.gp.get(gb.gp.gar.start)

	gb.totalEven += groundzero.getNumEven(
		gb.maxDepth, /*- pos.depthBefore*/
	)

	// extend up and down from the start
	gb.extendUp(
		gb.getRowAbove(&groundzero),
	)
	gb.extendDown(
		gb.getRowBelow(&groundzero),
	)

	// extend right, which extends up and down along the way
	gb.extendRight(
		gb.getColumnToTheRight(&groundzero),
	)

	// extend left, which extends up and down along the way
	gb.extendLeft(
		gb.getColumnToTheLeft(&groundzero),
	)
}

func (gb *galaxyBuilder) getColumnToTheRight(
	pg *precomputedGarden,
) [gridSize]int {
	var output [gridSize]int
	for row := 0; row < gridSize; row++ {
		output[row] = pg.distances[row][gridSize-1] + 1
	}
	return output
}

func (gb *galaxyBuilder) getColumnToTheLeft(
	pg *precomputedGarden,
) [gridSize]int {
	var output [gridSize]int
	for row := 0; row < gridSize; row++ {
		output[row] = pg.distances[row][0] + 1
	}
	return output
}

func (gb *galaxyBuilder) getRowAbove(
	pg *precomputedGarden,
) [gridSize]int {
	var output [gridSize]int
	for col := 0; col < gridSize; col++ {
		output[col] = pg.distances[0][col] + 1
	}
	return output
}

func (gb *galaxyBuilder) getRowBelow(
	pg *precomputedGarden,
) [gridSize]int {
	var output [gridSize]int
	for col := 0; col < gridSize; col++ {
		output[col] = pg.distances[gridSize-1][col] + 1
	}
	return output
}

func (gb *galaxyBuilder) extendRight(
	leftColumn [gridSize]int,
) {
	myplot := gb.gp.getRight(leftColumn)
	num := myplot.getNumEven(
		gb.maxDepth,
	)
	if num == 0 {
		return
	}

	gb.totalEven += num

	gb.extendUp(
		gb.getRowAbove(&myplot),
	)
	gb.extendDown(
		gb.getRowBelow(&myplot),
	)

	gb.extendRight(
		gb.getColumnToTheRight(&myplot),
	)

}

func (gb *galaxyBuilder) extendLeft(
	rightColumn [gridSize]int,
) {
	myplot := gb.gp.getLeft(rightColumn)
	num := myplot.getNumEven(
		gb.maxDepth,
	)
	if num == 0 {
		return
	}

	gb.totalEven += num

	gb.extendUp(
		gb.getRowAbove(&myplot),
	)
	gb.extendDown(
		gb.getRowBelow(&myplot),
	)

	gb.extendLeft(
		gb.getColumnToTheLeft(&myplot),
	)
}

func (gb *galaxyBuilder) extendUp(
	bottomRow [gridSize]int,
) {
	myplot := gb.gp.getUp(bottomRow)
	num := myplot.getNumEven(
		gb.maxDepth,
	)
	if num == 0 {
		return
	}

	gb.totalEven += num

	gb.extendUp(
		gb.getRowAbove(&myplot),
	)

}

func (gb *galaxyBuilder) extendDown(
	topRow [gridSize]int,
) {
	myplot := gb.gp.getDown(topRow)
	num := myplot.getNumEven(
		gb.maxDepth,
	)
	if num == 0 {
		return
	}

	gb.totalEven += num

	gb.extendDown(
		gb.getRowBelow(&myplot),
	)

}
