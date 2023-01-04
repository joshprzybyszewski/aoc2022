package sixteen

import (
	"strconv"
	"strings"
)

type valve struct {
	name  string
	dests []string
	flow  int
}

func newValve(
	name string,
	flow int,
) *valve {
	return &valve{
		name: name,
		flow: flow,
	}
}

func getValves(
	input string,
) ([]*valve, error) {
	// All valves have two char names
	// Valve MO has flow rate=0; tunnels lead to valves QM, ED
	lines := strings.Split(input, "\n")

	output := make([]*valve, 0, len(lines)-1)

	var name string
	var i1, i2, flow int
	var v *valve
	var err error
	for _, line := range lines {
		if line == `` {
			continue
		}
		name = line[6:8]

		i1 = strings.Index(line, `=`) + 1
		i2 = i1 + strings.Index(line[i1:], `;`)
		flow, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return nil, err
		}

		v = newValve(name, flow)

		i1 = strings.Index(line, `to valve`) + 8
		i1 += strings.Index(line[i1:], ` `) + 1
		for i1 < len(line) {
			i2 = i1 + strings.Index(line[i1:], `,`)
			if i2 < i1 {
				v.dests = append(v.dests, line[i1:])
				break
			}
			v.dests = append(v.dests, line[i1:i2])
			i1 = i2 + 2
		}

		output = append(output, v)
	}

	return output, nil
}
