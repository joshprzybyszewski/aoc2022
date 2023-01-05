package twentyone

import "fmt"

func One(
	input string,
) (int, error) {
	monkeys, nameToIndex, err := convertToMonkeys(input)
	if err != nil {
		return 0, err
	}

	return monkeys[nameToIndex[`root`]].eval(), nil
}

type operation uint8

const (
	mult operation = '*'
	div  operation = '/'
	add  operation = '+'
	sub  operation = '-'
)

type monkey struct {
	value int

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

func (m *monkey) eval() int {
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
	other *monkey,
) bool {
	if m == nil {
		return false
	}
	if m == other {
		return true
	}

	return m.left.dependsOn(other) || m.right.dependsOn(other)
}

func (m *monkey) reverseEval(
	prev uint64,
	other *monkey,
) (uint64, bool) {
	// fmt.Printf("Received: %d\n", prev)
	if m == other {
		return prev, true
	}

	if m.left == nil {
		// no evaluation
		return uint64(m.value), false
	}
	okl := m.left.dependsOn(other)
	okr := m.right.dependsOn(other)
	if !okl && !okr {
		// fmt.Printf("\tDoesn't depend.\n")
		return uint64(m.eval()), false
	}

	var known, next uint64

	if okl {
		known = uint64(m.right.eval())
	} else {
		known = uint64(m.left.eval())
	}

	switch m.op {
	case mult:
		next = prev / known
		// if prev%known != 0 {
		// 	panic(`int division`)
		// }
	case div:
		next = known * prev
		if next < known {
			panic(`mult overflow`)
		}
	case add:
		next = prev - known
		if next < 0 {
			panic(`negative`)
		}
	case sub:
		next = known + prev
		if next < known {
			panic(`add overflow`)
		}
	}

	if okl {
		return m.left.reverseEval(next, other)
	} else {
		return m.right.reverseEval(next, other)
	}
}
