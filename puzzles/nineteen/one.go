package nineteen

const (
	part1Minutes = 24
)

func One(
	input string,
) (int, error) {
	all, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := range all {
		total += ((i + 1) * harvest(all[i], part1Minutes))
	}

	return 0, nil
}

func harvest(
	b blueprint,
	minutes int,
) int {
	return 0
}
