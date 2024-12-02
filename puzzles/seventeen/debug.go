package seventeen

import (
	"fmt"
	"strings"
)

func (p position) String() string {
	return fmt.Sprintf("(%3d, %3d) %d", p.row, p.col, p.totalHeatLoss)
}

func (c city) String() string {
	var sb strings.Builder
	sb.WriteString("        ")
	for ci := 0; ci < citySize; ci++ {
		sb.WriteString(fmt.Sprintf("%4d ", ci))
	}
	sb.WriteByte('\n')
	for ri := 0; ri < citySize; ri++ {
		sb.WriteString(fmt.Sprintf("Row %3d:", ri))
		for ci := 0; ci < citySize; ci++ {
			val := c.getMinValAt(ri, ci)
			if val == 0 {
				sb.WriteString("     ")
			} else {
				sb.WriteString(fmt.Sprintf("%4d ", val))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

var (
	spaceBytes = []byte{
		'.',
		',',
		':',
		';',
		'i',
		'c',
		'o',
		'C',
		'I',
		'O',
		'0',
		'%',
		'#',
	}
)

func (c city) withPos(pos position) string {
	var headings [citySize][citySize]heading
	var numInDir [citySize][citySize]int

	min := position{
		row: citySize,
		col: citySize,
	}

	for cur := &pos; cur != nil; cur = cur.prev {
		if cur.row >= citySize ||
			cur.col >= citySize ||
			cur.row < 0 ||
			cur.col < 0 {
			continue
		}
		headings[cur.row][cur.col] = cur.heading
		numInDir[cur.row][cur.col] = int(cur.leftInDirection)
		if cur.row < min.row {
			min.row = cur.row
		}
		if cur.col < min.col {
			min.col = cur.col
		}
	}

	for ri := 0; ri < citySize; ri++ {
		for ci := 0; ci < citySize; ci++ {
			if c.getMinValAt(ri, ci) == 0 {
				continue
			}
			if ri < min.row {
				min.row = ri
			}
			if ci < min.col {
				min.col = ci
			}
		}
	}

	var sb strings.Builder
	sb.WriteString("        ")
	for ci := min.col; ci < citySize; ci++ {
		// sb.WriteString(fmt.Sprintf("   %4d ", ci))
		sb.WriteString(fmt.Sprintf(" %3d", ci))
	}
	sb.WriteByte('\n')
	for ri := min.row; ri < citySize; ri++ {
		sb.WriteString(fmt.Sprintf("Row %3d:", ri))
		for ci := min.col; ci < citySize; ci++ {
			if pos.row == ri && pos.col == ci {
				sb.WriteByte('X')
			} else {
				sb.WriteByte(' ')
			}
			switch headings[ri][ci] {
			case east:
				sb.WriteByte('>')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case west:
				sb.WriteByte('<')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case north:
				sb.WriteByte('^')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			case south:
				sb.WriteByte('v')
				sb.WriteByte('0' + byte(numInDir[ri][ci]))
			default:
				sb.WriteByte(' ')
				sb.WriteByte(' ')
			}
			val := c.getMinValAt(ri, ci)
			var b byte
			if val == 0 {
				b = ' '
			} else {
				i := (val * len(spaceBytes)) / 1200
				if i >= len(spaceBytes) {
					i = len(spaceBytes) - 1
				}
				b = spaceBytes[i]
			}
			sb.WriteByte(b)
			/*
				if val == 0 {
					sb.WriteString("     ")
					// sb.WriteString("---- ")
				} else {
					sb.WriteString(fmt.Sprintf("%4d ", val))
				}
			*/
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *city) getPathHeatLoss(pos *position) int {
	return c.getPathHeatLossWithPrint(``, pos)
}
func (c *city) getPathHeatLossWithPrint(
	prefix string,
	pos *position,
) int {
	if pos == nil {
		return 0
	}
	total := 0
	switch pos.heading {
	case south:
		fmt.Printf("%sSOUTH: %d\n", prefix, pos.leftInDirection)
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.row+n >= 0 && pos.row+n < citySize {
				fmt.Printf("%s+ %d\n", prefix, c.blocks[pos.row+n][pos.col])
				total += c.blocks[pos.row+n][pos.col]
			}
		}
	case north:
		fmt.Printf("%sNORTH: %d\n", prefix, pos.leftInDirection)
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.row-n >= 0 && pos.row-n < citySize {
				fmt.Printf("%s+ %d\n", prefix, c.blocks[pos.row-n][pos.col])
				total += c.blocks[pos.row-n][pos.col]
			}
		}
	case east:
		fmt.Printf("%sEAST: %d\n", prefix, pos.leftInDirection)
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.col+n >= 0 && pos.col+n < citySize {
				fmt.Printf("%s+ %d\n", prefix, c.blocks[pos.row][pos.col+n])
				total += c.blocks[pos.row][pos.col+n]
			}
		}
	case west:
		fmt.Printf("%sWEST: %d\n", prefix, pos.leftInDirection)
		for n := 1; n <= int(pos.leftInDirection); n++ {
			if pos.col-n >= 0 && pos.col-n < citySize {
				fmt.Printf("%s+ %d\n", prefix, c.blocks[pos.row][pos.col-n])
				total += c.blocks[pos.row][pos.col-n]
			}
		}
	}
	prefix += ` `
	return total + c.getPathHeatLossWithPrint(prefix, pos.prev)
}
