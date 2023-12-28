package twentyone

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (int, error) {
	for _, depth := range []int{
		6,
		10,
		50,
		100,
		500,
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
	if len(galaxy.gardensByCoord) < 100 {
		fmt.Printf("Galaxy:\n%s\n", galaxy.String())
	}

	return galaxy.totalEven
}

const (
	stepGoal = 26501365
)

type galaxyBuilder struct {
	gp *gardenProvider

	maxDepth int

	totalEven int

	gardensByCoord map[coord]*precomputedGarden
}

func newGalaxyBuilder(
	gp *gardenProvider,
) galaxyBuilder {
	return galaxyBuilder{
		gp:             gp,
		gardensByCoord: make(map[coord]*precomputedGarden, 100),
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
	gb.gardensByCoord[coord{0, 0}] = &groundzero

	gb.totalEven += groundzero.getNumEven(
		gb.maxDepth, /*- pos.depthBefore*/
	)

	// extend up and down from the start
	gb.extendUp(
		coord{row: -1, col: 0},
	)
	gb.extendDown(
		coord{row: 1, col: 0},
	)

	// extend right, which extends up and down along the way
	gb.extendRight(
		coord{row: 0, col: 1},
	)

	// extend left, which extends up and down along the way
	gb.extendLeft(
		coord{row: 0, col: -1},
	)

	// TODO go through the quartiles

}

func (gb *galaxyBuilder) extendRight(
	myCoord coord,
) {

	myplot := gb.gp.getWithNeighbors(neighbors{
		left: gb.gardensByCoord[myCoord.left()],
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot
	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendUpAndLeft(
		myCoord.up(),
	)
	gb.extendDownAndLeft(
		myCoord.down(),
	)

	gb.extendRight(
		myCoord.right(),
	)

}

func (gb *galaxyBuilder) extendLeft(
	myCoord coord,
) {

	myplot := gb.gp.getWithNeighbors(neighbors{
		right: gb.gardensByCoord[myCoord.right()],
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot
	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUpAndRight(
		myCoord.up(),
	)
	gb.extendDownAndRight(
		myCoord.down(),
	)

	gb.extendLeft(
		myCoord.left(),
	)
}

func (gb *galaxyBuilder) extendUp(
	myCoord coord,
) {

	myplot := gb.gp.getWithNeighbors(neighbors{
		below: gb.gardensByCoord[myCoord.down()],
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	num := myplot.getNumEven(
		gb.maxDepth,
	)

	gb.totalEven += num

	gb.extendUp(
		myCoord.up(),
	)

}

func (gb *galaxyBuilder) extendDown(
	myCoord coord,
) {
	myplot := gb.gp.getWithNeighbors(neighbors{
		above: gb.gardensByCoord[myCoord.up()],
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendDown(
		myCoord.down(),
	)

}

func (gb *galaxyBuilder) extendUpAndLeft(
	myCoord coord,
) {
	_, ok := gb.gardensByCoord[myCoord]
	if ok {
		return
	}

	left, ok := gb.gardensByCoord[myCoord.left()]
	if !ok {
		return
	}

	below, ok := gb.gardensByCoord[myCoord.down()]
	if !ok {
		return
	}

	myplot := gb.gp.getWithNeighbors(neighbors{
		left:  left,
		below: below,
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendUpAndLeft(
		myCoord.up().left(),
	)
}

func (gb *galaxyBuilder) extendDownAndLeft(
	myCoord coord,
) {
	_, ok := gb.gardensByCoord[myCoord]
	if ok {
		return
	}

	left, ok := gb.gardensByCoord[myCoord.left()]
	if !ok {
		return
	}

	above, ok := gb.gardensByCoord[myCoord.up()]
	if !ok {
		return
	}

	myplot := gb.gp.getWithNeighbors(neighbors{
		left:  left,
		above: above,
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendDownAndLeft(
		myCoord.down().left(),
	)
}

func (gb *galaxyBuilder) extendUpAndRight(
	myCoord coord,
) {
	_, ok := gb.gardensByCoord[myCoord]
	if ok {
		return
	}

	right, ok := gb.gardensByCoord[myCoord.right()]
	if !ok {
		return
	}

	below, ok := gb.gardensByCoord[myCoord.down()]
	if !ok {
		return
	}

	myplot := gb.gp.getWithNeighbors(neighbors{
		right: right,
		below: below,
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendUpAndRight(
		myCoord.up().right(),
	)
}

func (gb *galaxyBuilder) extendDownAndRight(
	myCoord coord,
) {
	_, ok := gb.gardensByCoord[myCoord]
	if ok {
		return
	}

	right, ok := gb.gardensByCoord[myCoord.right()]
	if !ok {
		return
	}

	above, ok := gb.gardensByCoord[myCoord.up()]
	if !ok {
		return
	}

	myplot := gb.gp.getWithNeighbors(neighbors{
		right: right,
		above: above,
	})
	if myplot.minDistance > gb.maxDepth {
		return
	}
	gb.gardensByCoord[myCoord] = &myplot

	gb.totalEven += myplot.getNumEven(
		gb.maxDepth,
	)

	gb.extendDownAndRight(
		myCoord.down().right(),
	)
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
