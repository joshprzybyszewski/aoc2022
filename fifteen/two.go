package fifteen

import (
	"errors"
)

func Two(
	input string,
) (int, error) {
	rs, err := getReports(input)
	if err != nil {
		return 0, err
	}

	const (
		max = 4000000
	)
	for y := 0; y <= max; y++ {
		ts := newTuples()

		for _, r := range rs {
			x1, x2, ok := r.seenInRow(y)
			if ok {
				ts.add(x1, x2)
			}
		}
		output := ts.generate()
		for _, t := range output {
			if t.t2 < 0 {
				continue
			}
			if t.t1 > max {
				break
			}
			if t.t1 > 0 {
				return (t.t1-1)*max + y, nil
			}
			if t.t2 < max {
				return (t.t2+1)*max + y, nil
			}
		}
	}

	return 0, errors.New("not found")
}
