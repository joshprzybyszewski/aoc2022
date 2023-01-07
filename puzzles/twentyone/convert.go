package twentyone

import (
	"strconv"
	"strings"
)

func convertToMonkeys(
	input string,
) ([]*monkey, map[string]int, error) {
	inputMonkeys, err := convertInput(input)
	if err != nil {
		return nil, nil, err
	}
	monkeys := make([]*monkey, len(inputMonkeys))

	nameToIndex := make(map[string]int, len(inputMonkeys))
	for i := range inputMonkeys {
		nameToIndex[inputMonkeys[i].name] = i
		monkeys[i] = &monkey{
			value: int64(inputMonkeys[i].value),
			op:    operation(inputMonkeys[i].operation),
		}
	}

	for i := range inputMonkeys {
		if inputMonkeys[i].left == `` {
			continue
		}
		monkeys[i].left = monkeys[nameToIndex[inputMonkeys[i].left]]
		if inputMonkeys[i].right != `` {
			monkeys[i].right = monkeys[nameToIndex[inputMonkeys[i].right]]
		}
	}
	return monkeys, nameToIndex, nil
}

type inputMonkey struct {
	name  string
	left  string
	right string

	operation byte

	value int
}

func convertInput(
	input string,
) ([]inputMonkey, error) {

	output := make([]inputMonkey, 0, 2163)

	var m inputMonkey
	var si int
	var x int
	var err error

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}
		m = inputMonkey{
			name: input[:4],
		}

		if '0' <= input[6] && input[6] <= '9' {
			// it's a value!
			x, err = strconv.Atoi(input[6:nli])
			if err != nil {
				return nil, err
			}
			m.value = x
		} else {
			// it has one or two other monkeys
			si = 6 + strings.Index(input[6:], ` `)
			if si > nli {
				m.left = input[6:nli]
			} else {
				m.left = input[6:si]
				m.operation = input[si+1]
				m.right = input[si+3 : nli]
			}
		}
		output = append(output, m)

		input = input[nli+1:]
	}

	return output, nil
}
