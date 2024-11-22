package twentytwo

import (
	"fmt"
	"slices"

	"github.com/joshprzybyszewski/aoc2022/util/strutil"
)

func convertInput(
	input string,
) []block {

	output := make([]block, 0, 1203)

	var skip int
	var x uint
	i := 0

	for {
		x, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 {
			break
		}
		output = append(output, block{})
		output[i].a.x = x
		if input[skip] != ',' {
			panic(`unexpected`)
		}
		input = input[skip+1:]
		output[i].a.y, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 || input[skip] != ',' {
			panic(`unexpected: ` + input[:4])
		}
		input = input[skip+1:]
		output[i].a.z, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 || input[skip] != '~' {
			panic(`unexpected: ` + input[:4])
		}
		input = input[skip+1:]

		output[i].b.x, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 || input[skip] != ',' {
			panic(`unexpected`)
		}
		input = input[skip+1:]
		output[i].b.y, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 || input[skip] != ',' {
			panic(`unexpected: ` + input[:4])
		}
		input = input[skip+1:]
		output[i].b.z, skip = strutil.UintBeforeNonInt(input)
		if skip == 0 || input[skip] != '\n' {
			panic(`unexpected: ` + input[:4])
		}
		input = input[skip+1:]

		i++
	}

	return output
}

type coord struct {
	x, y, z uint
}

type block struct {
	a, b coord
}

func (b block) minX() uint {
	return min(b.a.x, b.b.x)
}

func (b block) maxX() uint {
	return max(b.a.x, b.b.x)
}

func (b block) minY() uint {
	return min(b.a.y, b.b.y)
}

func (b block) maxY() uint {
	return max(b.a.y, b.b.y)
}

func (b block) minZ() uint {
	return min(b.a.z, b.b.z)
}

func (b block) maxZ() uint {
	return max(b.a.z, b.b.z)
}

func willStop(above, below block) bool {
	if above.minZ() <= below.maxZ() {
		return false
	}
	return intersectsInZPlane(above, below)
}

func intersectsInZPlane(one, two block) bool {
	if one.maxX() < two.minX() || one.minX() > two.maxX() {
		// one's x value cannot intersect two's
		return false
	}
	if one.maxY() < two.minY() || one.minY() > two.maxY() {
		// one's y value cannot intersect two's
		return false
	}

	return true
}

func moveToAbove(above, below block) block {
	zDiff := above.minZ() - 1 - below.maxZ()
	above.a.z -= zDiff
	above.b.z -= zDiff
	return above
}

func moveToGround(above block) block {
	zDiff := above.minZ() - 1
	above.a.z -= zDiff
	above.b.z -= zDiff
	return above
}

func settle(
	blocks []block,
) []block {

	slices.SortFunc(blocks, func(a, b block) int {
		if a.minZ() != b.minZ() {
			return int(a.minZ()) - int(b.minZ())
		}

		if a.maxZ() != b.maxZ() {
			return int(a.maxZ()) - int(b.maxZ())
		}

		if a.minX() != b.minX() {
			return int(a.minX()) - int(b.minX())
		}

		if a.maxX() != b.maxX() {
			return int(a.maxX()) - int(b.maxX())
		}

		if a.minY() != b.minY() {
			return int(a.minY()) - int(b.minY())
		}

		if a.maxY() != b.maxY() {
			return int(a.maxY()) - int(b.maxY())
		}
		panic(`wat`)
	})

	var j, t int
	for i := range blocks {
		// if blocks[i].minZ() == 1 {
		// 	continue
		// }
		t = -1
		for j = i - 1; j >= 0; j-- {
			if willStop(blocks[i], blocks[j]) {
				fmt.Printf("---STOPPED---\n")
				fmt.Printf("blocks[i]: %+v\n", blocks[i])
				fmt.Printf("blocks[j]: %+v\n", blocks[j])
				t = j
				break
			}
		}

		if t >= 0 {
			blocks[i] = moveToAbove(blocks[i], blocks[t])
		} else {
			blocks[i] = moveToGround(blocks[i])
		}
	}
	return blocks
}
