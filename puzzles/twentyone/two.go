package twentyone

func Two(
	input string,
) (int, error) {
	initial := newGarden(input)

	gardenProvider := newGardenProvider(initial)

	galaxy := newGalaxyBuilder()

	galaxy.populate(&gardenProvider, 1000)
	// galaxy.populate(&gardenProvider, stepGoal)

	return galaxy.totalEven, nil
}

const (
	stepGoal = 26501365
)

type plotPosition struct {
	plot coord

	entrance    coord
	depthBefore int
}

type galaxyBuilder struct {
	seenPlots map[coord]struct{}

	pending []plotPosition

	maxDepth int

	totalEven int
}

func newGalaxyBuilder() galaxyBuilder {
	return galaxyBuilder{
		seenPlots: make(map[coord]struct{}, 4096),
		pending:   make([]plotPosition, 0, 4096), // gonna be way more than this
	}
}

func (gb *galaxyBuilder) addInitial(
	gp *gardenProvider,
) {
	if gb.totalEven != 0 {
		panic(`unexpected`)
	}

	gb.totalEven += gp.start.numEven

	gb.pending = append(gb.pending,
		plotPosition{
			plot: coord{
				row: 0,
				col: 1,
			},
			entrance:    gp.start.exits.right,
			depthBefore: gp.start.rightDepth(),
		},
		plotPosition{
			plot: coord{
				row: 0,
				col: -1,
			},
			entrance:    gp.start.exits.left,
			depthBefore: gp.start.leftDepth(),
		},
		plotPosition{
			plot: coord{
				row: 1,
				col: 0,
			},
			entrance:    gp.start.exits.bottom,
			depthBefore: gp.start.bottomDepth(),
		},
		plotPosition{
			plot: coord{
				row: -1,
				col: 0,
			},
			entrance:    gp.start.exits.top,
			depthBefore: gp.start.topDepth(),
		},
	)
}

func (gb *galaxyBuilder) populate(
	gp *gardenProvider,
	maxDepth int,
) {
	if gb.maxDepth != 0 {
		panic(`unexpected`)
	}
	gb.maxDepth = maxDepth

	gb.addInitial(gp)

	for len(gb.pending) > 0 {
		p := gb.pending[0]
		gb.pending = gb.pending[1:]

		if _, ok := gb.seenPlots[p.plot]; ok {
			continue
		}

		gb.seenPlots[p.plot] = struct{}{}

		// Look at the next one above
		// above := p
		// above.plot.row--
		// above.plot.entrance =
		// pending = append(pending,
		// 	pos{
		// 		plot: coord{
		// 			row: 0,
		// 			col: 1,
		// 		},
		// 		entrance:    gp.start.exits.right,
		// 		depthBefore: gp.start.rightDepth(),
		// 	},
		// )

		// below

		// right

		// left

	}
}

func (gb *galaxyBuilder) process(
	gp *gardenProvider,
	pos plotPosition,
) {
	plot := gp.get(pos.entrance)

	// TODO process this plot.
	// add numEven
	// look up, right, down, and left, adding those plots to the pending queue
}
