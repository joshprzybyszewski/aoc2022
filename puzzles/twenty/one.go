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
	// fmt.Printf("machine: %+v\n", m)

	p := newProcessor(&m)

	for n := 0; n < 1000; n++ {
		p.pushButton()
	}

	return p.numLowPulsesSent * p.numHighPulsesSent, nil
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

func (p pulse) String() string {
	if p == highPulse {
		return `high`
	}
	return `low`
}

type cable struct {
	source      string
	destination string

	pulse pulse
}

func (c cable) String() string {
	return fmt.Sprintf("%s -%s-> %s", c.source, c.pulse, c.destination)
}

type machine struct {
	broadcaster broadcasterModule

	flipFlops []flipFlopModule

	conjunctions []conjunctionModule
}

func (m *machine) getFlipFlop(
	dest string,
) *flipFlopModule {
	for i := range m.flipFlops {
		if m.flipFlops[i].name == dest {
			return &m.flipFlops[i]
		}
	}
	return nil
}

func (m *machine) getConjunction(
	dest string,
) *conjunctionModule {
	for i := range m.conjunctions {
		if m.conjunctions[i].name == dest {
			return &m.conjunctions[i]
		}
	}
	return nil
}

type processor struct {
	m *machine

	numLowPulsesSent  int
	numHighPulsesSent int

	queue []cable
}

func newProcessor(
	m *machine,
) processor {
	return processor{
		m:     m,
		queue: make([]cable, 0, 5000),
	}
}

func (p *processor) pushButton() {
	const (
		buttonPulse = lowPulse
	)
	p.numLowPulsesSent += 1
	for _, dest := range p.m.broadcaster.destinations {
		p.addPulse(
			`broadcaster`,
			dest,
			buttonPulse,
		)
	}

	p.resolvePulses()
}

func (p *processor) addPulse(
	src string,
	dest string,
	state pulse,
) {
	p.queue = append(p.queue, cable{
		source:      src,
		destination: dest,
		pulse:       state,
	})
}

func (p *processor) resolvePulses() {

	var c cable
	var ffm *flipFlopModule
	var cm *conjunctionModule

	for len(p.queue) > 0 {
		c = p.queue[0]
		if c.pulse == lowPulse {
			p.numLowPulsesSent++
		} else if c.pulse == highPulse {
			p.numHighPulsesSent++
		} else {
			panic(`unexpected`)
		}
		p.queue = p.queue[1:]

		// fmt.Printf("Processing: %s\n", c)

		if ffm = p.m.getFlipFlop(c.destination); ffm != nil {
			p.resolveFlipFlop(c, ffm)
		} else if cm = p.m.getConjunction(c.destination); cm != nil {
			p.resolveConjunction(c, cm)
		} else {
			panic(`unexpected`)
		}
	}
}

func (p *processor) resolveFlipFlop(
	c cable,
	ffm *flipFlopModule,
) {
	if c.pulse == highPulse {
		return
	}

	ffm.isOn = !ffm.isOn

	pulse := lowPulse
	if ffm.isOn {
		pulse = highPulse
	}
	for _, dest := range ffm.outputs {
		p.addPulse(ffm.name, dest, pulse)
	}
}

func (p *processor) resolveConjunction(
	c cable,
	cm *conjunctionModule,
) {
	cm.inputsByName[c.source] = c.pulse

	pulse := highPulse
	if c.pulse == highPulse {
		pulse = lowPulse
		for _, p := range cm.inputsByName {
			if p == lowPulse {
				pulse = highPulse
				break
			}
		}
	}

	for _, dest := range cm.outputs {
		p.addPulse(cm.name, dest, pulse)
	}
}

type broadcasterModule struct {
	destinations []string
}

type flipFlopModule struct {
	name string
	isOn bool

	outputs []string
}

func newFlipFlopModule(name string, dests []string) flipFlopModule {
	return flipFlopModule{
		name:    name,
		isOn:    false,
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
