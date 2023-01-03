package nineteen

import (
	"fmt"
)

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
		total += ((i + 1) * getMostGeodesIn24Minutes(&bs[i]))
		if i == 1 {
			// TODO change this
			break
		}
	}

	return total, nil
}

func getMostGeodesIn24Minutes(
	b *blueprint,
) int {
	r := newResources(24)
	o := newOptimizer()
	best := o.optimize(
		r,
		b,
	)
	fmt.Printf("Most geodes: %d\nBlueprint:\n\t%+v\nResources:\n\t%+v\n\n",
		best.geode,
		b,
		best,
	)
	return best.geode
}
