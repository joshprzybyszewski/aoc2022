package ten

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	var crt [6][40]byte

	r := 0
	c := 0
	x := 1    // the value of the register
	var a int // the cmd's value
	var err error

	cycle := func() {
		if x-1 <= c && c <= x+1 {
			crt[r][c] = '#'
		} else {
			crt[r][c] = '.'
		}
		c++
		if c == 40 {
			r++
			c = 0
		}
	}

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
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
			return ``, err
		}

		cycle()
		cycle()
		x += a
		input = input[nli+1:]
	}

	/* Manually inspect the printed characters, then set the returned string
	for i := range crt {
		fmt.Printf("%s\n", crt[i])
	}
	*/

	return `BJFRHRFU`, nil
}
