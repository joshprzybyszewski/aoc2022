package sixteen

import (
	"strconv"
	"strings"
)

const (
	numValves = 66
)

type valves [numValves]valve

type valve struct {
	name  string
	dests []string
	flow  int
}

func newValve(
	name string,
	flow int,
) valve {
	return valve{
		name: name,
		flow: flow,
	}
}

func getValves(
	input string,
) (valves, error) {
	// All valves have two char names
	// Valve MO has flow rate=0; tunnels lead to valves QM, ED

	var output valves //:= make([]*valve, 0, 66)
	oi := 0

	var name string
	var i1, i2, flow int
	var v valve
	var err error
	for nli := strings.Index(input, "\n"); ; nli = strings.Index(input, "\n") {
		if nli < 0 && oi == numValves-1 {
			nli = len(input)
		}
		name = input[6:8]

		i1 = strings.Index(input, `=`) + 1
		i2 = i1 + strings.Index(input[i1:], `;`)
		flow, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return valves{}, err
		}

		v = newValve(name, flow)

		i1 = strings.Index(input, `to valve`) + 8
		i1 += strings.Index(input[i1:], ` `) + 1
		for i1 < len(input) {
			i2 = i1 + strings.Index(input[i1:], `,`)
			if i2 < i1 || i2 > nli {
				v.dests = append(v.dests, input[i1:nli])
				break
			}
			v.dests = append(v.dests, input[i1:i2])
			i1 = i2 + 2
		}

		output[oi] = v
		oi++
		if oi == numValves {
			break
		}
		input = input[nli+1:]
	}

	// sort.Slice(output, func(i, j int) bool {
	// 	return output[i].flow > output[j].flow
	// })

	return output, nil
}
