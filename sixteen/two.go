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

	var best, em pairPaths

	for n1, d1 := range g.startingPositions {
		for n2, d2 := range g.startingPositions {
			if n1 == n2 {
				// no use in going to the same node
				continue
			}
			// fmt.Printf("Starting from %d (spent %d time)\n", n, d)
			em = maximize2(
				g,
				pairPaths{
					one: traveller{
						cur:       node(n1),
						remaining: remaining - time(d1),
					},
					two: traveller{
						cur:       node(n2),
						remaining: remaining - time(d2),
					},
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

type traveller struct {
	cur       node
	remaining time
}

type pairPaths struct {
	one, two traveller

	valves valveState

	released pressure

	prev *pairPaths
}

func (s pairPaths) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s pairPaths) String(
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

func open2(
	g graph,
	s pairPaths,
) (pairPaths, bool) {
	if s.one.remaining <= 0 && s.two.remaining <= 0 {
		// fmt.Printf("\tCannot open. No time remaining\n")
		// no time
		return s, false
	}

	if s.one.remaining > s.two.remaining {
		if s.isOpen(s.two.cur) {
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
		s2.valves = s2.valves.open(s2.one.cur)
		s2.one.remaining -= 1
		s2.released += (pressure(s2.one.remaining) * pressure(g.getValue(s2.one.cur)))
		return s2, true
	}

	if s.isOpen(s.two.cur) {
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
	s2.valves = s2.valves.open(s2.two.cur)
	s2.two.remaining -= 1
	s2.released += (pressure(s2.two.remaining) * pressure(g.getValue(s2.two.cur)))

	return s2, true
}

func maximize2(
	g graph,
	s pairPaths,
) pairPaths {
	if s.one.remaining <= 0 && s.two.remaining <= 0 {
		// fmt.Printf("\tCannot open. No time remaining\n")
		// no time
		return s
	}

	if s.one.remaining != s.two.remaining {
		useOne := s.one.remaining > s.two.remaining
		s, ok := open2(g, s)
		if !ok {
			return s
		}

		// travel and maximize
		best := s
		var ts pairPaths
		for i := 0; i < numNodes; i++ {
			if s.isOpen(node(i)) {
				continue
			}
			if useOne && ts.two.cur == node(i) {
				continue
			} else if !useOne && ts.one.cur == node(i) {
				continue
			}

			ts = s
			ts.prev = &s
			if useOne {
				ts.one.remaining -= time(g.getDistance(ts.one.cur, node(i)))
				ts.one.cur = node(i)
			} else {
				ts.two.remaining -= time(g.getDistance(ts.two.cur, node(i)))
				ts.two.cur = node(i)
			}

			tmax := maximize2(g, ts)
			if tmax.released > best.released {
				best = tmax
			}
		}

		return best
	}

	if s.isOpen(s.one.cur) {
		s2 := s
		s2.prev = &s
		s2.valves = s2.valves.open(s2.one.cur)
		s2.one.remaining -= 1
		s2.released += (pressure(s2.one.remaining) * pressure(g.getValue(s2.one.cur)))
		s = s2
	}
	if s.isOpen(s.two.cur) {
		s2 := s
		s2.prev = &s
		s2.valves = s2.valves.open(s2.two.cur)
		s2.two.remaining -= 1
		s2.released += (pressure(s2.two.remaining) * pressure(g.getValue(s2.two.cur)))
		s = s2
	}

	best := s
	for i := 0; i < numNodes; i++ {
		if s.isOpen(node(i)) || s.one.cur == node(i) || s.two.cur == node(i) {
			continue
		}
		for j := 0; j < numNodes; j++ {
			if i == j {
				continue
			}
			if s.isOpen(node(j)) || s.one.cur == node(j) || s.two.cur == node(j) {
				continue
			}

			s2 := s
			s2.prev = &s
			s2.one.remaining -= time(g.getDistance(s2.one.cur, node(i)))
			s2.one.cur = node(i)
			s2.two.remaining -= time(g.getDistance(s2.two.cur, node(j)))
			s2.two.cur = node(j)

			tmax := maximize2(g, s2)
			if tmax.released > best.released {
				best = tmax
			}
		}
	}

	return best
}
