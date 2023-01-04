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
	for g.addSand(0, 500) {
		units++
		if g.check(0, 500) == sand {
			break
		}
	}

	return units, nil
}
