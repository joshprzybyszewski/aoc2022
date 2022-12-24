package twenty

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (int, error) {
	numbers := make([]int, 0, 1028)

	var val int
	var err error

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		val, err = strconv.Atoi(input[0:nli])
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, val)
		input = input[nli+1:]
	}

	mixed := mix(numbers)
	oneThou := mixed[1000%len(mixed)]
	twoThou := mixed[2000%len(mixed)]
	threeThou := mixed[3000%len(mixed)]

	return oneThou + twoThou + threeThou, nil
}

func mix(numbers []int) []int {
	indexes := make([]int, len(numbers))

	for i := range indexes {
		indexes[i] = i
	}

	var newIndex, j int
	for i, n := range numbers {
		newIndex = indexes[i] + n
		if n > 0 {
			for j = i + 1; j <= newIndex; j++ {
				indexes[j%len(numbers)]--
				if indexes[j%len(numbers)] < 0 {
					indexes[j%len(numbers)] = len(numbers) - 1
				}
			}
			indexes[i] = newIndex
			if indexes[i] > len(numbers)-1 {
				indexes[i] %= len(numbers)
			}
		} else {
			for j = i - 1; j >= newIndex; j-- {
				if j < 0 {
					// TODO figure out how to remove this if and get the same behavior
					indexes[len(numbers)-(j%len(numbers))]++
					if indexes[len(numbers)-(j%len(numbers))] > len(numbers)-1 {
						indexes[len(numbers)-(j%len(numbers))] = 0
					}
				} else {
					indexes[j]++
					if indexes[j] > len(numbers)-1 {
						indexes[j] = 0
					}
				}
			}
			indexes[i] = newIndex
			for indexes[i] < 0 {
				indexes[i] += len(numbers)
			}
		}
	}

	output := make([]int, len(numbers))
	for i := range output {
		output[indexes[i]] = numbers[i]
	}

	return output
}
