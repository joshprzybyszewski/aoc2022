package twentythree

func Two(
	input string,
) (int, error) {

	space, elves := convertInputToElfLocations(input)

	w := newWorkforce(&space, elves)
	w.start()
	defer w.stop()

	steady := false
	var ri uint8
	for r := 0; ; r++ {
		steady = runRound(&w, ri)
		if steady {
			return r + 1, nil
		}
		ri++
		ri &= 3
	}
	// return 0, nil
}
