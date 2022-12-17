package fifteen

import (
	"sort"
	"strconv"
	"strings"
)

type tuple struct {
	t1, t2 int
}

func newTuple(x1, x2 int) tuple {
	return tuple{
		t1: x1,
		t2: x2,
	}
}

type tuples struct {
	starts []int
	ends   []int
}

func newTuples() *tuples {
	return &tuples{
		starts: nil,
		ends:   nil,
	}
}

func (t *tuples) add(x1, x2 int) {
	t.starts = append(t.starts, x1)
	t.ends = append(t.ends, x2)
}

func (t *tuples) generate() []tuple {
	if len(t.starts) == 0 {
		return nil
	}

	sort.Ints(t.starts)
	sort.Ints(t.ends)

	output := make([]tuple, 0, len(t.starts))

	var si, ei int
	var curStart int

	active := 0

	for {
		if t.starts[si] <= t.ends[ei] {
			si++
			active++
			if active == 1 {
				curStart = t.starts[0]
			}
			if si == len(t.starts) {
				break
			}
		} else {
			active--
			if active == 0 {
				output = append(output, newTuple(curStart, t.ends[ei]))
				curStart = t.starts[0]
			}
			ei++
		}

	}
	output = append(output, newTuple(curStart, t.ends[len(t.ends)-1]))

	t.starts = t.starts[:0]
	t.ends = t.ends[:0]

	return output
}

func One(
	input string,
) (int, error) {
	rs, err := getReports(input)
	if err != nil {
		return 0, err
	}

	ts := newTuples()

	for _, r := range rs {
		x1, x2, ok := r.seenInRow(2000000)
		if ok {
			ts.add(x1, x2)
		}
	}

	// record known x positions for beacons in this row
	xs := make(map[int]struct{}, len(rs))
	for _, r := range rs {
		if r.by == 2000000 {
			xs[r.bx] = struct{}{}
		}
	}

	output := ts.generate()

	// total starts at len(output) instead of zero because every element will add
	// (t2 - t1 + 1) to the total
	total := len(output)

	for _, t := range output {
		total += (t.t2 - t.t1)

		// remove any known beacons in this range
		for x := range xs {
			if t.t1 <= x && x <= t.t2 {
				total--
			}
		}
	}

	return total, nil
}

type report struct {
	sx, sy int
	bx, by int
}

func (r report) beaconDistance() int {
	dx := r.sx - r.bx
	if dx < 0 {
		dx = -dx
	}
	dy := r.sy - r.by
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func (r report) seenInRow(y int) (int, int, bool) {
	ry := y - r.sy
	if ry < 0 {
		ry = -ry
	}

	dist := r.beaconDistance()

	if ry > dist {
		return 0, 0, false
	}

	rx := dist - ry

	return r.sx - rx, r.sx + rx, true
}

func getReports(
	input string,
) ([]report, error) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	lines := strings.Split(input, "\n")

	output := make([]report, 0, len(lines)-1)

	var i1, i2,
		sx, sy,
		bx, by int
	var err error
	for _, line := range lines {
		if line == `` {
			continue
		}
		i1 = strings.Index(line, `x=`) + 2
		i2 = i1 + strings.Index(line[i1:], `,`)
		sx, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return nil, err
		}

		i1 = strings.Index(line, `y=`) + 2
		i2 = i1 + strings.Index(line[i1:], `:`)
		sy, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return nil, err
		}

		i1 = strings.LastIndex(line, `x=`) + 2
		i2 = i1 + strings.Index(line[i1:], `,`)
		bx, err = strconv.Atoi(line[i1:i2])
		if err != nil {
			return nil, err
		}

		i1 = strings.LastIndex(line, `y=`) + 2
		// i2 = i1 + strings.Index(line[i1:], `,`)
		by, err = strconv.Atoi(line[i1:])
		if err != nil {
			return nil, err
		}

		output = append(output, report{
			sx: sx,
			sy: sy,
			bx: bx,
			by: by,
		})
	}

	return output, nil
}
