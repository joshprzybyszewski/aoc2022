package twentyone

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (int, error) {
	for _, depth := range []int{
		// 6,
		// 10,
		50,
		// 100,
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
	fmt.Printf("Galaxy:\n%s\n", galaxy.String())
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

	gardensByCoord map[coord]precomputedGarden
}

func newGalaxyBuilder(
	gp *gardenProvider,
) galaxyBuilder {
	return galaxyBuilder{
		gp:             gp,
		gardensByCoord: make(map[coord]precomputedGarden, 100),
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
	gb.gardensByCoord[coord{0, 0}] = groundzero

	gb.totalEven += groundzero.getNumEven(
		gb.maxDepth, /*- pos.depthBefore*/
	)

	// extend up and down from the start
	gb.extendUp(
		coord{row: -1, col: 0},
		gb.getRowAbove(&groundzero),
	)
	gb.extendDown(
		coord{row: 1, col: 0},
		gb.getRowBelow(&groundzero),
	)

	// extend right, which extends up and down along the way
	gb.extendRight(
		coord{row: 0, col: 1},
		gb.getColumnToTheRight(&groundzero),
	)

	// extend left, which extends up and down along the way
	gb.extendLeft(
		coord{row: 0, col: -1},
		gb.getColumnToTheLeft(&groundzero),
	)

	// TODO go through the quartiles

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
	myCoord coord,
	leftColumn [gridSize]int,
) {
	if gb.isBeyond(leftColumn) {
		return
	}

	myplot := gb.gp.getRight(leftColumn)
	gb.gardensByCoord[myCoord] = myplot
	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUpAndLeft(
		myCoord.up(),
	)
	gb.extendDownAndLeft(
		myCoord.down(),
	)

	gb.extendRight(
		myCoord.right(),
		gb.getColumnToTheRight(&myplot),
	)

}

func (gb *galaxyBuilder) extendLeft(
	myCoord coord,
	rightColumn [gridSize]int,
) {
	if gb.isBeyond(rightColumn) {
		return
	}

	myplot := gb.gp.getLeft(rightColumn)
	gb.gardensByCoord[myCoord] = myplot
	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUp(
		myCoord.up(),
		gb.getRowAbove(&myplot),
	)
	gb.extendDown(
		myCoord.down(),
		gb.getRowBelow(&myplot),
	)

	gb.extendLeft(
		myCoord.left(),
		gb.getColumnToTheLeft(&myplot),
	)
}

func (gb *galaxyBuilder) extendUp(
	myCoord coord,
	bottomRow [gridSize]int,
) {
	if gb.isBeyond(bottomRow) {
		return
	}

	myplot := gb.gp.getUp(bottomRow)
	gb.gardensByCoord[myCoord] = myplot

	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUp(
		myCoord.up(),
		gb.getRowAbove(&myplot),
	)

}

func (gb *galaxyBuilder) extendDown(
	myCoord coord,
	topRow [gridSize]int,
) {
	if gb.isBeyond(topRow) {
		return
	}

	myplot := gb.gp.getDown(topRow)
	gb.gardensByCoord[myCoord] = myplot
	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendDown(
		myCoord.down(),
		gb.getRowBelow(&myplot),
	)

}

func (gb *galaxyBuilder) extendUpAndLeft(
	myCoord coord,
) {
	left, ok := gb.gardensByCoord[myCoord.left()]
	if !ok {
		return
	}
	leftColumn := gb.getColumnToTheRight(&left)

	below, ok := gb.gardensByCoord[myCoord.down()]
	if !ok {
		return
	}
	bottomRow := gb.getRowAbove(&below)

	if gb.isBeyond(bottomRow) && gb.isBeyond(leftColumn) {
		return
	}

	myplot := gb.gp.getWithBottomRowAndLeftColumn(
		bottomRow,
		leftColumn,
	)
	gb.gardensByCoord[myCoord] = myplot

	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUpAndLeft(
		myCoord.up().left(),
	)
}

func (gb *galaxyBuilder) extendDownAndLeft(
	myCoord coord,
) {
	left, ok := gb.gardensByCoord[myCoord.left()]
	if !ok {
		return
	}
	leftColumn := gb.getColumnToTheRight(&left)

	above, ok := gb.gardensByCoord[myCoord.up()]
	if !ok {
		return
	}
	topRow := gb.getRowBelow(&above)

	if gb.isBeyond(topRow) && gb.isBeyond(leftColumn) {
		return
	}

	myplot := gb.gp.getWithTopRowAndLeftColumn(topRow, leftColumn)
	gb.gardensByCoord[myCoord] = myplot

	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendDownAndLeft(
		myCoord.down().left(),
	)
}

func (gb *galaxyBuilder) isBeyond(
	startingValues [gridSize]int,
) bool {
	for _, v := range startingValues {
		if v < gb.maxDepth {
			return false
		}
	}
	return true
}

func (gb *galaxyBuilder) String() string {
	var min, max coord
	for k := range gb.gardensByCoord {
		if k.row < min.row {
			min.row = k.row
		}
		if k.col < min.col {
			min.col = k.col
		}
		if k.row > max.row {
			max.row = k.row
		}
		if k.col > max.col {
			max.col = k.col
		}
	}

	var sb strings.Builder
	for ri := min.row; ri <= max.row; ri++ {
		gardens := make([][gridSize]string, 0, max.col-min.col+1)
		for ci := min.col; ci <= max.col; ci++ {
			pg, ok := gb.gardensByCoord[coord{row: ri, col: ci}]
			if ok {
				gardens = append(gardens, pg.lines(gb.maxDepth))
			} else {
				gardens = append(gardens, getEmptyLines())
			}
		}
		for row := 0; row < gridSize; row++ {
			for i := range gardens {
				sb.WriteString(gardens[i][row])
			}
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func getEmptyLines() [gridSize]string {
	var output [gridSize]string
	var sb strings.Builder
	for i := 0; i < gridSize; i++ {
		sb.WriteByte(' ')
	}
	for i := range output {
		output[i] = sb.String()
	}

	return output
}
