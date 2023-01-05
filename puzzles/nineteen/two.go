package nineteen

import (
	"runtime"
	"sync"
)

const (
	part2Minutes = 32
)

func Two(
	input string,
) (int, error) {

	all, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	var vals [3]int
	var wg sync.WaitGroup
	work := make(chan int, len(vals))
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				vals[w] = harvest(&all[w], part2Minutes)
				wg.Done()
			}
		}()
	}

	for i := 0; i < len(vals); i++ {
		wg.Add(1)
		work <- i
	}

	wg.Wait()

	return vals[0] * vals[1] * vals[2], nil
}
