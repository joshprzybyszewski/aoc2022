package five

import "fmt"

type stack struct {
	values []byte
}

func newStack() *stack {
	return &stack{
		values: nil,
	}
}

func (s *stack) top() byte {
	return s.values[len(s.values)-1]
}

func (s *stack) pop() (byte, error) {
	bs, err := s.popN(1)
	if err != nil {
		return 0, err
	}
	return bs[0], nil
}

func (s *stack) popN(n int) ([]byte, error) {
	if len(s.values) == 0 {
		if n == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf("too few elements to pop. %+v", s.values)
	}
	output := s.values[len(s.values)-n:]
	s.values = s.values[:len(s.values)-n]
	return output, nil
}

func (s *stack) push(ss ...byte) {
	s.values = append(s.values, ss...)
}
