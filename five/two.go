package five

import "strings"

func Two(
	input string,
) (string, error) {
	stacks := newStacks()

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
		values, err := stacks[inst.source-1].popN(inst.quantity)
		if err != nil {
			return ``, err
		}
		stacks[inst.dest-1].push(values...)
	}

	var sb strings.Builder
	for _, s := range stacks {
		sb.WriteByte(s.top())
	}
	return sb.String(), nil
}
