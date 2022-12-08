package four

import "strings"

type pair [2]assignment

func newPair(
	line string,
) (pair, error) {
	ranges := strings.Split(line, `,`)

	p := pair{}
	for i, r := range ranges {
		a, err := newAssignment(r)
		if err != nil {
			return pair{}, err
		}
		p[i] = a
	}
	return p, nil
}

func (p pair) fullyContained() bool {
	return (p[0].start >= p[1].start && p[0].end <= p[1].end) ||
		(p[1].start >= p[0].start && p[1].end <= p[0].end)
}

func (p pair) overlapping() bool {
	return (p[0].start >= p[1].start && p[0].start <= p[1].end) ||
		(p[0].end >= p[1].start && p[0].end <= p[1].end) ||
		(p[1].start >= p[0].start && p[1].start <= p[0].end) ||
		(p[1].end >= p[0].start && p[1].end <= p[0].end)
}
