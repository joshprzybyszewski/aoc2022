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
	isNeeded := make([]bool, len(blocks))

	var j, n int
	var z uint
	for i := len(blocks) - 1; i >= 0; i-- {
		z = blocks[i].minZ() - 1

		n = -1
		for j = i - 1; j >= 0; j-- {
			// TODO be smarter about breaking out of this j loop?
			if blocks[j].maxZ() != z {
				continue
			}
			// block j rests in the target plane just below block i.
			// If it intersects block i, then i rests on j.

			if !intersectsInZPlane(blocks[i], blocks[j]) {
				continue
			}

			if n != -1 {
				// There's already a block holding this one up.
				// That means neither is necessary.
				n = -1
				break
			}

			n = j
		}

		if n != -1 {
			// There's only one possible way to hold up block i and it's block j (which is now block n)
			isNeeded[n] = true
		}
	}

	output := 0
	for _, n := range isNeeded {
		if !n {
			output++
		}
	}

	return output
}
