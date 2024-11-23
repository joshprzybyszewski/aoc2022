package twentytwo

import "fmt"

func One(
	input string,
) (int, error) {

	blocks := convertInput(input)
	fmt.Print(getPrintableBlockString(blocks))
	blocks = settle(blocks)
	fmt.Print(getPrintableBlockString(blocks))

	// 837 and 412 is the wrong answer: it's too high

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
			// if n > 1 {
			// 	// There's more than one block holding this up. We can stop searching now.
			// 	break
			// }
		}

		fmt.Printf("blocks[%d] has %d supports\n", i, isSupportedBy[i])
	}
	fmt.Printf("\n")

	isNotSafe := make([]bool, len(blocks))

	// now look above
	for i := range blocks {
		z = blocks[i].maxZ() + 1

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
				fmt.Printf("blocks[%d] is supporting blocks[%d]\n", i, j)
				isNotSafe[i] = true
			}
		}
	}

	output := 0
	for i, n := range isNotSafe {
		if n {
			continue
		}
		fmt.Printf("safe to remove blocks[%d]: %+v\n", i, blocks[i])
		output++
	}

	return output
}
