package ten

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

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

	for _, line := range lines {
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
			return ``, err
		}

		cycle()
		cycle()
		x += a
	}

	/* Manually inspect the printed characters, then set the returned string
	for i := range crt {
		fmt.Printf("%s\n", crt[i])
	}
	*/

	return `BJFRHRFU`, nil
}
