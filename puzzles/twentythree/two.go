package twentythree

func Two(
	input string,
) (int, error) {

	elves := convertInputToElfLocations(input)
	es := make([]coord, len(elves))
	populateSlice(elves, es)
	steady := false
	var ri uint8
	for r := 0; ; r++ {
		steady = updateMap(elves, es, ri)
		if steady {
			return r + 1, nil
		}
		ri++
		ri &= 3
	}
	// return 0, nil
}
