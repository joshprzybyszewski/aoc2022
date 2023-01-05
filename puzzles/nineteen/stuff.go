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

func maximizeGeodes(
	b *blueprint,
	s stuff,
	remainingMinutes int,
) stuff {
	best := s
	elapse(&best, remainingMinutes)

	if remainingMinutes <= 1 {
		// if there's one minute remaining, it's not worth building a robot
		// if there's fewer than 1 minute remaining, elapse handles that.
		return best
	}

	other, ok := buildGeodeRobot(b, s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = buildObsidianRobot(b, s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = buildClayRobot(b, s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = buildOreRobot(b, s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	return best
}

func buildGeodeRobot(
	b *blueprint,
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.obsidian == 0 {
		// we'll never be able to because we cannot generate obsidian.
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, b.geodeRobotCost) {
		idleMinutes++

		elapse(&s, 1)
	}

	// it takes one minute to build the robot
	elapse(&s, 1)
	s.robots.geode++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return stuff{}, false
	}

	return maximizeGeodes(b, s, remainingMinutes-idleMinutes-1), true
}

func buildObsidianRobot(
	b *blueprint,
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.clay == 0 {
		// we'll never be able to because we cannot generate clay.
		return stuff{}, false
	} else if s.robots.obsidian >= b.geodeRobotCost.obsidian {
		// we shouldn't build an obsidian robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, b.obsidianRobotCost) {
		idleMinutes++

		elapse(&s, 1)
	}

	// it takes one minute to build the robot
	elapse(&s, 1)
	s.robots.obsidian++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return stuff{}, false
	}

	return maximizeGeodes(b, s, remainingMinutes-idleMinutes-1), true
}

func buildClayRobot(
	b *blueprint,
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.clay >= b.obsidianRobotCost.clay {
		// we shouldn't build a clay robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, b.clayRobotCost) {
		idleMinutes++

		elapse(&s, 1)
	}

	// it takes one minute to build the robot
	elapse(&s, 1)
	s.robots.clay++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return stuff{}, false
	}

	return maximizeGeodes(b, s, remainingMinutes-idleMinutes-1), true
}

func buildOreRobot(
	b *blueprint,
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.ore >= b.oreRobotCost.ore &&
		s.robots.ore >= b.clayRobotCost.ore &&
		s.robots.ore >= b.obsidianRobotCost.ore &&
		s.robots.ore >= b.geodeRobotCost.ore {
		// we shouldn't build an ore robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, b.oreRobotCost) {
		idleMinutes++

		elapse(&s, 1)
	}

	// it takes one minute to build the robot
	elapse(&s, 1)
	s.robots.ore++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return stuff{}, false
	}

	return maximizeGeodes(b, s, remainingMinutes-idleMinutes-1), true
}
