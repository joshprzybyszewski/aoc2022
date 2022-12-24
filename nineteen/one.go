package nineteen

import (
	"fmt"
	"strconv"
	"strings"
)

type blueprint struct {
	oreRobot      int
	clayRobot     int
	obsRobotOre   int
	obsRobotClay  int
	geodeRobotOre int
	geodeRobotObs int
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
}

func incrementTime(
	r resources,
) resources {
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
		24,
	)
	return best.geode
}

func checkFeasibility(
	b blueprint,
) bool {
	if 
}

func optimizeForGeodes(
	b blueprint,
	r resources,
	timeRemaining int,
) resources {
	r = incrementTime(r)
	fmt.Printf("\nTime: %d\n%+v\n%+v\n\n", timeRemaining, b, r)
	if timeRemaining == 0 {
		// no time left. the input is the best answer.
		return r
	}

	// don't use any resources this generation.
	best := optimizeForGeodes(b, r, timeRemaining-1)

	if r.ore >= b.geodeRobotOre && r.obs >= b.geodeRobotObs {
		// there's enough ore and obsidian to build a geode robot. Build it and then optimize with it.
		r := r
		r.ore -= b.geodeRobotOre
		r.obs -= b.geodeRobotObs
		r.geodeRobots++
		r = optimizeForGeodes(b, r, timeRemaining-1)
		if r.geode >= best.geode {
			best = r
		}
	}

	if r.ore >= b.obsRobotOre && r.clay >= b.obsRobotClay {
		// there's enough ore and clay to build an obsidian robot. Build it and then optimize with it.
		r := r
		r.ore -= b.obsRobotOre
		r.clay -= b.obsRobotClay
		r.obsRobots++
		r = optimizeForGeodes(b, r, timeRemaining-1)
		if r.geode >= best.geode {
			best = r
		}
	}

	if r.ore >= b.clayRobot {
		// there's enough ore to build a clay robot. Build it and then optimize with it.
		r := r
		r.ore -= b.clayRobot
		r.clayRobots++
		r = optimizeForGeodes(b, r, timeRemaining-1)
		if r.geode >= best.geode {
			best = r
		}
	}

	if r.ore >= b.oreRobot {
		// there's enough ore to build an ore robot. Build it and then optimize with it.
		r := r
		r.ore -= b.oreRobot
		r.oreRobots++
		r = optimizeForGeodes(b, r, timeRemaining-1)
		if r.geode >= best.geode {
			best = r
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
