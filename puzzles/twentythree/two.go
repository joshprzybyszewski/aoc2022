package twentythree

func Two(
	input string,
) (int, error) {

	elves := convertInputToElfLocations(input)
	numElves := len(elves)
	steady := false
	for r := 0; ; r++ {
		steady = updateMap(elves, numElves, r)
		if steady {
			return r + 1, nil
		}
	}
	// return 0, nil
}
