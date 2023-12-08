package eight

import (
	"strings"
)

type indexNode struct {
	left  int
	right int
	isA   bool
	isZ   bool
}

type allIndexNodes []indexNode

func populateAllIndexNodes(
	nodes allIndexNodes,
	input string,
) allIndexNodes {
	an := make(allNodes, len(nodes))

	ni := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		an[ni] = newNode(input[:nli])
		ni++

		input = input[nli+1:]
	}
	an = an[:ni]
	//indexOf works "fastest" if an is sorted
	an.sort()
	for i, n := range an {
		nodes[i] = indexNode{
			left:  an.indexOf(n.left),
			right: an.indexOf(n.right),
			isA:   n.name[len(n.name)-1] == 'A',
			isZ:   n.name[len(n.name)-1] == 'Z',
		}
	}

	return nodes[:len(an)]

}

func Two(
	input string,
) (uint64, error) {

	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	nodes := make(allIndexNodes, 752)
	nodes = populateAllIndexNodes(nodes, input)

	curIndexes := make([]int, len(nodes))
	ci := 0
	ni := 0
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

	firstZs := make([]int, len(curIndexes))
	for ci = range curIndexes {
		firstZs[ci] = getFirstZ(curIndexes[ci], nodes, lrs)
	}

	var gcf int
	firstZs, gcf = reduce(firstZs)

	mult := uint64(gcf)
	for _, v := range firstZs {
		mult *= uint64(v)
	}

	return mult, nil
}

func getFirstZ(
	index int,
	nodes allIndexNodes,
	lrs string,
) int {
	lri := 0
	for {
		if lrs[lri%len(lrs)] == 'L' {
			// go left
			index = nodes[index].left
		} else {
			// go right
			index = nodes[index].right
		}
		if nodes[index].isZ {
			return lri + 1
		}

		lri++
	}
}

func reduce(
	input []int,
) ([]int, int) {

	allDivisibleBy := func(d int) (bool, bool) {
		for i := range input {
			if d > input[i]/2 {
				return false, false
			}

			if input[i]%d != 0 {
				return false, true
			}
		}
		return true, true
	}
	gcf := 1

	for d := 2; ; {
		canDivide, canContinueUp := allDivisibleBy(d)
		if !canContinueUp {
			break
		}

		if !canDivide {
			d++
			continue
		}

		gcf *= d

		for i := range input {
			input[i] = input[i] / d
		}
	}

	return input, gcf
}
