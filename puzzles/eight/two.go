package eight

import (
	"strings"
)

const numEndingInA = 6

type indexNode struct {
	left  int
	right int
	isZ   bool
}

type allIndexNodes []indexNode

func populateAllIndexNodes(
	nodes allIndexNodes,
	input string,
) (int, int, ghosts) {
	an := make(allNodes, numNodes)

	allIndexes := [26][26][26]int{}
	var gs ghosts
	gi := 0

	ni := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		an[ni] = newNode(input[:nli])
		allIndexes[an[ni].name[0]-'A'][an[ni].name[1]-'A'][an[ni].name[2]-'A'] = ni
		if an[ni].name[2] == 'A' {
			gs[gi] = ni
			gi++
		}
		ni++

		input = input[nli+1:]
	}
	an = an[:ni]
	// if ni != numNodes {
	// 	panic(`mistake`)
	// }
	for i, n := range an {
		nodes[i] = indexNode{
			left:  allIndexes[n.left[0]-'A'][n.left[1]-'A'][n.left[2]-'A'],
			right: allIndexes[n.right[0]-'A'][n.right[1]-'A'][n.right[2]-'A'],
			isZ:   n.name[len(n.name)-1] == 'Z',
		}
	}

	return allIndexes[0][0][0], allIndexes[25][25][25], gs
}

func Two(
	input string,
) (int, error) {

	nli := strings.Index(input, "\n")
	lrs := input[:nli]

	input = input[nli+2:]

	nodes := make(allIndexNodes, numNodes)
	_, _, curIndexes := populateAllIndexNodes(nodes, input)

	var firstZs ghosts
	for ci := range curIndexes {
		firstZs[ci] = getFirstZ(curIndexes[ci], nodes, lrs)
	}

	var gcf int
	firstZs, gcf = reduce(firstZs)

	mult := gcf
	for _, v := range firstZs {
		mult *= v
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
	steps := 0
	for {
		if lrs[lri] == 'L' {
			// go left
			index = nodes[index].left
		} else {
			// go right
			index = nodes[index].right
		}
		steps++
		if nodes[index].isZ {
			return steps
		}
		lri++
		if lri == len(lrs) {
			lri = 0
		}
	}
}

func reduce(
	gs ghosts,
) (ghosts, int) {

	gi := 0
	var cannotDivide bool

	gcf := 1
	d := 2
	d2 := d + d

	for {
		for gi = range gs {
			if d2 > gs[gi] {
				// If the divisor is greater than gs[gi] / 2, then we know we've
				// hit our max divisor.
				// TODO I think that we could use the sqrt of gs[gi] to infer a stop
				// condition, but I'm not sure.
				return gs, gcf
			}

			if gs[gi]%d != 0 {
				cannotDivide = true
				break
			}
		}

		if cannotDivide {
			cannotDivide = false
			// Ideally, we'd be incrementing up to the next prime. But knowing primes is
			// hard, so we're just going to increment one at a time.
			d++
			d2 += 2
			continue
		}

		gcf *= d

		for gi = range gs {
			gs[gi] = gs[gi] / d
		}
	}
}
