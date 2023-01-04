package twentytwo

func Two(
	input string,
) (int, error) {
	start, dirs := convertInputToCube(input)
	s := start
	h := right
	for i, d := range dirs {
		s, h = moveInCube(s, d.dist, h)
		if i == len(dirs)-1 {
			// don't mutate h the last time.
			break
		}
		if d.clockwise {
			if h == 3 {
				h = 0
			} else {
				h++
			}
		} else {
			if h == 0 {
				h = 3
			} else {
				h--
			}
		}
	}

	// 111144, 111145 is too low
	return (1000 * int(s.row)) +
		(4 * int(s.col)) +
		int(h), nil
}
