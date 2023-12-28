package four

import (
	"strings"
)

func Two(
	input string,
) (int, error) {

	numCopies := make([]int, 200)
	total := 0

	tmp, cardCopies, j := 0, 0, 0

	var i int

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		// one original copy.
		cardCopies = numCopies[i] + 1
		total += cardCopies

		tmp = i + newCard(input[:nli]).numMatching() + 1
		for j = i + 1; j < tmp; j++ {
			numCopies[j] += cardCopies
		}

		input = input[nli+1:]
		i++
	}

	return total, nil
}
