package fourteen

func Two(
	input string,
) (int, error) {
	g, err := getGrid(input)
	if err != nil {
		return 0, err
	}

	g.addFloor()

	units := 0
	for g.addSand(500, 0) {
		units++
		if g.check(500, 0) == sand {
			break
		}
	}

	return units, nil
}
