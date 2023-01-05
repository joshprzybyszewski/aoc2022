package nineteen

const (
	maxMinutes = part2Minutes + 1
)

var (
	// this only has 13 entries instead of `maxMinutes` because the 12th minute
	// overflows uint8.
	possibleProductionByTimeRemaining [13]uint8
)

func init() {
	var sum uint8
	for i := 0; i < len(possibleProductionByTimeRemaining); i++ {
		possibleProductionByTimeRemaining[i] = sum
		for j := 1; j < i; j++ {
			sum += uint8(j)
		}
	}
}

type maximizer struct {
	b *blueprint

	knownBest stuff
}

func newMaximizer(
	b *blueprint,
) maximizer {
	m := maximizer{
		b: b,
	}

	return m
}

func (m *maximizer) setKnownBest(
	s stuff,
) {
	if s.bank.geode > m.knownBest.bank.geode {
		m.knownBest = s
	}
}

// canBeatKnownBest returns true if, by building robots, the given stuff could be the known best.
func (m *maximizer) canBeatKnownBest(
	now, terminal stuff,
	remainingMinutes uint8,
) bool {
	if remainingMinutes <= 1 {
		// if there's one or fewer minute(s) remaining, then no robot is worth building.
		return false
	}
	if int(remainingMinutes) >= len(possibleProductionByTimeRemaining) {
		// can't know yet.
		return true
	}

	// In theory, we could build a geode robot every minute for the remaining minutes
	// and that would be the most geodes we could produce.
	possibleProduction := possibleProductionByTimeRemaining[remainingMinutes]
	if terminal.bank.geode+possibleProduction < m.knownBest.bank.geode {
		return false
	}

	return true
}

func (m *maximizer) maximizeGeodes(
	s stuff,
	remainingMinutes uint8,
	// TODO use DP to keep track of seen states?
) {

	best := s
	elapseN(&best, remainingMinutes)
	m.setKnownBest(best)

	if !m.canBeatKnownBest(s, best, remainingMinutes) {
		return
	}

	m.buildGeodeRobot(s, remainingMinutes)
	m.buildObsidianRobot(s, remainingMinutes)
	m.buildClayRobot(s, remainingMinutes)
	m.buildOreRobot(s, remainingMinutes)

}

func (m *maximizer) buildGeodeRobot(
	s stuff,
	remainingMinutes uint8,
) {
	if s.robots.obsidian == 0 {
		// we'll never be able to because we cannot generate obsidian.
		return
	}

	var idleMinutes uint8
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.geodeRobotCost) {
		idleMinutes++

		elapse(&s)
	}

	// it takes one minute to build the robot
	elapse(&s)
	s.robots.geode++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return
	}

	m.maximizeGeodes(s, remainingMinutes-idleMinutes-1)
}

func (m *maximizer) buildObsidianRobot(
	s stuff,
	remainingMinutes uint8,
) {
	if s.robots.clay == 0 {
		// we'll never be able to because we cannot generate clay.
		return
	} else if s.robots.obsidian >= m.b.geodeRobotCost.obsidian {
		// we shouldn't build an obsidian robot because we already have enough
		return
	}

	var idleMinutes uint8
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.obsidianRobotCost) {
		idleMinutes++

		elapse(&s)
	}

	// it takes one minute to build the robot
	elapse(&s)
	s.robots.obsidian++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return
	}

	m.maximizeGeodes(s, remainingMinutes-idleMinutes-1)
}

func (m *maximizer) buildClayRobot(
	s stuff,
	remainingMinutes uint8,
) {
	if s.robots.clay >= m.b.obsidianRobotCost.clay {
		// we shouldn't build a clay robot because we already have enough
		return
	}

	var idleMinutes uint8
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.clayRobotCost) {
		idleMinutes++

		elapse(&s)
	}

	// it takes one minute to build the robot
	elapse(&s)
	s.robots.clay++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return
	}

	m.maximizeGeodes(s, remainingMinutes-idleMinutes-1)
}

func (m *maximizer) buildOreRobot(
	s stuff,
	remainingMinutes uint8,
) {
	if s.robots.ore >= m.b.oreRobotCost.ore &&
		s.robots.ore >= m.b.clayRobotCost.ore &&
		s.robots.ore >= m.b.obsidianRobotCost.ore &&
		s.robots.ore >= m.b.geodeRobotCost.ore {
		// we shouldn't build an ore robot because we already have enough
		return
	}

	var idleMinutes uint8
	for idleMinutes < remainingMinutes && !pay(&s.bank, m.b.oreRobotCost) {
		idleMinutes++

		elapse(&s)
	}

	// it takes one minute to build the robot
	elapse(&s)
	s.robots.ore++

	if idleMinutes+1 >= remainingMinutes {
		// could not generate then build the robot in time to use it.
		return
	}

	m.maximizeGeodes(s, remainingMinutes-idleMinutes-1)
}
