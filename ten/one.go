package ten

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	ci := -19 // the clock index (init to -19 = 20-1 to make modulo easier)
	x := 1    // the value of the register
	var a int // the cmd's value
	var err error

	sum := 0
	cycle := func() {
		if ci%40 == 0 {
			sum += (x * (ci + 20))
		}
		ci++
	}

	for _, line := range lines {
		if ci > 201 {
			break
		}
		if line == `` {
			continue
		}
		if line == `noop` {
			cycle()
			continue
		}

		// the instruction is "addx A"
		a, err = strconv.Atoi(line[5:])
		if err != nil {
			return 0, err
		}

		cycle()
		cycle()
		x += a
	}

	return sum, nil
}
