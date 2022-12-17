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

	ts := newTuples(len(rs))

	var x1, x2 int
	var ok bool
	var t tuple
	var r report

	output := make([]tuple, 0, len(rs))

	// TODO iterate the rows based on how far they are from a sensor, not
	// in 0 to max order.
	for y := 0; y <= max; y++ {
		for _, r = range rs {
			x1, x2, ok = r.getSeenInRow(y)
			if ok {
				ts.add(x1, x2)
			}
		}
		output = output[:0]
		output = ts.populate(output)
		// this assumes there's always output for the row
		for _, t = range output {
			if t.t2 < 0 {
				// don't look at tuples that end before 0
				continue
			}
			if t.t1 > max {
				// don't look at any more tuples that start after the max
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
