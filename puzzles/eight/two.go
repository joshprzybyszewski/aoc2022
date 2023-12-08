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
	return isAllZ
}

func Two(
	input string,
) (uint64, error) {

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

	firstZs := make([]int, len(curIndexes))
	for ci = range curIndexes {
		firstZs[ci] = getFirstZ(curIndexes[ci], nodes, lrs)
	}

	firstZs = reduce(firstZs)

	mult := uint64(1)
	for _, v := range firstZs {
		mult *= uint64(v)
	}

	// 17129231578253813679 is too high
	// 48823243223 is too low
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
) []int {

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

	for d := 2; ; {
		canDivide, canContinueUp := allDivisibleBy(d)
		if !canContinueUp {
			break
		}

		if !canDivide {
			d++
			continue
		}

		for i := range input {
			input[i] = input[i] / d
		}
	}

	return input
}
