package sixteen

type pressure int

type traveller struct {
	cur       node
	remaining distance
}

type soloPath struct {
	cur       node
	remaining distance

	valves valveState

	released pressure
}

func (s *soloPath) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

type duetPath struct {
	one, two traveller

	valves valveState

	released pressure
}

func (s *duetPath) isOpen(n node) bool {
	return s.valves.isOpen(n)
}

func (s *duetPath) numOpen() node {
	no := node(0)
	for n := node(0); n < numNodes; n++ {
		if s.valves.isOpen(n) {
			no++
		}
	}
	return no
}
