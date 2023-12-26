package twelve

import (
	"fmt"
)

type part uint8

const (
	safe    part = 0
	broken  part = 1
	unknown part = 2
)

func (p part) toString() byte {
	switch p {
	case safe:
		return '.'
	case broken:
		return '#'
	case unknown:
		return '?'
	}
	return 'X'
}

func One(
	input string,
) (int, error) {

	var p possibilities
	var groups []int

	var total int

	for len(input) > 0 {
		if input[0] == '\n' {
			input = input[1:]
			continue
		}

		p, groups, input = newPossibilities(input)
		p.build(groups)
		ans := p.answer(groups)
		fmt.Printf("answer: %d\n", ans)
		total += ans

	}

	return total, nil
}
