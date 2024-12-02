package sixteen

func Two(
	input string,
) (int, error) {
	c := newContraption(input)
	best, tmp := -1, 0
	for i := uint8(0); i < size; i++ {
		c.energized = [size][size]lightTile{}
		c.sendLight(coord{
			row: i,
			col: 0,
		}, right)
		if tmp = c.numEnergized(); tmp > best {
			best = tmp
		}

		c.energized = [size][size]lightTile{}
		c.sendLight(coord{
			row: 0,
			col: i,
		}, down)
		if tmp = c.numEnergized(); tmp > best {
			best = tmp
		}

		c.energized = [size][size]lightTile{}
		c.sendLight(coord{
			row: i,
			col: size - 1,
		}, left)
		if tmp = c.numEnergized(); tmp > best {
			best = tmp
		}

		c.energized = [size][size]lightTile{}
		c.sendLight(coord{
			row: size - 1,
			col: i,
		}, up)
		if tmp = c.numEnergized(); tmp > best {
			best = tmp
		}
	}

	return best, nil

}
