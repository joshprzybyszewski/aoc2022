package sixteen

import (
	"runtime"
	"sync"
)

func Two(
	input string,
) (int, error) {

	valves, err := getValves(input)
	if err != nil {
		return 0, err
	}
	g := buildGraph(startingNode, &valves)

	return getBestPathForDuet(
		&g,
		26,
	), nil
}

func getBestPathForDuet(
	g *graph,
	remaining distance,
) int {

	var wg sync.WaitGroup
	var best duetPath
	var bestLock sync.RWMutex
	checkBest := func(o duetPath) {
		bestLock.Lock()
		defer bestLock.Unlock()
		if o.released > best.released {
			best = o
		}
	}

	canBeatBest := func(dp duetPath) bool {
		bestLock.RLock()
		br := best.released
		bestLock.RUnlock()

		r1 := pressure(dp.one.remaining)
		r2 := pressure(dp.two.remaining)
		pot := dp.released

		for n := node(0); n < numNodes; n++ {
			if !dp.isOpen(n) {
				if r1 > r2 {
					r1--
					pot += r1 * pressure(g.getValue(n))
					if pot > br {
						return true
					}
					r1--
				} else {
					r2--
					pot += r2 * pressure(g.getValue(n))
					if pot > br {
						return true
					}
					r2--
					if r2 < 0 {
						break
					}
				}
			}
		}

		return false
	}

	// numNodes^2 is plenty of space
	work := make(chan struct{ i, j int }, numNodes*numNodes)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var p duetPath
			var d1, d2 distance
			for w := range work {
				d1 = g.startingPositions[w.i]
				d2 = g.startingPositions[w.j]
				p = maximizeDuet(
					g,
					duetPath{
						one: traveler{
							cur:       node(w.i),
							remaining: remaining - d1,
						},
						two: traveler{
							cur:       node(w.j),
							remaining: remaining - d2,
						},
					},
					canBeatBest,
				)
				checkBest(p)
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
