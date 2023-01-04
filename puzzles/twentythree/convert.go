package twentythree

func convertInputToElfLocations(input string) map[coord]bool {
	// there are 74 rows and cols. It's probably about half full of elves.
	// so they're likely to expand out (9/2)x in every direction
	output := make(map[coord]bool, 74*74*5)

	var x, y int16
	for _, ch := range input {
		switch ch {
		case '#':
			output[coord{
				x: x,
				y: y,
			}] = true
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

func populateSlice(
	m map[coord]bool,
	s []coord,
) {
	si := 0
	for c := range m {
		s[si] = c
		si++
	}
}
