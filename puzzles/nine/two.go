package nine

import "strings"

func Two(
	input string,
) (int, error) {

	var p puzzle

	total := 0

	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		p = newPuzzle(input[:nli])
		total += p.getPrev()

		input = input[nli+1:]
	}

	return total, nil
}

func (p *puzzle) getPrev() int {
	cur := 0
	for li := p.maxLayer; li >= 0; li-- {
		cur = p.numbers[li][0] - cur
	}
	return cur
}
