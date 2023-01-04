package nineteen

type optimizer struct {
	// seen map[resources]resources
}

func newOptimizer() *optimizer {
	return &optimizer{
		// seen: make(map[resources]resources, 256),
	}
}

func (o *optimizer) optimize(
	r resources,
	b *blueprint,
) resources {
	if r.remainingTime <= 0 {
		return r
	}

	best := r
	checkBest := func(o resources) {
		if o.geode > best.geode {
			best = o
		}
	}

	tmp, ok := r.buildGeodeRobot(b)
	if ok {
		checkBest(o.optimize(tmp, b))
	}

	tmp, ok = r.buildObsidianRobot(b)
	if ok {
		checkBest(o.optimize(tmp, b))
	}

	tmp, ok = r.buildClayRobot(b)
	if ok {
		checkBest(o.optimize(tmp, b))
	}

	tmp, ok = r.buildOreRobot(b)
	if ok {
		checkBest(o.optimize(tmp, b))
	}

	generate(&best, best.remainingTime)

	return best
}
