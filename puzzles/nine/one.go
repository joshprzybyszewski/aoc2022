package nine

import (
	"strings"

	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

const (
	newline = "\n"
	space   = " "
)

type puzzle struct {
	numbers [21][21]int

	layerZeroMaxIndex int
	allZerosLayer     int
}

func newPuzzle(line string) puzzle {
	p := puzzle{}
	i := 0

	for si := strings.Index(line, space); si >= 0; si = strings.Index(line, space) {
		p.numbers[0][i] = strutil.Int(line[:si])
		i++
		line = line[si+1:]
	}
	p.numbers[0][i] = strutil.Int(line)

	p.layerZeroMaxIndex = i

	p.populate()

	return p
}

func (p *puzzle) populate() {

	var i int
	li, pi := 1, 0

	for {
		for i = 0; i <= p.layerZeroMaxIndex-li; i++ {
			p.numbers[li][i] = p.numbers[pi][i+1] - p.numbers[pi][i]
		}

		if p.numbers[li] == p.numbers[20] {
			// the row is all zeros
			break
		}

		pi++
		li++
	}
	p.allZerosLayer = li
}

func (p *puzzle) getNext() int {
	sum := 0
	i := p.layerZeroMaxIndex
	for li := 0; li < p.allZerosLayer; li++ {
		sum += p.numbers[li][i]
		i--
	}
	return sum
}

func One(
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
		total += p.getNext()

		input = input[nli+1:]
	}

	return total, nil
}
