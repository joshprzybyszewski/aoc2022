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

	gb.process(
		gp,
		plotPosition{
			plot: coord{
				row: 0,
				col: 0,
			},
			entrance:    gp.start.entrance,
			depthBefore: 0,
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

		gb.process(gp, p)
	}
}

func (gb *galaxyBuilder) process(
	gp *gardenProvider,
	pos plotPosition,
) {
	plot := gp.get(pos.entrance)

	gb.totalEven += plot.getNumEven(
		gb.maxDepth - pos.depthBefore,
	)

	above := pos
	above.plot.row--
	above.entrance = plot.exits.top
	above.depthBefore += plot.topDepth()
	if above.depthBefore < gb.maxDepth {
		gb.pending = append(gb.pending, above)
	}

	below := pos
	below.plot.row++
	below.entrance = plot.exits.bottom
	below.depthBefore += plot.bottomDepth()
	if below.depthBefore < gb.maxDepth {
		gb.pending = append(gb.pending, below)
	}

	left := pos
	left.plot.row--
	left.entrance = plot.exits.left
	left.depthBefore += plot.leftDepth()
	if left.depthBefore < gb.maxDepth {
		gb.pending = append(gb.pending, left)
	}

	right := pos
	right.plot.row++
	right.entrance = plot.exits.right
	right.depthBefore += plot.rightDepth()
	if right.depthBefore < gb.maxDepth {
		gb.pending = append(gb.pending, right)
	}
}
