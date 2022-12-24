package twenty

import (
	"fmt"
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

	numbers = []int{
		1,
		2,
		-3,
		3,
		-2,
		0,
		4,
	}

	for i := range numbers {
		numbers[i] *= decryptionKey
	}

	// fmt.Printf("numbers: %+v\n", numbers)
	linkedList := convertToDoublyLinkedList(numbers)
	// fmt.Printf("linkedList: %+v\n", linkedList)
	// for i := range linkedList {
	// 	fmt.Printf("\tlinkedList[%d]: %+v\n", i, linkedList[i])
	// }

	for i := 0; i < numRoundsOfMixing; i++ {
		linkedList = mix(linkedList)
	}

	start := -1
	for i := range linkedList {
		if linkedList[i].val == 0 {
			start = i
			break
		}
	}
	if start == -1 {
		return 0, fmt.Errorf("did not have zero in the data set")
	}
	// fmt.Printf("linkedList: %+v\n", linkedList)
	oneThou := linkedList[(start+1000)%len(linkedList)].val
	// fmt.Printf("oneThou: %+v\n", oneThou)
	twoThou := linkedList[(start+2000)%len(linkedList)].val
	// fmt.Printf("twoThou: %+v\n", twoThou)
	threeThou := linkedList[(start+3000)%len(linkedList)].val
	// fmt.Printf("threeThou: %+v\n", threeThou)

	// 1596 is too low
	return oneThou + twoThou + threeThou, nil
}
