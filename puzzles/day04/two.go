package day04

import "github.com/joshprzybyszewski/aoc2022/util/lines"

func Two(
	input string,
) (int, error) {

	numCopies := make([]int, 200)
	total := 0

	tmp, cardCopies, j := 0, 0, 0

	var i int

	lines.ForEach(input, func(line string) {
		if len(line) == 0 {
			return
		}

		// one original copy.
		cardCopies = numCopies[i] + 1
		total += cardCopies

		tmp = i + newCard(line).numMatching() + 1
		for j = i + 1; j < tmp; j++ {
			numCopies[j] += cardCopies
		}

		i++
	})

	return total, nil
}
