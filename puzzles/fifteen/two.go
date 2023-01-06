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

	var x1, x2 int
	var ok bool
	var start [numReports]int
	var spans [numReports]tuple
	for i := range spans {
		x1, x2, ok = rs[i].getSeenInRow(0)
		if ok {
			spans[i].t1 = x1
			spans[i].t2 = x2
		} else {
			spans[i].t2 = -1
			start[i] = rs[i].sy - rs[i].dist
		}
	}

	ts := newTuples(len(rs))

	var gap int
	var i int

	for y := 0; y <= max; y++ {
		for i = range spans {
			if spans[i].t2 < spans[i].t1 {
				if y == start[i] {
					x1, x2, _ = rs[i].getSeenInRow(y)
					spans[i].t1 = x1
					spans[i].t2 = x2
					ts.add(spans[i].t1, spans[i].t2)
				}
			} else {
				if y <= rs[i].sy {
					spans[i].t1--
					spans[i].t2++
					ts.add(spans[i].t1, spans[i].t2)
				} else if spans[i].t1 == spans[i].t2 {
					spans[i].t1 = 1
					spans[i].t2 = 0
				} else {
					spans[i].t1++
					spans[i].t2--
					ts.add(spans[i].t1, spans[i].t2)
				}
			}
		}
		gap = ts.findGap(0, max)
		if gap >= 0 {
			return gap*max + y, nil
		}
	}

	return 0, errors.New("not found")
}
