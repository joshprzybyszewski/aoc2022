package twentytwo

func One(
	input string,
) (int, error) {

	start, dirs := convertInput(input)
	s := start
	for _, d := range dirs {
		s = move(s, d)
	}

	return (1000 * int(s.row)) +
		(4 * int(s.col)) +
		int(dirs[len(dirs)-1].heading), nil
}
