package twenty

import (
	"strconv"
	"strings"
)

const (
	decryptionKey = 811589153

	numRoundsOfMixing = 10
)

func Two(
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

	for i := range numbers {
		numbers[i] *= decryptionKey
	}

	linkedList, zero := convertToDoublyLinkedList(numbers)

	for i := 0; i < numRoundsOfMixing; i++ {
		mixSteps(linkedList)
	}

	oneThou := zero.getNthValue(1000 % len(numbers))
	twoThou := zero.getNthValue(2000 % len(numbers))
	threeThou := zero.getNthValue(3000 % len(numbers))

	return oneThou + twoThou + threeThou, nil
}
