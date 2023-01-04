package five

import "strings"

func Two(
	input string,
) (string, error) {
	stacks := newStacks()

	move := func(si, di int, q int) {
		for i := 0; i < q; i++ {
			stacks[di].values[stacks[di].length+i] = stacks[si].values[stacks[si].length-q+i]
		}
		stacks[si].length -= q
		stacks[di].length += q
	}

	var inst instruction
	var err error

	// jump past the stacks information
	input = input[strings.Index(input, "\n\n")+2:]

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			// skip empty lines
			input = input[1:]
			continue
		}
		inst, err = newInstruction(input[:nli])
		if err != nil {
			return ``, err
		}
		move(inst.source-1, inst.dest-1, inst.quantity)
		input = input[nli+1:]
	}

	output := make([]byte, 0, len(stacks))
	for i := range stacks {
		output = append(output, stacks[i].values[stacks[i].length-1])
	}
	return string(output), nil
}
