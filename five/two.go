package five

import "strings"

func Two(
	input string,
) (string, error) {
	stacks, ins, err := convertInputToStacksAndInstructions(input)
	if err != nil {
		return ``, err
	}

	for _, inst := range ins {
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
