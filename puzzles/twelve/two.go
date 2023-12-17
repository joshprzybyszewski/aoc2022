package twelve

func unfold(
	r *row,
	groups []int,
) []int {

	output := make([]int, 0, len(groups)*5)
	for i := 0; i < 5; i++ {
		output = append(output, groups...)
	}
	cpI := r.numParts
	for i := 1; i < 5; i++ {
		r.parts[cpI] = unknown
		cpI++
		for j := 0; j < r.numParts; j++ {
			r.parts[cpI] = r.parts[j]
			cpI++
		}
	}
	r.numParts = cpI

	return output
}

func Two(
	input string,
) (int, error) {
	total := 0
	var i, cur int
	var r row
	var addGroup bool
	groups := make([]int, 0, 40)
	for len(input) > 0 {
		if input[0] == '\n' {
			if r.numParts > 0 {
				groups = append(groups, cur)
				// fmt.Printf("  %-105s %v\n", r, groups)
				groups = unfold(&r, groups)
				// fmt.Printf("  %-105s %v\n", r, groups)
				num := getNum(
					r,
					0,
					groups,
					getRemainingRequired(groups),
				)
				// fmt.Printf("ANSWER: %d\n\n\n", num)
				total += num
			}
			i = 0
			cur = 0
			r = row{}
			addGroup = false
			groups = groups[:0]
		} else if addGroup {
			switch input[0] {
			case ',':
				// iterate past.
				groups = append(groups, cur)
				cur = 0
			default:
				cur *= 10
				cur += int(input[0] - '0')
			}
		} else {
			switch input[0] {
			case '?':
				r.parts[i] = unknown
			case '#':
				r.parts[i] = broken
			case ' ':
				r.numParts = i
				addGroup = true
			}
			i++
		}
		input = input[1:]
	}
	return total, nil
}
