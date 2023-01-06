package fifteen

func One(
	input string,
) (int, error) {
	rs, err := getReports(input)
	if err != nil {
		return 0, err
	}

	ts := newTuples(len(rs))

	var x1, x2 int
	var ok bool

	for _, r := range rs {
		x1, x2, ok = r.getSeenInRow(2000000)
		if ok {
			ts.add(x1, x2)
		}
	}

	// record known x positions for beacons in this row
	xs := make([]int, 0, len(rs))
	contains := false
	for _, r := range rs {
		if r.by != 2000000 {
			continue
		}
		for _, x := range xs {
			if x == r.bx {
				contains = true
				break
			}
		}
		if !contains {
			xs = append(xs, r.bx)
		}
	}

	output := make([]tuple, 0, len(rs))
	output = ts.populate(output)

	// total starts at len(output) instead of zero because every element will add
	// (t2 - t1 + 1) to the total
	total := len(output)

	for _, t := range output {
		total += t.t2 - t.t1

		// remove any known beacons in this range
		for _, x := range xs {
			if t.t1 <= x && x <= t.t2 {
				total--
			}
		}
	}

	return total, nil
}
