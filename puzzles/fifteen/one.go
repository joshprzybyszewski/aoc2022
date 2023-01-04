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

type sortedInts struct {
	vals []int
}

func newSortedInts(
	size int,
) *sortedInts {
	return &sortedInts{
		vals: make([]int, 0, size),
	}
}

func (s *sortedInts) clear() {
	s.vals = s.vals[:0]
}

func (s *sortedInts) add(v int) {
	// find the index to insert v
	i := s.getIndexFor(v)
	if i == len(s.vals) {
		// if it's at the end, just append
		s.vals = append(s.vals, v)
		return
	}

	// add an empty value at the end, to make space for the copy
	s.vals = append(s.vals, 0)

	// move the last elements one to the right
	copy(s.vals[i+1:], s.vals[i:])

	// insert the new value at the correct location
	s.vals[i] = v
}

func (s *sortedInts) getIndexFor(v int) int {
	// instead of using sort.SearchInts, we're going to iterate through the slice
	// backwards. This is because the reports are sorted in ascending x, then ascending y order.

	// return sort.SearchInts(s.vals, v)
	for i := len(s.vals) - 1; i > 0; i-- {
		if v >= s.vals[i] {
			return i + 1
		}
	}
	return 0
}

type tuples struct {
	starts *sortedInts
	ends   *sortedInts
}

func newTuples(size int) *tuples {
	return &tuples{
		starts: newSortedInts(size),
		ends:   newSortedInts(size),
	}
}

func (t *tuples) add(x1, x2 int) {
	t.starts.add(x1)
	t.ends.add(x2)
}

func (t *tuples) populate(
	rngs []tuple,
) []tuple {
	if len(t.starts.vals) == 0 {
		return nil
	}

	var si, ei int
	var curStart int

	active := 0

	for {
		if t.starts.vals[si] <= t.ends.vals[ei] {
			si++
			active++
			if active == 1 {
				curStart = t.starts.vals[0]
			}
			if si == len(t.starts.vals) {
				break
			}
		} else {
			active--
			if active == 0 {
				rngs = append(rngs, newTuple(curStart, t.ends.vals[ei]))
				curStart = t.starts.vals[0]
			}
			ei++
		}

	}
	rngs = append(rngs, newTuple(curStart, t.ends.vals[len(t.ends.vals)-1]))

	t.starts.clear()
	t.ends.clear()

	return rngs
}

func (t *tuples) findGap(
	start, end int,
) int {

	defer t.starts.clear()
	defer t.ends.clear()
	if len(t.starts.vals) == 0 {
		// no entries!
		return start - 1
	}

	var si, ei int

	active := 0

	for si < len(t.starts.vals) && ei < len(t.ends.vals) {
		if t.starts.vals[si] > end {
			// didn't find a gap
			return start - 1
		}

		if t.starts.vals[si] <= t.ends.vals[ei] {
			si++
			active++
			if si == len(t.starts.vals) {
				// didn't find a gap
				return start - 1
			}
		} else {
			active--
			if active == 0 {
				// found the end of the window
				if t.ends.vals[ei] > start {
					if t.ends.vals[ei] > end {
						// the end is past the range (we shouldn't hit this condition?)
						return start - 1
					}
					// it's within the [start:end] range!
					return t.ends.vals[ei] + 1
				}
			}
			ei++
		}

	}

	return start - 1
}

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
	xs := make(map[int]struct{}, len(rs))
	for _, r := range rs {
		if r.by == 2000000 {
			xs[r.bx] = struct{}{}
		}
	}

	output := make([]tuple, 0, len(rs))
	output = ts.populate(output)

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

	// calculated beacon distance
	dist int
}

func newReport(
	sx, sy int,
	bx, by int,
) report {
	dx := sx - bx
	if dx < 0 {
		dx = -dx
	}
	dy := sy - by
	if dy < 0 {
		dy = -dy
	}

	return report{
		sx:   sx,
		sy:   sy,
		bx:   bx,
		by:   by,
		dist: dx + dy,
	}
}

func (r report) getSeenInRow(y int) (int, int, bool) {
	ry := y - r.sy
	if ry < 0 {
		ry = -ry
	}

	if ry > r.dist {
		return 0, 0, false
	}

	return r.sx - r.dist + ry, r.sx + r.dist - ry, true
}

func getReports(
	input string,
) ([33]report, error) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15

	slice := make([]report, 0, 33)

	var i1, i2,
		sx, sy,
		bx, by int
	var err error
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}
		i1 = strings.Index(input, `x=`) + 2
		i2 = i1 + strings.Index(input[i1:], `,`)
		sx, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [33]report{}, err
		}

		i1 = strings.Index(input, `y=`) + 2
		i2 = i1 + strings.Index(input[i1:], `:`)
		sy, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [33]report{}, err
		}

		i1 = i2 + strings.Index(input[i2:], `x=`) + 2
		i2 = i1 + strings.Index(input[i1:], `,`)
		bx, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [33]report{}, err
		}

		i1 = i2 + strings.Index(input[i2:], `y=`) + 2
		by, err = strconv.Atoi(input[i1:nli])
		if err != nil {
			return [33]report{}, err
		}

		slice = append(slice, newReport(
			sx, sy,
			bx, by,
		))
		input = input[nli+1:]
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].sx == slice[j].sx {
			return slice[i].sy < slice[j].sy
		}
		return slice[i].sx < slice[j].sx
	})

	var output [33]report
	for i := range slice {
		output[i] = slice[i]
	}
	return output, nil
}
