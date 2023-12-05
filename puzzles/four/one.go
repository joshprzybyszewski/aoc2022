package four

import (
	"strings"
)

type card struct {
	winningNumbers map[int]struct{}
	shown          []int
}

func newCard() card {
	return card{
		winningNumbers: make(map[int]struct{}, 10),
		shown:          make([]int, 0, 25),
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

func (c card) numMatching() int {
	total := 0
	for _, s := range c.shown {
		if _, ok := c.winningNumbers[s]; ok {
			total++
		}
	}
	return total
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

	var tmpi, pi int
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

		total += c.value()

		input = input[nli+1:]
	}

	return total, nil
}
