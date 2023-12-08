package eight

import (
	"slices"
	"strings"
)

type indexNode struct {
	left  int
	right int
	isA   bool
	isZ   bool
}

type allIndexNodes []indexNode

func (ain allIndexNodes) populate(an allNodes) {
	//indexOf works "fastest" if an is sorted
	an.sort()

	for i, n := range an {
		ain[i] = indexNode{
			left:  an.indexOf(n.left),
			right: an.indexOf(n.right),
			isA:   n.name[len(n.name)-1] == 'A',
			isZ:   n.name[len(n.name)-1] == 'Z',
		}
	}
}

type allPositions []int

func (ap allPositions) goLeft(nodes allIndexNodes) bool {
	isAllZ := true
	for i := range ap {
		ap[i] = nodes[ap[i]].left
		if isAllZ && !nodes[ap[i]].isZ {
			isAllZ = false
		}
	}
	slices.Sort(ap)
	return isAllZ
}

func (ap allPositions) goRight(nodes allIndexNodes) bool {
	isAllZ := true
	for i := range ap {
		ap[i] = nodes[ap[i]].right
		if isAllZ && !nodes[ap[i]].isZ {
			isAllZ = false
		}
	}
	slices.Sort(ap)
	return isAllZ
}

func Two(
	input string,
) (int, error) {

	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	ni := 0
	nodes := make(allIndexNodes, 752)
	{
		an := make(allNodes, 750)

		for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
			an[ni] = newNode(input[:nli])
			ni++

			input = input[nli+1:]
		}
		an = an[:ni]

		nodes.populate(an)
		nodes = nodes[:len(an)]
	}

	curIndexes := make(allPositions, len(nodes))
	ci := 0
	for ni = 0; ni < len(nodes); ni++ {
		if !nodes[ni].isA {
			continue
		}
		curIndexes[ci] = ni
		ci++
		if ci == len(curIndexes) {
			break
		}
	}
	curIndexes = curIndexes[:ci]

	numSteps := 0
	lri := 0
	isAllZ := false
	for {
		numSteps++
		if lrs[lri] == 'L' {
			// go left
			isAllZ = curIndexes.goLeft(nodes)
		} else {
			// go right
			isAllZ = curIndexes.goRight(nodes)
		}
		if isAllZ {
			return numSteps, nil
		}

		lri++
		if lri >= len(lrs) {
			lri = 0
		}
	}

	return 0, nil
}
