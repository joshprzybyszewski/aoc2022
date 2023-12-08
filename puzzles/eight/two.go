package eight

import (
	"strings"
)

const numEndingInA = 6

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
) (int, int) {
	an := make(allNodes, numNodes)

	allIndexes := [26][26][26]int{}

	ni := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		an[ni] = newNode(input[:nli])
		allIndexes[an[ni].name[0]-'A'][an[ni].name[1]-'A'][an[ni].name[2]-'A'] = ni
		ni++

		input = input[nli+1:]
	}
	an = an[:ni]
	if ni != numNodes {
		panic(`mistake`)
	}
	for i, n := range an {
		nodes[i] = indexNode{
			left:  allIndexes[n.left[0]-'A'][n.left[1]-'A'][n.left[2]-'A'],
			right: allIndexes[n.right[0]-'A'][n.right[1]-'A'][n.right[2]-'A'],
			isA:   n.name[len(n.name)-1] == 'A',
			isZ:   n.name[len(n.name)-1] == 'Z',
		}
	}

	return allIndexes[0][0][0], allIndexes[25][25][25]
}

func Two(
	input string,
) (uint64, error) {

	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	nodes := make(allIndexNodes, numNodes)
	populateAllIndexNodes(nodes, input)

	curIndexes := ghosts{}
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

	var firstZs ghosts
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

type ghosts [numEndingInA]int

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
		lri++
		if nodes[index].isZ {
			return lri
		}
	}
}

func reduce(
	input ghosts,
) (ghosts, int) {

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
