package sixteen

type duetPath struct {
	one, two traveller

	valves  valveState
	numOpen node

	released pressure
}

func (s *duetPath) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s *duetPath) open(n node) {
	s.valves = s.valves.open(n)
	s.numOpen++
}

func (s *duetPath) openTwo(n1, n2 node) {
	s.valves = s.valves.openTwo(n1, n2)
	s.numOpen += 2
}

func maximizeDuet(
	g *graph,
	s duetPath,
) duetPath {
	if s.one.remaining == 0 && s.two.remaining == 0 {
		// no time
		return s
	}

	if s.one.remaining != s.two.remaining {
		useOne := s.one.remaining > s.two.remaining

		if useOne {
			s.open(s.one.cur)
			s.one.remaining -= 1
			s.released += (pressure(s.one.remaining) * pressure(g.getValue(s.one.cur)))
		} else {
			s.open(s.two.cur)
			s.two.remaining -= 1
			s.released += (pressure(s.two.remaining) * pressure(g.getValue(s.two.cur)))
		}

		if s.numOpen == numNodes {
			return s
		} else if s.numOpen == numNodes-1 {
			// This guy opened the second-to-last valve. Open the other one
			// by setting this one's remaining time to zero.
			if useOne {
				s.one.remaining = 0
			} else {
				s.two.remaining = 0
			}
			return maximizeDuet(g, s)
		}

		// travel and maximize
		best := s
		var ts, tmax duetPath
		var di distance
		for n := node(0); n < numNodes; n++ {
			if s.isOpen(n) || s.one.cur == n || s.two.cur == n {
				continue
			}

			ts = s
			if useOne {
				di = g.getDistance(s.one.cur, n)
				if di >= s.one.remaining {
					// it'll take longer to get there than it's worth
					continue
				}
				ts.one.remaining -= di
				ts.one.cur = n
			} else {
				di = g.getDistance(s.two.cur, n)
				if di >= s.two.remaining {
					// it'll take longer to get there than it's worth
					continue
				}
				ts.two.remaining -= di
				ts.two.cur = n
			}

			tmax = maximizeDuet(g, ts)
			if tmax.released > best.released {
				best = tmax
			}
		}

		return best
	}

	// open both valves
	s.openTwo(s.one.cur, s.two.cur)
	s.one.remaining -= 1
	s.two.remaining -= 1
	// s.one.remaining and s.two.remaining are the same
	s.released += (pressure(s.one.remaining) * pressure(g.getValue(s.one.cur)+g.getValue(s.two.cur)))

	if s.numOpen == numNodes {
		return s
	} else if s.numOpen == numNodes-1 {
		// We just opened two valves at once. Only one remains. Have the closer
		// operator move to the last valve.
		vi := node(0)
		for vi = 0; vi < numNodes; vi++ {
			if !s.isOpen(vi) {
				break
			}
		}
		d1 := g.getDistance(s.one.cur, vi)
		d2 := g.getDistance(s.two.cur, vi)
		ts := s
		if d1 < d2 && ts.one.remaining > d1 {
			ts.one.remaining -= d1
			ts.one.cur = vi
			ts.two.remaining = 0
		} else if ts.two.remaining > d2 {
			ts.two.remaining -= d2
			ts.two.cur = vi
			ts.one.remaining = 0
		} else {
			return s
		}
		return maximizeDuet(g, ts)
	}

	best := s
	var ts, tmax duetPath
	var d1, d2 distance
	for n1, n2 := node(0), node(0); n1 < numNodes; n1++ {
		if s.isOpen(n1) || s.one.cur == n1 || s.two.cur == n1 {
			// don't go to the opposite node or to an open one
			continue
		}
		d1 = g.getDistance(s.one.cur, n1)
		if d1 >= s.one.remaining {
			// it'll take longer to get there than it's worth
			continue
		}

		for n2 = 0; n2 < numNodes; n2++ {
			if n1 == n2 {
				// don't go to the same node
				continue
			}
			if s.isOpen(n2) || s.one.cur == n2 || s.two.cur == n2 {
				// don't go to the opposite node or to an open one
				continue
			}
			d2 = g.getDistance(s.two.cur, n2)
			if d2 >= s.two.remaining {
				// it'll take longer to get there than it's worth
				continue
			}

			ts = s
			// ts.prev = &s
			ts.one.remaining -= d1
			ts.one.cur = n1
			ts.two.remaining -= d2
			ts.two.cur = n2

			tmax = maximizeDuet(g, ts)
			if tmax.released > best.released {
				best = tmax
			}
		}
	}

	return best
}
