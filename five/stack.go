package five

type stack struct {
	length int
	values [64]byte
}

func (s *stack) top() byte {
	return s.values[s.length-1]
}
