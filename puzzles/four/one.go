package four

import (
	"strconv"
	"strings"
)

type card struct {
	winningNumbers map[int]struct{}
	shown          []int
}

func newCard(size int) card {
	return card{
		winningNumbers: make(map[int]struct{}, size),
	}
}

func (c card) addWinner(w int) card {
	c.winningNumbers[w] = struct{}{}
	return c
}

func (c card) addShown(w int) card {
	c.shown = append(c.shown, w)
	return c
}

func (c card) value() int {
	total := 0
	for _, s := range c.shown {
		if _, ok := c.winningNumbers[s]; !ok {
			continue
		}
		if total == 0 {
			total = 1
		} else {
			total <<= 1
		}
	}
	return total
}

func One(
	input string,
) (int, error) {

	total := 0

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

		total += c.value()

		input = input[nli+1:]
	}

	return total, nil
}
