package twentyfour

import "strings"

type square uint8

const (
	empty square = 0
	right square = 1 << iota
	down
	left
	up
	wall
)

type board [27][122]square

func next(
	b board,
) board {
	var n board
	var c int
	var s square
	for r := range b {
		for c, s = range b[r] {
			if s == wall {
				n[r][c] = wall
				continue
			}

			if s&right != 0 {
				if c == len(b[r])-2 {
					n[r][1] |= right
				} else {
					n[r][c+1] |= right
				}
			}
			if s&down != 0 {
				if r == len(b)-2 {
					n[1][c] |= down
				} else {
					n[r+1][c] |= down
				}
			}
			if s&left != 0 {
				if c == 1 {
					n[r][len(b[r])-2] |= left
				} else {
					n[r][c-1] |= left
				}
			}
			if s&up != 0 {
				if r == 1 {
					n[len(b)-2][c] |= up
				} else {
					n[r-1][c] |= up
				}
			}
		}
	}
	return n
}

func getBoard(
	input string,
) board {

	var b board
	r, c := 0, 0

	for _, ch := range input {
		if ch == '\n' {
			r++
			c = 0
			continue
		}

		switch ch {
		case '#':
			b[r][c] = wall
		case '.':
			b[r][c] = empty
		case '>':
			b[r][c] = right
		case 'v':
			b[r][c] = down
		case '<':
			b[r][c] = left
		case '^':
			b[r][c] = up
		}
		c++
	}
	return b
}

func prettyBoard(
	b *board,
	pos position,
) string {
	var sb strings.Builder
	var ch byte
	var c int
	var s square
	for r := range b {
		for c, s = range b[r] {
			if r == pos.row && c == pos.col {
				if s != empty {
					sb.WriteByte('!')
				} else {
					sb.WriteByte('E')
				}
				continue
			}
			if s == wall {
				sb.WriteByte('#')
				continue
			}
			ch = '.'
			if s&right != 0 {
				ch = '>'
			}
			if s&down != 0 {
				if ch != '.' {
					ch = '2'
				} else {
					ch = 'v'
				}
			}
			if s&left != 0 {
				if ch == '2' {
					ch = '3'
				} else if ch == '.' {
					ch = '<'
				} else {
					ch = '2'
				}
			}
			if s&up != 0 {
				if ch == '3' {
					ch = '4'
				} else if ch == '2' {
					ch = '3'
				} else if ch == '.' {
					ch = '^'
				} else {
					ch = '2'
				}
			}
			sb.WriteByte(ch)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
