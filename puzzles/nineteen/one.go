package nineteen

const (
	part1Minutes = 24
)

const (
	testInput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
`
)

func One(
	input string,
) (int, error) {
	// input = testInput
	all, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := range all {
		if input == testInput && i > 1 {
			break
		}
		total += ((i + 1) * harvest(&all[i], part1Minutes))
	}

	// 995 is too low
	return total, nil
}

func harvest(
	b *blueprint,
	minutes int,
) int {
	s := newInitialStuff()

	s = maximizeGeodes(
		b,
		s,
		minutes,
	)
	return s.bank.geode
}
