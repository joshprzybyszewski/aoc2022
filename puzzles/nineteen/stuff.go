package nineteen

import "fmt"

type raw struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type stuff struct {
	robots raw
	bank   raw
}

func newInitialStuff() stuff {
	return stuff{
		robots: raw{
			ore: 1,
		},
	}
}

func (s stuff) String() string {
	output := "Stuff:\n"
	output += "\tBank:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", s.bank.ore)
	output += fmt.Sprintf("\t\tClay:     %2d\n", s.bank.clay)
	output += fmt.Sprintf("\t\tObsidian: %2d\n", s.bank.obsidian)
	output += fmt.Sprintf("\t\tGeode:    %2d\n", s.bank.geode)
	output += "\tRobots:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", s.robots.ore)
	output += fmt.Sprintf("\t\tClay:     %2d\n", s.robots.clay)
	output += fmt.Sprintf("\t\tObsidian: %2d\n", s.robots.obsidian)
	output += fmt.Sprintf("\t\tGeode:    %2d\n", s.robots.geode)

	return output
}

func elapse(
	s *stuff,
	minutes int,
) {
	if minutes < 0 {
		panic(`josh come on`)
	}
	if minutes == 0 {
		// :confusion-intensifies:
		return
	}

	s.bank.ore += (minutes * s.robots.ore)
	s.bank.clay += (minutes * s.robots.clay)
	s.bank.obsidian += (minutes * s.robots.obsidian)
	s.bank.geode += (minutes * s.robots.geode)
}

func pay(
	bank *raw,
	cost raw,
) bool {
	if cost.ore > bank.ore ||
		cost.clay > bank.clay ||
		cost.obsidian > bank.obsidian {
		return false
	}
	bank.ore -= cost.ore
	bank.clay -= cost.clay
	bank.obsidian -= cost.obsidian
	return true
}
