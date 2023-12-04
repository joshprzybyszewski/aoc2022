package four

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (int, error) {

	// skip the zero index
	allCards := make([]card, 1, 200)

	var ci, pi int

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		ci = strings.Index(input, ":") + 1
		pi = strings.Index(input, "|")

		// TODO the numbers are not necessarily delimited by a
		// space characer because they're pretty-printed to line
		// up with columns above and below
		winners := strings.Split(
			strings.TrimSpace(input[ci:pi]),
			" ",
		)
		shown := strings.Split(
			strings.TrimSpace(input[pi+1:nli]),
			" ",
		)

		c := newCard(max(len(winners), len(shown)))
		for _, w := range winners {
			v, err := strconv.Atoi(w)
			if err != nil {
				// TODO reconsider
				continue
			}
			c = c.addWinner(v)
		}

		for _, s := range shown {
			v, err := strconv.Atoi(s)
			if err != nil {
				// TODO reconsider
				continue
			}
			c = c.addShown(v)
		}
		allCards = append(allCards, c)

		input = input[nli+1:]
	}

	numCopies := make([]int, len(allCards))

	tmp, cardCopies := 0, 0

	for i := 1; i < len(allCards); i++ {
		// one original copy.
		numCopies[i]++

		cardCopies = numCopies[i]
		tmp = i + allCards[i].numMatching() + 1
		for j := i + 1; j < tmp; j++ {
			numCopies[j] += cardCopies
		}
	}

	total := 0
	for _, n := range numCopies {
		total += n
	}

	return total, nil
}
