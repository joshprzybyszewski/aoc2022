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

func (b blueprint) minimumTimeRequiredForObsidian() int {
	total := 0
	for i := 1; ; i++ {
		total += i
		if total > b.geodeRobotObs {
			return i
		}
	}
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

	remainingTime int
}

func newResources(remainingTime int) resources {
	return resources{
		oreRobots:     1,
		remainingTime: remainingTime,
	}
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

func (r resources) maxUsableObsidian() int {
	return r.obs + (2*((r.remainingTime-1)/2)+1)*((r.remainingTime)/2)
}

func (r resources) maxUsableClay(dur int) int {
	return r.clay + (2*(dur/2)+1)*((dur+1)/2)
}

func (r resources) cannotGenerateGeode(b blueprint) bool {
	if r.geodeRobots > 0 {
		return false
	}
	if r.remainingTime < b.geodeRobotObs && r.maxUsableObsidian() < b.geodeRobotObs {
		// we cannot possibly generate enough obsidian to create even one geode robot
		return true
	}
	if r.remainingTime < b.geodeRobotObs+b.obsRobotClay {
		clayDur := r.remainingTime - 1 - b.minimumTimeRequiredForObsidian()
		if r.maxUsableClay(clayDur) < b.obsRobotClay {
			// we cannot possibly generate enough clay to create
			// even one obsidian robot which means we cannot generate a 	geode robot
			return true
		}
	}
	return false
}

func (r resources) canBuildObsidianRobot(b blueprint) bool {
	return r.ore >= b.obsRobotOre && r.clay >= b.obsRobotClay
}

func (r resources) shouldBuildObsidianRobot(b blueprint) bool {
	if r.obsRobots >= b.geodeRobotObs {
		// we'll generate more obsidian in one minute than we'll ever possibly use
		return false
	}
	// curObsTrajectory := r.obs + (r.obsRobots * r.remainingTime)
	// without := curObsTrajectory / b.geodeRobotObs
	// with := (curObsTrajectory + r.remainingTime) / b.geodeRobotObs
	// if with == 0 || with <= without {
	// 	// building this obsidian robot won't enable us to build another geode robot
	// 	// so it's not worth it.
	// 	return false
	// }
	return true
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

func (r resources) shouldBuildClayRobot(b blueprint) bool {
	if r.clayRobots >= b.obsRobotClay {
		// we'll generate more clay in one minute than we'll ever possibly use
		return false
	}
	// curClayTrajectory := r.clay + (r.clayRobots * r.remainingTime)
	// without := curClayTrajectory / b.obsRobotClay
	// with := (curClayTrajectory + r.remainingTime) / b.obsRobotClay

	// if with == 0 || with <= without {
	// 	// building this clay robot won't enable us to build another obsidian robot
	// 	// so it's not worth it.
	// 	return false
	// }

	return true
}

func (r resources) buildClayRobot(b blueprint) resources {
	r.ore -= b.clayRobot
	r.clayRobots++
	return r
}

func (r resources) canBuildOreRobot(b blueprint) bool {
	return r.ore >= b.oreRobot
}

func (r resources) shouldBuildOreRobot(b blueprint) bool {
	if r.oreRobots >= b.geodeRobotOre &&
		r.oreRobots >= b.obsRobotOre &&
		r.oreRobots >= b.clayRobot &&
		r.oreRobots >= b.oreRobot {
		// we'll generate more ore in one minute than we'll ever possibly use
		return false
	}
	// if b.oreRobot > r.remainingTime {
	// 	// this ore robot won't pay for itself.
	// 	return false
	// }
	return true
}

func (r resources) buildOreRobot(b blueprint) resources {
	r.ore -= b.oreRobot
	r.oreRobots++
	return r
}

func incrementTime(
	r resources,
) resources {
	r.remainingTime--
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

	bs[0] = blueprint{
		oreRobot:      4,
		clayRobot:     2,
		obsRobotOre:   3,
		obsRobotClay:  14,
		geodeRobotOre: 2,
		geodeRobotObs: 7,
	}
	bs[1] = blueprint{
		oreRobot:      2,
		clayRobot:     3,
		obsRobotOre:   3,
		obsRobotClay:  8,
		geodeRobotOre: 3,
		geodeRobotObs: 12,
	}

	total := 0
	for i := range bs {
		total += ((i + 1) * getMostGeodesIn24Minutes(bs[i]))
		if i == 1 {
			// TODO change this
			break
		}
	}

	return total, nil
}

func getMostGeodesIn24Minutes(
	b blueprint,
) int {
	r := newResources(24)
	best := optimizeForGeodes(
		b,
		r,
	)
	fmt.Printf("Most geodes:\n%+v\n\t%+v\n", b, best)
	return best.geode
}

func optimizeForGeodes(
	b blueprint,
	r resources,
) resources {
	if r.cannotGenerateGeode(b) {
		fmt.Printf("stopping iteration\n\t%+v\n\t%+v\n\n", r, b)
		return r
	}

	r = incrementTime(r)
	// fmt.Printf("\nRemaining Time: %d\n%+v\n%+v\n\n", r.remainingTime, b, r)
	if r.remainingTime == 0 {
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

	if r.canBuildObsidianRobot(b) && r.shouldBuildObsidianRobot((b)) {
		// there's enough ore and clay to build an obsidian robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildObsidianRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	if r.canBuildClayRobot(b) && r.shouldBuildClayRobot(b) {
		// there's enough ore to build a clay robot. Build it and then optimize with it.
		other := optimizeForGeodes(b, r.buildClayRobot(b))
		if other.geode >= best.geode {
			best = other
		}
	}

	if r.canBuildOreRobot(b) && r.shouldBuildOreRobot(b) {
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
