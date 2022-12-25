package twenty

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	exampleNumbers = []int{
		1,
		2,
		-3,
		3,
		-2,
		0,
		4,
	}
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

	linkedList, zero := convertToDoublyLinkedList(numbers)
	if zero == nil {
		return 0, fmt.Errorf("did not have zero in the data set")
	}

	mixSteps(linkedList, zero)

	oneThou := zero.getNthValue(1000 % len(numbers))
	twoThou := zero.getNthValue(2000 % len(numbers))
	threeThou := zero.getNthValue(3000 % len(numbers))

	return oneThou + twoThou + threeThou, nil
}

func convertToDoublyLinkedList(numbers []int) ([]*node, *node) {
	output := make([]*node, len(numbers))
	var zero *node

	for i := range numbers {
		output[i] = newNode(numbers[i])
		if output[i].val == 0 {
			if zero != nil {
				panic(`has more than one zero in the data-set`)
			}
			zero = output[i]
		}
	}

	for i := 1; i < len(output)-1; i++ {
		output[i].prev = output[i-1]
		output[i].next = output[i+1]
	}

	output[0].prev = output[len(output)-1]
	output[0].next = output[1]
	output[len(output)-1].prev = output[len(output)-2]
	output[len(output)-1].next = output[0]

	return output, zero
}

func mixSteps(
	nodes []*node,
	zero *node,
) {
	for _, n := range nodes {
		if n.val > 0 {
			n.forwardSteps(n.val % (len(nodes) - 1))
		} else if n.val < 0 {
			n.backwardSteps((-n.val) % (len(nodes) - 1))
		}
	}
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

func (n *node) forwardSteps(steps int) {
	if steps == 0 {
		// nothing to do
		return
	}
	// A <-> n <-> B ... C <-> D
	// A <-> B       ... C <-> n <-> D
	// keep in mind that C could equal A or B
	a := n.prev
	b := n.next

	c := n
	for s := 0; s < steps; s++ {
		c = c.next
	}
	d := c.next

	if c == a {
		// n moves after a, where it already is.
		return
	}

	a.next = b
	b.prev = a

	c.next = n
	n.prev = c

	n.next = d
	d.prev = n
}

func (n *node) backwardSteps(steps int) {
	if steps == 0 {
		// nothing to do
		return
	}
	// A <-> B       ... C <-> n <-> D
	// A <-> n <-> B ... C <-> D
	// keep in mind that C could equal A or B

	c := n.prev
	d := n.next

	b := n
	for s := 0; s < steps; s++ {
		b = b.prev
	}
	a := b.prev

	if b == d {
		// n moves before d, where it already is.
		return
	}

	a.next = n
	n.prev = a

	n.next = b
	b.prev = n

	c.next = d
	d.prev = c
}

func (n *node) getNthValue(steps int) int {
	if steps == 0 {
		// nothing to do
		return n.val
	}

	iter := n
	for s := 0; s < steps; s++ {
		iter = iter.next
	}
	return iter.val
}
