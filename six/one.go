package six

import (
	"fmt"
	"strconv"
)

func One(
	input string,
) (string, error) {
	// fmt.Printf("input is: %q\n", input)

	rs := []rune(input)
	for i := 3; i < len(rs); {
		if rs[i] == rs[i-1] { // xxAA -> Ayyy
			i += 3
		} else if rs[i] == rs[i-2] || // xAxA -> xAyy
			rs[i-1] == rs[i-3] || // AxAx -> Axyy
			rs[i-1] == rs[i-2] { // xAAx -> Axyy
			i += 2
		} else if rs[i] == rs[i-3] || // AxxA -> xxAy
			rs[i-3] == rs[i-2] { // AAxx -> Axxy
			i++
		} else {
			// fmt.Printf("4 unique at marker: %d\n", i+1)
			// fmt.Printf("String: %q\n", input[i-3:i+1])
			// fmt.Printf("Index: %d\n", i)
			return strconv.Itoa(i + 1), nil
		}
	}

	return ``, fmt.Errorf("Couldn't find unique window of 4")
}
