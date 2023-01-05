package nineteen

const (
	maxMinutes = part2Minutes
)

type maximizer struct {
	b *blueprint

	// [minutesRemaining][ore][clay][obsidian][geode]
	seen [maxMinutes]map[stuff]struct{}
}

func newMaximizer(
	b *blueprint,
) maximizer {
	m := maximizer{
		b: b,
	}

	for i := range m.seen {
		m.seen[i] = make(map[stuff]struct{})
	}

	return m
}

func (m *maximizer) hasSeen(
	s stuff,
	remainingMinutes int,
) bool {
	_, ok := m.seen[remainingMinutes][s]
	if ok {
		return true
	}
	m.seen[remainingMinutes][s] = struct{}{}
	return false
}

func (m *maximizer) maximizeGeodes(
	s stuff,
	remainingMinutes int,
	// TODO use DP to keep track of seen states?
) stuff {
	best := s
	elapse(&best, remainingMinutes)

	if remainingMinutes <= 1 {
		// if there's one minute remaining, it's not worth building a robot
		// if there's fewer than 1 minute remaining, elapse handles that.
		return best
	}

	other, ok := m.buildGeodeRobot(s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = m.buildObsidianRobot(s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = m.buildClayRobot(s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	other, ok = m.buildOreRobot(s, remainingMinutes)
	if ok && other.bank.geode > best.bank.geode {
		best = other
	}

	return best
}

func (m *maximizer) buildGeodeRobot(
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.obsidian == 0 {
		// we'll never be able to because we cannot generate obsidian.
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.geodeRobotCost) {
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

	return m.maximizeGeodes(s, remainingMinutes-idleMinutes-1), true
}

func (m *maximizer) buildObsidianRobot(
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.clay == 0 {
		// we'll never be able to because we cannot generate clay.
		return stuff{}, false
	} else if s.robots.obsidian >= m.b.geodeRobotCost.obsidian {
		// we shouldn't build an obsidian robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.obsidianRobotCost) {
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

	return m.maximizeGeodes(s, remainingMinutes-idleMinutes-1), true
}

func (m *maximizer) buildClayRobot(
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.clay >= m.b.obsidianRobotCost.clay {
		// we shouldn't build a clay robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.clayRobotCost) {
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

	return m.maximizeGeodes(s, remainingMinutes-idleMinutes-1), true
}

func (m *maximizer) buildOreRobot(
	s stuff,
	remainingMinutes int,
) (stuff, bool) {
	if s.robots.ore >= m.b.oreRobotCost.ore &&
		s.robots.ore >= m.b.clayRobotCost.ore &&
		s.robots.ore >= m.b.obsidianRobotCost.ore &&
		s.robots.ore >= m.b.geodeRobotCost.ore {
		// we shouldn't build an ore robot because we already have enough
		return stuff{}, false
	}

	idleMinutes := 0
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.oreRobotCost) {
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

	return m.maximizeGeodes(s, remainingMinutes-idleMinutes-1), true
}
