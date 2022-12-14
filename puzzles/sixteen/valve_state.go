package sixteen

// valveState is an array of bools. When true, the valve has been opened
type valveState uint16

func (vs valveState) isOpen(n node) bool {
	return vs&(1<<n) != 0
}

func (vs valveState) open(n node) valveState {
	return vs | (1 << n)
}

func (vs valveState) openTwo(n1, n2 node) valveState {
	return vs | (1 << n1) | (1 << n2)
}
