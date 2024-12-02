package twentytwo

func Two(
	input string,
) (int, error) {

	blocks := convertInput(input)
	blocks = settle(blocks)

	return getSumChainReactions(blocks), nil
}

func getSumChainReactions(
	blocks []block,
) int {
	nodes := make([]node, len(blocks))

	var j int
	var z uint
	for i := len(blocks) - 1; i >= 0; i-- {
		z = blocks[i].minZ() - 1
		if z == 0 {
			// there's no blocks below it. we can't check if it has a single support or not.
			continue
		}

		for j = i - 1; j >= 0; j-- {
			if blocks[j].maxZ() != z {
				// block j is not in the plane below us.
				continue
			}
			// block j rests in the target plane just below block i.
			// If it intersects block i, then i rests on j.

			if !intersectsInXYPlane(blocks[i], blocks[j]) {
				continue
			}

			nodes[i].restsOn = append(nodes[i].restsOn, j)
			nodes[j].supports = append(nodes[j].supports, i)
		}
	}

	output := 0
	for i := range nodes {
		output += getChainReactions(nodes, i)
	}

	return output
}

type node struct {
	supports []int
	restsOn  []int
}

func getChainReactions(
	nodes []node,
	ni int,
) int {

	removed := make(map[int]struct{}, 4)
	toCheck := make([]int, 0, 8)

	markFalling := func(ni int) {
		removed[ni] = struct{}{}
		toCheck = append(toCheck, nodes[ni].supports...)
	}

	markFalling(ni)

	var ok bool

	isFalling := func(ni int) bool {
		for _, support := range nodes[ni].restsOn {
			if _, ok = removed[support]; !ok {
				return false
			}
		}

		return true
	}

	numOthers := 0
	for len(toCheck) > 0 {
		if _, ok = removed[toCheck[0]]; ok {
			// ope, we already removed this node
			toCheck = toCheck[1:]
			continue
		}

		if isFalling(toCheck[0]) {
			numOthers++
			markFalling(toCheck[0])
		}
		toCheck = toCheck[1:]
	}

	return numOthers
}
