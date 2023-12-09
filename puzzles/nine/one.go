package nine

import (
	"strings"

	"github.com/joshprzybyszewski/aoc2022/util/itoa"
)

const (
	newline = "\n"
	space   = " "
)

func One(
	input string,
) (int, error) {

	total := 0
	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		total += getNextNumber(getLineOfVals(input[:nli]))
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
	var layers [][]int
	layers = append(layers, input)
	var isZeros bool
	var nextLayer []int

	for {
		nextLayer, isZeros = generateDiff(layers[len(layers)-1])
		if isZeros {
			break
		}
		layers = append(layers, nextLayer)
	}

	sum := 0
	for li := len(layers) - 1; li >= 0; li-- {
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
