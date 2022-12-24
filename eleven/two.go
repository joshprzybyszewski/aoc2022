package eleven

import (
	"runtime"
	"sync"
)

const (
	dividersProduct uint64 = 3 * 5 * 2 * 13 * 11 * 17 * 19 * 7
)

type item struct {
	value uint64

	cur         monkey
	inspections [numMonkeys]uint64
	numRounds   int
}

func newItem(
	startingMonkey monkey,
	e int,
) item {
	return item{
		cur:   startingMonkey,
		value: uint64(e),
	}
}

func (i *item) inspect() {
	i.inspections[i.cur]++

	// the monkeys that are most common to least: 5,4,6,2,1,3,0,7
	switch i.cur {
	case 5:
		i.value += 8
		if i.value%17 == 0 {
			i.cur = 0
		} else {
			i.cur = 2
		}
		i.numRounds++
	case 4:
		i.value += 7
		if i.value%11 == 0 {
			i.cur = 7
		} else {
			i.cur = 6
		}
	case 6:
		i.value += 5
		if i.value%19 == 0 {
			i.cur = 7
		} else {
			i.cur = 1
			i.numRounds++
		}
	case 2:
		i.value += 4
		if i.value%2 == 0 {
			i.cur = 6
		} else {
			i.cur = 4
		}
	case 1:
		i.value *= 11
		if i.value%5 == 0 {
			i.cur = 3
		} else {
			i.cur = 5
		}
	case 3:
		i.value *= i.value
		if i.value%13 == 0 {
			i.cur = 0
			i.numRounds++
		} else {
			i.cur = 5
		}
	case 0:
		i.value *= 17
		if i.value%3 == 0 {
			i.cur = 4
		} else {
			i.cur = 2
		}
	case 7:
		i.value += 3
		if i.value%7 == 0 {
			i.cur = 1
		} else {
			i.cur = 3
		}
		i.numRounds++
	}

	i.value %= dividersProduct
}

func Two(
	input string,
) (int64, error) {
	return runNRounds(numRounds2), nil
}

func runNRounds(
	numRounds int,
) int64 {
	var items [numItems]item
	ii := 0
	for i := range initialValues {
		for j := range initialValues[i] {
			items[ii] = newItem(monkey(i), initialValues[i][j])
			ii++
		}
	}

	var wg sync.WaitGroup
	inspections := [numMonkeys]uint64{}
	var inspectionsLock sync.Mutex
	transfer := func(item *item) {
		inspectionsLock.Lock()
		defer inspectionsLock.Unlock()
		for i := range inspections {
			inspections[i] += item.inspections[i]
		}
	}

	work := make(chan *item, numItems)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for item := range work {
				for item.numRounds < numRounds {
					item.inspect()
				}
				transfer(item)
				wg.Done()
			}
		}()
	}

	for ii := range items {
		wg.Add(1)
		work <- &items[ii]
	}
	wg.Wait()
	close(work)

	var m1, m2 uint64
	for _, ni := range inspections {
		if ni > m1 {
			m2 = m1
			m1 = ni
		} else if ni > m2 {
			m2 = ni
		}
	}

	return int64(m1 * m2)
}
