package thirteen

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {

	sum := 0
	var l, r interface{}
	var i, nli, nli2 int
	var line1, line2 string

	for len(input) > 0 {
		nli = strings.Index(input, "\n")
		if nli < 0 {
			break
		} else if nli == 0 {
			// skip empty line
			input = input[1:]
			continue
		} else {
			line1 = input[:nli]
		}

		nli2 = nli + 1 + strings.Index(input[nli+1:], "\n")
		if nli2 == nli+1 {
			// skip past empty line
			input = input[nli2+1:]
			continue
		} else if nli2 < nli {
			line2 = input[nli+1:]
			input = input[:0]
		} else {
			line2 = input[nli+1 : nli2]
			input = input[nli2+1:]
		}

		l, _ = parse(line1, 0)
		r, _ = parse(line2, 0)

		switch compare(l, r) {
		case valid:
			sum += i + 1
		case unknown:
			panic(`unexpected`)
		}
		i++
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
