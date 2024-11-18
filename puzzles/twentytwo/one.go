package twentytwo

func One(
	input string,
) (int, error) {

	blocks := convertInput(input)

	return len(blocks), nil
}
