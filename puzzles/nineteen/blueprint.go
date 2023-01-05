package nineteen

import (
	"strconv"
	"strings"
)

const (
	numBlueprints = 30 // number of lines in input
)

type allBlueprints [numBlueprints]blueprint

type blueprint struct {
	oreRobotCost      raw
	clayRobotCost     raw
	obsidianRobotCost raw
	geodeRobotCost    raw
}

func getBlueprints(
	input string,
) (allBlueprints, error) {
	var alli int
	var all allBlueprints

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
			return allBlueprints{}, err
		}
		all[alli].oreRobotCost.ore = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].clayRobotCost.ore = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].obsidianRobotCost.ore = tmp

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].obsidianRobotCost.clay = tmp

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].geodeRobotCost.ore = tmp

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].geodeRobotCost.obsidian = tmp

		input = input[nli+1:]
		alli++
	}

	return all, nil

}
