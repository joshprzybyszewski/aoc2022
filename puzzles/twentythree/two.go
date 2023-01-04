package twentythree

func Two(
	input string,
) (int, error) {

	elves := convertInputToElfLocations(input)
	steady := false
	for r := 0; ; r++ {
		elves, steady = getNextPositions(elves, r)
		if steady {
			return r + 1, nil
		}
	}
}
