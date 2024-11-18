package day09

import "github.com/joshprzybyszewski/aoc2022/util/lines"

func Two(
	input string,
) (int, error) {

	var p puzzle

	total := 0

	lines.ForEach(input, func(line string) {
		if len(line) == 0 {
			return
		}

		p = newPuzzle(line)
		total += p.getPrev()
	})

	return total, nil
}

func (p *puzzle) getPrev() int {
	cur := 0
	var li int
	for li = 0; li < p.allZerosLayer; li += 2 {
		cur += p.numbers[li][0]
	}
	for li = 1; li < p.allZerosLayer; li += 2 {
		cur -= p.numbers[li][0]
	}
	return cur
}
