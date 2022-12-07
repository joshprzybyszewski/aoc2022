package four

import (
	"fmt"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

func newAssignment(r string) (assignment, error) {
	vals := strings.Split(r, `-`)
	if len(vals) != 2 {
		return assignment{}, fmt.Errorf("should provide two values: %q", r)
	}

	start, err := strconv.Atoi(vals[0])
	if err != nil {
		return assignment{}, err
	}

	end, err := strconv.Atoi(vals[1])
	if err != nil {
		return assignment{}, err
	}

	return assignment{
		start: start,
		end:   end,
	}, nil
}

func (a assignment) isSubset(other assignment) bool {
	return a.start >= other.start && a.end <= other.end
}

func (a assignment) overlaps(other assignment) bool {
	return (a.start >= other.start && a.start <= other.end) ||
		(a.end >= other.start && a.end <= other.end)
}

func fullyContained(a, b assignment) bool {
	return a.isSubset(b) || b.isSubset(a)
}

func overlapping(a, b assignment) bool {
	return a.overlaps(b) || b.overlaps(a)
}

func One(
	input string,
) (string, error) {
	ass, err := convertInputToAssignments(input)
	if err != nil {
		return ``, err
	}

	total := 0
	for _, as := range ass {
		if fullyContained(as[0], as[1]) {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func convertInputToAssignments(
	input string,
) ([][2]assignment, error) {
	lines := strings.Split(input, "\n")

	output := make([][2]assignment, 0, len(lines))

	for _, line := range lines {
		if line == `` {
			continue
		}
		as := [2]assignment{}
		ranges := strings.Split(line, `,`)
		for i, r := range ranges {
			a, err := newAssignment(r)
			if err != nil {
				return nil, err
			}
			as[i] = a
		}
		output = append(output, as)
	}

	return output, nil
}
