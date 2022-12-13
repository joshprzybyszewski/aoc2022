package eleven

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	big "github.com/ncw/gmp"
	// "math/big"
)

type item struct {
	value *big.Int

	cur         monkey
	inspections [numMonkeys]int
	numRounds   int
}

func newItem(e int) *item {
	return &item{
		value: big.NewInt(int64(e)),
	}
}

func (i *item) inspect() bool {
	i.inspections[i.cur]++
	i.updateValue()
	d := i.getDest()
	if d < i.cur {
		i.numRounds++
	}
	i.cur = d

	return i.numRounds < numRounds2
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
	d int64,
) bool {
	return big.NewInt(0).Rem(i.value, big.NewInt(d)).Int64() == 0
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

func (i *item) add(v int64) {
	i.value.Add(i.value, big.NewInt(v))
}

func (i *item) mul(v int64) {
	i.value.Mul(i.value, big.NewInt(v))
}

func (i *item) square() {
	i.value.Mul(i.value, i.value)
}

func Two(
	input string,
) (string, error) {
	items := make([]*item, 0, numItems)
	for i := range initialValues {
		for j := range initialValues[i] {
			items = append(items, newItem(initialValues[i][j]))
		}
	}

	inspections := [numMonkeys]int{}

	work := make(chan *item, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for item := range work {
				t0 := time.Now()
				var dur time.Duration
				ni := 0
				for item.inspect() {
					ni++
					if ni%100 == 0 {
						dur = time.Since(t0)
						fmt.Printf("%13s: completed %d inspections to get to round %d\n", dur, ni, item.numRounds)
						if dur.Seconds() > 30 {
							break
						}
					}
				}
				for i := range inspections {
					inspections[i] += item.inspections[i]
				}
			}
		}()
	}

	for _, item := range items {
		item := item
		work <- item
	}

	var m1, m2 int
	for _, ni := range inspections {
		if ni > m1 {
			m2 = m1
			m1 = ni
		} else if ni > m2 {
			m2 = ni
		}
	}

	return strconv.Itoa(m1 * m2), nil
}
