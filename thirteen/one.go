package thirteen

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	sum := 0
	var l, r interface{}
	for i := 0; i < len(lines); i += 3 {
		if lines[i] == `` || lines[i+1] == `` {
			continue
		}
		l, _ = parse(lines[i], 0)
		r, _ = parse(lines[i+1], 0)

		switch compare(l, r) {
		case valid:
			sum += (i / 3) + 1
		case unknown:
			panic(`unexpected`)
		}
	}

	return sum, nil
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
		// compare ints
		if lv < rv {
			return valid
		} else if lv > rv {
			return invalid
		}
		return unknown
	} else if !okl && !okr {
		ls := l.([]interface{})
		rs := r.([]interface{})
		// compare lists
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

	var child interface{}
	var endIndex, ci, bi, v int
	var err error

	for i := start + 1; i < len(line); {
		switch line[i] {
		case '[':
			child, endIndex = parse(line, i)
			output = append(output, child)
			i = endIndex + 1
		case ']':
			return output, i
		case ',':
			i++
		default:
			ci = strings.Index(line[i:], `,`)
			bi = strings.Index(line[i:], `]`)
			if ci == -1 || (bi >= 0 && bi < ci) {
				ci = bi
			}
			v, err = strconv.Atoi(line[i : i+ci])
			if err != nil {
				panic(err)
			}
			output = append(output, v)
			i += ci
		}
	}

	panic(`should not have reached here`)
}
