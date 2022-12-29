package twentytwo

func Two(
	input string,
) (int, error) {
	start, dirs := convertInputToCube(input)
	s := start
	for _, d := range dirs {
		s = moveInCube(s, d)
	}

	return (1000 * int(s.row)) +
		(4 * int(s.col)) +
		int(dirs[len(dirs)-1].heading), nil
}
