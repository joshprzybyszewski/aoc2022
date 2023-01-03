package nineteen

func generate(
	r *resources,
	minutes int,
) {
	if minutes <= 0 {
		return
	}
	r.remainingTime -= minutes
	r.ore += (r.oreRobots * minutes)
	r.clay += (r.clayRobots * minutes)
	r.obs += (r.obsRobots * minutes)
	r.geode += (r.geodeRobots * minutes)
}

func generate1(
	r *resources,
) {
	r.remainingTime--
	r.ore += r.oreRobots
	r.clay += r.clayRobots
	r.obs += r.obsRobots
	r.geode += r.geodeRobots
}
