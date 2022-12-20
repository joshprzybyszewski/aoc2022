package sixteen

import (
	"fmt"
	"runtime"
	"sync"
	stdtime "time"
)

func Two(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(part1StartingNode, valves)

	// 1652, 2118, 2534 is too low
	// 2706 is not the right answer
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

	var wg sync.WaitGroup
	var best pairPaths
	var bestLock sync.Mutex
	checkBest := func(o pairPaths) {
		bestLock.Lock()
		defer bestLock.Unlock()
		if o.released > best.released {
			best = o
		}
	}

	// numNodes^2 is plenty of space
	work := make(chan struct{ i, j int }, numNodes*numNodes)
	t0 := stdtime.Now()
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var em pairPaths
			var d1, d2 distance
			for w := range work {
				if stdtime.Since(t0).Seconds() > 30 {
					// skip
					wg.Done()
					continue
				}

				d1, d2 = g.startingPositions[w.i], g.startingPositions[w.j]
				em = maximize2(
					&g,
					pairPaths{
						one: traveller{
							cur:       node(w.i),
							remaining: remaining - time(d1),
						},
						two: traveller{
							cur:       node(w.j),
							remaining: remaining - time(d2),
						},
					},
				)
				checkBest(em)
				wg.Done()
			}
		}()
	}

	for n1 := range g.startingPositions {
		for n2 := n1 + 1; n2 < len(g.startingPositions); n2++ {
			wg.Add(1)
			work <- struct {
				i int
				j int
			}{
				i: n1,
				j: n2,
			}
		}
	}

	wg.Wait()
	close(work)

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

	// prev *pairPaths
}

func (s pairPaths) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s pairPaths) numOpen() int {
	no := 0
	for n := 0; n < numNodes; n++ {
		if s.valves.isOpen(node(n)) {
			no++
		}
	}
	return no
}

func (s pairPaths) String(
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
	mr := s.one.remaining
	if s.two.remaining > mr {
		mr = s.two.remaining
	}
	return fmt.Sprintf(
		"Current Node: %s, %s\n"+
			"\t%s\n"+
			"\tMinute: %d\n"+
			"\tTime Remaining: %d, %d\n"+
			"\tPressure released: %d\n",
		names[s.one.cur], names[s.two.cur],
		openValves,
		26-mr,
		s.one.remaining, s.two.remaining,
		s.released,
	)
}

func open2(
	g *graph,
	s pairPaths,
) (pairPaths, bool) {
	if s.one.remaining <= 1 && s.two.remaining <= 1 {
		// no time
		return s, false
	}

	if s.one.remaining > s.two.remaining {
		// if s.isOpen(s.one.cur) {
		// 	panic(`wtf`)
		// }

		s2 := s
		// s2.prev = &s
		s2.valves = s2.valves.open(s2.one.cur)
		s2.one.remaining -= 1
		s2.released += (pressure(s2.one.remaining) * pressure(g.getValue(s2.one.cur)))
		return s2, true
	}

	// if s.isOpen(s.two.cur) {
	// 	panic(`wtf`)
	// }

	s2 := s
	// s2.prev = &s
	s2.valves = s2.valves.open(s2.two.cur)
	s2.two.remaining -= 1
	s2.released += (pressure(s2.two.remaining) * pressure(g.getValue(s2.two.cur)))

	return s2, true
}

func openBoth(
	g *graph,
	input pairPaths,
) pairPaths {
	// if input.isOpen(input.one.cur) {
	// 	panic(`wtf`)
	// }
	// if input.isOpen(input.two.cur) {
	// 	panic(`wtf`)
	// }

	s := input
	// s.prev = &input
	s.valves = s.valves.open(s.one.cur).open(s.two.cur)
	s.one.remaining -= 1
	s.two.remaining -= 1
	// s.one.remaining and s.two.remainig are the same
	// s.released += (pressure(s.one.remaining) * pressure(g.getValue(s.one.cur))) +
	// (pressure(s.two.remaining) * pressure(g.getValue(s.two.cur)))
	s.released += (pressure(s.one.remaining) * pressure(g.getValue(s.one.cur)+g.getValue(s.two.cur)))
	return s
}

