package fifteen

import (
	"errors"
)

func Two(
	input string,
) (int, error) {
	const (
		max = 4000000
	)

	rs, err := getReports(input)
	if err != nil {
		return 0, err
	}

	ts := newTuples()
	var output []tuple

	var x1, x2 int
	var ok bool
	var t tuple
	var r report

	for y := 0; y <= max; y++ {
		for _, r = range rs {
			x1, x2, ok = r.seenInRow(y)
			if ok {
				ts.add(x1, x2)
			}
		}
		output = ts.generate()
		// this assumes there's always output for the row
		for _, t = range output {
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
