package five

func Two(
	input string,
) (string, error) {
	stacks, ins, err := convertInputToStacksAndInstructions(input)
	if err != nil {
		return ``, err
	}

	for _, inst := range ins {
		values := make([]string, 0, inst.quantity)
		for i := 0; i < inst.quantity; i++ {
			v, err := stacks[inst.source-1].pop()
			if err != nil {
				return ``, err
			}
			values = append(values, v)
		}
		stacks[inst.dest-1].push(values...)
	}

	output := ``
	for _, s := range stacks {
		output += s.top()
	}
	return output, nil
}
