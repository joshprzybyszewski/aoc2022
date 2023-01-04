package sixteen

type beatBestChecker func(s soloPath) bool

type soloPath struct {
	cur       node
	remaining distance

	valves valveState

	released pressure
}

func (s *soloPath) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func maximize(
	g *graph,
	s soloPath,
	bbc beatBestChecker,
) soloPath {
	if s.remaining <= 1 {
		// no time
		return s
	}
	if !bbc(s) {
		// can't beat the best, so why try?
		return s
	}

	s.valves = s.valves.open(s.cur)
	s.remaining -= 1
	s.released += (pressure(s.remaining) * pressure(g.getValue(s.cur)))

	best := s
	var ts, tmax soloPath
	var d distance
	for n := node(0); n < numNodes; n++ {
		if s.isOpen(n) {
			// the destination is already open. Not worth going to.
			continue
		}
		d = g.getDistance(s.cur, n)
		if s.remaining <= d {
			// by the time we get there, we won't be able to open it
			continue
		}

		ts = s
		ts.remaining -= d
		ts.cur = n
		tmax = maximize(g, ts, bbc)
		if tmax.released > best.released {
			best = tmax
		}
	}

	return best
}
