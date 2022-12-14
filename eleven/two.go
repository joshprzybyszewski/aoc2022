package eleven

import (
	"fmt"
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
) *item {
	return &item{
		cur:   startingMonkey,
		value: uint64(e),
	}
}

func (i *item) inspect() {
	i.inspections[i.cur]++
	i.updateValue()
	next := i.getDest()
	if next < i.cur {
		i.numRounds++
	}
	i.cur = next
	i.value %= dividersProduct
}

func (i *item) getDest() monkey {
	switch i.cur {
	case 0:
		if i.divisibleBy(3) {
			return 4
		} else {
			return 2
		}
	case 1:
		if i.divisibleBy(5) {
			return 3
		} else {
			return 5
		}
	case 2:
		if i.divisibleBy(2) {
			return 6
		} else {
			return 4
		}
	case 3:
		if i.divisibleBy(13) {
			return 0
		} else {
			return 5
		}
	case 4:
		if i.divisibleBy(11) {
			return 7
		} else {
			return 6
		}
	case 5:
		if i.divisibleBy(17) {
			return 0
		} else {
			return 2
		}
	case 6:
		if i.divisibleBy(19) {
			return 7
		} else {
			return 1
		}
	case 7:
		if i.divisibleBy(7) {
			return 1
		} else {
			return 3
		}
	}
	panic(`unknown monkey`)
}

func (i *item) divisibleBy(
	d uint64,
) bool {
	return i.value%d == 0
}

func (i *item) updateValue() {
	switch i.cur {
	case 0:
		// old * 17
		i.mul(17)
	case 1:
		// old * 11
		i.mul(11)
	case 2:
		// old + 4
		i.add(4)
	case 3:
		// old * old
		i.square()
	case 4:
		// old + 7
		i.add(7)
	case 5:
		// old + 8
		i.add(8)
	case 6:
		// old + 5
		i.add(5)
	case 7:
		// old + 3
		i.add(3)
	default:
		panic(`unknown monkey`)
	}
}

func (i *item) add(v uint64) {
	i.value += v
}

func (i *item) mul(v uint64) {
	i.value *= v
}

func (i *item) square() {
	i.value *= i.value
}

func Two(
	input string,
) (string, error) {
	return runNRounds(numRounds2)
}

func runNRounds(
	numRounds int,
) (string, error) {
	items := make([]*item, 0, numItems)
	for i := range initialValues {
		for j := range initialValues[i] {
			items = append(items, newItem(monkey(i), initialValues[i][j]))
		}
	}

	var wg sync.WaitGroup

	work := make(chan int, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for ii := range work {
				for items[ii].numRounds < numRounds {
					items[ii].inspect()
				}
				wg.Done()
			}
		}()
	}

	for ii := range items {
		wg.Add(1)
		work <- ii
	}
	wg.Wait()
	close(work)

	inspections := [numMonkeys]uint64{}

	for _, item := range items {
		for i := range inspections {
			inspections[i] += item.inspections[i]
		}
	}

	var m1, m2 uint64
	for _, ni := range inspections {
		if ni > m1 {
			m2 = m1
			m1 = ni
		} else if ni > m2 {
			m2 = ni
		}
	}

	return fmt.Sprintf("%d", m1*m2), nil
}
