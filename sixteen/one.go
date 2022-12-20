package sixteen

import (
	"runtime"
	"sync"
)

const (
	// This is assumed in the puzzle.
	startingNode = `AA`
)

func One(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(startingNode, valves)

	return getBestPath(
		valves,
		&g,
		30,
	), nil
}

func getBestPath(
	valves []*valve,
	g *graph,
	remaining distance,
) int {

	var wg sync.WaitGroup
	var best soloPath
	var bestLock sync.Mutex
	checkBest := func(o soloPath) {
		bestLock.Lock()
		defer bestLock.Unlock()
		if o.released > best.released {
			best = o
		}
	}

	work := make(chan int, numNodes)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var em soloPath
			var d distance
			for w := range work {

				d = g.startingPositions[w]
				em = maximize(
					g,
					soloPath{
						cur:       node(w),
						remaining: remaining - d,
					},
				)
				checkBest(em)
				wg.Done()
			}
		}()
	}

	for n1 := range g.startingPositions {
		wg.Add(1)
		work <- n1
	}

	wg.Wait()
	close(work)

	return int(best.released)
}

func maximize(
	g *graph,
	s soloPath,
) soloPath {
	if s.remaining <= 1 {
		// no time
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
		tmax = maximize(g, ts)
		if tmax.released > best.released {
			best = tmax
		}
	}

	return best
}
