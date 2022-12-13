package eleven

import (
	"fmt"

	big "github.com/ncw/gmp"
	// big "github.com/ncw/gmp"
	// "math/big"
)

const (
	dividersProduct uint64 = 3 * 5 * 2 * 13 * 11 * 17 * 19 * 7
)

var (
	bigRemDiv = big.NewInt(int64(dividersProduct))
)

type item struct {
	value *big.Int

	cur         monkey
	inspections [numMonkeys]*big.Int
	numRounds   int
}

func newItem(
	startingMonkey monkey,
	e int,
) *item {
	item := item{
		cur:   startingMonkey,
		value: big.NewInt(int64(e)),
	}
	for i := range item.inspections {
		item.inspections[i] = big.NewInt(0)
	}

	return &item
}

// func (i *item) inspect() bool {
// 	i.inspections[i.cur].Add(i.inspections[i.cur], big.NewInt(1))
// 	i.updateValue()
// 	d := i.getDest()
// 	if d < i.cur {
// 		i.numRounds++
// 	}
// 	i.cur = d
// 	i.value.Rem(i.value, bigRemDiv)
// 	// i.value %= dividersProduct

// 	return i.numRounds <= numRounds2
// }

func (i *item) getDest(m monkey) monkey {
	switch m {
	// switch i.cur {
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
	d int64,
) bool {
	return big.NewInt(0).Rem(i.value, big.NewInt(d)).Int64() == 0
	// return i.value%d == 0
}

func (i *item) updateValue(m monkey) {
	switch m {
	// switch i.cur {
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

func (i *item) add(v int64) {
	i.value.Add(i.value, big.NewInt(v))
	// i.value += v
}

func (i *item) mul(v int64) {
	i.value.Mul(i.value, big.NewInt(v))
	// i.value *= v
}

func (i *item) square() {
	i.value.Mul(i.value, i.value)
	// i.value *= i.value
}

func Two(
	input string,
) (string, error) {
	return runNRounds(numRounds2)
}

func runNRounds(
	numRounds int,
) (string, error) {
	fmt.Printf("runNRounds : %d\n", numRounds)
	// items := make([]*item, 0, numItems)
	// for i := range initialValues {
	// 	for j := range initialValues[i] {
	// 		items = append(items, newItem(monkey(i), initialValues[i][j]))
	// 	}
	// }
	monkeys := [numMonkeys][]*item{}
	for i := range initialValues {
		monkeys[i] = make([]*item, 0, numItems)
		for j := range initialValues[i] {
			monkeys[i] = append(monkeys[i], newItem(monkey(i), initialValues[i][j]))
		}
	}
	fmt.Printf("monkeys : %+v\n", monkeys)

	inspections := [numMonkeys]*big.Int{}
	for i := range inspections {
		inspections[i] = big.NewInt(0)
	}
	fmt.Printf("inspections : %+v\n", inspections)

	/*
		var wg sync.WaitGroup

		work := make(chan int, runtime.NumCPU())
		for i := 0; i < runtime.NumCPU(); i++ {
			go func() {
				for ii := range work {
					for items[ii].numRounds <= numRounds {
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

		for _, item := range items {
			if item.numRounds != numRounds+1 {
				return ``, fmt.Errorf("item should have been inspected until it hit the target num rounds")
			}
			for i := range inspections {
				inspections[i].Add(inspections[i], item.inspections[i])
			}
		}
	*/
	var m, i, d int
	for r := 1; r <= numRounds; r++ {
		for m = 0; m < numMonkeys; m++ {
			for i = 0; i < len(monkeys[m]); i++ {
				inspections[m].Add(inspections[m], big.NewInt(1))
				// the monkey inspects the item, and I am NOT relieved
				monkeys[m][i].updateValue(monkey(m))
				monkeys[m][i].value.Rem(monkeys[m][i].value, bigRemDiv)
				d = int(monkeys[m][i].getDest(monkey(m)))
				if d == m {
					panic(`wtf`)
				}
				monkeys[d] = append(monkeys[d], monkeys[m][i])
			}
			monkeys[m] = monkeys[m][:0]
		}
		if r%10 == 0 {
			fmt.Printf("finished %d\n", r)
		}
	}

	var m1, m2 *big.Int
	for i, ni := range inspections {
		fmt.Printf("m%d -> %s\n", i, ni)
		if m1 == nil || ni.Cmp(m1) > 0 {
			m2 = m1
			m1 = ni
		} else if m2 == nil || ni.Cmp(m2) > 0 {
			m2 = ni
		}
	}
	fmt.Printf("1st : %s\n", m1)
	fmt.Printf("2nd : %s\n", m2)

	return big.NewInt(0).Mul(m1, m2).String(), nil
}
