package nine

import (
	"strings"

	"github.com/joshprzybyszewski/aoc2022/util/itoa"
)

const (
	newline = "\n"
	space   = " "
)

type puzzle struct {
	numbers [22][22]int

	layerZeroMaxIndex int
	maxLayer          int
}

func newPuzzle(line string) puzzle {
	p := puzzle{}
	i := 1

	for si := strings.Index(line, space); si >= 0; si = strings.Index(line, space) {
		p.numbers[0][i] = itoa.Int(line[:si])
		i++
		line = line[si+1:]
	}

	p.layerZeroMaxIndex = i - 1

	p.populate()

	return p
}

func (p *puzzle) populate() {
	hasNums := true
	var i int
	li, pi := 1, 0

	for hasNums {
		hasNums = false
		for i = 1; i <= p.layerZeroMaxIndex-li; i++ {
			p.numbers[li][i] = p.numbers[pi][i+1] - p.numbers[pi][i]
			if !hasNums && p.numbers[li][i] != 0 {
				hasNums = true
			}
		}

		if !hasNums {
			break
		}

		pi++
		li++
	}
	p.maxLayer = pi

}

func (p *puzzle) getNext() int {
	return 0
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

func getLineOfVals(
	line string,
) []int {
	vals := make([]int, 21)
	vi := 0

	for si := strings.Index(line, space); si >= 0; si = strings.Index(line, space) {
		vals[vi] = itoa.Int(line[:si])
		vi++
		line = line[si+1:]
	}

	vals[vi] = itoa.Int(line)

	return vals
}

func getNextNumber(
	input []int,
) int {
	layers := make([][]int, 20)
	li := 0
	layers[li] = input
	var isZeros bool
	var nextLayer []int

	for {
		nextLayer, isZeros = generateDiff(layers[li])
		if isZeros {
			break
		}
		li++
		layers[li] = nextLayer
	}

	sum := 0
	for ; li >= 0; li-- {
		sum += layers[li][len(layers[li])-1]
	}
	return sum
}

func generateDiff(
	input []int,
) ([]int, bool) {
	output := make([]int, len(input)-1)
	isZeros := true
	for i := range output {
		output[i] = input[i+1] - input[i]
		if isZeros && output[i] != 0 {
			isZeros = false
		}
	}
	return output, isZeros
}
