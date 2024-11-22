package twentytwo

import (
	"fmt"
	"strconv"
)

func getPrintableBlockString(
	blocks []block,
) string {

	getIndexIntersectingX := func(
		x, z uint,
	) int {
		var output int = -1
		for i, b := range blocks {
			if b.a.z > z || b.b.z < z {
				// TODO optimize based on assumptions about sorted blocks
				continue
			}

			if b.a.x > x || b.b.x < x {
				continue
			}

			if output != -1 {
				// sentinel value meaning "many"
				return -2
			}
			output = i
		}
		return output
	}

	getIndexIntersectingY := func(
		y, z uint,
	) int {
		var output int = -1
		for i, b := range blocks {
			if b.a.z > z || b.b.z < z {
				// TODO optimize based on assumptions about sorted blocks
				continue
			}

			if b.a.y > y || b.b.y < y {
				continue
			}

			if output != -1 {
				// sentinel value meaning "many"
				return -2
			}
			output = i
		}
		return output
	}

	var maxX, maxY, maxZ uint
	for _, b := range blocks {
		maxX = max(maxX, b.a.x, b.b.x)
		maxY = max(maxY, b.a.y, b.b.y)
		maxZ = max(maxZ, b.a.z, b.b.z)
	}
	output := `=========================`

	output += "\n"
	output += ` x ` + "\n"
	for x := uint(0); x <= maxX; x++ {
		output += string('0' + byte(x%10))
	}
	output += "\n"

	getXDisplay := func(x, z uint) byte {
		index := getIndexIntersectingX(x, z)
		if index == -1 {
			return '.'
		}
		if index == -2 {
			return '?'
		}
		return toChar(index)
	}

	for z := maxZ; z > 0; z-- {
		for x := uint(0); x <= maxX; x++ {
			output += string(getXDisplay(x, z))
		}
		output += ` `
		output += strconv.Itoa(int(z))
		output += "\n"
	}
	for x := uint(0); x <= maxX; x++ {
		output += `-`
	}
	output += ` 0`
	output += "\n"

	output += "\n"
	output += ` y ` + "\n"
	for y := uint(0); y <= maxY; y++ {
		output += string('0' + byte(y%10))
	}
	output += "\n"

	getYDisplay := func(y, z uint) byte {
		index := getIndexIntersectingY(y, z)
		if index == -1 {
			return '.'
		}
		if index == -2 {
			return '?'
		}
		return toChar(index)
	}

	for z := maxZ; z > 0; z-- {
		for y := uint(0); y <= maxY; y++ {
			output += string(getYDisplay(y, z))
		}
		output += ` `
		output += strconv.Itoa(int(z))
		output += "\n"
	}
	for y := uint(0); y <= maxY; y++ {
		output += `-`
	}
	output += ` 0`
	output += "\n"

	return output
}

func toChar(i int) byte {
	i %= (10 + 26 + 26)
	if i < 10 {
		return '0' + byte(i)
	}
	i -= 10
	if i < 26 {
		return 'a' + byte(i)
	}
	i -= 26
	return 'A' + byte(i)
}

func assertOrderedBlocks(blocks []block) {
	fmt.Printf("blocks[0] = %+v\n", blocks[0])
	for i := 1; i < len(blocks); i++ {
		fmt.Printf("blocks[%d] = %+v\n", i, blocks[i])
		assertOrdered(blocks[i-1], blocks[i])
	}
}

func assertOrdered(i, j block) {
	if i.minZ() > j.minZ() {
		panic(`problem`)
	}
	// if i.maxZ() > j.maxZ() {
	// 	panic(`problem`)
	// }
}
