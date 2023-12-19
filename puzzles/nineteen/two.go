package nineteen

import (
	"slices"
	"strings"
)

func Two(
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

	const (
		minVal = 1
		maxVal = 4000
	)

	numAccepted := 0

	r := rating{
		x: minVal,
		m: minVal,
		a: minVal,
		s: minVal,
	}

	// naive solution for part 2. this will take forever to run.
	for r.x = minVal; r.x <= maxVal; r.x++ {
		for r.m = minVal; r.m <= maxVal; r.m++ {
			for r.a = minVal; r.a <= maxVal; r.a++ {
				for r.s = minVal; r.s <= maxVal; r.s++ {
					if isAccepted(r) {
						numAccepted += r.sum()
					}
				}
			}
		}
	}

	return numAccepted, nil
}
