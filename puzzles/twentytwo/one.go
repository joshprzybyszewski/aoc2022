package twentytwo

func One(
	input string,
) (int, error) {

	blocks := convertInput(input)
	blocks = settle(blocks)

	return safeToDisintegrate(blocks), nil
}

func safeToDisintegrate(
	blocks []block,
) int {
	isSupportedBy := make([]int, len(blocks))

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

			isSupportedBy[i]++
			if isSupportedBy[i] > 1 {
				break
			}
		}
	}

	output := 0
	isSafe := false

	// now look above
	for i := range blocks {
		z = blocks[i].maxZ() + 1
		isSafe = true

		for j = i + 1; j < len(blocks); j++ {
			if blocks[j].minZ() != z {
				continue
			}
			// block j rests in the target plane just below block i.
			// If it intersects block i, then i rests on j.

			if !intersectsInXYPlane(blocks[i], blocks[j]) {
				continue
			}

			if isSupportedBy[j] == 1 {
				isSafe = false
				break
			}
		}
		if isSafe {
			output++
		}
	}

	return output
}
