package nineteen

import (
	"runtime"
	"sync"
)

const (
	part1Minutes = 24
)

func One(
	input string,
) (int, error) {
	all, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	var vals [numBlueprints]int
	var wg sync.WaitGroup
	work := make(chan int, len(vals))
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				vals[w] = harvest(&all[w], part1Minutes)
				wg.Done()
			}
		}()
	}

	for i := 0; i < len(vals); i++ {
		wg.Add(1)
		work <- i
	}

	wg.Wait()

	total := 0
	for i := range vals {
		total += ((i + 1) * vals[i])
	}

	return total, nil
}

func harvest(
	b *blueprint,
	minutes int,
) int {
	s := newInitialStuff()

	s = maximizeGeodes(
		b,
		s,
		minutes,
	)
	return s.bank.geode
}
