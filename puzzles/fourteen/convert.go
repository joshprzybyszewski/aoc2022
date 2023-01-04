package fourteen

import (
	"strconv"
	"strings"
)

func getGrid(input string) (grid, error) {
	const newline = "\n"
	const del = " -> "
	const comma = ","
	g := newGrid()

	var err error
	var prev coord
	var deli, commai, x, y int

	for nli := strings.Index(input, newline); nli >= 0; nli = strings.Index(input, newline) {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		commai = strings.Index(input, comma)
		x, err = strconv.Atoi(input[:commai])
		if err != nil {
			return grid{}, err
		}
		deli = strings.Index(input, del)
		y, err = strconv.Atoi(input[commai+1 : deli])
		if err != nil {
			return grid{}, err
		}
		input = input[deli+len(del):]
		prev = coord{x: x, y: y}

		for {
			deli = strings.Index(input, del)
			nli = strings.Index(input, newline)

			if nli == 0 {
				break
			}

			commai = strings.Index(input, comma)

			x, err = strconv.Atoi(input[:commai])
			if err != nil {
				return grid{}, err
			}
			if deli == -1 || deli > nli {
				y, err = strconv.Atoi(input[commai+1 : nli])
				input = input[nli:]
			} else {
				y, err = strconv.Atoi(input[commai+1 : deli])
				input = input[deli+len(del):]
			}
			if err != nil {
				return grid{}, err
			}

			for {
				g.addRock(prev.x, prev.y)

				if prev.x < x {
					prev.x++
				} else if prev.x > x {
					prev.x--
				} else if prev.y < y {
					prev.y++
				} else if prev.y > y {
					prev.y--
				} else {
					break
				}
			}

		}
		input = input[nli+1:]
	}

	return g, nil
}
