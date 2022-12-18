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

	for _, e := range es {
		fmt.Printf("Starting at %q\n\twith %d time remaining\n",
			e.dest.name,
			timeRemaining-e.weight,
		)
		em := maximize(
			e.dest.name,
			ns,
			timeRemaining-e.weight,
		)

		if em > max {
			max = em
		}
	}

	return max
}

func maximize(
	start string,
	ns map[string]*node,
	timeRemaining int,
) int {
	if timeRemaining <= 0 {
		return 0
	}

	all := make(map[string]struct{}, len(ns)-1)
	for name := range ns {
		if name == start {
			continue
		}
		all[name] = struct{}{}
	}

	return ((timeRemaining - 1) * ns[start].value) + next(
		start,
		ns,
		all,
		timeRemaining-1,
	)
}

func next(
	cur string,
	ns map[string]*node,
	closed map[string]struct{},
	timeRemaining int,
) int {
	// fmt.Printf("\tLocated at %q\n",
	// 	cur,
	// )
	// fmt.Printf("\t\twith %d time remaining\n",
	// 	timeRemaining,
	// )

	if timeRemaining <= 0 {
		return 0
	}

	max := 0

	// check the destinations we could move to
	for _, e := range ns[cur].edges {
		if e.weight >= timeRemaining {
			// it'll take longer to get there than any benefit we receive
			continue
		}
		if _, ok := closed[e.dest.name]; !ok {
			// the valve at this dest has already been opened!
			continue
		}

		// navigate to the destination and open the valve.
		// This will release a total pressure of the valve's
		// value * how long it will be open.
		// The valve will be open for the time remaining minus
		// the time it takes to get there minus one minute to open it.
		// == (timeRemaining - (e.weight + 1))
		val := (timeRemaining - e.weight - 1) * e.dest.value

		then := next(
			e.dest.name,
			ns,
			openValve(e.dest.name, closed),
			timeRemaining-e.weight-1,
		)
		if val+then > max {
			max = val + then
		}
	}

	return max
}

func openValve(
	d string,
	closed map[string]struct{},
) map[string]struct{} {
	// fmt.Printf("\tOPENING %q\n",
	// 	d,
	// )
	c2 := make(map[string]struct{}, len(closed)-1)
	for c := range closed {
		if c == d {
			continue
		}
		c2[c] = struct{}{}
	}
	return c2
}
