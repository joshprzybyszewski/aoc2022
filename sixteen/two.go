package sixteen

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	stdtime "time"
)

var (
	removeMe [numNodes]string
	removeT0 stdtime.Time
)

func Two(
	input string,
) (int, error) {

	removeT0 = stdtime.Now()

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(part1StartingNode, valves)

	// 1652, 2118 is too low
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
	removeMe = names

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

	work := make(chan struct{ i, j int }, numNodes*numNodes)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var em pairPaths
			var d1, d2 distance
			for w := range work {
				d1, d2 = g.startingPositions[w.i], g.startingPositions[w.j]
				fmt.Printf("Starting from <%s,%s> (spent <%d, %d> time)\n", names[w.i], names[w.j], d1, d2)
				em = maximize2(
					g,
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

	// fmt.Printf("info\n")

	// fmt.Println(fullPath(&dj, names))
	// fmt.Println(fullPath(&best, names))

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

func (s pairPaths) allOpen() bool {
	for n := 0; n < numNodes; n++ {
		if !s.valves.isOpen(node(n)) {
			return false
		}
	}
	return true
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

func fullPath(
	s *pairPaths,
	names [numNodes]string,
) string {
	paths := make([]*pairPaths, 100)
	pi := len(paths) - 1
	for ps := s; ps != nil && ps != ps.prev; ps = ps.prev {
		paths[pi] = ps
		pi--
		if pi < 0 {
			break
		}
	}
	paths = paths[pi+1:]

	var sb strings.Builder
	for _, ps := range paths {
		sb.WriteString("==============\n")
		sb.WriteString(fmt.Sprintf("%s\n", ps.String(names)))
	}
	return sb.String()
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
	g graph,
	s pairPaths,
) (pairPaths, bool) {
	if s.one.remaining <= 1 && s.two.remaining <= 1 {
		// fmt.Printf("\tCannot open. No time remaining\n")
		// no time
		return s, false
	}

	if s.one.remaining > s.two.remaining {
		if s.isOpen(s.one.cur) {
			panic(`wtf`)
			// fmt.Printf("\tCannot open. already open\n")
			// the current node is already open. Cannot open
			// return s, false
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
		panic(`wtf`)
		// return s, false
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

func openBoth(
	g graph,
	input pairPaths,
) pairPaths {
	if input.isOpen(input.one.cur) {
		panic(`wtf`)
	}
	if input.isOpen(input.two.cur) {
		panic(`wtf`)
	}

	s := input
	s.prev = &input
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
	g graph,
	s pairPaths,
) pairPaths {

	if stdtime.Since(removeT0).Seconds() > 45 {
		panic(`ahh`)
	}

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

		if no := s.numOpen(); no == numNodes {
			return s
		} else if no == numNodes-1 {
			// there is nowhere to travel to. Open the other one.
			s2 := s
			s2.prev = &s
			if useOne {
				s2.one.remaining = 0
				// open two
				// s2.valves = s2.valves.open(s2.two.cur)
				// s2.two.remaining -= 1
				// s2.released += (pressure(s2.two.remaining) * pressure(g.getValue(s2.two.cur)))
			} else {
				s2.two.remaining = 0
				// open one
				// s2.valves = s2.valves.open(s2.one.cur)
				// s2.one.remaining -= 1
				// s2.released += (pressure(s2.one.remaining) * pressure(g.getValue(s2.one.cur)))
			}
			// return s2
			return maximize2(g, s2)
		}

		// travel and maximize
		best := s
		for i := 0; i < numNodes; i++ {
			if s.one.cur == node(i) || s.two.cur == node(i) || s.isOpen(node(i)) {
				continue
			}

			ts := s
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

	// open both valves
	s = openBoth(g, s)

	if s.allOpen() {
		return s
	}

	// sfp := fullPath(&s, removeMe)
	debug := false
	// strings.Contains(sfp, `Current Node: DD, JJ`) &&
	// 	strings.Contains(sfp, `Current Node: HH, JJ`) &&
	// 	strings.Contains(sfp, `Current Node: HH, BB`) &&
	// 	strings.Contains(sfp, `Open Valves: DD, JJ`) &&
	// 	strings.Contains(sfp, `Time Remaining: 20, 20`)
	if debug {
		// fmt.Printf("\n\ndebugging DJ route:\n%s\n", sfp)
	}

	best := s
	for i := 0; i < numNodes; i++ {
		if s.one.cur == node(i) || s.two.cur == node(i) || s.isOpen(node(i)) {
			if debug {
				fmt.Printf("skipping i = %d = %s\n\n", i, removeMe[i])
			}
			// don't go to the opposite node or to an open one
			continue
		}

		if debug {
			fmt.Printf("Checking i = %d = %s\n", i, removeMe[i])
		}

		for j := 0; j < numNodes; j++ {
			if i == j {
				if debug {
					fmt.Printf("\ti == j = %d = %s\n", j, removeMe[j])
				}
				// don't go to the same node
				continue
			}
			if s.one.cur == node(j) || s.two.cur == node(j) || s.isOpen(node(j)) {
				if debug {
					fmt.Printf("\tskipping j = %d = %s\n", j, removeMe[j])
				}
				// don't go to the opposite node or to an open one
				continue
			}

			if debug {
				fmt.Printf("Checking j = %d = %s\n", j, removeMe[j])
				fmt.Printf("\tmaximizing movement:\n\t%s->%s\n\t%s->%s\n",
					removeMe[s.one.cur], removeMe[i],
					removeMe[s.two.cur], removeMe[j],
				)
			}

			s2 := s
			s2.prev = &s
			s2.one.remaining -= time(g.getDistance(s2.one.cur, node(i)))
			s2.one.cur = node(i)
			s2.two.remaining -= time(g.getDistance(s2.two.cur, node(j)))
			s2.two.cur = node(j)

			if debug {
				fmt.Printf("s2 = %+v\n", s2)
			}

			tmax := maximize2(g, s2)
			if debug {
				fmt.Printf("tmax.released = %d\n", tmax.released)
				fmt.Printf("tmax = %+v\n", tmax)
			}
			if tmax.released > best.released {
				best = tmax
			}
		}
	}

	return best
}

var (
	djPrefix = `==============
	Current Node: DD, JJ
		Open Valves: (none)
		Minute: 1
		Time Remaining: 25, 24
		Pressure released: 0

	==============
	Current Node: DD, JJ
		Open Valves: DD
		Minute: 2
		Time Remaining: 24, 24
		Pressure released: 480

	==============
	Current Node: HH, JJ
		Open Valves: DD
		Minute: 2
		Time Remaining: 20, 24
		Pressure released: 480

	==============
	Current Node: HH, JJ
		Open Valves: DD, JJ
		Minute: 3
		Time Remaining: 20, 23
		Pressure released: 963

	==============
	Current Node: HH, BB
		Open Valves: DD, JJ
		Minute: 6
		Time Remaining: 20, 20
		Pressure released: 963


`
)
