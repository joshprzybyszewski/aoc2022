package four

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (int, error) {

	// allCards := make([]card, 0, 200)
	numCopies := make([]int, 200)
	total := 0

	tmp, cardCopies, j := 0, 0, 0

	var i, ci, pi int

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		ci = strings.Index(input, ":") + 1
		pi = strings.Index(input, "|")

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
			if w == `` {
				continue
			}
			v, err := strconv.Atoi(w)
			if err != nil {
				panic(`ahh`)
			}
			c = c.addWinner(v)
		}

		for _, s := range shown {
			if s == `` {
				continue
			}
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			c = c.addShown(v)
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
