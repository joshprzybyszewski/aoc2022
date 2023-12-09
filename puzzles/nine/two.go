package nine

import "strings"

func Two(
	input string,
) (int, error) {

	total := 0

	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		total += getFirstNumber(getLineOfVals(input[:nli]))
		input = input[nli+1:]
	}

	return total, nil
}

func getFirstNumber(
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

	cur := 0
	for li := len(layers) - 1; li >= 0; li-- {
		cur = layers[li][0] - cur
	}
	return cur
}
