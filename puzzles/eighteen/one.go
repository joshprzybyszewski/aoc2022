package eighteen

import (
	"fmt"
	"slices"
	"strings"
)

const (
	maxNumLines = 625
)

func One(
	input string,
) (int, error) {
	l := newLagoon(input, false)

	return l.numDug(), nil
}

type coord struct {
	row int
	col int
}

type upperBound uint8

const (
	unset upperBound = 0
	is    upperBound = 1
	isNot upperBound = 2
)

type lagoon struct {
	lines    [maxNumLines]line
	numLines int

	rowEdgesOnly    [maxNumLines]line
	numRowEdges     int
	isRowUpperBound [maxNumLines]upperBound

	colEdgesOnly [maxNumLines]line
	numColEdges  int

	min, max coord
}

func newLagoon(
	input string,
	isPart2 bool,
) lagoon {
	cur := coord{}

	var l lagoon
	l.min = cur
	l.max = cur
	var i int
	for len(input) > 0 {
		if isPart2 {
			l.lines[i], cur, input = newLineV2(cur, input)
		} else {
			l.lines[i], cur, input = newLine(cur, input)
		}
		i++
		input = input[1:]

		if cur.row < l.min.row {
			l.min.row = cur.row
		} else if cur.row > l.max.row {
			l.max.row = cur.row
		}
		if cur.col < l.min.col {
			l.min.col = cur.col
		} else if cur.col > l.max.col {
			l.max.col = cur.col
		}
	}
	l.numLines = i

	l.separateEdges()

	return l
}

func (l *lagoon) isHole(ri, ci int) bool {
	for rowI := 0; rowI < l.numRowEdges; rowI++ {
		row := l.rowEdgesOnly[rowI]
		if row.val <= ri {
			continue
		}
		if row.low <= ci && ci <= row.high {
			return !l.isUpperBound(rowI)
		}
	}
	return false
}

func (l *lagoon) getWallChar(ri, ci int) (byte, bool) {
	for rowI := 0; rowI < l.numRowEdges; rowI++ {
		row := l.rowEdgesOnly[rowI]
		if row.val < ri {
			continue
		}
		if row.val > ri {
			break
		}
		if row.low <= ci && ci <= row.high {
			if l.isUpperBound(rowI) {
				return 'v', true
			}
			return '^', true
		}
	}

	for li := 0; li < l.numLines; li++ {
		if l.lines[li].isHorizontal {
			continue
		}

		if l.lines[li].val == ci {
			if l.lines[li].low <= ri && ri <= l.lines[li].high {
				return '#', true
			}
		}
	}

	return 0, false
}

func (l *lagoon) String() string {
	return l.bounds(l.min, l.max)
}

