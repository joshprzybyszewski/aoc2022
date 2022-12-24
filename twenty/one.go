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

	// numbers = []int{
	// 	1,
	// 	2,
	// 	-3,
	// 	3,
	// 	-2,
	// 	0,
	// 	4,
	// }

	// fmt.Printf("numbers: %+v\n", numbers)
	linkedList := convertToDoublyLinkedList(numbers)
	// fmt.Printf("linkedList: %+v\n", linkedList)
	// for i := range linkedList {
	// 	fmt.Printf("\tlinkedList[%d]: %+v\n", i, linkedList[i])
	// }

	mixed := mix(linkedList)
	// fmt.Printf("mixed: %+v\n", mixed)
	oneThou := mixed[1000%len(mixed)]
	twoThou := mixed[2000%len(mixed)]
	threeThou := mixed[3000%len(mixed)]

	// 1596 is too low
	return oneThou + twoThou + threeThou, nil
}

func mix(nodes []*node) []int {

	var j int
	head := nodes[0]
	for _, n := range nodes {
		if n.val > 0 {
			for j = 0; j < n.val; j++ {
				head = n.forward(head)
			}
		} else if n.val < 0 {
			for j = 0; j > n.val; j-- {
				head = n.backward(head)
			}
		}
	}

	output := make([]int, 0, len(nodes))
	output = append(output, head.val)
	for n := head.next; n != head; n = n.next {
		output = append(output, n.val)
	}

	return output
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

func (n *node) forward(head *node) *node {
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

	if n == head {
		return b
	}

	return head
}

func (n *node) backward(head *node) *node {
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

	if n == head {
		return b
	}

	return head
}
