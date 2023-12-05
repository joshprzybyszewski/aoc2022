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

	var i, tmpi, pi int
	var tmpVal int

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		pi = strings.Index(input, "|")
		c := newCard()

		tmpVal = 0
		for tmpi = strings.Index(input, ":") + 1; tmpi < pi; tmpi++ {
			if input[tmpi] == ' ' {
				if tmpVal != 0 {
					c = c.addWinner(tmpVal)
				}
				tmpVal = 0
				continue
			}
			tmpVal *= 10
			tmpVal += int(input[tmpi] - '0')
		}
		if tmpVal != 0 {
			c = c.addWinner(tmpVal)
		}

		tmpVal = 0
		for tmpi = pi + 1; tmpi < nli; tmpi++ {
			if input[tmpi] == ' ' {
				if tmpVal != 0 {
					c = c.addShown(tmpVal)
				}
				tmpVal = 0
				continue
			}
			tmpVal *= 10
			tmpVal += int(input[tmpi] - '0')
		}
		if tmpVal != 0 {
			c = c.addShown(tmpVal)
		}

		// one original copy.
		cardCopies = numCopies[i] + 1
		total += cardCopies

		tmp = i + c.numMatching() + 1
		for j = i + 1; j < tmp; j++ {
			numCopies[j] += cardCopies
		}

		input = input[nli+1:]
		i++
	}

	return total, nil
}
