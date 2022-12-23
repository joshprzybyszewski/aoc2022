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
	g := buildGraph(startingNode, valves)

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
	var bestLock sync.Mutex
	checkBest := func(o duetPath) {
		bestLock.Lock()
		defer bestLock.Unlock()
		if o.released > best.released {
			best = o
		}
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
						one: traveller{
							cur:       node(w.i),
							remaining: remaining - d1,
						},
						two: traveller{
							cur:       node(w.j),
							remaining: remaining - d2,
						},
					},
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
