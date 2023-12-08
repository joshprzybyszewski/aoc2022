package eight

import (
	"slices"
	"sort"
	"strings"
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

	nodes := make(allIndexNodes, 752)
	nodes = populateAllIndexNodes(nodes, input)

	lri := 0
	ni := 0
	zzzI := len(nodes) - 1
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
