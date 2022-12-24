package twenty

import (
	"fmt"
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

	linkedList := convertToDoublyLinkedList(numbers)

	mixed := mix(linkedList)
	start := -1
	for i := range mixed {
		if mixed[i].val == 0 {
			start = i
			break
		}
	}
	if start == -1 {
		return 0, fmt.Errorf("did not have zero in the data set")
	}
	oneThou := mixed[(start+1000)%len(mixed)].val
	twoThou := mixed[(start+2000)%len(mixed)].val
	threeThou := mixed[(start+3000)%len(mixed)].val

	return oneThou + twoThou + threeThou, nil
}

func mix(nodes []*node) []*node {

	var j int
	for _, n := range nodes {
		if n.val > 0 {
			for j = 0; j < n.val; j++ {
				n.forward()
			}
		} else if n.val < 0 {
			for j = 0; j > n.val; j-- {
				n.backward()
			}
		}
	}

	return nodes
}

func convertToDoublyLinkedList(numbers []int) []*node {
	output := make([]*node, len(numbers))

	for i := range numbers {
		output[i] = newNode(numbers[i])
	}

	for i := 1; i < len(output)-1; i++ {
		output[i].prev = output[i-1]
		output[i].next = output[i+1]
	}

	output[0].prev = output[len(output)-1]
	output[0].next = output[1]
	output[len(output)-1].prev = output[len(output)-2]
	output[len(output)-1].next = output[0]

	return output

}

type node struct {
	val int

	prev *node
	next *node
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

func (n *node) forward() {
	// A <-> n <-> B <-> C
	// A <-> B <-> n <-> C
	a := n.prev
	b := n.next
	c := n.next.next

	a.next = b
	b.prev = a

	b.next = n
	n.prev = b

	n.next = c
	c.prev = n
}

func (n *node) backward() {
	// A <-> B <-> n <-> C
	// A <-> n <-> B <-> C
	a := n.prev.prev
	b := n.prev
	c := n.next

	a.next = n
	n.prev = a

	n.next = b
	b.prev = n

	b.next = c
	c.prev = b
}
