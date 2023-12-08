package eight

import (
	"strings"
)

const numNodes = 750

type node struct {
	name string

	left  string
	right string
}

func newNode(line string) node {
	return node{
		name:  line[:3],
		left:  line[7:10],
		right: line[12:15],
	}
}

type allNodes []node

func One(
	input string,
) (int, error) {
	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	nodes := make(allIndexNodes, numNodes)
	ni, zzzI := populateAllIndexNodes(nodes, input)

	lri := 0
	for {
		if lrs[lri%len(lrs)] == 'L' {
			// go left
			ni = nodes[ni].left
		} else {
			// go right
			ni = nodes[ni].right
		}
		lri++

		if ni == zzzI {
			return lri, nil
		}

	}
}
