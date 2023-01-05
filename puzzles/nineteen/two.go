package nineteen

const (
	part2Minutes = 32
)

func Two(
	input string,
) (int, error) {
	all, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	a := harvest(&all[0], part2Minutes)
	b := harvest(&all[1], part2Minutes)
	c := harvest(&all[2], part2Minutes)

	return a * b * c, nil
}
