package twelve

func Two(
	input string,
) (int, error) {
	var p possibilities
	var groups []int

	var total int

	for len(input) > 0 {
		if input[0] == '\n' {
			input = input[1:]
			continue
		}

		p, groups, input = newPossibilities(input)
		groups = unfold(&p, groups)

		p.build(groups)

		total += p.answer(groups)
	}

	// 5728112261200 is too high
	// 5696803857515 is too high

	return total, nil
}

func unfold(
	p *possibilities,
	groups []int,
) []int {

	output := make([]int, 0, len(groups)*5)
	for i := 0; i < 5; i++ {
		output = append(output, groups...)
	}
	cpI := p.lineLength
	for i := 1; i < 5; i++ {
		p.line[cpI] = unknown
		cpI++
		for j := 0; j < p.lineLength; j++ {
			p.line[cpI] = p.line[j]
			cpI++
		}
	}
	p.lineLength = cpI

	return output
}
