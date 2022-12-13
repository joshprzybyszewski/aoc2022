package eleven

import (
	"fmt"
	"strconv"
	"time"

	big "github.com/ncw/gmp"
	// "math/big"
)

var (
	tmp   = big.NewInt(0)
	big1  = big.NewInt(1)
	big2  = big.NewInt(2)
	big3  = big.NewInt(3)
	big4  = big.NewInt(4)
	big5  = big.NewInt(5)
	big7  = big.NewInt(7)
	big8  = big.NewInt(8)
	big11 = big.NewInt(11)
	big13 = big.NewInt(13)
	big17 = big.NewInt(17)
	big19 = big.NewInt(19)
)

func Two(
	input string,
) (string, error) {
	monkeys := [numMonkeys][]*big.Int{}
	for i := range monkeys {
		monkeys[i] = make([]*big.Int, 0, numItems)
		for j := range initialValues[i] {
			monkeys[i] = append(monkeys[i], big.NewInt(int64(initialValues[i][j])))
		}
	}

	var m, i int
	var nw *big.Int
	var nm int

	inspections := [numMonkeys]int64{}
	// for i := range inspections {
	// 	inspections[i] = big.NewInt(0)
	// }

	t0 := time.Now()
	fmt.Printf("Starting the program at %s\n", t0)

	for r := 1; r <= numRounds2; r++ {
		for m = 0; m < numMonkeys; m++ {
			for i = 0; i < len(monkeys[m]); i++ {
				inspections[m]++
				// inspections[m].Add(inspections[m], big1)
				// the monkey inspects the item, and I am not relieved
				nw = newBigWorryLevel(monkey(m), monkeys[m][i])
				nm = int(getBigDestMonkey(
					monkey(m),
					nw,
				))
				monkeys[nm] = append(monkeys[nm], nw)
			}
			monkeys[m] = monkeys[m][:0]
		}
		fmt.Printf("round %d,%s", r, time.Since(t0))
		for m = 0; m < numMonkeys; m++ {
			fmt.Printf(",%d", inspections[m])
		}
		fmt.Printf("\n")
		if time.Since(t0) > 45*time.Second {
			break
		}
	}

	var m1, m2 int64
	// var m1, m2 *big.Int
	for _, ni := range inspections {
		if ni > m1 {
			m2 = m1
			m1 = ni
		} else if ni > m2 {
			m2 = ni
		}
		// if m1 == nil || ni.Cmp(m1) > 1 {
		// 	m2 = m1
		// 	m1 = ni
		// } else if m2 == nil || ni.Cmp(m2) > 1 {
		// 	m2 = ni
		// }
	}

	return strconv.Itoa(int(m1 * m2)), nil
	// return tmp.Mul(m1, m2).String(), nil
}

func newBigWorryLevel(
	m monkey,
	old *big.Int,
) *big.Int {
	switch m {
	case 0:
		// return old.Mul(old, big17)
		return old.MulInt32(old, 17)
	case 1:
		// return old.Mul(old, big11)
		return old.MulInt32(old, 11)
	case 2:
		// return old.Add(old, big4)
		return old.AddUint32(old, 4)
	case 3:
		return old.Mul(old, old)
	case 4:
		return old.AddUint32(old, 7)
		// return old.Add(old, big7)
	case 5:
		return old.AddUint32(old, 8)
		// return old.Add(old, big8)
	case 6:
		return old.AddUint32(old, 5)
		// return old.Add(old, big5)
	case 7:
		return old.AddUint32(old, 3)
		// return old.Add(old, big3)
	}
	panic(`unknown monkey`)
}

func getBigDestMonkey(
	m monkey,
	level *big.Int,
) monkey {
	switch m {
	case 0:
		if tmp.Rem(level, big3).Int64() == 0 {
			return 4
		} else {
			return 2
		}
	case 1:
		if tmp.Rem(level, big5).Int64() == 0 {
			return 3
		} else {
			return 5
		}
	case 2:
		if tmp.Rem(level, big2).Int64() == 0 {
			return 6
		} else {
			return 4
		}
	case 3:
		if tmp.Rem(level, big13).Int64() == 0 {
			return 0
		} else {
			return 5
		}
	case 4:
		if tmp.Rem(level, big11).Int64() == 0 {
			return 7
		} else {
			return 6
		}
	case 5:
		if tmp.Rem(level, big17).Int64() == 0 {
			return 0
		} else {
			return 2
		}
	case 6:
		if tmp.Rem(level, big19).Int64() == 0 {
			return 7
		} else {
			return 1
		}
	case 7:
		if tmp.Rem(level, big7).Int64() == 0 {
			return 1
		} else {
			return 3
		}
	}
	panic(`unknown monkey`)
}
