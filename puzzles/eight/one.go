package eight

import (
	"slices"
	"sort"
	"strings"
)

const (
	startingNode = `AAA`
	targetNode   = `ZZZ`
)

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

func (an allNodes) sort() {
	slices.SortFunc(an, func(a, b node) int {
		return strings.Compare(a.name, b.name)
	})
}

func (an allNodes) indexOf(name string) int {
	i, found := sort.Find(len(an), func(i int) int {
		if an[i].name == name {
			return 0
		}
		return strings.Compare(name, an[i].name)
	})
	if !found {
		panic(`couldn't find: ` + name)
	}
	return i
}

func One(
	input string,
) (int, error) {
	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	ni := 0
	nodes := make(allNodes, 750)

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		nodes[ni] = newNode(input[:nli])
		ni++

		input = input[nli+1:]
	}
	nodes = nodes[:ni]

	nodes.sort()

	numSteps := 0
	lri := 0
	ni = nodes.indexOf(startingNode)
	for {
		numSteps++
		if lrs[lri] == 'L' {
			// go left
			ni = nodes.indexOf(nodes[ni].left)
		} else {
			// go right
			ni = nodes.indexOf(nodes[ni].right)
		}
		if nodes[ni].name == targetNode {
			return numSteps, nil
		}

		lri++
		if lri >= len(lrs) {
			lri = 0
		}
	}
}
