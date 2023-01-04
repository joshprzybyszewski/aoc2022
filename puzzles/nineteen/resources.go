package nineteen

type resources struct {
	ore       int
	oreRobots int

	clay       int
	clayRobots int

	obs       int
	obsRobots int

	geode       int
	geodeRobots int

	remainingTime int
}

func newResources(remainingTime int) resources {
	return resources{
		oreRobots:     1,
		remainingTime: remainingTime,
	}
}

func (r resources) buildGeodeRobot(b *blueprint) (resources, bool) {
	if r.obsRobots == 0 {
		// we won't be able to generate a geode robot without an obsidian supply
		return resources{}, false
	}

	if b.geodeRobotObs > r.obs+(r.obsRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	if b.geodeRobotOre > r.ore+(r.oreRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	// // TODO I don't think this math is right
	// obsMinutes := (b.geodeRobotObs - r.obs) / r.obsRobots
	// oreMinutes := (b.geodeRobotOre - r.ore) / r.oreRobots
	// generate(&r, max3(obsMinutes, oreMinutes, r.remainingTime))
	for r.obs < b.geodeRobotObs && r.ore < b.geodeRobotOre && r.remainingTime > 0 {
		// TODO fix math to replace this for loop
		generate1(&r)
	}
	if r.remainingTime <= 0 {
		return r, false
	}
	r.obs -= b.geodeRobotObs
	r.ore -= b.geodeRobotOre
	generate1(&r)
	r.geodeRobots++

	return r, true
}

func (r resources) buildObsidianRobot(b *blueprint) (resources, bool) {
	if r.clayRobots == 0 {
		// we won't be able to generate an obsidian robot without a clay supply
		return resources{}, false
	}

	if r.obsRobots >= b.geodeRobotObs {
		// we have more obs robots than we'll ever be able to use
		return resources{}, false
	}

	if b.obsRobotClay > r.clay+(r.clayRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	if b.obsRobotOre > r.ore+(r.oreRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	// TODO I don't think this math is right
	// clayMinutes := (b.obsRobotClay - r.clay) / r.clayRobots
	// oreMinutes := (b.obsRobotOre - r.ore) / r.oreRobots
	// generate(&r, max(clayMinutes, oreMinutes))
	for r.clay < b.obsRobotClay && r.ore < b.obsRobotOre && r.remainingTime > 0 {
		// TODO fix math to replace this for loop
		generate1(&r)
	}
	if r.remainingTime <= 0 {
		return r, false
	}
	r.clay -= b.obsRobotClay
	r.ore -= b.obsRobotOre
	generate1(&r)
	r.obsRobots++

	return r, true
}

func (r resources) buildClayRobot(b *blueprint) (resources, bool) {
	if r.clayRobots >= b.obsRobotClay {
		// we have more clay robots than we'll ever be able to use
		return resources{}, false
	}

	if b.clayRobot > r.ore+(r.oreRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	// TODO I don't think this math is right
	// m := (b.clayRobot - r.ore) / r.oreRobots
	// generate(&r, m)
	for r.ore < b.clayRobot && r.remainingTime > 0 {
		// TODO fix math to replace this for loop
		generate1(&r)
	}
	if r.remainingTime <= 0 {
		return r, false
	}
	r.ore -= b.clayRobot
	generate1(&r)
	r.clayRobots++

	return r, true
}

func (r resources) buildOreRobot(b *blueprint) (resources, bool) {
	if b.oreRobot > r.ore+(r.oreRobots*r.remainingTime) {
		// won't be able to build one in time to use it
		return resources{}, false
	}

	if r.oreRobots >= b.maxOreNeed() {
		// we have more ore robots than we'll ever be able to use
		return resources{}, false
	}

	// TODO I don't think this math is right
	// m := (b.oreRobot - r.ore) / r.oreRobots
	// generate(&r, m)
	for r.ore < b.oreRobot && r.remainingTime > 0 {
		// TODO fix math to replace this for loop
		generate1(&r)
	}
	if r.remainingTime <= 0 {
		return r, false
	}
	r.ore -= b.oreRobot
	generate1(&r)
	r.oreRobots++

	return r, true
}
