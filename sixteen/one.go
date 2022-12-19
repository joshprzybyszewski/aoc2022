package sixteen

// TODO can be replaced with int8
type time int

// TODO can be a uint8, or maybe 16 or 32
type pressure int

func One(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(valves)

	// 1986 too high
	return getBestPath(
		g,
		30,
	), nil
}

func getBestPath(
	g graph,
	remaining time,
) int {

	max := 0
	var em pathState

	for n, d := range g.startingPositions {
		// fmt.Printf("Starting from %d (spent %d time)\n", n, d)
		em = maximize(
			g,
			pathState{
				cur:       node(n),
				remaining: remaining - time(d),
			},
		)
		// fmt.Printf("best path from %d is: %+v\n\n", n, em)

		if int(em.released) > max {
			max = int(em.released)
		}
	}

	return max
}

// valveState is an array of bools. When true, the valve has been opened
// TODO could replace with a bitmap
type valveState [numNodes]bool

func (vs valveState) isOpen(n node) bool {
	return ([numNodes]bool)(vs)[n]
}

func (vs valveState) open(n node) valveState {
	cpy := ([numNodes]bool)(vs)
	cpy[n] = true
	return cpy
}

type pathState struct {
	cur node

	valves valveState

	remaining time
	released  pressure
}

func (s pathState) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func travel(
	g graph,
	s pathState,
	dest node,
) (pathState, bool) {
	if s.remaining <= 0 {
		// fmt.Printf("\tCannot travel. No time remaining\n")
		// no time
		return s, false
	}

	if s.isOpen(dest) {
		// fmt.Printf("\tCannot travel. already open\n")
		// the destination is already open. Not worth going to.
		return s, false
	}

	// fmt.Printf("\tTraveling from %2d to %2d (distance %d)\n",
	// 	s.cur,
	// 	dest,
	// 	g.getDistance(s.cur, dest),
	// )

	s.remaining -= time(g.getDistance(s.cur, dest))
	s.cur = dest

	return s, true
}

func open(
	g graph,
	s pathState,
) (pathState, bool) {
	if s.remaining <= 0 {
		// fmt.Printf("\tCannot open. No time remaining\n")
		// no time
		return s, false
	}
	if s.isOpen(s.cur) {
		// fmt.Printf("\tCannot open. already open\n")
		// the current node is already open. Cannot open
		return s, false
	}
	// fmt.Printf("\tOpening %2d at %2d (with value %d)\n",
	// 	s.cur,
	// 	s.remaining,
	// 	g.getValue(s.cur),
	// )

	s.valves = s.valves.open(s.cur)
	s.remaining -= 1
	s.released += (pressure(s.remaining) * pressure(g.getValue(s.cur)))

	return s, true
}

func maximize(
	g graph,
	s pathState,
) pathState {
	s, ok := open(g, s)
	if !ok {
		return s
	}

	best := s
	var ts pathState
	for i := 0; i < numNodes; i++ {
		ts, ok = travel(
			g,
			s,
			node(i),
		)
		if !ok {
			continue
		}
		tmax := maximize(g, ts)
		if tmax.released > best.released {
			best = tmax
		}
	}

	return best
}
