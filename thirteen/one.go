package thirteen

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		if lines[i] == `` || lines[i+1] == `` {
			continue
		}
		l, _ := parse(lines[i], 0)
		r, _ := parse(lines[i+1], 0)

		// fmt.Printf("===========\n")
		// fmt.Printf("Comparing\n")
		// fmt.Printf("\tleft: %+v\n", l)
		// fmt.Printf("\trght: %+v\n", r)

		// fmt.Printf("-----------\n")
		switch compare(l, r) {
		case valid:
			// fmt.Printf("VALID\n")
			sum += (i / 3) + 1
			// case unknown:
			// 	fmt.Printf("UNKNOWN\n")
			// case invalid:
			// 	fmt.Printf("INVALID\n")
		}
		// fmt.Printf("===========\n")
	}

	return strconv.Itoa(sum), nil
}

type answer uint8

const (
	unknown answer = iota
	valid
	invalid
)

func compare(
	l, r interface{},
) answer {
	lv, okl := l.(int)
	rv, okr := r.(int)
	if okl && okr {
		// fmt.Printf("comparing ints\n\tl: %+v\n\tr: %+v\n", lv, rv)
		if lv < rv {
			return valid
		} else if lv > rv {
			return invalid
		}
		return unknown
	} else if !okl && !okr {
		ls := l.([]interface{})
		rs := r.([]interface{})
		// fmt.Printf("comparing slices\n\tl: %+v\n\tr: %+v\n", ls, rs)
		var a answer
		for i := range ls {
			if i >= len(rs) {
				return invalid
			}

			a = compare(ls[i], rs[i])
			if a != unknown {
				return a
			}
		}
		if len(ls) < len(rs) {
			return valid
		}
		return unknown
	} else if okl {
		return compare([]interface{}{lv}, r)
	} else if okr {
		return compare(l, []interface{}{rv})
	}
	panic(`cannot reach here`)
}

func parse(
	line string,
	start int,
) (interface{}, int) {
	if line[start] != '[' {
		panic(`got unexpected input`)
	}

	output := make([]interface{}, 0)

	for i := start + 1; i < len(line); {
		switch line[i] {
		case '[':
			child, endIndex := parse(line, i)
			output = append(output, child)
			i = endIndex + 1
		case ']':
			return output, i
		case ',':
			i++
		default:
			ci := strings.Index(line[i:], `,`)
			bi := strings.Index(line[i:], `]`)
			if ci == -1 || (bi >= 0 && bi < ci) {
				ci = bi
			}
			v, err := strconv.Atoi(line[i : i+ci])
			if err != nil {
				panic(err)
			}
			output = append(output, v)
			i += ci
		}
	}

	panic(`should not have reached here`)
}
