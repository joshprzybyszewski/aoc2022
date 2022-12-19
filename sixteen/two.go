package sixteen

import "fmt"

func Two(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(part1StartingNode, valves)

	// 1652 is too low
	return getBestPathForTwoTravellers(
		valves,
		g,
		26,
	), nil
}

func getBestPathForTwoTravellers(
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

	var best, em pathState2

	for n1, d1 := range g.startingPositions {
		for n2, d2 := range g.startingPositions {
			if n1 == n2 {
				// no use in going to the same node
				continue
			}
			// fmt.Printf("Starting from %d (spent %d time)\n", n, d)
			em = maximize2(
				g,
				pathState2{
					cur1:       node(n1),
					cur2:       node(n2),
					remaining1: remaining - time(d1),
					remaining2: remaining - time(d2),
				},
			)
			// fmt.Printf("best path from %d is: %+v\n\n", n, em)

			if em.released > best.released {
				best = em
			}
		}
	}

	var info string
	for ps := &best; ps != nil; ps = ps.prev {
		info = "=============\n" + ps.String(names) + "\n" + info
	}
	fmt.Println(info)

	return int(best.released)
}

type pathState2 struct {
	cur1, cur2             node
	remaining1, remaining2 time

	valves valveState

	released pressure

	prev *pathState2
}

func (s pathState2) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s pathState2) String(
	names [numNodes]string,
) string {
	return `TODO`
	// 	openValves := `Open Valves: `
	// 	no := 0
	// 	for i := 0; i < numNodes; i++ {
	// 		if !s.valves.isOpen(node(i)) {
	// 			continue
	// 		}
	// 		if no > 0 {
	// 			openValves += `, `
	// 		}
	// 		openValves += names[i]
	// 		no++
	// 	}
	// 	if no == 0 {
	// 		openValves += `(none)`
	// 	}
	// 	return fmt.Sprintf(
	// 		"Current Node: %s\n\t%s\n\tMinute: %d\n\tTime Remaining: %d\n\tPressure released: %d\n",
	// 		names[s.cur],
	// 		openValves,
	// 		30-s.remaining,
	// 		s.remaining,
	// 		s.released,
	// 	)
}

func travel2(
	g graph,
	s pathState2,
	dest node,
) (pathState2, bool) {
	if s.remaining1 <= 0 && s.remaining2 <= 0 {
		// fmt.Printf("\tCannot travel. No time remaining\n")
		// no time
		return s, false
	}

	if s.isOpen(dest) {
		// fmt.Printf("\tCannot travel. already open\n")
		// the destination is already open. Not worth going to.
		return s, false
	}

	if s.remaining1 >= s.remaining2 {
		s2 := s
		s2.prev = &s
		s2.remaining1 -= time(g.getDistance(s2.cur1, dest))
		s2.cur1 = dest

		return s2, true
	}

	// fmt.Printf("\tTraveling from %2d to %2d (distance %d)\n",
	// 	s.cur,
	// 	dest,
	// 	g.getDistance(s.cur, dest),
	// )

	s2 := s
	s2.prev = &s
	s2.remaining2 -= time(g.getDistance(s2.cur2, dest))
	s2.cur2 = dest

	return s2, true
}

func open2(
	g graph,
	s pathState2,
) (pathState2, bool) {
	if s.remaining1 <= 0 && s.remaining2 <= 0 {
		// fmt.Printf("\tCannot open. No time remaining\n")
		// no time
		return s, false
	}

	if s.remaining1 >= s.remaining2 {
		if s.isOpen(s.cur1) {
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
		s2.valves = s2.valves.open(s2.cur1)
		s2.remaining1 -= 1
		s2.released += (pressure(s2.remaining1) * pressure(g.getValue(s2.cur1)))
		return s2, true
	}

	if s.isOpen(s.cur2) {
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
	s2.valves = s2.valves.open(s2.cur2)
	s2.remaining2 -= 1
	s2.released += (pressure(s2.remaining2) * pressure(g.getValue(s2.cur2)))

	return s2, true
}

func maximize2(
	g graph,
	s pathState2,
) pathState2 {
	s, ok := open2(g, s)
	if !ok {
		return s
	}

	best := s
	var ts pathState2
	for i := 0; i < numNodes; i++ {
		ts, ok = travel2(
			g,
			s,
			node(i),
		)
		if !ok {
			continue
		}
		tmax := maximize2(g, ts)
		if tmax.released > best.released {
			best = tmax
		}
	}

	return best
}
