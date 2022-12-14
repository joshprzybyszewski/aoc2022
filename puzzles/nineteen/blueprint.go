package nineteen

import (
	"fmt"
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

func (b *blueprint) String() string {
	output := "Costs:\n"
	output += "\tOre Robot:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", b.oreRobotCost.ore)
	output += "\tClay Robot:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", b.clayRobotCost.ore)
	output += "\tObsidian Robot:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", b.obsidianRobotCost.ore)
	output += fmt.Sprintf("\t\tClay:     %2d\n", b.obsidianRobotCost.clay)
	output += "\tGeode Robot:\n"
	output += fmt.Sprintf("\t\tOre:      %2d\n", b.geodeRobotCost.ore)
	output += fmt.Sprintf("\t\tObsidian: %2d\n", b.geodeRobotCost.obsidian)

	return output
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
		all[alli].oreRobotCost.ore = uint8(tmp)

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].clayRobotCost.ore = uint8(tmp)

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].obsidianRobotCost.ore = uint8(tmp)

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].obsidianRobotCost.clay = uint8(tmp)

		i1 = i2 + strings.Index(input[i2:], `costs `) + 6
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].geodeRobotCost.ore = uint8(tmp)

		i1 = i2 + strings.Index(input[i2:], `and `) + 4
		i2 = i1 + strings.Index(input[i1:], ` `)
		tmp, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return allBlueprints{}, err
		}
		all[alli].geodeRobotCost.obsidian = uint8(tmp)

		input = input[nli+1:]
		alli++
	}

	return all, nil

}
