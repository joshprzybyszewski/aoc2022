package day03

import "fmt"

func One(
	inputString string,
) (int, error) {
	fmt.Printf("starting day 3 pt 1\n")
	sum := 0

	var l1, l2 int
	var ok bool

	input := []byte(inputString)
	maxI := len(input) - 8 // need space at the end for mul(x,y)

	getL1 := func(i *int) bool {
		l1 = 0
		for *i < maxI {
			if input[*i] == ',' {
				*i++
				return true
			}
			if input[*i] < '0' || input[*i] > '9' {
				return false
			}
			l1 *= 10
			l1 += int(input[*i] - '0')
			*i++
		}
		return false
	}

	getL2 := func(i *int) bool {
		l2 = 0
		for *i < maxI {
			if input[*i] == ')' {
				*i++
				return true
			}
			if input[*i] < '0' || input[*i] > '9' {
				return false
			}
			l2 *= 10
			l2 += int(input[*i] - '0')
			*i++
		}
		return false
	}

	for i := 0; i < maxI; {
		// fmt.Printf("%s", string(input[i]))
		if input[i] != 'm' {
			i++
			continue
		}
		if input[i+1] != 'u' {
			i += 1
			continue
		}
		if input[i+2] != 'l' {
			i += 2
			continue
		}
		if input[i+3] != '(' {
			i += 3
			continue
		}
		i += 4
		ok = getL1(&i)
		if !ok {
			fmt.Printf("\nnot l1. i = %d\n", i)
			continue
		}
		fmt.Printf("\nl1 = %d\n", l1)

		ok = getL2(&i)
		if !ok {
			fmt.Printf("\nnot l2. i = %d\n", i)
			continue
		}
		fmt.Printf("\nl2 = %d\n", l2)
		fmt.Printf("adding (%d * %d)\n", l1, l2)
		sum += (l1 * l2)
	}

	return sum, nil
}
