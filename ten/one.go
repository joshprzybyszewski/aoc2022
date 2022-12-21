package ten

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
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

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if ci > 201 {
			break
		}
		if nli == 0 {
			input = input[1:]
			continue
		}
		if input[:nli] == `noop` {
			cycle()
			input = input[nli+1:]
			continue
		}

		// the instruction is "addx A"
		a, err = strconv.Atoi(input[5:nli])
		if err != nil {
			return 0, err
		}

		cycle()
		cycle()
		x += a
		input = input[nli+1:]
	}

	return sum, nil
}
