package sixteen

import "fmt"

func One(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	vs, ns := simplify(valves)

	// 1986 too high
	return getBestPath(
		valves[0].name,
		vs,
		ns,
		30,
	), nil
}

func getBestPath(
	start string,
	vs map[string]*valve,
	ns map[string]*node,
	timeRemaining int,
) int {
	fmt.Printf("getBestPath(\n\t%q\n\t%+v\n\t%+v\n\t%d\n)\n",
		start,
		vs,
		ns,
		timeRemaining,
	)

	// the simplified graph only has nodes that have non-zero valves.
	// Therefore, we need to navigate to one of those nodes to start
	es := getEdges(
		start,
		vs,
		ns,
	)
	max := 0

	all := make(map[string]struct{}, len(ns)-1)
	for name := range ns {
		all[name] = struct{}{}
	}

	for _, e := range es {
		fmt.Printf("Starting at %q\n\twith %d time remaining\n",
			e.dest.name,
			timeRemaining-e.weight,
		)

		s := state{
			cur:           e.dest.name,
			ns:            ns,
			closed:        all,
			timeRemaining: timeRemaining - e.weight,
		}
		em := maximize(
			s,
		)

		if em > max {
			max = em
		}
	}

	return max
}

type state struct {
	cur string

	ns     map[string]*node
	closed map[string]struct{}

	timeRemaining    int
	pressureReleased int
}

func (s state) travel(
	e *edge,
) (state, bool) {
	if s.timeRemaining <= 0 {
		// no time
		return s, false
	}

	if _, ok := s.closed[e.dest.name]; !ok {
		// the destination isn't closed. Not worth going to.
		return s, false
	}

	return state{
		cur:           e.dest.name,
		ns:            s.ns,
		closed:        s.closed,
		timeRemaining: s.timeRemaining - e.weight,
	}, true
}

func (s state) open() (state, bool) {
	if s.timeRemaining <= 0 {
		// no time
		return s, false
	}

	s2 := state{
		cur:              s.cur,
		ns:               s.ns,
		closed:           openValve(s.cur, s.closed),
		timeRemaining:    s.timeRemaining - 1,
		pressureReleased: s.pressureReleased,
	}

	s2.pressureReleased += (s2.timeRemaining * s2.ns[s2.cur].value)

	return s2, true
}

func maximize(
	s state,
) int {
	os, ok := s.open()
	if !ok {
		return 0
	}

	max := os.pressureReleased
	for _, e := range s.ns[s.cur].edges {
		ts, ok := os.travel(e)
		if !ok {
			continue
		}
		tmax := maximize(ts)
		if tmax > max {
			max = tmax
		}
	}

	return max
}

func openValve(
	d string,
	closed map[string]struct{},
) map[string]struct{} {
	c2 := make(map[string]struct{}, len(closed)-1)
	for c := range closed {
		if c == d {
			continue
		}
		c2[c] = struct{}{}
	}
	return c2
}
