package eleven

import (
	"fmt"
	"strconv"
)

const (
	numMonkeys = 8
	numItems   = 34

	numRounds1 = 20
	numRounds2 = 10000
)

var (
	initialValues = [][]int{
		{99, 67, 92, 61, 83, 64, 98},
		{78, 74, 88, 89, 50},
		{98, 91},
		{59, 72, 94, 91, 79, 88, 94, 51},
		{95, 72, 78},
		{76},
		{69, 60, 53, 89, 71, 88},
		{72, 54, 63, 80},
	}
)

type monkey int

func One(
	input string,
) (string, error) {
	// 	return runNRounds(numRounds1)
	// }

	monkeys := [numMonkeys][]int{}
	for i := range monkeys {
		monkeys[i] = make([]int, 0, numItems)
		monkeys[i] = append(monkeys[i], initialValues[i]...)
	}

	var m, i int
	var nw int
	var nm int

	inspections := [numMonkeys]int{}

	for r := 1; r <= numRounds1; r++ {
		for m = 0; m < numMonkeys; m++ {
			for i = 0; i < len(monkeys[m]); i++ {
				inspections[m]++
				// the monkey inspects the item, and I am relieved
				nw = newWorryLevel(monkey(m), monkeys[m][i]) / 3
				nm = int(getDestMonkey(
					monkey(m),
					nw,
				))
				monkeys[nm] = append(monkeys[nm], nw)
			}
			monkeys[m] = monkeys[m][:0]
		}
	}

	var m1, m2 int
	for _, ni := range inspections {
		fmt.Printf("m%d -> %d\n", i, ni)
		if ni > m1 {
			m2 = m1
			m1 = ni
		} else if ni > m2 {
			m2 = ni
		}
	}
	fmt.Printf("1st : %d\n", m1)
	fmt.Printf("2nd : %d\n", m2)

	return strconv.Itoa(m1 * m2), nil
}

func newWorryLevel(
	m monkey,
	old int,
) int {
	switch m {
	case 0:
		return old * 17
	case 1:
		return old * 11
	case 2:
		return old + 4
	case 3:
		return old * old
	case 4:
		return old + 7
	case 5:
		return old + 8
	case 6:
		return old + 5
	case 7:
		return old + 3
	}
	panic(`unknown monkey`)
}

func getDestMonkey(
	m monkey,
	level int,
) monkey {
	switch m {
	case 0:
		if level%3 == 0 {
			return 4
		} else {
			return 2
		}
	case 1:
		if level%5 == 0 {
			return 3
		} else {
			return 5
		}
	case 2:
		if level%2 == 0 {
			return 6
		} else {
			return 4
		}
	case 3:
		if level%13 == 0 {
			return 0
		} else {
			return 5
		}
	case 4:
		if level%11 == 0 {
			return 7
		} else {
			return 6
		}
	case 5:
		if level%17 == 0 {
			return 0
		} else {
			return 2
		}
	case 6:
		if level%19 == 0 {
			return 7
		} else {
			return 1
		}
	case 7:
		if level%7 == 0 {
			return 1
		} else {
			return 3
		}
	}
	panic(`unknown monkey`)
}
