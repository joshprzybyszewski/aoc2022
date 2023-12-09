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
	iterations := [200][]int{}

	var ii int
	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		iterations[ii] = getLineOfVals(input[:nli])
		input = input[nli+1:]
		ii++
	}

	total := 0

	for i := range iterations {
		total += getNextNumer(iterations[i])
	}

	return total, nil
}

func getLineOfVals(
	line string,
) []int {
	vals := make([]int, 0, 25)

	for si := strings.Index(line, space); si >= 0; si = strings.Index(line, space) {
		vals = append(vals, itoa.Int(line[:si]))
		line = line[si+1:]
	}

	if line != `` {
		vals = append(vals, itoa.Int(line))
	}
	return vals
}

func getNextNumer(
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

	placeholder := 0
	for li := len(layers) - 1; li >= 0; li-- {
		placeholder += layers[li][len(layers[li])-1]
	}
	return placeholder
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
