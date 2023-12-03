package three

import (
	"strings"
)

type coord struct {
	row int
	col int
}

type symbols map[coord]byte

func newSymbols() symbols {
	return make(symbols)
}

func (s symbols) add(row, col int, c byte) {
	s[coord{
		row: row,
		col: col,
	}] = c
}

func (s symbols) isNextToSymbol(row, col int) bool {
	// check left, then below, then right, then above
	tmp := coord{
		row: row,
		col: col - 1,
	}
	_, ok := s[tmp]
	if ok {
		return true
	}
	tmp.row--
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.col++
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.row++
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		return true
	}
	tmp.col--
	_, ok = s[tmp]
	if ok {
		return true
	}
	return false
}

func One(
	fullInput string,
) (int, error) {
	input := fullInput

	var row, col int
	s := newSymbols()

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col < nli; col++ {
			if input[col] == '.' || (input[col] >= '0' && input[col] <= '9') {
				continue
			}
			s.add(row, col, input[col])
		}
		row++
		input = input[nli+1:]
	}

	total := 0
	curNum := 0
	shouldAdd := false
	row = 0
	input = fullInput
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col < nli; col++ {
			if input[col] < '0' || input[col] > '9' {
				if shouldAdd {
					total += curNum
					shouldAdd = false
				}
				curNum = 0
				continue
			}
			curNum *= 10
			curNum += int(input[col] - '0')
			shouldAdd = shouldAdd || s.isNextToSymbol(row, col)
		}
		input = input[nli+1:]
		row++
	}

	return total, nil
}
