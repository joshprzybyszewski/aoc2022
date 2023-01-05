package twentyone

import "fmt"

func One(
	input string,
) (int, error) {
	monkeys, nameToIndex, err := convertToMonkeys(input)
	if err != nil {
		return 0, err
	}

	return int(monkeys[nameToIndex[`root`]].eval()), nil
}

type operation uint8

const (
	mult operation = '*'
	div  operation = '/'
	add  operation = '+'
	sub  operation = '-'
)

type monkey struct {
	value int64

	left  *monkey
	right *monkey
	op    operation
}

func (m *monkey) Print() {
	m.print(0)
}

func (m *monkey) print(
	indents int,
) {
	bi := buildIndents(indents)
	if m.left == nil {
		fmt.Printf(
			"%svalue: %d\n",
			bi,
			m.value,
		)
		return
	}
	if m.right != nil {
		fmt.Printf(
			"%sop: %s\n",
			bi, string(m.op),
		)
		fmt.Printf(
			"%sleft:\n",
			bi,
		)
		m.left.print(indents + 1)
		fmt.Printf(
			"%sright:\n",
			bi,
		)
		m.right.print(indents + 1)
		return
	}

	fmt.Printf(
		"%snoop\n",
		bi,
	)
	fmt.Printf(
		"%sleft:\n",
		bi,
	)
	m.left.print(indents + 1)
}

func buildIndents(n int) string {
	out := ``
	for i := 0; i < n; i++ {
		out += ` `
	}
	return out
}

func (m *monkey) eval() int64 {
	if m.left == nil {
		return m.value
	}

	l := m.left.eval()
	if m.right == nil {
		return l
	}

	r := m.right.eval()

	switch m.op {
	case mult:
		return l * r
	case div:
		return l / r
	case add:
		return l + r
	case sub:
		return l - r
	default:
		panic(`ahhh`)
	}
}

func (m *monkey) dependsOn(
	target *monkey,
) bool {
	if m == nil {
		return false
	}
	if m == target {
		return true
	}

	return m.left.dependsOn(target) || m.right.dependsOn(target)
}

func (m *monkey) reverseEval(
	prev int64,
	target *monkey,
) (int64, bool) {
	if m == target {
		return prev, true
	}

	if m.left == nil {
		// no evaluation
		return int64(m.value), false
	}

	leftDep := m.left.dependsOn(target)
	if !leftDep {
		if !m.right.dependsOn(target) {
			// neither side depends on the target => standard evaluation
			return m.eval(), false
		}
	}

	var known, next int64

	if leftDep {
		known = int64(m.right.eval())
	} else {
		known = int64(m.left.eval())
	}

	switch m.op {
	case mult:
		next = prev / known
	case div:
		next = known * prev
	case add:
		next = prev - known
	case sub:
		// this tricked me:
		if leftDep {
			next = known + prev
		} else {
			next = known - prev
		}
	}

	if leftDep {
		return m.left.reverseEval(next, target)
	}
	return m.right.reverseEval(next, target)
}
