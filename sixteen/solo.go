package sixteen

type soloPath struct {
	cur       node
	remaining distance

	valves valveState

	released pressure
}

func (s *soloPath) isOpen(n node) bool {
	return s.valves.isOpen(n)
}
