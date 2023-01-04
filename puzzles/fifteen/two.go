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

	var x1, x2, gap int
	var ok bool
	var r report

	for y := 0; y <= max; y++ {
		for _, r = range rs {
			x1, x2, ok = r.getSeenInRow(y)
			if ok {
				ts.add(x1, x2)
			}
		}
		gap = ts.findGap(0, max)
		if gap >= 0 {
			return gap*max + y, nil
		}
	}

	return 0, errors.New("not found")
}
