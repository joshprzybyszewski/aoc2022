package nineteen

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxTime = 24
)

type blueprint struct {
	oreRobot      int
	clayRobot     int
	obsRobotOre   int
	obsRobotClay  int
	geodeRobotOre int
	geodeRobotObs int
}

func (r resources) canBuildGeodeRobot(b blueprint) bool {
	return r.ore >= b.geodeRobotOre && r.obs >= b.geodeRobotObs
}

func (r resources) buildGeodeRobot(b blueprint) resources {
	r.ore -= b.geodeRobotOre
	r.obs -= b.geodeRobotObs
	r.geodeRobots++
	return r
}

func (r resources) canBuildObsidianRobot(b blueprint) bool {
	return r.ore >= b.obsRobotOre && r.clay >= b.obsRobotClay
}

func (r resources) buildObsidianRobot(b blueprint) resources {
	r.ore -= b.obsRobotOre
	r.clay -= b.obsRobotClay
	r.obsRobots++
	return r
}

func (r resources) canBuildClayRobot(b blueprint) bool {
	return r.ore >= b.clayRobot
}

func (r resources) buildClayRobot(b blueprint) resources {
	r.ore -= b.clayRobot
	r.clayRobots++
	return r
}

func (r resources) canBuildOreRobot(b blueprint) bool {
	return r.ore >= b.oreRobot
}

func (r resources) buildOreRobot(b blueprint) resources {
	r.ore -= b.oreRobot
	r.oreRobots++
	return r
}

type resources struct {
	ore       int
	oreRobots int

	clay       int
	clayRobots int

	obs       int
	obsRobots int

	geode       int
	geodeRobots int

	time int
}

func incrementTime(
	r resources,
) resources {
	r.time++
	r.ore += r.oreRobots
	r.clay += r.clayRobots
	r.obs += r.obsRobots
	r.geode += r.geodeRobots
	return r
}

func One(
	input string,
) (int, error) {
	bs, err := getBlueprints(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := range bs {
		total += ((i + 1) * getMostGeodesIn24Minutes(bs[i]))
		if i == 1 {
			break
		}
	}

	return total, nil
}

func getMostGeodesIn24Minutes(
	b blueprint,
) int {
	r := resources{
		oreRobots: 1,
	}
	best := optimizeForGeodes(
		b,
		r,
	)
	return best.geode
}

func optimizeForGeodes(
	b blueprint,
	r resources,
) resources {
	r = incrementTime(r)
	fmt.Printf("\nTime: %d\n%+v\n%+v\n\n", r.time, b, r)
	if r.time == maxTime {
		// no time left. the input is the best answer.
		return r
	}

	// don't use any resources this generation.
	best := optimizeForGeodes(b, r)

	if r.canBuildGeodeRobot(b) {
		// there's enough ore and obsidian to build a geode robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildGeodeRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	if r.canBuildObsidianRobot(b) {
		// there's enough ore and clay to build an obsidian robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildObsidianRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	if r.canBuildClayRobot(b) {
		// there's enough ore to build a clay robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildClayRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	if r.canBuildOreRobot(b) {
		// there's enough ore to build an ore robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildOreRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	return best
}

func getBlueprints(
	input string,
) ([30]blueprint, error) {
	var bi int
	var bs [30]blueprint

	var i1, i2 int
	var tmp int
	var err error

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		i1 = strings.Index(input, `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].oreRobot = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].clayRobot = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].obsRobotOre = tmp

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].obsRobotClay = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].geodeRobotOre = tmp

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [30]blueprint{}, err
		}
		bs[bi].geodeRobotObs = tmp

		input = input[nli+1:]
		bi++
	}

	return bs, nil

}
