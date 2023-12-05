package four

import (
	"strings"
)

type card struct {
	winners [2]uint64
	shown   [25]int
}

func newCard(input string) card {
	tmpi := 0
	tmpVal := 0
	pi := strings.Index(input, "|")
	c := card{}
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

	tmpi = pi + 1
	tmpVal = 0
	pi = 0
	for ; tmpi < len(input); tmpi++ {
		if input[tmpi] == ' ' {
			if tmpVal != 0 {
				c.shown[pi] = tmpVal
				pi++
			}
			tmpVal = 0
			continue
		}
		tmpVal *= 10
		tmpVal += int(input[tmpi] - '0')
	}
	if tmpVal != 0 {
		c.shown[pi] = tmpVal
	}

	return c
}

func (c card) addWinner(w int) card {
	if w >= 64 {
		c.winners[1] |= 1 << (w - 64)
	} else {
		c.winners[0] |= 1 << w
	}
	return c
}

func (c card) isWinner(s int) bool {
	if s >= 64 {
		return c.winners[1]&(1<<(s-64)) != 0
	}
	return c.winners[0]&(1<<s) != 0
}

func (c card) numMatching() int {
	total := 0
	for _, s := range c.shown {
		if c.isWinner(s) {
			total++
		}
	}
	return total
}

func (c card) value() int {
	total := 0
	for _, s := range c.shown {
		if !c.isWinner(s) {
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

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		total += newCard(input[:nli]).value()

		input = input[nli+1:]
	}

	return total, nil
}
