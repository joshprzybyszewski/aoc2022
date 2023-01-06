package fifteen

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
) sortedInts {
	return sortedInts{
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
	for i := len(s.vals) - 1; i >= 0; i-- {
		if v >= s.vals[i] {
			return i + 1
		}
	}
	return 0
}

type tuples struct {
	starts sortedInts
	ends   sortedInts
}

func newTuples(size int) tuples {
	return tuples{
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
