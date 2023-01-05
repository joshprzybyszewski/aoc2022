package nineteen

import "fmt"

type raw struct {
	ore      uint8
	clay     uint8
	obsidian uint8
	geode    uint8
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
) {
	s.bank.ore += s.robots.ore
	s.bank.clay += s.robots.clay
	s.bank.obsidian += s.robots.obsidian
	s.bank.geode += s.robots.geode
}

func elapseN(
	s *stuff,
	remainingMinutes uint8,
) {
	s.bank.ore += (remainingMinutes * s.robots.ore)
	s.bank.clay += (remainingMinutes * s.robots.clay)
	s.bank.obsidian += (remainingMinutes * s.robots.obsidian)
	s.bank.geode += (remainingMinutes * s.robots.geode)
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
