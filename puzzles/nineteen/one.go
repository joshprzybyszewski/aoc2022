package nineteen

import (
	"slices"
	"strings"
)

func One(
	input string,
) (int, error) {

	allSteps := make([]step, 0, 522)
	var s step
	for input[0] != '\n' {
		s, input = newStep(input)
		allSteps = append(allSteps, s)
		input = input[1:]
	}
	input = input[1:]

	slices.SortFunc(allSteps, func(a, b step) int {
		return strings.Compare(a.name, b.name)
	})

	startingIndex := slices.IndexFunc(allSteps, func(s step) bool {
		return strings.Compare(s.name, `in`) >= 0
	})

	isAccepted := func(r rating) bool {
		index := startingIndex

		var next string
		var found bool
		var run process

		for {
			for _, run = range allSteps[index].run {
				next, found = run(r)
				if found {
					break
				}
			}
			if next == `A` {
				return true
			} else if next == `R` {
				return false
			}

			index = slices.IndexFunc(allSteps, func(s step) bool {
				return strings.Compare(s.name, next) >= 0
			})
			if index == -1 {
				panic(`unexpected`)
			}
		}
	}

	numAccepted := 0

	for len(input) > 0 && input[0] != '\n' {
		var r rating
		r, input = newRating(input)
		if isAccepted(r) {
			numAccepted += r.sum()
		}
		input = input[1:]
	}

	return numAccepted, nil
}

type process func(r rating) (string, bool)

func newProcess(input string) (process, string, bool) {
	if input[0] == '}' {
		return nil, input[1:], false
	}

	first := input[0]
	second := input[1]

	if second == '<' || second == '>' {
		input = input[2:]
		num := 0
		for input[0] != ':' {
			num *= 10
			num += int(input[0] - '0')
			input = input[1:]
		}
		input = input[1:]
		next := input
		i := 0
		for input[0] != ',' {
			i++
			input = input[1:]
		}
		next = next[:i]
		input = input[1:]
		return func(r rating) (string, bool) {
			if second == '<' {
				switch first {
				case 'x':
					return next, r.x < num
				case 'm':
					return next, r.m < num
				case 'a':
					return next, r.a < num
				case 's':
					return next, r.s < num
				}
			} else if second == '>' {
				switch first {
				case 'x':
					return next, r.x > num
				case 'm':
					return next, r.m > num
				case 'a':
					return next, r.a > num
				case 's':
					return next, r.s > num
				}
			}
			panic(`unexpected`)
		}, input, true
	}

	next := input
	i := 0
	for input[0] != '}' {
		i++
		input = input[1:]
	}
	next = next[:i]

	return func(r rating) (string, bool) {
		return next, true
	}, input, false
}

type step struct {
	name string

	run []process
}

func newStep(input string) (step, string) {
	s := step{
		name: input,
	}

	i := 0
	for input[0] != '{' {
		i++
		input = input[1:]
	}
	s.name = s.name[:i]
	input = input[1:]

	hasMore := true
	for hasMore {
		var p process
		p, input, hasMore = newProcess(input)
		s.run = append(s.run, p)
	}
	input = input[1:]

	return s, input
}

type rating struct {
	x int
	m int
	a int
	s int
}

func newRating(input string) (rating, string) {
	var r rating

	if input[:3] != `{x=` {
		panic(`unexpected`)
	}
	input = input[3:]

	num := 0
	for input[0] != ',' {
		num *= 10
		num += int(input[0] - '0')
		input = input[1:]
	}
	r.x = num

	if input[:3] != `,m=` {
		panic(`unexpected`)
	}
	input = input[3:]
	num = 0

	for input[0] != ',' {
		num *= 10
		num += int(input[0] - '0')
		input = input[1:]
	}
	r.m = num

	if input[:3] != `,a=` {
		panic(`unexpected`)
	}
	input = input[3:]
	num = 0

	for input[0] != ',' {
		num *= 10
		num += int(input[0] - '0')
		input = input[1:]
	}
	r.a = num

	if input[:3] != `,s=` {
		panic(`unexpected`)
	}
	input = input[3:]
	num = 0

	for input[0] != '}' {
		num *= 10
		num += int(input[0] - '0')
		input = input[1:]
	}
	r.s = num

	if input[:1] != `}` {
		panic(`unexpected: ` + input[:1])
	}
	input = input[1:]

	return r, input
}

func (r rating) sum() int {
	return r.x + r.m + r.a + r.s
}
