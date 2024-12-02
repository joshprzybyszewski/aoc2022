package fifteen

func One(
	input string,
) (int, error) {
	total := 0
	var cur uint8
	for len(input) > 0 {
		switch input[0] {
		case ',':
			total += int(cur)
			cur = 0
		case '\n':
			// do nothing
		default:
			cur += uint8(input[0])
			cur *= 17
		}
		input = input[1:]
	}
	total += int(cur)

	return total, nil
}
