package twelve

func Two(
	input string,
) (int, error) {
	var p possibilities

	var total int

	for len(input) > 0 {
		if input[0] == '\n' {
			input = input[1:]
			continue
		}

		p, input = newPossibilities(input)
		unfold(&p)

		p.build()

		total += p.answer()
	}

	return total, nil
}

func unfold(
	p *possibilities,
) {

	cpGI := p.numGroups
	for i := 1; i < 5; i++ {
		for j := 0; j < p.numGroups; j++ {
			p.groups[cpGI] = p.groups[j]
			cpGI++
		}
	}
	p.numGroups = cpGI

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
}
