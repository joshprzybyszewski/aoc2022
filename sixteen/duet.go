package sixteen

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
