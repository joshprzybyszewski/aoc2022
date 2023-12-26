package twenty

import (
	"fmt"
	"slices"
	"strings"
)

func One(
	input string,
) (int, error) {
	m := newMachine(input)
	fmt.Printf("machine: %+v\n", m)

	// TODO push button on the machine
	return 0, nil
}

func newMachine(input string) machine {
	m := machine{
		flipFlops:    make([]flipFlopModule, 0, 32),
		conjunctions: make([]conjunctionModule, 0, 32),
	}
	var src string
	var dests []string

	cables := make([]cable, 0, 256)

	for len(input) > 0 {
		if len(input) > 5 {
			fmt.Printf("input: %q\n", input[:5])
		}
		if input[0] == '\n' {
			break
		}
		if input[0] == 'b' {
			m.broadcaster.destinations, input = getDestinations(input)
			continue
		}

		mt := input[0]
		input = input[1:]

		src, dests, input = getSourceAndDestinations(input)
		for _, dest := range dests {
			cables = append(cables, cable{
				source:      src,
				destination: dest,
			})
		}

		switch mt {
		case '%':
			m.flipFlops = append(m.flipFlops, newFlipFlopModule(src, dests))
		case '&':
			m.conjunctions = append(m.conjunctions, newConjunctionModule(src, dests))
		default:
			fmt.Printf("mt: %s\n", string(mt))
			panic(`unexpected`)
		}
	}

	slices.SortFunc(cables, func(a, b cable) int {
		d := strings.Compare(a.destination, b.destination)
		if d != 0 {
			return d
		}
		return strings.Compare(a.source, b.source)
	})

	slices.SortFunc(m.flipFlops, func(a, b flipFlopModule) int {
		return strings.Compare(a.name, b.name)
	})

	slices.SortFunc(m.conjunctions, func(a, b conjunctionModule) int {
		return strings.Compare(a.name, b.name)
	})

	for cableIndex, cmi := 0, 0; cableIndex < len(cables) && cmi < len(m.conjunctions); {
		if cables[cableIndex].destination == m.conjunctions[cmi].name {
			m.conjunctions[cmi].addInput(cables[cableIndex].source)
			cableIndex++
			continue
		}
		if cmi < len(m.conjunctions)-1 && cables[cableIndex].destination == m.conjunctions[cmi+1].name {
			cmi++
			continue
		}
		cableIndex++
	}

	return m
}

func getSourceAndDestinations(input string) (string, []string, string) {
	var source string
	for i := 0; ; i++ {
		if input[i] == ' ' {
			source = input[:i]
			input = input[i:]
			break
		}
	}
	dests, input := getDestinations(input)
	return source, dests, input
}

func getDestinations(input string) ([]string, string) {
	dests := make([]string, 0, 5)

	for {
		if input[0] == '-' {
			if input[1] != '>' || input[2] != ' ' {
				panic(`unexpected`)
			}
			input = input[3:]
			break
		}
		input = input[1:]
	}

	i := 0
	for {
		if input[i] == '\n' {
			dests = append(dests, input[:i])
			input = input[i+1:]
			break
		}
		if input[i] == ',' {
			dests = append(dests, input[:i])
			if input[i+1] != ' ' {
				panic(`unexpected`)
			}
			input = input[i+2:]
			i = 0
			continue
		}
		i++
	}

	return dests, input
}

type pulse bool

const (
	highPulse pulse = true
	lowPulse  pulse = false
)

type cable struct {
	source      string
	destination string

	pulse pulse
}

type machine struct {
	broadcaster broadcasterModule

	flipFlops []flipFlopModule

	conjunctions []conjunctionModule

	// TODO pending pulses
}

type broadcasterModule struct {
	destinations []string
}

type flipFlopModule struct {
	name string
	last pulse

	outputs []string
}

func newFlipFlopModule(name string, dests []string) flipFlopModule {
	return flipFlopModule{
		name:    name,
		last:    lowPulse,
		outputs: dests,
	}
}

type conjunctionModule struct {
	name string

	inputsByName map[string]pulse

	outputs []string
}

func newConjunctionModule(name string, dests []string) conjunctionModule {
	return conjunctionModule{
		name:         name,
		inputsByName: make(map[string]pulse, 8),
		outputs:      dests,
	}
}

func (c *conjunctionModule) addInput(
	src string,
) {
	c.inputsByName[src] = lowPulse
}
