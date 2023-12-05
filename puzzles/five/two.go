package five

func Two(
	input string,
) (int, error) {

	seeds, mm, err := getSeedsAndMultiMapping(input)
	if err != nil {
		return 0, err
	}

	lowest, _ := mm.transformWithMax(seeds[0])
	tmp, s, maxS, max := 0, 0, 0, 0

	for i := 0; i < len(seeds); i += 2 {
		maxS = seeds[i] + seeds[i+1]
		for s = seeds[i]; s < maxS; {
			tmp, max = mm.transformWithMax(s)
			if tmp < lowest {
				lowest = tmp
			}
			if max < 1 {
				s++
			} else {
				// skip ahead
				s += (max - 1)
			}
		}
	}

	// 174137457 is too high
	return lowest, nil

}
