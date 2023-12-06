package three

import (
	"strings"
)

type coord struct {
	row int
	col int
}

type symbols struct {
	symbols [size][size]bool
}

func newSymbols() symbols {
	return symbols{}
}

func (s *symbols) addSymbol(row, col int, c byte) {
	if c == '.' || (c >= '0' && c <= '9') {
		return
	}

	s.symbols[row][col] = true
}

func (s *symbols) isNextToSymbol(row, col int, numDigits int) bool {
	minCol := col - 1 - numDigits
	if minCol < 0 {
		minCol = 0
	}

	if col == size {
		col--
	}

	if s.symbols[row][col] || s.symbols[row][minCol] {
		return true
	}

	var tmpCol int

	if row > 0 {
		for tmpCol = col; tmpCol >= minCol; tmpCol-- {
			if s.symbols[row-1][tmpCol] {
				return true
			}
		}
	}

	if row+1 < size {
		for tmpCol = col; tmpCol >= minCol; tmpCol-- {
			if s.symbols[row+1][tmpCol] {
				return true
			}
		}
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
			s.addSymbol(row, col, input[col])
		}
		row++
		input = input[nli+1:]
	}

	total := 0
	curNum, numDigits := 0, 0
	row = 0
	input = fullInput
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		for col = 0; col <= nli; col++ {
			if input[col] < '0' || input[col] > '9' {
				if curNum > 0 && s.isNextToSymbol(row, col, numDigits) {
					total += curNum
				}
				curNum = 0
				numDigits = 0
				continue
			}
			numDigits++
			curNum *= 10
			curNum += int(input[col] - '0')
		}
		input = input[nli+1:]
		row++
	}

	return total, nil
}
