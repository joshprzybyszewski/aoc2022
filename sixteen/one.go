package sixteen

import "fmt"

const (
	// This is assumed in the puzzle.
	part1StartingNode = `AA`
)

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
	g := buildGraph(part1StartingNode, valves)

	// 1963, 1986 is too high
	return getBestPath(
		valves,
		g,
		30,
	), nil
}

func getBestPath(
	valves []*valve,
	g graph,
	remaining time,
) int {

	var names [numNodes]string
	ni := 0
	for _, v := range valves {
		if v.flow > 0 {
			names[ni] = v.name
			ni++
		}
	}

	var best, em pathState

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

		if em.released > best.released {
			best = em
		}
	}

	// var info string
	// for ps := &best; ps != nil; ps = ps.prev {
	// 	info = "=============\n" + ps.String(names) + "\n" + info
	// }
	// fmt.Println(info)

	return int(best.released)
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

	prev *pathState
}

func (s pathState) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s pathState) String(
	names [numNodes]string,
) string {
	openValves := `Open Valves: `
	no := 0
	for i := 0; i < numNodes; i++ {
		if !s.valves.isOpen(node(i)) {
			continue
		}
		if no > 0 {
			openValves += `, `
		}
		openValves += names[i]
		no++
	}
	if no == 0 {
		openValves += `(none)`
	}
	return fmt.Sprintf(
		"Current Node: %s\n\t%s\n\tMinute: %d\n\tTime Remaining: %d\n\tPressure released: %d\n",
		names[s.cur],
		openValves,
		30-s.remaining,
		s.remaining,
		s.released,
	)
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

	s2 := s
	s2.prev = &s
	s2.remaining -= time(g.getDistance(s2.cur, dest))
	s2.cur = dest

	return s2, true
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

	s2 := s
	s2.prev = &s
	s2.valves = s2.valves.open(s2.cur)
	s2.remaining -= 1
	s2.released += (pressure(s2.remaining) * pressure(g.getValue(s2.cur)))

	return s2, true
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