func (l *lagoon) bounds(
	min, max coord,
) string {
	var sb strings.Builder

	// arbitrary bounds limit for drawing
	if max.row-min.row > 800 {
		max.row = min.row + 800
	}
	if max.col-min.col > 800 {
		max.col = min.col + 800
	}

	sb.WriteString(fmt.Sprintf("min: {%d, %d}, max {%d, %d}\n", min.row, min.col, max.row, max.col))

	for r := min.row; r <= max.row; r++ {
		for c := min.col; c <= max.col; c++ {
			wb, ok := l.getWallChar(r, c)
			if ok {
				sb.WriteByte(wb)
			} else if l.isHole(r, c) {
				sb.WriteByte('.')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (l *lagoon) numDug() int {

	n := 0

	for ri := 0; ri < l.numRowEdges; ri++ {
		if !l.isUpperBound(ri) {
			continue
		}
		below := l.getRowEdgesBelow(ri)

		top := l.rowEdgesOnly[ri]

		if l.hasDescendingLeft(top) {
			top.low++
		}
		if l.hasDescendingRight(top) {
			top.high--
		}

		n += getAreaUnder(top, below)
	}

	// iterate through all lines.
	//   The total will increase by their length.
	for li := 0; li < l.numLines; li++ {
		n += (l.lines[li].high - l.lines[li].low)
	}

	return n
}

func (l *lagoon) hasDescendingLeft(row line) bool {
	column := row.low
	for li := 0; li < l.numLines; li++ {
		if l.lines[li].isHorizontal {
			continue
		}

		if l.lines[li].val != column {
			continue
		}

		if l.lines[li].low == row.val {
			return true
		}
		if l.lines[li].high == row.val {
			return false
		}
	}

	return false
}

func (l *lagoon) hasDescendingRight(row line) bool {
	column := row.high
	for li := 0; li < l.numLines; li++ {
		if l.lines[li].isHorizontal {
			continue
		}

		if l.lines[li].val != column {
			continue
		}

		if l.lines[li].low == row.val {
			return true
		}
		if l.lines[li].high == row.val {
			return false
		}
	}

	return false
}

func (l *lagoon) separateEdges() {
	ri := 0
	ci := 0
	for i := 0; i < l.numLines; i++ {
		if l.lines[i].isHorizontal {
			l.rowEdgesOnly[ri] = l.lines[i]
			ri++
		} else {
			l.colEdgesOnly[ci] = l.lines[i]
			ci++
		}
	}
	l.numRowEdges = ri
	l.numColEdges = ci

	// sort the row edges so the first ones are at the top
	// and then they go left to right within the same row
	slices.SortFunc(l.rowEdgesOnly[:l.numRowEdges], func(a, b line) int {
		if a.val != b.val {
			return a.val - b.val
		}
		return a.low - b.low
	})

	slices.SortFunc(l.colEdgesOnly[:l.numColEdges], func(a, b line) int {
		if a.val != b.val {
			return a.val - b.val
		}
		return a.low - b.low
	})
}

func (l *lagoon) isUpperBound(
	ri int,
) (answer bool) {
	if l.isRowUpperBound[ri] != unset {
		return l.isRowUpperBound[ri] == is
	}

	defer func(ri int) {
		if answer {
			l.isRowUpperBound[ri] = is
		} else {
			l.isRowUpperBound[ri] = isNot
		}
	}(ri)

	var leftConnected, rightConnected bool
	var above line

	myLine := l.rowEdgesOnly[ri]
	for ri = ri - 1; ri >= 0; ri-- {
		above = l.rowEdgesOnly[ri]

		leftConnected = (above.low == myLine.low || above.high == myLine.low) && l.hasVerticalLine(
			myLine.low,
			above.val,
			myLine.val,
		)
		rightConnected = (above.low == myLine.high || above.high == myLine.high) && l.hasVerticalLine(
			myLine.high,
			above.val,
			myLine.val,
		)

		if leftConnected && rightConnected {
			// this is only true in the case of a single circle loop
			if l.numLines != 4 {
				panic(`unexpected`)
			}
			return !l.isUpperBound(ri)
		}
		if leftConnected {
			if above.low == myLine.low {
				return !l.isUpperBound(ri)
			}
			return l.isUpperBound(ri)
		}
		if rightConnected {
			if above.high == myLine.high {
				return !l.isUpperBound(ri)
			}
			return l.isUpperBound(ri)
		}

		if above.low <= myLine.low && myLine.low <= above.high {
			return !l.isUpperBound(ri)
		}
	}

	return true
}

func (l *lagoon) hasVerticalLine(
	column int,
	low, high int,
) bool {
	// TODO this could be faster if we do a slices.Index or sort.Search
	for ci := 0; ci < l.numColEdges; ci++ {
		if l.colEdgesOnly[ci].val < column {
			continue
		}
		if l.colEdgesOnly[ci].val > column {
			break
		}

		if l.colEdgesOnly[ci].low == low &&
			l.colEdgesOnly[ci].high == high {
			return true
		}
	}
	return false
}

func (l *lagoon) getRowEdgesBelow(
	ri int,
) []line {
	output := make([]line, 0, 26) // 26 is arbitrary, but it's the largest that my puzzle endpoint needs

	row := l.rowEdgesOnly[ri]
	for l.rowEdgesOnly[ri].val == row.val {
		ri++ // iterate past the ones on the same row
	}
	for ; ri < l.numRowEdges; ri++ {
		if row.overlaps(l.rowEdgesOnly[ri]) {
			output = append(output, l.rowEdgesOnly[ri])
		}
	}

	return output
}

func getAreaUnder(
	top line,
	below []line,
) int {
	total := 0

	getConstraintForColumn := func(
		column int,
	) line {
		for _, l := range below {
			if l.low <= column && column <= l.high {
				return l
			}
		}
		panic(`unexpected`)
	}

	getRightBounds := func(
		constraint line,
		startingCol int,
	) int {
		var min line
		found := false
		for _, l := range below {
			if l.val >= constraint.val {
				break
			}
			if l.low < startingCol {
				continue
			}
			if l.low > constraint.high {
				// starts beyond my line
				continue
			}
			if !found {
				min = l
				found = true
			} else if l.low < min.low {
				min = l
			}
		}

		if found {
			// there's something above. This box must end before it.
			return min.low - 1
		}

		// the next one must be below. We go all the way to the end
		// of this constraint.
		return constraint.high
	}

	bb := boxBounds{
		topRow:  top.val + 1,
		leftCol: top.low,
	}

	var bottom line
	for bb.leftCol <= top.high {
		bottom = getConstraintForColumn(bb.leftCol)

		bb.bottomRow = bottom.val - 1
		bb.rightCol = getRightBounds(bottom, bb.leftCol)
		if bb.rightCol > top.high {
			bb.rightCol = top.high
		}

		total += bb.area()

		bb.leftCol = bb.rightCol + 1
	}

	return total
}

type boxBounds struct {
	topRow    int
	bottomRow int
	leftCol   int
	rightCol  int
}

func (bb *boxBounds) area() int {
	return (bb.bottomRow - bb.topRow + 1) * (bb.rightCol - bb.leftCol + 1)
}

type line struct {
	low  int
	high int

	val int

	isHorizontal bool
}

func newLine(
	cur coord,
	input string,
) (line, coord, string) {
	dir := input[0]

	input = input[2:]
	num := 0
	for input[0] != ' ' {
		num *= 10
		num += int(input[0] - '0')
		input = input[1:]
	}
	input = input[10:]

	var l line
	switch dir {
	case 'R':
		l.isHorizontal = true
		l.val = cur.row
		l.low = cur.col
		cur.col += num
		l.high = cur.col
	case 'L':
		l.isHorizontal = true
		l.val = cur.row
		l.high = cur.col
		cur.col -= num
		l.low = cur.col
	case 'U':
		l.isHorizontal = false
		l.val = cur.col
		l.high = cur.row
		cur.row -= num
		l.low = cur.row
	case 'D':
		l.isHorizontal = false
		l.val = cur.col
		l.low = cur.row
		cur.row += num
		l.high = cur.row
	default:
		panic(`unexpected heading ` + string(dir))
	}

	return l, cur, input
}

func (l line) overlaps(other line) bool {
	// assume they're both rows.
	return l.low <= other.high && l.high >= other.low
}

func (l line) String() string {
	if l.isHorizontal {
		return fmt.Sprintf("Row at %d in cols [%d, %d]", l.val, l.low, l.high)
	}
	return fmt.Sprintf("Col at %d in rows [%d, %d]", l.val, l.low, l.high)
}
