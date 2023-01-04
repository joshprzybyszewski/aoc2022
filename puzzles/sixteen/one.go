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
	g := buildGraph(startingNode, &valves)

	return getBestPath(
		&g,
		30,
	), nil
}

func getBestPath(
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

	canBeatBest := func(s soloPath) bool {
		bestLock.Lock()
		br := best.released
		bestLock.Unlock()

		r := pressure(s.remaining)
		pot := pressure(0)

		for n := node(0); n < numNodes; n++ {
			if !s.isOpen(n) {
				pot += r * pressure(g.getValue(n))
				if pot > br {
					return true
				}
				r--
			}
		}

		return false
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
					canBeatBest,
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
