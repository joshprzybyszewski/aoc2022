package eighteen

import (
	"fmt"
	"strings"
)

const (
	maxRowCol = 10000
)

func One(
	input string,
) (int, error) {
	var l lagoon
	var i int
	for len(input) > 0 {
		l.edges[i], input = newEdge(input)
		i++
		input = input[1:]
	}
	l.numEdges = i

	l.dig()

	return l.numDug(), nil
}

type coord struct {
	row int
	col int
}

type lagoon struct {
	edges    [625]edge
	numEdges int

	holes [maxRowCol][maxRowCol]bool
	paths [maxRowCol][maxRowCol]int

	min, max coord
}

func (l *lagoon) String() string {
	var sb strings.Builder
	for r := l.min.row; r <= l.max.row; r++ {
		if l.holes[r] == [maxRowCol]bool{} {
			continue
		}
		for c := l.min.col; c <= l.max.col; c++ {
			if l.holes[r][c] {
				sb.WriteByte('#')
			} else if l.paths[r][c]&1 == 1 {
				sb.WriteByte('.')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (l *lagoon) dig() {
	cur := coord{
		row: maxRowCol / 2,
		col: maxRowCol / 2,
	}
	next := cur

	l.min = cur
	l.max = cur

	for i := 0; i < l.numEdges; i++ {
		e := l.edges[i]
		switch e.heading {
		case east:
			next.col += e.num
			for c := cur.col; c <= next.col; c++ {
				l.holes[cur.row][c] = true
			}
		case south:
			next.row += e.num
			for r := cur.row; r <= next.row; r++ {
				l.holes[r][cur.col] = true
			}
		case west:
			next.col -= e.num
			for c := cur.col; c >= next.col; c-- {
				l.holes[cur.row][c] = true
			}
		case north:
			next.row -= e.num
			for r := cur.row; r >= next.row; r-- {
				l.holes[r][cur.col] = true
			}
		default:
			panic(`ahhhh`)
		}
		cur = next
		if cur.row < l.min.row {
			l.min.row = cur.row
		}
		if cur.col < l.min.col {
			l.min.col = cur.col
		}
		if cur.row > l.max.row {
			l.max.row = cur.row
		}
		if cur.col > l.max.col {
			l.max.col = cur.col
		}
	}

	l.calcHoles()
}

func (l *lagoon) calcHoles() {

	for r := l.min.row - 1; r <= l.max.row; r++ {
		n := 0
		fromAbove := false
		fromBelow := false

		for c := l.min.col - 1; c <= l.max.col; c++ {
			if !l.holes[r][c] {
				l.paths[r][c] = n
				continue
			}

			if l.holes[r-1][c] {
				fromAbove = true
			} else if l.holes[r+1][c] {
				fromBelow = true
			} else {
				fmt.Printf("lagoon:\n%s\n", l.String())
				panic(`ahh`)
			}

			for l.holes[r][c] {
				l.paths[r][c] = n
				c++
			}

			if fromBelow {
				if l.holes[r-1][c-1] {
					n++
				}
			} else if fromAbove {
				if l.holes[r+1][c-1] {
					n++
				}
			} else {
				fmt.Printf("lagoon:\n%s\n", l.String())
				panic(`ahh`)
			}

			l.paths[r][c] = n
			fromAbove = false
			fromBelow = false
		}
	}
}

func (l *lagoon) numDug() int {
	n := 0
	for r := l.min.row; r <= l.max.row; r++ {
		for c := l.min.col; c <= l.max.col; c++ {
			if l.holes[r][c] {
				n++
			} else if l.paths[r][c]&1 == 1 {
				n++
			}
		}
	}
	return n
}

type edge struct {
	heading heading
	num     int
	color   string
}

func newEdge(input string) (edge, string) {
	var e edge
	switch input[0] {
	case 'R':
		e.heading = east
	case 'L':
		e.heading = west
	case 'U':
		e.heading = north
	case 'D':
		e.heading = south
	default:
		panic(`unexpected heading ` + string(input[0]) + ` from ` + input)
	}

	input = input[2:]
	for input[0] != ' ' {
		e.num *= 10
		e.num += int(input[0] - '0')
		input = input[1:]
	}
	input = input[2:]

	e.color = input[:7]
	input = input[8:]

	return e, input
}

type heading uint8

const (
	east  heading = 1
	south heading = 2
	west  heading = 3
	north heading = 4
)
