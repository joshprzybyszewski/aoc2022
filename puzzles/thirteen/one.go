package thirteen

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {

	sum := 0
	var l, r []thing
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

		switch compareSlices(l, r) {
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

func compareSlices(
	l, r []thing,
) answer {
	var a answer
	for i := range l {
		if i >= len(r) {
			return invalid
		}

		a = compare(l[i], r[i])
		if a != unknown {
			return a
		}
	}
	if len(l) < len(r) {
		return valid
	}
	return unknown
}

func compare(
	l, r thing,
) answer {
	if l.slice == nil && r.slice == nil {
		// compare ints
		if l.value < r.value {
			return valid
		} else if l.value > r.value {
			return invalid
		}
		return unknown
	} else if l.slice != nil && r.slice != nil {
		return compareSlices(l.slice, r.slice)
	} else if l.slice == nil {
		return compareSlices([]thing{l}, r.slice)
	} else if r.slice == nil {
		return compareSlices(l.slice, []thing{r})
	}
	panic(`cannot reach here`)
}

type thing struct {
	value int
	slice []thing
}

func parse(
	line string,
	start int,
) ([]thing, int) {
	if line[start] != '[' {
		panic(`got unexpected input`)
	}

	output := make([]thing, 0)

	var child []thing
	var endIndex, ci, bi, v int
	var err error

	for i := start + 1; i < len(line); {
		switch line[i] {
		case '[':
			child, endIndex = parse(line, i)
			output = append(output, thing{
				slice: child,
			})
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
			output = append(output, thing{value: v})
			i += ci
		}
	}

	panic(`should not have reached here`)
}
