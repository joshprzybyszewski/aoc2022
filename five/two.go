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

	for _, line := range strings.Split(
		input[strings.Index(input, "\n\n")+2:],
		"\n",
	) {
		if line == `` {
			continue
		}
		inst, err = newInstruction(line)
		if err != nil {
			return ``, err
		}
		move(inst.source-1, inst.dest-1, inst.quantity)
	}

	var sb strings.Builder
	for i := range stacks {
		sb.WriteByte(stacks[i].values[stacks[i].length-1])
	}
	return sb.String(), nil
}
