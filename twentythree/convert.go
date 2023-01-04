package twentythree

func convertInputToElfLocations(input string) []coord {
	// there are 74 rows and cols. It's probably about half full of elves.
	output := make([]coord, 0, 74*74/2)

	x, y := 0, 0
	for _, ch := range input {
		switch ch {
		case '#':
			output = append(output, coord{
				x: x,
				y: y,
			})
			x++
		case '.':
			x++
		case '\n':
			y++
			x = 0
		}
	}

	return output
}