func maximize2(
	g *graph,
	s pairPaths,
) pairPaths {
	if s.one.remaining <= 0 && s.two.remaining <= 0 {
		// no time
		return s
	}

	if s.one.remaining != s.two.remaining {
		useOne := s.one.remaining > s.two.remaining
		s, ok := open2(g, s)
		if !ok {
			return s
		}

		if no := s.numOpen(); no == numNodes {
			return s
		} else if no == numNodes-1 {
			// This guy opened the second-to-last valve. Open the other one
			// by setting this one's remaining time to zero.
			s2 := s
			// s2.prev = &s
			if useOne {
				s2.one.remaining = 0
			} else {
				s2.two.remaining = 0
			}
			return maximize2(g, s2)
		}

		// travel and maximize
		best := s
		var ts, tmax pairPaths
		var di distance
		for i := 0; i < numNodes; i++ {
			if s.isOpen(node(i)) || s.one.cur == node(i) || s.two.cur == node(i) {
				continue
			}

			ts = s
			// ts.prev = &s
			if useOne {
				di = g.getDistance(s.one.cur, node(i))
				if di >= distance(s.one.remaining) {
					// it'll take longer to get there than it's worth
					continue
				}
				ts.one.remaining -= time(di)
				ts.one.cur = node(i)
			} else {
				di = g.getDistance(s.two.cur, node(i))
				if di >= distance(s.two.remaining) {
					// it'll take longer to get there than it's worth
					continue
				}
				ts.two.remaining -= time(di)
				ts.two.cur = node(i)
			}

			tmax = maximize2(g, ts)
			if tmax.released > best.released {
				best = tmax
			}
		}

		return best
	}

	// open both valves
	s = openBoth(g, s)

	if no := s.numOpen(); no == numNodes {
		return s
	} else if no == numNodes-1 {
		// We just opened two valves at once. Only one remains. Have the closer
		// operator move to the last valve.
		vi := 0
		for vi = 0; vi < numNodes; vi++ {
			if !s.isOpen(node(vi)) {
				break
			}
		}
		d1 := g.getDistance(s.one.cur, node(vi))
		d2 := g.getDistance(s.two.cur, node(vi))
		ts := s
		// ts.prev = &s
		if d1 < d2 {
			ts.one.remaining -= time(d1)
			ts.one.cur = node(vi)
			ts.two.remaining = 0
		} else {
			ts.two.remaining -= time(d2)
			ts.two.cur = node(vi)
			ts.one.remaining = 0
		}
		return maximize2(g, ts)
	}

	best := s
	var ts, tmax pairPaths
	var di1, dj2 distance
	for i, j := 0, 0; i < numNodes; i++ {
		if s.isOpen(node(i)) || s.one.cur == node(i) || s.two.cur == node(i) {
			// don't go to the opposite node or to an open one
			continue
		}
		di1 = g.getDistance(s.one.cur, node(i))
		if di1 >= distance(s.one.remaining) {
			// it'll take longer to get there than it's worth
			continue
		}

		for j = 0; j < numNodes; j++ {
			if i == j {
				// don't go to the same node
				continue
			}
			if s.isOpen(node(j)) || s.one.cur == node(j) || s.two.cur == node(j) {
				// don't go to the opposite node or to an open one
				continue
			}
			dj2 = g.getDistance(s.two.cur, node(j))
			if dj2 >= distance(s.two.remaining) {
				// it'll take longer to get there than it's worth
				continue
			}

			ts = s
			// ts.prev = &s
			ts.one.remaining -= time(di1)
			ts.one.cur = node(i)
			ts.two.remaining -= time(dj2)
			ts.two.cur = node(j)

			tmax = maximize2(g, ts)
			if tmax.released > best.released {
				best = tmax
			}
		}
	}

	return best
}
