package nineteen

type blueprint struct {
	// ore cost to construct an ore robot
	oreRobot int

	// ore cost to construct a clay robot
	clayRobot int

	// the ore and clay cost to construct an obsidian robot
	obsRobotOre  int
	obsRobotClay int

	// the ore and obsidian cost to construct a geode robot
	geodeRobotOre int
	geodeRobotObs int
}

func (b *blueprint) maxOreNeed() int {
	// TODO store this as a field
	return max4(
		b.geodeRobotOre,
		b.obsRobotOre,
		b.clayRobot,
		b.oreRobot,
	)
}
